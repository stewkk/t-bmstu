#!/usr/bin/env sh

git clone --branch master --single-branch https://github.com/stewkk/t-bmstu.git \
    && cp t-bmstu/deploy/deploy.sh deploy.sh \
    && sudo mkdir -p /var/swag/nginx/proxy-confs \
    && sudo chown -R $USER /var/swag/nginx \
    && cp t-bmstu/deploy/t-bmstu.subdomain.conf /var/swag/nginx/proxy-confs/ \
    && docker compose -f t-bmstu/deploy/docker-compose.yaml up -d
