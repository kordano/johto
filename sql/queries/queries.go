package sql

// CustomerInsertQuery defines customer insertion sql
var CustomerInsertQuery = "INSERT INTO customers(name, contact, department, street, city, postal, country) VALUES($1, $2, $3, $4, $5, $6, $7)"

// CustomerSelectAllQuery defines customer all selection sql
var CustomerSelectAllQuery = "SELECT id, name, contact, department, street, city, postal, country from customers"

// CustomerUpdateQuery defines customer update sql
var CustomerUpdateQuery = "UPDATE customers SET name=$2, contact=$3, department=$4, street=$5, city=$6, postal=$7, country=$8 WHERE id=$1"

// CustomerDeleteQuery defines customer deletion sql
var CustomerDeleteQuery = "DELETE FROM customers where id=$1"
