package routes

import (
	"Users/pingjing/docker/goPractice/owning/app/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/owns", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/owns", controllers.PostProduct).Methods("POST")
	r.HandleFunc("/owns/{id}", controllers.GetProduct).Methods("GET")
	// r.HandleFunc("/owns/{id}/edit", controllers.EditProduct)
	// r.HandleFunc("/owns/{id}", controllers.UpdateProduct)
	// r.HandleFunc("/owns/{id}", controllers.DeleteProduct).Methods("DELETE")

	return r
}
