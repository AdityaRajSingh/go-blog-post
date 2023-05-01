FROM golang:1-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app


# RUN apk update

# RUN apk --update add postgresql-client git gcc musl-dev tar curl

RUN go mod download

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
