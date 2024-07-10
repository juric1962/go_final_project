
FROM golang:1.21.6

WORKDIR /app

ENV TODO_PASSWORD=123

ENV TODO_PORT=7540

ENV TODO_DB="./scheduler.db"

ENV CGO_ENABLED=0

ENV GOOS=linux

ENV GOARCH=amd64

EXPOSE ${TODO_PORT}

COPY . .

RUN go mod download

RUN   go build -o /todo_app

CMD ["/todo_app"]
