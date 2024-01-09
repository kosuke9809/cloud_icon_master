package repository

import (
	"github.com/kosuke9809/aws_icons_quiz_backend/model"
	"gorm.io/gorm"
)

type IIconRepository interface {
	GetIconByName(icon *model.Icon, name string) error
	GetIconById(icon *model.Icon, id uint) error
	GetAllIcons(icons *[]model.Icon) error
	GetRandomIcons(icons *[]model.Icon, size int) error
}

type iconRepository struct {
	db *gorm.DB
}

func NewIconRepository(db *gorm.DB) IIconRepository {
	return &iconRepository{db}
}

func (ir *iconRepository) GetIconByName(icon *model.Icon, name string) error {
	if err := ir.db.Where("name=?", name).First(icon).Error; err != nil {
		return err
	}
	return nil
}

func (ir *iconRepository) GetIconById(icon *model.Icon, id uint) error {
	if err := ir.db.Where("id=?", id).First(icon).Error; err != nil {
		return err
	}
	return nil
}

func (ir *iconRepository) GetAllIcons(icons *[]model.Icon) error {
	if err := ir.db.Order("id").Find(icons).Error; err != nil {
		return err
	}
	return nil
}

func (ir *iconRepository) GetRandomIcons(icons *[]model.Icon, size int) error {
	if err := ir.db.Order("RANDOM()").Limit(size).Find(icons).Error; err != nil {
		return err
	}
	return nil
}
