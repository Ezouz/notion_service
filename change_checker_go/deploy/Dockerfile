FROM golang:1.19.1 AS builder

ENV CGO_ENABLED off

COPY ./change_checker_go /go/src/gitlab.42paris.fr/notion_service
COPY .env /go/src/gitlab.42paris.fr/notion_service/.env
WORKDIR /go/src/gitlab.42paris.fr/notion_service

RUN go build -o /go/src/gitlab.42paris.fr/notion_service/tmp/change_checker /go/src/gitlab.42paris.fr/notion_service/cmd/db_change_checker/main.go

FROM debian:11-slim

COPY --from=builder /go/src/gitlab.42paris.fr/notion_service/tmp/change_checker /usr/local/bin/change_checker
COPY --from=builder /go/src/gitlab.42paris.fr/notion_service/.env .env

RUN apt-get update && apt-get install -y ca-certificates

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

CMD sh -c "/wait && /usr/local/bin/change_checker"