package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template

func main() {
	//confirmação de funcionamento no terminal
	fmt.Println("Ouvindo pela porta: 8080")
	//integração da página html no código
	templates = template.Must(template.ParseGlob("templates/*.html"))
	//endereçamento para a função de requisição
	http.HandleFunc("/", mainPage)
	//definição da porta do servidor
	http.ListenAndServe(":8080", nil)
}
func mainPage(w http.ResponseWriter, r *http.Request) {

	//requisição com form
	r.ParseForm()
	name := r.PostForm.Get("name")

	//executando o html que foi integrado na função principal
	templates.ExecuteTemplate(w, "index.html", struct {
		Name string
	}{
		Name: name,
	})
}
