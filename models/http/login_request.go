package http

type CheckAccountExistsRequest struct {
	PhoneNumber   string `json:"phoneNumber"`
	ReferalCardID string `json:"referalCardID"`
}
