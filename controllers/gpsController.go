package controllers

import (
	"encoding/json"
	"gps-tracker/database"
	"gps-tracker/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestGps struct {
	Link string `json:"link"`
}

func CreateGps(c *gin.Context) {
	db := database.GetDB()
	w := c.Writer

	var request RequestGps

	// verify json
	err := c.ShouldBindJSON(&request)

	// error invalid json
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid JSON data",
		})
		return
	}

	// define struct
	Gps := models.Gps{
		Datetime: time.Now(),
		Link:     request.Link,
	}

	// create
	errSave := db.Create(&Gps)

	// error saving
	if errSave != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error saving GPS data",
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
