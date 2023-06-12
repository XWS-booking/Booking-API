package saga_config

import (
	messaging "auth_service/common/messaging"
	events "auth_service/common/rate_host"
	"fmt"
)

type DeleteHostOrchestrator struct {
	commandPublisher messaging.Publisher
	replySubscriber  messaging.Subscriber
}

func NewDeleteHostOrchestrator(publisher messaging.Publisher, subscriber messaging.Subscriber) (*DeleteHostOrchestrator, error) {
	fmt.Println("inicijalno ", publisher, subscriber)
	o := &DeleteHostOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	fmt.Println("nakon inicijalno", o)
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *DeleteHostOrchestrator) Start(host string) error {
	event := &events.DeleteHostCommand{
		Type: events.FetchHostAccommodationCommand,
		Host: host,
	}
	return o.commandPublisher.Publish(event)
}
func (o *DeleteHostOrchestrator) handle(reply *events.DeleteHostReply) {
	command := events.DeleteHostCommand{Host: reply.Host, Accommodations: reply.Accommodations}
	command.Type = o.nextCommandType(reply.Type)

	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *DeleteHostOrchestrator) nextCommandType(reply events.DeleteHostReplyType) events.DeleteHostCommandType {
	switch reply {
	case events.FetchHostAccommodationSuccessReply:
		return events.ValidateReservationsCommand
	case events.ValidationSuccessReply:
		return events.DeleteAccommodationsCommand
	case events.AccommodationsDeletedReply:
		return events.RemoveHostCommand
	case events.FetchHostAccommodationFailedReply:
		return events.CancelDeleteHostCommand
	case events.ValidationFailedReply:
		return events.CancelDeleteHostCommand
	case events.AccommodationDeletionFailedReply:
		return events.CancelDeleteHostCommand
	case events.HostDeletionFailedReply:
		return events.RollbackDeleteAccommodationsCommand
	case events.RollbackDeleteAccommodationsReply:
		return events.CancelDeleteHostCommand
	default:
		return events.UnknownCommand
	}
}
