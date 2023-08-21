package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/model"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/common"
)

type SubmissionContoler struct {
	submissionUC usecase.SubmissionUseCase
	router       *gin.Engine
}

func (s *SubmissionContoler) createHandler(c *gin.Context) {
	var submission model.Submission
	submission.CreatedAt = time.Now()
	submission.UpdatedAt = time.Now()
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	submission.ID = common.GenerateID()
	if err := s.submissionUC.RegisterNewSubmission(submission); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	}
	c.JSON(201, submission)
}

func (s *SubmissionContoler) listHandler(c *gin.Context) {
	submissions, err := s.submissionUC.FindAllSubmission()
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
		"data":   submissions,
	})
}

func (s *SubmissionContoler) getHandler(c *gin.Context) {
	id := c.Param("id")
	submission, err := s.submissionUC.FindSubmissionById(id)
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
		"data":   submission,
	})
}

func (s *SubmissionContoler) updateHandler(c *gin.Context) {
	var submission model.Submission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	existingSubmission, _ := s.submissionUC.FindSubmissionById(submission.ID)
	submission.CreatedAt = existingSubmission.CreatedAt
	submission.UpdatedAt = time.Now()
	if err := s.submissionUC.UpdateSubmission(submission); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, submission)
}

func (s *SubmissionContoler) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := s.submissionUC.DeleteSubmission(id); err != nil {
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

func NewSubmissionController(usecase usecase.SubmissionUseCase, r *gin.Engine) *SubmissionContoler {
	controller := SubmissionContoler{
		router:       r,
		submissionUC: usecase,
	}
	rg := r.Group("api/v1")
	rg.POST("/submissions", controller.createHandler)
	rg.GET("/submissions", controller.listHandler)
	rg.GET("/submissions/:id", controller.getHandler)
	rg.PUT("/submissions/:id", controller.updateHandler)
	rg.DELETE("/submissions/:id", controller.deleteHandler)

	return &controller
}
