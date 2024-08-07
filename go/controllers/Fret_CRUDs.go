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
var __Fret__dummysDeclaration__ models.Fret
var __Fret_time__dummyDeclaration time.Duration

var mutexFret sync.Mutex

// An FretID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFret updateFret deleteFret
type FretID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FretInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFret updateFret
type FretInput struct {
	// The Fret to submit or modify
	// in: body
	Fret *orm.FretAPI
}

// GetFrets
//
// swagger:route GET /frets frets getFrets
//
// # Get all frets
//
// Responses:
// default: genericError
//
//	200: fretDBResponse
func (controller *Controller) GetFrets(c *gin.Context) {

	// source slice
	var fretDBs []orm.FretDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFrets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFret.GetDB()

	query := db.Find(&fretDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	fretAPIs := make([]orm.FretAPI, 0)

	// for each fret, update fields from the database nullable fields
	for idx := range fretDBs {
		fretDB := &fretDBs[idx]
		_ = fretDB
		var fretAPI orm.FretAPI

		// insertion point for updating fields
		fretAPI.ID = fretDB.ID
		fretDB.CopyBasicFieldsToFret_WOP(&fretAPI.Fret_WOP)
		fretAPI.FretPointersEncoding = fretDB.FretPointersEncoding
		fretAPIs = append(fretAPIs, fretAPI)
	}

	c.JSON(http.StatusOK, fretAPIs)
}

// PostFret
//
// swagger:route POST /frets frets postFret
//
// Creates a fret
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFret(c *gin.Context) {

	mutexFret.Lock()
	defer mutexFret.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFrets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFret.GetDB()

	// Validate input
	var input orm.FretAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create fret
	fretDB := orm.FretDB{}
	fretDB.FretPointersEncoding = input.FretPointersEncoding
	fretDB.CopyBasicFieldsFromFret_WOP(&input.Fret_WOP)

	query := db.Create(&fretDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFret.CheckoutPhaseOneInstance(&fretDB)
	fret := backRepo.BackRepoFret.Map_FretDBID_FretPtr[fretDB.ID]

	if fret != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), fret)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, fretDB)
}

// GetFret
//
// swagger:route GET /frets/{ID} frets getFret
//
// Gets the details for a fret.
//
// Responses:
// default: genericError
//
//	200: fretDBResponse
func (controller *Controller) GetFret(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFret", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFret.GetDB()

	// Get fretDB in DB
	var fretDB orm.FretDB
	if err := db.First(&fretDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var fretAPI orm.FretAPI
	fretAPI.ID = fretDB.ID
	fretAPI.FretPointersEncoding = fretDB.FretPointersEncoding
	fretDB.CopyBasicFieldsToFret_WOP(&fretAPI.Fret_WOP)

	c.JSON(http.StatusOK, fretAPI)
}

// UpdateFret
//
// swagger:route PATCH /frets/{ID} frets updateFret
//
// # Update a fret
//
// Responses:
// default: genericError
//
//	200: fretDBResponse
func (controller *Controller) UpdateFret(c *gin.Context) {

	mutexFret.Lock()
	defer mutexFret.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFret", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFret.GetDB()

	// Validate input
	var input orm.FretAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var fretDB orm.FretDB

	// fetch the fret
	query := db.First(&fretDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	fretDB.CopyBasicFieldsFromFret_WOP(&input.Fret_WOP)
	fretDB.FretPointersEncoding = input.FretPointersEncoding

	query = db.Model(&fretDB).Updates(fretDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	fretNew := new(models.Fret)
	fretDB.CopyBasicFieldsToFret(fretNew)

	// redeem pointers
	fretDB.DecodePointers(backRepo, fretNew)

	// get stage instance from DB instance, and call callback function
	fretOld := backRepo.BackRepoFret.Map_FretDBID_FretPtr[fretDB.ID]
	if fretOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), fretOld, fretNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the fretDB
	c.JSON(http.StatusOK, fretDB)
}

// DeleteFret
//
// swagger:route DELETE /frets/{ID} frets deleteFret
//
// # Delete a fret
//
// default: genericError
//
//	200: fretDBResponse
func (controller *Controller) DeleteFret(c *gin.Context) {

	mutexFret.Lock()
	defer mutexFret.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFret", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFret.GetDB()

	// Get model if exist
	var fretDB orm.FretDB
	if err := db.First(&fretDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&fretDB)

	// get an instance (not staged) from DB instance, and call callback function
	fretDeleted := new(models.Fret)
	fretDB.CopyBasicFieldsToFret(fretDeleted)

	// get stage instance from DB instance, and call callback function
	fretStaged := backRepo.BackRepoFret.Map_FretDBID_FretPtr[fretDB.ID]
	if fretStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), fretStaged, fretDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
