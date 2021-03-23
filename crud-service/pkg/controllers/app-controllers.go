package controllers

import (
	"fmt"
	"net/http"
)

func ReturnHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnHealth")
	fmt.Fprintf(w, "{\"status\":\"OK\"}")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: home")
	fmt.Fprintf(w, "{\"hello\":\"world\"}")
}
