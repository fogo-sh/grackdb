FROM node:16 AS builder

WORKDIR /app
COPY ./frontend .
RUN npm install && npm run build

FROM golang:1.16-buster

WORKDIR /build
COPY --from=builder /app/dist .
COPY . .
RUN go mod download &&\
    go mod verify &&\
    go build -o grackdb ./cmd/server

ENTRYPOINT ["/build/grackdb"]
