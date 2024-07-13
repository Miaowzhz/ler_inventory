package service

import (
	"github.com/gin-gonic/gin"
	"learning/domain"
	"learning/repository"
	"net/http"
)

// 李铭伟:查看所有
func QueryAll(c *gin.Context) {
	var todoList []domain.Todo
	err := repository.DB.Find(&todoList).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// 代宇航：添加待办
func AddTodo(c *gin.Context) {
	// 从请求中取出数据
	var todo domain.Todo
	c.Bind(&todo)
	// 存入数据库
	// 返回响应
	err := repository.DB.Create(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": todo})
	}
}

// 修改待办
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo domain.Todo
	err := repository.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	c.Bind(&todo)
	err = repository.DB.Save(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 马亿博:删除待办
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo domain.Todo
	err := repository.DB.Where("id = ?", id).Delete(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 李铭伟:id查询
func QueryById(c *gin.Context) {
	id := c.Param("id")
	var todo domain.Todo
	err := repository.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "id 不存在"})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
