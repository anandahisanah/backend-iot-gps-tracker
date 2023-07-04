package controllers

import (
	"encoding/json"
	"gps-tracker/database"
	"gps-tracker/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type requestGps struct {
	Link string `json:"link"`
}

func CreateGps(c *gin.Context) {
	db := database.GetDB()
	w := c.Writer

	// verify json
	var request requestGps
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":             400,
			"status":           "error",
			"message":          "Invalid JSON data",
			"original_message": err,
		})
		return
	}

	// define struct
	Gps := models.Gps{
		Datetime: time.Now(),
		Link:     request.Link,
	}

	// create
	if err := db.Create(&Gps).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":             400,
			"status":           "error",
			"message":          "Error saving GPS data",
			"original_message": err,
		})
		return
	}

	// response
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, _ := json.Marshal(gin.H{
		"code":    200,
		"status":  "success",
		"message": "Gps Saved",
		"data":    Gps,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
