package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	receipt ".."
)

var templates = template.Must(template.ParseGlob("templates/*"))

func receiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	decoder := json.NewDecoder(r.Body)
	// "Object than Array" is a best practice:)
	products := make(map[string]float64)
	if err := decoder.Decode(&products); err != nil {
		log.Printf("[request]json decoder error: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cart, err := receipt.NewCart(products)
	if err != nil {
		log.Printf("[cart]generate error: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}
	cart.Checkout()

	if err := templates.ExecuteTemplate(w, "text", cart); err != nil {
		log.Printf("[template]render error: %s", err)
	}
}

func init() {
	receipt.PrepareDB("./db.sqlite3")
}

func main() {
	http.HandleFunc("/", receiptHandler)
	http.ListenAndServe(":8080", nil)
}
