FROM golang:1.16 as base

FROM base as dev

WORKDIR /opt/app/api

CMD ["CHESS"]
