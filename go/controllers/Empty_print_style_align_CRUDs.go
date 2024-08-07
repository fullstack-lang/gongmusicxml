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
var __Empty_print_style_align__dummysDeclaration__ models.Empty_print_style_align
var __Empty_print_style_align_time__dummyDeclaration time.Duration

var mutexEmpty_print_style_align sync.Mutex

// An Empty_print_style_alignID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmpty_print_style_align updateEmpty_print_style_align deleteEmpty_print_style_align
type Empty_print_style_alignID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Empty_print_style_alignInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmpty_print_style_align updateEmpty_print_style_align
type Empty_print_style_alignInput struct {
	// The Empty_print_style_align to submit or modify
	// in: body
	Empty_print_style_align *orm.Empty_print_style_alignAPI
}

// GetEmpty_print_style_aligns
//
// swagger:route GET /empty_print_style_aligns empty_print_style_aligns getEmpty_print_style_aligns
//
// # Get all empty_print_style_aligns
//
// Responses:
// default: genericError
//
//	200: empty_print_style_alignDBResponse
func (controller *Controller) GetEmpty_print_style_aligns(c *gin.Context) {

	// source slice
	var empty_print_style_alignDBs []orm.Empty_print_style_alignDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_print_style_aligns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align.GetDB()

	query := db.Find(&empty_print_style_alignDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	empty_print_style_alignAPIs := make([]orm.Empty_print_style_alignAPI, 0)

	// for each empty_print_style_align, update fields from the database nullable fields
	for idx := range empty_print_style_alignDBs {
		empty_print_style_alignDB := &empty_print_style_alignDBs[idx]
		_ = empty_print_style_alignDB
		var empty_print_style_alignAPI orm.Empty_print_style_alignAPI

		// insertion point for updating fields
		empty_print_style_alignAPI.ID = empty_print_style_alignDB.ID
		empty_print_style_alignDB.CopyBasicFieldsToEmpty_print_style_align_WOP(&empty_print_style_alignAPI.Empty_print_style_align_WOP)
		empty_print_style_alignAPI.Empty_print_style_alignPointersEncoding = empty_print_style_alignDB.Empty_print_style_alignPointersEncoding
		empty_print_style_alignAPIs = append(empty_print_style_alignAPIs, empty_print_style_alignAPI)
	}

	c.JSON(http.StatusOK, empty_print_style_alignAPIs)
}

// PostEmpty_print_style_align
//
// swagger:route POST /empty_print_style_aligns empty_print_style_aligns postEmpty_print_style_align
//
// Creates a empty_print_style_align
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmpty_print_style_align(c *gin.Context) {

	mutexEmpty_print_style_align.Lock()
	defer mutexEmpty_print_style_align.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmpty_print_style_aligns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align.GetDB()

	// Validate input
	var input orm.Empty_print_style_alignAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create empty_print_style_align
	empty_print_style_alignDB := orm.Empty_print_style_alignDB{}
	empty_print_style_alignDB.Empty_print_style_alignPointersEncoding = input.Empty_print_style_alignPointersEncoding
	empty_print_style_alignDB.CopyBasicFieldsFromEmpty_print_style_align_WOP(&input.Empty_print_style_align_WOP)

	query := db.Create(&empty_print_style_alignDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmpty_print_style_align.CheckoutPhaseOneInstance(&empty_print_style_alignDB)
	empty_print_style_align := backRepo.BackRepoEmpty_print_style_align.Map_Empty_print_style_alignDBID_Empty_print_style_alignPtr[empty_print_style_alignDB.ID]

	if empty_print_style_align != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), empty_print_style_align)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, empty_print_style_alignDB)
}

