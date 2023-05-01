package accomodation

import (
	"accomodation_service/accomodation/dtos"
	"accomodation_service/accomodation/services/storage"
	. "accomodation_service/proto/accomodation"
	"accomodation_service/shared"
	. "context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"fmt"
	"github.com/google/uuid"

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
	fmt.Println(urls)

	accomodationDto := dtos.AccomodationDto{
		Name:           req.Name,
		Street:         req.Street,
		StreetNumber:   req.StreetNumber,
		City:           req.City,
		ZipCode:        req.ZipCode,
		Country:        req.Country,
		Wifi:           req.Wifi,
		Kitchen:        req.Kitchen,
		AirConditioner: req.AirConditioner,
		FreeParking:    req.FreeParking,
		MinGuests:      req.MinGuests,
		MaxGuests:      req.MaxGuests,
		OwnerId:        shared.StringToObjectId(req.OwnerId),
	}
	fmt.Println("evo me")
	fmt.Println(accomodationDto.Pictures)

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
