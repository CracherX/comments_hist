package handlers

import "github.com/CracherX/comments_hist/internal/entity"

type Validator interface {
	Validate(dto interface{}) error
}

type Logger interface {
	Info(msg string, field ...any)
	Error(msg string, field ...any)
	Debug(msg string, field ...any)
}

type CommentsUseCase interface {
	GetComments(prodID int) ([]entity.Comments, error)
	DeleteComment(id int, jwt string) error
	WriteComment(prodID int, jwt string, text string) error
	GetMyComment(prodID int, jwt string) (*entity.Comments, error)
}
