package rest_errors

import (
	"fmt"
	"errors"
	"net/http"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is a message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "this is a message", err.Message())
	assert.EqualValues(t, "message: this is a message - status: 400 - error: bad_request - causes: []", err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is a message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "this is a message", err.Message())
	assert.EqualValues(t, "message: this is a message - status: 404 - error: not_found - causes: []", err.Error())
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is a message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is a message", err.Message())
	assert.EqualValues(t, "message: this is a message - status: 500 - error: internal_server_error - causes: [database error]", err.Error())
	assert.NotNil(t, err.Causes())
	assert.Equal(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

func TestNewRestError(t *testing.T) {
	err := NewRestError("impleted me", http.StatusNotImplemented, "not_implemented", nil)
	assert.NotNil(t, err)
	assert.Nil(t, nil, err.Causes())
	assert.EqualValues(t, "impleted me", err.Message())
	assert.EqualValues(t, "message: impleted me - status: 501 - error: not_implemented - causes: []", err.Error())
}

func TestNewErrorFromBytes(t *testing.T){
	err := &restErr{
		ErrMessage:"no user found with given credintials",
		ErrStatus:404,
		ErrError:"not_found",
	}
	bytes, _ := json.Marshal(err)
	apiErr, jsonErr :=  NewRestErrorFromBytes(bytes)
    fmt.Println(apiErr)
	fmt.Println(jsonErr)
	assert.NotNil(t,apiErr)
	assert.Nil(t, jsonErr)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status())
	assert.EqualValues(t, "no user found with given credintials", apiErr.Message())
	assert.EqualValues(t, "message: no user found with given credintials - status: 404 - error: not_found - causes: []", apiErr.Error())
}
