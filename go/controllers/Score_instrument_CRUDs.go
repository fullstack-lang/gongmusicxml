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
var __Score_instrument__dummysDeclaration__ models.Score_instrument
var __Score_instrument_time__dummyDeclaration time.Duration

var mutexScore_instrument sync.Mutex

// An Score_instrumentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getScore_instrument updateScore_instrument deleteScore_instrument
type Score_instrumentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Score_instrumentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postScore_instrument updateScore_instrument
type Score_instrumentInput struct {
	// The Score_instrument to submit or modify
	// in: body
	Score_instrument *orm.Score_instrumentAPI
}

// GetScore_instruments
//
// swagger:route GET /score_instruments score_instruments getScore_instruments
//
// # Get all score_instruments
//
// Responses:
// default: genericError
//
//	200: score_instrumentDBResponse
func (controller *Controller) GetScore_instruments(c *gin.Context) {

	// source slice
	var score_instrumentDBs []orm.Score_instrumentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScore_instruments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_instrument.GetDB()

	query := db.Find(&score_instrumentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	score_instrumentAPIs := make([]orm.Score_instrumentAPI, 0)

	// for each score_instrument, update fields from the database nullable fields
	for idx := range score_instrumentDBs {
		score_instrumentDB := &score_instrumentDBs[idx]
		_ = score_instrumentDB
		var score_instrumentAPI orm.Score_instrumentAPI

		// insertion point for updating fields
		score_instrumentAPI.ID = score_instrumentDB.ID
		score_instrumentDB.CopyBasicFieldsToScore_instrument_WOP(&score_instrumentAPI.Score_instrument_WOP)
		score_instrumentAPI.Score_instrumentPointersEncoding = score_instrumentDB.Score_instrumentPointersEncoding
		score_instrumentAPIs = append(score_instrumentAPIs, score_instrumentAPI)
	}

	c.JSON(http.StatusOK, score_instrumentAPIs)
}

// PostScore_instrument
//
// swagger:route POST /score_instruments score_instruments postScore_instrument
//
// Creates a score_instrument
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostScore_instrument(c *gin.Context) {

	mutexScore_instrument.Lock()
	defer mutexScore_instrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostScore_instruments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_instrument.GetDB()

	// Validate input
	var input orm.Score_instrumentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create score_instrument
	score_instrumentDB := orm.Score_instrumentDB{}
	score_instrumentDB.Score_instrumentPointersEncoding = input.Score_instrumentPointersEncoding
	score_instrumentDB.CopyBasicFieldsFromScore_instrument_WOP(&input.Score_instrument_WOP)

	query := db.Create(&score_instrumentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoScore_instrument.CheckoutPhaseOneInstance(&score_instrumentDB)
	score_instrument := backRepo.BackRepoScore_instrument.Map_Score_instrumentDBID_Score_instrumentPtr[score_instrumentDB.ID]

	if score_instrument != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), score_instrument)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, score_instrumentDB)
}

// GetScore_instrument
//
// swagger:route GET /score_instruments/{ID} score_instruments getScore_instrument
//
// Gets the details for a score_instrument.
//
// Responses:
// default: genericError
//
//	200: score_instrumentDBResponse
func (controller *Controller) GetScore_instrument(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScore_instrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_instrument.GetDB()

	// Get score_instrumentDB in DB
	var score_instrumentDB orm.Score_instrumentDB
	if err := db.First(&score_instrumentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var score_instrumentAPI orm.Score_instrumentAPI
	score_instrumentAPI.ID = score_instrumentDB.ID
	score_instrumentAPI.Score_instrumentPointersEncoding = score_instrumentDB.Score_instrumentPointersEncoding
	score_instrumentDB.CopyBasicFieldsToScore_instrument_WOP(&score_instrumentAPI.Score_instrument_WOP)

	c.JSON(http.StatusOK, score_instrumentAPI)
}

// UpdateScore_instrument
//
// swagger:route PATCH /score_instruments/{ID} score_instruments updateScore_instrument
//
// # Update a score_instrument
//
// Responses:
// default: genericError
//
//	200: score_instrumentDBResponse
func (controller *Controller) UpdateScore_instrument(c *gin.Context) {

	mutexScore_instrument.Lock()
	defer mutexScore_instrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateScore_instrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_instrument.GetDB()

	// Validate input
	var input orm.Score_instrumentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var score_instrumentDB orm.Score_instrumentDB

	// fetch the score_instrument
	query := db.First(&score_instrumentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	score_instrumentDB.CopyBasicFieldsFromScore_instrument_WOP(&input.Score_instrument_WOP)
	score_instrumentDB.Score_instrumentPointersEncoding = input.Score_instrumentPointersEncoding

	query = db.Model(&score_instrumentDB).Updates(score_instrumentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	score_instrumentNew := new(models.Score_instrument)
	score_instrumentDB.CopyBasicFieldsToScore_instrument(score_instrumentNew)

	// redeem pointers
	score_instrumentDB.DecodePointers(backRepo, score_instrumentNew)

	// get stage instance from DB instance, and call callback function
	score_instrumentOld := backRepo.BackRepoScore_instrument.Map_Score_instrumentDBID_Score_instrumentPtr[score_instrumentDB.ID]
	if score_instrumentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), score_instrumentOld, score_instrumentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the score_instrumentDB
	c.JSON(http.StatusOK, score_instrumentDB)
}

// DeleteScore_instrument
//
// swagger:route DELETE /score_instruments/{ID} score_instruments deleteScore_instrument
//
// # Delete a score_instrument
//
// default: genericError
//
//	200: score_instrumentDBResponse
func (controller *Controller) DeleteScore_instrument(c *gin.Context) {

	mutexScore_instrument.Lock()
	defer mutexScore_instrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteScore_instrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_instrument.GetDB()

	// Get model if exist
	var score_instrumentDB orm.Score_instrumentDB
	if err := db.First(&score_instrumentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&score_instrumentDB)

	// get an instance (not staged) from DB instance, and call callback function
	score_instrumentDeleted := new(models.Score_instrument)
	score_instrumentDB.CopyBasicFieldsToScore_instrument(score_instrumentDeleted)

	// get stage instance from DB instance, and call callback function
	score_instrumentStaged := backRepo.BackRepoScore_instrument.Map_Score_instrumentDBID_Score_instrumentPtr[score_instrumentDB.ID]
	if score_instrumentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), score_instrumentStaged, score_instrumentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
