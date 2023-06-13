package api

import (
	"context"
	"gateway/infrastructure/services"
	"gateway/model"
	"gateway/model/mapper"
	"gateway/proto/gateway"
	"gateway/shared"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type SearchAccommodationHandler struct {
	authClientAddress          string
	accommodationClientAddress string
	reservationClientAddress   string
	ratingClientAddress        string
}
type PriceParams struct {
	From int32 `json:"from"`
	To   int32 `json:"to"`
}

type FilterParams struct {
	Price     PriceParams `json:"price"`
	Additions []string    `json:"additions"`
}
type SearchResult struct {
	Data       []model.Accommodation `json:"data"`
	TotalCount int32                 `json:"totalCount"`
}

func NewSearchAccommodationHandler(authClientAddress, accommodationClientAddress, reservationClientAddress string, ratingClientAddress string) Handler {
	return &SearchAccommodationHandler{
		authClientAddress:          authClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		reservationClientAddress:   reservationClientAddress,
		ratingClientAddress:        ratingClientAddress,
	}
}

func (handler *SearchAccommodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/accommodations/search/{city}/{guests}/{startDate}/{endDate}/{pageSize}/{pageNumber}", handler.Search)
	if err != nil {
		panic(err)
	}
}

func (handler *SearchAccommodationHandler) Search(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	city, guests, startDate, endDate, pageSize, pageNumber, err := handlePathParams(pathParams)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	authClient := services.NewAuthClient(handler.authClientAddress)
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	var filterParams FilterParams
	shared.DecodeBody(r, &filterParams)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	includingIds, err := handler.FindAllReservedAccommodations(startDate, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := &gateway.SearchAndFilterRequest{
		City:         city,
		Guests:       int32(guests),
		IncludingIds: includingIds,
		Page:         int32(pageNumber),
		Limit:        int32(pageSize),
		Price: &gateway.PriceRange{
			From: float32(filterParams.Price.From),
			To:   float32(filterParams.Price.To),
		},
		Filters: filterParams.Additions,
	}
	filtered, err := accommodationClient.SearchAndFilter(context.TODO(), request)

	data := make([]model.Accommodation, 0)

	for _, single := range filtered.Data {
		owner, err := authClient.FindById(context.TODO(), &gateway.FindUserByIdRequest{Id: single.OwnerId})
		if err != nil {
			shared.BadRequest(w, "Something wrong with capturing owner")
			return
		}
		averageRating, err := ratingClient.GetAverageAccommodationRating(context.TODO(), &gateway.GetAverageAccommodationRatingRequest{AccommodationId: single.Id})
		if err != nil {
			shared.BadRequest(w, "Something wrong with capturing rating")
			return
		}
		data = append(data, mapper.AccommodationFromAccomodationResponse(single, mapper.UserFromFindUserByIdResponse(owner), averageRating.Rating))
	}

	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}

	shared.Ok(&w, SearchResult{Data: data, TotalCount: filtered.TotalCount})
	return
}

func (handler *SearchAccommodationHandler) FindAllReservedAccommodations(startDate, endDate *timestamp.Timestamp) ([]string, error) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	ids, err := reservationClient.FindAllReservedAccommodations(context.TODO(), &gateway.FindAllReservedAccommodationsRequest{StartDate: startDate, EndDate: endDate})
	if err != nil {
		return nil, err
	}
	return ids.Ids, nil
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
