package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {

}
/*
 package http

type Handler interface {
     ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
*/
//(engine *Engine) 表示实现了 ServeHTTP 这个接口 ==》 具体是哪一个接口还需要看后面的
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
