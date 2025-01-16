FROM golang:1.23.4 AS builder

# TODO: make these variables dynamic depending on the deployment target
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

WORKDIR /app
 
COPY . .
 
RUN go mod download

RUN go build -o server
 
FROM alpine:3.21 AS production

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /server
COPY --from=builder /app/public /public

EXPOSE 3000

CMD ["./server"]