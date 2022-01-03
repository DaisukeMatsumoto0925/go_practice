package presenters

import "github.com/DaisukeMatsumoto0925/oapi_gen/src/domain/model"

type PetPresenter struct{}

func NewPet() *PetPresenter {
	return &PetPresenter{}
}

type PetResponse struct {
	Name string
	Age  int
}

func (c *PetPresenter) FindPets(pets []*model.Pet) []*PetResponse {

	resPets := make([]*PetResponse, len(pets))
	for i, p := range pets {
		petRes := &PetResponse{
			Name: p.Name,
			Age:  p.Age,
		}
		resPets[i] = petRes

	}

	return resPets
}
