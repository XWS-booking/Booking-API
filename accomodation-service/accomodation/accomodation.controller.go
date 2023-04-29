package accomodation

import (
	"accomodation_service/accomodation/services/storage"
	. "accomodation_service/proto/accomodation"
	. "context"
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

	data := req.Picture

	err := accomodationController.StorageService.UploadImage(data)
	if err != nil {
		return nil, status.Error(codes.Aborted, "Data upload filed!")
	}

	return &CreateAccomodationResponse{
		Id: "123",
	}, nil
}

func (accomodationController *AccomodationController) FindAll(ctx Context, req *FindAllAccomodationRequest) (*FindAllAccomodationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	accomodations := accomodationController.AccomodationService.FindAll(req.GetCity(), req.GetGuests())
	var accomodationResponses []*AccomodationResponse
	for _, a := range accomodations {
		accomodationResponses = append(accomodationResponses, NewAccomodationResponse(a))
	}
	return &FindAllAccomodationResponse{AccomodationResponses: accomodationResponses}, nil
}
