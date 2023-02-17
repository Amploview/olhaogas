package main

import (
	"net/http"
	"strings"
)

func cadastro_pedido(w http.ResponseWriter, r *http.Request, html string) {
	println("cadastro_pedido")
	println("glp : " + r.Form.Get("glp"))
	println("Operacao : " + r.Form.Get("operation"))
	println("Html : " + html)
	println("Action :" + strings.Trim(html, ".html"))
}
