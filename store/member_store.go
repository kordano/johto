package store

import ("database/sql"
	"fmt"
	model "github.com/kordano/johto/model"
	_ "github.com/lib/pq")

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://johto:boofar@localhost:5433/johto")

	if err != nil {
		fmt.Println(err)
	}
}


func createMember(member model.Member) {
}
