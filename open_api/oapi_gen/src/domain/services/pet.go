package services

import "github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/repositories"

type petService struct {
	PetRepository repositories.Pet
}

type PetService interface {
}

func NewPetService(repo repositories.Pet) PetService {
	return &petService{
		PetRepository: repo,
	}
}
