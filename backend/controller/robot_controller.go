package controller

import (
	"gin/config"
	"gin/model"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateRobot(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image is required"})
		return
	}

	lengthStr := c.PostForm("length")
	widthStr := c.PostForm("width")

	length, err1 := strconv.Atoi(lengthStr)
	width, err2 := strconv.Atoi(widthStr)

	if err1 != nil || err2 != nil || length < 0 || width < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Length or width invalid"})
		return
	}

	os.MkdirAll("./uploads", os.ModePerm)

	filename := time.Now().Format("20060102150405") + "_" + filepath.Base(file.Filename)
	savePath := filepath.Join("uploads", filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image"})
		return
	}

	robot := model.Robot{
		Length:   length,
		Width:    width,
		Filename: "/" + savePath,
	}

	config.DB.Create(&robot)

	c.JSON(http.StatusOK, gin.H{
		"length":     robot.Length,
		"width":      robot.Width,
		"filename":   robot.Filename,
		"created_at": robot.CreatedAt,
	})
}

func GetRobot(c *gin.Context) {
	var robots []model.Robot

	if err := config.DB.Find(&robots).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve robots"})
		return
	}

	c.JSON(http.StatusOK, robots)
}

func RemoveRobot(c *gin.Context) {
	id := c.Param("id")
	var robot model.Robot

	if err := config.DB.First(&robot, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Robot not found"})
		return
	}

	removePath := filepath.Join(".", robot.Filename)
	if err := os.Remove(removePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to remove image"})
		return
	}

	if err := config.DB.Delete(&robot).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete robot"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Robot deleted successfully"})
}

func EditRobot(c *gin.Context) {
	id := c.Param("id")
	var robot model.Robot

	if err := config.DB.First(&robot, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Robot not found"})
		return
	}

	file, err := c.FormFile("image")
	if err == nil {
		// 有上傳新圖就刪舊圖
		oldPath := filepath.Join(".", robot.Filename)
		if err := os.Remove(oldPath); err != nil && !os.IsNotExist(err) {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to remove old image"})
			return
		}

		os.MkdirAll("./uploads", os.ModePerm)

		filename := time.Now().Format("20060102150405") + "_" + filepath.Base(file.Filename)
		savePath := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save new image"})
			return
		}

		robot.Filename = "/" + savePath
	}

	lengthStr := c.PostForm("length")
	widthStr := c.PostForm("width")

	length, err1 := strconv.Atoi(lengthStr)
	width, err2 := strconv.Atoi(widthStr)

	if err1 != nil || err2 != nil || length < 0 || width < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Length or width invalid"})
		return
	}

	robot.Length = length
	robot.Width = width

	config.DB.Save(&robot)

	c.JSON(http.StatusOK, gin.H{
		"length":     robot.Length,
		"width":      robot.Width,
		"filename":   robot.Filename,
		"created_at": robot.CreatedAt,
	})
}
