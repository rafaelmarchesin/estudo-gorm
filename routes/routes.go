package routes

import (
	"net/http"

	"gorm_project/controllers"
)

// CarregaRotas : carrega as rotas da aplicação
func CarregaRotas() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/cria-noticia", controllers.NovaNoticia)
	http.HandleFunc("/formulario-noticia", controllers.Formulario)
}
