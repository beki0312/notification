package models

type ErrorResponse struct {
	ErrorCode     int         `json:"errorCode"`
	ErrorDesc     string      `json:"errorDescription"`
	AditionalInfo interface{} `json:"response,omitempty"`
}
