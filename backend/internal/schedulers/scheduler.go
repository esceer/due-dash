package scheduler

import (
	"time"

	"github.com/esceer/due-dash/backend/internal/service"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

type scheduler struct {
	scheduler *gocron.Scheduler

	taskService service.TaskService
}

func ScheduleTasksFromTemplates(ts service.TaskService) error {
	s := gocron.NewScheduler(time.UTC)
	scheduler := &scheduler{scheduler: s, taskService: ts}
	job, err := s.Every(12).Hour().StartImmediately().Do(scheduler.generateTasksFromTemplates)
	job.SingletonMode() // prevent overlaps
	log.Info().Msg("Starting scheduler(s)")
	s.StartAsync()
	return err
}

func (s *scheduler) generateTasksFromTemplates() {
	if err := s.taskService.GenerateFromTemplate(); err != nil {
		log.Error().Err(err).Msg("error while generating tasks from templates")
	}
}
