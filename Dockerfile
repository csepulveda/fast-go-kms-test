FROM golang:alpine3.15 AS builder
WORKDIR /app
COPY ./. ./
RUN --mount=type=cache,target=/root/.cache/ --mount=type=cache,target=/go/pkg/ go build -ldflags "-s -w" -o /runner ./main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /runner ./
CMD ["./runner"]  