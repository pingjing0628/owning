package main

import (
	"net/http"

	"Users/pingjing/docker/goPractice/owning/routes"
)

func main() {
	r := routes.Routes()

	http.ListenAndServe(":8080", r)

}
