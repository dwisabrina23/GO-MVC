package controller

import (
	"book-rentals/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUSerController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, newResponseArray(*users))
}

func GetByIdUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetByIDUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return c.JSON(http.StatusOK, *user)
}

func AddNewUserController(c echo.Context) error {
	var userReq RequestUser
	//catch data dari form
	c.Bind(&userReq)
	//dan simpan data ke user
	result, err := database.AddNewUser(userReq.toModel())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, NewResponse(*result))
}

func UpdateUserController(c echo.Context) error {
	var userReq RequestUser
	c.Bind(&userReq)
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := database.UpdateUser(id, userReq.toModel())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, NewResponse(*result))
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err1 := database.GetByIDUser(id)
	result, err2 := database.DeleteUser(id, *user)

	if err1 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err1.Error(),
		})
	} else if err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": result})
}
