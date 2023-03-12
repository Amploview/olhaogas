package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func cadastro_preco(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_preco")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))

	rows, err := d.db.Query("select id, id_glp, id_area, preco, ts from glp_preco_area")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var id_glp int
		var id_area int
		var preco float64
		var ts string
		err = rows.Scan(&id, &id_glp, &id_area, &preco, &ts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, id_glp, id_area, preco, ts)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
