package golangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Versi Text
func SimpleHtml(w http.ResponseWriter, req *http.Request) {
	templateText := "<html><body>{{.}}</body></html>"
	tmp, err := template.New("simpleHtml").Parse(templateText)
	if err != nil {
		panic(err)
	}
	tmp.ExecuteTemplate(w, "simpleHtml", "Hello World")
}

func TestSimpleTemplate(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, req)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// Versi File
func SimpleHtmlFile(w http.ResponseWriter, req *http.Request) {
	templateFile := template.Must(template.ParseFiles("./template/simple.gohtml"))
	templateFile.ExecuteTemplate(w, "simple.gohtml","Hello World With File")
}

func TestSimpleTemplateFile(t *testing.T)  {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlFile(recorder, req)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// Versi Inclusde All File Template
func SimpleHtmlFileDirectory(w http.ResponseWriter, req *http.Request) {
	templateFile := template.Must(template.ParseGlob("./template/*.gohtml"))
	templateFile.ExecuteTemplate(w, "simple.gohtml","Hello World With File Directory")
}

func TestSimpleTemplateFileDirectory(t *testing.T)  {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlFileDirectory(recorder, req)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

//go:embed template/*.gohtml
var templates embed.FS

// Versi Inclusde All File Template WITH GO EMBED
func SimpleHtmlWithGoEmbed(w http.ResponseWriter, req *http.Request) {
	templateFile := template.Must(template.ParseFS(templates,"template/*.gohtml"))
	templateFile.ExecuteTemplate(w, "simple.gohtml","Hello World With Go Embed")
}

func TestSimpleWithGoEmbed(t *testing.T)  {
	req := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlWithGoEmbed(recorder, req)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}