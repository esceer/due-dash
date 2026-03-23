package repository

import (
	"time"

	"github.com/esceer/due-dash/backend/internal"
	"github.com/esceer/due-dash/backend/internal/repository/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository interface {
	Create(*model.Task) error
	GetAll() ([]*model.Task, error)
	GetById(int) (*model.Task, error)
	Update(*model.Task) error
	Delete(int) error
	ExistsByTemplateAndDueDate(int, time.Time) (bool, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(b *model.Task) error {
	return r.db.Create(b).Error
}

func (r *taskRepository) GetAll() ([]*model.Task, error) {
	var tasks []*model.Task
	result := r.db.Preload(clause.Associations).Find(&tasks)
	return tasks, result.Error
}

func (r *taskRepository) GetById(id int) (*model.Task, error) {
	var task model.Task
	result := r.db.Preload(clause.Associations).First(&task, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, internal.ErrNotFound
	}
	return &task, result.Error
}

func (r *taskRepository) Update(b *model.Task) error {
	if _, err := r.GetById(b.Id); err != nil {
		return err
	}
	return r.db.Save(b).Error
}

func (r *taskRepository) Delete(id int) error {
	return r.db.Delete(&model.Task{}, id).Error
}

func (r *taskRepository) ExistsByTemplateAndDueDate(templateId int, dueDate time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&model.Task{}).
		Where("template_id = ? AND due_date = ?", templateId, dueDate).
		Count(&count).Error

	return count > 0, err
}
