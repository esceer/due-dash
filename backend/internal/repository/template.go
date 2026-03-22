package repository

import (
	"github.com/esceer/due-dash/backend/internal"
	"github.com/esceer/due-dash/backend/internal/repository/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TemplateRepository interface {
	Create(*model.Template) error
	GetAll() ([]*model.Template, error)
	GetById(int) (*model.Template, error)
	Update(*model.Template) error
	Delete(int) error
}

type templateRepository struct {
	db *gorm.DB
}

func NewTemplateRepository(db *gorm.DB) TemplateRepository {
	return &templateRepository{db: db}
}

func (r *templateRepository) Create(b *model.Template) error {
	return r.db.Create(b).Error
}

func (r *templateRepository) GetAll() ([]*model.Template, error) {
	var templates []*model.Template
	result := r.db.Find(&templates)
	return templates, result.Error
}

func (r *templateRepository) GetById(id int) (*model.Template, error) {
	var template model.Template
	result := r.db.Preload(clause.Associations).First(&template, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, internal.ErrNotFound
	}
	return &template, result.Error
}

func (r *templateRepository) Update(b *model.Template) error {
	if _, err := r.GetById(b.Id); err != nil {
		return err
	}
	return r.db.Updates(b).Error
}

func (r *templateRepository) Delete(id int) error {
	return r.db.Delete(&model.Template{}, id).Error
}
