# Justificativa

Esta aplicação foi desenvolvida para debugar requisições HTTP.

# Variáveis de ambiente

Esta aplicação espera receber as seguintes variáveis de ambiente:
- HTTPBIN_SERVICE_ID (default: "")
- HTTPBIN_PORT (default: "8888")

# Path /

## Features

### _Determine o status code_

Podemos determinar o _status code_ da resposta através do query param **_response_status_**.

### _Determine uma demora na respota_

Podemos determinar um delay, em segundos, na resposta através do query param **_sleep_** (em segundos).

## Output

Um JSON contendo o informações sobre a requisição e o Container onde a aplicação está sendo executada.
```
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

# Path /proxy


Usado para determinar o destino para o qual a requisição será proxiada.

```
curl -i "localhost:8888/proxy?to=http://www.google.com"
```

## output 
Replica os headers e body do destino proxiado 