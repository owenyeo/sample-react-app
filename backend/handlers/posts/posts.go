package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/owenyeo/sample-react-app/backend/api"
	"github.com/owenyeo/sample-react-app/backend/dataaccess"
	"github.com/owenyeo/sample-react-app/backend/database"
	"github.com/owenyeo/sample-react-app/backend/models"
)

const (
	ListPosts = "posts.HandlePostList"

	SuccessfulListPostsMessage = "Successfully listed posts"
	SuccessfulAddPostMessage   = "Successfully added post"
	ErrAddPost                 = "Failed to add post in %s"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrievePosts           = "Failed to retrieve posts in %s"
	ErrEncodeView              = "Failed to retrieve posts in %s"
)

func HandleListPosts(w http.ResponseWriter, r *http.Request) {
	db, err := database.GetDB()
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to load database"), http.StatusInternalServerError)
		return
	}

	posts, err := dataaccess.ListPosts(db)
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to retrieve posts"), http.StatusInternalServerError)
		return
	}

	fmt.Println(posts)
	api.WriteJSON(w, api.NewResponse(posts, "Successfully listed posts"), http.StatusOK)
}

func HandleAddPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Invalid request payload"), http.StatusBadRequest)
		return
	}

	db, err := database.GetDB()
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to load database"), http.StatusInternalServerError)
		return
	}

	if err := dataaccess.AddPost(db, post); err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to add post"), http.StatusInternalServerError)
		return
	}

	latestPost, err := dataaccess.GetLatestPost(db)
	if err != nil {
		api.WriteJSON(w, api.NewResponse(err, "Failed to retrieve latest post"), http.StatusInternalServerError)
		return
	}

	api.WriteJSON(w, api.NewResponse(latestPost, "Successfully added post"), http.StatusOK)
}
