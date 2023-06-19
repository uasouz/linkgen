package baseresponse

// BaseResponse - This is the default struct for a HTTP response
type BaseResponse struct {
	Success bool        `json:"success,omitempty"`
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}
