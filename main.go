// A simple webserver printing RFC3339 time.

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func outputTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%s\n", time.Now().Format(time.RFC3339))
}

func main() {
	http.HandleFunc("/now", outputTime)
	err := http.ListenAndServe(":9190", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
