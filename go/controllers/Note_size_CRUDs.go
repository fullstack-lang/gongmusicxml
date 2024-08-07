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
var __Note_size__dummysDeclaration__ models.Note_size
var __Note_size_time__dummyDeclaration time.Duration

var mutexNote_size sync.Mutex

// An Note_sizeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNote_size updateNote_size deleteNote_size
type Note_sizeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Note_sizeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNote_size updateNote_size
type Note_sizeInput struct {
	// The Note_size to submit or modify
	// in: body
	Note_size *orm.Note_sizeAPI
}

// GetNote_sizes
//
// swagger:route GET /note_sizes note_sizes getNote_sizes
//
// # Get all note_sizes
//
// Responses:
// default: genericError
//
//	200: note_sizeDBResponse
func (controller *Controller) GetNote_sizes(c *gin.Context) {

	// source slice
	var note_sizeDBs []orm.Note_sizeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNote_sizes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_size.GetDB()

	query := db.Find(&note_sizeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	note_sizeAPIs := make([]orm.Note_sizeAPI, 0)

	// for each note_size, update fields from the database nullable fields
	for idx := range note_sizeDBs {
		note_sizeDB := &note_sizeDBs[idx]
		_ = note_sizeDB
		var note_sizeAPI orm.Note_sizeAPI

		// insertion point for updating fields
		note_sizeAPI.ID = note_sizeDB.ID
		note_sizeDB.CopyBasicFieldsToNote_size_WOP(&note_sizeAPI.Note_size_WOP)
		note_sizeAPI.Note_sizePointersEncoding = note_sizeDB.Note_sizePointersEncoding
		note_sizeAPIs = append(note_sizeAPIs, note_sizeAPI)
	}

	c.JSON(http.StatusOK, note_sizeAPIs)
}

// PostNote_size
//
// swagger:route POST /note_sizes note_sizes postNote_size
//
// Creates a note_size
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNote_size(c *gin.Context) {

	mutexNote_size.Lock()
	defer mutexNote_size.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNote_sizes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_size.GetDB()

	// Validate input
	var input orm.Note_sizeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create note_size
	note_sizeDB := orm.Note_sizeDB{}
	note_sizeDB.Note_sizePointersEncoding = input.Note_sizePointersEncoding
	note_sizeDB.CopyBasicFieldsFromNote_size_WOP(&input.Note_size_WOP)

	query := db.Create(&note_sizeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNote_size.CheckoutPhaseOneInstance(&note_sizeDB)
	note_size := backRepo.BackRepoNote_size.Map_Note_sizeDBID_Note_sizePtr[note_sizeDB.ID]

	if note_size != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), note_size)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, note_sizeDB)
}

// GetNote_size
//
// swagger:route GET /note_sizes/{ID} note_sizes getNote_size
//
// Gets the details for a note_size.
//
// Responses:
// default: genericError
//
//	200: note_sizeDBResponse
func (controller *Controller) GetNote_size(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNote_size", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_size.GetDB()

	// Get note_sizeDB in DB
	var note_sizeDB orm.Note_sizeDB
	if err := db.First(&note_sizeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var note_sizeAPI orm.Note_sizeAPI
	note_sizeAPI.ID = note_sizeDB.ID
	note_sizeAPI.Note_sizePointersEncoding = note_sizeDB.Note_sizePointersEncoding
	note_sizeDB.CopyBasicFieldsToNote_size_WOP(&note_sizeAPI.Note_size_WOP)

	c.JSON(http.StatusOK, note_sizeAPI)
}

// UpdateNote_size
//
// swagger:route PATCH /note_sizes/{ID} note_sizes updateNote_size
//
// # Update a note_size
//
// Responses:
// default: genericError
//
//	200: note_sizeDBResponse
func (controller *Controller) UpdateNote_size(c *gin.Context) {

	mutexNote_size.Lock()
	defer mutexNote_size.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNote_size", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_size.GetDB()

	// Validate input
	var input orm.Note_sizeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var note_sizeDB orm.Note_sizeDB

	// fetch the note_size
	query := db.First(&note_sizeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	note_sizeDB.CopyBasicFieldsFromNote_size_WOP(&input.Note_size_WOP)
	note_sizeDB.Note_sizePointersEncoding = input.Note_sizePointersEncoding

	query = db.Model(&note_sizeDB).Updates(note_sizeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	note_sizeNew := new(models.Note_size)
	note_sizeDB.CopyBasicFieldsToNote_size(note_sizeNew)

	// redeem pointers
	note_sizeDB.DecodePointers(backRepo, note_sizeNew)

	// get stage instance from DB instance, and call callback function
	note_sizeOld := backRepo.BackRepoNote_size.Map_Note_sizeDBID_Note_sizePtr[note_sizeDB.ID]
	if note_sizeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), note_sizeOld, note_sizeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the note_sizeDB
	c.JSON(http.StatusOK, note_sizeDB)
}

// DeleteNote_size
//
// swagger:route DELETE /note_sizes/{ID} note_sizes deleteNote_size
//
// # Delete a note_size
//
// default: genericError
//
//	200: note_sizeDBResponse
func (controller *Controller) DeleteNote_size(c *gin.Context) {

	mutexNote_size.Lock()
	defer mutexNote_size.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNote_size", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_size.GetDB()

	// Get model if exist
	var note_sizeDB orm.Note_sizeDB
	if err := db.First(&note_sizeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&note_sizeDB)

	// get an instance (not staged) from DB instance, and call callback function
	note_sizeDeleted := new(models.Note_size)
	note_sizeDB.CopyBasicFieldsToNote_size(note_sizeDeleted)

	// get stage instance from DB instance, and call callback function
	note_sizeStaged := backRepo.BackRepoNote_size.Map_Note_sizeDBID_Note_sizePtr[note_sizeDB.ID]
	if note_sizeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), note_sizeStaged, note_sizeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
