FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o penweb .

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/penweb .
CMD ["./penweb"]
