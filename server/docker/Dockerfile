FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build cmd/api/main.go" --command=./main
