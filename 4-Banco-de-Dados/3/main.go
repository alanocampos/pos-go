package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Categoria struct {
	ID   int `gorm:"primaryKey"`
	Nome string
}

type Produto struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Price       float64
	CategoriaID int
	Categoria   Categoria
	gorm.Model
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Produto{})

	// create categoria
	//categoria := Categoria{Nome: "Eletronicos"}
	//db.Create(&categoria)
	//
	//// create produto
	//db.Create(&Produto{
	//	Name:        "Notebook",
	//	Price:       9000,
	//	CategoriaID: categoria.ID,
	//})

	// create produto
	db.Create(&Produto{
		Name:        "Mouse",
		Price:       9000,
		CategoriaID: 2,
	})

	var produtos []Produto
	db.Preload("Categoria").Find(&produtos)
	for _, produto := range produtos {
		fmt.Println(produto.Name, produto.Categoria.Nome)
	}
}
