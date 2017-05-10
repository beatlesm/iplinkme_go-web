package main

import (
	"net/http"
	"net/http/httptest"
)

type ModifierMiddleware struct {
	handler http.Handler
}

func (this *ModifierMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println(r.Host)
	rec := httptest.NewRecorder()
	this.handler.ServeHTTP(rec, r)

	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	w.Header().Set("go-web-foundation", "vip")
	w.WriteHeader(418)
	w.Write([]byte("Hey, this is middleware!"))
	w.Write(rec.Body.Bytes())
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
func main() {
	single := &ModifierMiddleware{http.HandlerFunc(myHandler)}
	http.ListenAndServe(":8080", single)
}
