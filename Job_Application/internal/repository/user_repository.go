package repository

import (
	"database/sql"
	"job_portal/internal/models"
)

func CreateUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", user.Username, user.Password, user.Email)
	return err
}

func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at, &user.IsAdmin, &user.ProfilePicture)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
