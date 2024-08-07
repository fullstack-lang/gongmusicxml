// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongmusicxml/go/models"
	"github.com/fullstack-lang/gongmusicxml/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Bar_style_color__dummysDeclaration__ models.Bar_style_color
var __Bar_style_color_time__dummyDeclaration time.Duration

var mutexBar_style_color sync.Mutex

// An Bar_style_colorID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBar_style_color updateBar_style_color deleteBar_style_color
type Bar_style_colorID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Bar_style_colorInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBar_style_color updateBar_style_color
type Bar_style_colorInput struct {
	// The Bar_style_color to submit or modify
	// in: body
	Bar_style_color *orm.Bar_style_colorAPI
}

// GetBar_style_colors
//
// swagger:route GET /bar_style_colors bar_style_colors getBar_style_colors
//
// # Get all bar_style_colors
//
// Responses:
// default: genericError
//
//	200: bar_style_colorDBResponse
func (controller *Controller) GetBar_style_colors(c *gin.Context) {

	// source slice
	var bar_style_colorDBs []orm.Bar_style_colorDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBar_style_colors", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBar_style_color.GetDB()

	query := db.Find(&bar_style_colorDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bar_style_colorAPIs := make([]orm.Bar_style_colorAPI, 0)

	// for each bar_style_color, update fields from the database nullable fields
	for idx := range bar_style_colorDBs {
		bar_style_colorDB := &bar_style_colorDBs[idx]
		_ = bar_style_colorDB
		var bar_style_colorAPI orm.Bar_style_colorAPI

		// insertion point for updating fields
		bar_style_colorAPI.ID = bar_style_colorDB.ID
		bar_style_colorDB.CopyBasicFieldsToBar_style_color_WOP(&bar_style_colorAPI.Bar_style_color_WOP)
		bar_style_colorAPI.Bar_style_colorPointersEncoding = bar_style_colorDB.Bar_style_colorPointersEncoding
		bar_style_colorAPIs = append(bar_style_colorAPIs, bar_style_colorAPI)
	}

	c.JSON(http.StatusOK, bar_style_colorAPIs)
}

// PostBar_style_color
//
// swagger:route POST /bar_style_colors bar_style_colors postBar_style_color
//
// Creates a bar_style_color
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBar_style_color(c *gin.Context) {

	mutexBar_style_color.Lock()
	defer mutexBar_style_color.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBar_style_colors", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBar_style_color.GetDB()

	// Validate input
	var input orm.Bar_style_colorAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bar_style_color
	bar_style_colorDB := orm.Bar_style_colorDB{}
	bar_style_colorDB.Bar_style_colorPointersEncoding = input.Bar_style_colorPointersEncoding
	bar_style_colorDB.CopyBasicFieldsFromBar_style_color_WOP(&input.Bar_style_color_WOP)

	query := db.Create(&bar_style_colorDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBar_style_color.CheckoutPhaseOneInstance(&bar_style_colorDB)
	bar_style_color := backRepo.BackRepoBar_style_color.Map_Bar_style_colorDBID_Bar_style_colorPtr[bar_style_colorDB.ID]

	if bar_style_color != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bar_style_color)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bar_style_colorDB)
}

// GetBar_style_color
//
// swagger:route GET /bar_style_colors/{ID} bar_style_colors getBar_style_color
//
// Gets the details for a bar_style_color.
//
// Responses:
// default: genericError
//
//	200: bar_style_colorDBResponse
func (controller *Controller) GetBar_style_color(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBar_style_color", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBar_style_color.GetDB()

	// Get bar_style_colorDB in DB
	var bar_style_colorDB orm.Bar_style_colorDB
	if err := db.First(&bar_style_colorDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bar_style_colorAPI orm.Bar_style_colorAPI
	bar_style_colorAPI.ID = bar_style_colorDB.ID
	bar_style_colorAPI.Bar_style_colorPointersEncoding = bar_style_colorDB.Bar_style_colorPointersEncoding
	bar_style_colorDB.CopyBasicFieldsToBar_style_color_WOP(&bar_style_colorAPI.Bar_style_color_WOP)

	c.JSON(http.StatusOK, bar_style_colorAPI)
}

// UpdateBar_style_color
//
// swagger:route PATCH /bar_style_colors/{ID} bar_style_colors updateBar_style_color
//
// # Update a bar_style_color
//
// Responses:
// default: genericError
//
//	200: bar_style_colorDBResponse
func (controller *Controller) UpdateBar_style_color(c *gin.Context) {

	mutexBar_style_color.Lock()
	defer mutexBar_style_color.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBar_style_color", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBar_style_color.GetDB()

	// Validate input
	var input orm.Bar_style_colorAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bar_style_colorDB orm.Bar_style_colorDB

	// fetch the bar_style_color
	query := db.First(&bar_style_colorDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bar_style_colorDB.CopyBasicFieldsFromBar_style_color_WOP(&input.Bar_style_color_WOP)
	bar_style_colorDB.Bar_style_colorPointersEncoding = input.Bar_style_colorPointersEncoding

	query = db.Model(&bar_style_colorDB).Updates(bar_style_colorDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bar_style_colorNew := new(models.Bar_style_color)
	bar_style_colorDB.CopyBasicFieldsToBar_style_color(bar_style_colorNew)

	// redeem pointers
	bar_style_colorDB.DecodePointers(backRepo, bar_style_colorNew)

	// get stage instance from DB instance, and call callback function
	bar_style_colorOld := backRepo.BackRepoBar_style_color.Map_Bar_style_colorDBID_Bar_style_colorPtr[bar_style_colorDB.ID]
	if bar_style_colorOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bar_style_colorOld, bar_style_colorNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bar_style_colorDB
	c.JSON(http.StatusOK, bar_style_colorDB)
}

// DeleteBar_style_color
//
// swagger:route DELETE /bar_style_colors/{ID} bar_style_colors deleteBar_style_color
//
// # Delete a bar_style_color
//
// default: genericError
//
//	200: bar_style_colorDBResponse
func (controller *Controller) DeleteBar_style_color(c *gin.Context) {

	mutexBar_style_color.Lock()
	defer mutexBar_style_color.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBar_style_color", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBar_style_color.GetDB()

	// Get model if exist
	var bar_style_colorDB orm.Bar_style_colorDB
	if err := db.First(&bar_style_colorDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bar_style_colorDB)

	// get an instance (not staged) from DB instance, and call callback function
	bar_style_colorDeleted := new(models.Bar_style_color)
	bar_style_colorDB.CopyBasicFieldsToBar_style_color(bar_style_colorDeleted)

	// get stage instance from DB instance, and call callback function
	bar_style_colorStaged := backRepo.BackRepoBar_style_color.Map_Bar_style_colorDBID_Bar_style_colorPtr[bar_style_colorDB.ID]
	if bar_style_colorStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bar_style_colorStaged, bar_style_colorDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
