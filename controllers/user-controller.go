package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nddat1811/golang_mongodb/models"
	"github.com/nddat1811/golang_mongodb/services"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

// CreateUser godoc
// @Summary Create user
// @Schemes
// @Description create Resource directory
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Success 400 {string} string "error"
// @Success 404 {string} string "error"
// @Success 500 {string} string "error"
// @Router /user/create [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetUser godoc
// @Summary Get user by name
// @Schemes
// @Description create Resource directory
// @Tags Users
// @Param name path string true "name"
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Success 400 {string} string "error"
// @Success 404 {string} string "error"
// @Success 500 {string} string "error"
// @Router /user/get/{name} [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	var username string = ctx.Param("name")
	fmt.Println(username)
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// GetAllUser godoc
// @Summary Get all user
// @Schemes
// @Description create Resource directory
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Success 400 {string} string "error"
// @Success 404 {string} string "error"
// @Success 500 {string} string "error"
// @Router /user/getall [get]
func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Get user by name
// @Schemes
// @Description create Resource directory
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Success 400 {string} string "error"
// @Success 404 {string} string "error"
// @Success 500 {string} string "error"
// @Router /user/update [put]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// DeleteUser godoc
// @Summary Delete user by name
// @Schemes
// @Description create Resource directory
// @Tags Users
// @Param name path string true "name"
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Success 400 {string} string "error"
// @Success 404 {string} string "error"
// @Success 500 {string} string "error"
// @Router /user/delete/{name} [delete]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var username string = ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PUT("/update", uc.UpdateUser)
	userroute.DELETE("/delete/:name", uc.DeleteUser)
}
