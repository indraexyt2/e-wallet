FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .
COPY .env .

RUN go build -o e-wallet-notification
RUN chmod +x e-wallet-notification
EXPOSE 7003
CMD [ "./e-wallet-notification" ]