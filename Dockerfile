FROM golang:1.18 as build

WORKDIR /app

ENV CGO_ENABLED 0
ENV GOOS linux

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -installsuffix cgo -o oryxtuiviewer oryxtuiviewer.go

FROM alpine:3.14

WORKDIR /src
COPY --from=build /app/oryxtuiviewer .

ENTRYPOINT ["./oryxtuiviewer"]
