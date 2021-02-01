FROM golang:1.15

RUN go get -u -d github.com/golang-migrate/migrate/cmd/migrate
RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

WORKDIR /go/src/
COPY . .
RUN go build main.go
EXPOSE 3000
CMD ["./main"]