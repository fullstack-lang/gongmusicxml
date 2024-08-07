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
var __Cancel__dummysDeclaration__ models.Cancel
var __Cancel_time__dummyDeclaration time.Duration

var mutexCancel sync.Mutex

// An CancelID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCancel updateCancel deleteCancel
type CancelID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CancelInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCancel updateCancel
type CancelInput struct {
	// The Cancel to submit or modify
	// in: body
	Cancel *orm.CancelAPI
}

// GetCancels
//
// swagger:route GET /cancels cancels getCancels
//
// # Get all cancels
//
// Responses:
// default: genericError
//
//	200: cancelDBResponse
func (controller *Controller) GetCancels(c *gin.Context) {

	// source slice
	var cancelDBs []orm.CancelDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCancels", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCancel.GetDB()

	query := db.Find(&cancelDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	cancelAPIs := make([]orm.CancelAPI, 0)

	// for each cancel, update fields from the database nullable fields
	for idx := range cancelDBs {
		cancelDB := &cancelDBs[idx]
		_ = cancelDB
		var cancelAPI orm.CancelAPI

		// insertion point for updating fields
		cancelAPI.ID = cancelDB.ID
		cancelDB.CopyBasicFieldsToCancel_WOP(&cancelAPI.Cancel_WOP)
		cancelAPI.CancelPointersEncoding = cancelDB.CancelPointersEncoding
		cancelAPIs = append(cancelAPIs, cancelAPI)
	}

	c.JSON(http.StatusOK, cancelAPIs)
}

// PostCancel
//
// swagger:route POST /cancels cancels postCancel
//
// Creates a cancel
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCancel(c *gin.Context) {

	mutexCancel.Lock()
	defer mutexCancel.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCancels", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCancel.GetDB()

	// Validate input
	var input orm.CancelAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create cancel
	cancelDB := orm.CancelDB{}
	cancelDB.CancelPointersEncoding = input.CancelPointersEncoding
	cancelDB.CopyBasicFieldsFromCancel_WOP(&input.Cancel_WOP)

	query := db.Create(&cancelDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCancel.CheckoutPhaseOneInstance(&cancelDB)
	cancel := backRepo.BackRepoCancel.Map_CancelDBID_CancelPtr[cancelDB.ID]

	if cancel != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), cancel)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, cancelDB)
}

// GetCancel
//
// swagger:route GET /cancels/{ID} cancels getCancel
//
// Gets the details for a cancel.
//
// Responses:
// default: genericError
//
//	200: cancelDBResponse
func (controller *Controller) GetCancel(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCancel", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCancel.GetDB()

	// Get cancelDB in DB
	var cancelDB orm.CancelDB
	if err := db.First(&cancelDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var cancelAPI orm.CancelAPI
	cancelAPI.ID = cancelDB.ID
	cancelAPI.CancelPointersEncoding = cancelDB.CancelPointersEncoding
	cancelDB.CopyBasicFieldsToCancel_WOP(&cancelAPI.Cancel_WOP)

	c.JSON(http.StatusOK, cancelAPI)
}

// UpdateCancel
//
// swagger:route PATCH /cancels/{ID} cancels updateCancel
//
// # Update a cancel
//
// Responses:
// default: genericError
//
//	200: cancelDBResponse
func (controller *Controller) UpdateCancel(c *gin.Context) {

	mutexCancel.Lock()
	defer mutexCancel.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCancel", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCancel.GetDB()

	// Validate input
	var input orm.CancelAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cancelDB orm.CancelDB

	// fetch the cancel
	query := db.First(&cancelDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	cancelDB.CopyBasicFieldsFromCancel_WOP(&input.Cancel_WOP)
	cancelDB.CancelPointersEncoding = input.CancelPointersEncoding

	query = db.Model(&cancelDB).Updates(cancelDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	cancelNew := new(models.Cancel)
	cancelDB.CopyBasicFieldsToCancel(cancelNew)

	// redeem pointers
	cancelDB.DecodePointers(backRepo, cancelNew)

	// get stage instance from DB instance, and call callback function
	cancelOld := backRepo.BackRepoCancel.Map_CancelDBID_CancelPtr[cancelDB.ID]
	if cancelOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), cancelOld, cancelNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the cancelDB
	c.JSON(http.StatusOK, cancelDB)
}

// DeleteCancel
//
// swagger:route DELETE /cancels/{ID} cancels deleteCancel
//
// # Delete a cancel
//
// default: genericError
//
//	200: cancelDBResponse
func (controller *Controller) DeleteCancel(c *gin.Context) {

	mutexCancel.Lock()
	defer mutexCancel.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCancel", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCancel.GetDB()

	// Get model if exist
	var cancelDB orm.CancelDB
	if err := db.First(&cancelDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&cancelDB)

	// get an instance (not staged) from DB instance, and call callback function
	cancelDeleted := new(models.Cancel)
	cancelDB.CopyBasicFieldsToCancel(cancelDeleted)

	// get stage instance from DB instance, and call callback function
	cancelStaged := backRepo.BackRepoCancel.Map_CancelDBID_CancelPtr[cancelDB.ID]
	if cancelStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), cancelStaged, cancelDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
