package main

import (
	"net/http"

	"gorm_project/db/migration"
	"gorm_project/routes"
)

func main() {

	migration.AutoMigration()
	routes.CarregaRotas()

	http.ListenAndServe(":8080", nil)
}
