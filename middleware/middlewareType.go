package main

import (
	"net/http"
)

type SingleHost struct {
	handler     http.Handler
	allowedHost string
}

func (this *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println(r.Host)
	if r.Host == this.allowedHost {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}

}
func myhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
func main() {
	single := &SingleHost{
		handler:     http.HandlerFunc(myhandler),
		allowedHost: "10.1.10.195:8080",
		// allowedHost: "example.com",
	}
	http.ListenAndServe(":8080", single)
}
