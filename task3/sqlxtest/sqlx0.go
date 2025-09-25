package sqlxtest

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func RunDepartment(db *sqlx.DB) {
	var employees []Employee
	sqlStr := "SELECT id, name, department, salary FROM employees WHERE department = ?"
	err := db.Select(&employees, sqlStr, "技术部")
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	// 输出结果
	fmt.Printf("技术部员工共%d人:\n", len(employees))
	for _, emp := range employees {
		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 薪资: %.2f\n",
			emp.ID, emp.Name, emp.Department, emp.Salary)
	}
}

func RunMaxSalary(db *sqlx.DB) {
	var employee Employee
	sqlStr := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(&employee, sqlStr)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}
	fmt.Printf("最高薪资的员工为: ID: %d, 姓名: %s, 部门: %s, 薪资: %.2f\n",
		employee.ID, employee.Name, employee.Department, employee.Salary)
}
