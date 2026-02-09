# Use a multi-stage build: builder uses the official Go image, final image is distroless (non-root)
ARG GOVERSION=1.24.9
FROM golang:${GOVERSION} AS builder

WORKDIR /src

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build a statically-linked binary suitable for distroless
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -ldflags='-s -w' -o /todo-api ./cmd

### Final image: minimal attack surface, runs as non-root
FROM gcr.io/distroless/static:nonroot

# Copy binary from builder
COPY --from=builder /todo-api /todo-api

# The distroless static:nonroot image runs as a non-root user by default.
# Expose the port the app listens on (adjust if your app uses a different port).
EXPOSE 8080

ENTRYPOINT ["/todo-api"]
