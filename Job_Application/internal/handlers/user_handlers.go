package handlers

import (
	"database/sql"
	"fmt"
	"job_portal/internal/services"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		user, err := services.GetUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserProfileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var userUpdate struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}

		// user/12 {userName: '', email: 'abc@gmail.com'}

		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetInt("userID")
		// isAdmin := c.GetBool("isAdmin")

		if userID != id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to update this user profile"})
			return
		}

		user, err := services.UpdateUserProfile(db, id, userUpdate.Username, userUpdate.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserProfilePcitureHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")

		if !isAdmin && userID != id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to update this user profile"})
			return
		}

		file, err := c.FormFile("profile_picture")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := os.MkdirAll(os.Getenv("UPLOAD_DIR"), os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating upload directory"})
			return
		}

		filename := fmt.Sprintf("%d-%s", id, filepath.Base(file.Filename))
		filePath := filepath.Join(os.Getenv("UPLOAD_DIR"), filename)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving uploaded file"})
			return
		}

		if err := services.UpdateUserProfilePicture(db, id, filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile picture updated successfully"})

	}
}
