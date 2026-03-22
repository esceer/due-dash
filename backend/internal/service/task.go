package service

import (
	"github.com/esceer/due-dash/backend/internal/adapter"
	apimodel "github.com/esceer/due-dash/backend/internal/api/model"
	"github.com/esceer/due-dash/backend/internal/repository"
)

type TaskService interface {
	GetAll() ([]*apimodel.Task, error)
	GetById(int) (*apimodel.Task, error)
	Create(*apimodel.NewTask) error
	GenerateFromTemplate() error
	UpdateStatus(int, string) error
	Delete(int) error
}

type taskService struct {
	repository repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{r}
}

func (s *taskService) GetAll() ([]*apimodel.Task, error) {
	tasks, err := s.repository.GetAll()
	return adapter.TaskSliceToApi(tasks), err
}

func (s *taskService) GetById(id int) (*apimodel.Task, error) {
	task, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return adapter.TaskToApi(task), err
}

func (s *taskService) Create(apiTask *apimodel.NewTask) error {
	task := adapter.NewTaskToDb(apiTask)
	return s.repository.Create(task)
}

func (s *taskService) GenerateFromTemplate() error {
	// dbTask := adapter.NewTaskToDb(apiTask)
	// return s.repository.Create(dbTask)
}

func (s *taskService) UpdateStatus(id int, newStatus string) error {
	task, err := s.repository.GetById(id)
	if err != nil {
		return err
	}

	task.Status = newStatus
	return s.repository.Update(task)
}

func (s *taskService) Delete(id int) error {
	return s.repository.Delete(id)
}
