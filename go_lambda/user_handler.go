package example

import (
	"fmt"

	"github.com/DaisukeMatsumoto0925/example/gen/models"
	"github.com/DaisukeMatsumoto0925/example/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func GetUsers(p operations.GetUsersParams) middleware.Responder {
	ctx := p.HTTPRequest.Context()
	users, err := scanUsers(ctx)
	if err != nil {
		return operations.NewGetUsersInternalServerError().WithPayload(&models.Error{
			Message: fmt.Sprintf("scan users error: %v", err),
		})
	}
	var resp models.Users
	for _, u := range users {
		u := u
		resp = append(resp, &models.User{
			UserID: &u.UserID,
			Name:   &u.UserName,
		})
	}
	return operations.NewGetUsersOK().WithPayload(resp)
}
