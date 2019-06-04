package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/kordano/johto/store"
)

// CreateMember inst a new Member
func CreateMember(w http.ResponseWriter, r *http.Request) {
	var dataResource MemberResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		fmt.Println(err)
		return
	}
	member := &dataResource.Data
	conn, err := store.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = conn.CreateMember(member)
	if err != nil {
		fmt.Println(err)
		return
	}

	j, err := json.Marshal(MemberResource{Data: *member})

	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
