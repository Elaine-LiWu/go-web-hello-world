FROM golang:alpine

LABEL maintainer="li.wu@aliyun.com"

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download


RUN ls -la

COPY *.go ./

RUN go build -v -o go-web-hello-world go-web-hello-world.go
EXPOSE 28082

ENTRYPOINT ["./go-web-hello-world"]
