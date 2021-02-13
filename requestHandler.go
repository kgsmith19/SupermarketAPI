package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Grocery Store Inventory API!")
}

func fetchAllInventory(w http.ResponseWriter, r *http.Request) {

	if len(Inventory) < 1 {
		fmt.Fprintf(w, "No Produce Items exist in the Inventory Database.")
	} else {
		fmt.Fprintf(w, "Endpoint Hit: fetchAllInventory")
		json.NewEncoder(w).Encode(Inventory)
	}
}

func fetchProduceItem(w http.ResponseWriter, r *http.Request) {

	if len(Inventory) < 1 {
		fmt.Fprintf(w, "No Produce Items exist in the Inventory Database.")
	} else {
		vars := mux.Vars(r)
		key := vars["produceCode"]

		if vars["produceCode"] == "" {
			fmt.Fprintf(w, "Produce Code must have value to return Produce Item.")
		} else {

			// Loop over all of our Inventory
			// if the produceItem equals the key we pass in
			// return the pruduceItem encoded as JSON
			var produceItemFound = false
			for _, produceItem := range Inventory {
				if strings.ToUpper(produceItem.ProduceCode) == strings.ToUpper(key) {
					produceItemFound = true
					json.NewEncoder(w).Encode(produceItem)
				}
			}
			if produceItemFound == false {
				fmt.Fprintf(w, "Produce Item does not exist in the Inventory Database.")
			}
		}
	}
}

func addProduceItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var produceItem ProduceItem
	json.Unmarshal(reqBody, &produceItem)

	if validateProduceItem(produceItem) == false {
		fmt.Fprintf(w, APIValidationError)
		APIValidationError = ""
	} else {
		produceItem.UnitPrice = parseUnitPrice(produceItem.UnitPrice)
		Inventory = append(Inventory, produceItem)
		json.NewEncoder(w).Encode(produceItem)
	}
}

func deleteProduceItem(w http.ResponseWriter, r *http.Request) {
	if len(Inventory) < 1 {
		fmt.Fprintf(w, "No Produce Items exist in the Inventory Database to Delete.")
	} else {
		vars := mux.Vars(r)
		id := vars["produceCode"]

		if vars["produceCode"] == "" {
			fmt.Fprintf(w, "Produce Code must have value to Delete Produce Item.")
		} else {

			// Loop over all of our Inventory
			// if the produceItem equals the key we pass in
			// delete produce item
			var produceItemFound = false
			for index, produceItem := range Inventory {
				if strings.ToUpper(produceItem.ProduceCode) == strings.ToUpper(id) {
					produceItemFound = true
					Inventory = append(Inventory[:index], Inventory[index+1:]...)
					fmt.Fprintf(w, "Produce Item deleted.")
				}
			}
			if produceItemFound == false {
				fmt.Fprintf(w, "Produce Item does not exist in the Inventory Database.")
			}
		}
	}
}
