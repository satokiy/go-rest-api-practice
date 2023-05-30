package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type taskUsecase struct {
	tr repository.ITaskRepository
}

// usecase definition
type ITaskUsecase interface {
	// modelを扱うのは、ユースケースの役割ぽいので、ここでレスポンスを明示する
	CreateTask(task model.Task) (model.TaskResponse, error)
}

// constractor
// 値を渡すのか、参照を渡すのか？ -> 参照を渡したい。つまりコピーではなく実態を渡したい。
func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{
		tr: tr,
	}
}

// implement
func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	// CreateTaskを実行するとtaskの値が変更されるので、参照を渡す
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	res := model.TaskResponse{
		ID: 		task.ID,
		Title: 		task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return res, nil
}
