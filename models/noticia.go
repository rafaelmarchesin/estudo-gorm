package models

import (
	"encoding/json"
	"log"
	"time"

	database "gorm_project/db"
)

type Noticia struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Titulo    string    `gorm:"size:255;not null;unique" json:"titulo"`
	Conteudo  string    `gorm:"size:255;not null;unique" json:"conteudo"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Tenta encontrar as notícias no Redis, se não encontrar, procura no MySQL
func (n *Noticia) BuscaTodasAsNoticias() ([]Noticia, error) {

	conexao := database.Connect()
	defer conexao.Close()

	noticias := []Noticia{}
	reply, err := database.Get("noticias")

	if err != nil {
		log.Println("Buscando no mysql")
		var err error
		noticias := []Noticia{}
		err = conexao.Debug().Model(&Noticia{}).Limit(100).Find(&noticias).Error
		if err != nil {
			return []Noticia{}, err
		}
		noticiasBytes, _ := json.Marshal(noticias)
		database.Set("noticias", noticiasBytes)
		return noticias, nil
	} else {
		log.Println("Buscando no redis")
		json.Unmarshal(reply, &noticias)
		return noticias, nil
	}
}

// Grava a notícia no MySQL e exclui a chave desatualizada do Redis
func (n *Noticia) AdicionaNoticia() (*Noticia, error) {

	conexao := database.Connect()
	defer conexao.Close()

	// var err error
	err := conexao.Debug().Model(&Noticia{}).Create(&n).Error
	if err != nil {
		return &Noticia{}, err
	}
	return n, nil
}
