package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/stretchr/testify/assert"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	data := string(body)
	fmt.Println(data)
}
