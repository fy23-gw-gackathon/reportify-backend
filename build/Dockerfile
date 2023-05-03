FROM golang:1.20

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.43.0 \
    && go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN go mod tidy

CMD ["air", "-c", ".air.toml"]