FROM golang:alpine AS builder
WORKDIR /snooper
COPY / ./
RUN go build

FROM alpine:latest
WORKDIR /snooper
COPY --from=builder /snooper/snooper ./snooper
ENTRYPOINT [ "./snooper" ]