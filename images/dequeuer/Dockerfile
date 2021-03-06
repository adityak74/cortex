# Build the manager binary
FROM golang:1.15 as builder

# Copy the Go Modules manifests
COPY go.mod go.sum /workspace/
WORKDIR /workspace
RUN go mod download

COPY pkg/config pkg/config
COPY pkg/consts pkg/consts
COPY pkg/lib pkg/lib
COPY pkg/dequeuer pkg/dequeuer
COPY pkg/probe pkg/probe
COPY pkg/types pkg/types
COPY pkg/crds pkg/crds
COPY pkg/workloads pkg/workloads
COPY cmd/dequeuer cmd/dequeuer

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o dequeuer ./cmd/dequeuer

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/dequeuer .
USER nonroot:nonroot

ENTRYPOINT ["/dequeuer"]
