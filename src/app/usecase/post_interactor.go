package usecase

import (
	"my-echo.com/src/app/domain"
)


type PostInteractor struct {
	PostRepository PostRepository
}

type PostAddInput struct {
	Title string
	Body string
}

func (interactor *PostInteractor) Add(input PostAddInput) (err error) {
	u, err := GeneratePost(input.Title, input.Body)
	_, err = interactor.PostRepository.Store(u)
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