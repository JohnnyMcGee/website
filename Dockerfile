FROM golang:1.23.4 AS builder
WORKDIR /app
COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -tags musl -o server .

FROM scratch

WORKDIR /app
COPY --from=builder /app/server /app/server
COPY --from=builder /app/public /app/public

EXPOSE 3000

ENTRYPOINT ["/app/server"]