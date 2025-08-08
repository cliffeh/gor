FROM golang:1.24.2 AS builder

WORKDIR /app
COPY go.mod ./

ENV CGO_ENABLED=0
RUN go mod download

COPY . .
RUN go build -o server main.go

# Copy the server binary into a distroless container
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app/server /

EXPOSE 8080

USER nonroot

CMD ["/server"]
