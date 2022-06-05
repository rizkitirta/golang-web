package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Versi Map
func TemplateDataMap(writer http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("template/template_data_map.gohtml"))
	t.ExecuteTemplate(writer, "template_data_map.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Golang Web Data Map",
		"Address": map[string]interface{}{
			"Street": "Jl. Raya 3",
		},
	})
}

func TestTemplateData(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, req)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// Versi Struct
type Page struct {
	Title   string
	Name    string
	Address Address
}

type Address struct {
	Street string
}

func TemplateDataMapStruct(writer http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("template/template_data_map.gohtml"))
	t.ExecuteTemplate(writer, "template_data_map.gohtml", Page{
		Title: "Template Data Map Struct",
		Name:  "Tirta",
		Address: Address{
			Street: "Jl. Raya 21",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMapStruct(recorder, req)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
