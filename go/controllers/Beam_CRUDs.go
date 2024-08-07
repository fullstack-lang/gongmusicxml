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
var __Beam__dummysDeclaration__ models.Beam
var __Beam_time__dummyDeclaration time.Duration

var mutexBeam sync.Mutex

// An BeamID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBeam updateBeam deleteBeam
type BeamID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BeamInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBeam updateBeam
type BeamInput struct {
	// The Beam to submit or modify
	// in: body
	Beam *orm.BeamAPI
}

// GetBeams
//
// swagger:route GET /beams beams getBeams
//
// # Get all beams
//
// Responses:
// default: genericError
//
//	200: beamDBResponse
func (controller *Controller) GetBeams(c *gin.Context) {

	// source slice
	var beamDBs []orm.BeamDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBeams", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeam.GetDB()

	query := db.Find(&beamDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	beamAPIs := make([]orm.BeamAPI, 0)

	// for each beam, update fields from the database nullable fields
	for idx := range beamDBs {
		beamDB := &beamDBs[idx]
		_ = beamDB
		var beamAPI orm.BeamAPI

		// insertion point for updating fields
		beamAPI.ID = beamDB.ID
		beamDB.CopyBasicFieldsToBeam_WOP(&beamAPI.Beam_WOP)
		beamAPI.BeamPointersEncoding = beamDB.BeamPointersEncoding
		beamAPIs = append(beamAPIs, beamAPI)
	}

	c.JSON(http.StatusOK, beamAPIs)
}

// PostBeam
//
// swagger:route POST /beams beams postBeam
//
// Creates a beam
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBeam(c *gin.Context) {

	mutexBeam.Lock()
	defer mutexBeam.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBeams", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeam.GetDB()

	// Validate input
	var input orm.BeamAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create beam
	beamDB := orm.BeamDB{}
	beamDB.BeamPointersEncoding = input.BeamPointersEncoding
	beamDB.CopyBasicFieldsFromBeam_WOP(&input.Beam_WOP)

	query := db.Create(&beamDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBeam.CheckoutPhaseOneInstance(&beamDB)
	beam := backRepo.BackRepoBeam.Map_BeamDBID_BeamPtr[beamDB.ID]

	if beam != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), beam)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, beamDB)
}

// GetBeam
//
// swagger:route GET /beams/{ID} beams getBeam
//
// Gets the details for a beam.
//
// Responses:
// default: genericError
//
//	200: beamDBResponse
func (controller *Controller) GetBeam(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBeam", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeam.GetDB()

	// Get beamDB in DB
	var beamDB orm.BeamDB
	if err := db.First(&beamDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var beamAPI orm.BeamAPI
	beamAPI.ID = beamDB.ID
	beamAPI.BeamPointersEncoding = beamDB.BeamPointersEncoding
	beamDB.CopyBasicFieldsToBeam_WOP(&beamAPI.Beam_WOP)

	c.JSON(http.StatusOK, beamAPI)
}

// UpdateBeam
//
// swagger:route PATCH /beams/{ID} beams updateBeam
//
// # Update a beam
//
// Responses:
// default: genericError
//
//	200: beamDBResponse
func (controller *Controller) UpdateBeam(c *gin.Context) {

	mutexBeam.Lock()
	defer mutexBeam.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBeam", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeam.GetDB()

	// Validate input
	var input orm.BeamAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var beamDB orm.BeamDB

	// fetch the beam
	query := db.First(&beamDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	beamDB.CopyBasicFieldsFromBeam_WOP(&input.Beam_WOP)
	beamDB.BeamPointersEncoding = input.BeamPointersEncoding

	query = db.Model(&beamDB).Updates(beamDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	beamNew := new(models.Beam)
	beamDB.CopyBasicFieldsToBeam(beamNew)

	// redeem pointers
	beamDB.DecodePointers(backRepo, beamNew)

	// get stage instance from DB instance, and call callback function
	beamOld := backRepo.BackRepoBeam.Map_BeamDBID_BeamPtr[beamDB.ID]
	if beamOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), beamOld, beamNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the beamDB
	c.JSON(http.StatusOK, beamDB)
}

// DeleteBeam
//
// swagger:route DELETE /beams/{ID} beams deleteBeam
//
// # Delete a beam
//
// default: genericError
//
//	200: beamDBResponse
func (controller *Controller) DeleteBeam(c *gin.Context) {

	mutexBeam.Lock()
	defer mutexBeam.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBeam", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeam.GetDB()

	// Get model if exist
	var beamDB orm.BeamDB
	if err := db.First(&beamDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&beamDB)

	// get an instance (not staged) from DB instance, and call callback function
	beamDeleted := new(models.Beam)
	beamDB.CopyBasicFieldsToBeam(beamDeleted)

	// get stage instance from DB instance, and call callback function
	beamStaged := backRepo.BackRepoBeam.Map_BeamDBID_BeamPtr[beamDB.ID]
	if beamStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), beamStaged, beamDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
