package main

import (
	"net/http"
	"pratice/WEB4/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
