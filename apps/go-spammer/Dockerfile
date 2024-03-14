FROM golang:1.22-alpine

WORKDIR /app

# ./apps/go-spammer/go.sum
COPY ./apps/go-spammer/go.mod ./

RUN go mod download

COPY ./apps/go-spammer .

RUN go build -o go-spammer

CMD ["./go-spammer"]