package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jeffail/gabs/v2"
)

var port int = 8080

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	printHeaders(r)
	fmt.Println("")
	printBody(r)
}

func printHeaders(r *http.Request) {
	fmt.Println("HEADERS:")
	fmt.Println("--------")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
}

func printBody(r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	jsonParsed, err := gabs.ParseJSON(body)

	if err != nil {
		panic(err)
	}

	fmt.Println("BODY:")
	fmt.Println("-----")
	fmt.Println(jsonParsed.StringIndent("", "  "))
}

func main() {
	log.Println(fmt.Sprintf("server started on port: %d", port))
	http.HandleFunc("/", handleWebhook)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
