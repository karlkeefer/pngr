package post

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Post struct {
	ID       int64     `json:"id"`
	AuthorID int64     `json:"author_id" db:"author_id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Status   int       `json:"status"`
	Created  time.Time `json:"created"`
}

const (
	PostStatusPrivate = 0
	PostStatusPublic  = 1
)

// REPO stuff
// TODO: consider moving repo to separate package
type Repo struct {
	create     *sqlx.NamedStmt
	getForUser *sqlx.Stmt
}

// NewRepo prepares statements, and panics if a statement fails to create
func NewRepo(db *sqlx.DB) *Repo {
	create, err := db.PrepareNamed(`INSERT INTO posts (author_id, title, body, status) VALUES (:author_id, :title, :body, :status) RETURNING *`)
	if err != nil {
		panic(err)
	}
	getForUser, err := db.Preparex(`SELECT * FROM posts WHERE author_id = $1 ORDER BY id DESC`)
	if err != nil {
		panic(err)
	}
	return &Repo{
		create,
		getForUser,
	}
}

func (r *Repo) GetPostsForUser(userID int64) (posts []*Post, err error) {
	posts = []*Post{} // always at least return empty list
	err = r.getForUser.Select(&posts, userID)
	return
}

func (r *Repo) CreatePost(p *Post) (post *Post, err error) {
	post = &Post{}
	err = r.create.Get(post, p)
	return
}
