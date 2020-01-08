package main

import (
	"fmt"
	"net/http"
)

const (
	port = 8080
)

func main() {
	http.ListenAndServe(fmt.Sprintf(":%d", port), IntializeEngine(10, 15))
}
