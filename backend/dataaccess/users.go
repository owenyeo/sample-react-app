package users

import (
	"github.com/owenyeo/sample-react-app/backend/database"
	"github.com/owenyeo/sample-react-app/backend/models"
)

func List(db *database.Database) ([]models.User, error) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// AddUser adds a new user to the database
func AddUser(db *database.Database, newUser models.User) error {
	
	_, err := db.Query("INSERT INTO users (name) VALUES ($1)", newUser.Name)

	if err != nil {
		return err
	}

	return nil
}