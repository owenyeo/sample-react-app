package users

import (
	"encoding/json"
	"net/http"

	"github.com/owenyeo/sample-react-app/backend/api"
	"github.com/owenyeo/sample-react-app/backend/auth"
	"github.com/owenyeo/sample-react-app/backend/dataaccess"
	"github.com/owenyeo/sample-react-app/backend/database"
	"github.com/owenyeo/sample-react-app/backend/models"
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

func HandleList(w http.ResponseWriter, r *http.Request) {
	db, err := database.GetDB()

	if err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
		return
	}

	users, err := dataaccess.ListUsers(db)
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to retrieve users"), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to retrieve users"), http.StatusInternalServerError)
		return
	}

	api.WriteJSON(w, api.NewResponse(data, "Successfully listed users"), http.StatusOK)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Invalid request payload"), http.StatusBadRequest)
		return
	}

	db, err := database.GetDB()
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to load database"), http.StatusInternalServerError)
		return
	}

	exists, err := dataaccess.UserExists(db, user.Name)
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to retrieve users"), http.StatusInternalServerError)
		return
	}

	if exists {
		token, err := auth.GenerateToken(user.Name)

		if err != nil {
			api.WriteJSON(w, api.NewResponse(err, "Failed to generate token"), http.StatusInternalServerError)
			return
		}
		api.WriteJSON(w, api.NewResponse([]byte(token), "Successfully logged in"), http.StatusOK)

	} else {
		if err := dataaccess.AddUser(db, user); err != nil {
			api.WriteJSON(w, api.NewResponse(nil, "Failed to add user"), http.StatusInternalServerError)
			return
		}

		token, err := auth.GenerateToken(user.Name)

		if err != nil {
			api.WriteJSON(w, api.NewResponse(nil, "Failed to generate token"), http.StatusInternalServerError)
			return
		}
		api.WriteJSON(w, api.NewResponse([]byte(token), "Successfully added user"), http.StatusOK)
		return
	}
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.WriteJSON(w, api.NewResponse(nil, "Invalid request payload"), http.StatusBadRequest)
		return
	}

	db, err := database.GetDB()
	if err != nil {
		api.WriteJSON(w, api.NewResponse(nil, "Failed to load database"), http.StatusInternalServerError)
		return
	}

	if err := dataaccess.AddUser(db, user); err != nil {
		api.WriteJSON(w, api.NewResponse(nil, "Failed to add user"), http.StatusInternalServerError)
		return
	}

	token, err := auth.GenerateToken(user.Name)
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to generate token"), http.StatusInternalServerError)
		return
	}
	

	api.WriteJSON(w, api.NewResponse([]byte(token), "Successfully added user"), http.StatusOK)
}
