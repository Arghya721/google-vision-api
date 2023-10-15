FROM golang:1.20-alpine AS builder

# Set the current working directory.
RUN mkdir -p /google-vision
WORKDIR /google-vision

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download
COPY . .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o google-vision app/main.go

# Build the final image
FROM alpine:latest

# Set the current working directory.
WORKDIR /google-vision
COPY --from=builder /google-vision /google-vision

EXPOSE 1323

# Command to run the executable
CMD ["./google-vision"]