package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

/**
* repositoryはデータベースへのアクセスを行う処理を実装する
 */

// これはユースケースの責務では？？
//	-> usecaseを見ればわかるが、usecaseではこのインターフェースに依存する実装をしている
// インターフェースの実態（？）はRepository側に所属する
type IUserRepository interface {
	GetUserbyEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

// CreateUser implements IUserRepository
func (ur *userRepository) CreateUser(user *model.User) error {
	// TODO: implement
	ur.db.Create(user)
	panic("unimplemented")
}

// GetUserbyEmail implements IUserRepository
func (ur *userRepository) GetUserbyEmail(user *model.User, email string) error {
	// TODO: implement
	panic("unimplemented")
}

// constructor
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}
