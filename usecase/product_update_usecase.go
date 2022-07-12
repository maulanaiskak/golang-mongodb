package usecase

import (
	"golang-mongodb/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductUpdateUseCase interface {
	Update(id string, by bson.D) error
}

type productUpdateUseCase struct {
	repo repository.ProductRepository
}

func (p *productUpdateUseCase) Update(id string, by bson.D) error {
	return p.repo.Update(id, by)
}

func NewProductUpdateUseCase(repo repository.ProductRepository) ProductUpdateUseCase {
	return &productUpdateUseCase{repo: repo}
}
