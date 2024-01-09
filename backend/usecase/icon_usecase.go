package usecase

import (
	"github.com/kosuke9809/aws_icons_quiz_backend/model"
	"github.com/kosuke9809/aws_icons_quiz_backend/repository"
)

type IIconUsecase interface {
	GetAllIcons() ([]model.Icon, error)
	GetIconById(id uint) (model.Icon, error)
	GetIconByName(name string) (model.Icon, error)
	GetRandomIcons(size int) ([]model.Icon, error)
}

type iconUsecase struct {
	ir repository.IIconRepository
}

func NewIconUsecase(ir repository.IIconRepository) IIconUsecase {
	return &iconUsecase{ir}
}

func (iu *iconUsecase) GetIconById(id uint) (model.Icon, error) {
	icon := model.Icon{}
	if err := iu.ir.GetIconById(&icon, id); err != nil {
		return model.Icon{}, err
	}
	resIcon := model.Icon{
		ID:        icon.ID,
		Name:      icon.Name,
		Category:  icon.Category,
		ObjectKey: icon.ObjectKey,
	}
	return resIcon, nil
}

func (iu *iconUsecase) GetIconByName(name string) (model.Icon, error) {
	icon := model.Icon{}
	if err := iu.ir.GetIconByName(&icon, name); err != nil {
		return model.Icon{}, err
	}
	resIcon := model.Icon{
		ID:        icon.ID,
		Name:      icon.Name,
		Category:  icon.Category,
		ObjectKey: icon.ObjectKey,
	}
	return resIcon, nil
}

func (iu *iconUsecase) GetAllIcons() ([]model.Icon, error) {
	icons := []model.Icon{}
	if err := iu.ir.GetAllIcons(&icons); err != nil {
		return nil, err
	}

	resIcons := []model.Icon{}
	for _, i := range icons {
		icon := model.Icon{
			ID:        i.ID,
			Name:      i.Name,
			Category:  i.Category,
			ObjectKey: i.ObjectKey,
		}
		resIcons = append(resIcons, icon)
	}
	return resIcons, nil
}

func (iu *iconUsecase) GetRandomIcons(size int) ([]model.Icon, error) {
	icons := []model.Icon{}
	if err := iu.ir.GetRandomIcons(&icons, size); err != nil {
		return nil, err
	}
	resIcons := []model.Icon{}
	for _, i := range icons {
		icon := model.Icon{
			ID:        i.ID,
			Name:      i.Name,
			Category:  i.Category,
			ObjectKey: i.ObjectKey,
		}
		resIcons = append(resIcons, icon)
	}
	return resIcons, nil

}
