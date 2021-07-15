package repository

import "rest-api/models"

type UserRepo interface {
	Query() ([]models.User, error)
	Save(u models.User) error
	Update(u models.User) error
	Delete(u models.User) error
}