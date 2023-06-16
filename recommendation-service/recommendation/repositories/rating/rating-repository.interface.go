package rating

import "recommendation_service/recommendation/model"

type IRatingRepository interface {
	Create(rating model.Rating) error
	Delete(rating model.Rating) error
	Update(rating model.Rating) error
}
