package main

import (
	"fmt"
	"log"
	"net/http"
)

var port int = 8080

func main() {
	log.Println(fmt.Sprintf("server started on port: %d", port))
	http.HandleFunc("/", HandleWebhook)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
