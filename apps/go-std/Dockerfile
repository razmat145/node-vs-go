FROM golang:1.22-alpine

WORKDIR /app

COPY ./apps/go-std/go.mod ./apps/go-std/go.sum ./

RUN go mod download

COPY ./apps/go-std .

RUN go build -o go-std

CMD ["./go-std"]