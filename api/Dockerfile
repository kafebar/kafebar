## Build
FROM golang:1.22-bullseye as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY src src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main work/cmd/server/main.go

## Deploy
FROM alpine:3.15
WORKDIR /
COPY --from=build /app/main /usr/bin/
ENTRYPOINT ["main"]
