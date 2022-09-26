#build stage
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN make build

#final stage
FROM alpine:3.10
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/bin /app
RUN chmod +x /app/demo
ENTRYPOINT /app/demo
EXPOSE 8005