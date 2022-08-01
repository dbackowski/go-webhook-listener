# Webhook listener ![Tests](https://github.com/dbackowski/go-webhook-listener/actions/workflows/test.yml/badge.svg)

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

![screenshot](https://i.imgur.com/7P0dlkP.png)
```
