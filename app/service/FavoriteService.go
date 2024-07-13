package service

import (
	"github.com/gin-gonic/gin"
	"learning/domain"
	"learning/repository"
	"net/http"
)

// 孙文乐：查看所有收藏
func GetFavorites(c *gin.Context) {
	var favorites []domain.Todo
	err := repository.DB.Where("favorite = ?", true).Find(&favorites).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, favorites)
	}
}

// 收藏
func ToggleFavorite(c *gin.Context) {
	id := c.Param("id")
	var todo domain.Todo

	// 根据ID查找待办事项
	err := repository.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "id 不存在"})
		return
	}

	// 切换收藏状态
	todo.Favorite = !todo.Favorite

	// 保存更新后的待办事项
	err = repository.DB.Save(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "favorite": todo.Favorite})
	}
}
