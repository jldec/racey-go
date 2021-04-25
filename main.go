package main

import (
	"fmt"
	"net/http"

	counter "github.com/jldec/counter-go"
)

func main() {
	count := new(counter.CounterAtomic)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count.Inc()
		fmt.Fprintln(w, count.Get())
	})

	fmt.Println("Go listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
