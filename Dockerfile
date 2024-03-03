
FROM node:20-buster as build-frontend


WORKDIR /ui
COPY ui/package.json ui/pnpm-lock.yaml ./

RUN corepack enable

RUN pnpm install --frozen-lockfile

COPY ui/.env.production ui/.env ui/index.html ui/postcss.config.js ui/tailwind.config.js ui/tsconfig.json ui/tsconfig.node.json ui/vite.config.ts ./
COPY ui/src ./src

ARG MODE=production
RUN pnpm build --mode $MODE

FROM golang:1.22-bullseye as build-server
WORKDIR /api
COPY api/go.mod api/go.sum ./
RUN go mod download

COPY api/kafebar kafebar

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./kafebar/cmd

FROM alpine:3.15

WORKDIR /
COPY --from=build-server /api/server /usr/bin/
COPY --from=build-frontend /ui/dist /dist
COPY db/migrations /db/migrations

ENV UI_PATH /dist
ENV MIGRATIONS_DIRECTORY=/db/migrations

ENTRYPOINT ["server"]