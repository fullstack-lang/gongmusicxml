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
var __Note_type__dummysDeclaration__ models.Note_type
var __Note_type_time__dummyDeclaration time.Duration

var mutexNote_type sync.Mutex

// An Note_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNote_type updateNote_type deleteNote_type
type Note_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Note_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNote_type updateNote_type
type Note_typeInput struct {
	// The Note_type to submit or modify
	// in: body
	Note_type *orm.Note_typeAPI
}

// GetNote_types
//
// swagger:route GET /note_types note_types getNote_types
//
// # Get all note_types
//
// Responses:
// default: genericError
//
//	200: note_typeDBResponse
func (controller *Controller) GetNote_types(c *gin.Context) {

	// source slice
	var note_typeDBs []orm.Note_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNote_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_type.GetDB()

	query := db.Find(&note_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	note_typeAPIs := make([]orm.Note_typeAPI, 0)

	// for each note_type, update fields from the database nullable fields
	for idx := range note_typeDBs {
		note_typeDB := &note_typeDBs[idx]
		_ = note_typeDB
		var note_typeAPI orm.Note_typeAPI

		// insertion point for updating fields
		note_typeAPI.ID = note_typeDB.ID
		note_typeDB.CopyBasicFieldsToNote_type_WOP(&note_typeAPI.Note_type_WOP)
		note_typeAPI.Note_typePointersEncoding = note_typeDB.Note_typePointersEncoding
		note_typeAPIs = append(note_typeAPIs, note_typeAPI)
	}

	c.JSON(http.StatusOK, note_typeAPIs)
}

// PostNote_type
//
// swagger:route POST /note_types note_types postNote_type
//
// Creates a note_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNote_type(c *gin.Context) {

	mutexNote_type.Lock()
	defer mutexNote_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNote_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_type.GetDB()

	// Validate input
	var input orm.Note_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create note_type
	note_typeDB := orm.Note_typeDB{}
	note_typeDB.Note_typePointersEncoding = input.Note_typePointersEncoding
	note_typeDB.CopyBasicFieldsFromNote_type_WOP(&input.Note_type_WOP)

	query := db.Create(&note_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNote_type.CheckoutPhaseOneInstance(&note_typeDB)
	note_type := backRepo.BackRepoNote_type.Map_Note_typeDBID_Note_typePtr[note_typeDB.ID]

	if note_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), note_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, note_typeDB)
}

// GetNote_type
//
// swagger:route GET /note_types/{ID} note_types getNote_type
//
// Gets the details for a note_type.
//
// Responses:
// default: genericError
//
//	200: note_typeDBResponse
func (controller *Controller) GetNote_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNote_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_type.GetDB()

	// Get note_typeDB in DB
	var note_typeDB orm.Note_typeDB
	if err := db.First(&note_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var note_typeAPI orm.Note_typeAPI
	note_typeAPI.ID = note_typeDB.ID
	note_typeAPI.Note_typePointersEncoding = note_typeDB.Note_typePointersEncoding
	note_typeDB.CopyBasicFieldsToNote_type_WOP(&note_typeAPI.Note_type_WOP)

	c.JSON(http.StatusOK, note_typeAPI)
}

// UpdateNote_type
//
// swagger:route PATCH /note_types/{ID} note_types updateNote_type
//
// # Update a note_type
//
// Responses:
// default: genericError
//
//	200: note_typeDBResponse
func (controller *Controller) UpdateNote_type(c *gin.Context) {

	mutexNote_type.Lock()
	defer mutexNote_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNote_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_type.GetDB()

	// Validate input
	var input orm.Note_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var note_typeDB orm.Note_typeDB

	// fetch the note_type
	query := db.First(&note_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	note_typeDB.CopyBasicFieldsFromNote_type_WOP(&input.Note_type_WOP)
	note_typeDB.Note_typePointersEncoding = input.Note_typePointersEncoding

	query = db.Model(&note_typeDB).Updates(note_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	note_typeNew := new(models.Note_type)
	note_typeDB.CopyBasicFieldsToNote_type(note_typeNew)

	// redeem pointers
	note_typeDB.DecodePointers(backRepo, note_typeNew)

	// get stage instance from DB instance, and call callback function
	note_typeOld := backRepo.BackRepoNote_type.Map_Note_typeDBID_Note_typePtr[note_typeDB.ID]
	if note_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), note_typeOld, note_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the note_typeDB
	c.JSON(http.StatusOK, note_typeDB)
}

// DeleteNote_type
//
// swagger:route DELETE /note_types/{ID} note_types deleteNote_type
//
// # Delete a note_type
//
// default: genericError
//
//	200: note_typeDBResponse
func (controller *Controller) DeleteNote_type(c *gin.Context) {

	mutexNote_type.Lock()
	defer mutexNote_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNote_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNote_type.GetDB()

	// Get model if exist
	var note_typeDB orm.Note_typeDB
	if err := db.First(&note_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&note_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	note_typeDeleted := new(models.Note_type)
	note_typeDB.CopyBasicFieldsToNote_type(note_typeDeleted)

	// get stage instance from DB instance, and call callback function
	note_typeStaged := backRepo.BackRepoNote_type.Map_Note_typeDBID_Note_typePtr[note_typeDB.ID]
	if note_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), note_typeStaged, note_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
