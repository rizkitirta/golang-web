package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func IfStatement(w http.ResponseWriter,req *http.Request) {
	t := template.Must(template.ParseFiles("template/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{
		"Name": "Tirta Dev",
		"Age":  20,
	})
}

func TestIfstatement(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	IfStatement(recorder, req)
	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

// If Statement With Operator
func IfStatementWithOperator(w http.ResponseWriter,req *http.Request) {
	t := template.Must(template.ParseFiles("template/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{
		"Name": "Tirta Dev",
		"Age":  17,
	})
}

func TestIfstatementWithOperator(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	IfStatementWithOperator(recorder, req)
	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}