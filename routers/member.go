package routers

import (
	"github.com/gorilla/mux"
	"github.com/kordano/johto/controllers"
)

// SetMemberRoutes registers routes for member entity
func SetMemberRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/members", controllers.CreateMember).Methods("POST")
	return router
}