package main

import "database/sql"

const sizeCols = 15
const sizeRows = 30

type data struct {
	Title       string
	Header      string
	Host        string
	db          *sql.DB
	TabelaDados [sizeRows][sizeCols]string //comportar rownum e colunas das tabelas do sistema
	Reload      int
	Area        [sizeRows][3]string //comportar rownum, id, descricao
	Glp         [sizeRows][3]string //comportar rownum, id, descricao
	Cliente     [sizeRows][5]string //comportar rownum, id, nome, telefone, email
	Usuario     [sizeRows][5]string //comportar rownum, id, login, nome, tipo_usuario
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}
