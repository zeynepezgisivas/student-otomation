package mhttp

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"net/http"
)

func ServerError(w http.ResponseWriter, err error, statusCode int) {
	var logMessage string
	http.Error(w, fmt.Sprintf("%v:%s", err, logMessage), statusCode)
}

func NotFound(w http.ResponseWriter) {
	UsageErrorWithStatus(w, "", nil, http.StatusNotFound)
}

func UsageError(w http.ResponseWriter, message string, data interface{}) {
	UsageErrorWithStatus(w, message, data, http.StatusPreconditionFailed)
}

func UsageErrorWithStatus(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	m, err := NewResponse(false, data)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
		return
	}
	err = m.SendWithStatus(w, statusCode)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
	}
}

func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	ar, err := NewResponse(true, data)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
		return
	}
	err = ar.Send(w)
}

func ResponseSuccessWithID(w http.ResponseWriter, id int64) {
	data := struct {
		ID int64
	}{
		ID: id,
	}

	ar, err := NewResponse(true, data)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
		return
	}

	err = ar.Send(w)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
	}
}

func ResponseSuccessWithKSUID(w http.ResponseWriter, id ksuid.KSUID) {
	data := struct {
		ID ksuid.KSUID
	}{
		ID: id,
	}

	ar, err := NewResponse(true, data)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
		return
	}

	err = ar.Send(w)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
	}
}

func ResponseFail(w http.ResponseWriter, msg string, statusCode int) {
	ar, err := NewResponse(false, msg)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
		if err != nil {

		}
		return
	}
	err = ar.SendWithStatus(w, statusCode)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
	}
}

func ResponseFailWithStruct(w http.ResponseWriter, data interface{}, statusCode int) {
	ar, err := NewResponse(false, data)
	if err != nil {
		ServerError(w, err, http.StatusInternalServerError)
		return
	}
	err = ar.SendWithStatus(w, statusCode)

}
