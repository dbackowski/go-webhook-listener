package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var responseBody = ""

func main() {
	portPtr := flag.Int("port", 8080, "port number")
	responsePtr := flag.String("response", "", "response body that will be sent back")

	flag.Parse()
	port := *portPtr
	responseBody = *responsePtr

	log.Println(fmt.Sprintf("server started on port: %d", port))
	http.HandleFunc("/", HandleWebhook)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
