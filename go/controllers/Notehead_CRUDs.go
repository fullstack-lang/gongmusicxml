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
var __Notehead__dummysDeclaration__ models.Notehead
var __Notehead_time__dummyDeclaration time.Duration

var mutexNotehead sync.Mutex

// An NoteheadID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNotehead updateNotehead deleteNotehead
type NoteheadID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NoteheadInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNotehead updateNotehead
type NoteheadInput struct {
	// The Notehead to submit or modify
	// in: body
	Notehead *orm.NoteheadAPI
}

// GetNoteheads
//
// swagger:route GET /noteheads noteheads getNoteheads
//
// # Get all noteheads
//
// Responses:
// default: genericError
//
//	200: noteheadDBResponse
func (controller *Controller) GetNoteheads(c *gin.Context) {

	// source slice
	var noteheadDBs []orm.NoteheadDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNoteheads", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead.GetDB()

	query := db.Find(&noteheadDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	noteheadAPIs := make([]orm.NoteheadAPI, 0)

	// for each notehead, update fields from the database nullable fields
	for idx := range noteheadDBs {
		noteheadDB := &noteheadDBs[idx]
		_ = noteheadDB
		var noteheadAPI orm.NoteheadAPI

		// insertion point for updating fields
		noteheadAPI.ID = noteheadDB.ID
		noteheadDB.CopyBasicFieldsToNotehead_WOP(&noteheadAPI.Notehead_WOP)
		noteheadAPI.NoteheadPointersEncoding = noteheadDB.NoteheadPointersEncoding
		noteheadAPIs = append(noteheadAPIs, noteheadAPI)
	}

	c.JSON(http.StatusOK, noteheadAPIs)
}

// PostNotehead
//
// swagger:route POST /noteheads noteheads postNotehead
//
// Creates a notehead
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNotehead(c *gin.Context) {

	mutexNotehead.Lock()
	defer mutexNotehead.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNoteheads", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead.GetDB()

	// Validate input
	var input orm.NoteheadAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create notehead
	noteheadDB := orm.NoteheadDB{}
	noteheadDB.NoteheadPointersEncoding = input.NoteheadPointersEncoding
	noteheadDB.CopyBasicFieldsFromNotehead_WOP(&input.Notehead_WOP)

	query := db.Create(&noteheadDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNotehead.CheckoutPhaseOneInstance(&noteheadDB)
	notehead := backRepo.BackRepoNotehead.Map_NoteheadDBID_NoteheadPtr[noteheadDB.ID]

	if notehead != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), notehead)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, noteheadDB)
}

// GetNotehead
//
// swagger:route GET /noteheads/{ID} noteheads getNotehead
//
// Gets the details for a notehead.
//
// Responses:
// default: genericError
//
//	200: noteheadDBResponse
func (controller *Controller) GetNotehead(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNotehead", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead.GetDB()

	// Get noteheadDB in DB
	var noteheadDB orm.NoteheadDB
	if err := db.First(&noteheadDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var noteheadAPI orm.NoteheadAPI
	noteheadAPI.ID = noteheadDB.ID
	noteheadAPI.NoteheadPointersEncoding = noteheadDB.NoteheadPointersEncoding
	noteheadDB.CopyBasicFieldsToNotehead_WOP(&noteheadAPI.Notehead_WOP)

	c.JSON(http.StatusOK, noteheadAPI)
}

// UpdateNotehead
//
// swagger:route PATCH /noteheads/{ID} noteheads updateNotehead
//
// # Update a notehead
//
// Responses:
// default: genericError
//
//	200: noteheadDBResponse
func (controller *Controller) UpdateNotehead(c *gin.Context) {

	mutexNotehead.Lock()
	defer mutexNotehead.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNotehead", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead.GetDB()

	// Validate input
	var input orm.NoteheadAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var noteheadDB orm.NoteheadDB

	// fetch the notehead
	query := db.First(&noteheadDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	noteheadDB.CopyBasicFieldsFromNotehead_WOP(&input.Notehead_WOP)
	noteheadDB.NoteheadPointersEncoding = input.NoteheadPointersEncoding

	query = db.Model(&noteheadDB).Updates(noteheadDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	noteheadNew := new(models.Notehead)
	noteheadDB.CopyBasicFieldsToNotehead(noteheadNew)

	// redeem pointers
	noteheadDB.DecodePointers(backRepo, noteheadNew)

	// get stage instance from DB instance, and call callback function
	noteheadOld := backRepo.BackRepoNotehead.Map_NoteheadDBID_NoteheadPtr[noteheadDB.ID]
	if noteheadOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), noteheadOld, noteheadNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the noteheadDB
	c.JSON(http.StatusOK, noteheadDB)
}

// DeleteNotehead
//
// swagger:route DELETE /noteheads/{ID} noteheads deleteNotehead
//
// # Delete a notehead
//
// default: genericError
//
//	200: noteheadDBResponse
func (controller *Controller) DeleteNotehead(c *gin.Context) {

	mutexNotehead.Lock()
	defer mutexNotehead.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNotehead", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead.GetDB()

	// Get model if exist
	var noteheadDB orm.NoteheadDB
	if err := db.First(&noteheadDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&noteheadDB)

	// get an instance (not staged) from DB instance, and call callback function
	noteheadDeleted := new(models.Notehead)
	noteheadDB.CopyBasicFieldsToNotehead(noteheadDeleted)

	// get stage instance from DB instance, and call callback function
	noteheadStaged := backRepo.BackRepoNotehead.Map_NoteheadDBID_NoteheadPtr[noteheadDB.ID]
	if noteheadStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), noteheadStaged, noteheadDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
