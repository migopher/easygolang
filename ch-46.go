package main

import (
	"fmt"
	"net/http"
)

/**
包
 */
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "golang")
	fmt.Fprint(w, `hell world!`)
}
