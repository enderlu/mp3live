package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	e := http.ListenAndServe(":8080", nil)
	fmt.Println(e)
}
