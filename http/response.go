package mhttp

import (
	"encoding/json"
	"errors"
)

import (
	"github.com/segmentio/ksuid"
	"net/http"
)

type GrpcError struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	ErrData interface{} `json:"err_data"`
}

func GenerateGrpcError(errCode int, errMsg string, data interface{}) error {
	var m = GrpcError{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		ErrData: data,
	}

	byte, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return errors.New(string(byte))
}

type Response struct {
	ID      string
	Success bool
	Data    interface{}
}

func NewResponse(success bool, data interface{}) (*Response, error) {

	m := &Response{
		ID:      ksuid.New().String(),
		Success: success,
		Data:    data,
	}

	return m, nil
}

func (m *Response) SendWithStatus(w http.ResponseWriter, statusCode int) error {
	encjson, err := json.Marshal(m)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	_, err = w.Write(encjson)
	return err
}

func (m *Response) Send(w http.ResponseWriter) error {
	return m.SendWithStatus(w, http.StatusOK)
}
