package main

import (
	"fmt"
	"log"
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
	println("Carregando FK's")
	for i := 0; i < sizeRows; i++ { //Incicialização tabelas fk Glp e area
		d.Glp[i][0] = ""  //id
		d.Glp[i][1] = ""  //descricao
		d.Area[i][0] = "" //id
		d.Area[i][1] = "" //descricao
	}
	cmd = "select * from (select row_number() over (order by id) rownum, id as id_glp, descricao as descricao_glp from glp order by id)"
	rows, err := d.db.Query(cmd)
	println(cmd)
	if err != nil {
		log.Fatal(err)
	}
	var row int32
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
		log.Fatal(err)
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
		cmd = "insert into glp_preco_area(preco, id_glp, id_area) values ('" + r.Form.Get("preco") + ", " + r.Form.Get("id_glp") + ", " + r.Form.Get("id_area") + "')"
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
			where = " where preco = '" + r.Form.Get("preco")
		}
		if r.Form.Get("id_glp") != "" {
			if where != "" {
				where = " and "
			} else {
				where = " where "
			}
			where = "id_glp = " + r.Form.Get("id_glp")
		}
		if r.Form.Get("id_area") != "" {
			if where != "" {
				where = " and "
			} else {
				where = " where "
			}
			where = "id_area = " + r.Form.Get("id_area")
		}
	}
	cmd = "select * from (select row_number() over (order by glp_preco_area.id) rownum, glp_preco_area.id as id, preco, id_glp, glp.descricao as descricao_glp, id_area, area.descricao as descricao_area, glp_preco_area.ts as ts from glp_preco_area, glp, area where glp_preco_area.id_glp = glp.id and glp_preco_area.id_area = area.id order by glp_preco_area.id) " + where
	rows, err = d.db.Query(cmd)
	println(cmd)
	if err != nil {
		log.Fatal(err)
	}
	row = 0
	for rows.Next() {
		var id int
		var preco float32
		var id_glp int
		var descricao_glp string
		var id_area int
		var descricao_area string
		var ts string
		var rownum int
		err = rows.Scan(&rownum, &id, &preco, &id_glp, &descricao_glp, &id_area, &descricao_area, &ts)
		if err != nil {
			println(err)
		}
		//fmt.Println(id, preco, ts)
		d.TabelaDados[row][0] = strconv.Itoa(rownum - 1)
		d.TabelaDados[row][1] = strconv.Itoa(id)
		d.TabelaDados[row][2] = fmt.Sprint(preco)
		d.TabelaDados[row][3] = strconv.Itoa(id_glp)
		d.TabelaDados[row][4] = descricao_glp
		d.TabelaDados[row][5] = strconv.Itoa(id_area)
		d.TabelaDados[row][6] = descricao_area
		if err != nil {
			println(err)
		}
		d.TabelaDados[row][7] = ts
		row++
	}
	err = rows.Err()
	if err != nil {
		println(err)
	}
	defer rows.Close()
	var Tot_elementos = row
	for row := 0; int(row) < int(Tot_elementos); row++ {
		var id string
		var preco string
		var id_glp string
		var id_area string
		id = d.TabelaDados[row][1]
		preco = r.Form.Get("preco")
		id_glp = r.Form.Get("id_glp")
		id_area = r.Form.Get("id_area")
		if contains(radioSelected, d.TabelaDados[row][0]) {
			if r.Form.Get("operation") == "Alterar" {
				println(d.TabelaDados[row][1] + " foi selecionado!")
				cmd = "update glp_preco_area set preco = '" + preco + "' , id_glp = " + id_glp + ", id_area = " + id_area + ", ts = CURRENT_TIMESTAMP where id = " + id
				_, err := d.db.Exec(cmd)
				println(cmd)
				if err != nil {
					println(err)
				}
				println(d.TabelaDados[row][1] + " foi alterado de " + d.TabelaDados[row][2] + "/" + d.TabelaDados[row][3] + "/" + d.TabelaDados[row][5] + " para " + id_glp + "/" + id_area + "/" + preco + "!")
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
				d.TabelaDados[row][6] = ""
				d.TabelaDados[row][7] = ""
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
				d.TabelaDados[row][6] = ""
				d.TabelaDados[row][7] = ""
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
