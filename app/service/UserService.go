package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"learning/domain"
	"learning/repository"
	"net/http"
)

// 仵明雨：登录
func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 检查用户名和密码是否为空
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "需要用户名和密码"})
		return
	}
	// 查找用户
	var user domain.User
	if err := repository.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
		return
	}
	// 检查密码是否正确
	if user.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
		return
	}
	// 登录成功，创建session
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()
	// 登录成功，重定向到主页面
	c.Redirect(http.StatusFound, "/")
}

// 杨士涵：注册
func UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 检查用户名和密码是否为空
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "需要用户名和密码"})
		return
	}
	// 检查用户名是否已存在
	var existingUser domain.User
	if err := repository.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "用户名已被占用"})
		return
	}
	// 创建用户模型
	user := domain.User{
		Username: username,
		Password: password,
	}
	// 保存用户到数据库
	if err := repository.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "\n注册失败", "error": err.Error()})
		return
	}
	// 注册成功后创建session，存入redis数据库中
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	// 重定向到登录页面
	c.Redirect(http.StatusFound, "/login")
}

// 仵明雨：处理主页面请求
func HomePage(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	// 校验登录态
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "请登录"})
		return
	}
	// 重定向到主页面
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 仵明雨：检查用户是否已登录
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

// 更新用户信息
func UpdateUser(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "请登录"})
		return
	}

	var user domain.User
	if err := repository.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "用户不存在"})
		return
	}

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效数据"})
		return
	}

	if err := repository.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "无法修改用户信息", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户信息修改成功"})
}

// 获取当前用户信息
func GetCurrentUser(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "请登录"})
		return
	}

	var user domain.User
	if err := repository.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"sex":      user.Sex,
		"email":    user.Email,
		"phone":    user.Phone,
	})
}
