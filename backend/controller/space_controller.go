package controller

import (
	"gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigSpace(c *gin.Context) {
	var req model.Space
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if req.Length < 0 || req.Width < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Space size invalid"})
	}

	c.JSON(http.StatusOK, gin.H{"Space length": (req.Length), "Space width": (req.Width)})
}
