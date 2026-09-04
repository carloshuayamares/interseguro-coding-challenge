FROM golang:alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /matrix-service .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /matrix-service /matrix-service
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/matrix-service"]
