# docker build -f Containerfile --progress=plain --no-cache -t go-practice .
# docker run --network=host go-practice

# Build stage
FROM jetpackio/devbox:latest AS builder

WORKDIR /devbox

# Copy the source code
COPY . .

USER root
RUN chown -R devbox .
USER devbox

# Ensure proper permissions for devbox
RUN whoami

RUN devbox install

# Build the application using devbox
RUN devbox run -- go build -o app ./cmd/main.go

# Runtime stage
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

WORKDIR /devbox

# Copy the compiled binary from builder
COPY --from=builder /devbox/app .

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./app"]
