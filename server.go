package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")

		data := make(map[string]string)
		data["message"] = "Endpoint hit! (Go)"

		json.NewEncoder(w).Encode(data)
		
	})

	fmt.Printf("Starting server at port 8000\n")
	err := http.ListenAndServe(":8000", nil)
	if (err != nil){
		log.Fatal(err)
	}
}
