package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/owenyeo/sample-react-app/backend/api"
	"github.com/owenyeo/sample-react-app/backend/dataaccess"
	"github.com/owenyeo/sample-react-app/backend/database"
	"github.com/owenyeo/sample-react-app/backend/models"
	"github.com/pkg/errors"
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

func HandleAddPost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrAddPost, ListPosts))
	}

	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrAddPost, ListPosts))
	}

	err = dataaccess.AddPost(db, post)
	if err != nil {
		http.Error(w, "Failed to add post", http.StatusInternalServerError)
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrAddPost, ListPosts))
	}

	return &api.Response{
		Messages: []string{SuccessfulAddPostMessage},
	}, nil
}

func HandleListPosts(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListPosts))
	}

	posts, err := dataaccess.ListPosts(db)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts))
	}

	data, err := json.Marshal(posts)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListPosts))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil
}
