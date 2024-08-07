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
var __Repeat__dummysDeclaration__ models.Repeat
var __Repeat_time__dummyDeclaration time.Duration

var mutexRepeat sync.Mutex

// An RepeatID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRepeat updateRepeat deleteRepeat
type RepeatID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RepeatInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRepeat updateRepeat
type RepeatInput struct {
	// The Repeat to submit or modify
	// in: body
	Repeat *orm.RepeatAPI
}

// GetRepeats
//
// swagger:route GET /repeats repeats getRepeats
//
// # Get all repeats
//
// Responses:
// default: genericError
//
//	200: repeatDBResponse
func (controller *Controller) GetRepeats(c *gin.Context) {

	// source slice
	var repeatDBs []orm.RepeatDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRepeats", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepeat.GetDB()

	query := db.Find(&repeatDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	repeatAPIs := make([]orm.RepeatAPI, 0)

	// for each repeat, update fields from the database nullable fields
	for idx := range repeatDBs {
		repeatDB := &repeatDBs[idx]
		_ = repeatDB
		var repeatAPI orm.RepeatAPI

		// insertion point for updating fields
		repeatAPI.ID = repeatDB.ID
		repeatDB.CopyBasicFieldsToRepeat_WOP(&repeatAPI.Repeat_WOP)
		repeatAPI.RepeatPointersEncoding = repeatDB.RepeatPointersEncoding
		repeatAPIs = append(repeatAPIs, repeatAPI)
	}

	c.JSON(http.StatusOK, repeatAPIs)
}

// PostRepeat
//
// swagger:route POST /repeats repeats postRepeat
//
// Creates a repeat
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRepeat(c *gin.Context) {

	mutexRepeat.Lock()
	defer mutexRepeat.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRepeats", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepeat.GetDB()

	// Validate input
	var input orm.RepeatAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create repeat
	repeatDB := orm.RepeatDB{}
	repeatDB.RepeatPointersEncoding = input.RepeatPointersEncoding
	repeatDB.CopyBasicFieldsFromRepeat_WOP(&input.Repeat_WOP)

	query := db.Create(&repeatDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRepeat.CheckoutPhaseOneInstance(&repeatDB)
	repeat := backRepo.BackRepoRepeat.Map_RepeatDBID_RepeatPtr[repeatDB.ID]

	if repeat != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), repeat)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, repeatDB)
}

// GetRepeat
//
// swagger:route GET /repeats/{ID} repeats getRepeat
//
// Gets the details for a repeat.
//
// Responses:
// default: genericError
//
//	200: repeatDBResponse
func (controller *Controller) GetRepeat(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRepeat", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepeat.GetDB()

	// Get repeatDB in DB
	var repeatDB orm.RepeatDB
	if err := db.First(&repeatDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var repeatAPI orm.RepeatAPI
	repeatAPI.ID = repeatDB.ID
	repeatAPI.RepeatPointersEncoding = repeatDB.RepeatPointersEncoding
	repeatDB.CopyBasicFieldsToRepeat_WOP(&repeatAPI.Repeat_WOP)

	c.JSON(http.StatusOK, repeatAPI)
}

// UpdateRepeat
//
// swagger:route PATCH /repeats/{ID} repeats updateRepeat
//
// # Update a repeat
//
// Responses:
// default: genericError
//
//	200: repeatDBResponse
func (controller *Controller) UpdateRepeat(c *gin.Context) {

	mutexRepeat.Lock()
	defer mutexRepeat.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRepeat", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepeat.GetDB()

	// Validate input
	var input orm.RepeatAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var repeatDB orm.RepeatDB

	// fetch the repeat
	query := db.First(&repeatDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	repeatDB.CopyBasicFieldsFromRepeat_WOP(&input.Repeat_WOP)
	repeatDB.RepeatPointersEncoding = input.RepeatPointersEncoding

	query = db.Model(&repeatDB).Updates(repeatDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	repeatNew := new(models.Repeat)
	repeatDB.CopyBasicFieldsToRepeat(repeatNew)

	// redeem pointers
	repeatDB.DecodePointers(backRepo, repeatNew)

	// get stage instance from DB instance, and call callback function
	repeatOld := backRepo.BackRepoRepeat.Map_RepeatDBID_RepeatPtr[repeatDB.ID]
	if repeatOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), repeatOld, repeatNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the repeatDB
	c.JSON(http.StatusOK, repeatDB)
}

// DeleteRepeat
//
// swagger:route DELETE /repeats/{ID} repeats deleteRepeat
//
// # Delete a repeat
//
// default: genericError
//
//	200: repeatDBResponse
func (controller *Controller) DeleteRepeat(c *gin.Context) {

	mutexRepeat.Lock()
	defer mutexRepeat.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRepeat", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepeat.GetDB()

	// Get model if exist
	var repeatDB orm.RepeatDB
	if err := db.First(&repeatDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&repeatDB)

	// get an instance (not staged) from DB instance, and call callback function
	repeatDeleted := new(models.Repeat)
	repeatDB.CopyBasicFieldsToRepeat(repeatDeleted)

	// get stage instance from DB instance, and call callback function
	repeatStaged := backRepo.BackRepoRepeat.Map_RepeatDBID_RepeatPtr[repeatDB.ID]
	if repeatStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), repeatStaged, repeatDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
