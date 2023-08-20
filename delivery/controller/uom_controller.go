package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/model"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/common"
)

type UomController struct {
	uomUC  usecase.UomUseCase
	router *gin.Engine
}

func (u *UomController) createHandler(c *gin.Context) {
	var uom model.Uom
	uom.CreatedAt = time.Now()
	uom.UpdatedAt = time.Now()
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	uom.Id = common.GenerateID()
	if err := u.uomUC.RegisterNewUom(uom); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	}
	c.JSON(201, uom)
}

func (u *UomController) listHandler(c *gin.Context) {
	uoms, err := u.uomUC.FindAllUom()
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
		"data":   uoms,
	})
}

func (u *UomController) getHandler(c *gin.Context) {
	id := c.Param("id")
	uom, err := u.uomUC.FindUomById(id)
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
		"data":   uom,
	})
}

func (u *UomController) updateHandler(c *gin.Context) {
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	existingUom, _ := u.uomUC.FindUomById(uom.Id)
	uom.CreatedAt = existingUom.CreatedAt
	uom.UpdatedAt = time.Now()
	if err := u.uomUC.UpdateUom(uom); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, uom)
}

func (u *UomController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := u.uomUC.DeleteUom(id); err != nil {
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

func NewUomUseController(usecase usecase.UomUseCase, r *gin.Engine) *UomController {
	controller := UomController{
		router: r,
		uomUC:  usecase,
	}
	rg := r.Group("api/v1")
	rg.POST("/uoms", controller.createHandler)
	rg.GET("/uoms", controller.listHandler)
	rg.GET("/uoms/:id", controller.getHandler)
	rg.PUT("/uoms/:id", controller.updateHandler)
	rg.DELETE("/uoms/:id", controller.deleteHandler)

	return &controller
}
