package main

import (
	"net/http"
	"strconv"
	"strings"
)

func cadastro_preco(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_preco")
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
		cmd = "insert into preco(preco) values ('" + r.Form.Get("preco") + "')"
		_, err := d.db.Exec(cmd)
		println(cmd)
		if err != nil {
			println(err)
		}
		println(r.Form.Get("preco") + " foi incluido!")
	}
	var where string
	where = ""
	if r.Form.Get("operation") == "Localizar" {
		if r.Form.Get("preco") != "" {
			where = " where preco like '" + r.Form.Get("preco") + "%'"
		}
	}
	cmd = "select * from (select row_number() over (order by id) rownum, id, preco, ts from preco order by id) " + where
	rows, err := d.db.Query(cmd)
	println(cmd)
	if err != nil {
		log.Fatal(err)
	}
	var row int32
	row = 0
	for rows.Next() {
		var id int
		var id_glp int
		var id_area int
		var preco float
		var ts string
		var rownum int
		err = rows.Scan(&rownum, &id, &id_glp, &id_area, &preco, &ts)
		if err != nil {
			println(err)
		}
		//fmt.Println(id, preco, ts)
		d.TabelaDados[row][0] = strconv.Itoa(rownum - 1)
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = strconv.Itoa(id_glp)
		d.TabelaDados[row][3] = strconv.Itoa(id_area)
		d.TabelaDados[row][4], err := strconv.ParseFloat(preco float, 32);
		if err != nil {
			println(err)
		}
		d.TabelaDados[row][5] = ts
		row++
		d.Tot_elementos = row
	}
	err = rows.Err()
	if err != nil {
		println(err)
	}
	defer rows.Close()
	for row := 0; int(row) < int(d.Tot_elementos); row++ {
		var id string
		var preco float
		var id_glp float
		var id_area float
		id = d.TabelaDados[row][1]
		preco = r.Form.Get("preco")
		id_glp = r.Form.Get("id_glp")
		id_area = r.Form.Get("id_area")
		if contains(radioSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Alterar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "update glp_preco_area set preco = '" + preco + "' ,ts = CURRENT_TIMESTAMP where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				println(d.TabelaDados[row][1] + " foi alterado de " + d.TabelaDados[row][2] + " para " + preco + "!")
				d.TabelaDados[row][2] = preco
			} else if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from glp_preco_area where id = " + id
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
				d.TabelaDados[row][4] = ""
				d.TabelaDados[row][5] = ""
			}
		}
		if contains(checkboxSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from glp_preco_area where id = " + id
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
				d.TabelaDados[row][4] = ""
				d.TabelaDados[row][5] = ""
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
