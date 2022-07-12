package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"
)

type ProductGetByIdUseCase interface {
	GetById(id string) (model.Product, error)
}

type productGetByIdUseCase struct {
	repo repository.ProductRepository
}

func (p *productGetByIdUseCase) GetById(id string) (model.Product, error){
	return p.repo.GetById(id)
}

func NewProductGetByIdUseCase(repo repository.ProductRepository) ProductGetByIdUseCase {
	return &productGetByIdUseCase{repo: repo}
}
