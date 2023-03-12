package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func cadastro_cliente(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_cliente")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))

	rows, err := d.db.Query("select id, nome, id_area, cep, endereco, ponto_de_referencia, ddi, ddd, coalesce(telefone,0), email, flg_aviso_gas_final, ts from cliente")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var nome string
		var id_area int
		var cep int
		var endereco string
		var ponto_de_referencia string
		var ddi int
		var ddd int
		var telefone int
		var email string
		var flg_aviso_gas_final int
		var ts string
		err = rows.Scan(&id, &nome, &id_area, &cep, &endereco, &ponto_de_referencia, &ddi, &ddd, &telefone, &email, &flg_aviso_gas_final, &ts)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, nome, id_area, cep, endereco, ponto_de_referencia, ddi, ddd, telefone, email, flg_aviso_gas_final, ts)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
