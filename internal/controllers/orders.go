package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	customerrors "github.com/mateusrlopez/go-market/internal/errors"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/services"
)

type OrdersController struct {
	service services.OrdersService
}

func NewOrdersController(service services.OrdersService) OrdersController {
	return OrdersController{
		service: service,
	}
}

func (c OrdersController) Create(ctx *gin.Context) {
	var input inputs.CreateOrderInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := c.service.Create(input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"order": order})
}

func (c OrdersController) Index(ctx *gin.Context) {
	var input inputs.QueryOrdersInput

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := c.service.FindMany(input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (c OrdersController) GetOneByID(ctx *gin.Context) {
	id := ctx.Param("id")

	order, err := c.service.FindOneByID(id)

	if err != nil {
		if errors.Is(err, customerrors.ErrOrderNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"order": order})
}

func (c OrdersController) UpdateOneByID(ctx *gin.Context) {
	var input inputs.UpdateOrderInput

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.service.UpdateOneByID(id, input)

	if err != nil {
		if errors.Is(err, customerrors.ErrOrderNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"order": updated})
}

func (c OrdersController) RemoveOneByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.DeleteOneByID(id); err != nil {
		if errors.Is(err, customerrors.ErrOrderNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
