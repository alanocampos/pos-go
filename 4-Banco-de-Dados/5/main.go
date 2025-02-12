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

	//create categoria
	categoria := Categoria{Nome: "Eletronicos"}
	db.Create(&categoria)

	// create produto
	db.Create(&Produto{
		Name:        "Notebook",
		Price:       9000,
		CategoriaID: 1,
	})

	// create Serial Number
	db.Create(&SerialNumber{
		Number:    "12345",
		ProdutoID: 1,
	})

	var produtos []Produto
	db.Preload("Categoria").Preload("SerialNumber").Find(&produtos)
	for _, produto := range produtos {
		fmt.Println(produto.Name, produto.Categoria.Nome, produto.SerialNumber.Number)
	}
}
