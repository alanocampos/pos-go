package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	tx := db.Begin()
	var c Categoria
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Nome = "Electronic"
	tx.Debug().Save(&c)
	tx.Commit()
}
