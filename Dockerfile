FROM golang:1.21.6 as builder

WORKDIR /app
COPY go.mod ./

RUN go mod download

COPY . .
RUN go build -o main .
FROM busybox:latest  
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
