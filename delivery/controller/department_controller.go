package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/model"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/common"
)

type DepartmentController struct {
	departmentUC usecase.DepartmentUseCase
	router       *gin.Engine
}

func (d *DepartmentController) createHandler(c *gin.Context) {
	// inisiasi struct kosong untuk di lakukan bind di body json (POSTMAN)
	var department model.Department
	department.CreatedAt = time.Now()
	department.UpdatedAt = time.Now()
	// cek error ketika melakukan bind body JSON, keluarkan status code 400 (bad request - CLIENT)
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return // ini harus ada supaya gak diteruskan ke bawah
	}
	// cek error ketikan server tidak merespon atau ada kesalahan, keluarkan status code 500 (internal server error - SERVER)
	department.Id = common.GenerateID()
	if err := d.departmentUC.RegisterNewDepartment(department); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return // ini harus ada supaya gak diteruskan ke bawah
	}
	// jika semua aman dan tidak ada error
	c.JSON(201, department)
}

func (d *DepartmentController) listHandler(c *gin.Context) {
	departments, err := d.departmentUC.FindAllDepartment()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	// status : code, description
	// data : uoms
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   departments,
	})
}

func (d *DepartmentController) getHandler(c *gin.Context) {
	id := c.Param("id")
	department, err := d.departmentUC.FindByIdDepartment(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get Data by ID Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   department,
	})
}

func (d *DepartmentController) updateHandler(c *gin.Context) {
	var department model.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	existingDepartment, _ := d.departmentUC.FindByIdDepartment(department.Id)
	department.CreatedAt = existingDepartment.CreatedAt
	department.UpdatedAt = time.Now()
	if err := d.departmentUC.UpdateDepartment(department); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, department)
}

func (d *DepartmentController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := d.departmentUC.DeleteDepartment(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	}
	status := map[string]any{
		"code":        204,
		"description": "Delete Data By Id Success",
	}
	c.JSON(204, gin.H{
		"status": status,
	})
}

func NewDepartmentController(usecase usecase.DepartmentUseCase, r *gin.Engine) *DepartmentController {
	controller := DepartmentController{
		router:       r,
		departmentUC: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/departments", controller.createHandler)
	rg.GET("/departments", controller.listHandler)
	rg.GET("/departments/:id", controller.getHandler)
	rg.PUT("/departments/:id", controller.updateHandler)
	rg.DELETE("/departments/:id", controller.deleteHandler)

	return &controller
}
