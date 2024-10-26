FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/myapp .

FROM scratch

COPY --from=builder /app/myapp /myapp

EXPOSE 8080

CMD ["/myapp"]