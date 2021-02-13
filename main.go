package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// APIValidationError used to check validation for acceptance criteria of API inputs
var APIValidationError string

//Inventory is here
var Inventory []ProduceItem

// ProduceItem - Our struct for all inventories
type ProduceItem struct {
	ProduceCode string `json:"ProduceCode"`
	Name        string `json:"Name"`
	UnitPrice   string `json:"UnitPrice"`
}

func main() {

	createInventoryArray()
	handleRequests()
}

func handleRequests() {
	//Put API calls below
	myRouter := mux.NewRouter().StrictSlash(true)

	// home page
	myRouter.HandleFunc("/", homePage).Methods("GET")

	// fetch All Inventory
	myRouter.HandleFunc("/inventory", fetchAllInventory).Methods("GET")

	// fetch produce item
	myRouter.HandleFunc("/produceItem/{produceCode}", fetchProduceItem).Methods("GET")

	// add new produce item
	myRouter.HandleFunc("/addProduceItem", addProduceItem).Methods("POST")

	// delete new produce item
	myRouter.HandleFunc("/deleteProduceItem/{produceCode}", deleteProduceItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createInventoryArray() {
	Inventory = []ProduceItem{
		{ProduceCode: "A12T-4GH7-QPL9-3N4M", Name: "Lettuce", UnitPrice: "$3.46"},
		{ProduceCode: "E5T6-9UI3-TH15-QR88", Name: "Peach", UnitPrice: "$2.99"},
		{ProduceCode: "YRT6-72AS-K736-L4AR", Name: "Green Pepper", UnitPrice: "$0.79"},
		{ProduceCode: "TQ4C-VV6T-75ZX-1RMR", Name: "Gala Apple", UnitPrice: "$3.59"},
	}
}
