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
var __Beat_unit_tied__dummysDeclaration__ models.Beat_unit_tied
var __Beat_unit_tied_time__dummyDeclaration time.Duration

var mutexBeat_unit_tied sync.Mutex

// An Beat_unit_tiedID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBeat_unit_tied updateBeat_unit_tied deleteBeat_unit_tied
type Beat_unit_tiedID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Beat_unit_tiedInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBeat_unit_tied updateBeat_unit_tied
type Beat_unit_tiedInput struct {
	// The Beat_unit_tied to submit or modify
	// in: body
	Beat_unit_tied *orm.Beat_unit_tiedAPI
}

// GetBeat_unit_tieds
//
// swagger:route GET /beat_unit_tieds beat_unit_tieds getBeat_unit_tieds
//
// # Get all beat_unit_tieds
//
// Responses:
// default: genericError
//
//	200: beat_unit_tiedDBResponse
func (controller *Controller) GetBeat_unit_tieds(c *gin.Context) {

	// source slice
	var beat_unit_tiedDBs []orm.Beat_unit_tiedDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBeat_unit_tieds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_unit_tied.GetDB()

	query := db.Find(&beat_unit_tiedDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	beat_unit_tiedAPIs := make([]orm.Beat_unit_tiedAPI, 0)

	// for each beat_unit_tied, update fields from the database nullable fields
	for idx := range beat_unit_tiedDBs {
		beat_unit_tiedDB := &beat_unit_tiedDBs[idx]
		_ = beat_unit_tiedDB
		var beat_unit_tiedAPI orm.Beat_unit_tiedAPI

		// insertion point for updating fields
		beat_unit_tiedAPI.ID = beat_unit_tiedDB.ID
		beat_unit_tiedDB.CopyBasicFieldsToBeat_unit_tied_WOP(&beat_unit_tiedAPI.Beat_unit_tied_WOP)
		beat_unit_tiedAPI.Beat_unit_tiedPointersEncoding = beat_unit_tiedDB.Beat_unit_tiedPointersEncoding
		beat_unit_tiedAPIs = append(beat_unit_tiedAPIs, beat_unit_tiedAPI)
	}

	c.JSON(http.StatusOK, beat_unit_tiedAPIs)
}

// PostBeat_unit_tied
//
// swagger:route POST /beat_unit_tieds beat_unit_tieds postBeat_unit_tied
//
// Creates a beat_unit_tied
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBeat_unit_tied(c *gin.Context) {

	mutexBeat_unit_tied.Lock()
	defer mutexBeat_unit_tied.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBeat_unit_tieds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_unit_tied.GetDB()

	// Validate input
	var input orm.Beat_unit_tiedAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create beat_unit_tied
	beat_unit_tiedDB := orm.Beat_unit_tiedDB{}
	beat_unit_tiedDB.Beat_unit_tiedPointersEncoding = input.Beat_unit_tiedPointersEncoding
	beat_unit_tiedDB.CopyBasicFieldsFromBeat_unit_tied_WOP(&input.Beat_unit_tied_WOP)

	query := db.Create(&beat_unit_tiedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBeat_unit_tied.CheckoutPhaseOneInstance(&beat_unit_tiedDB)
	beat_unit_tied := backRepo.BackRepoBeat_unit_tied.Map_Beat_unit_tiedDBID_Beat_unit_tiedPtr[beat_unit_tiedDB.ID]

	if beat_unit_tied != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), beat_unit_tied)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, beat_unit_tiedDB)
}

