package main

import (
	"net/http"
	"strings"
)

func relatorio_area(w http.ResponseWriter, r *http.Request, html string) {
	println("relatorio_area")
	println("Descricao : " + r.Form.Get("descricao"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))
}
