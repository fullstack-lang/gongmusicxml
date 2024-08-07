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
var __Frame_note__dummysDeclaration__ models.Frame_note
var __Frame_note_time__dummyDeclaration time.Duration

var mutexFrame_note sync.Mutex

// An Frame_noteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFrame_note updateFrame_note deleteFrame_note
type Frame_noteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Frame_noteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFrame_note updateFrame_note
type Frame_noteInput struct {
	// The Frame_note to submit or modify
	// in: body
	Frame_note *orm.Frame_noteAPI
}

// GetFrame_notes
//
// swagger:route GET /frame_notes frame_notes getFrame_notes
//
// # Get all frame_notes
//
// Responses:
// default: genericError
//
//	200: frame_noteDBResponse
func (controller *Controller) GetFrame_notes(c *gin.Context) {

	// source slice
	var frame_noteDBs []orm.Frame_noteDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFrame_notes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame_note.GetDB()

	query := db.Find(&frame_noteDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	frame_noteAPIs := make([]orm.Frame_noteAPI, 0)

	// for each frame_note, update fields from the database nullable fields
	for idx := range frame_noteDBs {
		frame_noteDB := &frame_noteDBs[idx]
		_ = frame_noteDB
		var frame_noteAPI orm.Frame_noteAPI

		// insertion point for updating fields
		frame_noteAPI.ID = frame_noteDB.ID
		frame_noteDB.CopyBasicFieldsToFrame_note_WOP(&frame_noteAPI.Frame_note_WOP)
		frame_noteAPI.Frame_notePointersEncoding = frame_noteDB.Frame_notePointersEncoding
		frame_noteAPIs = append(frame_noteAPIs, frame_noteAPI)
	}

	c.JSON(http.StatusOK, frame_noteAPIs)
}

// PostFrame_note
//
// swagger:route POST /frame_notes frame_notes postFrame_note
//
// Creates a frame_note
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFrame_note(c *gin.Context) {

	mutexFrame_note.Lock()
	defer mutexFrame_note.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFrame_notes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame_note.GetDB()

	// Validate input
	var input orm.Frame_noteAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create frame_note
	frame_noteDB := orm.Frame_noteDB{}
	frame_noteDB.Frame_notePointersEncoding = input.Frame_notePointersEncoding
	frame_noteDB.CopyBasicFieldsFromFrame_note_WOP(&input.Frame_note_WOP)

	query := db.Create(&frame_noteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFrame_note.CheckoutPhaseOneInstance(&frame_noteDB)
	frame_note := backRepo.BackRepoFrame_note.Map_Frame_noteDBID_Frame_notePtr[frame_noteDB.ID]

	if frame_note != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), frame_note)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, frame_noteDB)
}

// GetFrame_note
//
// swagger:route GET /frame_notes/{ID} frame_notes getFrame_note
//
// Gets the details for a frame_note.
//
// Responses:
// default: genericError
//
//	200: frame_noteDBResponse
func (controller *Controller) GetFrame_note(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFrame_note", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame_note.GetDB()

	// Get frame_noteDB in DB
	var frame_noteDB orm.Frame_noteDB
	if err := db.First(&frame_noteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var frame_noteAPI orm.Frame_noteAPI
	frame_noteAPI.ID = frame_noteDB.ID
	frame_noteAPI.Frame_notePointersEncoding = frame_noteDB.Frame_notePointersEncoding
	frame_noteDB.CopyBasicFieldsToFrame_note_WOP(&frame_noteAPI.Frame_note_WOP)

	c.JSON(http.StatusOK, frame_noteAPI)
}

// UpdateFrame_note
//
// swagger:route PATCH /frame_notes/{ID} frame_notes updateFrame_note
//
// # Update a frame_note
//
// Responses:
// default: genericError
//
//	200: frame_noteDBResponse
func (controller *Controller) UpdateFrame_note(c *gin.Context) {

	mutexFrame_note.Lock()
	defer mutexFrame_note.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFrame_note", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame_note.GetDB()

	// Validate input
	var input orm.Frame_noteAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var frame_noteDB orm.Frame_noteDB

	// fetch the frame_note
	query := db.First(&frame_noteDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	frame_noteDB.CopyBasicFieldsFromFrame_note_WOP(&input.Frame_note_WOP)
	frame_noteDB.Frame_notePointersEncoding = input.Frame_notePointersEncoding

	query = db.Model(&frame_noteDB).Updates(frame_noteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	frame_noteNew := new(models.Frame_note)
	frame_noteDB.CopyBasicFieldsToFrame_note(frame_noteNew)

	// redeem pointers
	frame_noteDB.DecodePointers(backRepo, frame_noteNew)

	// get stage instance from DB instance, and call callback function
	frame_noteOld := backRepo.BackRepoFrame_note.Map_Frame_noteDBID_Frame_notePtr[frame_noteDB.ID]
	if frame_noteOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), frame_noteOld, frame_noteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the frame_noteDB
	c.JSON(http.StatusOK, frame_noteDB)
}

// DeleteFrame_note
//
// swagger:route DELETE /frame_notes/{ID} frame_notes deleteFrame_note
//
// # Delete a frame_note
//
// default: genericError
//
//	200: frame_noteDBResponse
func (controller *Controller) DeleteFrame_note(c *gin.Context) {

	mutexFrame_note.Lock()
	defer mutexFrame_note.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFrame_note", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame_note.GetDB()

	// Get model if exist
	var frame_noteDB orm.Frame_noteDB
	if err := db.First(&frame_noteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&frame_noteDB)

	// get an instance (not staged) from DB instance, and call callback function
	frame_noteDeleted := new(models.Frame_note)
	frame_noteDB.CopyBasicFieldsToFrame_note(frame_noteDeleted)

	// get stage instance from DB instance, and call callback function
	frame_noteStaged := backRepo.BackRepoFrame_note.Map_Frame_noteDBID_Frame_notePtr[frame_noteDB.ID]
	if frame_noteStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), frame_noteStaged, frame_noteDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
