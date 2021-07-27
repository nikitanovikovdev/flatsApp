FROM golang:latest
WORKDIR /goProjects/src/flatsApp
COPY ./ ./
RUN go build -o bin/main cmd/web/main.go
RUN go build -o bin/main cmd/users/main.go
CMD ["bin/main"]