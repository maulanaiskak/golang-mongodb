package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"
)

type ProductGetByCategoryUseCase interface {
	GetByCategory(category string) ([]model.Product, error)
}

type productGetByCategoryUseCase struct {
	repo repository.ProductRepository
}

func (p *productGetByCategoryUseCase) GetByCategory(category string) ([]model.Product, error){
	return p.repo.GetByCategory(category)
}

func NewProductGetByCategoryUseCase(repo repository.ProductRepository) ProductGetByCategoryUseCase {
	return &productGetByCategoryUseCase{repo: repo}
}
