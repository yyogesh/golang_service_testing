package repository

import (
	"database/sql"
	"job_portal/internal/models"
)

func GetAllJobs(db *sql.DB) ([]models.Job, error) {
	rows, err := db.Query("SELECT * FROM jobs")

	// SELECT * FROM jobs where user_id = ?

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Company, &job.Location, &job.Salary, &job.Experience, &job.CreatedAt, &job.UserID); err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func CreateJob(db *sql.DB, job *models.Job) (*models.Job, error) {
	stmt, err := db.Prepare("INSERT INTO jobs (title, description, company, location, salary, experience, user_id) VALUES (?,?,?,?,?,?,?)")

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(job.Title, job.Description, job.Company, job.Location, job.Salary, job.Experience, job.UserID)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	job.ID = int(id)

	return job, err
}
