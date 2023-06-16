package accommodation

import "recommendation_service/recommendation/model"

type IAccommodationRepository interface {
	Create(accommodation model.Accommodation) error
	Delete(accommodation model.Accommodation) error
}
