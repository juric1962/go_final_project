
FROM golang:1.22

WORKDIR /app

ENV TODO_PASSWORD=123

ENV TODO_PORT=7540

ENV TODO_DB ="./scheduler.db"

EXPOSE ${TODO_PORT}

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /my_app

CMD ["/my_app"]