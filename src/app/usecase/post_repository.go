package usecase

import "my-echo.com/src/app/domain"

type PostRepository interface {
	Store(post domain.Post) (s string, e error)
	ResolveBy(string) (domain.Post, error)
	ResolveAll() (domain.Posts, error)
}