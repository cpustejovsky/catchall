# syntax=docker/dockerfile:1
FROM golang:1.17

COPY . /catchall
WORKDIR /catchall
RUN go mod download

WORKDIR /catchall/app/v1
RUN go build -o catchall

CMD [ "./catchall" ]