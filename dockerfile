FROM golang:1.20-alpine3.17 as builder

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o mydocu ./cmd/MyDocu/main.go


FROM alpine:3.17

WORKDIR /app
COPY --from=builder /app/mydocu .
COPY --from=builder /app/config.json .
EXPOSE 5000
ENTRYPOINT ["/app/mydocu"]
#CMD [ "config.json" ]