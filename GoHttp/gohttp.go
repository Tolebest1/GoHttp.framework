package GoHttp

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine{
	return &Engine {router:newRouter()}
}

func (e *Engine) addRoute(method string,pattern string,handler HandlerFunc){
	e.router.addRoute(method,pattern,handler)
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
	c:=newContext(w,r)
	e.router.handle(c)
}


