package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func cadastro_pedido(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("cadastro_pedido")
	println("glp : " + r.Form.Get("glp"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))

	rows, err := d.db.Query("select id, id_cliente, quantidade, id_glp_preco_area, valor_pedido, data_pedido, flg_prioridade_pedido, flg_origem_pedido, flg_hora, flg_status, data_status, ts from pedido")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var id_cliente int
		var quantidade int
		var id_glp_preco_area int
		var valor_pedido float64
		var data_pedido string
		var flg_prioridade_pedido int
		var flg_origem_pedido int
		var flg_hora int
		var flg_status int
		var data_status int
		var ts string
		err = rows.Scan(&id, &id_cliente, &quantidade, &id_glp_preco_area, &valor_pedido, &data_pedido, &flg_prioridade_pedido, &flg_origem_pedido, &flg_hora, &flg_status, &data_status, &ts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, id_cliente, quantidade, id_glp_preco_area, valor_pedido, data_pedido, flg_prioridade_pedido, flg_origem_pedido, flg_hora, flg_status, data_status, ts)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
