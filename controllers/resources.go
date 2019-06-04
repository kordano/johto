package controllers

import (
	"github.com/kordano/johto/model"
)

type(
	// MemberResource for Post and Put - /members
	// For Get - /members/id
	MemberResource struct {
		Data model.Member `json:"data"`
	}

	// MembersResource for Get - /members
	MembersResource struct {
		Data []model.Member `json:"data"`
	}
)
