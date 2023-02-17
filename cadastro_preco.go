package main

import (
	"net/http"
	"strings"
)

func cadastro_preco(w http.ResponseWriter, r *http.Request, html string) {
	println("cadastro_preco")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))
}
