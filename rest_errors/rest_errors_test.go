package rest_errors

import (
	"net/http"
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)


func TestNewError(t *testing.T)  {
}

func TestNewBadRequestError(t *testing.T)  {
	err := NewBadRequestError("this is a message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T)  {
	err := NewNotFoundError("this is a message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}


func TestNewInternalServerError(t *testing.T)  {
	err := NewInternalServerError("this is a message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
	assert.NotNil(t, err.Causes)
	assert.Equal(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])
}