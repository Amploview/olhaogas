package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var porta = ":8080"

var host = "http://localhost" + porta
var title = "OlhaOgas"
var header = "Olhaogas | Gas online"

/* uso tmpl generico */
func genericExecTemplate(w http.ResponseWriter, r *http.Request, html string, tmplgeneric template.Template, d *data) {
	println("r.Host = '" + r.Host + "'")
	println("\n")
	//println("d.Record = '" + d.Record + "'")
	println("\n")
	if r.Host == "localhost:8080" {
		println("Execucao local, d.host não foi alterado!")
		println("\n")
	} else {
		println("Execucao não está local, entao reconfigurando d.Host para")
		d.Host = "https://" + r.Host
		println("'" + d.Host + "'")
		println("\n")
	}
	if err := tmplgeneric.ExecuteTemplate(w, html, d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintln(w, "Descricao : ", r.Form.Get("descricao"))
	//fmt.Fprintln(w, "Operacao : ", r.Form.Get("operation"))
	//fmt.Fprintln(w, "Html : ", html)
	//println("Descricao : " + r.Form.Get("descricao"))
	//println("Operacao : " + r.Form.Get("operation"))
	//println("Html : " + html)
	//println("Action :" + strings.Trim(html, ".html"))
	/*
		if r.Form.Get("operation") == "" {
			for i := 0; i < sizeRows; i++ {
				for j := 0; j < sizeCols; j++ {
					d.TabelaDados[i][j] = ""
				}
			}
		}
	*/
	switch strings.Trim(html, ".html") {
	case "olhaogas":
		println("é olhaogas")
	case "cadastro":
		println("é cadastros")
	case "relatorios":
		println("é relatorios")
	case "sobre":
		println("é sobre")
	case "cadastro_area":
		cadastro_area(w, r, html, d)
		fmt.Fprintf(w, "<input type=\"hidden\" name=\"reload\" value="+strconv.Itoa(d.Reload)+">")
		if d.Reload != 0 {
			d.Reload = -1
		}
	case "cadastro_glp":
		cadastro_glp(w, r, html, d)
		fmt.Fprintf(w, "<input type=\"hidden\" name=\"reload\" value="+strconv.Itoa(d.Reload)+">")
		if d.Reload != 0 {
			d.Reload = -1
		}
	case "cadastro_preco":
		cadastro_preco(w, r, html, d)
		fmt.Fprintf(w, "<input type=\"hidden\" name=\"reload\" value="+strconv.Itoa(d.Reload)+">")
		if d.Reload != 0 {
			d.Reload = -1
		}
	case "cadastro_cliente":
		cadastro_cliente(w, r, html, d)
	case "cadastro_cobertura":
		cadastro_cobertura(w, r, html, d)
	case "cadastro_usuario":
		cadastro_usuario(w, r, html, d)
	case "cadastro_pedido":
		cadastro_pedido(w, r, html, d)
	case "relatorio_area":
		relatorio_area(w, r, html, d)
	case "relatorio_glp":
		relatorio_glp(w, r, html, d)
	case "relatorio_preco":
		relatorio_preco(w, r, html, d)
	case "relatorio_cliente":
		relatorio_cliente(w, r, html, d)
	case "relatorio_pedido":
		relatorio_pedido(w, r, html, d)
		//case "relatorio_cobertura":
		//	relatorio_cobertura(w, r, html, d)
		//case "relatorio_usuario":
		//	relatorio_usuario(w, r, html, d)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "./data/olhaogas.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	//Declare Tabela
	var TabelaDados [sizeRows][sizeCols]string
	//Declare FKs
	var Area [sizeRows][3]string    //rownum, id, descricao
	var Glp [sizeRows][3]string     //rownum, id, descricao
	var Cliente [sizeRows][5]string //rownum, id, nome, telefone, email
	var Usuario [sizeRows][5]string //rownum, id, login, nome, tipo_usuario
	//Menu e demais
	data_olhaogas := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplolhaogas := template.Must(template.ParseFiles("templates/olhaogas.html"))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "olhaogas.html", *tmplolhaogas, &data_olhaogas)
	})
	data_cadastros := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplcadastros := template.Must(template.ParseFiles("templates/cadastros.html"))
	r.HandleFunc("/Cadastros", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "cadastros.html", *tmplcadastros, &data_cadastros)
	})
	data_relatorios := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplrelatorios := template.Must(template.ParseFiles("templates/relatorios.html"))
	r.HandleFunc("/Relatorios", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "relatorios.html", *tmplrelatorios, &data_relatorios)
	})
	data_sobre := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplsobre := template.Must(template.ParseFiles("templates/sobre.html"))
	r.HandleFunc("/Sobre", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "sobre.html", *tmplsobre, &data_sobre)
	})

	//Cadastros
	data_cadastroarea := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplcadastroarea := template.Must(template.ParseFiles("templates/cadastro_area.html"))
	r.HandleFunc("/CadastroArea", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "cadastro_area.html", *tmplcadastroarea, &data_cadastroarea)
	})
	data_cadastroglp := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplcadastroglp := template.Must(template.ParseFiles("templates/cadastro_glp.html"))
	r.HandleFunc("/CadastroGlp", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "cadastro_glp.html", *tmplcadastroglp, &data_cadastroglp)
	})
	data_cadastropreco := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplcadastropreco := template.Must(template.ParseFiles("templates/cadastro_preco.html"))
	r.HandleFunc("/CadastroPreco", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "cadastro_preco.html", *tmplcadastropreco, &data_cadastropreco)
	})
	data_cadastrocliente := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplcadastrocliente := template.Must(template.ParseFiles("templates/cadastro_cliente.html"))
	r.HandleFunc("/CadastroCliente", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "cadastro_cliente.html", *tmplcadastrocliente, &data_cadastrocliente)
	})
	data_cadastrocobertura := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplcadastrocobertura := template.Must(template.ParseFiles("templates/cadastro_cobertura.html"))
	r.HandleFunc("/CadastroCobertura", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "cadastro_cobertura.html", *tmplcadastrocobertura, &data_cadastrocobertura)
	})
	data_cadastrousuario := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplcadastrousuario := template.Must(template.ParseFiles("templates/cadastro_usuario.html"))
	r.HandleFunc("/CadastroUsuario", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "cadastro_usuario.html", *tmplcadastrousuario, &data_cadastrousuario)
	})
	data_cadastropedido := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplcadastropedido := template.Must(template.ParseFiles("templates/cadastro_pedido.html"))
	r.HandleFunc("/CadastroPedido", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "cadastro_pedido.html", *tmplcadastropedido, &data_cadastropedido)
	})

	//Relatorios
	data_relatorioarea := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplrelatorioarea := template.Must(template.ParseFiles("templates/relatorio_area.html"))
	r.HandleFunc("/RelatorioArea", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "relatorio_area.html", *tmplrelatorioarea, &data_relatorioarea)
	})
	data_relatorioglp := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplrelatorioglp := template.Must(template.ParseFiles("templates/relatorio_glp.html"))
	r.HandleFunc("/RelatorioGlp", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "relatorio_glp.html", *tmplrelatorioglp, &data_relatorioglp)
	})
	data_relatoriopreco := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplrelatoriopreco := template.Must(template.ParseFiles("templates/relatorio_preco.html"))
	r.HandleFunc("/RelatorioPreco", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "relatorio_preco.html", *tmplrelatoriopreco, &data_relatoriopreco)
	})
	data_relatoriocliente := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplrelatoriocliente := template.Must(template.ParseFiles("templates/relatorio_cliente.html"))
	r.HandleFunc("/RelatorioCliente", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "relatorio_cliente.html", *tmplrelatoriocliente, &data_relatoriocliente)
	})
	data_relatoriopedido := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	tmplrelatoriopedido := template.Must(template.ParseFiles("templates/relatorio_pedido.html"))
	r.HandleFunc("/RelatorioPedido", func(w http.ResponseWriter, r *http.Request) {
		genericExecTemplate(w, r, "relatorio_pedido.html", *tmplrelatoriopedido, &data_relatoriopedido)
	})
	//data_relatoriocobertura := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	//tmplrelatoriocobertura := template.Must(template.ParseFiles("templates/relatorio_cobertura.html"))
	//r.HandleFunc("/RelatorioCobertura", func(w http.ResponseWriter, r *http.Request) {
	//	genericExecTemplate(w, r, "cadastro_relatorioa.html", *tmplrelatoriocobertura, &data_relatoriocobertura)
	//})
	//data_relatoriousuario := data{title, header, host, db, TabelaDados, -1, Area, Glp, Cliente, Usuario}
	//tmplrelatoriousuario := template.Must(template.ParseFiles("templates/relatorio_usuario.html"))
	//r.HandleFunc("/RelatorioUsuario", func(w http.ResponseWriter, r *http.Request) {
	//	genericExecTemplate(w, r, "cadastro_relatoriohtml", *tmplrelatoriousuario, &data_relatoriousuario)
	//})

	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("templates/styles/"))))

	http.Handle("/", r)
	http.Handle("/Cadastros", r)
	http.Handle("/Relatorios", r)
	http.Handle("/Sobre", r)
	http.Handle("/CadastroArea", r)
	http.Handle("/CadastroGlp", r)
	http.Handle("/CadastroPreco", r)
	http.Handle("/CadastroCliente", r)
	http.Handle("/CadastroCobertura", r)
	http.Handle("/CadastroUsuario", r)
	http.Handle("/CadastroPedido", r)
	http.Handle("/RelatorioArea", r)
	http.Handle("/RelatorioGlp", r)
	http.Handle("/RelatorioPreco", r)
	http.Handle("/RelatorioCliente", r)
	http.Handle("/RelatorioPedido", r)
	//http.Handle("/RelatorioCobertura", r)
	//http.Handle("/RelatorioUsuario", r)
	log.Fatalln(http.ListenAndServe(porta, nil))
}
