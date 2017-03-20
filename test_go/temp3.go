package main

import (
	"io"
	"log"
	"net/http"
	"time"
)
var mux map[string]func(http.ResponseWriter, *http.Request)
func main() {
	server:=&http.Server{
		Addr:         ":8080",
		Handler:      &myHandler{},
		ReadTimeout:  5 * time.Second,
	}
	mux = make(map[string]func (http.ResponseWriter, *http.Request))
	mux["/hello"] = SayHello
	mux["/bye"] = SayBye

	err:=server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h,ok:= mux[r.URL.String()]; ok{
		h(w, r)
		return
	}
	io.WriteString(w, "URL:"+r.URL.String())
}
func SayHello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello World!! this is version 3.")
}
func SayBye(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Bye bye!! this is version 3.")
}

