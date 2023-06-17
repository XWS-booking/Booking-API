package recommendation

import (
	. "context"
	"fmt"
	"log"
	"recommendation_service/messaging"
	. "recommendation_service/opentelementry"
	. "recommendation_service/proto/recommendation"
	"recommendation_service/recommendation/model"
	"recommendation_service/recommendation/services"
)

type RecommendationController struct {
	UnimplementedRecommendationServiceServer
	RecommendationService services.RecommendationService
}

func NewRecommendationController(
	recommendationService services.RecommendationService,
	accommodationListener messaging.Subscriber,
	userListener messaging.Subscriber,
	ratingListener messaging.Subscriber,
) *RecommendationController {
	controller := RecommendationController{RecommendationService: recommendationService}
	err := accommodationListener.Subscribe(controller.AccommodationEventHandler)
	if err != nil {
		log.Fatal("Accommodation event subscription failed!")
	}
	err = userListener.Subscribe(controller.UserEventHandler)
	if err != nil {
		log.Fatal("User event subscription failed!")
	}
	err = ratingListener.Subscribe(controller.RatingEventHandler)
	if err != nil {
		log.Fatal("Rating event subscription failed!")
	}
	return &controller
}

func (recommendationController *RecommendationController) AccommodationEventHandler(message *messaging.AccommodationMessage) {
	fmt.Println("Accommodation event triggered", message)
	switch message.Type {
	case messaging.CREATE:
		displayError(recommendationController.RecommendationService.CreateAccommodation(message.Accommodation))
		break
	case messaging.DELETE:
		displayError(recommendationController.RecommendationService.DeleteAccommodation(message.Accommodation))
		break
	}
}

func (recommendationController *RecommendationController) UserEventHandler(message *messaging.UserMessage) {
	fmt.Println("User event triggered", message)
	switch message.Type {
	case messaging.CREATE:
		displayError(recommendationController.RecommendationService.CreateUser(message.User))
		break
	case messaging.DELETE:
		displayError(recommendationController.RecommendationService.DeleteUser(message.User))
		break
	default:
		fmt.Println(message.Type, "not fitting")
		break
	}
}

func displayError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (recommendationController *RecommendationController) RatingEventHandler(message *messaging.RatingMessage) {
	fmt.Println("Rating event triggered", message)
	switch message.Type {
	case messaging.CREATE:
		displayError(recommendationController.RecommendationService.CreateRating(message.Rating))
		break
	case messaging.DELETE:
		displayError(recommendationController.RecommendationService.DeleteRating(message.Rating))
		break
	case messaging.UPDATE:
		displayError(recommendationController.RecommendationService.UpdateRating(message.Rating))
		break
	}
}

func (recommendationController *RecommendationController) GetRecommendedAccommodations(ctx Context, req *RecommendationRequest) (*RecommendationResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "getRecommendedAccommodations")
	defer func() { span.End() }()
	result, err := recommendationController.RecommendationService.GetRecommended(req.UserId)
	if result == nil {
		result = []model.Accommodation{}
	}

	if err != nil {
		return nil, err
	}

	ids := make([]string, 0)
	for _, single := range result {
		ids = append(ids, single.Id)
	}

	return &RecommendationResponse{
		Accommodations: ids,
	}, nil
}
