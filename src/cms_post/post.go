package cms_post

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"micky-svr/db"
	"encoding/json"
	//"database/sql"
	_ "github.com/lib/pq"
)

type Post struct {
	Id                  int    `json:"id" xml:"id" form:"id" query:"id"`
	Title               string `json:"title" xml:"title" form:"title" query:"title"`
	Description         string `json:"description" xml:"description" form:"description" query:"description"`
}
var ctx = context.Background()

func GetPost(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer db.Close()
		sqlQuery := `SELECT * FROM mk_post;`

		rows, err := db.QueryContext(ctx, sqlQuery)
		if err != nil {
			fmt.Println("loi o day nef ============>")
			panic(err)
			w.WriteHeader(http.StatusInternalServerError)
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
		fmt.Println("err =>",err)
		fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
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
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer db.Close()

		sqlQuery := `
			INSERT INTO mk_post (title, description)
			VALUES ($1, $2);`
		_, err = db.Exec(sqlQuery, title, description)

		if err != nil {
			panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]string{"status": "ok"}
		js, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return

	} else {
		return
	}
}