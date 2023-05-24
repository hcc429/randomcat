FROM golang:alpine AS builder 

WORKDIR /app
COPY . .
RUN go build -o app 

FROM alpine AS production 
ENV GIN_MODE=release
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
