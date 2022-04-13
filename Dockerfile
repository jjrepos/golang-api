# syntax=docker/dockerfile:1

FROM alpine:latest as builder
RUN apk --update add \
    go \
    musl-dev 

RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o lib .


FROM alpine:latest
RUN mkdir /api
COPY --from=builder /build/lib /api
EXPOSE 8080
WORKDIR /api
ENTRYPOINT [ "sh", "-c", "./lib" ]