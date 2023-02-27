package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	portPtr := flag.Int("port", 8080, "port number")
	flag.Parse()
	port := *portPtr

	log.Println(fmt.Sprintf("server started on port: %d", port))
	http.HandleFunc("/", HandleWebhook)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
