package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartUp() *mux.Router {

	mux := mux.NewRouter()

	//home
	mux.HandleFunc("/test", Index)
	return mux
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Moin Moin beim Lotse!")
}