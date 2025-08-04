/* 题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。 */

package main
import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"// Sqlx扩展库
	_ "github.com/mattn/go-sqlite3" // SQLite驱动
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func main() {
	// 连接数据库
	db, err := sqlx.Connect("sqlite3", "employees.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	createTableSQL := `CREATE TABLE IF NOT EXISTS employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		department TEXT,
		salary REAL
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}

	// 查询所有技术部员工
	var techEmployees []Employee
	err = db.Select(&techEmployees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("技术部员工信息：")
	for _, emp := range techEmployees {
		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 查询工资最高的员工
	var highestSalaryEmployee Employee
	err = db.Get(&highestSalaryEmployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("工资最高的员工信息：ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
		highestSalaryEmployee.ID, highestSalaryEmployee.Name,
		highestSalaryEmployee.Department, highestSalaryEmployee.Salary)
}