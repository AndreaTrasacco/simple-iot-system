FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /measurementsaggregator ./cmd/api

EXPOSE 8081
CMD ["/measurementsaggregator"]