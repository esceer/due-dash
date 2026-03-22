package adapter

import (
	apimodel "github.com/esceer/due-dash/backend/internal/api/model"
	dbmodel "github.com/esceer/due-dash/backend/internal/repository/model"
)

func NewTaskToDb(apiModel *apimodel.NewTask) *dbmodel.Task {
	if apiModel == nil {
		return nil
	}
	return &dbmodel.Task{
		Name:    apiModel.Name,
		DueDate: apiModel.DueDate,
	}
}

func TaskToApi(dbModel *dbmodel.Task) *apimodel.Task {
	if dbModel == nil {
		return nil
	}
	task := &apimodel.Task{
		Id:          dbModel.Id,
		CreatedAt:   dbModel.CreatedAt,
		CompletedAt: dbModel.CompletedAt,
		Status:      apimodel.Status(dbModel.Status),
		NewTask: apimodel.NewTask{
			Name:    dbModel.Name,
			DueDate: dbModel.DueDate,
		},
	}
	if dbModel.Template != nil {
		task.Template = TemplateToApi(dbModel.Template)
	}
	return task
}

func TaskSliceToApi(dbModels []*dbmodel.Task) []*apimodel.Task {
	apiModel := make([]*apimodel.Task, len(dbModels))
	for i, dbm := range dbModels {
		apiModel[i] = TaskToApi(dbm)
	}
	return apiModel
}
