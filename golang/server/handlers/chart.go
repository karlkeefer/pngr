package handlers

import (
	"net/http"

	mdb "github.com/QuinnMain/infograph/golang/db"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/QuinnMain/infograph/golang/env"
	"github.com/QuinnMain/infograph/golang/errors"
	"github.com/QuinnMain/infograph/golang/server/write"
)

// d = ""
// r := &github.Repository{Description:&d}
// client.Repositories.Edit("user", "repo", r)

func GetChart(env env.Env, user *mdb.MUser, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	// чекаем пользователя на жизнь
	if *user.Status != mdb.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	id, err := getInt64("id", r)
	if err != nil {
		return write.Error(errors.RouteNotFound)
	}
	filter := bson.D{
		{"_id", id},
	}
	coll := env.Collection("commodityCharts")
	var result mdb.MCommodity

	err = coll.FindOne(r.Context(), filter).Decode(&result)
	if err != nil {
		if isNotFound(err) {
			return write.Error(errors.PostNotFound)
		}
		return write.Error(err)
	}

	return write.JSON(result)

}

// func GetCharts(env env.Env, user *mdb.MUser, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
// 	if *user.Status != mdb.UserStatusActive {
// 		return write.Error(errors.RouteUnauthorized)
// 	}
// }

// func UpdatePost(env env.Env, user *mdb.MUser, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
// 	if *user.Status != mdb.UserStatusActive {
// 		return write.Error(errors.RouteUnauthorized)
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	c := &mdb.MCommodity{}
// 	err := decoder.Decode(c)
// 	if err != nil || &c == nil {
// 		return write.Error(errors.NoJSONBody)
// 	}

// 	// check authority
// 	if p.AuthorID != user.ID {
// 		return write.Error(errors.RouteUnauthorized)
// 	}

// 	return write.JSONorErr(env.DB().UpdatePost(r.Context(), db.UpdatePostParams{
// 		ID:       p.ID,
// 		AuthorID: p.AuthorID,
// 		Title:    p.Title,
// 		Body:     p.Body,
// 	}))
// }
