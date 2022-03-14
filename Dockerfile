FROM golang:latest as builder
WORKDIR /app
COPY ./ /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./tnYT ./cmd/main.go

FROM scratch
COPY --from=builder /app/tnYT /usr/bin/tnYT


ENTRYPOINT ["/usr/bin/tnYT"]
EXPOSE 50051