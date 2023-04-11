package main

import (
	"net/http"
	"strconv"
	"strings"
)

func cadastro_glp(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_glp")
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
		if r.Form.Get("glp_default") == "1" {
			cmd = "update glp set glp_default = 0"
			_, err := d.db.Exec(cmd)
			println(cmd)
			if err != nil {
				println(err)
			}
			println("Efetuado reset do glp Default para " + r.Form.Get("descricao") + " se tornar o Default!")
		}
		cmd = "insert into glp(descricao, glp_default) values ('" + r.Form.Get("descricao") + "', " + r.Form.Get("glp_default") + ")"
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
		if r.Form.Get("descricao") != "" || r.Form.Get("glp_default") != "" {
			where = " where "
			if r.Form.Get("descricao") != "" {
				where = " where descricao like '" + r.Form.Get("descricao") + "%'"
			}
			if r.Form.Get("glp_default") != "" {
				if where != " where " {
					where = " and "
				}
				where = "glp_default = " + r.Form.Get("glp_default")
			}
		}
	}
	cmd = "select * from (select row_number() over (order by id) rownum, id, descricao, glp_default, ts from glp order by id) " + where
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
		var glp_default int
		var ts string
		var rownum int
		err = rows.Scan(&rownum, &id, &descricao, &glp_default, &ts)
		if err != nil {
			println(err)
		}
		//fmt.Println(id, descricao, ts)
		d.TabelaDados[row][0] = strconv.Itoa(rownum - 1)
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = descricao
		if glp_default == 1 {
			d.TabelaDados[row][3] = "1"
		} else {
			d.TabelaDados[row][3] = "0"
		}
		d.TabelaDados[row][4] = ts
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
		var descricao string
		var glp_default string
		id = d.TabelaDados[row][1]
		descricao = r.Form.Get("descricao")
		glp_default = r.Form.Get("glp_default")
		if contains(radioSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Alterar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				if r.Form.Get("glp_default") == "1" {
					cmd = "update glp set glp_default = 0"
					_, err := d.db.Exec(cmd)
					println(cmd)
					if err != nil {
						println(err)
					}
					println("Efetuado reset do glp Default para " + r.Form.Get("descricao") + " se tornar o Default!")
					d.TabelaDados[row][3] = "1"
				}
				cmd = "update glp set descricao = '" + descricao + "' , glp_default = " + glp_default + " ,ts = CURRENT_TIMESTAMP where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				println(d.TabelaDados[row][1] + " foi alterado de " + d.TabelaDados[row][2] + " para " + descricao + "!")
				d.TabelaDados[row][2] = descricao
			} else if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from glp where id = " + id
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
			}
		} else {
			if r.Form.Get("operation") == "Alterar" {
				if r.Form.Get("glp_default") == "1" {
					d.TabelaDados[row][3] = "0"
				}
			}
		}
		if contains(checkboxSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from glp where id = " + id
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
