package handlers

import (
	. "auth_service/auth"
	"auth_service/auth/model"
	"auth_service/common/messaging"
	events "auth_service/common/rate_host"
	"auth_service/shared"
)

type DeleteHostCommandHandler struct {
	authService       *AuthService
	replyPublisher    messaging.Publisher
	commandSubscriber messaging.Subscriber
}

func NewDeleteHostCommandHandler(authService *AuthService, publisher messaging.Publisher, subscriber messaging.Subscriber) (*DeleteHostCommandHandler, error) {
	o := &DeleteHostCommandHandler{
		authService:       authService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
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
	case events.RemoveHostCommand:
		err := handler.authService.Delete(shared.StringToObjectId(command.Host))
		if err != nil {
			reply.Type = events.HostDeletionFailedReply
			break
		}
		reply.Type = events.HostDeletedReply
	case events.CancelDeleteHostCommand:
		user, _ := handler.authService.FindById(shared.StringToObjectId(command.Host))
		user.DeleteStatus = model.ACTIVE
		handler.authService.UpdatePersonalInfo(user)
		reply.Type = events.CancelDeleteHostReply

	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
