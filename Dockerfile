FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/ExerciseDiary/ && CGO_ENABLED=0 go build -o /ExerciseDiary .


FROM scratch

WORKDIR /data/ExerciseDiary
WORKDIR /app

COPY --from=builder /ExerciseDiary /app/

ENTRYPOINT ["./ExerciseDiary"]