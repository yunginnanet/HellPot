FROM golang:1.23 AS build
WORKDIR /go/src/app

COPY go.* .
RUN go mod download

COPY . .

RUN go vet -v ./...
RUN go test -v ./...
RUN \
    CGO_ENABLED=0 \
    VERSION=`git tag --sort=-version:refname | head -n 1` \
    go build -trimpath \
    -ldflags "-s -w -X main.version=$VERSION" \
    cmd/HellPot/*.go


FROM gcr.io/distroless/static-debian11
LABEL org.opencontainers.image.source https://github.com/yunginnanet/HellPot

COPY --from=build /go/src/app/HellPot /app
COPY --from=build /go/src/app/docker_config.toml /config
EXPOSE 8080
ENTRYPOINT ["/app", "-c", "/config"]
