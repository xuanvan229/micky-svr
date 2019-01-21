package cms_session

type Session struct {
	Id      int    `json : "id"  xml: "id" form: "id" query: "id"`
	Title   string `json: "title" xml: "title" form : "title" query: "title"`
	Content string `json: "content" xml: "content" form : "content" query: "content"`
}
