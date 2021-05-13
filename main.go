package main

import (
	"algorithm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

// SaveFileHandler : handles a POST form request for file
// https://ramezanpour.net/post/2020/09/12/file-upload-using-go-gin

func upload3d(c *gin.Context) {
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

	// from testAlgorithm.go
	algorithm.RunQuickHull3(newFileName, "done-" + newFileName)

	c.JSON(http.StatusOK, gin.H {
		"message": "file successfully uploaded",
		"filename": "done-" + newFileName,
	})
}

func upload2d(c *gin.Context) {
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

	// from testAlgorithm.go
	algorithm.RunQuickHull(newFileName, "done-" + newFileName)

	c.JSON(http.StatusOK, gin.H {
		"message": "file successfully uploaded",
		"filename": "done-" + newFileName,
	})
}

func main() {
	router := gin.Default()

	router.Static("/text", "./tmp")

	router.POST("/upload2d", upload2d)
	router.POST("/upload3d", upload3d)

	router.Run(":8080")
}