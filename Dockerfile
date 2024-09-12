FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/kish1n/KhOn
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/KhOn /go/src/github.com/kish1n/KhOn


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/KhOn /usr/local/bin/KhOn
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["KhOn"]
