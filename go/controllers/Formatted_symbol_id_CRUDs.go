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
var __Formatted_symbol_id__dummysDeclaration__ models.Formatted_symbol_id
var __Formatted_symbol_id_time__dummyDeclaration time.Duration

var mutexFormatted_symbol_id sync.Mutex

// An Formatted_symbol_idID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormatted_symbol_id updateFormatted_symbol_id deleteFormatted_symbol_id
type Formatted_symbol_idID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Formatted_symbol_idInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormatted_symbol_id updateFormatted_symbol_id
type Formatted_symbol_idInput struct {
	// The Formatted_symbol_id to submit or modify
	// in: body
	Formatted_symbol_id *orm.Formatted_symbol_idAPI
}

// GetFormatted_symbol_ids
//
// swagger:route GET /formatted_symbol_ids formatted_symbol_ids getFormatted_symbol_ids
//
// # Get all formatted_symbol_ids
//
// Responses:
// default: genericError
//
//	200: formatted_symbol_idDBResponse
func (controller *Controller) GetFormatted_symbol_ids(c *gin.Context) {

	// source slice
	var formatted_symbol_idDBs []orm.Formatted_symbol_idDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormatted_symbol_ids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol_id.GetDB()

	query := db.Find(&formatted_symbol_idDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formatted_symbol_idAPIs := make([]orm.Formatted_symbol_idAPI, 0)

	// for each formatted_symbol_id, update fields from the database nullable fields
	for idx := range formatted_symbol_idDBs {
		formatted_symbol_idDB := &formatted_symbol_idDBs[idx]
		_ = formatted_symbol_idDB
		var formatted_symbol_idAPI orm.Formatted_symbol_idAPI

		// insertion point for updating fields
		formatted_symbol_idAPI.ID = formatted_symbol_idDB.ID
		formatted_symbol_idDB.CopyBasicFieldsToFormatted_symbol_id_WOP(&formatted_symbol_idAPI.Formatted_symbol_id_WOP)
		formatted_symbol_idAPI.Formatted_symbol_idPointersEncoding = formatted_symbol_idDB.Formatted_symbol_idPointersEncoding
		formatted_symbol_idAPIs = append(formatted_symbol_idAPIs, formatted_symbol_idAPI)
	}

	c.JSON(http.StatusOK, formatted_symbol_idAPIs)
}

// PostFormatted_symbol_id
//
// swagger:route POST /formatted_symbol_ids formatted_symbol_ids postFormatted_symbol_id
//
// Creates a formatted_symbol_id
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormatted_symbol_id(c *gin.Context) {

	mutexFormatted_symbol_id.Lock()
	defer mutexFormatted_symbol_id.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormatted_symbol_ids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol_id.GetDB()

	// Validate input
	var input orm.Formatted_symbol_idAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formatted_symbol_id
	formatted_symbol_idDB := orm.Formatted_symbol_idDB{}
	formatted_symbol_idDB.Formatted_symbol_idPointersEncoding = input.Formatted_symbol_idPointersEncoding
	formatted_symbol_idDB.CopyBasicFieldsFromFormatted_symbol_id_WOP(&input.Formatted_symbol_id_WOP)

	query := db.Create(&formatted_symbol_idDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormatted_symbol_id.CheckoutPhaseOneInstance(&formatted_symbol_idDB)
	formatted_symbol_id := backRepo.BackRepoFormatted_symbol_id.Map_Formatted_symbol_idDBID_Formatted_symbol_idPtr[formatted_symbol_idDB.ID]

	if formatted_symbol_id != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formatted_symbol_id)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formatted_symbol_idDB)
}

