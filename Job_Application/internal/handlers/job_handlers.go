package handlers

import (
	"database/sql"
	"job_portal/internal/models"
	"job_portal/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllJobsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		jobs, err := services.GetAllJobs(db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, jobs)
	}
}

func CreateJobHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var job models.Job
		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetInt("userID")
		job.UserID = userID

		cratedJob, err := services.CreateJob(db, &job)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, cratedJob)
	}
}
