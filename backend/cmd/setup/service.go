package setup

import (
	"github.com/esceer/due-dash/backend/cmd/config"
	"github.com/esceer/due-dash/backend/internal/repository"
	"github.com/esceer/due-dash/backend/internal/service"
	"gorm.io/gorm"
)

type serviceBundle struct {
	taskService     service.TaskService
	templateService service.TemplateService
}

func Services(cfg *config.Config, db *gorm.DB) *serviceBundle {
	return &serviceBundle{
		taskService:     TaskService(db),
		templateService: TemplateService(db),
	}
}

func TaskService(db *gorm.DB) service.TaskService {
	taskRepo := repository.NewTaskRepository(db)
	templateRepo := repository.NewTemplateRepository(db)
	return service.NewTaskService(taskRepo, templateRepo)
}

func TemplateService(db *gorm.DB) service.TemplateService {
	repo := repository.NewTemplateRepository(db)
	return service.NewTemplateService(repo)
}
