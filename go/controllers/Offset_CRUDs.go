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
var __Offset__dummysDeclaration__ models.Offset
var __Offset_time__dummyDeclaration time.Duration

var mutexOffset sync.Mutex

// An OffsetID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOffset updateOffset deleteOffset
type OffsetID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// OffsetInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOffset updateOffset
type OffsetInput struct {
	// The Offset to submit or modify
	// in: body
	Offset *orm.OffsetAPI
}

// GetOffsets
//
// swagger:route GET /offsets offsets getOffsets
//
// # Get all offsets
//
// Responses:
// default: genericError
//
//	200: offsetDBResponse
func (controller *Controller) GetOffsets(c *gin.Context) {

	// source slice
	var offsetDBs []orm.OffsetDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOffsets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOffset.GetDB()

	query := db.Find(&offsetDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	offsetAPIs := make([]orm.OffsetAPI, 0)

	// for each offset, update fields from the database nullable fields
	for idx := range offsetDBs {
		offsetDB := &offsetDBs[idx]
		_ = offsetDB
		var offsetAPI orm.OffsetAPI

		// insertion point for updating fields
		offsetAPI.ID = offsetDB.ID
		offsetDB.CopyBasicFieldsToOffset_WOP(&offsetAPI.Offset_WOP)
		offsetAPI.OffsetPointersEncoding = offsetDB.OffsetPointersEncoding
		offsetAPIs = append(offsetAPIs, offsetAPI)
	}

	c.JSON(http.StatusOK, offsetAPIs)
}

// PostOffset
//
// swagger:route POST /offsets offsets postOffset
//
// Creates a offset
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOffset(c *gin.Context) {

	mutexOffset.Lock()
	defer mutexOffset.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOffsets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOffset.GetDB()

	// Validate input
	var input orm.OffsetAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create offset
	offsetDB := orm.OffsetDB{}
	offsetDB.OffsetPointersEncoding = input.OffsetPointersEncoding
	offsetDB.CopyBasicFieldsFromOffset_WOP(&input.Offset_WOP)

	query := db.Create(&offsetDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOffset.CheckoutPhaseOneInstance(&offsetDB)
	offset := backRepo.BackRepoOffset.Map_OffsetDBID_OffsetPtr[offsetDB.ID]

	if offset != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), offset)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, offsetDB)
}

// GetOffset
//
// swagger:route GET /offsets/{ID} offsets getOffset
//
// Gets the details for a offset.
//
// Responses:
// default: genericError
//
//	200: offsetDBResponse
func (controller *Controller) GetOffset(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOffset", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOffset.GetDB()

	// Get offsetDB in DB
	var offsetDB orm.OffsetDB
	if err := db.First(&offsetDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var offsetAPI orm.OffsetAPI
	offsetAPI.ID = offsetDB.ID
	offsetAPI.OffsetPointersEncoding = offsetDB.OffsetPointersEncoding
	offsetDB.CopyBasicFieldsToOffset_WOP(&offsetAPI.Offset_WOP)

	c.JSON(http.StatusOK, offsetAPI)
}

// UpdateOffset
//
// swagger:route PATCH /offsets/{ID} offsets updateOffset
//
// # Update a offset
//
// Responses:
// default: genericError
//
//	200: offsetDBResponse
func (controller *Controller) UpdateOffset(c *gin.Context) {

	mutexOffset.Lock()
	defer mutexOffset.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOffset", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOffset.GetDB()

	// Validate input
	var input orm.OffsetAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var offsetDB orm.OffsetDB

	// fetch the offset
	query := db.First(&offsetDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	offsetDB.CopyBasicFieldsFromOffset_WOP(&input.Offset_WOP)
	offsetDB.OffsetPointersEncoding = input.OffsetPointersEncoding

	query = db.Model(&offsetDB).Updates(offsetDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	offsetNew := new(models.Offset)
	offsetDB.CopyBasicFieldsToOffset(offsetNew)

	// redeem pointers
	offsetDB.DecodePointers(backRepo, offsetNew)

	// get stage instance from DB instance, and call callback function
	offsetOld := backRepo.BackRepoOffset.Map_OffsetDBID_OffsetPtr[offsetDB.ID]
	if offsetOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), offsetOld, offsetNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the offsetDB
	c.JSON(http.StatusOK, offsetDB)
}

// DeleteOffset
//
// swagger:route DELETE /offsets/{ID} offsets deleteOffset
//
// # Delete a offset
//
// default: genericError
//
//	200: offsetDBResponse
func (controller *Controller) DeleteOffset(c *gin.Context) {

	mutexOffset.Lock()
	defer mutexOffset.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOffset", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOffset.GetDB()

	// Get model if exist
	var offsetDB orm.OffsetDB
	if err := db.First(&offsetDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&offsetDB)

	// get an instance (not staged) from DB instance, and call callback function
	offsetDeleted := new(models.Offset)
	offsetDB.CopyBasicFieldsToOffset(offsetDeleted)

	// get stage instance from DB instance, and call callback function
	offsetStaged := backRepo.BackRepoOffset.Map_OffsetDBID_OffsetPtr[offsetDB.ID]
	if offsetStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), offsetStaged, offsetDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
