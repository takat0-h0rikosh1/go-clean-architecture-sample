package controller

import (
	"github.com/google/uuid"
	"my-echo.com/src/app/domain"
	"time"
)

func GeneratePost(title, body string) (domain.Post, error) {
	p := new(domain.Post)
	id, err := uuid.NewRandom()
	if err != nil {
		return *p, err
	}
	p.Id = id.String()
	p.Created = time.Now().UnixNano()
	p.Title = title
	p.Body = body
	return *p, err
}

