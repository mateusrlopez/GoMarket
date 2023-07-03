package services

import (
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(input inputs.CreateUserInput) (models.User, error)
	ValidateLogin(input inputs.LoginInput) (models.User, error)
	AssignToken(userId string) (string, error)
	ValidateToken(token string) (models.User, error)
	Logout(userId string) error
}

type authServiceImplementation struct {
	tokenService TokenService
	usersService UsersService
}

func NewAuthService(tokenService TokenService, usersService UsersService) AuthService {
	return authServiceImplementation{
		tokenService: tokenService,
		usersService: usersService,
	}
}

func (s authServiceImplementation) Register(input inputs.CreateUserInput) (models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	input.Password = string(hash)

	return s.usersService.Create(input)
}

func (s authServiceImplementation) ValidateLogin(input inputs.LoginInput) (models.User, error) {
	user, err := s.usersService.FindOneByEmail(input.Email)

	if err != nil {
		return models.User{}, err
	}

	if err = user.ComparePassword(input.Password); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s authServiceImplementation) AssignToken(userId string) (string, error) {
	return s.tokenService.AssignToken(userId)
}

func (s authServiceImplementation) ValidateToken(token string) (models.User, error) {
	id, err := s.tokenService.ValidateToken(token)

	if err != nil {
		return models.User{}, err
	}

	user, err := s.usersService.FindOneByID(id)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s authServiceImplementation) Logout(userId string) error {
	return s.tokenService.UnassignToken(userId)
}
