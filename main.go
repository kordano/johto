package main

import (
	"fmt"
	model "github.com/kordano/johto/model"
)

func main() {
	member := model.Member{ID: 1, Firstname: "Konrad", Lastname: "Kühne", Email: "konrad.kuehne@lambdaforge.io", Password: "swordfish"}

	fmt.Println(member)
	fmt.Println("Hello World!")
}
