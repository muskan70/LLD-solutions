package response

type RootResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	isSet   bool
	// ExceptionCode
}

func Fail(status bool, msg string, err error) RootResponse {
	var errstr string
	if err != nil {
		errstr = err.Error()
	}
	return RootResponse{
		Status:  status,
		Error:   errstr,
		Message: msg,
		isSet:   true,
	}
}

func Success(msg string) RootResponse {
	return RootResponse{
		Status:  true,
		Message: msg,
		isSet:   true,
	}
}
