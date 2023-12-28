package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/owenyeo/sample-react-app/backend/api"
	"github.com/owenyeo/sample-react-app/backend/auth"
	"github.com/owenyeo/sample-react-app/backend/dataaccess"
	"github.com/owenyeo/sample-react-app/backend/database"
	"github.com/owenyeo/sample-react-app/backend/models"
	"github.com/pkg/errors"
)

const (
	ListUsers = "users.HandleList"

	SuccessfulListUsersMessage = "Successfully listed users"
	SuccessfulUserExistsMessage= "User Exists"
	SuccessfulAddUserMessage   = "Successfully added user"
	ErrAddUser				   = "Failed to add user in %s"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveUsers           = "Failed to retrieve users in %s"
	ErrEncodeView              = "Failed to retrieve users in %s"
)

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListUsers))
	}

	users, err := dataaccess.ListUsers(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	data, err := json.Marshal(users)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListUsers))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	exists, err := dataaccess.UserExists(db, user.Name)
	if err != nil {
		http.Error(w, "Failed to check user existence", http.StatusInternalServerError)
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	if exists {
		token, err := auth.GenerateToken(user.Name)

		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			fmt.Println(err)
			return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
		}
		return &api.Response{
			Messages: []string{token},
		}, nil

	} else {
		if err := dataaccess.AddUser(db, user); err != nil {
			http.Error(w, "Failed to add user to database", http.StatusInternalServerError)
			fmt.Println(err)
			return nil, errors.Wrap(err, fmt.Sprintf(ErrAddUser, ListUsers))
		}

		token, err := auth.GenerateToken(user.Name)

		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			fmt.Println(err)
			return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
		}
		return &api.Response{
			Messages: []string{token},
		}, nil
	}
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	if err := dataaccess.AddUser(db, user); err != nil {
		http.Error(w, "Failed to add user to database", http.StatusInternalServerError)
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrAddUser, ListUsers))
	}

	token, err := auth.GenerateToken(user.Name)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrAddUser, ListUsers))
	}
	

	return &api.Response{
		Messages: []string{token},
	}, nil
}
