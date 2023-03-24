package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func cadastro_area(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_area")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))

	if r.Form.Get("operation") == "Inserir" {

		_, err := d.db.Exec("insert into area(descricao) values ('" + r.Form.Get("descricao") + "')")
		if err != nil {
			log.Fatal(err)
		}

	}

	rows, err := d.db.Query("select id, descricao, ts from area")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var row int32
	row = 0
	for rows.Next() {
		var id int
		var descricao string
		var ts string
		err = rows.Scan(&id, &descricao, &ts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, descricao, ts)
		d.TabelaIds[row] = strconv.Itoa(id)
		d.TabelaDados[row][0] = descricao
		d.TabelaDados[row][1] = descricao
		d.TabelaDados[row][2] = ts
		row++
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
