package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/kordano/johto/store"
	"github.com/kordano/johto/model"
)

// CreateCustomer inserts a new Customer
// Handler for POST - "/customers"
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var dataResource CustomerResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		fmt.Println(err)
		return
	}
	customer := &dataResource.Data
	conn, err := store.GetConnection()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = conn.CreateCustomer(customer)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetCustomers retrieves all Customers
// Handler for GET  - "/customers"
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	conn, err := store.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	customerRefs, err := conn.GetCustomers()
	if err != nil {
		fmt.Println(err)
		return
	}
	customers := make([]model.Customer, len(customerRefs))
	for i, v := range customerRefs {
		customers[i] = *v
	}

	j, err := json.Marshal(CustomersResource{Data: customers})
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// UpdateCustomer updates an existing customer
// Handler for PUT - "/customers/{id}"
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}
	var dataResource CustomerResource
	err = json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		fmt.Println(err)
		return
	}
	customer := &dataResource.Data
	customer.ID = id
	conn, err := store.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.UpdateCustomer(customer)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// DeleteCustomer deletes an existing customer given an ID
// Handler for DELETE - "/customers/{id}"
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := store.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.DeleteCustomer(id)
	if err != nil {
		fmt.Println(err)
		return
	}
}
