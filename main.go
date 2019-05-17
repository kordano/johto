package main

import (
	"fmt"
	model "github.com/kordano/johto/model"
	store "github.com/kordano/johto/store"
)

func main() {

	members := []*model.Member {
		&model.Member{Firstname: "Konrad", Lastname: "Kühne", Email: "konrad.kuehne@lambdaforge.io", Password: "swordfish"},
		&model.Member{Firstname: "Alice", Lastname: "From Resident Evil", Email: "alice@umbrella.corp", Password: "zombie"}	,
		&model.Member{Firstname: "Bob", Lastname: "der Baumeister", Email: "bob@baustel.le", Password: "cement"},
	}

	for _, mbr := range members {
		err := store.CreateMember(mbr)
		if err != nil {
			fmt.Println(err)
		}
	}

	sqlMembers, err := store.GetMembers()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, mbr := range sqlMembers {
			fmt.Println(mbr.Email)
		}
	}
}
