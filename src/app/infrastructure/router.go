package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"my-echo.com/src/app/interface/controller"
	"net/http"
)

// Echo instance
var Router = echo.New()
var postController = controller.NewPostController(newSQLHandler())

func init() {

	// Middleware
	Router.Use(middleware.Logger())
	Router.Use(middleware.Recover())

	// Routes
	Router.GET("/healthCheck", healthCheck)
	Router.POST("/posts", create)
	Router.GET("/posts", index)
	Router.GET("/posts/:id", show)
}

// Handler
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "I'm alive")
}

func create(c echo.Context) error {
	return postController.Create(c)
}

func index(c echo.Context) error {
	return postController.Index(c)
}

func show(c echo.Context) error {
	return postController.Show(c)
}
