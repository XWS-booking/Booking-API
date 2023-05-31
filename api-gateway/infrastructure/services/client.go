package services

import (
	. "gateway/proto/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewAuthClient(address string) AuthServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Auth service: %v", err)
	}
	return NewAuthServiceClient(conn)
}

func NewAccommodationClient(address string) AccomodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return NewAccomodationServiceClient(conn)
}

func NewReservationClient(address string) ReservationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Reservation service: %v", err)
	}
	return NewReservationServiceClient(conn)
}

func NewRatingClient(address string) RatingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Rating service: %v", err)
	}
	return NewRatingServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
