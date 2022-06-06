package golangweb

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	_ "embed"
)

func FormUpload(w http.ResponseWriter,req *http.Request) {
	t := template.Must(template.ParseFiles("./template/file_upload.gohtml"))
	t.ExecuteTemplate(w, "file_upload.gohtml",nil)
}

func TestFileUpload(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", FormUpload)
	mux.HandleFunc("/file-upload", UploadFile)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func UploadFile(w http.ResponseWriter, req *http.Request)  {
	//req.ParseMultipartForm(10 << 20) // 10 MB

	file,fileHeader,err := req.FormFile("file")
	if err != nil {
		panic(err)
	}
	destination,err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(destination, file)
	if err != nil {
		panic(err)
	}

	name := req.PostFormValue("name")

	t := template.Must(template.ParseFiles("./template/success_upload.gohtml"))
	t.ExecuteTemplate(w, "success_upload.gohtml",map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

//go:embed resources/restapi2.png
var image []byte
func TestUploadFile(t *testing.T)  {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "test upload file")
	file,err := writer.CreateFormFile("file", "contoh1.png")
	if err != nil {
		panic(err)
	}
	file.Write(image)
	writer.Close()

	req := httptest.NewRequest("POST", "http://localhost:8080/file-upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	UploadFile(recorder, req)

	response,_ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(response))
}