// GetEmpty_print_style_align
//
// swagger:route GET /empty_print_style_aligns/{ID} empty_print_style_aligns getEmpty_print_style_align
//
// Gets the details for a empty_print_style_align.
//
// Responses:
// default: genericError
//
//	200: empty_print_style_alignDBResponse
func (controller *Controller) GetEmpty_print_style_align(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_print_style_align", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align.GetDB()

	// Get empty_print_style_alignDB in DB
	var empty_print_style_alignDB orm.Empty_print_style_alignDB
	if err := db.First(&empty_print_style_alignDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var empty_print_style_alignAPI orm.Empty_print_style_alignAPI
	empty_print_style_alignAPI.ID = empty_print_style_alignDB.ID
	empty_print_style_alignAPI.Empty_print_style_alignPointersEncoding = empty_print_style_alignDB.Empty_print_style_alignPointersEncoding
	empty_print_style_alignDB.CopyBasicFieldsToEmpty_print_style_align_WOP(&empty_print_style_alignAPI.Empty_print_style_align_WOP)

	c.JSON(http.StatusOK, empty_print_style_alignAPI)
}

// UpdateEmpty_print_style_align
//
// swagger:route PATCH /empty_print_style_aligns/{ID} empty_print_style_aligns updateEmpty_print_style_align
//
// # Update a empty_print_style_align
//
// Responses:
// default: genericError
//
//	200: empty_print_style_alignDBResponse
func (controller *Controller) UpdateEmpty_print_style_align(c *gin.Context) {

	mutexEmpty_print_style_align.Lock()
	defer mutexEmpty_print_style_align.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEmpty_print_style_align", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align.GetDB()

	// Validate input
	var input orm.Empty_print_style_alignAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var empty_print_style_alignDB orm.Empty_print_style_alignDB

	// fetch the empty_print_style_align
	query := db.First(&empty_print_style_alignDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	empty_print_style_alignDB.CopyBasicFieldsFromEmpty_print_style_align_WOP(&input.Empty_print_style_align_WOP)
	empty_print_style_alignDB.Empty_print_style_alignPointersEncoding = input.Empty_print_style_alignPointersEncoding

	query = db.Model(&empty_print_style_alignDB).Updates(empty_print_style_alignDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	empty_print_style_alignNew := new(models.Empty_print_style_align)
	empty_print_style_alignDB.CopyBasicFieldsToEmpty_print_style_align(empty_print_style_alignNew)

	// redeem pointers
	empty_print_style_alignDB.DecodePointers(backRepo, empty_print_style_alignNew)

	// get stage instance from DB instance, and call callback function
	empty_print_style_alignOld := backRepo.BackRepoEmpty_print_style_align.Map_Empty_print_style_alignDBID_Empty_print_style_alignPtr[empty_print_style_alignDB.ID]
	if empty_print_style_alignOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), empty_print_style_alignOld, empty_print_style_alignNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the empty_print_style_alignDB
	c.JSON(http.StatusOK, empty_print_style_alignDB)
}

// DeleteEmpty_print_style_align
//
// swagger:route DELETE /empty_print_style_aligns/{ID} empty_print_style_aligns deleteEmpty_print_style_align
//
// # Delete a empty_print_style_align
//
// default: genericError
//
//	200: empty_print_style_alignDBResponse
func (controller *Controller) DeleteEmpty_print_style_align(c *gin.Context) {

	mutexEmpty_print_style_align.Lock()
	defer mutexEmpty_print_style_align.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmpty_print_style_align", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align.GetDB()

	// Get model if exist
	var empty_print_style_alignDB orm.Empty_print_style_alignDB
	if err := db.First(&empty_print_style_alignDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&empty_print_style_alignDB)

	// get an instance (not staged) from DB instance, and call callback function
	empty_print_style_alignDeleted := new(models.Empty_print_style_align)
	empty_print_style_alignDB.CopyBasicFieldsToEmpty_print_style_align(empty_print_style_alignDeleted)

	// get stage instance from DB instance, and call callback function
	empty_print_style_alignStaged := backRepo.BackRepoEmpty_print_style_align.Map_Empty_print_style_alignDBID_Empty_print_style_alignPtr[empty_print_style_alignDB.ID]
	if empty_print_style_alignStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), empty_print_style_alignStaged, empty_print_style_alignDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
