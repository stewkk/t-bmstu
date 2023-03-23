#!/usr/bin/env python3

async def test_ping(app_client):
    response = await app_client.get(
        '/ping',
    )
    assert response.status_code == 200
