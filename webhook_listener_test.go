package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleWebhook(t *testing.T) {
	t.Run("for valid json payload it returns propers headers and payload", func(t *testing.T) {
		var jsonData = []byte(`{ "name": "morpheus", "job": "leader" }`)
		request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response := httptest.NewRecorder()

		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		HandleWebhook(response, request)

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		fmt.Println(string(out))

		got := string(out)
		want := "\x1b[32mHEADERS:\x1b[0m\n--------\nContent-Type: application/json; charset=UTF-8\n\n\x1b[32mBODY:\x1b[0m\n-----\n{ \"name\": \"morpheus\", \"job\": \"leader\" }\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("for invalid json payload it returns propers headers and payload", func(t *testing.T) {
		var jsonData = []byte(`{ "name": "morpheus", }`)
		request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response := httptest.NewRecorder()

		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		HandleWebhook(response, request)

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		fmt.Println(string(out))

		got := string(out)
		want := "\x1b[32mHEADERS:\x1b[0m\n--------\nContent-Type: application/json; charset=UTF-8\n\n\x1b[31mCannot decode JSON payload!\x1b[0m\n{ \"name\": \"morpheus\", }\ninvalid character '}' looking for beginning of object key string\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
