FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o urlshortener .

EXPOSE 8080

CMD ["./urlshortener"]