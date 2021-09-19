package routes

import (
	"book-rentals/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	//routing
	e.GET("/user", controller.GetAllUSerController)
	e.GET("/user/:id", controller.GetByIdUserController)
	e.POST("/user", controller.AddNewUserController)
	e.PUT("/user/:id", controller.UpdateUserController)
	e.DELETE("/user/:id", controller.DeleteUserController)
	return e
}
