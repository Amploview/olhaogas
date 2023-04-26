package main

import (
	"net/http"
	"strconv"
	"strings"
)

func cadastro_cobertura(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_cobertura")
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
		cmd = "insert into cobertura(id_usuario, id_area) values (" + r.Form.Get("id_usuario") + ", " + r.Form.Get("id_area") + ")"
		_, err := d.db.Exec(cmd)
		println(cmd)
		if err != nil {
			println(err)
		}
		println(r.Form.Get("id_usuario") + "/" + r.Form.Get("id_area") + " foi incluido!")
	}
	var where string
	where = ""
	if r.Form.Get("operation") == "Localizar" {
		if r.Form.Get("id_usuario") != "" {
			where += " where id_usuario = " + r.Form.Get("id_motorista")
		}
		if r.Form.Get("id_area") != "" {
			if where != "" {
				where += " and "
			} else {
				where += " where "
			}
			where += "id_area = " + r.Form.Get("id_area")
		}
	}
	cmd = "select * from (select row_number() over (order by cobertura.id) rownum, cobertura.id as id, id_usuario, usuario.login as login_motorista, usuario.nome as nome_motorista, id_area, area.descricao as descricao_area, cobertura.ts as ts from cobertura, usuario, area where cobertura.id_area = area.id and cobertura.id_usuario = usuario.id and usuario.flg_tipo_usuario = 3 order by cobertura.id) " + where + " order by rownum desc"
	rows, err := d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	var row_sav int
	row_sav = -1
	var row int32
	row = 0
	for rows.Next() {
		var id int
		var id_usuario int
		var login_motorista string
		var nome_motorista string
		var id_area int
		var descricao_area string
		var ts string
		var rownum int
		err = rows.Scan(&rownum, &id, &id_usuario, &login_motorista, &nome_motorista, &id_area, &descricao_area, &ts)
		if err != nil {
			println(err)
		}
		d.TabelaDados[row][0] = strconv.Itoa(rownum - 1)
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = strconv.Itoa(id_usuario)
		d.TabelaDados[row][3] = login_motorista
		d.TabelaDados[row][4] = nome_motorista
		d.TabelaDados[row][5] = strconv.Itoa(id_area)
		d.TabelaDados[row][6] = descricao_area
		d.TabelaDados[row][7] = ts
		row++
		if row >= sizeRows {
			rows.Close()
			break
		}
	}
	err = rows.Err()
	if err != nil {
		println(err)
	}
	defer rows.Close()
	var Tot_elementos = row
	for row := 0; int(row) < int(Tot_elementos); row++ {
		var id string
		var id_usuario string
		var id_area string
		id = d.TabelaDados[row][1]
		id_usuario = r.Form.Get("id_usuario")
		id_area = r.Form.Get("id_area")
		if contains(radioSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Alterar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				println(d.TabelaDados[row][1] + " esta sendo alterado de " + d.TabelaDados[row][2] + "/" +
					d.TabelaDados[row][5] +
					" para " +
					id_usuario + "/" +
					id_area + "!")
				cmd = "update cobertura set id_usuario = " + id_usuario +
					", id_area = " + id_area +
					", ts = CURRENT_TIMESTAMP where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				d.TabelaDados[row][2] = id_usuario
				d.TabelaDados[row][5] = id_area
				row_sav = row
			} else if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from cobertura where id = " + id
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
				d.TabelaDados[row][6] = ""
				d.TabelaDados[row][7] = ""
			}
		}
		if contains(checkboxSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from cobertura where id = " + id
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
				d.TabelaDados[row][6] = ""
				d.TabelaDados[row][7] = ""
			}
		}
	}
	println("Carregando FK's")
	for i := 0; i < sizeRows; i++ { //Incicialização tabelas fk Area
		d.Usuario[i][0] = "" //id
		d.Usuario[i][1] = "" //login
		d.Usuario[i][2] = "" //nome
		d.Area[i][0] = ""    //id
		d.Area[i][1] = ""    //descricao
	}
	cmd = "select * from (select row_number() over (order by id) rownum, id as id_usuario, login as login_usuario, nome as nome_usuario from usuario where flg_tipo_usuario = 3 order by id)"
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	row = 0
	for rows.Next() {
		var id_usuario int
		var login_usuario string
		var nome_usuario string
		var rownum int
		err = rows.Scan(&rownum, &id_usuario, &login_usuario, &nome_usuario)
		if err != nil {
			println(err)
		}
		d.Usuario[row][0] = strconv.Itoa(rownum - 1)
		d.Usuario[row][1] = strconv.Itoa(id_usuario)
		d.Usuario[row][2] = login_usuario
		if d.Usuario[row][1] == r.Form.Get("id_usuario") && row_sav > -1 {
			d.TabelaDados[row_sav][3] = login_usuario
			d.TabelaDados[row_sav][4] = nome_usuario
		}
		if err != nil {
			println(err)
		}
		row++
	}
	err = rows.Err()
	if err != nil {
		println(err)
	}
	defer rows.Close()
	cmd = "select * from (select row_number() over (order by id) rownum, id as id_area, descricao as descricao_area from area order by id)"
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	row = 0
	for rows.Next() {
		var id_area int
		var descricao_area string
		var rownum int
		err = rows.Scan(&rownum, &id_area, &descricao_area)
		if err != nil {
			println(err)
		}
		d.Area[row][0] = strconv.Itoa(rownum - 1)
		d.Area[row][1] = strconv.Itoa(id_area)
		d.Area[row][2] = descricao_area
		if d.Area[row][1] == r.Form.Get("id_area") && row_sav > -1 {
			d.TabelaDados[row_sav][6] = descricao_area
		}
		if err != nil {
			println(err)
		}
		row++
	}
	err = rows.Err()
	if err != nil {
		println(err)
	}
	defer rows.Close()
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
