# Step 1 - Compile the binary source
FROM golang:alpine as builder

# Install git to allow go mod to install dev deps
RUN apk --update add git less openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Install Go modules. Go modules files are copied seperately from the rest of the source code to allow Docker to cache
# this layer seperatly, allowing us to skip `go mod download` unless the mod files have actually changed.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the remaining source code
COPY . /app
RUN go build -o monitor .

# Step 2 - Build a small runtime image
FROM scratch

COPY --from=builder /app/monitor /
ENTRYPOINT ["/monitor"]