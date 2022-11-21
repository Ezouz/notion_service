FROM golang:1.19.1 AS builder

ENV CGO_ENABLED off

COPY ./change_checker_go /go/src/gitlab.42paris.fr/notion_service
COPY .env /go/src/gitlab.42paris.fr/notion_service/.env
WORKDIR /go/src/gitlab.42paris.fr/notion_service

RUN go build -o /go/src/gitlab.42paris.fr/notion_service/tmp/api /go/src/gitlab.42paris.fr/notion_service/.

FROM debian:11-slim

COPY --from=builder /go/src/gitlab.42paris.fr/notion_service/tmp/api /usr/local/bin/api
COPY --from=builder /go/src/gitlab.42paris.fr/notion_service/.env .env

RUN apt-get update && apt-get install -y ca-certificates

CMD /usr/local/bin/api