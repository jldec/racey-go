package main

import (
	"fmt"
	"net/http"
)

func main() {
	var count uint64 = 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Fprintln(w, count)
	})

	fmt.Println("Go listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
