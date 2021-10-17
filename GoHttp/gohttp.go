package GoHttp

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter,*http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine{
	return &Engine {router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method string,pattern string,handler HandlerFunc){
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) GET (pattern string, handler HandlerFunc){
	e.addRoute("GET",pattern,handler)
}

func (e *Engine) POST (pattern string, handler HandlerFunc){
	e.addRoute("POST",pattern,handler)
}

func (e *Engine) RUN (address string)(err error){
	return http.ListenAndServe(address,e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method+"-"+r.URL.Path
	if hanlder,ok :=e.router[key];ok{
		hanlder(w,r)
	}else{
		fmt.Fprintf(w,"404 NOT FOUND %s",r.URL)
	}
}


