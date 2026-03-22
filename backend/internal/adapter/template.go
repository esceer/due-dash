package adapter

import (
	apimodel "github.com/esceer/due-dash/backend/internal/api/model"
	dbmodel "github.com/esceer/due-dash/backend/internal/repository/model"
)

func TemplateToDb(apiModel *apimodel.Template) *dbmodel.Template {
	if apiModel == nil {
		return nil
	}
	dbModel := NewTemplateToDb(&apiModel.NewTemplate)
	dbModel.Id = apiModel.Id
	dbModel.Enabled = apiModel.Enabled
	return dbModel
}

func NewTemplateToDb(apiModel *apimodel.NewTemplate) *dbmodel.Template {
	if apiModel == nil {
		return nil
	}
	return &dbmodel.Template{
		Name:       apiModel.Name,
		Frequency:  apiModel.Frequency,
		DayOfMonth: apiModel.DayOfMonth,
	}
}

func TemplateToApi(dbModel *dbmodel.Template) *apimodel.Template {
	if dbModel == nil {
		return nil
	}
	return &apimodel.Template{
		Id:      dbModel.Id,
		Enabled: dbModel.Enabled,
		NewTemplate: apimodel.NewTemplate{
			Name:       dbModel.Name,
			Frequency:  dbModel.Frequency,
			DayOfMonth: dbModel.DayOfMonth,
		},
	}
}

func TemplateSliceToApi(dbModels []*dbmodel.Template) []*apimodel.Template {
	apiModel := make([]*apimodel.Template, len(dbModels))
	for i, dbm := range dbModels {
		apiModel[i] = TemplateToApi(dbm)
	}
	return apiModel
}
