FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o server ./server

EXPOSE 50051

CMD ["./server"]
