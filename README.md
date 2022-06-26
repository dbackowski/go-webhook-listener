# Webhook listener

Very simple webhook listener written in go.

It starts the HTTP server on the 8080 port and will print received headers and JSON payload.

## Usage

* clone the repo
* go to cloned repo directory and run:

```
go build main.go
./webhook-listener
```

* next using curl send some test webhook
```sh
curl -k -X POST -H 'Content-Type: application/json' -H 'APIKEY: test' -d '{
    "data": [
        {
            "id": 1,
            "status": 2,
            "test": {
                "id": 1,
                "description": "test"
            }
        }
    ]
}' http://localhost:8080
```

* you should see received webhook in the output:

```
2022/06/19 17:21:38 server started on port: 8080
HEADERS:
--------
Content-Length: 188
User-Agent: curl/7.79.1
Accept: */*
Content-Type: application/json
Apikey: test
BODY:
-----
{
  "data": [
    {
      "id": 1,
      "status": 2,
      "test": {
        "description": "test",
        "id": 1
      }
    }
  ]
}
```
