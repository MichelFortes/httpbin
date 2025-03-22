FROM golang:1.24.1 AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o httpbin ./cmd/httpbin/main.go

FROM alpine
RUN apk update
RUN apk upgrade -U --no-interactive
RUN apk add curl --no-interactive
COPY --from=build /app/httpbin /usr/local/bin/httpbin
CMD [ "httpbin" ]
EXPOSE 8888
