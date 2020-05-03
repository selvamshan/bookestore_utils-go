package rest_errors

import (
	"net/http"	
	"errors"
)

type RestErr struct {
	Message string  	 `json:"message"`
	Status int        	 `json:"status"`
	Error string    	 `json:"error"`
	Causes []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestError(msg string, status int, err string, causes []interface{}) *RestErr{
	return &RestErr{
		Message: msg, 
		Status: status, 
		Error: err, 
		Causes: causes,		
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
		
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not_found",
	}
}

func NewUnauthorizedError() *RestErr {
	return &RestErr{
		Message: "uanable to retrive user information from given access_token",
		Status: http.StatusUnauthorized,
		Error: "unauthorized",		
	}
}


func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal_server_error",		
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}