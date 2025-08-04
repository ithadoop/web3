/* 题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。 */

package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

// User 用户模型
type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;unique;not null"`
	Password string `gorm:"size:255;not null"`
	Posts    []Post // 一对多关系：一个用户有多篇文章
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title    string    `gorm:"size:255;not null"`
	Content  string    `gorm:"type:text;not null"`
	UserID   uint      // 外键，关联User模型
	User     User      `gorm:"foreignKey:UserID"` // 定义外键关系
	Comments []Comment // 一对多关系：一篇文章有多个评论
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	PostID  uint   // 外键，关联Post模型
	Post    Post   `gorm:"foreignKey:PostID"` // 定义外键关系
	UserID  uint   // 外键，关联User模型
	User    User   `gorm:"foreignKey:UserID"` // 定义外键关系
}

func main() {
	//数据库连接和自动迁移代码
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	    panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
	