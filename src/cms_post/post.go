package cms_post

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"micky-svr/db"
	"micky-svr/helper"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)

type Post struct {
	Id                  int    		`json:"id" xml:"id" form:"id" query:"id"`
	Title               string 		`json:"title" xml:"title" form:"title" query:"title"`
	Description         string 		`json:"description" xml:"description" form:"description" query:"description"`
	Session 			[]Session 	`json:"session" xml:"session" form:"session" query:"session"`
}

type Session struct {
	Id 			int    			`json:"id" xml:"id" form:"id" query:"id"`
	Content  	string 			`json:"content" xml:"content" form:"content" query:"content"`
	PostId  	string 			`json:"-" xml:"post_id" form:"post_id" query:"post_id"`
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



func GetPost(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		db, err := sql.Open("postgres", db.DbInfo())
		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusForbidden)
			return
		}

		sqlQuery := `SELECT * FROM mk_post;`

		postRows, err := db.QueryContext(ctx, sqlQuery)
		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusInternalServerError)
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
				helper.FailRequest(&w,"false", http.StatusInternalServerError)
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
		fmt.Println("run code inside")
		newPost := Post{}
		err = json.Unmarshal(body, &newPost)
		if err != nil {
			panic(err)
		}

		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusInternalServerError)
			return
		}

		defer db.Close()

		sqlQuery := `INSERT INTO mk_post (title, description) VALUES ($1, $2);`
		_, err = db.Exec(sqlQuery, newPost.Title, newPost.Description)

		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusInternalServerError)
			return
		}

		sqlQueryPost := `SELECT id FROM mk_post WHERE title=$1 AND description=$2 LIMIT 1;`
		row := db.QueryRowContext(ctx, sqlQueryPost, newPost.Title, newPost.Description)

		var id string
		err = row.Scan(
			&id,
		)

		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusInternalServerError)
			return
		}
		sqlInsert := `INSERT INTO mk_session (content, post_id) VALUES ($1, $2);`

		if len(newPost.Session) > 0{
			for _, session := range newPost.Session {
				_, err = db.Exec(sqlInsert, session.Content, id)
				if err != nil {
					panic(err)
					helper.FailRequest(&w, "false", http.StatusInternalServerError)
					return
				}
			}
		}


		response := map[string]string{"status": "ok"}
		js, _ := json.Marshal(response)
		helper.WriteResponse(&w, js)
		return

	} else {
		helper.FailRequest(&w, "false", http.StatusInternalServerError)
		return
	}
}

func SinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if r.Method == "GET" {
		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusForbidden)
			return
		}

		queryPost := `SELECT * FROM mk_post WHERE id=$1;`

		postRow := db.QueryRowContext(ctx,queryPost,id)
		post := Post{}
		err = postRow.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
		)

		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusForbidden)
			return
		}

		querySession := `SELECT * FROM mk_session WHERE post_id=$1;`

		sessionRows, err := db.QueryContext(ctx,querySession,id)
		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusForbidden)
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
			helper.FailRequest(&w, "false", http.StatusInternalServerError)
			return
		}

		defer db.Close()

		deletePostQuery := `DELETE FROM mk_post WHERE id=$1;`
		_, err = db.Exec(deletePostQuery, id)
		deleteSessionQuery := `DELETE FROM mk_session WHERE post_id=$1;`
		_, err = db.Exec(deleteSessionQuery, id)
	}

	return
}