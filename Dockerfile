FROM golang:1.19.9 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o albion_killboard .

# Create a new lightweight image for deployment
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the build image
COPY --from=build /app/albion_killboard .

# Expose the port your Go application listens on
EXPOSE 8080

# Define the command to run your Go application
CMD ["./albion_killboard"]