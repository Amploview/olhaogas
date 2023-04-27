package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func cadastro_pedido(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_pedido")
	var cmd string
	r.ParseForm()
	checkboxSelected := r.Form["name_checkbox"]
	radioSelected := r.Form["name_radio"]
	println("Operacao         : " + r.Form.Get("operation"))
	println("Action           : " + strings.Trim(html, ".html"))
	//Aqui sera codificado a identificacao do cliente passada na URL para pegar o preco deste em sua area
	if r.Form.Get("id_cliente") != "" {
		d.id_cliente = r.Form.Get("id_cliente")
	} else if r.Form.Get("id_cliente_hidden") != "" {
		d.id_cliente = r.Form.Get("id_cliente_hidden")
	}
	if r.Form.Get("key_id") != "" {
		d.key_id = r.Form.Get("key_id")
	} else if r.Form.Get("key_id_hidden") != "" {
		d.key_id = r.Form.Get("key_id_hidden")
	}
	println("Cliente          : " + d.id_cliente)
	println("Key_id           : " + d.key_id)
	for i := 0; i < sizeRows; i++ {
		for j := 0; j < sizeCols; j++ {
			d.TabelaDados[i][j] = ""
		}
	}
	if r.Form.Get("operation") == "Inserir" {
		cmd = "insert into pedido(id_cliente, quantidade, id_glp_preco_area, valor_pedido, flg_prioridade_pedido, flg_hora) values (" + r.Form.Get("id_cliente_hidden") + ", " + r.Form.Get("quantidade") + ", " + r.Form.Get("id_glp_preco_area_hidden") + ", " + r.Form.Get("valor_pedido") + ", " + r.Form.Get("flg_prioridade_pedido") + ", " + r.Form.Get("flg_hora") + ")"
		_, err := d.db.Exec(cmd)
		println(cmd)
		if err != nil {
			println(err)
		}
		println("pedido " + r.Form.Get("id_glp") + "/" + r.Form.Get("id_cliente") + "/" + r.Form.Get("quantidade") + "/" + r.Form.Get("id_glp_preco_area") + "/" + r.Form.Get("valor_pedido") + "/" + r.Form.Get("flg_prioridade_pedido") + "/" + r.Form.Get("flg_hora") + " foi incluido!")
	}
	var where string
	where = ""
	if r.Form.Get("operation") == "Localizar" {
		if r.Form.Get("id_cliente") != "" {
			where += " where id_cliente = " + r.Form.Get("id_cliente")
		}
		if r.Form.Get("id_glp_preco") != "" {
			if r.Form.Get("id_glp_preco") != "" {
				where += " where id_glp_preco = " + r.Form.Get("id_glp_preco")
			}
		}
		if r.Form.Get("quantidade") != "" {
			if where != "" {
				where += " and "
			} else {
				where += " where "
			}
			where += "quantidade = " + r.Form.Get("quantidade")
		}
		if r.Form.Get("flg_prioridade_pedido") != "" {
			if where != "" {
				where += " and "
			} else {
				where += " where "
			}
			where += "flg_prioridade_pedido = " + r.Form.Get("flg_prioridade_pedido")
		}
		if r.Form.Get("flg_hora") != "" {
			if where != "" {
				where += " and "
			} else {
				where += " where "
			}
			where += "flg_hora = " + r.Form.Get("flg_hora")
		}

	}
	cmd = "select * from (select row_number() over (order by pedido.id) rownum, id_cliente, cliente.nome as nome_cliente, quantidade, id_glp_preco_area, glp_preco_area.preco as preco_area, valor_pedido, flg_prioridade_pedido, txt_flg_prioridade_pedido.descricao as descricao_flg_prioridade_pedido, flg_hora, txt_flg_hora.descricao as descricao_flg_hora, pedido.ts as ts from pedido, glp_preco_area, cliente, (SELECT 0 as id, 'Hoje' as descricao union SELECT 1 as id, 'Amanhã' as descricao) txt_flg_prioridade_pedido, (SELECT 0 as id, 'Hoje' as descricao union SELECT 1 as id, 'Amanhã' as descricao) txt_flg_hora where pedido.id_glp_preco_area = glp_preco_area.id and pedido.id_cliente = cliente.id and pedido.flg_prioridade_pedido = txt_flg_prioridade_pedido.id and pedido.flg_hora = txt_flg_hora.id order by pedido.id) " + where + " order by rownum desc"
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
		var id_cliente int
		var nome_cliente string
		var quantidade int
		var id_preco_area int
		var preco_area float32
		var valor_pedido float32
		var flg_prioridade_pedido int
		var descricao_flg_prioridade_pedido string
		var flg_hora int
		var descricao_flg_hora string
		var ts string
		var rownum int
		err = rows.Scan(&rownum, &id, &id_cliente, &nome_cliente, &quantidade, &id_preco_area, &preco_area, &valor_pedido, &flg_prioridade_pedido, &descricao_flg_prioridade_pedido, &flg_hora, &descricao_flg_hora, &ts)
		if err != nil {
			println(err)
		}
		d.TabelaDados[row][0] = strconv.Itoa(rownum - 1)
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = strconv.Itoa(id_cliente)
		d.TabelaDados[row][3] = nome_cliente
		d.TabelaDados[row][4] = strconv.Itoa(quantidade)
		d.TabelaDados[row][5] = strconv.Itoa(id_preco_area)
		d.TabelaDados[row][6] = fmt.Sprint(preco_area)
		d.TabelaDados[row][7] = fmt.Sprint(valor_pedido)
		d.TabelaDados[row][8] = strconv.Itoa(flg_prioridade_pedido)
		d.TabelaDados[row][9] = descricao_flg_prioridade_pedido
		d.TabelaDados[row][10] = strconv.Itoa(flg_hora)
		d.TabelaDados[row][11] = descricao_flg_hora
		d.TabelaDados[row][12] = ts
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
		var id_cliente string
		var quantidade string
		var id_preco_area string
		var valor_pedido string
		var flg_prioridade_pedido string
		var descricao_flg_prioridade_pedido string
		var flg_hora string
		var descricao_flg_hora string
		id = d.TabelaDados[row][1]
		quantidade = r.Form.Get("quantidade")
		flg_prioridade_pedido = r.Form.Get("flg_prioridade_pedido")
		descricao_flg_prioridade_pedido = r.Form.Get("descricao_flg_prioridade_pedido ")
		if contains(radioSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Alterar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				println(d.TabelaDados[row][1] + " esta sendo alterado de " + d.TabelaDados[row][3] + "/" +
					d.TabelaDados[row][4] + "/" +
					d.TabelaDados[row][6] + "/" +
					d.TabelaDados[row][7] + "/" +
					d.TabelaDados[row][9] + "/" +
					d.TabelaDados[row][11] +
					" para " +
					id_cliente + "/" +
					quantidade + "/" +
					id_preco_area + "/" +
					valor_pedido + "/" +
					descricao_flg_prioridade_pedido + "/" +
					descricao_flg_hora + "!")
				cmd = "update pedido set id_cliente = " + id_cliente +
					", quantidade = " + quantidade +
					", id_preco_area = " + id_preco_area +
					", valor_pedido = " + valor_pedido +
					", flg_prioridade_pedido = " + flg_prioridade_pedido +
					", flg_hora = " + flg_hora +
					", ts = CURRENT_TIMESTAMP where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				d.TabelaDados[row][2] = id_cliente
				d.TabelaDados[row][4] = quantidade
				d.TabelaDados[row][5] = id_preco_area
				d.TabelaDados[row][7] = valor_pedido
				d.TabelaDados[row][8] = flg_prioridade_pedido
				d.TabelaDados[row][10] = flg_hora
				row_sav = row
			} else if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from pedido where id = " + id
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
			}
		}
		if contains(checkboxSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Eliminar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "delete from pedido where id = " + id
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
			}
		}
	}
	println("Carregando FK's")
	for i := 0; i < sizeRows; i++ { //Incicialização tabelas fk Glp e area
		d.Glp[i][0] = ""     //id
		d.Glp[i][1] = ""     //descricao
		d.Cliente[i][0] = "" //id
		d.Cliente[i][1] = "" //nome
	}
	cmd = "select * from (select row_number() over (order by id) rownum, id as id_glp, descricao as descricao_glp from glp order by id)"
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	row = 0
	for rows.Next() {
		var id_glp int
		var descricao_glp string
		var rownum int
		err = rows.Scan(&rownum, &id_glp, &descricao_glp)
		if err != nil {
			println(err)
		}
		d.Glp[row][0] = strconv.Itoa(rownum - 1)
		d.Glp[row][1] = strconv.Itoa(id_glp)
		d.Glp[row][2] = descricao_glp
		if d.Glp[row][1] == r.Form.Get("id_glp") && row_sav > -1 {
			d.TabelaDados[row_sav][6] = descricao_glp
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
	cmd = "select * from (select row_number() over (order by id) rownum, id as id_cliente, nome as nome_cliente from cliente order by id)"
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	row = 0
	for rows.Next() {
		var id_cliente int
		var nome_cliente string
		var rownum int
		err = rows.Scan(&rownum, &id_cliente, &nome_cliente)
		if err != nil {
			println(err)
		}
		d.Cliente[row][0] = strconv.Itoa(rownum - 1)
		d.Cliente[row][1] = strconv.Itoa(id_cliente)
		d.Cliente[row][2] = nome_cliente
		if d.Cliente[row][1] == r.Form.Get("id_cliente") && row_sav > -1 {
			d.TabelaDados[row_sav][4] = nome_cliente
		}
		if err != nil {
			println(err)
		}
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
	for i := 0; i < sizeRows; i++ { //Incicialização tabelas fk GlpPrecoArea
		d.GlpPrecoArea[i][0] = "" //id
		d.GlpPrecoArea[i][1] = "" //id_area
		d.GlpPrecoArea[i][2] = "" //id_glp
		d.GlpPrecoArea[i][3] = "" //preco_area
	}
	cmd = "select * from (select row_number() over (order by id) rownum, id, id_area, id_glp, preco as preco_area from glp_preco_area order by id)"
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	row = 0
	for rows.Next() {
		var id int
		var id_area int
		var id_glp int
		var preco_area float32
		var rownum int
		err = rows.Scan(&rownum, &id, &id_area, &id_glp, &preco_area)
		if err != nil {
			println(err)
		}
		d.GlpPrecoArea[row][0] = strconv.Itoa(rownum - 1)
		d.GlpPrecoArea[row][1] = strconv.Itoa(id)
		d.GlpPrecoArea[row][2] = strconv.Itoa(id_area)
		d.GlpPrecoArea[row][3] = strconv.Itoa(id_glp)
		d.GlpPrecoArea[row][4] = fmt.Sprint(preco_area)
		if d.GlpPrecoArea[row][1] == r.Form.Get("id_preco_area") && row_sav > -1 {
			d.TabelaDados[row_sav][6] = fmt.Sprint(preco_area)
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
	for i := 0; i < 2; i++ { //Incicialização tabelas fk Flg_Prioridade_Pedido
		d.Flg_Prioridade_Pedido[i][0] = "" //id
		d.Flg_Prioridade_Pedido[i][1] = "" //descricao
	}
	cmd = "select * from (SELECT 0 as rownum, 0 as id, 'Hoje' as descricao union SELECT 1 as rownum, 1 as id, 'Amanhã' as descricao order by id)"
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	row = 0
	for rows.Next() {
		var flg_prioridade_pedido int
		var descricao_flg_prioridade_pedido string
		var rownum int
		err = rows.Scan(&rownum, &flg_prioridade_pedido, &descricao_flg_prioridade_pedido)
		if err != nil {
			println(err)
		}
		d.Flg_Prioridade_Pedido[row][0] = strconv.Itoa(rownum - 1)
		d.Flg_Prioridade_Pedido[row][1] = strconv.Itoa(flg_prioridade_pedido)
		d.Flg_Prioridade_Pedido[row][2] = descricao_flg_prioridade_pedido
		if d.Flg_Prioridade_Pedido[row][1] == r.Form.Get("flg_prioridade_pedido") && row_sav > -1 {
			d.TabelaDados[row_sav][9] = descricao_flg_prioridade_pedido
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
	for i := 0; i < 2; i++ { //Incicialização tabelas fk Flg_Hora
		d.Flg_Hora[i][0] = "" //id
		d.Flg_Hora[i][1] = "" //descricao
	}
	cmd = "select * from (SELECT 0 as rownum, 0 as id, 'Qualquer hora' as descricao union SELECT 1 as rownum, 1 as id, '8 hora' as descricao union SELECT 2 as rownum, 2 as id, '9 hora' as descricao union SELECT 3 as rownum, 3 as id, '10 horas' as descricao union SELECT 4 as rownum, 4 as id, '11 horas' as descricao union SELECT 5 as rownum, 5 as id, '12 horas' as descricao union SELECT 6 as rownum, 6 as id, '13 horas' as descricao union SELECT 7 as rownum, 7 as id, '14 horas' as descricao union SELECT 8 as rownum, 8 as id, '15 horas' as descricao union SELECT 9 as rownum, 9 as id, '16 horas' as descricao order by id)"
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		println(err)
	}
	row = 0
	for rows.Next() {
		var flg_hora int
		var descricao_flg_hora string
		var rownum int
		err = rows.Scan(&rownum, &flg_hora, &descricao_flg_hora)
		if err != nil {
			println(err)
		}
		d.Flg_Hora[row][0] = strconv.Itoa(rownum - 1)
		d.Flg_Hora[row][1] = strconv.Itoa(flg_hora)
		d.Flg_Hora[row][2] = descricao_flg_hora
		if d.Flg_Hora[row][1] == r.Form.Get("flg_hora") && row_sav > -1 {
			d.TabelaDados[row_sav][11] = descricao_flg_hora
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
