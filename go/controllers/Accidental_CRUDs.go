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
var __Accidental__dummysDeclaration__ models.Accidental
var __Accidental_time__dummyDeclaration time.Duration

var mutexAccidental sync.Mutex

// An AccidentalID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAccidental updateAccidental deleteAccidental
type AccidentalID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AccidentalInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAccidental updateAccidental
type AccidentalInput struct {
	// The Accidental to submit or modify
	// in: body
	Accidental *orm.AccidentalAPI
}

// GetAccidentals
//
// swagger:route GET /accidentals accidentals getAccidentals
//
// # Get all accidentals
//
// Responses:
// default: genericError
//
//	200: accidentalDBResponse
func (controller *Controller) GetAccidentals(c *gin.Context) {

	// source slice
	var accidentalDBs []orm.AccidentalDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccidentals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental.GetDB()

	query := db.Find(&accidentalDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	accidentalAPIs := make([]orm.AccidentalAPI, 0)

	// for each accidental, update fields from the database nullable fields
	for idx := range accidentalDBs {
		accidentalDB := &accidentalDBs[idx]
		_ = accidentalDB
		var accidentalAPI orm.AccidentalAPI

		// insertion point for updating fields
		accidentalAPI.ID = accidentalDB.ID
		accidentalDB.CopyBasicFieldsToAccidental_WOP(&accidentalAPI.Accidental_WOP)
		accidentalAPI.AccidentalPointersEncoding = accidentalDB.AccidentalPointersEncoding
		accidentalAPIs = append(accidentalAPIs, accidentalAPI)
	}

	c.JSON(http.StatusOK, accidentalAPIs)
}

// PostAccidental
//
// swagger:route POST /accidentals accidentals postAccidental
//
// Creates a accidental
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAccidental(c *gin.Context) {

	mutexAccidental.Lock()
	defer mutexAccidental.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAccidentals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental.GetDB()

	// Validate input
	var input orm.AccidentalAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create accidental
	accidentalDB := orm.AccidentalDB{}
	accidentalDB.AccidentalPointersEncoding = input.AccidentalPointersEncoding
	accidentalDB.CopyBasicFieldsFromAccidental_WOP(&input.Accidental_WOP)

	query := db.Create(&accidentalDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAccidental.CheckoutPhaseOneInstance(&accidentalDB)
	accidental := backRepo.BackRepoAccidental.Map_AccidentalDBID_AccidentalPtr[accidentalDB.ID]

	if accidental != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), accidental)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, accidentalDB)
}

// GetAccidental
//
// swagger:route GET /accidentals/{ID} accidentals getAccidental
//
// Gets the details for a accidental.
//
// Responses:
// default: genericError
//
//	200: accidentalDBResponse
func (controller *Controller) GetAccidental(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccidental", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental.GetDB()

	// Get accidentalDB in DB
	var accidentalDB orm.AccidentalDB
	if err := db.First(&accidentalDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var accidentalAPI orm.AccidentalAPI
	accidentalAPI.ID = accidentalDB.ID
	accidentalAPI.AccidentalPointersEncoding = accidentalDB.AccidentalPointersEncoding
	accidentalDB.CopyBasicFieldsToAccidental_WOP(&accidentalAPI.Accidental_WOP)

	c.JSON(http.StatusOK, accidentalAPI)
}

// UpdateAccidental
//
// swagger:route PATCH /accidentals/{ID} accidentals updateAccidental
//
// # Update a accidental
//
// Responses:
// default: genericError
//
//	200: accidentalDBResponse
func (controller *Controller) UpdateAccidental(c *gin.Context) {

	mutexAccidental.Lock()
	defer mutexAccidental.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAccidental", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental.GetDB()

	// Validate input
	var input orm.AccidentalAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var accidentalDB orm.AccidentalDB

	// fetch the accidental
	query := db.First(&accidentalDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	accidentalDB.CopyBasicFieldsFromAccidental_WOP(&input.Accidental_WOP)
	accidentalDB.AccidentalPointersEncoding = input.AccidentalPointersEncoding

	query = db.Model(&accidentalDB).Updates(accidentalDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	accidentalNew := new(models.Accidental)
	accidentalDB.CopyBasicFieldsToAccidental(accidentalNew)

	// redeem pointers
	accidentalDB.DecodePointers(backRepo, accidentalNew)

	// get stage instance from DB instance, and call callback function
	accidentalOld := backRepo.BackRepoAccidental.Map_AccidentalDBID_AccidentalPtr[accidentalDB.ID]
	if accidentalOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), accidentalOld, accidentalNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the accidentalDB
	c.JSON(http.StatusOK, accidentalDB)
}

// DeleteAccidental
//
// swagger:route DELETE /accidentals/{ID} accidentals deleteAccidental
//
// # Delete a accidental
//
// default: genericError
//
//	200: accidentalDBResponse
func (controller *Controller) DeleteAccidental(c *gin.Context) {

	mutexAccidental.Lock()
	defer mutexAccidental.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAccidental", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental.GetDB()

	// Get model if exist
	var accidentalDB orm.AccidentalDB
	if err := db.First(&accidentalDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&accidentalDB)

	// get an instance (not staged) from DB instance, and call callback function
	accidentalDeleted := new(models.Accidental)
	accidentalDB.CopyBasicFieldsToAccidental(accidentalDeleted)

	// get stage instance from DB instance, and call callback function
	accidentalStaged := backRepo.BackRepoAccidental.Map_AccidentalDBID_AccidentalPtr[accidentalDB.ID]
	if accidentalStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), accidentalStaged, accidentalDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
