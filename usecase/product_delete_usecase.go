package usecase

import (
	"golang-mongodb/repository"
)

type ProductDeleteUseCase interface {
	Delete(id string) error
}

type productDeleteUseCase struct {
	repo repository.ProductRepository
}

func (p *productDeleteUseCase) Delete(id string) error {
	return p.repo.Delete(id)
}

func NewProductDeleteUseCase(repo repository.ProductRepository) ProductDeleteUseCase {
	return &productDeleteUseCase{repo: repo}
}
