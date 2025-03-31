package dto

import (
	"encoding/json"
	"github.com/CracherX/comments_hist/internal/entity"
	"net/http"
)

type e struct {
	Status      int    `json:"status"`
	Translation string `json:"translation"`
	Message     string `json:"message"`
	Details     string `json:"details,omitempty"`
}

// Response возвращает сообщение об успехе или ошибке клиенту в json формате.
func Response(w http.ResponseWriter, status int, msg string, details ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := e{
		Status:      status,
		Translation: http.StatusText(status),
		Message:     msg,
	}
	if len(details) > 0 {
		errorResponse.Details = details[0]
	}
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(errorResponse)
}

type GetCommentsResponse struct {
	Comments []entity.Comments `json:"comments"`
}

type GetMyCommentResponse struct {
	Comment *entity.Comments `json:"comments"`
}
