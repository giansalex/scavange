FROM golang:alpine3.12
WORKDIR /app

RUN apk update && \
    apk add \
    build-base \
    gcc && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates

# Copy the local package files to the container's workspace.
COPY . .

# Build executable
RUN make
