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
	for i := 0; i < sizeRows; i++ {
		for j := 0; j < sizeCols; j++ {
			d.TabelaDados[i][j] = ""
		}
	}
	if r.Form.Get("operation") == "Inserir" {
		_, err := d.db.Exec("insert into area(descricao) values ('" + r.Form.Get("descricao") + "')")
		if err != nil {
			log.Fatal(err)
		}
		println(r.Form.Get("descricao") + " foi incluida!")
	}
	var where string
	where = ""
	if r.Form.Get("operation") == "Localizar" {
		if r.Form.Get("descricao") != "" {
			where = " where descricao like '" + r.Form.Get("descricao") + "%'"
		}
	}
	rows, err := d.db.Query("select id, descricao, ts from area" + where)
	if err != nil {
		log.Fatal(err)
	}
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
		row++
		d.Tot_elementos = row
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for row := 0; int(row) < int(d.Tot_elementos); row++ {
		var id string
		var descricao string
		id = d.TabelaDados[row][1]
		descricao = d.TabelaDados[row][2]
		if contains(radioSelected, d.TabelaDados[row][0]) {
			println(d.TabelaDados[row][0] + " foi selecionado!")
			if r.Form.Get("operation") == "Alterar" {
				_, err := d.db.Exec("update area set descricao = '" + descricao +
					"' ,ts = CURRENT_TIMESTAMP where id = " + id)
				if err != nil {
					log.Fatal(err)
				}
				println(d.TabelaDados[row][0] + " foi alterado!")
			} else if r.Form.Get("operation") == "Eliminar" {
				_, err := d.db.Exec("delete from area where id = " + id)
				if err != nil {
					log.Fatal(err)
				}
				println("delete from area where id = " + id)
				println(d.TabelaDados[row][0] + " foi eliminado!")
			}
		}
		if contains(checkboxSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][0] + " foi selecionado!")
				_, err := d.db.Exec("delete from area where id = " + id)
				if err != nil {
					log.Fatal(err)
				}
				println(d.TabelaDados[row][0] + " foi eliminado!")
			}
		}
	}
	defer rows.Close()
}
