FROM golang:alpine

ENV CGO_ENABLED=0

RUN go install github.com/cortesi/modd/cmd/modd@latest

WORKDIR /code

COPY go.mod go.sum ./
RUN go mod download -x

STOPSIGNAL SIGKILL

CMD [ "modd" ]
