package exercise1

import (
	"gorm.io/gorm"
)

// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、
// age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
type Student struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Age   int
	Grade string `gorm:"size:50"`
}

func Run(db *gorm.DB) {
	// db.AutoMigrate(&Student{})
	// student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	// result := db.Create(&student)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }

	// var students []Student
	// result := db.Where("age > ?", 18).Find(&students)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// for _, student := range students {
	// 	fmt.Printf("Name: %s, Age: %d, Grade: %s\n", student.Name, student.Age, student.Grade)
	// }

	// result := db.Model(&Student{}).
	// 	Where("name = ?", "张三").
	// 	Update("grade", "四年级")
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// student := Student{Name: "老五", Age: 12, Grade: "一年级"}
	// result := db.Create(&student)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }

	// result := db.Where("age < ?", 15).Delete(&Student{})
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// var students []Student
	// findResult := db.Find(&students)
	// if findResult.Error != nil {
	// 	panic(findResult.Error)
	// }
	// for _, student := range students {
	// 	fmt.Printf("Name: %s, Age: %d, Grade: %s\n", student.Name, student.Age, student.Grade)
	// }

}
