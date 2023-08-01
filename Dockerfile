# Get Go image from DockerHub.
FROM golang:1.21rc3-alpine3.18 AS api

# Set working directory.
WORKDIR /app

# Copy dependency locks so we can cache.
COPY go.mod go.sum ./

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Run application and expose port 8080.
EXPOSE 8080

CMD ["go", "run", "main.go"]
