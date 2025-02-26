# syntax=docker/dockerfile:1

FROM golang:1.23.3-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o ./api.exe ./cmd/api

EXPOSE 8080

# Run the application
CMD ["./api.exe"]
