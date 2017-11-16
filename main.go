package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
)

type Response struct {
	Status int
	Data   string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/card/generate", GenerateCard).Methods("POST")
	router.HandleFunc("/card/verify", VerifyCard).Methods("POST")
	router.HandleFunc("/card/identify/bank", IdentifyBank).Methods("POST")
	log.Fatal(http.ListenAndServe(":8092", router))
}

func GenerateCard(w http.ResponseWriter, r *http.Request) {
	response := Response{200, "success"}
	rs, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(rs)
}

func VerifyCard(w http.ResponseWriter, r *http.Request) {

}

func IdentifyBank(w http.ResponseWriter, r *http.Request) {

}
