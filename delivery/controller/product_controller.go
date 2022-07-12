package controller

import (
	"fmt"
	"golang-mongodb/model"
	"golang-mongodb/usecase"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductController struct {
	router                      *gin.Engine
	productUseCase              usecase.ProductRegistrationUseCase
	productFindAllUseCase       usecase.ProductFindAllWithPaginationUseCase
	productUpdateUseCase        usecase.ProductUpdateUseCase
	productDeleteUseCase        usecase.ProductDeleteUseCase
	productGetByIdUseCase       usecase.ProductGetByIdUseCase
	productGetByCategoryUseCase usecase.ProductGetByCategoryUseCase
}

func (pc *ProductController) registerNewProduct(ctx *gin.Context) {
	var newProduct model.Product
	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = pc.productUseCase.Register(&newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    newProduct,
	})
}

func (pc *ProductController) findAllProductWithPagination(ctx *gin.Context) {
	page, err := strconv.ParseInt(ctx.Param("page"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	totalDoc, err := strconv.ParseInt(ctx.Param("totalDoc"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	products, err := pc.productFindAllUseCase.FindAll(page, totalDoc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    products,
	})
}

func (pc *ProductController) updateProduct(ctx *gin.Context) {
	var by bson.D

	id := ctx.Param("id")

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = bson.UnmarshalExtJSON([]byte(body), true, &by)
	fmt.Println(by, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = pc.productUpdateUseCase.Update(id, by)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product, err := pc.productGetByIdUseCase.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    product,
	})
}

func (pc *ProductController) deleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := pc.productGetByIdUseCase.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = pc.productDeleteUseCase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    product,
	})
}

func (pc *ProductController) getProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := pc.productGetByIdUseCase.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    product,
	})
}

func (pc *ProductController) getProductByCategory(ctx *gin.Context) {
	category := ctx.Param("category")

	products, err := pc.productGetByCategoryUseCase.GetByCategory(category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if products == nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    products,
	})
}

func NewProductController(router *gin.Engine, productUseCase usecase.ProductRegistrationUseCase,
	productFindAllUseCase usecase.ProductFindAllWithPaginationUseCase, productUpdateUseCase usecase.ProductUpdateUseCase,
	productDeleteUseCase usecase.ProductDeleteUseCase, productGetByIdUseCase usecase.ProductGetByIdUseCase,
	productGetByCategoryUseCase usecase.ProductGetByCategoryUseCase) *ProductController {

	controller := ProductController{
		router:                      router,
		productUseCase:              productUseCase,
		productFindAllUseCase:       productFindAllUseCase,
		productUpdateUseCase:        productUpdateUseCase,
		productDeleteUseCase:        productDeleteUseCase,
		productGetByIdUseCase:       productGetByIdUseCase,
		productGetByCategoryUseCase: productGetByCategoryUseCase,
	}
	router.POST("/product", controller.registerNewProduct)
	router.GET("/product/:page/:totalDoc", controller.findAllProductWithPagination)
	router.PATCH("/product/:id", controller.updateProduct)
	router.DELETE("/product/:id", controller.deleteProduct)
	router.GET("/product/id/:id", controller.getProductById)
	router.GET("/product/category/:category", controller.getProductByCategory)
	return &controller
}
