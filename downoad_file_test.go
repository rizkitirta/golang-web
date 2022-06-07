package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, req *http.Request) {
	fileName := req.URL.Query().Get("file")

	if fileName == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "File not found")
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	http.ServeFile(w,req,"./resources/"+fileName)
}

func TestDownloadFile(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {panic(err)}
}