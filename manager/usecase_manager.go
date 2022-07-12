package manager

import "golang-mongodb/usecase"

type UseCaseManager interface {
	ProductRegistrationUseCase() usecase.ProductRegistrationUseCase
	ProductFindAllWithPaginationUseCase() usecase.ProductFindAllWithPaginationUseCase
	ProductUpdateUseCase() usecase.ProductUpdateUseCase
	ProductDeleteUseCase() usecase.ProductDeleteUseCase
	ProductGetByIdUseCase() usecase.ProductGetByIdUseCase
	ProductGetByCategoryUseCase() usecase.ProductGetByCategoryUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) ProductRegistrationUseCase() usecase.ProductRegistrationUseCase {
	return usecase.NewProductRegistrationUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) ProductFindAllWithPaginationUseCase() usecase.ProductFindAllWithPaginationUseCase {
	return usecase.NewProductFindAllWithPaginationUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) ProductUpdateUseCase() usecase.ProductUpdateUseCase {
	return usecase.NewProductUpdateUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) ProductDeleteUseCase() usecase.ProductDeleteUseCase {
	return usecase.NewProductDeleteUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) ProductGetByIdUseCase() usecase.ProductGetByIdUseCase {
	return usecase.NewProductGetByIdUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) ProductGetByCategoryUseCase() usecase.ProductGetByCategoryUseCase {
	return usecase.NewProductGetByCategoryUseCase(u.repoManager.ProductRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
