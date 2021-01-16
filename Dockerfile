FROM golang:1.15

WORKDIR /go/src/
COPY . .
RUN go build main.go
CMD ["./main"]