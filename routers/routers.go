package routers

import ("github.com/gorilla/mux")

// InitRoutes registers all routes for the application
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetMemberRoutes(router)
	return router
}
