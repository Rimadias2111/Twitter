package models

type RequestId struct {
	Id string `json:"id"`
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
