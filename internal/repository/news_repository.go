package repository

import (
	"go-web-scaffold/internal/models"

	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(news *models.News) error
	GetByID(id uint) (*models.News, error)
	GetAll() ([]*models.News, error)
	Update(news *models.News) error
	Delete(id uint) error
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db: db}
}

func (r *newsRepository) Create(news *models.News) error {
	return r.db.Create(news).Error
}

func (r *newsRepository) GetByID(id uint) (*models.News, error) {
	var news models.News
	if err := r.db.First(&news, id).Error; err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *newsRepository) GetAll() ([]*models.News, error) {
	var newsList []*models.News
	if err := r.db.Find(&newsList).Error; err != nil {
		return nil, err
	}
	return newsList, nil
}

func (r *newsRepository) Update(news *models.News) error {
	return r.db.Save(news).Error
}

func (r *newsRepository) Delete(id uint) error {
	return r.db.Delete(&models.News{}, id).Error
}
