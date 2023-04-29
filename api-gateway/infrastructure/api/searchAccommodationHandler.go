package api

import (
	"context"
	"encoding/json"
	"gateway/infrastructure/services"
	"gateway/proto/gateway"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"strconv"
	"time"
)

type SearchAccommodationHandler struct {
	accommodationClientAddress string
	reservationClientAddress   string
}

func NewSearchAccommodationHandler(accommodationClientAddress, reservationClientAddress string) Handler {
	return &SearchAccommodationHandler{
		accommodationClientAddress: accommodationClientAddress,
		reservationClientAddress:   reservationClientAddress,
	}
}

func (handler *SearchAccommodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/searchAccommodation/{city}/{guests}/{startDate}/{endDate}", handler.Search)
	if err != nil {
		panic(err)
	}
}

func (handler *SearchAccommodationHandler) Search(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	city := pathParams["city"]
	guests, _ := strconv.Atoi(pathParams["guests"])
	startDate := pathParams["startDate"]
	endDate := pathParams["endDate"]

	accommodations, err := handler.SearchByCityAndGuests(city, guests)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ids, err := handler.SearchByDateRange(startDate, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	removeIds := make(map[string]bool)
	for _, id := range ids {
		removeIds[id] = true
	}

	availableAccommodations := []gateway.AccomodationResponse{}
	for _, obj := range accommodations.AccomodationResponses {
		if !removeIds[obj.Id] {
			availableAccommodations = append(availableAccommodations, *obj)
		}
	}

	response, err := json.Marshal(availableAccommodations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *SearchAccommodationHandler) SearchByCityAndGuests(city string, guests int) (*gateway.FindAllAccomodationResponse, error) {
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	accommodations, err := accommodationClient.FindAll(context.TODO(), &gateway.FindAllAccomodationRequest{City: city, Guests: int32(guests)})
	if err != nil {
		return &gateway.FindAllAccomodationResponse{}, err
	}
	return accommodations, nil
}

func (handler *SearchAccommodationHandler) SearchByDateRange(startDate, endDate string) ([]string, error) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	ids, err := reservationClient.FindAllReservedAccommodations(context.TODO(), &gateway.FindAllReservedAccommodationsRequest{StartDate: parseTimestamp(startDate), EndDate: parseTimestamp(endDate)})
	if err != nil {
		return nil, err
	}
	return ids.Ids, nil
}

func parseTimestamp(str string) *timestamp.Timestamp {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return nil
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil
	}
	return ts
}
