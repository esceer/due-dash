package service

import (
	"github.com/esceer/due-dash/backend/internal/adapter"
	apimodel "github.com/esceer/due-dash/backend/internal/api/model"
	"github.com/esceer/due-dash/backend/internal/repository"
)

type TemplateService interface {
	GetAll() ([]*apimodel.Template, error)
	GetById(int) (*apimodel.Template, error)
	Create(*apimodel.NewTemplate) error
	Update(*apimodel.Template) error
	Delete(int) error
}

type templateService struct {
	repository repository.TemplateRepository
}

func NewTemplateService(r repository.TemplateRepository) TemplateService {
	return &templateService{r}
}

func (s *templateService) GetAll() ([]*apimodel.Template, error) {
	templates, err := s.repository.GetAll()
	return adapter.TemplateSliceToApi(templates), err
}

func (s *templateService) GetById(id int) (*apimodel.Template, error) {
	template, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return adapter.TemplateToApi(template), err
}

func (s *templateService) Create(apiTemplate *apimodel.NewTemplate) error {
	template := adapter.NewTemplateToDb(apiTemplate)
	return s.repository.Create(template)
}

func (s *templateService) Update(apiInvoice *apimodel.Template) error {
	template := adapter.TemplateToDb(apiInvoice)
	return s.repository.Update(template)
}

func (s *templateService) Delete(id int) error {
	return s.repository.Delete(id)
}
