FROM golang:1.25
EXPOSE 8080
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /books

CMD ["/books"]