package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// Version 1
	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")

	// Version 2
	firstNameV2 := r.FormValue("firstName")
	lastNameV2 := r.FormValue("lastName")
	
	fmt.Fprintf(w, "%s %s", firstName, lastName)
	fmt.Fprintf(w, "%s %s", firstNameV2, lastNameV2)
}

func TestFormPost(t *testing.T) {
	data := strings.NewReader("firstName=Rizki&lastName=Tirta")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", data)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
