FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/server



FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

ENTRYPOINT [ "/app/app" ]