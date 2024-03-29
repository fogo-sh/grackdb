FROM node:16 AS frontend-builder

WORKDIR /app
COPY ./frontend .
RUN npm install && npm run build

FROM golang:1.18-bullseye AS go-builder

WORKDIR /build
COPY --from=frontend-builder /app/dist ./frontend/dist
COPY . .
RUN go mod download &&\
    go mod verify &&\
    go build -o grackdb ./cmd/server

FROM debian:bullseye-slim

WORKDIR /app
COPY --from=frontend-builder /app/dist ./frontend/dist
COPY --from=go-builder /build/grackdb ./grackdb
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/grackdb"]
