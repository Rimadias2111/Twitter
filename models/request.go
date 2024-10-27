package models

import "github.com/google/uuid"

type RequestId struct {
	Id uuid.UUID `json:"id"`
}

type ResponseId struct {
	Id string `json:"id"`
}

type ResponseError struct {
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
}

type ResponseSuccess struct {
	Message string
}
