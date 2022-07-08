package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jeffail/gabs/v2"
	"github.com/dbackowski/colors"
)

var port int = 8080

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	printHeaders(r)
	fmt.Println("")
	printBody(w, r)
}

func printHeaders(r *http.Request) {
	fmt.Println(colors.Colorize("HEADERS:", colors.FgGreen))
	fmt.Println("--------")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
}

func printBody(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	jsonParsed, err := gabs.ParseJSON(body)

	fmt.Println(colors.Colorize("BODY:", colors.FgGreen))
	fmt.Println("-----")

	if err != nil {
		fmt.Println("Cannot decode JSON payload!")
		fmt.Println(string(body))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(jsonParsed.StringIndent("", "  "))
}

func main() {
	log.Println(fmt.Sprintf("server started on port: %d", port))
	http.HandleFunc("/", handleWebhook)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
