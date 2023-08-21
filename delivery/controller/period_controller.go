package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/model"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/common"
)

type PeriodController struct {
	periodUC usecase.PeriodUseCase
	router   *gin.Engine
}

func (p *PeriodController) createHandler(c *gin.Context) {
	var period model.Period
	period.CreatedAt = time.Now()
	period.UpdatedAt = time.Now()
	if err := c.ShouldBindJSON(&period); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	period.ID = common.GenerateID()
	if err := p.periodUC.RegisterNewPeriod(period); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	}
	c.JSON(201, period)
}

func (p *PeriodController) listHandler(c *gin.Context) {
	periods, err := p.periodUC.FindAllPeriod()
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
		"data":   periods,
	})
}

func (p *PeriodController) getHandler(c *gin.Context) {
	id := c.Param("id")
	period, err := p.periodUC.FindByIdPeriod(id)
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
		"data":   period,
	})
}

func (p *PeriodController) updateHandler(c *gin.Context) {
	var period model.Period
	if err := c.ShouldBindJSON(&period); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	existingPeriod, _ := p.periodUC.FindByIdPeriod(period.ID)
	period.CreatedAt = existingPeriod.CreatedAt
	period.UpdatedAt = time.Now()
	if err := p.periodUC.UpdatePeriod(period); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, period)
}

func (p *PeriodController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := p.periodUC.DeletePeriod(id); err != nil {
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

func NewPeriodController(usecase usecase.PeriodUseCase, r *gin.Engine) *PeriodController {
	controller := PeriodController{
		router:   r,
		periodUC: usecase,
	}
	rg := r.Group("api/v1")
	rg.POST("/periods", controller.createHandler)
	rg.GET("/periods", controller.listHandler)
	rg.GET("/periods/:id", controller.getHandler)
	rg.PUT("/periods/:id", controller.updateHandler)
	rg.DELETE("/periods/:id", controller.deleteHandler)

	return &controller
}
