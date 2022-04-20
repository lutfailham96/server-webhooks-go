FROM ubuntu:focal

ENV PATH=$PATH:/opt/go/bin \
    GOPATH=/opt/go/bin \
    GOLANG_VERSION=1.18.1

WORKDIR /usr/src/app

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get autoremove -y --purge \
    && apt-get clean -y

COPY ./go /opt/go
