package entity

import "time"

type Comments struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Grade     int       `json:"grade"`
	Date      time.Time `json:"date"`
	ProductID int       `json:"productID"`
	UserID    int       `json:"userID"`
}
