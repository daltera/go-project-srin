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
		data["message"] = "Endpoint hit!"

		json.NewEncoder(w).Encode(data)
		
	})

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if (err != nil){
		log.Fatal(err)
	}
}