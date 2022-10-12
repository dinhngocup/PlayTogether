FROM golang:1.19 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main.go

# This container exposes port 8088 to the outside world
EXPOSE 8088

# Run the executable
CMD ["main"]
