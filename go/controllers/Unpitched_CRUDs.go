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
var __Unpitched__dummysDeclaration__ models.Unpitched
var __Unpitched_time__dummyDeclaration time.Duration

var mutexUnpitched sync.Mutex

// An UnpitchedID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUnpitched updateUnpitched deleteUnpitched
type UnpitchedID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UnpitchedInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUnpitched updateUnpitched
type UnpitchedInput struct {
	// The Unpitched to submit or modify
	// in: body
	Unpitched *orm.UnpitchedAPI
}

// GetUnpitcheds
//
// swagger:route GET /unpitcheds unpitcheds getUnpitcheds
//
// # Get all unpitcheds
//
// Responses:
// default: genericError
//
//	200: unpitchedDBResponse
func (controller *Controller) GetUnpitcheds(c *gin.Context) {

	// source slice
	var unpitchedDBs []orm.UnpitchedDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUnpitcheds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUnpitched.GetDB()

	query := db.Find(&unpitchedDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	unpitchedAPIs := make([]orm.UnpitchedAPI, 0)

	// for each unpitched, update fields from the database nullable fields
	for idx := range unpitchedDBs {
		unpitchedDB := &unpitchedDBs[idx]
		_ = unpitchedDB
		var unpitchedAPI orm.UnpitchedAPI

		// insertion point for updating fields
		unpitchedAPI.ID = unpitchedDB.ID
		unpitchedDB.CopyBasicFieldsToUnpitched_WOP(&unpitchedAPI.Unpitched_WOP)
		unpitchedAPI.UnpitchedPointersEncoding = unpitchedDB.UnpitchedPointersEncoding
		unpitchedAPIs = append(unpitchedAPIs, unpitchedAPI)
	}

	c.JSON(http.StatusOK, unpitchedAPIs)
}

// PostUnpitched
//
// swagger:route POST /unpitcheds unpitcheds postUnpitched
//
// Creates a unpitched
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostUnpitched(c *gin.Context) {

	mutexUnpitched.Lock()
	defer mutexUnpitched.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostUnpitcheds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUnpitched.GetDB()

	// Validate input
	var input orm.UnpitchedAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create unpitched
	unpitchedDB := orm.UnpitchedDB{}
	unpitchedDB.UnpitchedPointersEncoding = input.UnpitchedPointersEncoding
	unpitchedDB.CopyBasicFieldsFromUnpitched_WOP(&input.Unpitched_WOP)

	query := db.Create(&unpitchedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoUnpitched.CheckoutPhaseOneInstance(&unpitchedDB)
	unpitched := backRepo.BackRepoUnpitched.Map_UnpitchedDBID_UnpitchedPtr[unpitchedDB.ID]

	if unpitched != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), unpitched)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, unpitchedDB)
}

// GetUnpitched
//
// swagger:route GET /unpitcheds/{ID} unpitcheds getUnpitched
//
// Gets the details for a unpitched.
//
// Responses:
// default: genericError
//
//	200: unpitchedDBResponse
func (controller *Controller) GetUnpitched(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUnpitched", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUnpitched.GetDB()

	// Get unpitchedDB in DB
	var unpitchedDB orm.UnpitchedDB
	if err := db.First(&unpitchedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var unpitchedAPI orm.UnpitchedAPI
	unpitchedAPI.ID = unpitchedDB.ID
	unpitchedAPI.UnpitchedPointersEncoding = unpitchedDB.UnpitchedPointersEncoding
	unpitchedDB.CopyBasicFieldsToUnpitched_WOP(&unpitchedAPI.Unpitched_WOP)

	c.JSON(http.StatusOK, unpitchedAPI)
}

// UpdateUnpitched
//
// swagger:route PATCH /unpitcheds/{ID} unpitcheds updateUnpitched
//
// # Update a unpitched
//
// Responses:
// default: genericError
//
//	200: unpitchedDBResponse
func (controller *Controller) UpdateUnpitched(c *gin.Context) {

	mutexUnpitched.Lock()
	defer mutexUnpitched.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateUnpitched", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUnpitched.GetDB()

	// Validate input
	var input orm.UnpitchedAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var unpitchedDB orm.UnpitchedDB

	// fetch the unpitched
	query := db.First(&unpitchedDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	unpitchedDB.CopyBasicFieldsFromUnpitched_WOP(&input.Unpitched_WOP)
	unpitchedDB.UnpitchedPointersEncoding = input.UnpitchedPointersEncoding

	query = db.Model(&unpitchedDB).Updates(unpitchedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	unpitchedNew := new(models.Unpitched)
	unpitchedDB.CopyBasicFieldsToUnpitched(unpitchedNew)

	// redeem pointers
	unpitchedDB.DecodePointers(backRepo, unpitchedNew)

	// get stage instance from DB instance, and call callback function
	unpitchedOld := backRepo.BackRepoUnpitched.Map_UnpitchedDBID_UnpitchedPtr[unpitchedDB.ID]
	if unpitchedOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), unpitchedOld, unpitchedNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the unpitchedDB
	c.JSON(http.StatusOK, unpitchedDB)
}

// DeleteUnpitched
//
// swagger:route DELETE /unpitcheds/{ID} unpitcheds deleteUnpitched
//
// # Delete a unpitched
//
// default: genericError
//
//	200: unpitchedDBResponse
func (controller *Controller) DeleteUnpitched(c *gin.Context) {

	mutexUnpitched.Lock()
	defer mutexUnpitched.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteUnpitched", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUnpitched.GetDB()

	// Get model if exist
	var unpitchedDB orm.UnpitchedDB
	if err := db.First(&unpitchedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&unpitchedDB)

	// get an instance (not staged) from DB instance, and call callback function
	unpitchedDeleted := new(models.Unpitched)
	unpitchedDB.CopyBasicFieldsToUnpitched(unpitchedDeleted)

	// get stage instance from DB instance, and call callback function
	unpitchedStaged := backRepo.BackRepoUnpitched.Map_UnpitchedDBID_UnpitchedPtr[unpitchedDB.ID]
	if unpitchedStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), unpitchedStaged, unpitchedDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
