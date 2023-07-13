package controllers

import (
	"backend-iot-gps-tracker/database"
	"backend-iot-gps-tracker/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type requestCreate struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func GetChat(c *gin.Context) {
	db := database.GetDB()

	paramLimit := c.Query("limit")
	var limit int

	// paramLimit into int
	if paramLimit != "" {
		limit64, err := strconv.ParseInt(paramLimit, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"status":  "failed",
				"message": "Invalid limit value",
				"data":    nil,
			})
			return
		}
		limit = int(limit64)
	}

	// get data with limit
	var chats []models.Chat
	if limit > 0 {
		db = db.Limit(limit)
	}
	if err := db.Order("id desc").Find(&chats).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"status":  "failed",
			"message": "Failed to get Chat",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"status":  "success",
		"message": "Chat retrieved successfully",
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

func CreateChatByUrl(c *gin.Context) {
	db := database.GetDB()

	// param
	paramLat := c.Param("lat")
	paramLon := c.Param("lon")

	// define model
	chat := models.Chat{
		Username: "System",
		Message:  "https://maps.google.com/?q=" + paramLat + "," + paramLon,
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
