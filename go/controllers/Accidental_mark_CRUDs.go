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
var __Accidental_mark__dummysDeclaration__ models.Accidental_mark
var __Accidental_mark_time__dummyDeclaration time.Duration

var mutexAccidental_mark sync.Mutex

// An Accidental_markID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAccidental_mark updateAccidental_mark deleteAccidental_mark
type Accidental_markID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Accidental_markInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAccidental_mark updateAccidental_mark
type Accidental_markInput struct {
	// The Accidental_mark to submit or modify
	// in: body
	Accidental_mark *orm.Accidental_markAPI
}

// GetAccidental_marks
//
// swagger:route GET /accidental_marks accidental_marks getAccidental_marks
//
// # Get all accidental_marks
//
// Responses:
// default: genericError
//
//	200: accidental_markDBResponse
func (controller *Controller) GetAccidental_marks(c *gin.Context) {

	// source slice
	var accidental_markDBs []orm.Accidental_markDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccidental_marks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_mark.GetDB()

	query := db.Find(&accidental_markDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	accidental_markAPIs := make([]orm.Accidental_markAPI, 0)

	// for each accidental_mark, update fields from the database nullable fields
	for idx := range accidental_markDBs {
		accidental_markDB := &accidental_markDBs[idx]
		_ = accidental_markDB
		var accidental_markAPI orm.Accidental_markAPI

		// insertion point for updating fields
		accidental_markAPI.ID = accidental_markDB.ID
		accidental_markDB.CopyBasicFieldsToAccidental_mark_WOP(&accidental_markAPI.Accidental_mark_WOP)
		accidental_markAPI.Accidental_markPointersEncoding = accidental_markDB.Accidental_markPointersEncoding
		accidental_markAPIs = append(accidental_markAPIs, accidental_markAPI)
	}

	c.JSON(http.StatusOK, accidental_markAPIs)
}

// PostAccidental_mark
//
// swagger:route POST /accidental_marks accidental_marks postAccidental_mark
//
// Creates a accidental_mark
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAccidental_mark(c *gin.Context) {

	mutexAccidental_mark.Lock()
	defer mutexAccidental_mark.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAccidental_marks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_mark.GetDB()

	// Validate input
	var input orm.Accidental_markAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create accidental_mark
	accidental_markDB := orm.Accidental_markDB{}
	accidental_markDB.Accidental_markPointersEncoding = input.Accidental_markPointersEncoding
	accidental_markDB.CopyBasicFieldsFromAccidental_mark_WOP(&input.Accidental_mark_WOP)

	query := db.Create(&accidental_markDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAccidental_mark.CheckoutPhaseOneInstance(&accidental_markDB)
	accidental_mark := backRepo.BackRepoAccidental_mark.Map_Accidental_markDBID_Accidental_markPtr[accidental_markDB.ID]

	if accidental_mark != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), accidental_mark)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, accidental_markDB)
}

// GetAccidental_mark
//
// swagger:route GET /accidental_marks/{ID} accidental_marks getAccidental_mark
//
// Gets the details for a accidental_mark.
//
// Responses:
// default: genericError
//
//	200: accidental_markDBResponse
func (controller *Controller) GetAccidental_mark(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccidental_mark", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_mark.GetDB()

	// Get accidental_markDB in DB
	var accidental_markDB orm.Accidental_markDB
	if err := db.First(&accidental_markDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var accidental_markAPI orm.Accidental_markAPI
	accidental_markAPI.ID = accidental_markDB.ID
	accidental_markAPI.Accidental_markPointersEncoding = accidental_markDB.Accidental_markPointersEncoding
	accidental_markDB.CopyBasicFieldsToAccidental_mark_WOP(&accidental_markAPI.Accidental_mark_WOP)

	c.JSON(http.StatusOK, accidental_markAPI)
}

// UpdateAccidental_mark
//
// swagger:route PATCH /accidental_marks/{ID} accidental_marks updateAccidental_mark
//
// # Update a accidental_mark
//
// Responses:
// default: genericError
//
//	200: accidental_markDBResponse
func (controller *Controller) UpdateAccidental_mark(c *gin.Context) {

	mutexAccidental_mark.Lock()
	defer mutexAccidental_mark.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAccidental_mark", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_mark.GetDB()

	// Validate input
	var input orm.Accidental_markAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var accidental_markDB orm.Accidental_markDB

	// fetch the accidental_mark
	query := db.First(&accidental_markDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	accidental_markDB.CopyBasicFieldsFromAccidental_mark_WOP(&input.Accidental_mark_WOP)
	accidental_markDB.Accidental_markPointersEncoding = input.Accidental_markPointersEncoding

	query = db.Model(&accidental_markDB).Updates(accidental_markDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	accidental_markNew := new(models.Accidental_mark)
	accidental_markDB.CopyBasicFieldsToAccidental_mark(accidental_markNew)

	// redeem pointers
	accidental_markDB.DecodePointers(backRepo, accidental_markNew)

	// get stage instance from DB instance, and call callback function
	accidental_markOld := backRepo.BackRepoAccidental_mark.Map_Accidental_markDBID_Accidental_markPtr[accidental_markDB.ID]
	if accidental_markOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), accidental_markOld, accidental_markNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the accidental_markDB
	c.JSON(http.StatusOK, accidental_markDB)
}

// DeleteAccidental_mark
//
// swagger:route DELETE /accidental_marks/{ID} accidental_marks deleteAccidental_mark
//
// # Delete a accidental_mark
//
// default: genericError
//
//	200: accidental_markDBResponse
func (controller *Controller) DeleteAccidental_mark(c *gin.Context) {

	mutexAccidental_mark.Lock()
	defer mutexAccidental_mark.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAccidental_mark", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_mark.GetDB()

	// Get model if exist
	var accidental_markDB orm.Accidental_markDB
	if err := db.First(&accidental_markDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&accidental_markDB)

	// get an instance (not staged) from DB instance, and call callback function
	accidental_markDeleted := new(models.Accidental_mark)
	accidental_markDB.CopyBasicFieldsToAccidental_mark(accidental_markDeleted)

	// get stage instance from DB instance, and call callback function
	accidental_markStaged := backRepo.BackRepoAccidental_mark.Map_Accidental_markDBID_Accidental_markPtr[accidental_markDB.ID]
	if accidental_markStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), accidental_markStaged, accidental_markDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
