package main
import (
	"fmt"
	"goblog/bootstrap"
	"goblog/pkg/config"
	"goblog/pkg/model"
	"os"
)

func main() {
	fmt.Println("加载文件: golandTask4/blog/main.go 执行 main() 方法 ====== 启动应用")

	// 初始化配置
	config.Initialize()

	// 初始化数据库和 ORM
	bootstrap.SetupDB()

	// 连接数据库
	if err := model.ConnectDB(); err != nil {
		fmt.Println("连接数据库失败:", err)
		os.Exit(1)
	}

	fmt.Println("应用启动成功")
}