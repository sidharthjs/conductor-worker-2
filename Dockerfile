FROM golang:alpine3.16 as builder
COPY go.mod go.sum main.go /worker/
COPY workers/ /worker/workers/
WORKDIR /worker
RUN CGO_ENABLED=0 GOOS=linux go build -a -o conductor-worker-2 .

FROM alpine:latest
COPY --from=builder /worker/conductor-worker-2 /app/conductor-worker
CMD ["/app/conductor-worker"]
