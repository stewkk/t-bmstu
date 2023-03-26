#!/usr/bin/env python3

from bs4 import BeautifulSoup


async def test_api_get_problem(app_client):
    response = await app_client.get(
        '/api/problems/9aa25460-c0f6-4a5c-aba4-ef59bc22e3a2',
    )
    assert response.status_code == 200
    assert response.json() == {'statement': 'TODO'}

    response = await app_client.get(
        '/api/problems/wrong-id'
    )
    body = response.json()

    assert response.status_code == 400
    assert body['message'] is not None
    assert body['code'] == 400


async def test_view_problem(app_client):
    response = await app_client.get(
        '/problems/9aa25460-c0f6-4a5c-aba4-ef59bc22e3a2',
    )
    assert response.status_code == 200
    assert BeautifulSoup(
        response.content, features='html.parser').text == 'TODO'

    response = await app_client.get(
        '/problems/wrong-id'
    )
    assert response.status_code == 400
    body = BeautifulSoup(response.content, features='html.parser').text
    assert body is not None and body != ""
