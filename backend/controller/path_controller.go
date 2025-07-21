package controller

import (
	"gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalculatePath(c *gin.Context) {
	var req model.PathRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if req.Start.X < 0 || req.Start.Y < 0 || req.End.X < 0 || req.End.Y < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "coordinate invalid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pathX": (req.End.X - req.Start.X), "pathY": (req.End.Y - req.Start.Y)})
}
