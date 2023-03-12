package main

import (
	"net/http"
	"strings"
)

func relatorio_preco(w http.ResponseWriter, r *http.Request, html string, d *data) {
	println("relatorio_preco")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))
}
