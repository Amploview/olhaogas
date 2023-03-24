package main

import "database/sql"

const sizeCols = 10
const sizeRows = 30

type data struct {
	Title       string
	Header      string
	Host        string
	db          *sql.DB
	TabelaIds   [sizeRows]string
	TabelaDados [sizeRows][sizeCols]string
}
