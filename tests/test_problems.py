#!/usr/bin/env python3

from bs4 import BeautifulSoup


async def test_api_get_problem(app_client, load_json):
    response = await app_client.get(
        '/api/problems/7a81f316-add0-4dc9-85a7-28b8fd747f8f',
    )
    assert response.status_code == 200
    assert response.json() == load_json(
        '7a81f316-add0-4dc9-85a7-28b8fd747f8f.json')


async def test_view_problem(app_client, load):
    response = await app_client.get(
        '/problems/7a81f316-add0-4dc9-85a7-28b8fd747f8f',
    )
    assert response.status_code == 200
    assert str(BeautifulSoup(
        response.content, features='html.parser').find(id='statement')).strip() == load('7a81f316-add0-4dc9-85a7-28b8fd747f8f.html')[:-1]
