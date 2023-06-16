package services

import (
	"recommendation_service/recommendation/model"
	. "recommendation_service/recommendation/repositories/accommodation"
	. "recommendation_service/recommendation/repositories/rating"
	. "recommendation_service/recommendation/repositories/user"
)

type RecommendationService struct {
	UserRepository          IUserRepository
	AccommodationRepository IAccommodationRepository
	RatingRepository        IRatingRepository
}

func (recommendationService *RecommendationService) GetRecommended(id string) ([]model.Accommodation, error) {
	return recommendationService.UserRepository.GetRecommended(model.User{Id: id})
}
func (recommendationService *RecommendationService) CreateUser(user model.User) error {
	return recommendationService.UserRepository.Create(user)
}
func (recommendationService *RecommendationService) DeleteUser(user model.User) error {
	return recommendationService.UserRepository.Delete(user)
}
func (recommendationService *RecommendationService) CreateAccommodation(accommodation model.Accommodation) error {
	return recommendationService.AccommodationRepository.Create(accommodation)
}
func (recommendationService *RecommendationService) DeleteAccommodation(accommodation model.Accommodation) error {
	return recommendationService.AccommodationRepository.Delete(accommodation)
}
func (recommendationService *RecommendationService) CreateRating(rating model.Rating) error {
	return recommendationService.RatingRepository.Create(rating)
}
func (recommendationService *RecommendationService) DeleteRating(rating model.Rating) error {
	return recommendationService.RatingRepository.Delete(rating)
}
func (recommendationService *RecommendationService) UpdateRating(rating model.Rating) error {
	return recommendationService.RatingRepository.Update(rating)
}
