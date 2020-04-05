FROM golang:1.14-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get	github.com/dgrijalva/jwt-go
RUN go get	github.com/gorilla/handlers
RUN go get	github.com/gorilla/mux
RUN go get	github.com/mattn/go-sqlite3
RUN go build -o main .
CMD ["/app/main"]