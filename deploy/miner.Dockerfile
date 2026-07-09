FROM golang:alpine AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/

RUN GOOS=linux go build -v -o miner ./cmd/miner

FROM alpine:3.22.1
WORKDIR /app
RUN touch .env
RUN apk add ca-certificates
COPY --from=build /app/miner miner
CMD ["/app/miner"]
