FROM ghcr.io/fiskaly/docker.oapi-codegen:latest as oapi
FROM golang:bullseye

# Set current timezone
RUN echo "Europe/Moscow" > /etc/timezone
RUN ln -sf /usr/share/zoneinfo/Europe/Moscow /etc/localtime

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update

RUN apt-get install -y --allow-unauthenticated \
    binutils-dev \
    build-essential \
    curl \
    python3-dev \
    vim \
    sudo \
    wget \
    dirmngr \
    python3-pip


RUN apt-get clean all

RUN pip3 install yandex-taxi-testsuite[postgresql-binary]
COPY --from=oapi /oapi-codegen /usr/bin/oapi-codegen

EXPOSE 8080

ENV PATH /usr/sbin:/usr/bin:/sbin:/bin:${PATH}
