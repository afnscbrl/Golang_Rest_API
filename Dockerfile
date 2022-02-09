FROM golang:alpine as builder

RUN mkdir /app

COPY . /app

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/* ./
# COPY --from=builder /app/templates ./templates
# COPY --from=builder /app/pkg ./pkg
# COPY --from=builder /app/data/basemigrations ./migrations

EXPOSE 8000

CMD ["/app/main"]