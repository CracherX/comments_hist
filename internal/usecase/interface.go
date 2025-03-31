package usecase

import (
	"github.com/CracherX/comments_hist/internal/entity"
	"net/http"
)

type CommentsRepo interface {
	WriteComment(prodID, userID int, text string) error
	DeleteComment(id int) error
	GetComments(prodID int) ([]entity.Comments, error)
	GetMyComment(prodID, userID int) (*entity.Comments, error)
	GetComment(id int) (*entity.Comments, error)
}

type Client interface {
	Get(path string, queryParams ...map[string]string) (*http.Response, error)
}
