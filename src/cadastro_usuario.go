package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func cadastro_usuario(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_usuario")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))

	rows, err := d.db.Query("select id, login, senha, nome, flg_tipo_usuario, ts from usuario")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var login string
		var senha string
		var nome string
		var flg_tipo_usuario int
		var ts string
		err = rows.Scan(&id, &login, &senha, &nome, &flg_tipo_usuario, &ts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, login, senha, nome, flg_tipo_usuario, ts)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
