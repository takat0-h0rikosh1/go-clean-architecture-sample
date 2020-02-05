package domain

type Post struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id      string `json:"id" db:"id, primarykey"`
	Created int64  `json:"created" db:"created"`
	Title   string `json:"title" db:"title,size:50"`         // Column size set to 50
	Body    string `json:"body" db:"article_body,size:1024"` // Set both column name and size
}

type Posts []Post
