package sqltest

import (
	"fmt"

	"gorm.io/gorm"
)

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

type Employee1 struct {
	ID         int
	Name       string
	Department string
	Salary     int
}

func SqlCreateTable(db *gorm.DB) {
	db.AutoMigrate(&Employee{})
	employees := []Employee{
		{Name: "张三", Department: "技术部", Salary: 15000},
		{Name: "李四", Department: "技术部", Salary: 18000},
		{Name: "王五", Department: "技术部", Salary: 22000},
		{Name: "赵六", Department: "销售部", Salary: 12000},
		{Name: "钱七", Department: "技术部", Salary: 25000},
		{Name: "孙八", Department: "市场部", Salary: 13000},
		{Name: "周九", Department: "技术部", Salary: 19000},
	}
	result := db.CreateInBatches(employees, len(employees))
	fmt.Printf("分批批量创建成功: 新增了%d条记录\n", result.RowsAffected)

}

func SqlSelect(db *gorm.DB) {
	var employees []Employee
	db.Debug().Table("employees").Select("id", "name", "department", "salary").Where("salary > ?", 15000).Find(&employees)
}

func SqlDepartment(db *gorm.DB) {
	fmt.Println("查询技术部员工信息")
	var employees []Employee
	result := db.Debug().Raw("select id,name,department,salary from employees where department = ?", "技术部").Scan(&employees)
	fmt.Println("查询结果为：", result)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("开始循环输出结果")
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %d\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}
}

func MaxSalary(db *gorm.DB) {
	var employee Employee
	db.Debug().Raw("select id,name,department,salary from employees order by salary desc limit 1").Scan(&employee)
	fmt.Printf("最高工资员工: ID: %d, Name: %s, Department: %s, Salary: %d\n",
		employee.ID, employee.Name, employee.Department, employee.Salary)
}
