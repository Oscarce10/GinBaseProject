FROM golang:1.19-bullseye

RUN apt update && apt upgrade -y
WORKDIR /go/src/app

ADD . .
COPY go.mod ./
COPY go.sum ./

RUN go mod download
WORKDIR /go/src/app/cmd
RUN go build -o main .
EXPOSE 8000
CMD ["./main"]
