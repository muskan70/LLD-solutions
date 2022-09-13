package requests

type RegisterUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type PushScoreRequest struct {
	UserId int `json:"userId" binding:"gt=0"`
	Score  int `json:"score" binding:"gt=0"`
}
