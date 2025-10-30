package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

type Product struct {
	Id           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	//var products []Product
	//db.Find(&products) // select * from
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//db.Where("price > ?", 2000).Find(&products) // select com where
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//db.Where("name LIKE ?", "%a%").Find(&products) // select com LIKE
	//for _, product := range products {
	//	fmt.Println(product)

	// update com upsert
	//var product Product
	//db.First(&product, 1) // acha o primeiro com a condição
	//product.Name = "Celular 1"
	//db.Save(&product) // save é um upsert

	//db.Model(&Product{}).Where("id = ?", &product.Id).Update("price", 250) // update de coluna específica

	//db.Delete(&product)

	//var product Product
	//db.First(&product, 1) // acha o primeiro com a condição
	//product.Name = "Celular 1"
	//db.Save(&product) // save é um upsert
	//db.Delete(&product)

	// create category
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	// create produto
	db.Create(&Product{
		Name:       "Celular",
		Price:      50,
		CategoryID: 1,
	})

	// create serial number
	db.Create(&SerialNumber{
		Number:    "123",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products) // belongs to é o preload
	for _, product := range products {
		fmt.Println(product.Name)
		fmt.Println(product.Price)
		fmt.Println(product.Category.Name)
		fmt.Println(product.SerialNumber.Number)
	}

}
