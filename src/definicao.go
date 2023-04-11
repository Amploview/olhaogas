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
	Reload        int
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}
