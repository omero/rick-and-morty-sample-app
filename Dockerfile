FROM golang:alpine
WORKDIR /go/src/app
COPY . .

ENV PORT=8080
CMD ["go", "run", "main.go"]