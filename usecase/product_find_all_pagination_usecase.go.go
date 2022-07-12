package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"
)

type ProductFindAllWithPaginationUseCase interface {
	FindAll(page, totalDoc int64) ([]model.Product, error)
}

type productFindAllWithPaginationUseCase struct {
	repo repository.ProductRepository
}

func (p *productFindAllWithPaginationUseCase) FindAll(page, totalDoc int64) ([]model.Product, error){
	return p.repo.FindAllProductWithPagination(page, totalDoc)
}

func NewProductFindAllWithPaginationUseCase(repo repository.ProductRepository) ProductFindAllWithPaginationUseCase {
	return &productFindAllWithPaginationUseCase{repo: repo}
}
