package route

import (
	"github.com/labstack/echo/v4"
	"github.com/yunkon-kim/my-template/pkg/api/rest/controller"
)


// /my-template/sample/*
func RegisterSampleRoutes(g *echo.Group) {
	g.GET("/users", controller.GetUsers)
	g.GET("/users/:id", controller.GetUser)
	g.POST("/users", controller.CreateUser)
	g.PUT("/users/:id", controller.UpdateUser)
	g.DELETE("/users/:id", controller.DeleteUser)
}
