package service

import (
	"go-web-scaffold/internal/models"
	"go-web-scaffold/internal/repository"
)

type NewsService interface {
	CreateNews(news *models.News) (*models.News, error)
	GetNewsByID(id uint) (*models.News, error)
	GetAllNews() ([]*models.News, error)
	UpdateNews(news *models.News) (*models.News, error)
	DeleteNews(id uint) error
}

type newsService struct {
	newsRepo repository.NewsRepository
}

func NewNewsService(newsRepo repository.NewsRepository) NewsService {
	return &newsService{newsRepo: newsRepo}
}

func (s *newsService) CreateNews(news *models.News) (*models.News, error) {
	if err := s.newsRepo.Create(news); err != nil {
		return nil, err
	}
	return news, nil
}

func (s *newsService) GetNewsByID(id uint) (*models.News, error) {
	return s.newsRepo.GetByID(id)
}

func (s *newsService) GetAllNews() ([]*models.News, error) {
	return s.newsRepo.GetAll()
}

func (s *newsService) UpdateNews(news *models.News) (*models.News, error) {
	if err := s.newsRepo.Update(news); err != nil {
		return nil, err
	}
	return news, nil
}

func (s *newsService) DeleteNews(id uint) error {
	return s.newsRepo.Delete(id)
}
