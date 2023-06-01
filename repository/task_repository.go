package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// interface definition
// Repositoryではタスクの実態を引数として処理を行う。実態が変更されるので、戻り値はエラーのみ
// Read系は引数に空のモデルのポインタを渡し、そこに詰める
type ITaskRepository interface {
	CreateTask(task *model.Task) error
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}
type taskRepository struct {
	db *gorm.DB
}

// constructor
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{
		db,
	}
}

// DeleteTask implements ITaskRepository
func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist. Delete failed")
	}
	return nil
}

// GetAllTasks implements ITaskRepository
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").
		Where("user_id = ?", userId).
		Order("created_at desc").
		Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// GetTaskById implements ITaskRepository
func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil

}

// UpdateTask implements ITaskRepository
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.
		Model(task).
		Clauses(clause.Returning{}).
		Where("id=? AND user_id=?", taskId, userId).
		Update("title", task.Title)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist. Update failed")
	}
	return nil
}

// implement
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}
