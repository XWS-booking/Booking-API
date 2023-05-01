package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/infrastructure/services"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"io"
	"net/http"
	"strconv"
)

type CreateAccomodationHandler struct {
	accommodationClientAddress string
	authClientAddress          string
}

type AccomodationDto struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Street         string `json:"street"`
	StreetNumber   string `json:"streetNumber"`
	City           string `json:"city"`
	ZipCode        string `json:"zipCode"`
	Country        string `json:"country"`
	Wifi           bool   `json:"wifi"`
	Kitchen        bool   `json:"kitchen"`
	AirConditioner bool   `json:"airConditioner"`
	FreeParking    bool   `json:"freeParking"`
	MinGuests      int32  `json:"minGuests"`
	MaxGuests      int32  `json:"maxGuests"`
	OwnerId        string `json:"ownerId"`
}

func NewCreateAccomodationHandler(accommodationClientAddress, authClientAddress string) Handler {
	return &CreateAccomodationHandler{
		accommodationClientAddress: accommodationClientAddress,
		authClientAddress:          authClientAddress,
	}
}

func (handler *CreateAccomodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/accomodation/create", handler.Create)
	if err != nil {
		panic(err)
	}
}

func (handler *CreateAccomodationHandler) Create(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	token := r.Header["Authorization"][0]

	authClient := services.NewAuthClient(handler.authClientAddress)
	user, e := authClient.GetUser(context.TODO(), &gateway.GetUserRequest{Token: token})
	if e != nil {
		panic(e)
	}
	accomodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	var dto AccomodationDto

	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	files := MapFilesFromRequest(r, "attachment")
	dto.OwnerId = user.Id
	wifi, _ := strconv.ParseBool(r.FormValue("wifi"))
	kitchen, _ := strconv.ParseBool(r.FormValue("wifi"))
	airConditioner, _ := strconv.ParseBool(r.FormValue("wifi"))
	freeParking, _ := strconv.ParseBool(r.FormValue("wifi"))
	minGuests, _ := strconv.Atoi(r.FormValue("minGuests"))
	maxGuests, _ := strconv.Atoi(r.FormValue("maxGuests"))

	accReq := &gateway.CreateAccomodationRequest{
		Name:           r.FormValue("name"),
		Street:         r.FormValue("street"),
		StreetNumber:   r.FormValue("streetNumber"),
		City:           r.FormValue("city"),
		ZipCode:        r.FormValue("street"),
		Country:        r.FormValue("street"),
		Wifi:           wifi,
		Kitchen:        kitchen,
		AirConditioner: airConditioner,
		FreeParking:    freeParking,
		MinGuests:      int32(minGuests),
		MaxGuests:      int32(maxGuests),
		OwnerId:        user.Id,
		Pictures:       files,
	}
	res, e := accomodationClient.Create(context.TODO(), accReq)

	if e != nil {
		panic(e)
	}
	shared.Ok(&w, res)
}

func DecodeBody(req *http.Request, v interface{}) error {
	err := json.NewDecoder(req.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func MapFilesFromRequest(r *http.Request, fieldName string) []*gateway.ImageInfo {
	files := make([]*gateway.ImageInfo, 0)
	r.ParseMultipartForm(10 << 20)
	for _, fh := range r.MultipartForm.File[fieldName] {
		f, err := fh.Open()
		if err != nil {
			continue
		}
		file, err := io.ReadAll(f)
		if err != nil {
			return []*gateway.ImageInfo{}
		}

		info := &gateway.ImageInfo{
			Data:     file,
			Filename: fh.Filename,
		}
		f.Close()
		files = append(files, info)
	}
	return files
}
