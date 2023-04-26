package main

import (
	"net/http"
	"strconv"
	"strings"
)

func cadastro_usuario(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_usuario")
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
		cmd = "insert into usuario(login, senha, nome, flg_tipo_usuario) values ('" + r.Form.Get("login") + "', '" + r.Form.Get("senha") + "', '" + r.Form.Get("nome") + "', " + r.Form.Get("flg_tipo_usuario") + ")"
		_, err := d.db.Exec(cmd)
		println(cmd)
		if err != nil {
			println(err)
		}
		txt_flg_tipo_usuario := ""
		switch r.Form.Get("flg_tipo_usuario") {
		case "0":
			txt_flg_tipo_usuario = "Administrador"
			break
		case "1":
			txt_flg_tipo_usuario = "Light (Deposito e Motorista)"
			break
		case "2":
			txt_flg_tipo_usuario = "Deposito"
			break
		case "3":
			txt_flg_tipo_usuario = "Motorista"
			break
		default:
			txt_flg_tipo_usuario = "Erro usuario nao reconhecido"
		}
		println(txt_flg_tipo_usuario + "/" + r.Form.Get("login") + "/" + r.Form.Get("nome") + " foi incluido!")
	}
	var where string
	where = ""
	if r.Form.Get("operation") == "Localizar" {
		if r.Form.Get("nome") != "" {
			where += " where nome like '%" + r.Form.Get("nome") + "%'"
		}
		if r.Form.Get("login") != "" {
			if where != "" {
				where += " and "
			} else {
				where += " where "
			}
			where += "login like '%" + r.Form.Get("login") + "%'"
		}
		if r.Form.Get("flg_tipo_usuario") != "" {
			if where != "" {
				where += " and "
			} else {
				where += " where "
			}
			where += "flg_tipo_usuario = " + r.Form.Get("flg_tipo_usuario")
		}

	}
	cmd = "select * from (select row_number() over (order by usuario.id) rownum, usuario.id as id, login, nome, senha, flg_tipo_usuario, txt_flg_tipo_usuario.descricao as descricao_flg_tipo_usuario, usuario.ts as ts from usuario, (SELECT 0 as id, 'Administrador' as descricao union SELECT 1 as id, 'Light (Deposito e Motorista)' as descricao union SELECT 2 as id, 'Deposito' as descricao union SELECT 3 as id, 'Motorista' as descricao) txt_flg_tipo_usuario where usuario.flg_tipo_usuario = txt_flg_tipo_usuario.id order by usuario.id) " + where + " order by rownum desc"
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
		var login string
		var senha string
		var nome string
		var flg_tipo_usuario int
		var descricao_flg_tipo_usuario string
		var ts string
		var rownum int
		err = rows.Scan(&rownum, &id, &login, &senha, &nome, &flg_tipo_usuario, &descricao_flg_tipo_usuario, &ts)
		if err != nil {
			println(err)
		}
		d.TabelaDados[row][0] = strconv.Itoa(rownum - 1)
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = login
		d.TabelaDados[row][3] = senha
		d.TabelaDados[row][4] = nome
		d.TabelaDados[row][5] = strconv.Itoa(flg_tipo_usuario)
		d.TabelaDados[row][6] = descricao_flg_tipo_usuario
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
		var login string
		var senha string
		var nome string
		var flg_tipo_usuario string
		var descricao_flg_tipo_usuario string
		id = d.TabelaDados[row][1]
		login = r.Form.Get("login")
		senha = r.Form.Get("senha")
		nome = r.Form.Get("nome")
		flg_tipo_usuario = r.Form.Get("flg_tipo_usuario")
		descricao_flg_tipo_usuario = r.Form.Get("descricao_flg_tipo_usuario")
		if contains(radioSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Alterar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				println(d.TabelaDados[row][1] + " esta sendo alterado de " + d.TabelaDados[row][2] + "/" +
					d.TabelaDados[row][3] + "/" +
					d.TabelaDados[row][4] + "/" +
					d.TabelaDados[row][5] + "/" +
					d.TabelaDados[row][6] +
					" para " +
					login + "/" +
					senha + "/" +
					nome + "/" +
					flg_tipo_usuario + "/" +
					descricao_flg_tipo_usuario + "!")
				cmd = "update usuario set login = '" + login +
					"', senha = '" + senha +
					"', nome = '" + nome +
					"', flg_tipo_usuario = " + flg_tipo_usuario +
					", ts = CURRENT_TIMESTAMP where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				d.TabelaDados[row][2] = login
				d.TabelaDados[row][3] = senha
				d.TabelaDados[row][4] = nome
				d.TabelaDados[row][5] = flg_tipo_usuario
				d.TabelaDados[row][6] = descricao_flg_tipo_usuario
				row_sav = row
			} else if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from usuario where id = " + id
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
				cmd = "delete from usuario where id = " + id
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
	for i := 0; i < 4; i++ { //Incicialização tabelas fk Flg_Tipo_Usuario
		d.Flg_Tipo_Usuario[i][0] = "" //id
		d.Flg_Tipo_Usuario[i][1] = "" //descricao
	}
	cmd = "select * from (SELECT 0 as rownum, 0 as id, 'Administrador' as descricao union SELECT 1 as rownum, 1 as id, 'Light (Deposito e Motorista)' as descricao union SELECT 2 as rownum, 2 as id, 'Deposito' as descricao union SELECT 3 as rownum, 3 as id, 'Motorista' as descricao order by id)"
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	row = 0
	for rows.Next() {
		var flg_tipo_usuario int
		var descricao_flg_tipo_usuario string
		var rownum int
		err = rows.Scan(&rownum, &flg_tipo_usuario, &descricao_flg_tipo_usuario)
		if err != nil {
			println(err)
		}
		d.Flg_Tipo_Usuario[row][0] = strconv.Itoa(rownum - 1)
		d.Flg_Tipo_Usuario[row][1] = strconv.Itoa(flg_tipo_usuario)
		d.Flg_Tipo_Usuario[row][2] = descricao_flg_tipo_usuario
		if d.Flg_Tipo_Usuario[row][1] == r.Form.Get("flg_tipo_usuario") && row_sav > -1 {
			d.TabelaDados[row_sav][6] = descricao_flg_tipo_usuario
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
