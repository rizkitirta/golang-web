package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, req *http.Request){
		fmt.Fprintf(writer,"Hello World")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)  
	}
}

func TestServeMux(t *testing.T)  {
	mux := http.NewServeMux()

	mux.HandleFunc("/",func (writer http.ResponseWriter, req *http.Request)  {
		fmt.Fprint(writer,"Welcome To Golang Web")
	})

	mux.HandleFunc("/hello-world",func(writer http.ResponseWriter,req *http.Request){
		fmt.Fprint(writer,"Hello World")
	})

	mux.HandleFunc("/images/",func(writer http.ResponseWriter,req *http.Request){
		fmt.Fprint(writer,"images")
	})
	mux.HandleFunc("/images/thumbnails/",func(writer http.ResponseWriter,req *http.Request){
		fmt.Fprint(writer,"Thumbnails")
	})

	server := http.Server {
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)	
	}
}