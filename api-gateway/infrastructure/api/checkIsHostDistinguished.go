package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	. "gateway/model"
	"gateway/model/mapper"
	"gateway/proto/gateway"
	"math"
	"time"
)

type HostDistinguishedChecker struct {
	ratingClientAddress        string
	accommodationClientAddress string
	reservationClientAddress   string
	authClientAddress          string
	notificationClientAddress  string
}

func NewIsHostDistinguishedFunc(notificationClientAddress string, authClientAddress string, ratingClientAddress string, reservationClientAddress string, accommodationClientAddress string) *HostDistinguishedChecker {
	return &HostDistinguishedChecker{
		ratingClientAddress:        ratingClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		reservationClientAddress:   reservationClientAddress,
		notificationClientAddress:  notificationClientAddress,
		authClientAddress:          authClientAddress,
	}
}

func (checker *HostDistinguishedChecker) CheckIsHostDistinguishedFunc(id string) (*bool, error) {
	ratingClient := services.NewRatingClient(checker.ratingClientAddress)
	authClient := services.NewAuthClient(checker.authClientAddress)
	notificationClient := services.NewNotificationClient(checker.notificationClientAddress)
	host, err := authClient.FindById(context.TODO(), &gateway.FindUserByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	hostOldDistinguishedStatus := host.Distinguished
	reservations, err := checker.FindAllReservationsByOwner(id)

	if err != nil {
		return nil, err
	}
	hostRating, e := ratingClient.GetAverageHostRating(context.TODO(), &gateway.GetAverageHostRatingRequest{HostId: id})
	if e != nil {
		return nil, err
	}
	isHostDistinguished := false
	if !math.IsNaN(hostRating.Rating) {
		isThereLessThan5PercentCanceledReservations := checker.CheckIsThereLessThan5PercentCanceledReservations(reservations)
		fmt.Println(isThereLessThan5PercentCanceledReservations)
		isThereMoreThan50DaysOfReservations := checker.CheckIsThereMoreThan50DaysOfReservations(reservations)
		fmt.Println(isThereMoreThan50DaysOfReservations)
		isThereMoreThan5ReservationsInPast := checker.CheckIsThereMoreThan5ReservationsInPast(reservations)
		fmt.Println(isThereMoreThan5ReservationsInPast)
		fmt.Println(hostRating.Rating > 4.7)
		isHostDistinguished = hostRating.Rating > 4.7 &&
			isThereLessThan5PercentCanceledReservations &&
			isThereMoreThan50DaysOfReservations &&
			isThereMoreThan5ReservationsInPast
	}
	if hostOldDistinguishedStatus != isHostDistinguished {
		fmt.Println("usao sam!")
		authClient.ChangeHostDistinguishedStatus(context.TODO(), &gateway.ChangeHostDistinguishedStatusRequest{Id: id})
		if isHostDistinguished {
			_, err = notificationClient.SendNotification(context.TODO(), &gateway.SendNotificationRequest{NotificationType: "distinguished_host", UserId: id, Message: "You are now distinguished host!"})
		} else {
			_, err = notificationClient.SendNotification(context.TODO(), &gateway.SendNotificationRequest{NotificationType: "distinguished_host", UserId: id, Message: "You are no longed distinguished host!"})
		}
	}
	return &isHostDistinguished, nil
}

func (checker *HostDistinguishedChecker) CheckIsThereLessThan5PercentCanceledReservations(reservations []ReservationWithCancellation) bool {
	canceledReservations := make([]ReservationWithCancellation, 0)
	for _, reservation := range reservations {
		if reservation.Status == 3 {
			canceledReservations = append(canceledReservations, reservation)
		}
	}
	numberOfCanceledReservations := float64(len(canceledReservations))
	numberOfReservations := float64(len(reservations))
	return ((numberOfCanceledReservations / numberOfReservations) * 100) < 5
}

func (checker *HostDistinguishedChecker) CheckIsThereMoreThan50DaysOfReservations(reservations []ReservationWithCancellation) bool {
	totalDays := 0

	for _, reservation := range reservations {
		duration := reservation.EndDate.Sub(reservation.StartDate)
		days := int(duration.Hours() / 24)
		totalDays += days
	}
	return totalDays > 50
}

func (checker *HostDistinguishedChecker) CheckIsThereMoreThan5ReservationsInPast(reservations []ReservationWithCancellation) bool {
	now := time.Now().UTC()
	pastReservations := make([]ReservationWithCancellation, 0)
	for _, reservation := range reservations {
		startDateInPast := reservation.StartDate.Before(now)
		endDateInPast := reservation.EndDate.Before(now)
		if startDateInPast && endDateInPast {
			pastReservations = append(pastReservations, reservation)
		}
	}

	return len(pastReservations) > 5
}

func (checker *HostDistinguishedChecker) FindAllReservationsByOwner(id string) ([]ReservationWithCancellation, error) {
	accommodationClient := services.NewAccommodationClient(checker.accommodationClientAddress)
	accommodations, e := accommodationClient.FindAllAccommodationIdsByOwnerId(context.TODO(), &gateway.FindAllAccommodationIdsByOwnerIdRequest{OwnerId: id})
	if e != nil {
		return []ReservationWithCancellation{}, e
	}
	reservationClient := services.NewReservationClient(checker.reservationClientAddress)
	var reservationsWithAccommodation []ReservationWithCancellation
	for _, accommId := range accommodations.Ids {
		accommodation, e := accommodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{Id: accommId})
		if e != nil {
			return []ReservationWithCancellation{}, e
		}
		reservations, e := reservationClient.FindAllByAccommodationId(context.TODO(), &gateway.FindAllReservationsByAccommodationIdRequest{AccommodationId: accommId})
		if e != nil {
			return []ReservationWithCancellation{}, e
		}
		for _, reservation := range reservations.Reservations {
			numberOfCancellation, e := reservationClient.FindNumberOfBuyersCancellations(context.TODO(), &gateway.NumberOfCancellationRequest{BuyerId: reservation.BuyerId})
			if e != nil {
				return []ReservationWithCancellation{}, e
			}
			reservationsWithAccommodation = append(reservationsWithAccommodation, ReservationWithCancellation{
				Id:                   reservation.Id,
				Accommodation:        mapper.AccommodationFromAccomodationResponse(accommodation, User{}, 0),
				BuyerId:              reservation.BuyerId,
				StartDate:            reservation.StartDate.AsTime(),
				EndDate:              reservation.EndDate.AsTime(),
				Guests:               reservation.Guests,
				Status:               reservation.Status,
				NumberOfCancellation: numberOfCancellation.CancellationNumber,
			})
		}
	}
	return reservationsWithAccommodation, nil
}
