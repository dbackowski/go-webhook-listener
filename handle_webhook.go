package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jeffail/gabs/v2"
	"github.com/dbackowski/colors"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println(colors.Colorize("Cannot decode JSON payload!", colors.FgRed))
		fmt.Println(string(body))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(jsonParsed.StringIndent("", "  "))
}
