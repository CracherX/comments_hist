package handlers

import (
	"encoding/json"
	"errors"
	"github.com/CracherX/comments_hist/internal/controller/dto"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CommentsHandler struct {
	uc  CommentsUseCase
	val Validator
	log Logger
}

func NewCommentsHandler(uc CommentsUseCase, val Validator, log Logger) *CommentsHandler {
	return &CommentsHandler{uc: uc, val: val, log: log}
}

func (c *CommentsHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	pid, err := strconv.Atoi(q.Get("productID"))
	if err != nil {
		c.log.Debug("Bad Request", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Ознакомьтесь с документацией и повторите отправку запроса")
		return
	}

	data := dto.GetCommentsRequest{ProductID: pid}

	err = c.val.Validate(&data)
	if err != nil {
		c.log.Debug("Bad Request", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Ознакомьтесь с документацией и повторите отправку запроса")
		return
	}

	comments, err := c.uc.GetComments(pid)
	if err != nil {
		c.log.Error("Ошибка работы UC", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Внутренняя ошибка сервера!")
		return
	}

	res := dto.GetCommentsResponse{Comments: comments}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		c.log.Error("Проблема с энкодером", "Ошибка", err.Error())
	}
}

func (c *CommentsHandler) GetMyComment(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	pid, err := strconv.Atoi(q.Get("productID"))
	if err != nil {
		c.log.Debug("Bad Request", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Ознакомьтесь с документацией и повторите отправку запроса")
		return
	}

	data := dto.GetMyCommentRequest{ProductID: pid, JWT: q.Get("jwt")}

	err = c.val.Validate(&data)
	if err != nil {
		c.log.Debug("Bad Request", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Ознакомьтесь с документацией и повторите отправку запроса")
		return
	}

	comment, err := c.uc.GetMyComment(pid, data.JWT)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.log.Debug("Пользователь еще не оставил отзыв", "Ошибка", err.Error())
			dto.Response(w, http.StatusNoContent, "Пользователь еще не оставил отзыв")
		} else {
			c.log.Error("Ошибка в работе UC", "Ошибка", err.Error())
			dto.Response(w, http.StatusInternalServerError, "Внутренняя ошибка сервера!")
		}
		return
	}

	res := dto.GetMyCommentResponse{Comment: comment}

	w.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		c.log.Error("Проблема с энкодером", "Ошибка", err.Error())
	}
}
