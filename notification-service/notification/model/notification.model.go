package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	UserId                             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	GuestCreatedReservationRequest     bool               `bson:"guest_created_reservation_request" json:"guestCreatedReservationRequest"`
	GuestCanceledReservation           bool               `bson:"guest_canceled_reservation" json:"guestCanceledReservation"`
	GuestRatedHost                     bool               `bson:"guest_rated_host" json:"guestRatedHost"`
	GuestRatedAccommodation            bool               `bson:"guest_rated_accommodation" json:"guestRatedAccommodation"`
	DistinguishedHost                  bool               `bson:"distinguished_host" json:"distinguishedHost"`
	HostConfirmedOrRejectedReservation bool               `bson:"host_confirmed_or_rejected_reservation" json:"hostConfirmedOrRejectedReservation"`
}
