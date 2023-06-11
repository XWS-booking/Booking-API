package saga_config

import (
	messaging "github.com/XWS-booking/Booking-API/tree/feature/saga-rate-host/common/messaging"
	events "github.com/XWS-booking/Booking-API/tree/feature/saga-rate-host/common/rate_host"
)

type DeleteHostOrchestrator struct {
	commandPublisher messaging.Publisher
	replySubscriber  messaging.Subscriber
}

func NewCreateOrderOrchestrator(publisher messaging.Publisher, subscriber messaging.Subscriber) (*DeleteHostOrchestrator, error) {
	o := &CreateOrderOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *CreateOrderOrchestrator) Start(order *domain.Order, address string) error {
	event := &events.CreateOrderCommand{
		Type: events.UpdateInventory,
		Order: events.OrderDetails{
			Id:      order.Id.Hex(),
			Items:   make([]events.OrderItem, 0),
			Address: address,
		},
	}
	for _, item := range order.Items {
		eventItem := events.OrderItem{
			Product: events.Product{
				Id:    item.Product.Id,
				Color: events.Color{Code: item.Product.Color.Code},
			},
			Quantity: item.Quantity,
		}
		event.Order.Items = append(event.Order.Items, eventItem)
	}
	return o.commandPublisher.Publish(event)
}

func (o *CreateOrderOrchestrator) handle(reply *events.CreateOrderReply) {
	command := events.CreateOrderCommand{Order: reply.Order}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *CreateOrderOrchestrator) nextCommandType(reply events.CreateOrderReplyType) events.CreateOrderCommandType {
	switch reply {
	case events.InventoryUpdated:
		return events.ShipOrder
	case events.InventoryNotUpdated:
		return events.CancelOrder
	case events.InventoryRolledBack:
		return events.CancelOrder
	case events.OrderShippingScheduled:
		return events.ApproveOrder
	case events.OrderShippingNotScheduled:
		return events.RollbackInventory
	default:
		return events.UnknownCommand
	}
}