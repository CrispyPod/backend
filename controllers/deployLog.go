package controllers

import (
	"database/sql"
	"net/http"

	"crispypod.com/crispypod-backend/db"
	"crispypod.com/crispypod-backend/dbModels"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// THIS ROUTE SHOULD NOT BE PROXIED TO FRONTEND!!

type LogUploadStruct struct {
	LogID      string
	LogContent string
}

func DeployLogUpload(c *gin.Context) {
	// WE WILL NOT AUTH CHECK THIS, BE WARE!!
	var uploadStruct LogUploadStruct
	c.Bind(&uploadStruct)

	logID, err := uuid.Parse(uploadStruct.LogID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "log not found"})
		return
	}

	var dbDeployLog dbModels.DeployLog
	if err = db.DB.Model(dbModels.DeployLog{ID: logID}).Find(&dbDeployLog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "log not found"})
		return
	}

	dbDeployLog.Log = sql.NullString{String: uploadStruct.LogContent, Valid: true}

	db.DB.Save(dbDeployLog)

	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}
