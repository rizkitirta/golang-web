package golangweb

import (
	_"embed"
	"fmt"	
	"net/http"
	"testing"
)

func Servefile(w http.ResponseWriter,req *http.Request) {
	name := req.URL.Query().Get("name")
	if name != "" {
		http.ServeFile(w, req, "resources/index.html")
	}else{
		http.ServeFile(w,req,"resources/notFound.html")
	}
}


func TestServeFile(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(Servefile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Versi golang embed
//go:embed resources/index.html
var resourcesOk string
//go:embed resources/notFound.html
var resourcesNotFound string

func ServefileEmbed(w http.ResponseWriter,req *http.Request) {
	name := req.URL.Query().Get("name")
	if name != "" {
		fmt.Fprint(w,resourcesOk)
	}else{
		fmt.Fprint(w,resourcesNotFound)
	}
}


func TestServeFileEmbed(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServefileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}