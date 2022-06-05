package golangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	dir := http.Dir("./resources")
	fileServer := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/static/",http.StripPrefix("/static/",fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


// With golang embed
//go:embed resources
var resources embed.FS

func TestFileServerWithGoEmbed(t *testing.T) {
	dir,_ := fs.Sub(resources,"resources")
	fileServer := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()
	mux.Handle("/static/",http.StripPrefix("/static/",fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}