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
var __Inversion__dummysDeclaration__ models.Inversion
var __Inversion_time__dummyDeclaration time.Duration

var mutexInversion sync.Mutex

// An InversionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getInversion updateInversion deleteInversion
type InversionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// InversionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postInversion updateInversion
type InversionInput struct {
	// The Inversion to submit or modify
	// in: body
	Inversion *orm.InversionAPI
}

// GetInversions
//
// swagger:route GET /inversions inversions getInversions
//
// # Get all inversions
//
// Responses:
// default: genericError
//
//	200: inversionDBResponse
func (controller *Controller) GetInversions(c *gin.Context) {

	// source slice
	var inversionDBs []orm.InversionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInversions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInversion.GetDB()

	query := db.Find(&inversionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	inversionAPIs := make([]orm.InversionAPI, 0)

	// for each inversion, update fields from the database nullable fields
	for idx := range inversionDBs {
		inversionDB := &inversionDBs[idx]
		_ = inversionDB
		var inversionAPI orm.InversionAPI

		// insertion point for updating fields
		inversionAPI.ID = inversionDB.ID
		inversionDB.CopyBasicFieldsToInversion_WOP(&inversionAPI.Inversion_WOP)
		inversionAPI.InversionPointersEncoding = inversionDB.InversionPointersEncoding
		inversionAPIs = append(inversionAPIs, inversionAPI)
	}

	c.JSON(http.StatusOK, inversionAPIs)
}

// PostInversion
//
// swagger:route POST /inversions inversions postInversion
//
// Creates a inversion
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostInversion(c *gin.Context) {

	mutexInversion.Lock()
	defer mutexInversion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostInversions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInversion.GetDB()

	// Validate input
	var input orm.InversionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create inversion
	inversionDB := orm.InversionDB{}
	inversionDB.InversionPointersEncoding = input.InversionPointersEncoding
	inversionDB.CopyBasicFieldsFromInversion_WOP(&input.Inversion_WOP)

	query := db.Create(&inversionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoInversion.CheckoutPhaseOneInstance(&inversionDB)
	inversion := backRepo.BackRepoInversion.Map_InversionDBID_InversionPtr[inversionDB.ID]

	if inversion != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), inversion)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, inversionDB)
}

// GetInversion
//
// swagger:route GET /inversions/{ID} inversions getInversion
//
// Gets the details for a inversion.
//
// Responses:
// default: genericError
//
//	200: inversionDBResponse
func (controller *Controller) GetInversion(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInversion", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInversion.GetDB()

	// Get inversionDB in DB
	var inversionDB orm.InversionDB
	if err := db.First(&inversionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var inversionAPI orm.InversionAPI
	inversionAPI.ID = inversionDB.ID
	inversionAPI.InversionPointersEncoding = inversionDB.InversionPointersEncoding
	inversionDB.CopyBasicFieldsToInversion_WOP(&inversionAPI.Inversion_WOP)

	c.JSON(http.StatusOK, inversionAPI)
}

// UpdateInversion
//
// swagger:route PATCH /inversions/{ID} inversions updateInversion
//
// # Update a inversion
//
// Responses:
// default: genericError
//
//	200: inversionDBResponse
func (controller *Controller) UpdateInversion(c *gin.Context) {

	mutexInversion.Lock()
	defer mutexInversion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateInversion", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInversion.GetDB()

	// Validate input
	var input orm.InversionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var inversionDB orm.InversionDB

	// fetch the inversion
	query := db.First(&inversionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	inversionDB.CopyBasicFieldsFromInversion_WOP(&input.Inversion_WOP)
	inversionDB.InversionPointersEncoding = input.InversionPointersEncoding

	query = db.Model(&inversionDB).Updates(inversionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	inversionNew := new(models.Inversion)
	inversionDB.CopyBasicFieldsToInversion(inversionNew)

	// redeem pointers
	inversionDB.DecodePointers(backRepo, inversionNew)

	// get stage instance from DB instance, and call callback function
	inversionOld := backRepo.BackRepoInversion.Map_InversionDBID_InversionPtr[inversionDB.ID]
	if inversionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), inversionOld, inversionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the inversionDB
	c.JSON(http.StatusOK, inversionDB)
}

// DeleteInversion
//
// swagger:route DELETE /inversions/{ID} inversions deleteInversion
//
// # Delete a inversion
//
// default: genericError
//
//	200: inversionDBResponse
func (controller *Controller) DeleteInversion(c *gin.Context) {

	mutexInversion.Lock()
	defer mutexInversion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteInversion", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInversion.GetDB()

	// Get model if exist
	var inversionDB orm.InversionDB
	if err := db.First(&inversionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&inversionDB)

	// get an instance (not staged) from DB instance, and call callback function
	inversionDeleted := new(models.Inversion)
	inversionDB.CopyBasicFieldsToInversion(inversionDeleted)

	// get stage instance from DB instance, and call callback function
	inversionStaged := backRepo.BackRepoInversion.Map_InversionDBID_InversionPtr[inversionDB.ID]
	if inversionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), inversionStaged, inversionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
