package controllers

import (
	openapi "github.com/DaisukeMatsumoto0925/oapi_gen/openApi"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/services"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/interfaces"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/interfaces/database"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/interfaces/presenters"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/usecases"
	"github.com/labstack/echo/v4"
)

type PetController struct {
	Usecase   usecases.PetInteractor
	Presenter *presenters.PetPresenter
}

func NewPetController(sqlhandler interfaces.SQLHandler) *PetController {
	petRepo := database.NewPetRepository(sqlhandler)
	presenter := presenters.NewPet()
	petService := services.NewPetService(petRepo)
	interactor := usecases.NewPetInteractor(
		sqlhandler,
		petRepo,
		petService,
	)

	return &PetController{Usecase: interactor, Presenter: presenter}
}

func (pc *PetController) FindPets(ctx echo.Context, params openapi.FindPetsParams) error {
	return nil
}

// Creates a new pet
// (POST /pets)
func (pc *PetController) AddPet(ctx echo.Context) error {
	return nil
}

// Deletes a pet by ID
// (DELETE /pets/{id})
func (pc *PetController) DeletePet(ctx echo.Context, id int64) error {
	return nil
}

// Returns a pet by ID
// (GET /pets/{id})
func (pc *PetController) FindPetByID(ctx echo.Context, id int64) error {
	return nil
}
