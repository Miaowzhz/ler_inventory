package main

import (
	"learning/app/http"
	"learning/domain"
	"learning/repository"
)

func main() {
	// 连接数据库
	err := repository.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 模型绑定
	repository.DB.AutoMigrate(&domain.Todo{})
	repository.DB.AutoMigrate(&domain.User{})
	// 初始化路由
	r := http.InitRouter()
	// 运行
	r.Run()
}
