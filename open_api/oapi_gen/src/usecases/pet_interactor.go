package usecases

import (
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/model"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/repositories"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/services"
)

type PetInteractor interface {
	GetPets(PetInput) ([]*model.Pet, error)
}

type petInteractor struct {
	SQLHandler    repositories.TransactionHandler
	petRepository repositories.Pet
	petService    services.PetService
}

func NewPetInteractor(
	sqlHandler repositories.TransactionHandler,
	petRepository repositories.Pet,
	petService services.PetService,
) PetInteractor {
	return &petInteractor{
		SQLHandler:    sqlHandler,
		petRepository: petRepository,
		petService:    petService,
	}
}

// Index
type PetInput struct {
	UserID int
}

func (ca *petInteractor) GetPets(input PetInput) ([]*model.Pet, error) {
	return ca.petRepository.GetAll(nil, input.UserID)
}
