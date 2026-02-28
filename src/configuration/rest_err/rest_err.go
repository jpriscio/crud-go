package resterr

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequesValidatetErr(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: "",
		Err:     "Bad Request Validete",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not Found Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewForbiddendErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}
