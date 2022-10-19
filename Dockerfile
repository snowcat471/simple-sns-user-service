# Build Stage
FROM golang:1.19.2-alpine3.16 AS builder

WORKDIR /app
COPY . .

ENV GO111MODULE=on

RUN go get 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o app .

# Running Stage
FROM alpine

WORKDIR /app
COPY --from=builder /app .

EXPOSE 3000

ENTRYPOINT [ "./app" ]