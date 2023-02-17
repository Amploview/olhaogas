package main

import (
	"net/http"
	"strings"
)

func relatorio_usuario(w http.ResponseWriter, r *http.Request, html string) {
	println("relatorio_usuario")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))
}
