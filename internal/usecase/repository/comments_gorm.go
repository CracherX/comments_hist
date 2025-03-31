package repository

import (
	"github.com/CracherX/comments_hist/internal/entity"
	"gorm.io/gorm"
)

type CommentsRepo struct {
	db *gorm.DB
}

func NewCommentsGorm(db *gorm.DB) *CommentsRepo {
	return &CommentsRepo{db: db}
}

func (c *CommentsRepo) GetComments(prodID int) ([]entity.Comments, error) {
	var comments []entity.Comments

	if err := c.db.Find(&comments).Where("ProductID = ?", prodID).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *CommentsRepo) WriteComment(prodID, userID int, text string) error {
	var comment entity.Comments

	comment.Text = text
	comment.ProductID = prodID
	comment.UserID = userID

	if err := c.db.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (c *CommentsRepo) DeleteComment(id int) error {

	comment, err := c.GetComment(id)
	if err != nil {
		return err
	}

	if err = c.db.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}

// Ошибки в GetMyComment следует обрабатывать сначала на gorm.ErrRecordNorFound
func (c *CommentsRepo) GetMyComment(prodID, userID int) (*entity.Comments, error) {
	var comment entity.Comments

	if err := c.db.First(&comment).Where("ProductID = ? AND UserID = ?", prodID, userID).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *CommentsRepo) GetComment(id int) (*entity.Comments, error) {
	var comment entity.Comments

	if err := c.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}
