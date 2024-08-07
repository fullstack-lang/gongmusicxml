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
var __Ending__dummysDeclaration__ models.Ending
var __Ending_time__dummyDeclaration time.Duration

var mutexEnding sync.Mutex

// An EndingID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEnding updateEnding deleteEnding
type EndingID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EndingInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEnding updateEnding
type EndingInput struct {
	// The Ending to submit or modify
	// in: body
	Ending *orm.EndingAPI
}

// GetEndings
//
// swagger:route GET /endings endings getEndings
//
// # Get all endings
//
// Responses:
// default: genericError
//
//	200: endingDBResponse
func (controller *Controller) GetEndings(c *gin.Context) {

	// source slice
	var endingDBs []orm.EndingDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEndings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnding.GetDB()

	query := db.Find(&endingDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	endingAPIs := make([]orm.EndingAPI, 0)

	// for each ending, update fields from the database nullable fields
	for idx := range endingDBs {
		endingDB := &endingDBs[idx]
		_ = endingDB
		var endingAPI orm.EndingAPI

		// insertion point for updating fields
		endingAPI.ID = endingDB.ID
		endingDB.CopyBasicFieldsToEnding_WOP(&endingAPI.Ending_WOP)
		endingAPI.EndingPointersEncoding = endingDB.EndingPointersEncoding
		endingAPIs = append(endingAPIs, endingAPI)
	}

	c.JSON(http.StatusOK, endingAPIs)
}

// PostEnding
//
// swagger:route POST /endings endings postEnding
//
// Creates a ending
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEnding(c *gin.Context) {

	mutexEnding.Lock()
	defer mutexEnding.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEndings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnding.GetDB()

	// Validate input
	var input orm.EndingAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create ending
	endingDB := orm.EndingDB{}
	endingDB.EndingPointersEncoding = input.EndingPointersEncoding
	endingDB.CopyBasicFieldsFromEnding_WOP(&input.Ending_WOP)

	query := db.Create(&endingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEnding.CheckoutPhaseOneInstance(&endingDB)
	ending := backRepo.BackRepoEnding.Map_EndingDBID_EndingPtr[endingDB.ID]

	if ending != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), ending)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, endingDB)
}

// GetEnding
//
// swagger:route GET /endings/{ID} endings getEnding
//
// Gets the details for a ending.
//
// Responses:
// default: genericError
//
//	200: endingDBResponse
func (controller *Controller) GetEnding(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEnding", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnding.GetDB()

	// Get endingDB in DB
	var endingDB orm.EndingDB
	if err := db.First(&endingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var endingAPI orm.EndingAPI
	endingAPI.ID = endingDB.ID
	endingAPI.EndingPointersEncoding = endingDB.EndingPointersEncoding
	endingDB.CopyBasicFieldsToEnding_WOP(&endingAPI.Ending_WOP)

	c.JSON(http.StatusOK, endingAPI)
}

// UpdateEnding
//
// swagger:route PATCH /endings/{ID} endings updateEnding
//
// # Update a ending
//
// Responses:
// default: genericError
//
//	200: endingDBResponse
func (controller *Controller) UpdateEnding(c *gin.Context) {

	mutexEnding.Lock()
	defer mutexEnding.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEnding", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnding.GetDB()

	// Validate input
	var input orm.EndingAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var endingDB orm.EndingDB

	// fetch the ending
	query := db.First(&endingDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	endingDB.CopyBasicFieldsFromEnding_WOP(&input.Ending_WOP)
	endingDB.EndingPointersEncoding = input.EndingPointersEncoding

	query = db.Model(&endingDB).Updates(endingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	endingNew := new(models.Ending)
	endingDB.CopyBasicFieldsToEnding(endingNew)

	// redeem pointers
	endingDB.DecodePointers(backRepo, endingNew)

	// get stage instance from DB instance, and call callback function
	endingOld := backRepo.BackRepoEnding.Map_EndingDBID_EndingPtr[endingDB.ID]
	if endingOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), endingOld, endingNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the endingDB
	c.JSON(http.StatusOK, endingDB)
}

// DeleteEnding
//
// swagger:route DELETE /endings/{ID} endings deleteEnding
//
// # Delete a ending
//
// default: genericError
//
//	200: endingDBResponse
func (controller *Controller) DeleteEnding(c *gin.Context) {

	mutexEnding.Lock()
	defer mutexEnding.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEnding", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEnding.GetDB()

	// Get model if exist
	var endingDB orm.EndingDB
	if err := db.First(&endingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&endingDB)

	// get an instance (not staged) from DB instance, and call callback function
	endingDeleted := new(models.Ending)
	endingDB.CopyBasicFieldsToEnding(endingDeleted)

	// get stage instance from DB instance, and call callback function
	endingStaged := backRepo.BackRepoEnding.Map_EndingDBID_EndingPtr[endingDB.ID]
	if endingStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), endingStaged, endingDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
