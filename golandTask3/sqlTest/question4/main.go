/* 题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。 */

package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"       // Sqlx扩展库
	_ "github.com/mattn/go-sqlite3" // SQLite驱动
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	// 这里可以添加代码来实现题目要求
	// 例如，连接数据库、创建表、插入数据等
	// 连接数据库
	db, err := sqlx.Connect("sqlite3", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	createTableSQL := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		price REAL
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}

	// 插入示例数据
	insertSQL := `INSERT INTO books (title, author, price) VALUES (?, ?, ?)`
	if _, err := db.Exec(insertSQL, "Go语言编程", "张三", 59.99); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(insertSQL, "Python编程", "李四", 45.00); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(insertSQL, "Java编程", "王五", 75.50); err != nil {
		log.Fatal(err)
	}
	// 查询所有书籍
	// var books []Book // Removed unused variable
	// 查询价格大于50元的书籍
	var expensiveBooks []Book
	err = db.Select(&expensiveBooks, "SELECT * FROM books WHERE price > ?", 50)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("价格大于50元的书籍信息：")
	for _, book := range expensiveBooks {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}
