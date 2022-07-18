package dto

type ErrorResponse struct {
	Code     uint   `json:"code"`
	Error    string `json:"error"`
	Messages string `json:"messages"`
}
