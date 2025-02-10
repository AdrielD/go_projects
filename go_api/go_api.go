package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

type Router struct {}

func (router *Router) handleRequet(method, path string, action func() string) {
	http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			http.NotFound(w, req)
			return
		}

		log.Printf("[%s] %q\n", req.Method, html.EscapeString(req.URL.Path))
		// w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, action())
	})
}

func (router *Router) Get(path string, action func() string) {
	router.handleRequet(http.MethodGet, path, action)
}

func (router *Router) Post(path string, action func() string) {
	router.handleRequet(http.MethodPost, path, action)
}

func (router *Router) Put(path string, action func() string) {
	router.handleRequet(http.MethodPut, path, action)
}

func (router *Router) Patch(path string, action func() string) {
	router.handleRequet(http.MethodPatch, path, action)
}

func (router *Router) Delete(path string, action func() string) {
	router.handleRequet(http.MethodDelete, path, action)
}


func sayAhoy() string {
	return "Ahoy!"
}

func main() {
	var router Router

	router.Get("/foo", sayAhoy)
	router.Post("/bar", sayAhoy)
	router.Put("/woo", sayAhoy)
	router.Patch("/wab", sayAhoy)
	router.Delete("/del", sayAhoy)

	portFlag := flag.Int("port", 8090, "server port")
	flag.Parse()

	port := strconv.Itoa(*portFlag)
	log.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
