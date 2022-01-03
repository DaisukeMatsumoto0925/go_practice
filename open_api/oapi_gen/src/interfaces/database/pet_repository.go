package database

import (
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/model"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/repositories"
	"github.com/DaisukeMatsumoto0925/oapi_gen/src/interfaces"
)

type PetRepository struct {
	SQLHandler interfaces.SQLHandler
}

func NewPetRepository(sqlhandler interfaces.SQLHandler) repositories.Pet {
	return &PetRepository{
		SQLHandler: sqlhandler,
	}
}

func (r *PetRepository) GetAll(tx repositories.Transaction, userID int) (pets []*model.Pet, err error) {
	sqlhandler := r.SQLHandler.FromTransaction(tx)
	if err := sqlhandler.Where("user_id = ?").Find(pets).GetError(); err != nil {
		return nil, err
	}

	return pets, nil
}
