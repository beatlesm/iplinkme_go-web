package main

import (
	"net/http"
)

type AppendMiddleware struct {
	handler http.Handler
}

func (this *AppendMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.handler.ServeHTTP(w, r)
	w.Write([]byte("Hey, this is middleware!"))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
func main() {
	single := &AppendMiddleware{http.HandlerFunc(myHandler)}
	http.ListenAndServe(":8080", single)
}
