package database

import (
	"log"
	"my-echo.com/src/app/domain"
)

type PostRepository struct {
	SQLHandler
}

func (repo *PostRepository) Store(u domain.Post) (id string, err error) {
	err = repo.Execute("INSERT INTO \"Posts\" VALUES ($1,$2,$3,$4)", u.Id, u.Created, u.Title, u.Body)
	if err != nil {
		return
	}
	return
}

func (repo *PostRepository) ResolveBy(identifier string) (Post domain.Post, err error) {
	row, err := repo.Query("SELECT * FROM \"Posts\" WHERE id = $1", identifier)
	defer func() {
		err := row.Close()
		if err != nil {
			log.Fatal()
		}
	}()
	Post = domain.Post{}
	row.Next()
	if err = row.Scan(&Post.Id, &Post.Created, &Post.Title, &Post.Body); err != nil {
		return
	}
	return
}

func (repo *PostRepository) ResolveAll() (Posts domain.Posts, err error) {
	rows, err := repo.Query("SELECT * from \"Posts\"")
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal()
		}
	}()
	if err != nil {
		return
	}
	for rows.Next() {
		Post := domain.Post{}
		if err = rows.Scan(&Post.Id, &Post.Created, &Post.Title, &Post.Body); err != nil {
			return
		}
		Posts = append(Posts, Post)
	}
	return
}
