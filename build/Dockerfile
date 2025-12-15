FROM golang:1.25.4 AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o httpbin ./cmd/httpbin/main.go

FROM alpine:3.21.3
RUN apk update && apk add --no-cache curl --no-interactive && adduser -D -u 1001 appuser
COPY --from=build /app/httpbin /usr/local/bin/httpbin
USER appuser
CMD [ "httpbin" ]
EXPOSE 8888
