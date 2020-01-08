package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	log.Println(*r)
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Println(http.ListenAndServe(":3000", nil))
}
