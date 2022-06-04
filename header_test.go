package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	data := string(body)
	fmt.Println(data)
}

// Response Header From Server
func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Power-By", "Tirta")
	fmt.Fprint(w, "ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	data := string(body)
	fmt.Println(data)

	fmt.Println(response.Header.Get("X-Power-By"))
}
