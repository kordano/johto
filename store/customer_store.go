package store

import (
	model "github.com/kordano/johto/model"
	queries "github.com/kordano/johto/sql/queries"
)

// CreateCustomer creates new customer in sql db
func (conn Connection) CreateCustomer(customer *model.Customer) error {
	_, err := conn.db.Exec(queries.CustomerInsertQuery,
		&customer.Name, &customer.Contact, &customer.Department, &customer.Street,
		&customer.City, &customer.Postal, &customer.Country)
	if err != nil {
		return err
	}
	return err
}

// GetCustomers retrieves all customers from sql db
func (conn Connection) GetCustomers() ([]*model.Customer, error) {
	rows, err := conn.db.Query(queries.CustomerSelectAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*model.Customer

	for rows.Next() {
		cstmr := &model.Customer{}
		err := rows.Scan(&cstmr.ID, &cstmr.Name, &cstmr.Contact, &cstmr.Department, &cstmr.Street, &cstmr.City, &cstmr.Postal, &cstmr.Country)
		if err != nil {
			return nil, err
		}
		customers = append(customers, cstmr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

// UpdateCustomer updates a Customer given its ID and data
func (conn Connection) UpdateCustomer(customer *model.Customer) error {
	_, err := conn.db.Exec(queries.CustomerUpdateQuery, &customer.ID, &customer.Name, &customer.Contact, &customer.Department, &customer.Street, &customer.City, &customer.Postal, &customer.Country)
	return err
}

// DeleteCustomer deletes a Customer given an ID
func (conn Connection) DeleteCustomer(id int) error {
	_, err := conn.db.Exec(queries.CustomerDeleteQuery, id)
	return err
}
