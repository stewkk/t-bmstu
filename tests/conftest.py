#!/usr/bin/env python3

import pathlib
import sys

import pytest

pytest_plugins = ['testsuite.pytest_plugin']


def pytest_addoption(parser):
    group = parser.getgroup('app service')
    group.addoption(
        '--app-service-port',
        help='Bind app service to this port (default is %(default)s)',
        default=8080,
        type=int,
    )


@pytest.fixture
async def app_service(
    ensure_daemon_started,
    # Service process holder
    app_service_scope,
    # Service dependencies
    mockserver,
):
    # Start service if not started yet
    await ensure_daemon_started(app_service_scope)


@pytest.fixture
async def app_client(
    create_service_client,
    app_service_baseurl,
    app_service,
):
    # Create service client instance
    return create_service_client(app_service_baseurl)


@pytest.fixture(scope='session')
def app_service_baseurl(pytestconfig):
    return f'http://localhost:{pytestconfig.option.app_service_port}/'


@pytest.fixture(scope='session')
def app_root():
    """Path to app service root."""
    return pathlib.Path(__file__).parent.parent


@pytest.fixture(scope='session')
async def app_service_scope(
    pytestconfig,
    create_daemon_scope,
    app_root,
    app_service_baseurl,
):
    async with create_daemon_scope(
        args=[
            str(app_root.joinpath('main')),
            '--port',
            str(pytestconfig.option.app_service_port),
        ],
        ping_url=app_service_baseurl + 'ping',
    ) as scope:
        yield scope
