FROM golang:1.21-alpine as build
WORKDIR /app
COPY go.mod server.go ./
COPY handlers ./handlers/
RUN go mod tidy
RUN go build -o httpbin-go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/httpbin-go ./
RUN apk update && apk upgrade -U --no-interactive
CMD [ "./httpbin-go" ]
EXPOSE 8888