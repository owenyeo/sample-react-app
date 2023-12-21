package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/owenyeo/sample-react-app/backend/api"
	users "github.com/owenyeo/sample-react-app/backend/dataaccess"
	"github.com/owenyeo/sample-react-app/backend/database"
	"github.com/owenyeo/sample-react-app/backend/models"
	"github.com/pkg/errors"
)

const (
	ListUsers = "users.HandleList"

	SuccessfulListUsersMessage = "Successfully listed users"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveUsers           = "Failed to retrieve users in %s"
	ErrEncodeView              = "Failed to retrieve users in %s"
)

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListUsers))
	}

	users, err := users.List(db)
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	exists, err := db.UserExists(user.Name)
	if err != nil {
		http.Error(w, "Failed to check user existence", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	if exists {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User exists"))
		fmt.Println("User exists")
	} else {
		if err := db.AddUser(user); err != nil {
			http.Error(w, "Failed to add user to database", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User added to database"))
		fmt.Println("User added to database")
	}
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	if err := db.AddUser(user); err != nil {
		http.Error(w, "Failed to add user to database", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User added to database"))
}
