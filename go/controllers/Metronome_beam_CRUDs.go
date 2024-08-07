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
var __Metronome_beam__dummysDeclaration__ models.Metronome_beam
var __Metronome_beam_time__dummyDeclaration time.Duration

var mutexMetronome_beam sync.Mutex

// An Metronome_beamID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMetronome_beam updateMetronome_beam deleteMetronome_beam
type Metronome_beamID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Metronome_beamInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMetronome_beam updateMetronome_beam
type Metronome_beamInput struct {
	// The Metronome_beam to submit or modify
	// in: body
	Metronome_beam *orm.Metronome_beamAPI
}

// GetMetronome_beams
//
// swagger:route GET /metronome_beams metronome_beams getMetronome_beams
//
// # Get all metronome_beams
//
// Responses:
// default: genericError
//
//	200: metronome_beamDBResponse
func (controller *Controller) GetMetronome_beams(c *gin.Context) {

	// source slice
	var metronome_beamDBs []orm.Metronome_beamDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronome_beams", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_beam.GetDB()

	query := db.Find(&metronome_beamDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	metronome_beamAPIs := make([]orm.Metronome_beamAPI, 0)

	// for each metronome_beam, update fields from the database nullable fields
	for idx := range metronome_beamDBs {
		metronome_beamDB := &metronome_beamDBs[idx]
		_ = metronome_beamDB
		var metronome_beamAPI orm.Metronome_beamAPI

		// insertion point for updating fields
		metronome_beamAPI.ID = metronome_beamDB.ID
		metronome_beamDB.CopyBasicFieldsToMetronome_beam_WOP(&metronome_beamAPI.Metronome_beam_WOP)
		metronome_beamAPI.Metronome_beamPointersEncoding = metronome_beamDB.Metronome_beamPointersEncoding
		metronome_beamAPIs = append(metronome_beamAPIs, metronome_beamAPI)
	}

	c.JSON(http.StatusOK, metronome_beamAPIs)
}

// PostMetronome_beam
//
// swagger:route POST /metronome_beams metronome_beams postMetronome_beam
//
// Creates a metronome_beam
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMetronome_beam(c *gin.Context) {

	mutexMetronome_beam.Lock()
	defer mutexMetronome_beam.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMetronome_beams", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_beam.GetDB()

	// Validate input
	var input orm.Metronome_beamAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create metronome_beam
	metronome_beamDB := orm.Metronome_beamDB{}
	metronome_beamDB.Metronome_beamPointersEncoding = input.Metronome_beamPointersEncoding
	metronome_beamDB.CopyBasicFieldsFromMetronome_beam_WOP(&input.Metronome_beam_WOP)

	query := db.Create(&metronome_beamDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMetronome_beam.CheckoutPhaseOneInstance(&metronome_beamDB)
	metronome_beam := backRepo.BackRepoMetronome_beam.Map_Metronome_beamDBID_Metronome_beamPtr[metronome_beamDB.ID]

	if metronome_beam != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), metronome_beam)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, metronome_beamDB)
}

// GetMetronome_beam
//
// swagger:route GET /metronome_beams/{ID} metronome_beams getMetronome_beam
//
// Gets the details for a metronome_beam.
//
// Responses:
// default: genericError
//
//	200: metronome_beamDBResponse
func (controller *Controller) GetMetronome_beam(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronome_beam", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_beam.GetDB()

	// Get metronome_beamDB in DB
	var metronome_beamDB orm.Metronome_beamDB
	if err := db.First(&metronome_beamDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var metronome_beamAPI orm.Metronome_beamAPI
	metronome_beamAPI.ID = metronome_beamDB.ID
	metronome_beamAPI.Metronome_beamPointersEncoding = metronome_beamDB.Metronome_beamPointersEncoding
	metronome_beamDB.CopyBasicFieldsToMetronome_beam_WOP(&metronome_beamAPI.Metronome_beam_WOP)

	c.JSON(http.StatusOK, metronome_beamAPI)
}

// UpdateMetronome_beam
//
// swagger:route PATCH /metronome_beams/{ID} metronome_beams updateMetronome_beam
//
// # Update a metronome_beam
//
// Responses:
// default: genericError
//
//	200: metronome_beamDBResponse
func (controller *Controller) UpdateMetronome_beam(c *gin.Context) {

	mutexMetronome_beam.Lock()
	defer mutexMetronome_beam.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMetronome_beam", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_beam.GetDB()

	// Validate input
	var input orm.Metronome_beamAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var metronome_beamDB orm.Metronome_beamDB

	// fetch the metronome_beam
	query := db.First(&metronome_beamDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	metronome_beamDB.CopyBasicFieldsFromMetronome_beam_WOP(&input.Metronome_beam_WOP)
	metronome_beamDB.Metronome_beamPointersEncoding = input.Metronome_beamPointersEncoding

	query = db.Model(&metronome_beamDB).Updates(metronome_beamDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	metronome_beamNew := new(models.Metronome_beam)
	metronome_beamDB.CopyBasicFieldsToMetronome_beam(metronome_beamNew)

	// redeem pointers
	metronome_beamDB.DecodePointers(backRepo, metronome_beamNew)

	// get stage instance from DB instance, and call callback function
	metronome_beamOld := backRepo.BackRepoMetronome_beam.Map_Metronome_beamDBID_Metronome_beamPtr[metronome_beamDB.ID]
	if metronome_beamOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), metronome_beamOld, metronome_beamNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the metronome_beamDB
	c.JSON(http.StatusOK, metronome_beamDB)
}

// DeleteMetronome_beam
//
// swagger:route DELETE /metronome_beams/{ID} metronome_beams deleteMetronome_beam
//
// # Delete a metronome_beam
//
// default: genericError
//
//	200: metronome_beamDBResponse
func (controller *Controller) DeleteMetronome_beam(c *gin.Context) {

	mutexMetronome_beam.Lock()
	defer mutexMetronome_beam.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMetronome_beam", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_beam.GetDB()

	// Get model if exist
	var metronome_beamDB orm.Metronome_beamDB
	if err := db.First(&metronome_beamDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&metronome_beamDB)

	// get an instance (not staged) from DB instance, and call callback function
	metronome_beamDeleted := new(models.Metronome_beam)
	metronome_beamDB.CopyBasicFieldsToMetronome_beam(metronome_beamDeleted)

	// get stage instance from DB instance, and call callback function
	metronome_beamStaged := backRepo.BackRepoMetronome_beam.Map_Metronome_beamDBID_Metronome_beamPtr[metronome_beamDB.ID]
	if metronome_beamStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), metronome_beamStaged, metronome_beamDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
