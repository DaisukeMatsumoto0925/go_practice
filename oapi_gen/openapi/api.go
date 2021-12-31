package openapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

// Returns all pets
// (GET /pets)
func (s *Handler) FindPets(ctx echo.Context, params FindPetsParams) error {
	type User struct {
		Name  string `json:"name" xml:"name"`
		Email string `json:"email" xml:"email"`
	}
	u := &User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	ctx.JSON(http.StatusOK, u)
	return nil
}

// Creates a new pet
// (POST /pets)
func (s *Handler) AddPet(ctx echo.Context) error {
	return nil
}

// Deletes a pet by ID
// (DELETE /pets/{id})
func (s *Handler) DeletePet(ctx echo.Context, id int64) error {
	return nil
}

// Returns a pet by ID
// (GET /pets/{id})
func (s *Handler) FindPetByID(ctx echo.Context, id int64) error {
	return nil
}
