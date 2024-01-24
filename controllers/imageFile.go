package controllers

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"crispypod.com/crispypod-backend/db"
	"crispypod.com/crispypod-backend/dbModels"
	"crispypod.com/crispypod-backend/helpers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

const ThumbnailFolder = "ImageFile"

var acceptedImageFormat []string = []string{
	"jpeg",
	"jpg",
	"png",
}

func ImageFileUpload(c *gin.Context) {
	userName := helpers.JWTFromContext(c.Request.Context())
	if len(userName) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please login"})
	}

	var uploadStruct UploadStruct
	c.Bind(&uploadStruct)

	var dbEpisode dbModels.Episode
	if err := db.DB.Model(dbModels.Episode{ID: uuid.Must(uuid.Parse(uploadStruct.EpisodeId))}).Find(&dbEpisode).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "episode not found"})
	}

	file, err := c.FormFile("file")
	if err != nil || file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error fetching file"})
	}

	fileNameSplited := strings.Split(file.Filename, ".")
	fileExt := strings.ToLower(fileNameSplited[len(fileNameSplited)-1])
	if !slices.Contains(acceptedImageFormat, fileExt) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported audio file format"})
	}

	savedFileName := uuid.New().String() + "." + fileExt
	fileFolder := filepath.Join(FolderPath, ThumbnailFolder)
	filePath := filepath.Join(fileFolder, savedFileName)

	if _, err := os.Stat(FolderPath); os.IsNotExist(err) {
		if err := os.Mkdir(FolderPath, os.ModePerm); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to create folder"})
		}
	}

	if _, err := os.Stat(fileFolder); os.IsNotExist(err) {
		if err := os.Mkdir(fileFolder, os.ModePerm); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to create folder"})
		}
	}

	c.SaveUploadedFile(file, filePath)

	dbEpisode.ThumbnailFileName = sql.NullString{String: savedFileName, Valid: true}
	dbEpisode.ThumbnailUploadName = sql.NullString{String: file.Filename, Valid: true}
	db.DB.Save(&dbEpisode)

	c.JSON(http.StatusOK, gin.H{
		"thumbnailFileName": savedFileName,
	})
}

func UploadFile(c *gin.Context) {
	userName := helpers.JWTFromContext(c.Request.Context())
	if len(userName) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please login"})
	}

	file, err := c.FormFile("file")
	if err != nil || file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error fetching file"})
	}

	fileNameSplited := strings.Split(file.Filename, ".")
	fileExt := strings.ToLower(fileNameSplited[len(fileNameSplited)-1])
	if !slices.Contains(acceptedImageFormat, fileExt) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported audio file format"})
	}

	savedFileName := uuid.New().String() + "." + fileExt
	fileFolder := filepath.Join(FolderPath, ThumbnailFolder)
	filePath := filepath.Join(fileFolder, savedFileName)

	if _, err := os.Stat(FolderPath); os.IsNotExist(err) {
		if err := os.Mkdir(FolderPath, os.ModePerm); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to create folder"})
		}
	}

	if _, err := os.Stat(fileFolder); os.IsNotExist(err) {
		if err := os.Mkdir(fileFolder, os.ModePerm); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to create folder"})
		}
	}

	c.SaveUploadedFile(file, filePath)

	c.JSON(http.StatusOK, gin.H{
		"fileName": savedFileName,
	})

}

func GetImageFile(c *gin.Context) {
	fileName := c.Param("fileName")
	thumbnailFilePath := filepath.Join(FolderPath, ThumbnailFolder, fileName)
	if _, err := os.Stat(thumbnailFilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Thumbnail not found"})
	}
	c.File(thumbnailFilePath)
}
