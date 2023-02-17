package main

import (
	"net/http"
	"strings"
)

func relatorio_cobertura(w http.ResponseWriter, r *http.Request, html string) {
	println("relatorio_cobertura")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))
}
