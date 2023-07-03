package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	customerrors "github.com/mateusrlopez/go-market/internal/errors"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/services"
)

type ProductsController struct {
	service services.ProductsService
}

func NewProductsController(service services.ProductsService) ProductsController {
	return ProductsController{
		service: service,
	}
}

func (c ProductsController) Create(ctx *gin.Context) {
	var input inputs.CreateProductInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.service.Create(input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"product": product})
}

func (c ProductsController) Index(ctx *gin.Context) {
	products, err := c.service.FindMany()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

func (c ProductsController) GetOneById(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := c.service.FindOneByID(id)

	if err != nil {
		if errors.Is(err, customerrors.ErrProductNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"product": product})
}

func (c ProductsController) UpdateOneById(ctx *gin.Context) {
	var input inputs.UpdateProductInput

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.service.UpdateOneByID(id, input)

	if err != nil {
		if errors.Is(err, customerrors.ErrProductNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Product": updated})
}

func (c ProductsController) RemoveOneById(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.DeleteOneByID(id); err != nil {
		if errors.Is(err, customerrors.ErrProductNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
