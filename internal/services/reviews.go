package services

import (
	"github.com/jinzhu/copier"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/models"
	"github.com/mateusrlopez/go-market/internal/repositories"
)

type ReviewsService interface {
	Create(input inputs.CreateReviewInput) (models.Review, error)
	FindMany(input inputs.QueryReviewsInput) ([]models.Review, error)
	FindOneByID(id string) (models.Review, error)
	UpdateOneByID(id string, input inputs.UpdateReviewInput) (models.Review, error)
	DeleteOneByID(id string) error
}

type reviewsServiceImplementation struct {
	repository repositories.ReviewsRepository
}

func NewReviewsService(repository repositories.ReviewsRepository) ReviewsService {
	return reviewsServiceImplementation{
		repository: repository,
	}
}

func (s reviewsServiceImplementation) Create(input inputs.CreateReviewInput) (models.Review, error) {
	var data models.Review

	if err := copier.Copy(&data, &input); err != nil {
		return models.Review{}, err
	}

	return s.repository.Create(data)
}

func (s reviewsServiceImplementation) FindMany(input inputs.QueryReviewsInput) ([]models.Review, error) {
	var filter models.Review

	if err := copier.Copy(&filter, &input); err != nil {
		return nil, err
	}

	return s.repository.FindMany(filter)
}

func (s reviewsServiceImplementation) FindOneByID(id string) (models.Review, error) {
	return s.repository.FindOneByID(id)
}

func (s reviewsServiceImplementation) UpdateOneByID(id string, input inputs.UpdateReviewInput) (models.Review, error) {
	var data models.Review

	if err := copier.Copy(&data, &input); err != nil {
		return models.Review{}, err
	}

	return s.repository.UpdateOneByID(id, data)
}

func (s reviewsServiceImplementation) DeleteOneByID(id string) error {
	return s.repository.DeleteOneByID(id)
}
