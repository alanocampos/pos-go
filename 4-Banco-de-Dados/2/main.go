package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Produto struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Produto{})

	//Create One
	//db.Create(&Produto{
	//	Name:  "Notebook",
	//	Price: 90,
	//})

	// create batch
	//produtos := []Produto{
	//	{Name: "Notebook", Price: 100},
	//	{Name: "Mouse", Price: 200},
	//	{Name: "Teclado", Price: 300},
	//}
	//db.Create(&produtos)

	// select one
	//var produto Produto
	//db.First(&produto, 1)
	//fmt.Println(produto)
	//db.First(&produto, "name = ?", "Mouse")
	//fmt.Println(produto)

	// select all
	//var produtos []Produto
	//db.Limit(2).Offset(2).Find(&produtos)
	//for _, produto := range produtos {
	//	fmt.Println(produto)
	//}

	//var produtos []Produto
	//db.Where("name like ?", "%book").Find(&produtos)
	//for _, produto := range produtos {
	//	fmt.Println(produto)
	//}

	//var p Produto
	//db.First(&p, 1)
	//p.Name = "New Mouse"
	//db.Save(&p)

	var p2 Produto
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)

}
