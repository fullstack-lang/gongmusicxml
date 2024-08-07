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
var __Metronome_note__dummysDeclaration__ models.Metronome_note
var __Metronome_note_time__dummyDeclaration time.Duration

var mutexMetronome_note sync.Mutex

// An Metronome_noteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMetronome_note updateMetronome_note deleteMetronome_note
type Metronome_noteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Metronome_noteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMetronome_note updateMetronome_note
type Metronome_noteInput struct {
	// The Metronome_note to submit or modify
	// in: body
	Metronome_note *orm.Metronome_noteAPI
}

// GetMetronome_notes
//
// swagger:route GET /metronome_notes metronome_notes getMetronome_notes
//
// # Get all metronome_notes
//
// Responses:
// default: genericError
//
//	200: metronome_noteDBResponse
func (controller *Controller) GetMetronome_notes(c *gin.Context) {

	// source slice
	var metronome_noteDBs []orm.Metronome_noteDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronome_notes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_note.GetDB()

	query := db.Find(&metronome_noteDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	metronome_noteAPIs := make([]orm.Metronome_noteAPI, 0)

	// for each metronome_note, update fields from the database nullable fields
	for idx := range metronome_noteDBs {
		metronome_noteDB := &metronome_noteDBs[idx]
		_ = metronome_noteDB
		var metronome_noteAPI orm.Metronome_noteAPI

		// insertion point for updating fields
		metronome_noteAPI.ID = metronome_noteDB.ID
		metronome_noteDB.CopyBasicFieldsToMetronome_note_WOP(&metronome_noteAPI.Metronome_note_WOP)
		metronome_noteAPI.Metronome_notePointersEncoding = metronome_noteDB.Metronome_notePointersEncoding
		metronome_noteAPIs = append(metronome_noteAPIs, metronome_noteAPI)
	}

	c.JSON(http.StatusOK, metronome_noteAPIs)
}

// PostMetronome_note
//
// swagger:route POST /metronome_notes metronome_notes postMetronome_note
//
// Creates a metronome_note
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMetronome_note(c *gin.Context) {

	mutexMetronome_note.Lock()
	defer mutexMetronome_note.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMetronome_notes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_note.GetDB()

	// Validate input
	var input orm.Metronome_noteAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create metronome_note
	metronome_noteDB := orm.Metronome_noteDB{}
	metronome_noteDB.Metronome_notePointersEncoding = input.Metronome_notePointersEncoding
	metronome_noteDB.CopyBasicFieldsFromMetronome_note_WOP(&input.Metronome_note_WOP)

	query := db.Create(&metronome_noteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMetronome_note.CheckoutPhaseOneInstance(&metronome_noteDB)
	metronome_note := backRepo.BackRepoMetronome_note.Map_Metronome_noteDBID_Metronome_notePtr[metronome_noteDB.ID]

	if metronome_note != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), metronome_note)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, metronome_noteDB)
}

// GetMetronome_note
//
// swagger:route GET /metronome_notes/{ID} metronome_notes getMetronome_note
//
// Gets the details for a metronome_note.
//
// Responses:
// default: genericError
//
//	200: metronome_noteDBResponse
func (controller *Controller) GetMetronome_note(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronome_note", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_note.GetDB()

	// Get metronome_noteDB in DB
	var metronome_noteDB orm.Metronome_noteDB
	if err := db.First(&metronome_noteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var metronome_noteAPI orm.Metronome_noteAPI
	metronome_noteAPI.ID = metronome_noteDB.ID
	metronome_noteAPI.Metronome_notePointersEncoding = metronome_noteDB.Metronome_notePointersEncoding
	metronome_noteDB.CopyBasicFieldsToMetronome_note_WOP(&metronome_noteAPI.Metronome_note_WOP)

	c.JSON(http.StatusOK, metronome_noteAPI)
}

// UpdateMetronome_note
//
// swagger:route PATCH /metronome_notes/{ID} metronome_notes updateMetronome_note
//
// # Update a metronome_note
//
// Responses:
// default: genericError
//
//	200: metronome_noteDBResponse
func (controller *Controller) UpdateMetronome_note(c *gin.Context) {

	mutexMetronome_note.Lock()
	defer mutexMetronome_note.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMetronome_note", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_note.GetDB()

	// Validate input
	var input orm.Metronome_noteAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var metronome_noteDB orm.Metronome_noteDB

	// fetch the metronome_note
	query := db.First(&metronome_noteDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	metronome_noteDB.CopyBasicFieldsFromMetronome_note_WOP(&input.Metronome_note_WOP)
	metronome_noteDB.Metronome_notePointersEncoding = input.Metronome_notePointersEncoding

	query = db.Model(&metronome_noteDB).Updates(metronome_noteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	metronome_noteNew := new(models.Metronome_note)
	metronome_noteDB.CopyBasicFieldsToMetronome_note(metronome_noteNew)

	// redeem pointers
	metronome_noteDB.DecodePointers(backRepo, metronome_noteNew)

	// get stage instance from DB instance, and call callback function
	metronome_noteOld := backRepo.BackRepoMetronome_note.Map_Metronome_noteDBID_Metronome_notePtr[metronome_noteDB.ID]
	if metronome_noteOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), metronome_noteOld, metronome_noteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the metronome_noteDB
	c.JSON(http.StatusOK, metronome_noteDB)
}

// DeleteMetronome_note
//
// swagger:route DELETE /metronome_notes/{ID} metronome_notes deleteMetronome_note
//
// # Delete a metronome_note
//
// default: genericError
//
//	200: metronome_noteDBResponse
func (controller *Controller) DeleteMetronome_note(c *gin.Context) {

	mutexMetronome_note.Lock()
	defer mutexMetronome_note.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMetronome_note", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_note.GetDB()

	// Get model if exist
	var metronome_noteDB orm.Metronome_noteDB
	if err := db.First(&metronome_noteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&metronome_noteDB)

	// get an instance (not staged) from DB instance, and call callback function
	metronome_noteDeleted := new(models.Metronome_note)
	metronome_noteDB.CopyBasicFieldsToMetronome_note(metronome_noteDeleted)

	// get stage instance from DB instance, and call callback function
	metronome_noteStaged := backRepo.BackRepoMetronome_note.Map_Metronome_noteDBID_Metronome_notePtr[metronome_noteDB.ID]
	if metronome_noteStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), metronome_noteStaged, metronome_noteDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
