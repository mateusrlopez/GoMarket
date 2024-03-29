package services

import (
	"github.com/jinzhu/copier"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/models"
	"github.com/mateusrlopez/go-market/internal/repositories"
)

type OrdersService interface {
	Create(input inputs.CreateOrderInput) (models.Order, error)
	FindMany(input inputs.QueryOrdersInput) ([]models.Order, error)
	FindOneByID(id string) (models.Order, error)
	UpdateOneByID(id string, input inputs.UpdateOrderInput) (models.Order, error)
	DeleteOneByID(id string) error
}

type ordersServiceImplementation struct {
	repository repositories.OrdersRepository
}

func NewOrdersService(repository repositories.OrdersRepository) OrdersService {
	return ordersServiceImplementation{
		repository: repository,
	}
}

func (s ordersServiceImplementation) Create(input inputs.CreateOrderInput) (models.Order, error) {
	var data models.Order

	if err := copier.Copy(&data, &input); err != nil {
		return models.Order{}, err
	}

	return s.repository.Create(data)
}

func (s ordersServiceImplementation) FindMany(input inputs.QueryOrdersInput) ([]models.Order, error) {
	var filter models.Order

	if err := copier.Copy(&filter, &input); err != nil {
		return nil, err
	}

	return s.repository.FindMany(filter)
}

func (s ordersServiceImplementation) FindOneByID(id string) (models.Order, error) {
	return s.repository.FindOneByID(id)
}

func (s ordersServiceImplementation) UpdateOneByID(id string, input inputs.UpdateOrderInput) (models.Order, error) {
	var data models.Order

	if err := copier.Copy(&data, &input); err != nil {
		return models.Order{}, err
	}

	return s.repository.UpdateOneByID(id, data)
}

func (s ordersServiceImplementation) DeleteOneByID(id string) error {
	return s.repository.DeleteOneByID(id)
}
