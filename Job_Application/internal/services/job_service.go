package services

import (
	"database/sql"
	"errors"
	"job_portal/internal/models"
	"job_portal/internal/repository"
)

func GetAllJobs(db *sql.DB) ([]models.Job, error) {
	return repository.GetAllJobs(db)
}

func CreateJob(db *sql.DB, job *models.Job) (*models.Job, error) {
	return repository.CreateJob(db, job)
}

func GetJobById(db *sql.DB, id int) (*models.Job, error) {
	return repository.GetJobById(db, id)
}

func UpdateJob(db *sql.DB, job *models.Job, userID int, isAdmin bool) (*models.Job, error) {
	existingJob, err := repository.GetJobById(db, job.ID)

	if err != nil {
		return nil, err
	}

	if !isAdmin && existingJob.UserID != userID {
		return nil, errors.New("unauthorized to update this job")
	}
	return repository.UpdateJob(db, job)
}

func DeleteJob(db *sql.DB, id int, userID int, isAdmin bool) error {
	existingJob, err := repository.GetJobById(db, id)

	if err != nil {
		return err
	}

	if !isAdmin && existingJob.UserID != userID {
		return errors.New("unauthorized to delete this job")
	}
	return repository.DeleteJob(db, id)
}
