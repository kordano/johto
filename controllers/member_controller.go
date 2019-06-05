package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/kordano/johto/store"
	"github.com/kordano/johto/model"
)

// CreateMember inserts a new Member
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

// GetMembers retrieves all Members
func GetMembers(w http.ResponseWriter, r *http.Request) {
	conn, err := store.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	memberRefs, err := conn.GetMembers()
	if err != nil {
		fmt.Println(err)
		return
	}

	members := make([]model.Member, len(memberRefs))

	for i, v := range memberRefs {
		members[i] = *v
	}

	j, err := json.Marshal(MembersResource{Data: members})
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// UpdateMember update an existing member
// Handler for HTTP Put - "/members/{id}"
func UpdateMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}
	var dataResource MemberResource
	err = json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		fmt.Println(err)
		return
	}
	member := &dataResource.Data
	member.ID = id
	conn, err := store.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.UpdateMember(member)
}
