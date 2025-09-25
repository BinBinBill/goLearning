package main

import (
	"task3/gormtest"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// exercise1.Run(db)
	// exercise1.Run1(db)
	// exercise1.TransferMoney(db, 1, 2, 200)

	// sqltest.SqlCreateTable(db)
	// sqltest.SqlDepartment(db)
	// sqltest.MaxSalary(db)

	// sqltest.SqlCreate(db)
	// sqltest.PriceSelect(db, 50)

	// sqltest.QueryBooks(db)

	// sqlx
	// dsn := "root:root@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True"
	// db, err := sqlx.Connect("mysql", dsn)
	// if err != nil {
	// 	log.Fatalf("数据库连接失败: %v", err)
	// }
	// defer db.Close()
	// // sqlxtest.RunDepartment(db)
	// // sqlxtest.RunMaxSalary(db)
	// sqlxtest.RunQueryBooks(db)

	//gorm
	// gormtest.RunCreate(db)
	// gormtest.RunQueryByUserName(db, "Alice")
	// mostCommentedPost, _ := gormtest.GetMostCommentedPost(db)
	// fmt.Printf("评论最多的文章: %s\n", mostCommentedPost.Title)

	gormtest.DataCreate(db)

}
