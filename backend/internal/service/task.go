package service

import (
	"time"

	"github.com/esceer/due-dash/backend/internal/adapter"
	apimodel "github.com/esceer/due-dash/backend/internal/api/model"
	"github.com/esceer/due-dash/backend/internal/repository"
	"github.com/esceer/due-dash/backend/internal/repository/model"
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
	taskRepository     repository.TaskRepository
	templateRepository repository.TemplateRepository
}

func NewTaskService(taskRepo repository.TaskRepository, templateRepo repository.TemplateRepository) TaskService {
	return &taskService{
		taskRepository:     taskRepo,
		templateRepository: templateRepo,
	}
}

func (s *taskService) GetAll() ([]*apimodel.Task, error) {
	tasks, err := s.taskRepository.GetAll()
	return adapter.TaskSliceToApi(tasks), err
}

func (s *taskService) GetById(id int) (*apimodel.Task, error) {
	task, err := s.taskRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return adapter.TaskToApi(task), err
}

func (s *taskService) Create(apiTask *apimodel.NewTask) error {
	task := adapter.NewTaskToDb(apiTask)
	return s.taskRepository.Create(task)
}

func (s *taskService) GenerateFromTemplate() error {
	templates, err := s.templateRepository.GetAllActive()
	if err != nil {
		return err
	}

	targetMonths := getTargetMonths()

	for _, template := range templates {
		for _, month := range targetMonths {
			dueDate := s.adjustToMonthEnd(month.Year(), month.Month(), template.DayOfMonth)

			exists, err := s.taskRepository.ExistsByTemplateAndDueDate(template.Id, dueDate)
			if err != nil {
				continue
			}

			if !exists {
				newGeneratedTask := &model.Task{
					Template: template,
					Name:     template.Name,
					DueDate:  dueDate,
					Status:   string(apimodel.StatusPending),
				}
				if err := s.taskRepository.Create(newGeneratedTask); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *taskService) UpdateStatus(id int, newStatus string) error {
	task, err := s.taskRepository.GetById(id)
	if err != nil {
		return err
	}

	task.Status = newStatus
	return s.taskRepository.Update(task)
}

func (s *taskService) Delete(id int) error {
	return s.taskRepository.Delete(id)
}

func getTargetMonths() []time.Time {
	// Current and next month
	now := time.Now()
	return []time.Time{
		now,
		now.AddDate(0, 1, 0),
	}
}

func (s *taskService) adjustToMonthEnd(year int, month time.Month, day int) time.Time {
	// The 0 index of a month is the previous month's last day
	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()

	actualDay := day
	if actualDay > lastDayOfMonth {
		// E.g. if month has only 28 or 30 days
		actualDay = lastDayOfMonth
	}

	return time.Date(year, month, actualDay, 0, 0, 0, 0, time.Local)
}
