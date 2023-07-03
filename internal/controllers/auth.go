package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/models"
	"github.com/mateusrlopez/go-market/internal/services"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (c AuthController) Register(ctx *gin.Context) {
	var input inputs.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.authService.Register(input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := c.authService.AssignToken(user.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Token", token)
	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func (c AuthController) Login(ctx *gin.Context) {
	var input inputs.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.authService.ValidateLogin(input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := c.authService.AssignToken(user.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Token", token)
	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func (c AuthController) Me(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (c AuthController) Logout(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	if err := c.authService.Logout(user.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
