package controllers

import (
	"backend-iot-gps-tracker/database"
	"backend-iot-gps-tracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestCreate struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func GetChat(c *gin.Context) {
	db := database.GetDB()

	// get
	var chats []models.Chat
	if err := db.Find(&chats).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  "failed",
			"message": "Failed to get Chat",
			"data":    nil,
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"status":  "success",
		"message": "Success",
		"data":    chats,
	})
}

func CreateChat(c *gin.Context) {
	db := database.GetDB()

	var request requestCreate
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":             400,
			"status":           "failed",
			"message":          "Invalid request body",
			"original_message": err,
			"data":             nil,
		})
		return
	}

	// define model
	chat := models.Chat{
		Username: request.Username,
		Message:  request.Message,
	}

	// create
	if err := db.Create(&chat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  "failed",
			"message": "Failed to create Chat",
			"data":    nil,
		})
		return
	}

	// get
	var chats []models.Chat
	if err := db.Find(&chats).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  "failed",
			"message": "Failed to get Chat",
			"data":    nil,
		})
		return
	}

	// response
	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"status":  "success",
		"message": "Success",
		"data":    chats,
	})
}

func DeleteChat(c *gin.Context) {
	db := database.GetDB()

	// get
	var chats []models.Chat
	if err := db.Find(&chats).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  "failed",
			"message": "Failed to get Chat",
			"data":    nil,
		})
		return
	}
	
	for _, i := range chats {
		db.Delete(i)
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"status":  "success",
		"message": "Deleted Success",
		"data":    nil,
	})
}
