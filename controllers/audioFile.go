package controllers

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"crispypod.com/crispypod-backend/db"
	"crispypod.com/crispypod-backend/helpers"
	"crispypod.com/crispypod-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type UploadStruct struct {
	EpisodeId string `form:"episodeId"`
}

const FolderPath = "UploadFile"
const AudioFileFolder = "AudioFile"

var acceptedAudioFormat []string = []string{
	"mp3",
	"wav",
	"aac",
}

func AudioFileUpload(c *gin.Context) {
	userName := helpers.JWTFromContext(c.Request.Context())
	if len(userName) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please login"})
	}

	var uploadStruct UploadStruct
	c.Bind(&uploadStruct)

	var dbEpisode models.Episode
	var epId uuid.UUID
	var err error
	if epId, err = uuid.Parse(uploadStruct.EpisodeId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "episode not found"})
		return
	}

	if err := db.DB.Model(models.Episode{ID: epId}).Find(&dbEpisode).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "episode not found"})
		return
	}

	file, err := c.FormFile("file")

	if err != nil || file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error fetching file"})
	}

	fileNameSplited := strings.Split(file.Filename, ".")
	fileExt := strings.ToLower(fileNameSplited[len(fileNameSplited)-1])
	if !slices.Contains(acceptedAudioFormat, fileExt) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported audio file format"})
	}

	savedFileName := uuid.New().String() + "." + fileExt
	fileFolder := filepath.Join(FolderPath, AudioFileFolder)
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

	dbEpisode.AudioFileName = sql.NullString{String: savedFileName, Valid: true}
	dbEpisode.AudioFileUploadName = sql.NullString{String: file.Filename, Valid: true}

	db.DB.Save(&dbEpisode)

	c.JSON(http.StatusOK, gin.H{
		"audioFileName": savedFileName,
	})

}

func GetAudioFile(c *gin.Context) {
	fileName := c.Param("fileName")
	audioFilePath := filepath.Join(FolderPath, AudioFileFolder, fileName)
	if _, err := os.Stat(audioFilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Audio file not found"})
	}
	c.File(audioFilePath)
}
