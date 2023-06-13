package handlers

import (
	"reservation_service/common/messaging"
	events "reservation_service/common/rate_host"
	. "reservation_service/reservation"
)

type DeleteHostCommandHandler struct {
	reservationService *ReservationService
	replyPublisher     messaging.Publisher
	commandSubscriber  messaging.Subscriber
}

func NewDeleteHostCommandHandler(reservationService *ReservationService, publisher messaging.Publisher, subscriber messaging.Subscriber) (*DeleteHostCommandHandler, error) {
	o := &DeleteHostCommandHandler{
		reservationService: reservationService,
		replyPublisher:     publisher,
		commandSubscriber:  subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeleteHostCommandHandler) handle(command *events.DeleteHostCommand) {
	reply := events.DeleteHostReply{Host: command.Host, Accommodations: command.Accommodations}
	accommodationIds := make([]string, 0)

	for _, accommodation := range command.Accommodations {
		accommodationIds = append(accommodationIds, accommodation.Id.Hex())
	}

	switch command.Type {
	case events.ValidateReservationsCommand:
		reservationsExist, err := handler.reservationService.CheckActiveReservationsForAccommodations(accommodationIds)
		if err != nil || reservationsExist {
			reply.Type = events.ValidationFailedReply
			break
		}
		reply.Type = events.ValidationSuccessReply

	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
