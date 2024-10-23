# Find Index Service

Find Index REST API service searches its data for a given value and returns the index of that value and the value.

If the target value is not found, but there is a value within a tolerance of 10% of the target,
it returns that value and its index.

If there is no value within 10% target value, it returns NotFound error.


## Install

    git clone https://github.com/marcin-mc/find-index
    go mod tidy


Install golangci-lint for linting:
`https://golangci-lint.run/welcome/install/`

## Setup
Set port and logging level in config.yaml file:

    port:
      3000
    log_level:
      DEBUG



## Linting:
    make lint

## Testing:
    make test


## Usage 

    make run

### Request

`GET /endpoint/<target:int>`

    curl -i -H 'Accept: application/json' http://localhost:3000/endpoint/100

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Wed, 23 Oct 2024 11:25:58 GMT
    Content-Length: 60
    Connection: close

    {
        "Target": 100,
        "Index": 1,
        "Value": 100,
        "Message": "Value found"
    }


### Error Response (Not Found)

    HTTP/1.1 404 Not Found
    Content-Type: application/json; charset=utf-8
    Date: Wed, 23 Oct 2024 11:27:19 GMT
    Content-Length: 39
    Connection: close

    {
        "Message": "Value not found: 20000000"
    }

### Error Response (Bad Request)

    HTTP/1.1 400 Bad Request
    Content-Type: application/json; charset=utf-8
    Date: Wed, 23 Oct 2024 11:28:04 GMT
    Content-Length: 36
    Connection: close

    {
        "Message": "Bad input value: 'abc'"
    }
