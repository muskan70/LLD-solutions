package response

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func Fail(message string, err error) Response {
	var errstr string
	if err != nil {
		errstr = err.Error()
	}
	return Response{
		Error:   errstr,
		Message: message,
		Status:  false,
	}
}

func Success(message string) Response {
	return Response{
		Status:  true,
		Message: message,
	}
}

type RegisterUserResponse struct {
	Response
	UserId int `json:"userId"`
}

func SuccessWithUserId(message string, id int) RegisterUserResponse {
	return RegisterUserResponse{
		Response: Success(message),
		UserId:   id,
	}
}

type GetTopScoresResponse struct {
	Response
	TopScores interface{} `json:"topScores"`
}

func SuccessWithTopScores(message string, scores interface{}) GetTopScoresResponse {
	return GetTopScoresResponse{
		Response:  Success(message),
		TopScores: scores,
	}
}
