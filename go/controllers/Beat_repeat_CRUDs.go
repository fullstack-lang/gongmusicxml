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
var __Beat_repeat__dummysDeclaration__ models.Beat_repeat
var __Beat_repeat_time__dummyDeclaration time.Duration

var mutexBeat_repeat sync.Mutex

// An Beat_repeatID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBeat_repeat updateBeat_repeat deleteBeat_repeat
type Beat_repeatID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Beat_repeatInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBeat_repeat updateBeat_repeat
type Beat_repeatInput struct {
	// The Beat_repeat to submit or modify
	// in: body
	Beat_repeat *orm.Beat_repeatAPI
}

// GetBeat_repeats
//
// swagger:route GET /beat_repeats beat_repeats getBeat_repeats
//
// # Get all beat_repeats
//
// Responses:
// default: genericError
//
//	200: beat_repeatDBResponse
func (controller *Controller) GetBeat_repeats(c *gin.Context) {

	// source slice
	var beat_repeatDBs []orm.Beat_repeatDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBeat_repeats", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_repeat.GetDB()

	query := db.Find(&beat_repeatDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	beat_repeatAPIs := make([]orm.Beat_repeatAPI, 0)

	// for each beat_repeat, update fields from the database nullable fields
	for idx := range beat_repeatDBs {
		beat_repeatDB := &beat_repeatDBs[idx]
		_ = beat_repeatDB
		var beat_repeatAPI orm.Beat_repeatAPI

		// insertion point for updating fields
		beat_repeatAPI.ID = beat_repeatDB.ID
		beat_repeatDB.CopyBasicFieldsToBeat_repeat_WOP(&beat_repeatAPI.Beat_repeat_WOP)
		beat_repeatAPI.Beat_repeatPointersEncoding = beat_repeatDB.Beat_repeatPointersEncoding
		beat_repeatAPIs = append(beat_repeatAPIs, beat_repeatAPI)
	}

	c.JSON(http.StatusOK, beat_repeatAPIs)
}

// PostBeat_repeat
//
// swagger:route POST /beat_repeats beat_repeats postBeat_repeat
//
// Creates a beat_repeat
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBeat_repeat(c *gin.Context) {

	mutexBeat_repeat.Lock()
	defer mutexBeat_repeat.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBeat_repeats", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_repeat.GetDB()

	// Validate input
	var input orm.Beat_repeatAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create beat_repeat
	beat_repeatDB := orm.Beat_repeatDB{}
	beat_repeatDB.Beat_repeatPointersEncoding = input.Beat_repeatPointersEncoding
	beat_repeatDB.CopyBasicFieldsFromBeat_repeat_WOP(&input.Beat_repeat_WOP)

	query := db.Create(&beat_repeatDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBeat_repeat.CheckoutPhaseOneInstance(&beat_repeatDB)
	beat_repeat := backRepo.BackRepoBeat_repeat.Map_Beat_repeatDBID_Beat_repeatPtr[beat_repeatDB.ID]

	if beat_repeat != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), beat_repeat)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, beat_repeatDB)
}

// GetBeat_repeat
//
// swagger:route GET /beat_repeats/{ID} beat_repeats getBeat_repeat
//
// Gets the details for a beat_repeat.
//
// Responses:
// default: genericError
//
//	200: beat_repeatDBResponse
func (controller *Controller) GetBeat_repeat(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBeat_repeat", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_repeat.GetDB()

	// Get beat_repeatDB in DB
	var beat_repeatDB orm.Beat_repeatDB
	if err := db.First(&beat_repeatDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var beat_repeatAPI orm.Beat_repeatAPI
	beat_repeatAPI.ID = beat_repeatDB.ID
	beat_repeatAPI.Beat_repeatPointersEncoding = beat_repeatDB.Beat_repeatPointersEncoding
	beat_repeatDB.CopyBasicFieldsToBeat_repeat_WOP(&beat_repeatAPI.Beat_repeat_WOP)

	c.JSON(http.StatusOK, beat_repeatAPI)
}

// UpdateBeat_repeat
//
// swagger:route PATCH /beat_repeats/{ID} beat_repeats updateBeat_repeat
//
// # Update a beat_repeat
//
// Responses:
// default: genericError
//
//	200: beat_repeatDBResponse
func (controller *Controller) UpdateBeat_repeat(c *gin.Context) {

	mutexBeat_repeat.Lock()
	defer mutexBeat_repeat.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBeat_repeat", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_repeat.GetDB()

	// Validate input
	var input orm.Beat_repeatAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var beat_repeatDB orm.Beat_repeatDB

	// fetch the beat_repeat
	query := db.First(&beat_repeatDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	beat_repeatDB.CopyBasicFieldsFromBeat_repeat_WOP(&input.Beat_repeat_WOP)
	beat_repeatDB.Beat_repeatPointersEncoding = input.Beat_repeatPointersEncoding

	query = db.Model(&beat_repeatDB).Updates(beat_repeatDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	beat_repeatNew := new(models.Beat_repeat)
	beat_repeatDB.CopyBasicFieldsToBeat_repeat(beat_repeatNew)

	// redeem pointers
	beat_repeatDB.DecodePointers(backRepo, beat_repeatNew)

	// get stage instance from DB instance, and call callback function
	beat_repeatOld := backRepo.BackRepoBeat_repeat.Map_Beat_repeatDBID_Beat_repeatPtr[beat_repeatDB.ID]
	if beat_repeatOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), beat_repeatOld, beat_repeatNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the beat_repeatDB
	c.JSON(http.StatusOK, beat_repeatDB)
}

// DeleteBeat_repeat
//
// swagger:route DELETE /beat_repeats/{ID} beat_repeats deleteBeat_repeat
//
// # Delete a beat_repeat
//
// default: genericError
//
//	200: beat_repeatDBResponse
func (controller *Controller) DeleteBeat_repeat(c *gin.Context) {

	mutexBeat_repeat.Lock()
	defer mutexBeat_repeat.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBeat_repeat", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeat_repeat.GetDB()

	// Get model if exist
	var beat_repeatDB orm.Beat_repeatDB
	if err := db.First(&beat_repeatDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&beat_repeatDB)

	// get an instance (not staged) from DB instance, and call callback function
	beat_repeatDeleted := new(models.Beat_repeat)
	beat_repeatDB.CopyBasicFieldsToBeat_repeat(beat_repeatDeleted)

	// get stage instance from DB instance, and call callback function
	beat_repeatStaged := backRepo.BackRepoBeat_repeat.Map_Beat_repeatDBID_Beat_repeatPtr[beat_repeatDB.ID]
	if beat_repeatStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), beat_repeatStaged, beat_repeatDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
