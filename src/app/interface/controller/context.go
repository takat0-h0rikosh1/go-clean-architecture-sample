package controller

type Context interface {
	String(code int, s string) error
	Param(name string) string
	Bind(i interface{}) error
	JSON(code int, i interface{}) error
	Error(err error)
}