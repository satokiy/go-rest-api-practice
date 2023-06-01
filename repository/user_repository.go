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
//	-> usecaseを見ればわかるが、usecaseではこのインターフェースに依存する実装をしている
//
// インターフェースの実態（？）はRepository側に所属する
// ここでは、値ではなく参照を渡す必要がある。なぜならデータの変更まで行うから
type IUserRepository interface {
	GetUserbyEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// repositoryはDBに依存…する？
// 本来はDBインターフェースに依存させるべきなのかもしれない
type userRepository struct {
	db *gorm.DB
}

// CreateUser implements IUserRepository
func (ur *userRepository) CreateUser(user *model.User) error {

	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	// userの値は変更されているので、エラー以外返す必要はない
	return nil
}

// GetUserbyEmail implements IUserRepository
func (ur *userRepository) GetUserbyEmail(user *model.User, email string) error {
	// DBアクセス
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// constructor
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}
