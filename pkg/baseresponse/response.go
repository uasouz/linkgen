package baseresponse

import (
	"encoding/json"
	"net/http"
)

// Succeed - Function to create a successful response with code 200 and the given data, using BaseResponse
func (base *BaseResponse) Succeed(w http.ResponseWriter, data interface{}) error {
	w.WriteHeader(http.StatusOK)
	base.Code = http.StatusOK
	base.Data = data
	return json.NewEncoder(w).Encode(*base)
}

// Fail - Function to create a failed response with code 500 and the given data, using BaseResponse
func (base *BaseResponse) Fail(w http.ResponseWriter, data interface{}) error {
	w.WriteHeader(http.StatusInternalServerError)
	base.Code = http.StatusInternalServerError
	base.Data = data
	return json.NewEncoder(w).Encode(*base)
}

// FailWithCode - Function to create a failed response with the given code and the given data, using BaseResponse
func (base *BaseResponse) FailWithCode(w http.ResponseWriter, data interface{}, code int) error {
	w.WriteHeader(code)
	base.Code = code
	base.Data = data
	return json.NewEncoder(w).Encode(*base)
}
