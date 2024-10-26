package services

import (
	"database/sql"
	"job_portal/internal/models"
	"job_portal/internal/repository"
)

func GetAllJobs(db *sql.DB) ([]models.Job, error) {
	return repository.GetAllJobs(db)
}

func CreateJob(db *sql.DB, job *models.Job) (*models.Job, error) {
	return repository.CreateJob(db, job)
}
