FROM golang:1.17.3-alpine3.14 AS builder

WORKDIR /app
COPY ./ ./
RUN go mod download

RUN go build -o main


#2
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder app.env /app
COPY --from=builder /app/helper/email_templates/ ./helper/email_templates/

EXPOSE 8000

CMD [ "./main" ]

# FROM golang:1.17-alpine3.14

# WORKDIR /invoice-api

# COPY . .

# RUN go mod download


# RUN go build -o mainfile

# EXPOSE 8000

# CMD ["./mainfile"]
