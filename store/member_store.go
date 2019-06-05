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

// UpdateMember updates a Member given an id and member
func (conn Connection) UpdateMember(member *model.Member) error {
	_, err := conn.db.Exec("UPDATE members SET first_name=$1, last_name=$2, email=$3 WHERE id=$4", &member.Firstname, &member.Lastname, &member.Email, &member.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}


// DeleteMember removes a member given an id
func (conn Connection) DeleteMember(id int) error {
	_, err := conn.db.Exec("DELETE FROM members WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
