# REST Overview

## What is REST?
- REpresentational State Transfer
- Stateless client-service communication architecture
- Communicate resources over the internet
- CRUD endpoints (Create, Read, Update, Delete)
- HTTP as protocol

## HTTP
- HTTP Requests
  - Method
    - POST, GET, PUT, DELETE... (CRUD)
  - URL - `https://www.golangdojo.com:80/?key=value`
    - HTTP/HTTPS - `https://`
    - Host - `www.golangdojo.com`
    - Port - `:80`(default/omitted)
    - Query parameters - `/?key=value`
  - Header
    - Accept: text/html (or application/json, etc.)
    - Accept-Language: en-US
    - User-Agent: Mozilla/5.0
  - Body
    - JSON
- HTTP Responses
  - Status Codes
    - 1XX Informational
    - 2XX Success
    - 3XX Redirection
    - 4XX Client error
    - 5XX Server error
  - HTTP version
    - HTTP 1.0
    - HTTP 2.0
    - ...
  - Header
    - Content-Type: text/html; charset=UTF-8
    - Content-Length: 
  - Body (interpreted based on "Accept" header in the request)
    - HTML
    - JSON
    - ...

## Package http