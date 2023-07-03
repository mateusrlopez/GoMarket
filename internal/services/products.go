package services

import (
	"github.com/jinzhu/copier"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/models"
	"github.com/mateusrlopez/go-market/internal/repositories"
)

type ProductsService interface {
	Create(input inputs.CreateProductInput) (models.Product, error)
	FindMany() ([]models.Product, error)
	FindOneByID(id string) (models.Product, error)
	UpdateOneByID(id string, input inputs.UpdateProductInput) (models.Product, error)
	DeleteOneByID(id string) error
}

type ProductsServiceImplementation struct {
	repository repositories.ProductsRepository
}

func NewProductsService(repository repositories.ProductsRepository) ProductsService {
	return ProductsServiceImplementation{
		repository: repository,
	}
}

func (s ProductsServiceImplementation) Create(input inputs.CreateProductInput) (models.Product, error) {
	var data models.Product

	if err := copier.Copy(&data, &input); err != nil {
		return models.Product{}, err
	}

	return s.repository.Create(data)
}

func (s ProductsServiceImplementation) FindMany() ([]models.Product, error) {
	return s.repository.FindMany()
}

func (s ProductsServiceImplementation) FindOneByID(id string) (models.Product, error) {
	return s.repository.FindOneByID(id)
}

func (s ProductsServiceImplementation) UpdateOneByID(id string, input inputs.UpdateProductInput) (models.Product, error) {
	var data models.Product

	if err := copier.Copy(&data, &input); err != nil {
		return models.Product{}, err
	}

	return s.repository.UpdateOneByID(id, data)
}

func (s ProductsServiceImplementation) DeleteOneByID(id string) error {
	return s.repository.DeleteOneByID(id)
}
