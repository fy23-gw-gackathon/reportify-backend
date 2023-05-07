FROM golang:1.20

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.43.0

ADD go.mod go.sum ./

RUN go mod download

ADD . .

CMD ["air", "-c", "./worker/.air.toml"]