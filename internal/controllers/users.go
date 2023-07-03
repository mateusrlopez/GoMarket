package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	customerrors "github.com/mateusrlopez/go-market/internal/errors"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/models"
	"github.com/mateusrlopez/go-market/internal/services"
)

type UsersController struct {
	usersService services.UsersService
}

func NewUsersController(usersService services.UsersService) UsersController {
	return UsersController{
		usersService: usersService,
	}
}

func (c UsersController) Index(ctx *gin.Context) {
	users, err := c.usersService.FindMany()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (c UsersController) GetOneById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.usersService.FindOneByID(id)

	if err != nil {
		if errors.Is(err, customerrors.ErrUserNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (c UsersController) UpdateOneByID(ctx *gin.Context) {
	var input inputs.UpdateUserInput

	id := ctx.Param("id")
	user := ctx.MustGet("user").(models.User)

	if user.ID != id {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "cannot update another user's data"})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.usersService.UpdateOneByID(id, input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": updated})
}

func (c UsersController) RemoveOneByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user := ctx.MustGet("user").(models.User)

	if user.ID != id {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "cannot delete another user's data"})
		return
	}

	if err := c.usersService.DeleteOneByID(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
