FROM golang:1.12-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

EXPOSE 38080

CMD ["./main"]