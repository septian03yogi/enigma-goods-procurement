package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/model"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/common"
)

type EmployeeController struct {
	employeeUC usecase.EmployeeUseCase
	router     *gin.Engine
}

func (e *EmployeeController) createHandler(c *gin.Context) {
	// inisiasi struct kosong untuk di lakukan bind di body json (POSTMAN)
	var employee model.Employee
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()
	// cek error ketika melakukan bind body JSON, keluarkan status code 400 (bad request - CLIENT)
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return // ini harus ada supaya gak diteruskan ke bawah
	}
	// cek error ketikan server tidak merespon atau ada kesalahan, keluarkan status code 500 (internal server error - SERVER)
	employee.ID = common.GenerateID()
	if err := e.employeeUC.RegisterNewEmployee(employee); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return // ini harus ada supaya gak diteruskan ke bawah
	}
	// jika semua aman dan tidak ada error
	c.JSON(201, employee)
}

func (e *EmployeeController) listHandler(c *gin.Context) {
	employees, err := e.employeeUC.FindAllEmployee()
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
		"data":   employees,
	})
}

func (e *EmployeeController) getHandler(c *gin.Context) {
	id := c.Param("id")
	employee, err := e.employeeUC.FindByIdEmployee(id)
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
		"data":   employee,
	})
}

func (e *EmployeeController) updateHandler(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	existingEmployee, _ := e.employeeUC.FindByIdEmployee(employee.ID)
	employee.CreatedAt = existingEmployee.CreatedAt
	employee.UpdatedAt = time.Now()
	if err := e.employeeUC.UpdateEmployee(employee); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, employee)
}

func (e *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := e.employeeUC.DeleteEmployee(id); err != nil {
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

func NewEmployeeController(usecase usecase.EmployeeUseCase, r *gin.Engine) *EmployeeController {
	controller := EmployeeController{
		router:     r,
		employeeUC: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/employees", controller.createHandler)
	rg.GET("/employees", controller.listHandler)
	rg.GET("/employees/:id", controller.getHandler)
	rg.PUT("/employees/:id", controller.updateHandler)
	rg.DELETE("/employees/:id", controller.deleteHandler)

	return &controller
}
