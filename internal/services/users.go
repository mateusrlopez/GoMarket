package services

import (
	"github.com/jinzhu/copier"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/models"
	"github.com/mateusrlopez/go-market/internal/repositories"
)

type UsersService interface {
	Create(input inputs.CreateUserInput) (models.User, error)
	FindMany() ([]models.User, error)
	FindOneByID(id string) (models.User, error)
	FindOneByEmail(email string) (models.User, error)
	UpdateOneByID(id string, input inputs.UpdateUserInput) (models.User, error)
	DeleteOneByID(id string) error
}

type usersServiceImplementation struct {
	repository repositories.UsersRepository
}

func NewUsersService(repository repositories.UsersRepository) UsersService {
	return usersServiceImplementation{
		repository: repository,
	}
}

func (s usersServiceImplementation) Create(input inputs.CreateUserInput) (models.User, error) {
	var data models.User

	if err := copier.Copy(&data, &input); err != nil {
		return models.User{}, err
	}

	return s.repository.Create(data)
}

func (s usersServiceImplementation) FindMany() ([]models.User, error) {
	return s.repository.FindMany()
}

func (s usersServiceImplementation) FindOneByID(id string) (models.User, error) {
	return s.repository.FindOneByID(id)
}

func (s usersServiceImplementation) FindOneByEmail(email string) (models.User, error) {
	return s.repository.FindOneByEmail(email)
}

func (s usersServiceImplementation) UpdateOneByID(id string, input inputs.UpdateUserInput) (models.User, error) {
	var data models.User

	if err := copier.Copy(&data, &input); err != nil {
		return models.User{}, err
	}

	return s.repository.UpdateOneByID(id, data)
}

func (s usersServiceImplementation) DeleteOneByID(id string) error {
	return s.repository.DeleteOneByID(id)
}
