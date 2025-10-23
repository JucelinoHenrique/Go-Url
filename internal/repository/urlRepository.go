package repository

import (
	"go-url-shortener/internal/model"

	"gorm.io/gorm"
)

type URLRepository interface {
	Save(url *model.URL) error
	FindByCode(shortCode string) (*model.URL, error)
	FindAll() ([]model.URL, error)
	IncrementClicks(code string) error
}

type URLRepositoryImpl struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepositoryImpl {
	return &URLRepositoryImpl{db: db}
}

func (r *URLRepositoryImpl) Save(url *model.URL) error {
	result := r.db.Create(url)
	return result.Error
}
func (r *URLRepositoryImpl) FindByCode(shortCode string) (*model.URL, error) {
	url := &model.URL{}
	result := r.db.Where("short_code = ?", shortCode).First(url)
	if result.Error != nil {
		return nil, result.Error
	}
	return url, nil
}

func (r *URLRepositoryImpl) IncrementClicks(code string) error {
	result := r.db.Model(&model.URL{}).Where("short_code = ?", code).UpdateColumn("clicks", gorm.Expr("clicks + ?", 1))
	return result.Error
}

func (r *URLRepositoryImpl) FindAll() ([]model.URL, error) {
	var urls []model.URL
	result := r.db.Find(&urls)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []model.URL{}, nil
		}
		return nil, result.Error
	}
	return urls, nil
}
