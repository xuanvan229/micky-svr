package cms_post

import (
	"micky-svr/cms_session"
)

type Post struct {
	Id                  int    `json : "id"  xml: "id" form: "id" query: "id"`
	Title               string `json: "title" xml: "title" form : "title" query: "title"`
	cms_session.Session `json: "session" xml: "session" form : "session" query: "session"`
}
