package handlers

import (
	. "accomodation_service/accomodation"
	"accomodation_service/accomodation/model"
	"accomodation_service/common/messaging"
	events "accomodation_service/common/rate_host"
	"accomodation_service/shared"
)

type DeleteHostCommandHandler struct {
	accommodationService *AccomodationService
	replyPublisher       messaging.Publisher
	commandSubscriber    messaging.Subscriber
}

func NewDeleteHostCommandHandler(accommodationService *AccomodationService, publisher messaging.Publisher, subscriber messaging.Subscriber) (*DeleteHostCommandHandler, error) {
	o := &DeleteHostCommandHandler{
		accommodationService: accommodationService,
		replyPublisher:       publisher,
		commandSubscriber:    subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeleteHostCommandHandler) handle(command *events.DeleteHostCommand) {
	reply := events.DeleteHostReply{Host: command.Host, Accommodations: command.Accommodations}

	switch command.Type {
	case events.FetchHostAccommodationCommand:
		accommodations, err := handler.accommodationService.FindAllByOwnerId(shared.StringToObjectId(command.Host))
		if err != nil {
			reply.Type = events.FetchHostAccommodationFailedReply
			break
		}
		reply.Accommodations = mapAccommodations(accommodations)
		reply.Type = events.FetchHostAccommodationSuccessReply
	case events.DeleteAccommodationsCommand:
		err := handler.accommodationService.DeleteByOwnerId(shared.StringToObjectId(command.Host))
		if err != nil {
			reply.Type = events.AccommodationDeletionFailedReply
			break
		}
		reply.Type = events.AccommodationsDeletedReply
	case events.RollbackDeleteAccommodationsCommand:
		for _, accommodation := range reply.Accommodations {
			mappedAcc := model.Accomodation{
				Id:              accommodation.Id,
				OwnerId:         accommodation.OwnerId,
				Name:            accommodation.Name,
				Street:          accommodation.Street,
				StreetNumber:    accommodation.StreetNumber,
				City:            accommodation.City,
				ZipCode:         accommodation.ZipCode,
				Country:         accommodation.Country,
				Wifi:            accommodation.Wifi,
				Kitchen:         accommodation.Kitchen,
				AirConditioner:  accommodation.AirConditioner,
				FreeParking:     accommodation.FreeParking,
				AutoReservation: accommodation.AutoReservation,
				MinGuests:       accommodation.MinGuests,
				MaxGuests:       accommodation.MaxGuests,
				PictureUrls:     accommodation.PictureUrls,
				Pricing:         mapPricingEventToModel(accommodation.Pricing),
			}
			handler.accommodationService.Create(mappedAcc)
		}
		reply.Type = events.RollbackDeleteAccommodationsReply
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
