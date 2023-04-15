package main

import (
	"net/http"
	"strconv"
	"strings"
)

func cadastro_area(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_area")
	var cmd string
	r.ParseForm()
	checkboxSelected := r.Form["name_checkbox"]
	radioSelected := r.Form["name_radio"]
	println("Operacao         : " + r.Form.Get("operation"))
	println("Action           : " + strings.Trim(html, ".html"))
	for i := 0; i < sizeRows; i++ {
		for j := 0; j < sizeCols; j++ {
			d.TabelaDados[i][j] = ""
		}
	}
	if r.Form.Get("operation") == "Inserir" {
		cmd = "insert into area(descricao) values ('" + r.Form.Get("descricao") + "')"
		_, err := d.db.Exec(cmd)
		println(cmd)
		if err != nil {
			println(err)
		}
		println(r.Form.Get("descricao") + " foi incluido!")
	}
	var where string
	where = ""
	if r.Form.Get("operation") == "Localizar" {
		if r.Form.Get("descricao") != "" {
			where = " where descricao like '%" + r.Form.Get("descricao") + "%'"
		}
	}
	cmd = "select * from (select row_number() over (order by id) rownum, id, descricao, ts from area order by id) " + where
	rows, err := d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	var row int32
	row = 0
	for rows.Next() {
		var id int
		var descricao string
		var ts string
		var rownum int
		err = rows.Scan(&rownum, &id, &descricao, &ts)
		if err != nil {
			println(err)
		}
		d.TabelaDados[row][0] = strconv.Itoa(rownum - 1)
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = descricao
		d.TabelaDados[row][3] = ts
		row++
	}
	err = rows.Err()
	if err != nil {
		println(err)
	}
	defer rows.Close()
	//rows.Close()
	var Tot_elementos = row
	for row := 0; int(row) < int(Tot_elementos); row++ {
		var id string
		var descricao string
		id = d.TabelaDados[row][1]
		descricao = r.Form.Get("descricao")
		if contains(radioSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Alterar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "update area set descricao = '" + descricao + "' , ts = CURRENT_TIMESTAMP where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				println(d.TabelaDados[row][1] + " foi alterado de " + d.TabelaDados[row][2] + " para " + descricao + "!")
				d.TabelaDados[row][2] = descricao
			} else if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from area where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				println(d.TabelaDados[row][1] + " foi eliminado!")
				d.TabelaDados[row][0] = ""
				d.TabelaDados[row][1] = ""
				d.TabelaDados[row][2] = ""
				d.TabelaDados[row][3] = ""
			}
		}
		if contains(checkboxSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from area where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				println(d.TabelaDados[row][1] + " foi eliminado!")
				d.TabelaDados[row][0] = ""
				d.TabelaDados[row][1] = ""
				d.TabelaDados[row][2] = ""
				d.TabelaDados[row][3] = ""
			}
		}
	}
	if r.Form.Get("operation") == "Inserir" {
		d.Reload++
	} else if r.Form.Get("operation") == "Alterar" {
		d.Reload++
	} else if r.Form.Get("operation") == "Eliminar" {
		d.Reload++
	} else {
		d.Reload++
	}
	if d.Reload > 0 {
		d.Reload = -1
	}
	defer rows.Close()
}
