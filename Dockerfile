FROM golang:1.25
EXPOSE 8080
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /books

CMD ["/books"]