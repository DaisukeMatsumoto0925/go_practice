package repositories

import "github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/model"

type Pet interface {
	GetAll(tx Transaction, userID int) ([]*model.Pet, error)
}
