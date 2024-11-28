FROM golang:1.21-alpine

WORKDIR /app
COPY go.mod main.go ./
RUN go mod tidy
RUN GOOS=linux go build -o main .

EXPOSE 8080
CMD ["./main"]
