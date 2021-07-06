FROM golang:latest
WORKDIR /goProjects/src/flatsApp
COPY . .
RUN go build -o bin/main cmd/web/main.go
CMD ["bin/main"]