package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home!")
}

func Detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DetailID := vars["DetailID"]
	fmt.Fprintln(w, "最新消息 Detail ID:", DetailID)
}
