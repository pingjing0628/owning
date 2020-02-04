package controllers

import (
	"Users/pingjing/docker/goPractice/owning/app/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	model := model.Product{}

	if err := model.FindAll(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&model)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	model := model.Product{}
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

	product := model.Product{}

	_ = json.NewDecoder(r.Body).Decode(&product)

	// fmt.Println(model.ProductName)

	if err := product.Insert(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	product := model.Product{}

	_ = json.NewDecoder(r.Body).Decode(&product)

	// fmt.Println(model.ProductName)

	if err := product.Update(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&product)
}
