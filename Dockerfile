FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/AppTemplate/ && CGO_ENABLED=0 go build -o /AppTemplate .


FROM alpine

WORKDIR /app

RUN apk add --no-cache arp-scan tzdata \
    && mkdir /data/AppTemplate

COPY --from=builder /AppTemplate /app/

ENTRYPOINT ["./AppTemplate"]