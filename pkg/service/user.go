package service

import (
	"CV_MANAGER/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	LoginUser(email, password string) (models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) LoginUser(email, password string) (models.User, error) {
	var user models.User
	result := s.db.Where("correo = ?", email).First(&user)
	if result.Error != nil {
		return models.User{}, errors.New("usuario no encontrado")
	}

	// Comparamos la contraseña ingresada con la almacenada
	err := bcrypt.CompareHashAndPassword([]byte(user.Contraseña), []byte(password))
	if err != nil {
		return models.User{}, errors.New("contraseña incorrecta")
	}

	return user, nil
}
