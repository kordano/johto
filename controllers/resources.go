package controllers

import (
	"github.com/kordano/johto/model"
)

type (
	// MemberResource for Post and Put - /members
	MemberResource struct {
		Data model.Member `json:"data"`
	}

	// MembersResource for Get - /members
	MembersResource struct {
		Data []model.Member `json:"data"`
	}

	// CustomerResource for POST and PUT - /customers
	CustomerResource struct {
		Data model.Customer `json:"data"`
	}

	// CustomersResource for GET - /customers
	CustomersResource struct {
		Data []model.Customer `json:"data"`
	}
)
