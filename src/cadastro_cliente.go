package main

import (
	"net/http"
	"strconv"
	"strings"
)

func cadastro_cliente(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_cliente")
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
	println("Carregando FK's")
	for i := 0; i < sizeRows; i++ { //Incicialização tabelas fk Area
		d.Area[i][0] = "" //id
		d.Area[i][1] = "" //descricao
	}
	cmd = "select * from (select row_number() over (order by id) rownum, id as id_area, descricao as descricao_area from area order by id)"
	rows, err := d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	var row int32
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
		cmd = "insert into Cliente(nome) values ('" + r.Form.Get("nome") + "')"
		_, err := d.db.Exec(cmd)
		println(cmd)
		if err != nil {
			println(err)
		}
		println(r.Form.Get("nome") + " foi incluido!")
	}
	var where string
	where = ""
	if r.Form.Get("operation") == "Localizar" {
		if r.Form.Get("nome") != "" {
			where += " where nome like '" + r.Form.Get("nome") + "%'"
		}
	}
	cmd = "select * from (select row_number() over (order by cliente.id) rownum, cliente.id as id, nome, id_area, area.descricao as descricao_area, cep, endereco, ponto_de_referencia, ddi, ddd, telefone, email, flg_aviso_gas_final, cliente.ts as ts from cliente, area where cliente.id_area = area.id order by cliente.id) " + where
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	row = 0
	for rows.Next() {
		var id int
		var nome string
		var id_area int
		var descricao_area string
		var cep int
		var endereco string
		var ponto_de_referencia string
		var ddi int
		var ddd int
		var telefone int
		var email string
		var flg_aviso_gas_final int
		var ts string
		var rownum int
		err = rows.Scan(&rownum, &id, &nome, &id_area, &descricao_area, &cep, &endereco, &ponto_de_referencia, &ddi, &ddd, &telefone, &email, &flg_aviso_gas_final, &ts)
		if err != nil {
			println(err)
		}
		d.TabelaDados[row][0] = strconv.Itoa(rownum - 1)
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = nome
		d.TabelaDados[row][3] = strconv.Itoa(id_area)
		d.TabelaDados[row][4] = descricao_area
		d.TabelaDados[row][5] = strconv.Itoa(cep)
		d.TabelaDados[row][6] = endereco
		d.TabelaDados[row][7] = ponto_de_referencia
		d.TabelaDados[row][8] = strconv.Itoa(ddi)
		d.TabelaDados[row][9] = strconv.Itoa(ddd)
		d.TabelaDados[row][10] = strconv.Itoa(telefone)
		d.TabelaDados[row][11] = email
		d.TabelaDados[row][12] = strconv.Itoa(flg_aviso_gas_final)
		d.TabelaDados[row][13] = ts
		row++
		if row >= sizeRows {
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
		var nome string
		var id_area string
		var descricao_area string
		var cep string
		var endereco string
		var ponto_de_referencia string
		var ddi string
		var ddd string
		var telefone string
		var email string
		var flg_aviso_gas_final string
		id = d.TabelaDados[row][1]
		nome = r.Form.Get("nome")
		id_area = r.Form.Get("id_area")
		descricao_area = r.Form.Get("descricao_area")
		cep = r.Form.Get("cep")
		endereco = r.Form.Get("endereco")
		ponto_de_referencia = r.Form.Get("ponto_de_referencia")
		ddi = r.Form.Get("ddi")
		ddd = r.Form.Get("ddd")
		telefone = r.Form.Get("telefone")
		email = r.Form.Get("email")
		flg_aviso_gas_final = r.Form.Get("flg_aviso_gas_final")
		if contains(radioSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Alterar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "update cliente set nome = '" + nome +
					", id_area = " + id_area +
					", cep = " + cep +
					", endereco = '" + endereco +
					"' , ponto_de_referencia = '" + ponto_de_referencia +
					"' , ddi = " + ddi +
					", ddd = " + ddd +
					", telefone = " + telefone +
					", email = '" + email +
					"' , flg_aviso_gas_final = " + flg_aviso_gas_final +
					" ,ts = CURRENT_TIMESTAMP where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				println(d.TabelaDados[row][1] + " foi alterado de " + d.TabelaDados[row][2] + "/" +
					d.TabelaDados[row][3] + "/" +
					d.TabelaDados[row][4] + "/" +
					d.TabelaDados[row][5] + "/" +
					d.TabelaDados[row][6] + "/" +
					d.TabelaDados[row][7] + "/" +
					d.TabelaDados[row][8] + "/" +
					d.TabelaDados[row][9] + "/" +
					d.TabelaDados[row][10] + "/" +
					d.TabelaDados[row][11] + "/" +
					d.TabelaDados[row][12] +
					" para " +
					nome + "/" +
					id_area + "/" +
					cep + "/" +
					endereco + "/" +
					ponto_de_referencia + "/" +
					ddi + "/" +
					ddd + "/" +
					telefone + "/" +
					email + "/" +
					flg_aviso_gas_final + "!")
				d.TabelaDados[row][2] = nome
				d.TabelaDados[row][3] = id_area
				d.TabelaDados[row][4] = descricao_area
				d.TabelaDados[row][5] = cep
				d.TabelaDados[row][6] = endereco
				d.TabelaDados[row][7] = ponto_de_referencia
				d.TabelaDados[row][8] = ddi
				d.TabelaDados[row][9] = ddd
				d.TabelaDados[row][10] = telefone
				d.TabelaDados[row][11] = email
				d.TabelaDados[row][12] = flg_aviso_gas_final
			} else if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from cliente where id = " + id
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
				d.TabelaDados[row][8] = ""
				d.TabelaDados[row][9] = ""
				d.TabelaDados[row][10] = ""
				d.TabelaDados[row][11] = ""
				d.TabelaDados[row][12] = ""
				d.TabelaDados[row][13] = ""
			}
		}
		if contains(checkboxSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from cliente where id = " + id
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
				d.TabelaDados[row][8] = ""
				d.TabelaDados[row][9] = ""
				d.TabelaDados[row][10] = ""
				d.TabelaDados[row][11] = ""
				d.TabelaDados[row][12] = ""
				d.TabelaDados[row][13] = ""
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
