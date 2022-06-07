FROM golang:1.18

ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app
RUN go mod tidy

COPY ./ /app

# binary will be $(go env GOPATH)/bin/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]