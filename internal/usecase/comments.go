package usecase

import (
	"encoding/json"
	"errors"
	"github.com/CracherX/comments_hist/internal/entity"
	"github.com/CracherX/comments_hist/internal/usecase/clientDTO"
)

type CommentsUseCase struct {
	repo   CommentsRepo
	client Client
}

func NewCommentsUseCase(repo CommentsRepo) *CommentsUseCase {
	return &CommentsUseCase{repo: repo}
}

func (c *CommentsUseCase) GetComments(prodID int) ([]entity.Comments, error) {
	return c.repo.GetComments(prodID)
}

func (c *CommentsUseCase) GetMyComment(prodID int, jwt string) (*entity.Comments, error) {
	user, err := c.auth(jwt)

	if err != nil {
		return nil, err
	}
	return c.repo.GetMyComment(prodID, user.ID)
}

func (c *CommentsUseCase) DeleteComment(id int, jwt string) error {
	user, err := c.auth(jwt)
	if err != nil {
		return err
	}
	comment, err := c.repo.GetComment(id)

	if !user.IsAdmin || user.ID != comment.UserID {
		return errors.New("Пользователь не админ и отзыв принадлежит не ему ")
	}
	err = c.repo.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentsUseCase) WriteComment(prodID int, jwt string, text string) error {
	user, err := c.auth(jwt)
	if err != nil {
		return err
	}
	err = c.repo.WriteComment(prodID, user.ID, text)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentsUseCase) auth(jwt string) (*clientDTO.ProfileResponse, error) {
	var user clientDTO.ProfileResponse
	params := map[string]string{
		"jwt": jwt,
	}

	res, err := c.client.Get("/auth/profile", params)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
