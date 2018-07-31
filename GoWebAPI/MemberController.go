package main

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	login := logindata{
		Account:  "rdtest1233",
		Password: "123456"}

	if err := json.NewEncoder(w).Encode(login); err != nil {
		panic(err)
	}
}
