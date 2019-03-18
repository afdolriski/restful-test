package image

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/afdolriski/restful-test/database/models"
)

type Image = models.Image

func getJSON(target interface{}) error {
	url := "https://randomfox.ca/floof/"
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}

// Response struct
type Response struct {
	Image string
}

func saveImage(url string) string {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	imageDir := fmt.Sprintf("%s/images", currentDir)

	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		os.Mkdir(imageDir, 0755)
	}

	urlSplit := strings.Split(url, "/")

	fileName := urlSplit[len(urlSplit)-1]

	filePath := fmt.Sprintf("%s/%s", imageDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return fileName
}

// Fetch image
func Fetch(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	response := new(Response)
	getJSON(&response)

	fileName := saveImage(response.Image)
	image := Image{
		Name: fileName,
	}

	db.Create(&image)

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
