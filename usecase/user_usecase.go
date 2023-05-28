package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

/**
* Implment Usecase Layer
* Usecase called by Controller
*/

// ユースケースの定義
type IUserUsecase interface {
	// ここでは、ポインタを受け取る必要がない
	// なぜなら、ポインタを受け取ると、ポインタの中身を変更してしまう可能性があるため
	SignUp(user model.User) (model.User, error)
	Login(user model.User) (string, error) // loginでは認証用のJWTを返却する
}

// usecaseはrepositoryを呼び出す
// そのためのIFを構造体に持つ
type userUsecase struct {
	ur repository.IUserRepository
}

// dependency injectionによるコンストラクタ
// DIにおいては、依存する具体オブジェクトを外部から引数として渡す（注入）
// これによって外部から注入された抽象的なオブジェクトとして扱うことができる
// 引数はあくまでインターフェース。インターフェースにのみ依存させる
// 戻り値をインターフェースとすることで、userUsecaseは実質インターフェースを矯正される？
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		ur: ur,
	}
}


func (uu *userUsecase) SignUp(user model.User) (model.User, error) {
	// TODO: implement.最初は仮の実装で良い
	// この後、リポジトリのCreateUserを呼び出す
	return user, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	// TODO: implement.今は仮の実装で良い
	// この後、リポジトリのCreateUserを呼び出す
	return "", nil
}