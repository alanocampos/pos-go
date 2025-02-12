package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Categoria struct {
	ID       int `gorm:"primaryKey"`
	Nome     string
	Produtos []Produto `gorm:"many2many:produtos_categorias"`
}

type Produto struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categorias []Categoria `gorm:"many2many:produtos_categorias"`
	gorm.Model
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Produto{}, &Categoria{})

	//create categoria
	categoria := Categoria{Nome: "Cozinha"}
	db.Create(&categoria)

	//create categoria
	categoria2 := Categoria{Nome: "Eletronicos"}
	db.Create(&categoria2)

	// create produto
	db.Create(&Produto{
		Name:       "Panela",
		Price:      9000,
		Categorias: []Categoria{categoria, categoria2},
	})

	var categorias []Categoria
	err = db.Model(&Categoria{}).Preload("Produtos").Find(&categorias).Error
	if err != nil {
		panic(err)
	}
	for _, categoria := range categorias {
		fmt.Println(categoria.Nome, ":")
		for _, produto := range categoria.Produtos {
			println("-", produto.Name)
		}
	}
}
