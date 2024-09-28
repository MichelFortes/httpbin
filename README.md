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

Description: Determine the _status code_ of the response.

### _Set a response delay_

Header: X-HttpBin-Sleep

Description: Determine a response delay in seconds.

## Proxing request

Header: X-HttpBin-Proxy-To

Description: Use the special path "/proxy" to set the destination to which the request should be proxied. The response replicates the headers and body the destination server return.

```
curl -i -H "X-HttpBin-Proxy-To: https://google.com" "localhost:8888/proxy"
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

```bash
curl -i "localhost:8888/some-path?products=notebook&products=tablet&customer=john"
```
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 09 Jan 2024 13:41:57 GMT
Content-Length: 267
```
```json
{
  "serviceId": "container-01",
  "remoteAddr": "127.0.0.1:57968",
  "method": "GET",
  "path": "/some-path",
  "queryParams": {
    "customer": [
      "john"
    ],
    "products": [
      "notebook",
      "tablet"
    ]
  },
  "headers": {
    "Accept": [
      "*/*"
    ],
    "User-Agent": [
      "curl/7.81.0"
    ]
  }
}

```