// GetBeat_unit_tied
//
// swagger:route GET /beat_unit_tieds/{ID} beat_unit_tieds getBeat_unit_tied
//
// Gets the details for a beat_unit_tied.
//
// Responses:
// default: genericError
//
//	200: beat_unit_tiedDBResponse
func (controller *Controller) GetBeat_unit_tied(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBeat_unit_tied", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_unit_tied.GetDB()

	// Get beat_unit_tiedDB in DB
	var beat_unit_tiedDB orm.Beat_unit_tiedDB
	if err := db.First(&beat_unit_tiedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var beat_unit_tiedAPI orm.Beat_unit_tiedAPI
	beat_unit_tiedAPI.ID = beat_unit_tiedDB.ID
	beat_unit_tiedAPI.Beat_unit_tiedPointersEncoding = beat_unit_tiedDB.Beat_unit_tiedPointersEncoding
	beat_unit_tiedDB.CopyBasicFieldsToBeat_unit_tied_WOP(&beat_unit_tiedAPI.Beat_unit_tied_WOP)

	c.JSON(http.StatusOK, beat_unit_tiedAPI)
}

// UpdateBeat_unit_tied
//
// swagger:route PATCH /beat_unit_tieds/{ID} beat_unit_tieds updateBeat_unit_tied
//
// # Update a beat_unit_tied
//
// Responses:
// default: genericError
//
//	200: beat_unit_tiedDBResponse
func (controller *Controller) UpdateBeat_unit_tied(c *gin.Context) {

	mutexBeat_unit_tied.Lock()
	defer mutexBeat_unit_tied.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBeat_unit_tied", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_unit_tied.GetDB()

	// Validate input
	var input orm.Beat_unit_tiedAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var beat_unit_tiedDB orm.Beat_unit_tiedDB

	// fetch the beat_unit_tied
	query := db.First(&beat_unit_tiedDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	beat_unit_tiedDB.CopyBasicFieldsFromBeat_unit_tied_WOP(&input.Beat_unit_tied_WOP)
	beat_unit_tiedDB.Beat_unit_tiedPointersEncoding = input.Beat_unit_tiedPointersEncoding

	query = db.Model(&beat_unit_tiedDB).Updates(beat_unit_tiedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	beat_unit_tiedNew := new(models.Beat_unit_tied)
	beat_unit_tiedDB.CopyBasicFieldsToBeat_unit_tied(beat_unit_tiedNew)

	// redeem pointers
	beat_unit_tiedDB.DecodePointers(backRepo, beat_unit_tiedNew)

	// get stage instance from DB instance, and call callback function
	beat_unit_tiedOld := backRepo.BackRepoBeat_unit_tied.Map_Beat_unit_tiedDBID_Beat_unit_tiedPtr[beat_unit_tiedDB.ID]
	if beat_unit_tiedOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), beat_unit_tiedOld, beat_unit_tiedNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the beat_unit_tiedDB
	c.JSON(http.StatusOK, beat_unit_tiedDB)
}

// DeleteBeat_unit_tied
//
// swagger:route DELETE /beat_unit_tieds/{ID} beat_unit_tieds deleteBeat_unit_tied
//
// # Delete a beat_unit_tied
//
// default: genericError
//
//	200: beat_unit_tiedDBResponse
func (controller *Controller) DeleteBeat_unit_tied(c *gin.Context) {

	mutexBeat_unit_tied.Lock()
	defer mutexBeat_unit_tied.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBeat_unit_tied", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_unit_tied.GetDB()

	// Get model if exist
	var beat_unit_tiedDB orm.Beat_unit_tiedDB
	if err := db.First(&beat_unit_tiedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&beat_unit_tiedDB)

	// get an instance (not staged) from DB instance, and call callback function
	beat_unit_tiedDeleted := new(models.Beat_unit_tied)
	beat_unit_tiedDB.CopyBasicFieldsToBeat_unit_tied(beat_unit_tiedDeleted)

	// get stage instance from DB instance, and call callback function
	beat_unit_tiedStaged := backRepo.BackRepoBeat_unit_tied.Map_Beat_unit_tiedDBID_Beat_unit_tiedPtr[beat_unit_tiedDB.ID]
	if beat_unit_tiedStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), beat_unit_tiedStaged, beat_unit_tiedDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
