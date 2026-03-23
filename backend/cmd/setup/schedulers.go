package setup

import (
	scheduler "github.com/esceer/due-dash/backend/internal/schedulers"
	"github.com/esceer/due-dash/backend/internal/service"
)

func Schedulers(s service.TaskService) error {
	return scheduler.ScheduleTasksFromTemplates(s)
}
