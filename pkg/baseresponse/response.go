package baseresponse

import (
	"encoding/json"
	"net/http"
)

func (base *BaseResponse) Succeed(w http.ResponseWriter, data interface{}) error {
	w.WriteHeader(http.StatusOK)
	base.Code = http.StatusOK
	base.Data = data
	return json.NewEncoder(w).Encode(*base)
}

func (base *BaseResponse) Fail(w http.ResponseWriter, data interface{}) error {
	w.WriteHeader(http.StatusInternalServerError)
	base.Code = http.StatusInternalServerError
	base.Data = data
	return json.NewEncoder(w).Encode(*base)
}

func (base *BaseResponse) FailWithCode(w http.ResponseWriter, data interface{}, code int) error {
	w.WriteHeader(code)
	base.Code = code
	base.Data = data
	return json.NewEncoder(w).Encode(*base)
}
