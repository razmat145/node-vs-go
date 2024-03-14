FROM golang:1.22-alpine

WORKDIR /app

COPY ./apps/go-echo/go.mod ./apps/go-echo/go.sum ./

RUN go mod download

COPY ./apps/go-echo .

RUN go build -o go-echo

CMD ["./go-echo"]