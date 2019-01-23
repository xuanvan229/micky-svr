package cms_post

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"micky-svr/db"
	"micky-svr/helper"
	"encoding/json"
	//"database/sql"
	_ "github.com/lib/pq"
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
	PostId  	string 			`json:"post_id" xml:"post_id" form:"post_id" query:"post_id"`
}
var ctx = context.Background()

func GetPost(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusForbidden)
			return
		}

		defer db.Close()
		sqlQuery := `SELECT * FROM mk_post;`

		rows, err := db.QueryContext(ctx, sqlQuery)
		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusInternalServerError)
			return
		}
		posts := []Post{}
		for rows.Next() {
			post := Post{}
			err = rows.Scan(
				&post.Id,
				&post.Title,
				&post.Description,
			)
			posts = append(posts, post)
		}

		fmt.Println(posts)
		js, err := json.Marshal(posts)
		helper.WriteResponse(&w, js)
		return

	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		title := r.FormValue("title")
		description := r.FormValue("description")

		db, err := sql.Open("postgres", db.DbInfo())
		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusInternalServerError)
			return
		}

		defer db.Close()

		sqlQuery := `INSERT INTO mk_post (title, description) VALUES ($1, $2);`
		_, err = db.Exec(sqlQuery, title, description)

		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusInternalServerError)
			return
		}

		//sqlQueryPost = `SELECT id FROM mk_post WHERE title=$1 AND description=$2;`
		//row = db.QueryRowContext(ctx, sqlQuery, username)
		response := map[string]string{"status": "ok"}
		js, _ := json.Marshal(response)
		helper.WriteResponse(&w, js)
		return

	} else {
		helper.FailRequest(&w, "false", http.StatusInternalServerError)
		return
	}
}