## Build api
FROM golang:1.21-bullseye as server
WORKDIR /app
COPY api/go.mod ./
RUN go mod download

COPY api/src src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main src/main.go

## Build ui
FROM node:20-slim as client
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app
COPY ui/package.json ui/pnpm-lock.yaml  ./
RUN pnpm install

COPY ui/index.html ui/tsconfig.json  ui/vite.config.ts ui/uno.config.ts ./
COPY ui/src src

RUN pnpm build


## Deploy
FROM alpine:3.15
WORKDIR /
COPY --from=server /app/main /usr/bin/
COPY --from=client /app/dist /frontend
ENV "UI_PATH" "/frontend"
ENTRYPOINT ["main"]
