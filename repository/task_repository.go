package repository

import (
	"go-rest-api/model"
	"gorm.io/gorm"
)

// interface definition
type ITaskRepository interface {
	CreateTask(task *model.Task) error
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

// implement
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}