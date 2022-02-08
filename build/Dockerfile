FROM golang:latest AS builder
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /shortener ./cmd/api/url-shortener

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /shortener ./
RUN chmod +x ./shortener
ENTRYPOINT ["./shortener"]
EXPOSE 8080