package requests

import "github.com/go-playground/validator"

type Request struct {
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate
