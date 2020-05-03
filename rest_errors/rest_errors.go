package rest_errors

import (
	"net/http"	
	"errors"
)


type RestErr interface {
	Message() string 
	Status() int
	Error() string
	Causes() []interface {}	
}

type restErr struct {
	message string  	 `json:"message"`
	status  int        	 `json:"status"`
	err     string    	 `json:"error"`
	causes []interface{} `json:"causes"`
}



func (e *restErr) Error() string {
	return fmt.Sprintf("messages: %s - status: %d - err: %s - causes: [ %v ]",
			e.message, e.status, e.err, e.causes)
}

func (e *restErr) Message() string {
	return e.message
}

func (e *restErr) Status() int {
	return e.status
}
func (e *restErr) Causes() []interface{} {
	return e.causes
}


func NewRestError(msg string, status int, err string, causes []interface{}) RestErr{
	return &restErr{
		message: msg, 
		status: status, 
		err: err, 
		causes: causes,		
	}
}

func NewBadRequestError(message string) RestErr {
	return &restErr{
		message: message,
		status: http.StatusBadRequest,
		err: "bad_request",
		
	}
}

func NewNotFoundError(message string) RestErr {
	return &restErr{
		message: message,
		status: http.StatusNotFound,
		err: "not_found",
	}
}

func NewUnauthorizedError() RestErr {
	return &restErr{
		message: "uanable to retrive user information from given access_token",
		status: http.StatusUnauthorized,
		err: "unauthorized",		
	}
}


func NewInternalServerError(message string, err_  error) RestErr {
	result := &restErr{
		message: message,
		status: http.StatusInternalServerError,
		err: "internal_server_error",		
	}
	if err != nil {
		result.causes = append(result.causes, err_.Error())
	}
	return result
}