/* 题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。 */
package main
import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3" // SQLite驱动
	_ "gorm.io/driver/sqlite" // GORM SQLite驱动
)

func main() {
	// 连接数据库
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	createTableSQL := `CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER,
		grade TEXT
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}

	// 插入新记录
	insertSQL := `INSERT INTO students (name, age, grade) VALUES (?, ?, ?)`
	if _, err := db.Exec(insertSQL, "张三", 20, "三年级"); err != nil {
		log.Fatal(err)
	}

	// 查询年龄大于18岁的学生信息
	querySQL := `SELECT * FROM students WHERE age > ?`
	rows, err := db.Query(querySQL, 18)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("年龄大于18岁的学生信息：")
	for rows.Next() {
		var id int
		var name string
		var age int
		var grade string
		if err := rows.Scan(&id, &name, &age, &grade); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n", id, name, age, grade)
	}

	// 更新姓名为"张三"的学生年级为"四年级"
	updateSQL := `UPDATE students SET grade = ? WHERE name = ?`
	if _, err := db.Exec(updateSQL, "四年级", "张三"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("已将张三的年级更新为四年级。")

	// 删除年龄小于15岁的学生记录
	deleteSQL := `DELETE FROM students WHERE age < ?`
	if _, err := db.Exec(deleteSQL, 15); err != nil {
		log.Fatal(err)
	}

	fmt.Println("已删除年龄小于15岁的学生记录。")
}

	