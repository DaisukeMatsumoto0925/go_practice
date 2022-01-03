package controllers

import (
	openapi "github.com/DaisukeMatsumoto0925/oapi_gen/openApi"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/interfaces"
	"github.com/labstack/echo/v4"
)

type UserController struct {
}

func NewUserController(sqlhandler interfaces.SQLHandler) *UserController {
	return &UserController{}
}

func (u *UserController) FindUsers(ctx echo.Context, params openapi.FindUsersParams) error {
	return nil
}
