package posts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models"
	"github.com/karlkeefer/pngr/golang/server/write"
	"github.com/karlkeefer/pngr/golang/utils"
)

func Handler(env env.Env, w http.ResponseWriter, r *http.Request, u *models.User) (handler http.HandlerFunc) {
	// minimum authorization requirement
	if u.Status < models.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		if id, err := strconv.ParseInt(head, 10, 64); err == nil {
			return getPost(env, u, id)
		} else {
			return getPosts(env, u)
		}
	case http.MethodPost:
		return createPost(env, u, r)
	case http.MethodPut:
		return updatePost(env, u, r)
	case http.MethodDelete:
		if id, err := strconv.ParseInt(head, 10, 64); err == nil {
			return deletePost(env, u, id)
		}
		return write.Error(errors.RouteNotFound)
	default:
		return write.Error(errors.BadRequestMethod)
	}
}

func createPost(env env.Env, u *models.User, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	p := &models.Post{}
	err := decoder.Decode(p)
	if err != nil || &p == nil {
		return write.Error(errors.NoJSONBody)
	}

	// set author to current user
	p.AuthorID = u.ID

	return write.JSONorErr(env.PostRepo().Create(p))
}

func getPost(env env.Env, u *models.User, id int64) http.HandlerFunc {
	return write.JSONorErr(env.PostRepo().GetForUserByID(u.ID, id))
}

func getPosts(env env.Env, u *models.User) http.HandlerFunc {
	return write.JSONorErr(env.PostRepo().GetForUser(u.ID))
}

func updatePost(env env.Env, u *models.User, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	p := &models.Post{}
	err := decoder.Decode(p)
	if err != nil || &p == nil {
		return write.Error(errors.NoJSONBody)
	}

	// set author to current user
	p.AuthorID = u.ID

	return write.JSONorErr(env.PostRepo().Update(p))
}

func deletePost(env env.Env, u *models.User, id int64) http.HandlerFunc {
	return write.SuccessOrErr(env.PostRepo().DeleteForUser(u.ID, id))
}
