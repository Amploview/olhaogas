package main

import (
	"net/http"
	"strings"
)

func cadastro_cobertura(w http.ResponseWriter, r *http.Request, html string) {
	println("cadastro_cobertura")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))
}
