package store

import (
	"fmt"
	model "github.com/kordano/johto/model"
	)


// CreateMember creates new member in sql db
func (conn Connection) CreateMember(member *model.Member) error {
	result, err := conn.db.Exec("INSERT INTO members(first_name, last_name, email, salt, passhash) VALUES($1, $2, $3, $4, $5)",
		&member.Firstname, &member.Lastname, &member.Email,"foobar", &member.Password)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Member with id=%d created successfully (%d row affected)\n", lastInsertID, rowsAffected)
	return err
}

// GetMembers retrieves all members from sql db
func (conn Connection) GetMembers() ([]*model.Member, error) {
	rows, err := conn.db.Query("SELECT id, first_name, last_name, email FROM members")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*model.Member
	for rows.Next() {
		mbr := &model.Member{}
		err := rows.Scan(&mbr.ID, &mbr.Firstname, &mbr.Lastname, &mbr.Email)
		if err != nil {
			return nil, err
		}
		members = append(members, mbr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}
