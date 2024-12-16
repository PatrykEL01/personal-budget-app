FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY .env /app/.env

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o personal-budget

FROM golang:1.23-alpine AS runner

RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

COPY --from=builder /app /app

COPY --from=builder /app/personal-budget /personal-budget

COPY --from=builder /app/.env /app/.env

RUN chown -R nonroot:nonroot /app /app/personal-budget /personal-budget


USER nonroot:nonroot

ENTRYPOINT ["/personal-budget"]

EXPOSE 8080
