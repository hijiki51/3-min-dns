package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Write([]byte(fmt.Sprintf("Hello, I'm %s!\n", hostname)))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	return
}
