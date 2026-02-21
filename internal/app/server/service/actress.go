package service

import (
	"github.com/samber/lo"
	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/repository"
	"github.com/sssjuing/media-vault/internal/app/server/types"

	"gorm.io/datatypes"
)

type ActressService interface {
	GetList() ([]types.ActressDTO, error)
	Create(*types.CreateActressDTO) (*types.ActressDTO, error)
	Update(*model.Actress, *types.UpdateActressDTO) (*types.ActressDTO, error)
}

type ActressServiceImpl struct {
	actressRepo repository.ActressRepository
}

func NewActressService(actressRepo repository.ActressRepository) *ActressServiceImpl {
	return &ActressServiceImpl{
		actressRepo: actressRepo,
	}
}

func (svc *ActressServiceImpl) GetList() ([]types.ActressDTO, error) {
	actresses, err := svc.actressRepo.FindAll()
	if err != nil {
		return nil, err
	}
	list := lo.Map(actresses, func(a model.Actress, _ int) types.ActressDTO {
		return *a.DTO()
	})
	return list, nil
}

func (svc *ActressServiceImpl) Create(cad *types.CreateActressDTO) (*types.ActressDTO, error) {
	a := model.Actress{
		UniqueName:   cad.UniqueName,
		ChineseName:  cad.ChineseName,
		EnglishName:  cad.EnglishName,
		OtherNames:   (*datatypes.JSON)(cad.OtherNames),
		BirthDate:    cad.BirthDate,
		BirthPlace:   cad.BirthPlace,
		Height:       cad.Height,
		Weight:       cad.Weight,
		Measurements: (*datatypes.JSON)(cad.Measurements),
		Cup:          cad.Cup,
		BloodGroup:   cad.BloodGroup,
		DebutDate:    cad.DebutDate,
		Hobbies:      cad.Hobbies,
		Notes:        cad.Notes,
	}
	if err := svc.actressRepo.Create(&a); err != nil {
		return nil, err
	}
	return a.DTO(), nil
}

func (svc *ActressServiceImpl) Update(a *model.Actress, uad *types.UpdateActressDTO) (*types.ActressDTO, error) {
	a.UniqueName = uad.UniqueName
	a.ChineseName = uad.ChineseName
	a.EnglishName = uad.EnglishName
	a.OtherNames = (*datatypes.JSON)(uad.OtherNames)
	a.BirthDate = uad.BirthDate
	a.BirthPlace = uad.BirthPlace
	a.Height = uad.Height
	a.Weight = uad.Weight
	a.Measurements = (*datatypes.JSON)(uad.Measurements)
	a.Cup = uad.Cup
	a.BloodGroup = uad.BloodGroup
	a.DebutDate = uad.DebutDate
	a.Hobbies = uad.Hobbies
	a.Notes = uad.Notes
	if err := svc.actressRepo.Update(a); err != nil {
		return nil, err
	}
	return a.DTO(), nil
}
