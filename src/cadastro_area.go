package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func cadastro_area(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_area")
	println("Descricao        : " + r.Form.Get("descricao"))
	println("Id_hidden        : " + r.Form.Get("id_hidden_"+r.Form.Get("name_radio")))
	println("Descricao_hidden : " + r.Form.Get("descricao_hidden_"+r.Form.Get("name_radio")))
	println("Ts_hidden        : " + r.Form.Get("ts_hidden_"+r.Form.Get("name_radio")))
	println("Operacao         : " + r.Form.Get("operation"))
	println("Html             : " + html)
	println("Action           : " + strings.Trim(html, ".html"))
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
		//fmt.Println(id, descricao, ts)
		d.TabelaDados[row][0] = strconv.Itoa(id)
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = descricao
		d.TabelaDados[row][3] = ts
		row++
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
