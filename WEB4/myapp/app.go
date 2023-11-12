package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get userinfo by /user/{id}")
}

func getUserInfo89Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "User Id:", vars["id"])
}

func NewHandler() http.Handler {
	mux := mux.NewRouter()
	//mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfo89Handler)
	return mux
}
