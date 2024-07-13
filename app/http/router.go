package http

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"learning/app/service"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 创建服务
	r := gin.Default()

	// 静态资源
	r.Static("/static", "./static")

	// 静态模板
	r.LoadHTMLGlob("templates/*")

	// 配置Redis连接和Session存储
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions("mysession", store))

	// 主页面
	protected := r.Group("/")
	// 仵明雨：检查是否登录
	protected.Use(service.AuthMiddleware())
	{
		protected.GET("/", service.HomePage)
		protected.GET("/about", func(c *gin.Context) {
			c.HTML(http.StatusOK, "about.html", gin.H{})
		})
		protected.GET("/my", func(c *gin.Context) {
			c.HTML(http.StatusOK, "my.html", gin.H{})
		})
		// 获取收藏的待办事项
		protected.GET("/favorites", func(c *gin.Context) {
			c.HTML(http.StatusOK, "favorite.html", gin.H{})
		})
		protected.POST("/update", service.UpdateUser)  // 新增更新用户信息路由
		protected.GET("/user", service.GetCurrentUser) // 新增获取当前用户信息的路由
	}

	// 仵明雨：登录页面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	// 杨士涵：注册页面
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})

	// 仵明雨：登录路由
	r.POST("/login", service.UserLogin)

	// 杨士涵：注册路由
	r.POST("/register", service.UserRegister)

	// 路由
	v1Group := r.Group("v1")
	{
		// 李铭伟:查看所有
		v1Group.GET("/todo", service.QueryAll)
		// 代宇航：添加
		v1Group.POST("/todo", service.AddTodo)
		// 孙文乐：修改
		v1Group.PUT("/todo/:id", service.UpdateTodo)
		// 马亿博:删除
		v1Group.DELETE("/todo/:id", service.DeleteTodo)
		// 孙文乐：收藏
		v1Group.PUT("/todo/favorite/:id", service.ToggleFavorite)
	}
	return r
}
