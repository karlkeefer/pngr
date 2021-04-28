package posts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/post"
	"github.com/karlkeefer/pngr/golang/models/user"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
	"github.com/karlkeefer/pngr/golang/utils"
)

func Handler(env env.Env, w http.ResponseWriter, r *http.Request) (handler http.HandlerFunc) {
	// protect all of these routes with auth wrapper
	return jwt.RequireAuth(user.StatusActive, env, w, r, func(u *user.User) http.HandlerFunc {
		var head string
		head, r.URL.Path = utils.ShiftPath(r.URL.Path)

		switch r.Method {
		case http.MethodGet:
			if id, err := strconv.ParseInt(head, 10, 64); err == nil {
				return getPost(u, id, env, w)
			} else {
				return getPosts(u, env)
			}
		case http.MethodPost:
			return createPost(u, env, r)
		case http.MethodPut:
			return updatePost(u, env, r)
		case http.MethodDelete:
			if id, err := strconv.ParseInt(head, 10, 64); err == nil {
				return deletePost(u, id, env)
			}
			return write.Error(errors.RouteNotFound)
		default:
			return write.Error(errors.BadRequestMethod)
		}
	})
}

func createPost(u *user.User, env env.Env, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	p := &post.Post{}
	err := decoder.Decode(p)
	if err != nil || &p == nil {
		return write.Error(errors.NoJSONBody)
	}

	// set author to current user
	p.AuthorID = u.ID

	return write.JSONorErr(env.PostRepo().Create(p))
}

func getPost(u *user.User, id int64, env env.Env, w http.ResponseWriter) http.HandlerFunc {
	return write.JSONorErr(env.PostRepo().GetForUserByID(u.ID, id))
}

func getPosts(u *user.User, env env.Env) http.HandlerFunc {
	return write.JSONorErr(env.PostRepo().GetForUser(u.ID))
}

func updatePost(u *user.User, env env.Env, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	p := &post.Post{}
	err := decoder.Decode(p)
	if err != nil || &p == nil {
		return write.Error(errors.NoJSONBody)
	}

	// set author to current user
	p.AuthorID = u.ID

	return write.JSONorErr(env.PostRepo().Update(p))
}

func deletePost(u *user.User, id int64, env env.Env) http.HandlerFunc {
	return write.SuccessOrErr(env.PostRepo().DeleteForUser(u.ID, id))
}
