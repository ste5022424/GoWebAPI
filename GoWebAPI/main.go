package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
