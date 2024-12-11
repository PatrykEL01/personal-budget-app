FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY .env /app/.env

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o personal-budget

FROM gcr.io/distroless/base-debian10 AS runner

COPY --from=builder /app/personal-budget /personal-budget

COPY --from=builder /app/.env /app/.env

USER nonroot:nonroot

ENTRYPOINT ["/personal-budget"]

EXPOSE 8080