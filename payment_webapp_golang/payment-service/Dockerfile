FROM golang:1.23-alpine

WORKDIR /app

# Copy only the Go module files first
COPY go.mod .
RUN go mod download

# copy the application code
COPY main.go .

# Build the go binary
RUN go build -o payment-service .

# Expose the service port
EXPOSE 8082

# Run the application
CMD ["./payment-service"]