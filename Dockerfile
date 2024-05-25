# Use the official Golang image to create a build artifact.
FROM golang:1.17 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the binary.
RUN go build -o main .

# Use a minimal base image to package the binary.
FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/main /app/
WORKDIR /app
CMD ["./main"]
