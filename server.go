package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", handlerWebServiceLog)
	server := http.Server{Addr: ":8080", Handler: handler}
	server.ListenAndServe()
}

func handlerWebServiceLog(w http.ResponseWriter, r *http.Request) {
	var responseBody string
	responseBody = fmt.Sprintln(r.Method)
	responseBody += fmt.Sprintln(r.RequestURI)
	responseBody += fmt.Sprintln("")
	for i, m := range r.Header {
		responseBody += fmt.Sprintln(i, m)
	}
	responseBody += fmt.Sprintln("")
	body, _ := io.ReadAll(r.Body)
	responseBody += fmt.Sprintln(string(body))
	//w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseBody))
}
