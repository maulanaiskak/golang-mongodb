package delivery

import (
	"fmt"
	"golang-mongodb/config"
	"golang-mongodb/delivery/controller"
	"golang-mongodb/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.useCaseManager.ProductRegistrationUseCase(),
		a.useCaseManager.ProductFindAllWithPaginationUseCase(), a.useCaseManager.ProductUpdateUseCase(),
		a.useCaseManager.ProductDeleteUseCase(), a.useCaseManager.ProductGetByIdUseCase(),
		a.useCaseManager.ProductGetByCategoryUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func NewServer() *appServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
	}
}
