package main

import (
	"fmt"
	"net/http"

	"github.com/kordano/johto/routers"
)


func main() {

	router := routers.InitRoutes()
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: router,
	}
	fmt.Println("Listening...")
	server.ListenAndServe()


	/*
	members := []*model.Member {
		&model.Member{Firstname: "Konrad", Lastname: "Kühne", Email: "konrad.kuehne@lambdaforge.io", Password: "swordfish"},
		&model.Member{Firstname: "Alice", Lastname: "From Resident Evil", Email: "alice@umbrella.corp", Password: "zombie"}	,
		&model.Member{Firstname: "Bob", Lastname: "der Baumeister", Email: "bob@baustel.le", Password: "cement"},
	}

	conn, err := store.GetConnection()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, mbr := range members {
		err := conn.CreateMember(mbr)
		if err != nil {
			fmt.Println(err)
		}
	}

	sqlMembers, err := conn.GetMembers()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, mbr := range sqlMembers {
			fmt.Println(mbr.Email)
		}
	}
	*/
}
