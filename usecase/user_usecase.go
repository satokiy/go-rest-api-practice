package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

/**
* Implment Usecase Layer
* Usecase called by Controller
 */

// ユースケースの定義
type IUserUsecase interface {
	// ここでは、ポインタを受け取る必要がない
	// なぜなら、ポインタを受け取ると、ポインタの中身を変更してしまう可能性があるため
	// ここでモデルが引数になっているのも、DIなのかもしれない
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error) // loginでは認証用のJWTを返却する
}

// usecaseはrepositoryを呼び出す
// そのためのIFを構造体に持つ
type userUsecase struct {
	// RepositoryのIFへの依存。ただし実態に依存させないのがポイント
	// これによってDIPを実現している
	// Repositoryの実態がなくても、インターフェースさえ満たせばOKという状態
	ur repository.IUserRepository
}

// dependency injectionによるコンストラクタ
// DIにおいては、依存する具体オブジェクトを外部から引数として渡す（注入）
// これによって外部から注入された抽象的なオブジェクトとして扱うことができる
// 引数はあくまでインターフェース。インターフェースにのみ依存させる
// 戻り値をインターフェースとすることで、userUsecaseは実質インターフェースを強制される？
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		ur: ur,
	}
}

// signupユースケースの具体
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	// インスタンス生成
	newUser := model.User{
		Email:    user.Email,
		Password: string(hash),
	}
	// 永続化
	// ポインタを指定して渡すので、newUserの値は置き換わる。だからそれをresponseにする
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	// response
	res := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return res, nil
}

// loginユースケースの具体
func (uu *userUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	// check email
	if err := uu.ur.GetUserbyEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
