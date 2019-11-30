package posts

import (
	"encoding/json"
	"net/http"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/post"
	"github.com/karlkeefer/pngr/golang/models/user"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
)

func Handler(env env.Env, w http.ResponseWriter, r *http.Request) (handler http.HandlerFunc) {
	// protect all of these routes with auth wrapper
	return jwt.RequireAuth(user.StatusActive, env, w, r, func(u *user.User) http.HandlerFunc {
		switch r.Method {
		case http.MethodGet:
			return getPosts(u, env)
		case http.MethodPost:
			return createPost(u, env, r)
		default:
			return write.Error(errors.BadRequestMethod)
		}
	})
}

func getPosts(u *user.User, env env.Env) http.HandlerFunc {
	return write.JSONorErr(env.PostRepo().GetPostsForUser(u.ID))
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

	return write.JSONorErr(env.PostRepo().CreatePost(p))
}
