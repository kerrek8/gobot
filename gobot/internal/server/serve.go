package server

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", alive)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func alive(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Alive")
	}
}
