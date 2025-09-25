package sqlxtest

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     uint
	Title  string
	Author string
	Price  float64
}

func RunQueryBooks(db *sqlx.DB) {
	var books []Book
	sqlStr := "select id,title,author,price from books where price > ?"
	err := db.Select(&books, sqlStr, 50.0)
	if err != nil {
		fmt.Printf("Failed to query books:%d \n", err)
	}
	for _, book := range books {
		fmt.Printf("id: %d,Title: %s, Author: %s, Price: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}
