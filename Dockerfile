FROM golang:1.21 AS builder
WORKDIR /source
COPY go.mod /source/go.mod
COPY main.go /source/main.go
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /fwd main.go

FROM scratch
COPY --from=builder /fwd /fwd
CMD ["/fwd"]
