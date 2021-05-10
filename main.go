package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

// SaveFileHandler : handles a POST form request for file
// https://ramezanpour.net/post/2020/09/12/file-upload-using-go-gin
func SaveFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "no file is received",
		})
		return
	}

	extension:= filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "./tmp/" + newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"message": "unable to save file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"message": "file successfully uploaded",
		"filename": newFileName,
	})
}

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})

	router.POST("/upload", SaveFileHandler)

	router.Run(":8080")
}