package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func cadastro_cobertura(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_cobertura")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))

	rows, err := d.db.Query("select id, id_usuario, id_area, ts from cobertura")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var id_usuario int
		var id_area int
		var ts string
		err = rows.Scan(&id, &id_usuario, &id_area, &ts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, id_usuario, id_area, ts)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
