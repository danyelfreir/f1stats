# syntax=docker/dockerfile:1

FROM golang:1.22
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o /f1stats
EXPOSE 8080
CMD ["/f1-stats-service"]
