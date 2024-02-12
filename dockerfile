FROM golang:alpine

WORKDIR /ascii

COPY . ./

RUN go build -o dockerize

RUN apk add bash

EXPOSE 8080

CMD [ "./dockerize" ]