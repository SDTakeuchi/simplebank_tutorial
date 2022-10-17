# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

RUN chmod +x /app/start.sh

EXPOSE 8080
# when CMD is used with ENTRYPOINT,
# it will act as an additional params
# in this case, this is equivalent to ENTRYPOINT [ "/app/start.sh", "/app/main" ]
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
