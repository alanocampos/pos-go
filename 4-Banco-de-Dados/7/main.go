package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Categoria struct {
	ID       int `gorm:"primaryKey"`
	Nome     string
	Produtos []Produto
}

type Produto struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoriaID  int
	Categoria    Categoria
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProdutoID int
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Produto{}, &Categoria{}, &SerialNumber{})

	////create categoria
	//categoria := Categoria{Nome: "Cozinha"}
	//db.Create(&categoria)
	//
	//// create produto
	//db.Create(&Produto{
	//	Name:        "Panela",
	//	Price:       9000,
	//	CategoriaID: 1,
	//})
	//
	//// create Serial Number
	//db.Create(&SerialNumber{
	//	Number:    "12345",
	//	ProdutoID: 1,
	//})

	var categorias []Categoria
	err = db.Model(&Categoria{}).Preload("Produtos.SerialNumber").Find(&categorias).Error
	if err != nil {
		panic(err)
	}
	for _, categoria := range categorias {
		fmt.Println(categoria.Nome, ":")
		for _, produto := range categoria.Produtos {
			println("-", produto.Name, "- Serial Number: ", produto.SerialNumber.Number)
		}
	}
}
