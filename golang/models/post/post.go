package post

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/karlkeefer/pngr/golang/errors"
)

type Post struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Status int

const (
	StatusPrivate Status = 0
	StatusPublic         = 1
)

// REPO stuff
// TODO: consider moving repo to separate package
type Repo interface {
	Create(p *Post) (post *Post, err error)
	Update(p *Post) (post *Post, err error)
	GetForUser(author_id int64) (posts []*Post, err error)
	GetForUserByID(author_id int64, id int64) (post *Post, err error)
	DeleteForUser(author_id int64, id int64) error
}

type repo struct {
	create         *sqlx.NamedStmt
	update         *sqlx.NamedStmt
	getForUser     *sqlx.Stmt
	getForUserByID *sqlx.Stmt
	deleteForUser  *sqlx.Stmt
}

// NewRepo prepares statements, and panics if a statement fails to create
func NewRepo(db *sqlx.DB) Repo {
	create, err := db.PrepareNamed(`INSERT INTO posts (author_id, title, body, status) VALUES (:author_id, :title, :body, :status) RETURNING *`)
	if err != nil {
		panic(err)
	}
	update, err := db.PrepareNamed(`UPDATE posts SET title = :title, body = :body, updated_at = :updated_at WHERE id = :id AND author_id = :author_id RETURNING *`)
	if err != nil {
		panic(err)
	}
	getForUser, err := db.Preparex(`SELECT * FROM posts WHERE author_id = $1 ORDER BY id DESC`)
	if err != nil {
		panic(err)
	}
	getForUserByID, err := db.Preparex(`SELECT * FROM posts WHERE author_id = $1 AND id = $2 LIMIT 1`)
	if err != nil {
		panic(err)
	}
	deleteForUser, err := db.Preparex(`DELETE FROM posts WHERE author_id = $1 AND id = $2`)
	if err != nil {
		panic(err)
	}

	return &repo{
		create,
		update,
		getForUser,
		getForUserByID,
		deleteForUser,
	}
}

func (r *repo) GetForUser(author_id int64) (posts []*Post, err error) {
	posts = []*Post{} // always at least return empty list
	err = r.getForUser.Select(&posts, author_id)
	return
}

func (r *repo) GetForUserByID(author_id int64, id int64) (post *Post, err error) {
	post = &Post{}
	err = r.getForUserByID.Get(post, author_id, id)
	if err == sql.ErrNoRows {
		err = errors.PostNotFound
	}
	return
}

func (r *repo) Create(p *Post) (post *Post, err error) {
	post = &Post{}
	err = r.create.Get(post, p)
	return
}

func (r *repo) Update(in *Post) (out *Post, err error) {
	out = &Post{}
	in.UpdatedAt = time.Now()
	err = r.update.Get(out, in)
	return
}

func (r *repo) DeleteForUser(author_id int64, id int64) (err error) {
	_, err = r.deleteForUser.Exec(author_id, id)
	return
}
