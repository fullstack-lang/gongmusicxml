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
var __Empty_print_style__dummysDeclaration__ models.Empty_print_style
var __Empty_print_style_time__dummyDeclaration time.Duration

var mutexEmpty_print_style sync.Mutex

// An Empty_print_styleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmpty_print_style updateEmpty_print_style deleteEmpty_print_style
type Empty_print_styleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Empty_print_styleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmpty_print_style updateEmpty_print_style
type Empty_print_styleInput struct {
	// The Empty_print_style to submit or modify
	// in: body
	Empty_print_style *orm.Empty_print_styleAPI
}

// GetEmpty_print_styles
//
// swagger:route GET /empty_print_styles empty_print_styles getEmpty_print_styles
//
// # Get all empty_print_styles
//
// Responses:
// default: genericError
//
//	200: empty_print_styleDBResponse
func (controller *Controller) GetEmpty_print_styles(c *gin.Context) {

	// source slice
	var empty_print_styleDBs []orm.Empty_print_styleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_print_styles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style.GetDB()

	query := db.Find(&empty_print_styleDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	empty_print_styleAPIs := make([]orm.Empty_print_styleAPI, 0)

	// for each empty_print_style, update fields from the database nullable fields
	for idx := range empty_print_styleDBs {
		empty_print_styleDB := &empty_print_styleDBs[idx]
		_ = empty_print_styleDB
		var empty_print_styleAPI orm.Empty_print_styleAPI

		// insertion point for updating fields
		empty_print_styleAPI.ID = empty_print_styleDB.ID
		empty_print_styleDB.CopyBasicFieldsToEmpty_print_style_WOP(&empty_print_styleAPI.Empty_print_style_WOP)
		empty_print_styleAPI.Empty_print_stylePointersEncoding = empty_print_styleDB.Empty_print_stylePointersEncoding
		empty_print_styleAPIs = append(empty_print_styleAPIs, empty_print_styleAPI)
	}

	c.JSON(http.StatusOK, empty_print_styleAPIs)
}

// PostEmpty_print_style
//
// swagger:route POST /empty_print_styles empty_print_styles postEmpty_print_style
//
// Creates a empty_print_style
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmpty_print_style(c *gin.Context) {

	mutexEmpty_print_style.Lock()
	defer mutexEmpty_print_style.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmpty_print_styles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style.GetDB()

	// Validate input
	var input orm.Empty_print_styleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create empty_print_style
	empty_print_styleDB := orm.Empty_print_styleDB{}
	empty_print_styleDB.Empty_print_stylePointersEncoding = input.Empty_print_stylePointersEncoding
	empty_print_styleDB.CopyBasicFieldsFromEmpty_print_style_WOP(&input.Empty_print_style_WOP)

	query := db.Create(&empty_print_styleDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmpty_print_style.CheckoutPhaseOneInstance(&empty_print_styleDB)
	empty_print_style := backRepo.BackRepoEmpty_print_style.Map_Empty_print_styleDBID_Empty_print_stylePtr[empty_print_styleDB.ID]

	if empty_print_style != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), empty_print_style)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, empty_print_styleDB)
}

// GetEmpty_print_style
//
// swagger:route GET /empty_print_styles/{ID} empty_print_styles getEmpty_print_style
//
// Gets the details for a empty_print_style.
//
// Responses:
// default: genericError
//
//	200: empty_print_styleDBResponse
func (controller *Controller) GetEmpty_print_style(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_print_style", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style.GetDB()

	// Get empty_print_styleDB in DB
	var empty_print_styleDB orm.Empty_print_styleDB
	if err := db.First(&empty_print_styleDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var empty_print_styleAPI orm.Empty_print_styleAPI
	empty_print_styleAPI.ID = empty_print_styleDB.ID
	empty_print_styleAPI.Empty_print_stylePointersEncoding = empty_print_styleDB.Empty_print_stylePointersEncoding
	empty_print_styleDB.CopyBasicFieldsToEmpty_print_style_WOP(&empty_print_styleAPI.Empty_print_style_WOP)

	c.JSON(http.StatusOK, empty_print_styleAPI)
}

// UpdateEmpty_print_style
//
// swagger:route PATCH /empty_print_styles/{ID} empty_print_styles updateEmpty_print_style
//
// # Update a empty_print_style
//
// Responses:
// default: genericError
//
//	200: empty_print_styleDBResponse
func (controller *Controller) UpdateEmpty_print_style(c *gin.Context) {

	mutexEmpty_print_style.Lock()
	defer mutexEmpty_print_style.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEmpty_print_style", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style.GetDB()

	// Validate input
	var input orm.Empty_print_styleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var empty_print_styleDB orm.Empty_print_styleDB

	// fetch the empty_print_style
	query := db.First(&empty_print_styleDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	empty_print_styleDB.CopyBasicFieldsFromEmpty_print_style_WOP(&input.Empty_print_style_WOP)
	empty_print_styleDB.Empty_print_stylePointersEncoding = input.Empty_print_stylePointersEncoding

	query = db.Model(&empty_print_styleDB).Updates(empty_print_styleDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	empty_print_styleNew := new(models.Empty_print_style)
	empty_print_styleDB.CopyBasicFieldsToEmpty_print_style(empty_print_styleNew)

	// redeem pointers
	empty_print_styleDB.DecodePointers(backRepo, empty_print_styleNew)

	// get stage instance from DB instance, and call callback function
	empty_print_styleOld := backRepo.BackRepoEmpty_print_style.Map_Empty_print_styleDBID_Empty_print_stylePtr[empty_print_styleDB.ID]
	if empty_print_styleOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), empty_print_styleOld, empty_print_styleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the empty_print_styleDB
	c.JSON(http.StatusOK, empty_print_styleDB)
}

// DeleteEmpty_print_style
//
// swagger:route DELETE /empty_print_styles/{ID} empty_print_styles deleteEmpty_print_style
//
// # Delete a empty_print_style
//
// default: genericError
//
//	200: empty_print_styleDBResponse
func (controller *Controller) DeleteEmpty_print_style(c *gin.Context) {

	mutexEmpty_print_style.Lock()
	defer mutexEmpty_print_style.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmpty_print_style", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style.GetDB()

	// Get model if exist
	var empty_print_styleDB orm.Empty_print_styleDB
	if err := db.First(&empty_print_styleDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&empty_print_styleDB)

	// get an instance (not staged) from DB instance, and call callback function
	empty_print_styleDeleted := new(models.Empty_print_style)
	empty_print_styleDB.CopyBasicFieldsToEmpty_print_style(empty_print_styleDeleted)

	// get stage instance from DB instance, and call callback function
	empty_print_styleStaged := backRepo.BackRepoEmpty_print_style.Map_Empty_print_styleDBID_Empty_print_stylePtr[empty_print_styleDB.ID]
	if empty_print_styleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), empty_print_styleStaged, empty_print_styleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
