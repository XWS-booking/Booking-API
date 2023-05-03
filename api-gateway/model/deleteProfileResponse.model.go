package model

type DeleteProfileResponse struct {
	Message string `json:"message"`
	Deleted bool   `json:"deleted"`
}
