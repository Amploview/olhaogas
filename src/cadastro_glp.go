package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func cadastro_glp(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_glp")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))

	rows, err := d.db.Query("select id, descricao, glp_default, ts from glp")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var row int32
	row = 0
	for rows.Next() {
		var id int
		var descricao string
		var glp_default int
		var ts string
		err = rows.Scan(&id, &descricao, &glp_default, &ts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, descricao, glp_default, ts)
		d.TabelaDados[row][0] = strconv.Itoa(id)
		d.TabelaDados[row][1] = descricao
		if glp_default == 1 {
			d.TabelaDados[row][2] = "Default"
		} else {
			d.TabelaDados[row][2] = ""
		}
		d.TabelaDados[row][3] = ts
		row++
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
