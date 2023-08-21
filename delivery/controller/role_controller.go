package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/model"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/common"
)

type RoleController struct {
	roleUC usecase.RoleUserUseCase
	router *gin.Engine
}

func (r *RoleController) createHandler(c *gin.Context) {
	var role model.RoleUser
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	role.ID = common.GenerateID()
	if err := r.roleUC.RegisterNewRole(role); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	}
	c.JSON(201, role)
}

func (r *RoleController) listHandler(c *gin.Context) {
	roles, err := r.roleUC.FindAllRole()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   roles,
	})
}

func (r *RoleController) getHandler(c *gin.Context) {
	id := c.Param("id")
	role, err := r.roleUC.FindByIdRole(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get Data By Id Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   role,
	})
}

func (r *RoleController) updateHandler(c *gin.Context) {
	var role model.RoleUser
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	existingRole, _ := r.roleUC.FindByIdRole(role.ID)
	role.CreatedAt = existingRole.CreatedAt
	role.UpdatedAt = time.Now()
	if err := r.roleUC.UpdateRole(role); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, role)
}

func (r *RoleController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := r.roleUC.DeleteRole(id); err != nil {
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

func NewRoleController(usecase usecase.RoleUserUseCase, r *gin.Engine) *RoleController {
	controller := RoleController{
		router: r,
		roleUC: usecase,
	}
	rg := r.Group("api/v1")
	rg.POST("/roles", controller.createHandler)
	rg.GET("/roles", controller.listHandler)
	rg.GET("/roles/:id", controller.getHandler)
	rg.PUT("/roles/:id", controller.updateHandler)
	rg.DELETE("/roles/:id", controller.deleteHandler)
	return &controller
}
