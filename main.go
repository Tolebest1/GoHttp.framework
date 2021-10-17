package main

import (
	"GoHttp"
	"fmt"
	"net/http"
)

func main() {
	r:=GoHttp.New()
	r.GET("/",func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
	})

	r.GET("/hello",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
	})

	r.RUN(":9999")
}
