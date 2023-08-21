package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/model"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/common"
)

type UserController struct {
	userUC usecase.UserUseCase
	router *gin.Engine
}

func (u *UserController) createHandler(c *gin.Context) {
	var user model.UserCredential
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	user.ID = common.GenerateID()
	if err := u.userUC.RegisterNewUser(user); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	}
	c.JSON(201, user)
}

func (u *UserController) listHandler(c *gin.Context) {
	users, err := u.userUC.FindAllUser()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Get All Data Succesfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   users,
	})
}

func (u *UserController) getHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userUC.FindUserById(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get Data By Id Succesfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   user,
	})
}

func (u *UserController) updateHandler(c *gin.Context) {
	var user model.UserCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	existingUser, _ := u.userUC.FindUserById(user.ID)
	user.CreatedAt = existingUser.CreatedAt
	user.UpdatedAt = time.Now()
	if err := u.userUC.UpdateUser(user); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (u *UserController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := u.userUC.Delete(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	}
	status := map[string]any{
		"code":        204,
		"description": "Delete Data By Id Successfully",
	}
	c.JSON(204, gin.H{
		"status": status,
	})
}

func NewUserController(usecase usecase.UserUseCase, r *gin.Engine) *UserController {
	controller := UserController{
		router: r,
		userUC: usecase,
	}
	rg := r.Group("api/v1")
	rg.POST("/users", controller.createHandler)
	rg.GET("/users", controller.listHandler)
	rg.GET("/users/:id", controller.getHandler)
	rg.PUT("/users/:id", controller.updateHandler)
	rg.DELETE("/users/:id", controller.deleteHandler)

	return &controller
}
