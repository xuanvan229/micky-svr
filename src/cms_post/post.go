package cms_post

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"micky-svr/db"
	"micky-svr/helper"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Post struct {
	Id          int       `json:"id" xml:"id" form:"id" query:"id"`
	Title       string    `json:"title" xml:"title" form:"title" query:"title"`
	Description string    `json:"description" xml:"description" form:"description" query:"description"`
	Session     []Session `json:"session" xml:"session" form:"session" query:"session"`
}

type Session struct {
	Id      int    `json:"id" xml:"id" form:"id" query:"id"`
	Content string `json:"content" xml:"content" form:"content" query:"content"`
	PostId  string `json:"-" xml:"post_id" form:"post_id" query:"post_id"`
}

var ctx = context.Background()

func CheckJson(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	newPost := Post{}
	err = json.Unmarshal(body, &newPost)
	if err != nil {
		panic(err)
	}
	fmt.Println(newPost)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db, err := sql.Open("postgres", db.DbInfo())
		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusForbidden)
			return
		}

		sqlQuery := `SELECT * FROM mk_post;`

		postRows, err := db.QueryContext(ctx, sqlQuery)
		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusInternalServerError)
			return
		}

		posts := []Post{}

		for postRows.Next() {
			post := Post{}
			err = postRows.Scan(
				&post.Id,
				&post.Title,
				&post.Description,
			)

			sqlQuerySession := `SELECT * FROM mk_session WHERE post_id=$1;`
			sessionRows, err := db.QueryContext(ctx, sqlQuerySession, post.Id)
			if err != nil {
				panic(err)
				helper.SetResponse(&w, "false", http.StatusInternalServerError)
				return
			}

			sessions := []Session{}
			for sessionRows.Next() {
				session := Session{}
				err = sessionRows.Scan(
					&session.Id,
					&session.Content,
					&session.PostId,
				)
				sessions = append(sessions, session)
			}
			post.Session = sessions
			posts = append(posts, post)
		}

		db.Close()

		response := map[string][]Post{"data": posts}
		js, err := json.Marshal(response)
		helper.WriteResponse(&w, js)
		return

	} else if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		newPost := Post{}
		err = json.Unmarshal(body, &newPost)
		if err != nil {
			panic(err)
		}

		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusInternalServerError)
			return
		}

		defer db.Close()

		sqlQuery := `INSERT INTO mk_post (title, description) VALUES ($1, $2) RETURNING *;`
		insertRow := db.QueryRow(sqlQuery, newPost.Title, newPost.Description)
    
		insertPost := Post{}
		err = insertRow.Scan(
			&insertPost.Id,
			&insertPost.Title,
			&insertPost.Description,
		)
		fmt.Println(insertPost)
		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusInternalServerError)
			return
		}
		sqlInsert := `INSERT INTO mk_session (content, post_id) VALUES ($1, $2);`

		if len(newPost.Session) > 0 {
			for _, session := range newPost.Session {
				_, err = db.Exec(sqlInsert, session.Content, insertPost.Id)
				if err != nil {
					panic(err)
					helper.SetResponse(&w, "false", http.StatusInternalServerError)
					return
				}
			}
		}
		
		sqlSelectSession := `SELECT * FROM mk_session WHERE post_id=$1;`
		sessionRows, err := db.QueryContext(ctx, sqlSelectSession, insertPost.Id)
		sessions := []Session{}
			for sessionRows.Next() {
				session := Session{}
				err = sessionRows.Scan(
					&session.Id,
					&session.Content,
					&session.PostId,
				)
				sessions = append(sessions, session)
			}
			insertPost.Session = sessions

		response := map[string]Post{"status": insertPost}
		js, _ := json.Marshal(response)
		helper.WriteResponse(&w, js)
		return

	} else {
		helper.SetResponse(&w, "false", http.StatusInternalServerError)
		return
	}
}

func SinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_string := vars["id"]
	id, _ := strconv.Atoi(id_string)
	if r.Method == "GET" {
		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusForbidden)
			return
		}

		queryPost := `SELECT * FROM mk_post WHERE id=$1;`

		postRow := db.QueryRowContext(ctx, queryPost, id)
		post := Post{}
		err = postRow.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
		)

		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusForbidden)
			return
		}

		querySession := `SELECT * FROM mk_session WHERE post_id=$1;`

		sessionRows, err := db.QueryContext(ctx, querySession, id)
		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusForbidden)
			return
		}

		sessions := []Session{}

		for sessionRows.Next() {
			session := Session{}
			err = sessionRows.Scan(
				&session.Id,
				&session.Content,
				&session.PostId,
			)
			sessions = append(sessions, session)
		}

		post.Session = sessions
		response := map[string]Post{"data": post}
		js, err := json.Marshal(response)
		helper.WriteResponse(&w, js)
		return
	} else if r.Method == "DELETE" {
		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusInternalServerError)
			return
		}

		deleteSessionQuery := `DELETE FROM mk_session WHERE post_id=$1;`
		_, err = db.Exec(deleteSessionQuery, id)
		if err != nil {
			panic(err)
			helper.SetResponse(&w, "BadRequest", http.StatusBadRequest)
			return
		}
		defer db.Close()
		deletePostQuery := `DELETE FROM mk_post WHERE id=$1;`
		_, err = db.Exec(deletePostQuery, id)
		if err != nil {
			panic(err)
			helper.SetResponse(&w, "BadRequest", http.StatusBadRequest)
			return
		}
		helper.SetResponse(&w, "Susscess", http.StatusOK)
		return

	} else if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
			helper.SetResponse(&w, "BadRequest", http.StatusBadRequest)
			return
		}

		newPost := Post{}
		err = json.Unmarshal(body, &newPost)

		if err != nil {
			panic(err)
		}

		if id != newPost.Id {
			panic(err)
			helper.SetResponse(&w, "BadRequest", http.StatusBadRequest)
			return
		}

		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusInternalServerError)
			return
		}

		updatePost := `UPDATE mk_post SET title=$1, description=$2 WHERE id=$3;`
		_, err = db.Exec(updatePost, newPost.Title, newPost.Description, newPost.Id)

		updateSession := `UPDATE mk_session SET content=$1 WHERE id=$2;`
		insertSession := `INSERT INTO mk_session (content, post_id) VALUES  ($1,$2);`
		if len(newPost.Session) > 0 {
			for _, session := range newPost.Session {
				if session.Id == 0 {
					_, err = db.Exec(insertSession, session.Content, id)
					if err != nil {
						panic(err)
						helper.SetResponse(&w, "false", http.StatusInternalServerError)
					}
				} else {
					_, err = db.Exec(updateSession, session.Content, session.Id)
					if err != nil {
						panic(err)
						helper.SetResponse(&w, "false", http.StatusInternalServerError)
						return
					}
				}
			}
		}

		helper.SetResponse(&w, "Susscess", http.StatusOK)
		return
	}

	return
}
