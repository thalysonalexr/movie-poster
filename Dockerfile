FROM golang:1.15 AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch
COPY --from=builder /app/movie-poster /app/
EXPOSE 8080
ENTRYPOINT ["/app/movie-poster"]
