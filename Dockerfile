FROM golang:1.19.9 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./out/albion_killboard ./cmd/main.go

# Define the command to run your Go application
CMD ["./out/albion_killboard"]