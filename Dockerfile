FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod .
RUN go mod download

COPY . .
COPY /sites.csv .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cli-healthcheck ./cmd

FROM scratch
WORKDIR /app
COPY --from=builder /app/cli-healthcheck .

ENTRYPOINT ["./cli-healthcheck"]