package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHelloParam(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello %s", name)
}

func TestParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=tirta", nil)
	recorder := httptest.NewRecorder()

	SayHelloParam(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	data := string(body)
	fmt.Println(data)
	fmt.Println("Done")
}

// Multiple parameter
func SayHelloMultipleParam(w http.ResponseWriter, req *http.Request) {
	firstName := req.URL.Query().Get("first_name")
	lastName := req.URL.Query().Get("last_name")
	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleParam(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello?first_name=Rizki&last_name=Tirta", nil)
	rec := httptest.NewRecorder()

	SayHelloMultipleParam(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// multiple value with same name parameter
func SayHalloMultipleValue(w http.ResponseWriter, req *http.Request) {
	names := req.URL.Query()["name"]
	fmt.Fprintf(w, "Hello %s", strings.Join(names, " "))
}

func TestMultipleValue(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Tirta&name=Joko&name=Arif", nil)
	rec := httptest.NewRecorder()

	SayHalloMultipleValue(rec, req)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
