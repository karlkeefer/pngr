package posts

import (
	"encoding/json"
	"net/http"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/post"
	"github.com/karlkeefer/pngr/golang/models/user"
	"github.com/karlkeefer/pngr/golang/server/jwt"
)

func Handler(env *env.Env, w http.ResponseWriter, r *http.Request) (http.HandlerFunc, error) {
	switch r.Method {
	case "GET":
		return getPosts(env, w, r), nil
	case "POST":
		return createPost(env, w, r), nil
	default:
		return nil, errors.BadRequestMethod
	}
}

func getPosts(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return jwt.RequireAuth(r, func(u *user.User) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			posts, err := env.PostRepo().GetPostsForUser(u.ID)
			if err != nil {
				errors.Write(w, err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts)
		}
	})
}

func createPost(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return jwt.RequireAuth(r, func(u *user.User) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			decoder := json.NewDecoder(r.Body)
			p := &post.Post{}
			err := decoder.Decode(p)
			if err != nil || &p == nil {
				errors.Write(w, errors.NoJSONBody)
				return
			}

			// set author to current user
			p.AuthorID = u.ID

			post, err := env.PostRepo().CreatePost(p)
			if err != nil {
				errors.Write(w, err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
		}
	})
}