// GetFormatted_symbol_id
//
// swagger:route GET /formatted_symbol_ids/{ID} formatted_symbol_ids getFormatted_symbol_id
//
// Gets the details for a formatted_symbol_id.
//
// Responses:
// default: genericError
//
//	200: formatted_symbol_idDBResponse
func (controller *Controller) GetFormatted_symbol_id(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormatted_symbol_id", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol_id.GetDB()

	// Get formatted_symbol_idDB in DB
	var formatted_symbol_idDB orm.Formatted_symbol_idDB
	if err := db.First(&formatted_symbol_idDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formatted_symbol_idAPI orm.Formatted_symbol_idAPI
	formatted_symbol_idAPI.ID = formatted_symbol_idDB.ID
	formatted_symbol_idAPI.Formatted_symbol_idPointersEncoding = formatted_symbol_idDB.Formatted_symbol_idPointersEncoding
	formatted_symbol_idDB.CopyBasicFieldsToFormatted_symbol_id_WOP(&formatted_symbol_idAPI.Formatted_symbol_id_WOP)

	c.JSON(http.StatusOK, formatted_symbol_idAPI)
}

// UpdateFormatted_symbol_id
//
// swagger:route PATCH /formatted_symbol_ids/{ID} formatted_symbol_ids updateFormatted_symbol_id
//
// # Update a formatted_symbol_id
//
// Responses:
// default: genericError
//
//	200: formatted_symbol_idDBResponse
func (controller *Controller) UpdateFormatted_symbol_id(c *gin.Context) {

	mutexFormatted_symbol_id.Lock()
	defer mutexFormatted_symbol_id.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormatted_symbol_id", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol_id.GetDB()

	// Validate input
	var input orm.Formatted_symbol_idAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formatted_symbol_idDB orm.Formatted_symbol_idDB

	// fetch the formatted_symbol_id
	query := db.First(&formatted_symbol_idDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formatted_symbol_idDB.CopyBasicFieldsFromFormatted_symbol_id_WOP(&input.Formatted_symbol_id_WOP)
	formatted_symbol_idDB.Formatted_symbol_idPointersEncoding = input.Formatted_symbol_idPointersEncoding

	query = db.Model(&formatted_symbol_idDB).Updates(formatted_symbol_idDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formatted_symbol_idNew := new(models.Formatted_symbol_id)
	formatted_symbol_idDB.CopyBasicFieldsToFormatted_symbol_id(formatted_symbol_idNew)

	// redeem pointers
	formatted_symbol_idDB.DecodePointers(backRepo, formatted_symbol_idNew)

	// get stage instance from DB instance, and call callback function
	formatted_symbol_idOld := backRepo.BackRepoFormatted_symbol_id.Map_Formatted_symbol_idDBID_Formatted_symbol_idPtr[formatted_symbol_idDB.ID]
	if formatted_symbol_idOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formatted_symbol_idOld, formatted_symbol_idNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formatted_symbol_idDB
	c.JSON(http.StatusOK, formatted_symbol_idDB)
}

// DeleteFormatted_symbol_id
//
// swagger:route DELETE /formatted_symbol_ids/{ID} formatted_symbol_ids deleteFormatted_symbol_id
//
// # Delete a formatted_symbol_id
//
// default: genericError
//
//	200: formatted_symbol_idDBResponse
func (controller *Controller) DeleteFormatted_symbol_id(c *gin.Context) {

	mutexFormatted_symbol_id.Lock()
	defer mutexFormatted_symbol_id.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormatted_symbol_id", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol_id.GetDB()

	// Get model if exist
	var formatted_symbol_idDB orm.Formatted_symbol_idDB
	if err := db.First(&formatted_symbol_idDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&formatted_symbol_idDB)

	// get an instance (not staged) from DB instance, and call callback function
	formatted_symbol_idDeleted := new(models.Formatted_symbol_id)
	formatted_symbol_idDB.CopyBasicFieldsToFormatted_symbol_id(formatted_symbol_idDeleted)

	// get stage instance from DB instance, and call callback function
	formatted_symbol_idStaged := backRepo.BackRepoFormatted_symbol_id.Map_Formatted_symbol_idDBID_Formatted_symbol_idPtr[formatted_symbol_idDB.ID]
	if formatted_symbol_idStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formatted_symbol_idStaged, formatted_symbol_idDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
