package main

import (
	"net/http"
	"strings"
)

func relatorio_pedido(w http.ResponseWriter, r *http.Request, html string) {
	println("relatorio_pedido")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))
}
