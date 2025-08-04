/* 题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
 
 */
package main
import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3" // SQLite驱动
)

func main() {
	// 连接数据库
	db, err := sql.Open("sqlite3", "bank.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	createAccountsSQL := `CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		balance INTEGER
	);`
	createTransactionsSQL := `CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		from_account_id INTEGER,
		to_account_id INTEGER,
		amount INTEGER,
		FOREIGN KEY(from_account_id) REFERENCES accounts(id),
		FOREIGN KEY(to_account_id) REFERENCES accounts(id)
	);`
	if _, err := db.Exec(createAccountsSQL); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(createTransactionsSQL); err != nil {
		log.Fatal(err)
	}

	// 假设账户 A 和 B 已经存在，余额分别为 200 元和 100 元
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	var fromAccountID, toAccountID int = 1, 2
	var transferAmount int = 100

	var balance int
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromAccountID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if balance < transferAmount {
		fmt.Println("余额不足，无法转账")
		tx.Rollback()
		return
	}

	if _, err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", transferAmount, fromAccountID); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if _, err := tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", transferAmount, toAccountID); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if _, err := tx.Exec("INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (?, ?, ?)", fromAccountID, toAccountID, transferAmount); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("转账成功")
}