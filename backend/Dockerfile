# BUILDER
FROM golang:latest as builder

LABEL maintainer="xavier@franquet.es"

COPY src/ /go/src/

WORKDIR /go/src
RUN go mod download
RUN go mod verify
RUN go build -o /go/bin/backend

FROM alpine as runner
RUN apk add --no-cache libc6-compat
RUN mkdir /go
COPY --chown=0:0 --from=builder /go/bin/backend /go
WORKDIR /go/conf

CMD ["/go/backend"]
