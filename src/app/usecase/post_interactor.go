package usecase

import (
	"my-echo.com/src/app/domain"
)


type PostInteractor struct {
	PostRepository PostRepository
}

type PostCreateInput struct {
	Title string
	Body string
}

func (interactor *PostInteractor) Add(post domain.Post) (err error) {
	_, err = interactor.PostRepository.Store(post)
	return
}

func (interactor *PostInteractor) Posts() (posts domain.Posts, err error) {
	posts, err = interactor.PostRepository.ResolveAll()
	return
}

func (interactor *PostInteractor) PostById(identifier string) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.ResolveBy(identifier)
	return
}