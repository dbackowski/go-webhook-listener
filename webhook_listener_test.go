package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func captureHandleWebhookConsoleOutput(response http.ResponseWriter, request *http.Request) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	HandleWebhook(response, request)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}

func TestHandleWebhook(t *testing.T) {
	t.Run("for valid json payload it returns propers headers and payload", func(t *testing.T) {
		var jsonData = []byte(`{ "name": "morpheus", "job": "leader" }`)
		request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response := httptest.NewRecorder()

		got := captureHandleWebhookConsoleOutput(response, request)
		want := "\x1b[32mHEADERS:\x1b[0m\n--------\nContent-Type: application/json; charset=UTF-8\n\n\x1b[32mBODY:\x1b[0m\n-----\n{ \"name\": \"morpheus\", \"job\": \"leader\" }\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		gotStatusCode := response.Result().StatusCode
		wantStatusCode := http.StatusOK

		if gotStatusCode != wantStatusCode {
			t.Errorf("got %v, want %v", gotStatusCode, wantStatusCode)
		}
	})

	t.Run("for invalid json payload it returns propers headers and payload", func(t *testing.T) {
		var jsonData = []byte(`{ "name": "morpheus", }`)
		request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response := httptest.NewRecorder()

		got := captureHandleWebhookConsoleOutput(response, request)
		want := "\x1b[32mHEADERS:\x1b[0m\n--------\nContent-Type: application/json; charset=UTF-8\n\n\x1b[31mCannot decode JSON payload!\x1b[0m\n{ \"name\": \"morpheus\", }\ninvalid character '}' looking for beginning of object key string\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		gotStatusCode := response.Result().StatusCode
		wantStatusCode := http.StatusBadRequest

		if gotStatusCode != wantStatusCode {
			t.Errorf("got %v, want %v", gotStatusCode, wantStatusCode)
		}
	})
}
