package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// Middleware
type Middleware struct {
	Handler http.Handler
}

func (middleware *Middleware) ServeHTTP(w http.ResponseWriter,req *http.Request)  {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(w,req)
	fmt.Println("After execute handler")
}

// Error handling
type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("Recovered in f",err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s",err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w,req)
}

func TestMiddleware(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
		fmt.Fprint(w, "Hello middleware")
	})
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		panic("ups terjadi error")
	})
	
	middleware := &Middleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: middleware,
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}