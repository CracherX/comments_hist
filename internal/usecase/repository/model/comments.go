package model

import "time"

type Comments struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Text      string    `gorm:"type:text;not null" json:"text"`
	Grade     int       `gorm:"not null;" json:"grade"`
	Date      time.Time `gorm:"autoCreateTime" json:"date"`
	ProductID int       `json:"productID"`
	UserID    int       `json:"userID"`
}
