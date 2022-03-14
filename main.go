package main

import "net/http"

func main() {
	var h http.Handler
	http.ListenAndServe(":8080", h)
}
