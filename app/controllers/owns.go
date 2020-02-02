package controllers

import (
	"Users/pingjing/docker/goPractice/owning/app/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	model := model.Products{}

	if err := model.FindAll(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&model)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	model := model.Products{}
	params := mux.Vars(r)

	if err := model.FindOne(bson.M{"productId": params["id"]}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&model)

}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	model := model.Products{}

	q := json.NewDecoder(r.Body).Decode(&model)
	fmt.Println(q)

	if err := model.Insert(q); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&model)
}

// func Store() {

// }
