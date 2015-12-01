package main

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

const (
	QNewUser  string = "INSERT INTO users (email, username, password, created_at, updated_at) VALUES (:email, :username, :password, :created_at, :updated_at) RETURNING id"
	QNewList  string = "INSERT INTO lists (title, user_id, created_at, updated_at) VALUES (:title, :user_id, :created_at, :updated_at)"
	QNewMovie string = "INSERT INTO movies (imdb_id, user_rating, list_id) VALUES (:IMDB, :UserRating, :ID)"
)

type Timestamps struct {
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type User struct {
	ID       int64 `json:"-"`
	Username string
	Password string `json:"-"`
	Email    string
	Lists    []*List
	Timestamps
}
type List struct {
	ID      int64    `json:"id"`
	Title   string   `json:"title"`
	Movies  []*Movie `json:"movies"`
	User_ID int64    `json:"-"`
	Timestamps
}

type Movie struct {
	ID         int64 `json:"id"`
	IMDB       string
	UserRating float32
	List_ID    int64 `json:"-"`
	Watched    bool
	Timestamps
}

func (u *User) DeleteList(db *sqlx.DB, id int64) error {
	stmt := "DELETE FROM lists WHERE id=$1 AND user_id=$2"
	_, err := db.Exec(stmt, id, u.ID)
	return err
}

func (u *User) GenSignedString(key []byte) string {
	token := jwt.New(jwt.SigningMethodHS256)
	duration := time.Now().Add(time.Hour * 96).Unix()

	// Set a header and a claim
	token.Header["typ"] = "JWT"
	token.Claims["exp"] = duration
	token.Claims["auth"] = true
	token.Claims["uid"] = u.ID
	token.Claims["username"] = u.Username

	// Generate encoded token
	t, err := token.SignedString(key)

	if err != nil {
		log.Printf("Error generating signed string out of token: %s", err)

		// return invalid token [temporarily]
		return ""
	}
	return t
}

func (u *User) ListsAll() []*List {
	return u.Lists
}

func (u *User) NewList(db *sqlx.DB, l List) error {
	l.User_ID = u.ID
	l.Created_At = time.Now()
	l.Updated_At = time.Now()
	_, err := db.NamedExec(QNewList, l)
	// query, args, err := ParseQueryStruct(QNewList, l, db)
	if err != nil {
		return err
	}
	return nil
}

func (l *List) MoviesAll() []*Movie {
	return l.Movies
}

func (l *List) AddMovie(db *sqlx.DB, m Movie) error {
	m.List_ID = l.ID
	_, err := db.NamedExec(QNewMovie, m)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(db *sqlx.DB, u User) (int64, error) {
	u.Created_At = time.Now()
	u.Updated_At = time.Now()
	query, args, err := ParseQueryStruct(QNewUser, u, db)

	if err != nil {
		return 0, err
	}

	// Postgres does not return last inserted row id in sql.Result
	// so we have to do this
	var id int64
	row := db.QueryRowx(query, args...)
	row.Scan(&id)

	if row.Err() != nil {
		return 0, row.Err()
	}

	return id, nil
}

// func NewMovie (db *sqlx.DB, m
func FindUserByUsername(db *sqlx.DB, u string) (*User, error) {
	user := User{}
	// query := "SELECT * FROM users WHERE username=$1"
	err := db.Get(&user, "SELECT * FROM users WHERE username=$1", u)
	err = user.QueryResources(db)
	if err != nil {
		log.Println(err)
	}
	return &user, err
}
func FindUserByID(db *sqlx.DB, uid int64) (*User, error) {
	user := User{}
	// query := "SELECT * FROM users WHERE username=$1"
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", uid)
	err = user.QueryResources(db)
	if err != nil {
		log.Println(err)
	}
	return &user, err
}
func (u *User) QueryResources(db *sqlx.DB) error {
	lists := []*List{}
	err := db.Select(&lists, "SELECT * FROM lists WHERE user_id=$1", u.ID)
	if err != nil {
		return err
	}
	for _, l := range lists {
		l.QueryResources(db)
	}
	u.Lists = lists
	return nil
}
func (l *List) QueryResources(db *sqlx.DB) error {
	var movies []*Movie
	err := db.Select(&movies, "SELECT * FROM lists WHERE list_id=$1", l.ID)
	if err != nil {
		return err
	}
	l.Movies = movies
	return nil
}
func ParseQueryStruct(q string, d interface{}, db *sqlx.DB) (string, []interface{}, error) {
	errors := make([]error, 2)
	query, args, err := sqlx.Named(q, d)
	errors = append(errors, err)
	query, args, err = sqlx.In(query, args...)
	errors = append(errors, err)
	query = db.Rebind(query)
	for _, e := range errors {
		if e != nil {
			return "", nil, e
		}
	}
	return query, args, nil
}
