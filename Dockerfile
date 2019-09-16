# build step
FROM golang:1.12 as builder

LABEL maintainer="sunnydog0826@gmail.com"
COPY . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

# run step
FROM alpine

RUN apk update \
    && apk add --no-cache ca-certificates \
    && rm -rf /var/cache/apk/*

# copy bin from build step
COPY --from=builder /build/drone-oss /bin/

ENTRYPOINT ["/bin/drone-oss"]