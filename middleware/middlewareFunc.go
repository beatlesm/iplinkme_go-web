package main

import (
	"net/http"
)

func SingleHost(handler http.Handler, allowedHost string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Host == allowedHost {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}
	return http.HandlerFunc(fn)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
func main() {
	single := SingleHost(http.HandlerFunc(myHandler), "10.1.10.195:8080")
	http.ListenAndServe(":8080", single)
}
