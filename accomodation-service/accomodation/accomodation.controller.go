package accomodation

import (
	"accomodation_service/accomodation/dtos"
	"accomodation_service/accomodation/model"
	"accomodation_service/accomodation/services/storage"
	. "accomodation_service/proto/accomodation"
	"accomodation_service/shared"
	. "context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"fmt"
	"github.com/google/uuid"

	. "accomodation_service/opentelementry"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAccomodationController(accomodationService *AccomodationService, storageService storage.IStorageService) *AccomodationController {
	controller := &AccomodationController{AccomodationService: accomodationService, StorageService: storageService}
	return controller
}

type AccomodationController struct {
	UnimplementedAccomodationServiceServer
	AccomodationService *AccomodationService
	StorageService      storage.IStorageService
}

func (accomodationController *AccomodationController) Create(ctx Context, req *CreateAccomodationRequest) (*CreateAccomodationResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "create")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with accomodation data")
	}

	data := req.Pictures
	urls := make([]string, 0)
	for _, info := range data {
		uid := uuid.New()
		info.Filename = uid.String() + "-" + info.Filename
		url, err := accomodationController.StorageService.UploadImage(info.Data, info.Filename)
		if err != nil {

			fmt.Println("Something wrong when uploading!", err)
			return nil, status.Error(codes.Aborted, "File upload failed!")
		}
		urls = append(urls, url)
	}

	pricing := make([]model.Pricing, 0)

	for _, pr := range req.Pricing {
		single := PricingRequestToPricing(*pr)
		pricing = append(pricing, *single)
	}

	accomodationDto := dtos.AccomodationDto{
		Name:            req.Name,
		Street:          req.Street,
		StreetNumber:    req.StreetNumber,
		City:            req.City,
		ZipCode:         req.ZipCode,
		Country:         req.Country,
		Wifi:            req.Wifi,
		Kitchen:         req.Kitchen,
		AirConditioner:  req.AirConditioner,
		AutoReservation: req.AutoReservation,
		FreeParking:     req.FreeParking,
		MinGuests:       req.MinGuests,
		MaxGuests:       req.MaxGuests,
		OwnerId:         shared.StringToObjectId(req.OwnerId),
		Pricing:         pricing,
	}

	accomodation := AccomodationDtoToAccomodation(accomodationDto)
	accomodation.PictureUrls = urls

	created, e := accomodationController.AccomodationService.Create(accomodation)

	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}

	return &CreateAccomodationResponse{
		Id: created.Id.Hex(),
	}, nil
}

func (accomodationController *AccomodationController) FindAll(ctx Context, req *FindAllAccomodationRequest) (*FindAllAccomodationResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "findAll")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	accomodations, e := accomodationController.AccomodationService.FindAll(req.GetCity(), req.GetGuests())
	if e != nil {
		return nil, status.Error(codes.Internal, e.Message)
	}
	var accomodationResponses []*AccomodationResponse
	for _, a := range accomodations {
		accomodationResponses = append(accomodationResponses, NewAccomodationResponse(a))
	}
	return &FindAllAccomodationResponse{AccomodationResponses: accomodationResponses}, nil
}

func (accomodationController *AccomodationController) FindAllAccommodationIdsByOwnerId(ctx Context, req *FindAllAccommodationIdsByOwnerIdRequest) (*FindAllAccommodationIdsByOwnerIdResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "findAllAccommodationIdsByOwnerId")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.GetOwnerId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	accommodations, e := accomodationController.AccomodationService.FindAllByOwnerId(id)
	if e != nil {
		return nil, status.Error(codes.Internal, e.Message)
	}
	var accommodationIds []string
	for _, a := range accommodations {
		accommodationIds = append(accommodationIds, a.Id.Hex())
	}
	return &FindAllAccommodationIdsByOwnerIdResponse{Ids: accommodationIds}, nil
}

func (accomodationController *AccomodationController) DeleteByOwnerId(ctx Context, req *DeleteByOwnerIdRequest) (*DeleteByOwnerIdResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "deleteByOwnerId")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.GetOwnerId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	e := accomodationController.AccomodationService.DeleteByOwnerId(id)
	if e != nil {
		return nil, status.Error(codes.Internal, e.Message)
	}
	return &DeleteByOwnerIdResponse{Deleted: true}, nil
}

func (accomodationController *AccomodationController) FindById(ctx Context, req *FindAccommodationByIdRequest) (*AccomodationResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "findById")
	defer func() { span.End() }()

	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	accommodation, e := accomodationController.AccomodationService.FindById(id)
	if e != nil {
		return nil, status.Error(codes.Internal, e.Message)
	}
	return NewAccomodationResponse(accommodation), nil
}

func (accomodationController *AccomodationController) UpdatePricing(ctx Context, req *UpdatePricingRequest) (*UpdatePricingResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "updatePricing")
	defer func() { span.End() }()

	id := shared.StringToObjectId(req.Id)
	pricing := make([]model.Pricing, 0)

	for _, pr := range req.Pricing {
		single := PricingRequestToPricing(*pr)
		pricing = append(pricing, *single)
	}
	acc := model.Accomodation{
		Pricing: pricing,
		OwnerId: shared.StringToObjectId(req.UserId),
		Id:      id,
	}
	err := accomodationController.AccomodationService.UpdatePricing(acc)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Message)
	}
	return &UpdatePricingResponse{}, nil
}

func (accomodationController *AccomodationController) GetBookingPrice(ctx Context, req *GetBookingPriceRequest) (*GetBookingPriceResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "getBookingPrice")
	defer func() { span.End() }()

	from := time.Unix(req.From.Seconds, int64(req.From.Nanos)).UTC()
	to := time.Unix(req.To.Seconds, int64(req.To.Nanos)).UTC()
	interval := model.TimeInterval{From: from, To: to}
	params := model.BookingPriceParams{
		Guests:         req.Guests,
		Interval:       interval,
		AccomodationId: shared.StringToObjectId(req.AccomodationId),
	}
	fmt.Println(params)

	price, err := accomodationController.AccomodationService.GetBookingPrice(params)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Message)
	}

	return &GetBookingPriceResponse{
		Price: price,
	}, nil
}

func (accomodationController *AccomodationController) SearchAndFilter(ctx Context, req *SearchAndFilterRequest) (*SearchAndFilterResponse, error) {

	params := dtos.SearchDto{
		Limit:   req.Limit,
		Page:    req.Page,
		Filters: req.Filters,
		Guests:  req.Guests,
		Price: dtos.PriceRange{
			From: float64(req.Price.From),
			To:   float64(req.Price.To),
		},
		IncludingIds: req.IncludingIds,
		City:         req.City,
	}

	results, err := accomodationController.AccomodationService.SearchAndFilter(params)

	if err != nil {
		return nil, status.Error(codes.Aborted, err.Message)
	}

	return MapSearchAndFilterResponse(results.Data, results.TotalCount), nil
}

func (accommodationController *AccomodationController) PopulateRecommended(ctx Context, req *PopulateRecommendedRequest) (*PopulateRecommendedResponse, error) {

	idsMapped := make([]primitive.ObjectID, 0)

	for _, id := range req.Ids {
		idsMapped = append(idsMapped, shared.StringToObjectId(id))
	}

	result, err := accommodationController.AccomodationService.PopulateRecommended(idsMapped)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return MapRecommendedPopulatedResponse(result), nil
}
