package mygee

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	routers map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{routers: make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	e.routers[key] = handler
}

func (e *Engine) Get(route string, handler HandlerFunc) {
	e.addRoute("get", route, handler)
}
func (e *Engine) Post(route string, handler HandlerFunc) {
	e.addRoute("post", route, handler)
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.routers[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
