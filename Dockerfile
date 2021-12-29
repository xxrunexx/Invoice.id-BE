FROM golang:1.17-alpine3.14

WORKDIR /invoice-api

COPY . .

RUN go mod download


RUN go build -o mainfile

EXPOSE 8000

CMD ["./mainfile"]