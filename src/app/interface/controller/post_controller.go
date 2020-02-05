package controller

import (
	"my-echo.com/src/app/interface/database"
	"my-echo.com/src/app/usecase"
)

type PostController struct {
	Interactor usecase.PostInteractor
}

func NewPostController(sqlHandler database.SQLHandler) *PostController {
	return &PostController{
		Interactor: usecase.PostInteractor{
			PostRepository: &database.PostRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *PostController) Create(c Context) error {
	input := new(usecase.PostAddInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(400, err)
	}
	err := controller.Interactor.Add(*input)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, input)
}

func (controller *PostController) Index(c Context) error {
	users, err := controller.Interactor.Posts()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, users)
}

func (controller *PostController) Show(c Context) error {
	id := c.Param("id")
	post, err := controller.Interactor.PostById(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, post)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

