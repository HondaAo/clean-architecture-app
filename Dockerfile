FROM golang:1.18

ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app
RUN go mod tidy

COPY . /app

RUN go install github.com/cosmtrek/air@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

COPY ./start.sh /app/start.sh
RUN chmod +x /app/start.sh

CMD ["air","-c", ".air.toml"]