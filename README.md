[![Docker Pulls](https://img.shields.io/docker/pulls/michelfortes/httpbin.svg)](https://hub.docker.com/r/michelfortes/httpbin)

# What is it?

HTTPBin is a server that helps you debug HTTP requests.

# Environment variables

This application expects to receive the following environment variables:
- HTTPBIN_SERVICE_ID (default: "")
- HTTPBIN_SERVICE_PORT (default: "8888")

# Features

You can determine the server's behavior for a specific request-response cycle by sending predefined configuration headers, as follows:

### _Set the status code_

Header: X-HttpBin-Status

Description: Specifies the _status code_ of the response.

### _Set a response delay_

Header: X-HttpBin-Sleep

Description: Specifies a response delay in seconds.

### _Restrict the content-type accepted by the server_

Header: X-HttpBin-Content-Type

Description: Requests with payloads must declare a content type that matches the one accepted by the server, otherwise the server will return a 415 (Unsupported Media Type) status code.

## Special Features

### Proxying request

Use the special path "/proxy" with the query param "to" to set the destination to which the request should be proxied. The response replicates the headers and body the destination server returns.

```
curl -i "http://localhost:8888/proxy?to=https://google.com"
```

### Health Checks

Use the special path "/health" to check service health.
It returns a simple 200 status code when the service is up and ready.

```
curl -i "http://localhost:8888/health"
```


# Running

## Using go run command

```bash
go run cmd/httpbin/main.go
```

## Using docker compose

```bash
docker compose -f deployments/docker-compose.yml up --build
```

# Output Example

A JSON containing information about the request and the Container where the application is running.

Conteúdo do arquivo `payload.json`:

```json
{
    "name": "John Doe",
    "email": "john.doe@example.com"
}
```

```bash
curl -i -H "Content-Type: application/json" -X POST "localhost:8888/some-path?fruits=apple&fruits=orange&drink=juice" -d @payload.json
```
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 09 Jan 2024 13:41:57 GMT
Content-Length: 340
```
```json
{
  "serviceId": "container-01",
  "remoteAddr": "127.0.0.1:57968",
  "method": "POST",
  "path": "/some-path",
  "queryParams": {
    "drink": [
      "juice"
    ],
    "fruits": [
      "apple",
      "orange"
    ]
  },
  "headers": {
    "Accept": [
      "*/*"
    ],
    "Content-Length": [
      "60"
    ],
    "Content-Type": [
      "application/json"
    ],
    "User-Agent": [
      "curl/8.7.1"
    ]
  },
  "payload": "{    \"name\": \"Jhon Doe\",    \"email\": \"john.doe@example.com\"}"
}
