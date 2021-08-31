package server

import (
	"encoding/json"
	"net/http"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models"
	"github.com/karlkeefer/pngr/golang/server/write"
)

func CreatePost(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status < models.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	p := &models.Post{}
	err := decoder.Decode(p)
	if err != nil || &p == nil {
		return write.Error(errors.NoJSONBody)
	}

	// set author to current user
	p.AuthorID = user.ID

	return write.JSONorErr(env.PostRepo().Create(p))
}

func GetPost(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status < models.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	id, err := getID(r)
	if err != nil {
		return write.Error(errors.RouteNotFound)
	}
	return write.JSONorErr(env.PostRepo().GetForUserByID(user.ID, id))
}

func GetPosts(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status < models.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	return write.JSONorErr(env.PostRepo().GetForUser(user.ID))
}

func UpdatePost(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status < models.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	p := &models.Post{}
	err := decoder.Decode(p)
	if err != nil || &p == nil {
		return write.Error(errors.NoJSONBody)
	}

	// check authority
	if p.AuthorID != user.ID {
		return write.Error(errors.RouteUnauthorized)
	}

	return write.JSONorErr(env.PostRepo().Update(p))
}

func DeletePost(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status < models.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	id, err := getInt64(r)
	if err != nil {
		return write.Error(errors.RouteNotFound)
	}

	return write.SuccessOrErr(env.PostRepo().DeleteForUser(user.ID, id))
}
