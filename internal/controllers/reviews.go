package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	customerrors "github.com/mateusrlopez/go-market/internal/errors"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/services"
)

type ReviewsController struct {
	service services.ReviewsService
}

func NewReviewsController(service services.ReviewsService) ReviewsController {
	return ReviewsController{
		service: service,
	}
}

func (c ReviewsController) Create(ctx *gin.Context) {
	var input inputs.CreateReviewInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := c.service.Create(input)

	if err != nil {
		if errors.Is(err, customerrors.ErrReviewNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"review": review})
}

func (c ReviewsController) Index(ctx *gin.Context) {
	var input inputs.QueryReviewsInput

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reviews, err := c.service.FindMany(input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func (c ReviewsController) GetOneById(ctx *gin.Context) {
	id := ctx.Param("id")

	review, err := c.service.FindOneByID(id)

	if err != nil {
		if errors.Is(err, customerrors.ErrReviewNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"review": review})
}

func (c ReviewsController) UpdateOneByID(ctx *gin.Context) {
	var input inputs.UpdateReviewInput

	id := ctx.Param("id")

	updated, err := c.service.UpdateOneByID(id, input)

	if err != nil {
		if errors.Is(err, customerrors.ErrReviewNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"review": updated})
}

func (c ReviewsController) RemoveOneByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.DeleteOneByID(id); err != nil {
		if errors.Is(err, customerrors.ErrReviewNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
