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
var __Instrument_link__dummysDeclaration__ models.Instrument_link
var __Instrument_link_time__dummyDeclaration time.Duration

var mutexInstrument_link sync.Mutex

// An Instrument_linkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getInstrument_link updateInstrument_link deleteInstrument_link
type Instrument_linkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Instrument_linkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postInstrument_link updateInstrument_link
type Instrument_linkInput struct {
	// The Instrument_link to submit or modify
	// in: body
	Instrument_link *orm.Instrument_linkAPI
}

// GetInstrument_links
//
// swagger:route GET /instrument_links instrument_links getInstrument_links
//
// # Get all instrument_links
//
// Responses:
// default: genericError
//
//	200: instrument_linkDBResponse
func (controller *Controller) GetInstrument_links(c *gin.Context) {

	// source slice
	var instrument_linkDBs []orm.Instrument_linkDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInstrument_links", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_link.GetDB()

	query := db.Find(&instrument_linkDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	instrument_linkAPIs := make([]orm.Instrument_linkAPI, 0)

	// for each instrument_link, update fields from the database nullable fields
	for idx := range instrument_linkDBs {
		instrument_linkDB := &instrument_linkDBs[idx]
		_ = instrument_linkDB
		var instrument_linkAPI orm.Instrument_linkAPI

		// insertion point for updating fields
		instrument_linkAPI.ID = instrument_linkDB.ID
		instrument_linkDB.CopyBasicFieldsToInstrument_link_WOP(&instrument_linkAPI.Instrument_link_WOP)
		instrument_linkAPI.Instrument_linkPointersEncoding = instrument_linkDB.Instrument_linkPointersEncoding
		instrument_linkAPIs = append(instrument_linkAPIs, instrument_linkAPI)
	}

	c.JSON(http.StatusOK, instrument_linkAPIs)
}

// PostInstrument_link
//
// swagger:route POST /instrument_links instrument_links postInstrument_link
//
// Creates a instrument_link
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostInstrument_link(c *gin.Context) {

	mutexInstrument_link.Lock()
	defer mutexInstrument_link.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostInstrument_links", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_link.GetDB()

	// Validate input
	var input orm.Instrument_linkAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create instrument_link
	instrument_linkDB := orm.Instrument_linkDB{}
	instrument_linkDB.Instrument_linkPointersEncoding = input.Instrument_linkPointersEncoding
	instrument_linkDB.CopyBasicFieldsFromInstrument_link_WOP(&input.Instrument_link_WOP)

	query := db.Create(&instrument_linkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoInstrument_link.CheckoutPhaseOneInstance(&instrument_linkDB)
	instrument_link := backRepo.BackRepoInstrument_link.Map_Instrument_linkDBID_Instrument_linkPtr[instrument_linkDB.ID]

	if instrument_link != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), instrument_link)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, instrument_linkDB)
}

// GetInstrument_link
//
// swagger:route GET /instrument_links/{ID} instrument_links getInstrument_link
//
// Gets the details for a instrument_link.
//
// Responses:
// default: genericError
//
//	200: instrument_linkDBResponse
func (controller *Controller) GetInstrument_link(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInstrument_link", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_link.GetDB()

	// Get instrument_linkDB in DB
	var instrument_linkDB orm.Instrument_linkDB
	if err := db.First(&instrument_linkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var instrument_linkAPI orm.Instrument_linkAPI
	instrument_linkAPI.ID = instrument_linkDB.ID
	instrument_linkAPI.Instrument_linkPointersEncoding = instrument_linkDB.Instrument_linkPointersEncoding
	instrument_linkDB.CopyBasicFieldsToInstrument_link_WOP(&instrument_linkAPI.Instrument_link_WOP)

	c.JSON(http.StatusOK, instrument_linkAPI)
}

// UpdateInstrument_link
//
// swagger:route PATCH /instrument_links/{ID} instrument_links updateInstrument_link
//
// # Update a instrument_link
//
// Responses:
// default: genericError
//
//	200: instrument_linkDBResponse
func (controller *Controller) UpdateInstrument_link(c *gin.Context) {

	mutexInstrument_link.Lock()
	defer mutexInstrument_link.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateInstrument_link", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_link.GetDB()

	// Validate input
	var input orm.Instrument_linkAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var instrument_linkDB orm.Instrument_linkDB

	// fetch the instrument_link
	query := db.First(&instrument_linkDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	instrument_linkDB.CopyBasicFieldsFromInstrument_link_WOP(&input.Instrument_link_WOP)
	instrument_linkDB.Instrument_linkPointersEncoding = input.Instrument_linkPointersEncoding

	query = db.Model(&instrument_linkDB).Updates(instrument_linkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	instrument_linkNew := new(models.Instrument_link)
	instrument_linkDB.CopyBasicFieldsToInstrument_link(instrument_linkNew)

	// redeem pointers
	instrument_linkDB.DecodePointers(backRepo, instrument_linkNew)

	// get stage instance from DB instance, and call callback function
	instrument_linkOld := backRepo.BackRepoInstrument_link.Map_Instrument_linkDBID_Instrument_linkPtr[instrument_linkDB.ID]
	if instrument_linkOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), instrument_linkOld, instrument_linkNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the instrument_linkDB
	c.JSON(http.StatusOK, instrument_linkDB)
}

// DeleteInstrument_link
//
// swagger:route DELETE /instrument_links/{ID} instrument_links deleteInstrument_link
//
// # Delete a instrument_link
//
// default: genericError
//
//	200: instrument_linkDBResponse
func (controller *Controller) DeleteInstrument_link(c *gin.Context) {

	mutexInstrument_link.Lock()
	defer mutexInstrument_link.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteInstrument_link", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_link.GetDB()

	// Get model if exist
	var instrument_linkDB orm.Instrument_linkDB
	if err := db.First(&instrument_linkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&instrument_linkDB)

	// get an instance (not staged) from DB instance, and call callback function
	instrument_linkDeleted := new(models.Instrument_link)
	instrument_linkDB.CopyBasicFieldsToInstrument_link(instrument_linkDeleted)

	// get stage instance from DB instance, and call callback function
	instrument_linkStaged := backRepo.BackRepoInstrument_link.Map_Instrument_linkDBID_Instrument_linkPtr[instrument_linkDB.ID]
	if instrument_linkStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), instrument_linkStaged, instrument_linkDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
