package sqltest

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	ID     uint
	Title  string
	Author string
	Price  float64
}

func SqlCreate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
	books := []Book{
		{Title: "Go编程语言", Author: "Alan A. A. Donovan", Price: 75.50},
		{Title: "Go语言实战", Author: "William Kennedy", Price: 45.00},
		{Title: "Go Web编程", Author: "Sau Sheong Chang", Price: 60.00},
		{Title: "分布式系统", Author: "Martin Kleppmann", Price: 89.90},
		{Title: "算法导论", Author: "Thomas H. Cormen", Price: 120.00},
		{Title: "数据结构", Author: "Mark Allen Weiss", Price: 40.00},
	}

	result := db.Create(&books)
	if result.Error != nil {
		fmt.Printf("Failed to insert books:%v\n", result.Error)
	}

	fmt.Printf("Inserted %d books successfully\n", result.RowsAffected)
}

func PriceSelect(db *gorm.DB, price float64) {
	var books []Book
	result := db.Where("price > ?", price).Find(&books)
	if result.Error != nil {
		fmt.Printf("Failed to query books:%v\n", result.Error)
	}
	fmt.Printf("Found %d books with price > %.2f\n", result.RowsAffected, price)
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Price: %.2f\n", book.Title, book.Author, book.Price)
	}
}

func QueryBooks(db *gorm.DB) {
	var books []Book
	query := "select id,title,author,price from books where price > ?"
	result := db.Raw(query, 50).Scan(&books)
	if result.Error != nil {
		fmt.Printf("Failed to query books:%v\n", result.Error)
	}
	fmt.Printf("Found %d books with price > 50\n", result.RowsAffected)
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Price: %.2f\n", book.Title, book.Author, book.Price)
	}

}
