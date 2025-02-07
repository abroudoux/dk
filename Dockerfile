FROM golang:1.22.2-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /dk ./cmd/main.go

FROM golang:1.22.2-alpine

WORKDIR /app

COPY --from=builder /dk /dk

CMD ["/dk"]