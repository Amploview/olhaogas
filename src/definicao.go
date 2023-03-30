package main

import "database/sql"

const sizeCols = 10
const sizeRows = 30

type data struct {
	Title         string
	Header        string
	Host          string
	db            *sql.DB
	TabelaDados   [sizeRows][sizeCols]string
	Tot_elementos int32
}
