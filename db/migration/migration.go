package migration

import (
	"gorm_project/db"
	"gorm_project/models"
)

// AutoMigration cria a tabela no Banco de Dados se ela não existir
func AutoMigration() {
	db := db.Connect()
	defer db.Close()
	db.AutoMigrate(models.Noticia{})
}
