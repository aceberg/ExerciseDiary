FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/ExerciseDiary/ && CGO_ENABLED=0 go build -o /ExerciseDiary .


FROM alpine

WORKDIR /app

RUN apk add --no-cache arp-scan tzdata \
    && mkdir /data/ExerciseDiary

COPY --from=builder /ExerciseDiary /app/

ENTRYPOINT ["./ExerciseDiary"]