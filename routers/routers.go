package routers

import ("github.com/gorilla/mux")

// InitRoutes registers all routes for the application
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetMemberRoutes(router)
	SetCustomerRoutes(router)
	return router
}
