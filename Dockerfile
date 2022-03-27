FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Assign working directory to /usr/src/app
WORKDIR /usr/src/app

# Copy and download dependency using go mod
COPY go.mod go.sum ./

# RUN go mod download
RUN go mod download && go mod verify

# Copy the code into the container to build
COPY . .

RUN go build -v -o main .

# Build a small image
FROM golang:alpine

COPY --from=builder /usr/src/app/main /

# Expose port 8080 to host machine
EXPOSE 8080 

# Command to run application
ENTRYPOINT ["/main"]
#CMD ["/main"]