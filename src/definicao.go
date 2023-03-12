package main

import "database/sql"

const sizeCols = 10
const sizeRows = 30

type data struct {
	Title  string
	Header string
	Host   string
	db     *sql.DB
	Tabela [sizeRows][sizeCols]string
}
