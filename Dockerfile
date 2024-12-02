FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o upassed-answer-service ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir -p /upassed-answer-service/config
RUN mkdir -p /upassed-answer-service/migration/scripts
COPY --from=builder /app/upassed-answer-service /upassed-answer-service/upassed-answer-service
COPY --from=builder /app/config/* /upassed-answer-service/config
COPY --from=builder /app/migration/scripts/* /upassed-answer-service/migration/scripts
RUN chmod +x /upassed-answer-service/upassed-answer-service
ENV APP_CONFIG_PATH="/upassed-answer-service/config/local.yml"
EXPOSE 44048
CMD ["/upassed-answer-service/upassed-answer-service"]
