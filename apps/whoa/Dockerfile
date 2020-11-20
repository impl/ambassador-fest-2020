FROM golang:1.13 as builder

WORKDIR /workspace
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o whoa main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /workspace/whoa .

ENTRYPOINT ["/whoa"]
