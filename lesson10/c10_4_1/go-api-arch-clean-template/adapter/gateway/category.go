package gateway

import (
	"go-api-arch-clean-template/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetOrCreate(category *entity.Category) (*entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// GetOrCreate implements CategoryRepository.
func (c *categoryRepository) GetOrCreate(category *entity.Category) (*entity.Category, error) {
	var getOrCreatedCategory entity.Category
	tx := c.db.FirstOrCreate(&getOrCreatedCategory, category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &getOrCreatedCategory, nil
}
