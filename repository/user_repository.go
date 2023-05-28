package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

/**
* repositoryはデータベースへのアクセスを行う処理を実装する
 */

// これはユースケースの責務では？？
//
//	-> 境界線としては、ユースケース側にある.
type IUserRepository interface {
	GetUserbyEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

// CreateUser implements IUserRepository
func (*userRepository) CreateUser(user *model.User) error {
	// TODO: implement
	panic("unimplemented")
}

// GetUserbyEmail implements IUserRepository
func (*userRepository) GetUserbyEmail(user *model.User, email string) error {
	// TODO: implement
	panic("unimplemented")
}

// constructor
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}
