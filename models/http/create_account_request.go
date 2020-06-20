package http

type CreateAccountRequest struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Gender        bool   `json:"gender"`
	PhoneNumber   string `json:"phoneNumber"`
	ReferalCardID string `json:"referalCardID"`
}
