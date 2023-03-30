package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func cadastro_area(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_area")
	println("Descricao        : " + r.Form.Get("descricao"))
	println("************* radio/checkbox ***************")
	r.ParseForm()
	checkboxSelected := r.Form["name_checkbox"]
	radioSelected := r.Form["name_radio"]
	println("********************************************")
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
		d.TabelaDados[row][0] = strconv.Itoa(int(row))
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = descricao
		d.TabelaDados[row][3] = ts
		if contains(radioSelected, d.TabelaDados[row][0]) {
			println(d.TabelaDados[row][0] + " foi selecionado!")
			//Codificar aqui update ou delete
		}
		if contains(checkboxSelected, d.TabelaDados[row][0]) {
			println(d.TabelaDados[row][0] + " foi selecionado!")
			//Codificar aqui delete
		}
		row++
		d.Tot_elementos = row
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
