package entities

type HTTPResponse struct {
	ApiVersion int         `json:"apiVersion"`
	Context    string      `json:"context,omitempty"`
	ID         string      `json:"id,omitempty"`
	Method     string      `json:"method"`
	Params     interface{} `json:"params,omitempty"`
	Data       interface{} `json:"data"`
	Error      *HTTPError  `json:"error"`
}

type HTTPError struct {
	Message string             `json:"message"`
	Code    int                `json:"code"`
	Errors  []HTTPErrorDetails `json:"errors"`
}

type HTTPErrorDetails struct {
	Domain       string `json:"domain,omitempty"`
	Reason       string `json:"reason"`
	Message      string `json:"message"`
	Location     string `json:"location"`
	LocationType string `json:"locationType"`
	ExtendedHelp string `json:"extendedHelp,omitempty"`
	SendReport   bool   `json:"sendReport,omitempty"`
}

func WrapErrors(message string, code int, errors ...HTTPErrorDetails) *HTTPError {
	return &HTTPError{
		Message: message,
		Code:    code,
		Errors:  errors,
	}
}

func NewHTTPResponse(data interface{}, error *HTTPError, apiVersion int, method string) *HTTPResponse {
	return &HTTPResponse{
		ApiVersion: apiVersion,
		Method:     method,
		Data:       data,
		Error:      error,
	}
}
