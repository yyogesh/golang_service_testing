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
	var profilePicture sql.NullString // Use sql.NullString to handle NULL values
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at, &user.IsAdmin, &profilePicture)
	if err != nil {
		return nil, err
	}

	// If profilePicture is valid (not NULL), assign it to the user struct
	if profilePicture.Valid {
		user.ProfilePicture = &profilePicture.String
	} else {
		user.ProfilePicture = nil // Explicitly set to nil if the database value is NULL
	}
	return &user, nil
}

func GetUserByUserName(db *sql.DB, username string) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at, &user.IsAdmin, &user.ProfilePicture)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserProfile(db *sql.DB, user *models.User) (*models.User, error) {
	_, err := db.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", user.Username, user.Email, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserProfilePicture(db *sql.DB, id int, profilePicture string) error {
	_, err := db.Exec("UPDATE users SET profile_picture = ? WHERE id = ?", profilePicture, id)
	if err != nil {
		return err
	}
	return nil
}
