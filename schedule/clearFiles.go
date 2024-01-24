package schedule

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"crispypod.com/crispypod-backend/db"
	"crispypod.com/crispypod-backend/dbModels"
	"golang.org/x/exp/slices"
)

func ClearFiles() {
	ClearAudioFile()
	ClearImageFile()
}

func ClearAudioFile() {
	fmt.Println("Clearing audio files started")
	var audioFileNames []string
	// var episodes []dbModels.Episode
	if err := db.DB.Model(&dbModels.Episode{}).Select("audio_file_name").Find(&audioFileNames).Error; err != nil {
		fmt.Println("Failed to fetch episode data")
		return
	}

	// fmt.Println(audioFileNames)
	folderPath := filepath.Join("UploadFile", "AudioFile")
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Failed to read audio file directory")
	}
	for _, e := range entries {
		if !e.IsDir() && !slices.Contains(audioFileNames, e.Name()) {
			var fileInfo fs.FileInfo
			fileInfo, err = e.Info()
			if err != nil {
				fmt.Println("Failed to get fileinfo for file ", e.Name())
			}
			if time.Since(fileInfo.ModTime()).Hours() > 24 {
				os.Remove(filepath.Join(folderPath, e.Name()))
				fmt.Println("Deleted ", e.Name())
			}
		}
	}

	fmt.Println("Clearing audio files finished.")

}

func ClearImageFile() {
	fmt.Println("Clearing image files started")
	var imageFileNames []string
	// var episodes []dbModels.Episode
	if err := db.DB.Model(&dbModels.Episode{}).Select("thumbnail_file_name").Find(&imageFileNames).Error; err != nil {
		fmt.Println("Failed to fetch episode data")
		return
	}

	var siteConfig dbModels.SiteConfig
	if err := db.DB.First(&siteConfig).Error; err != nil {
		fmt.Println("Failed to fetch site config")
		return
	}

	// fmt.Println(audioFileNames)
	folderPath := filepath.Join("UploadFile", "ImageFile")
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Failed to read audio file directory")
	}
	for _, e := range entries {
		if !e.IsDir() && !slices.Contains(imageFileNames, e.Name()) {
			var fileInfo fs.FileInfo
			fileInfo, err = e.Info()
			if err != nil {
				fmt.Println("Failed to get fileinfo for file ", e.Name())
			}
			if time.Since(fileInfo.ModTime()).Hours() > 24 {
				os.Remove(filepath.Join(folderPath, e.Name()))
				fmt.Println("Deleted ", e.Name())
			}
		}
	}

	fmt.Println("Clearing image files finished.")

}
