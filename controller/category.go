//Not used yet
package controller

import (
	"github.com/afnscbrl/Golang_Rest_API/models"
	"gorm.io/gorm"
)

type CategoriaController struct {
	Db *gorm.DB
}

func NewCategoria(db *gorm.DB) *CategoriaController {
	return &CategoriaController{Db: db}
}

func (c *CategoriaController) SeedCategorias() {
	c.Db.Model(&models.Category{}).Create([]map[string]interface{}{
		{"nome": "Alimentação"},
		{"nome": "Saúde"},
		{"nome": "Moradia"},
		{"nome": "Transporte"},
		{"nome": "Educação"},
		{"nome": "Lazer"},
		{"nome": "Imprevistos"},
		{"nome": "Outras"},
	})
}
