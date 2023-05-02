package api

import (
	"context"
	"encoding/json"
	"gateway/infrastructure/services"
	"gateway/model"
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
	err := mux.HandlePath("GET", "/api/searchAccommodation/{city}/{guests}/{startDate}/{endDate}/{pageSize}/{pageNumber}", handler.Search)
	if err != nil {
		panic(err)
	}
}

func (handler *SearchAccommodationHandler) Search(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	city, guests, startDate, endDate, pageSize, pageNumber, err := handlePathParams(pathParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	accommodations, err := handler.SearchByCityAndGuests(city, guests)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	reservedAccommodationIds, err := handler.FindAllReservedAccommodations(startDate, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	removeIds := make(map[string]bool)
	for _, id := range reservedAccommodationIds {
		removeIds[id] = true
	}

	availableAccommodations := []gateway.AccomodationResponse{}
	for _, obj := range accommodations.AccomodationResponses {
		if !removeIds[obj.Id] {
			availableAccommodations = append(availableAccommodations, *obj)
		}
	}

	data := pagination(pageSize, pageNumber, availableAccommodations)

	response, err := json.Marshal(model.AccommodationPage{Data: data, TotalCount: len(availableAccommodations)})
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

func (handler *SearchAccommodationHandler) FindAllReservedAccommodations(startDate, endDate *timestamp.Timestamp) ([]string, error) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	ids, err := reservationClient.FindAllReservedAccommodations(context.TODO(), &gateway.FindAllReservedAccommodationsRequest{StartDate: startDate, EndDate: endDate})
	if err != nil {
		return nil, err
	}
	return ids.Ids, nil
}

func pagination(pageSize int, pageNumber int, accommodations []gateway.AccomodationResponse) []model.Accommodation {
	startIndex := (pageNumber - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > len(accommodations) {
		endIndex = len(accommodations)
	}
	paginationData := accommodations[startIndex:endIndex]
	var data []model.Accommodation
	for _, e := range paginationData {
		data = append(data, model.Accommodation{
			Id:             e.Id,
			Name:           e.Name,
			Street:         e.Street,
			StreetNumber:   e.StreetNumber,
			City:           e.City,
			ZipCode:        e.ZipCode,
			Country:        e.Country,
			Wifi:           e.Wifi,
			Kitchen:        e.Kitchen,
			AirConditioner: e.AirConditioner,
			FreeParking:    e.FreeParking,
			MinGuests:      e.MinGuests,
			MaxGuests:      e.MaxGuests,
			PictureUrls:    e.Pictures,
			OwnerId:        e.OwnerId,
		})
	}
	return data
}

func handlePathParams(pathParams map[string]string) (string, int, *timestamp.Timestamp, *timestamp.Timestamp, int, int, error) {
	city := pathParams["city"]
	guests, err := strconv.Atoi(pathParams["guests"])
	if err != nil {
		return city, guests, nil, nil, -1, -1, err
	}
	startDate, err := parseTimestamp(pathParams["startDate"])
	if err != nil {
		return city, guests, startDate, nil, -1, -1, err
	}
	endDate, err := parseTimestamp(pathParams["endDate"])
	if err != nil {
		return city, guests, startDate, endDate, -1, -1, err
	}
	pageSize, err := strconv.Atoi(pathParams["pageSize"])
	if err != nil {
		return city, guests, startDate, endDate, pageSize, -1, err
	}
	pageNumber, err := strconv.Atoi(pathParams["pageNumber"])
	if err != nil {
		return city, guests, startDate, endDate, pageSize, pageNumber, err
	}
	return city, guests, startDate, endDate, pageSize, pageNumber, nil
}

func parseTimestamp(str string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return nil, err
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}
