FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o upassed-submission-service ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir -p /upassed-submission-service/config
RUN mkdir -p /upassed-submission-service/migration/scripts
COPY --from=builder /app/upassed-submission-service /upassed-submission-service/upassed-submission-service
COPY --from=builder /app/config/* /upassed-submission-service/config
COPY --from=builder /app/migration/scripts/* /upassed-submission-service/migration/scripts
RUN chmod +x /upassed-submission-service/upassed-submission-service
ENV APP_CONFIG_PATH="/upassed-submission-service/config/local.yml"
EXPOSE 44048
CMD ["/upassed-submission-service/upassed-submission-service"]
