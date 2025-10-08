package controllers

import (
	"backend/configs"
	"backend/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTodoRequest รับข้อมูลจาก frontend
type CreateTodoRequest struct {
	Title			string	`json:"title" binding:"required"`
	Description string	`json:"description"`
}

// C
func CreateTodo(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthirzed"})
		return
	}

	userID := userId.(uint)

	var req CreateTodoRequest
	// อ่านข้อมูล JSON จาก Request Body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := entity.Todo{
		Title: req.Title,
		Description: req.Description,
		Status:	"pending",
		UserID: userID,
	}

	if err := configs.DB().Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"massage": "todo created", "todo": todo})
}

// R
func GetTodos(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := userId.(uint)
	
	var todos []entity.Todo
	if err := configs.DB().Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

// U
func UpdateTodo(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthrized"})
		return
	}

	userID := userId.(uint)

	id, _ := strconv.Atoi(c.Param("id"))
	var todo entity.Todo
	if err := configs.DB().First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	// ตรวจสิทธิ์
	if todo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := configs.DB().Model(&todo).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo updated", "todo": todo})
}

// D
func DeleteTodo(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := userId.(uint)

	id, _ := strconv.Atoi(c.Param("id"))
	var todo entity.Todo
	if err := configs.DB().First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	// ตรวจเจ้าของ
	if todo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	if err := configs.DB().Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
}