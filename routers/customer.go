package routers

import (
	"github.com/gorilla/mux"
	"github.com/kordano/johto/controllers"
)

// SetCustomerRoutes registers routes for customer entity
func SetCustomerRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")
	return router
}
