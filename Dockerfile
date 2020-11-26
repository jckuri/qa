# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.15 as builder

# Add Maintainer Info
LABEL maintainer="Juan Carlos Kuri <juan.kuri@bairesdev.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o qaserver cmd/server/main.go


######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/qaserver .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./qaserver"]
