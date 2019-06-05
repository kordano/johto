package routers

import (
	"github.com/gorilla/mux"
	"github.com/kordano/johto/controllers"
)

// SetMemberRoutes registers routes for member entity
func SetMemberRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/members", controllers.CreateMember).Methods("POST")
	router.HandleFunc("/members", controllers.GetMembers).Methods("GET")
	router.HandleFunc("/members/{id}", controllers.UpdateMember).Methods("PUT")
	router.HandleFunc("/members/{id}", controllers.DeleteMember).Methods("DELETE")
	return router
}
