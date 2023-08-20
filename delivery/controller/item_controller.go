package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/model"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/common"
)

type ItemController struct {
	itemUC usecase.ItemUseCase
	router *gin.Engine
}

func (i *ItemController) createHandler(c *gin.Context) {
	// inisiasi struct kosong untuk di lakukan bind di body json (POSTMAN)
	var item model.Item
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	// cek error ketika melakukan bind body JSON, keluarkan status code 400 (bad request - CLIENT)
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return // ini harus ada supaya gak diteruskan ke bawah
	}
	// cek error ketikan server tidak merespon atau ada kesalahan, keluarkan status code 500 (internal server error - SERVER)
	item.Id = common.GenerateID()
	if err := i.itemUC.RegisterNewItem(item); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return // ini harus ada supaya gak diteruskan ke bawah
	}
	// jika semua aman dan tidak ada error
	c.JSON(201, item)
}

func (i *ItemController) listHandler(c *gin.Context) {
	items, err := i.itemUC.FindAllItem()
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
		"data":   items,
	})
}

func (i *ItemController) getHandler(c *gin.Context) {
	id := c.Param("id")
	item, err := i.itemUC.FindByIdItem(id)
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
		"data":   item,
	})
}

func (i *ItemController) updateHandler(c *gin.Context) {
	var item model.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	existingItem, _ := i.itemUC.FindByIdItem(item.Id)
	item.CreatedAt = existingItem.CreatedAt
	item.UpdatedAt = time.Now()
	if err := i.itemUC.UpdateItem(item); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, item)
}

func (i *ItemController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := i.itemUC.DeleteItem(id); err != nil {
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

func NewItemController(usecase usecase.ItemUseCase, r *gin.Engine) *ItemController {
	controller := ItemController{
		router: r,
		itemUC: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/items", controller.createHandler)
	rg.GET("/items", controller.listHandler)
	rg.GET("/items/:id", controller.getHandler)
	rg.PUT("/items/:id", controller.updateHandler)
	rg.DELETE("/items/:id", controller.deleteHandler)

	return &controller
}
