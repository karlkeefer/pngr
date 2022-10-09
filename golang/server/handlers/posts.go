package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/server/write"
)

func CreatePost(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	p := &db.Post{}
	err := decoder.Decode(p)
	if err != nil || p == nil {
		return write.Error(errors.NoJSONBody)
	}

	// set author to current user
	p.AuthorID = user.ID

	return write.JSONorErr(env.DB().CreatePost(r.Context(), db.CreatePostParams{
		AuthorID: p.AuthorID,
		Title:    p.Title,
		Body:     p.Body,
		Status:   db.PostStatusPublished,
	}))
}

func GetPost(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	id, err := getID(r)
	if err != nil {
		return write.Error(errors.RouteNotFound)
	}

	post, err := env.DB().FindPostByIDs(r.Context(), db.FindPostByIDsParams{
		AuthorID: user.ID,
		ID:       id,
	})
	if err != nil {
		if isNotFound(err) {
			return write.Error(errors.PostNotFound)
		}
		return write.Error(err)
	}

	return write.JSON(post)
}

func GetPosts(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	return write.JSONorErr(env.DB().FindPostsByAuthor(r.Context(), user.ID))
}

func UpdatePost(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	p := &db.Post{}
	err := decoder.Decode(p)
	if err != nil || p == nil {
		return write.Error(errors.NoJSONBody)
	}

	// check authority
	if p.AuthorID != user.ID {
		return write.Error(errors.RouteUnauthorized)
	}

	return write.JSONorErr(env.DB().UpdatePost(r.Context(), db.UpdatePostParams{
		ID:       p.ID,
		AuthorID: p.AuthorID,
		Title:    p.Title,
		Body:     p.Body,
	}))
}

func DeletePost(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	id, err := getID(r)
	if err != nil {
		return write.Error(errors.RouteNotFound)
	}

	return write.SuccessOrErr(env.DB().DeletePostByIDs(r.Context(), db.DeletePostByIDsParams{
		AuthorID: user.ID,
		ID:       id,
	}))
}
