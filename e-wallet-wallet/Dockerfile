FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .
COPY .env .

RUN go build -o e-wallet-wallet
RUN chmod +x e-wallet-wallet

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/e-wallet-wallet .
COPY --from=builder /app/.env .

EXPOSE 8081

CMD [ "./e-wallet-wallet" ]