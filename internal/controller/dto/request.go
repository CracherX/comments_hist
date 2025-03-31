package dto

type GetCommentsRequest struct {
	ProductID int `validate:"required"`
}

type GetMyCommentRequest struct {
	ProductID int    `validate:"required"`
	JWT       string `validate:"required,jwt"`
}
