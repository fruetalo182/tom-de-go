package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ServeImage(imagePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate the file extension
		ext := filepath.Ext(imagePath)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
			return
		}

		// Set the appropriate content type
		switch ext {
		case ".jpg", ".jpeg":
			c.Header("Content-Type", "image/jpeg")
		case ".png":
			c.Header("Content-Type", "image/png")
		case ".gif":
			c.Header("Content-Type", "image/gif")
		}

		// Serve the file
		c.File(imagePath)
	}
}