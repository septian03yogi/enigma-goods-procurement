package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/model"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/common"
)

type SubStatusController struct {
	subStatusUC usecase.SubmisisonStatusUseCase
	router      *gin.Engine
}

func (s *SubStatusController) createHandler(c *gin.Context) {
	var subStatus model.SubmisisonStatus
	subStatus.CreatedAt = time.Now()
	subStatus.UpdatedAt = time.Now()
	if err := c.ShouldBindJSON(&subStatus); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	subStatus.ID = common.GenerateID()
	if err := s.subStatusUC.RegisterNewSubStatus(subStatus); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	}
	c.JSON(201, subStatus)
}

func (s *SubStatusController) listHandler(c *gin.Context) {
	subStatus, err := s.subStatusUC.FindAllSubStatus()
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
		"data":   subStatus,
	})
}

func (s *SubStatusController) getHandler(c *gin.Context) {
	id := c.Param("id")
	subStatus, err := s.subStatusUC.FindSubStatusById(id)
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
		"data":   subStatus,
	})
}

func (s *SubStatusController) updateHandler(c *gin.Context) {
	var subStatus model.SubmisisonStatus
	if err := c.ShouldBindJSON(&subStatus); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	existingSubStatus, _ := s.subStatusUC.FindSubStatusById(subStatus.ID)
	subStatus.CreatedAt = existingSubStatus.CreatedAt
	subStatus.UpdatedAt = time.Now()
	if err := s.subStatusUC.UpdateSubStatus(subStatus); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, subStatus)
}

func (s *SubStatusController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := s.subStatusUC.DeleteSubStatus(id); err != nil {
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

func NewSubStatusController(usecase usecase.SubmisisonStatusUseCase, r *gin.Engine) *SubStatusController {
	controller := SubStatusController{
		router:      r,
		subStatusUC: usecase,
	}
	rg := r.Group("api/v1")
	rg.POST("/substatus", controller.createHandler)
	rg.GET("/substatus", controller.listHandler)
	rg.GET("/substatus/:id", controller.getHandler)
	rg.PUT("/substatus/:id", controller.updateHandler)
	rg.DELETE("/substatus/:id", controller.deleteHandler)

	return &controller
}
