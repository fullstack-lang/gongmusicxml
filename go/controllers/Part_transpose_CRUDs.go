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
var __Part_transpose__dummysDeclaration__ models.Part_transpose
var __Part_transpose_time__dummyDeclaration time.Duration

var mutexPart_transpose sync.Mutex

// An Part_transposeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPart_transpose updatePart_transpose deletePart_transpose
type Part_transposeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Part_transposeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPart_transpose updatePart_transpose
type Part_transposeInput struct {
	// The Part_transpose to submit or modify
	// in: body
	Part_transpose *orm.Part_transposeAPI
}

// GetPart_transposes
//
// swagger:route GET /part_transposes part_transposes getPart_transposes
//
// # Get all part_transposes
//
// Responses:
// default: genericError
//
//	200: part_transposeDBResponse
func (controller *Controller) GetPart_transposes(c *gin.Context) {

	// source slice
	var part_transposeDBs []orm.Part_transposeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_transposes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_transpose.GetDB()

	query := db.Find(&part_transposeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	part_transposeAPIs := make([]orm.Part_transposeAPI, 0)

	// for each part_transpose, update fields from the database nullable fields
	for idx := range part_transposeDBs {
		part_transposeDB := &part_transposeDBs[idx]
		_ = part_transposeDB
		var part_transposeAPI orm.Part_transposeAPI

		// insertion point for updating fields
		part_transposeAPI.ID = part_transposeDB.ID
		part_transposeDB.CopyBasicFieldsToPart_transpose_WOP(&part_transposeAPI.Part_transpose_WOP)
		part_transposeAPI.Part_transposePointersEncoding = part_transposeDB.Part_transposePointersEncoding
		part_transposeAPIs = append(part_transposeAPIs, part_transposeAPI)
	}

	c.JSON(http.StatusOK, part_transposeAPIs)
}

// PostPart_transpose
//
// swagger:route POST /part_transposes part_transposes postPart_transpose
//
// Creates a part_transpose
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPart_transpose(c *gin.Context) {

	mutexPart_transpose.Lock()
	defer mutexPart_transpose.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPart_transposes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_transpose.GetDB()

	// Validate input
	var input orm.Part_transposeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create part_transpose
	part_transposeDB := orm.Part_transposeDB{}
	part_transposeDB.Part_transposePointersEncoding = input.Part_transposePointersEncoding
	part_transposeDB.CopyBasicFieldsFromPart_transpose_WOP(&input.Part_transpose_WOP)

	query := db.Create(&part_transposeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPart_transpose.CheckoutPhaseOneInstance(&part_transposeDB)
	part_transpose := backRepo.BackRepoPart_transpose.Map_Part_transposeDBID_Part_transposePtr[part_transposeDB.ID]

	if part_transpose != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), part_transpose)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, part_transposeDB)
}

// GetPart_transpose
//
// swagger:route GET /part_transposes/{ID} part_transposes getPart_transpose
//
// Gets the details for a part_transpose.
//
// Responses:
// default: genericError
//
//	200: part_transposeDBResponse
func (controller *Controller) GetPart_transpose(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_transpose", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_transpose.GetDB()

	// Get part_transposeDB in DB
	var part_transposeDB orm.Part_transposeDB
	if err := db.First(&part_transposeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var part_transposeAPI orm.Part_transposeAPI
	part_transposeAPI.ID = part_transposeDB.ID
	part_transposeAPI.Part_transposePointersEncoding = part_transposeDB.Part_transposePointersEncoding
	part_transposeDB.CopyBasicFieldsToPart_transpose_WOP(&part_transposeAPI.Part_transpose_WOP)

	c.JSON(http.StatusOK, part_transposeAPI)
}

// UpdatePart_transpose
//
// swagger:route PATCH /part_transposes/{ID} part_transposes updatePart_transpose
//
// # Update a part_transpose
//
// Responses:
// default: genericError
//
//	200: part_transposeDBResponse
func (controller *Controller) UpdatePart_transpose(c *gin.Context) {

	mutexPart_transpose.Lock()
	defer mutexPart_transpose.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePart_transpose", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_transpose.GetDB()

	// Validate input
	var input orm.Part_transposeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var part_transposeDB orm.Part_transposeDB

	// fetch the part_transpose
	query := db.First(&part_transposeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	part_transposeDB.CopyBasicFieldsFromPart_transpose_WOP(&input.Part_transpose_WOP)
	part_transposeDB.Part_transposePointersEncoding = input.Part_transposePointersEncoding

	query = db.Model(&part_transposeDB).Updates(part_transposeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	part_transposeNew := new(models.Part_transpose)
	part_transposeDB.CopyBasicFieldsToPart_transpose(part_transposeNew)

	// redeem pointers
	part_transposeDB.DecodePointers(backRepo, part_transposeNew)

	// get stage instance from DB instance, and call callback function
	part_transposeOld := backRepo.BackRepoPart_transpose.Map_Part_transposeDBID_Part_transposePtr[part_transposeDB.ID]
	if part_transposeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), part_transposeOld, part_transposeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the part_transposeDB
	c.JSON(http.StatusOK, part_transposeDB)
}

// DeletePart_transpose
//
// swagger:route DELETE /part_transposes/{ID} part_transposes deletePart_transpose
//
// # Delete a part_transpose
//
// default: genericError
//
//	200: part_transposeDBResponse
func (controller *Controller) DeletePart_transpose(c *gin.Context) {

	mutexPart_transpose.Lock()
	defer mutexPart_transpose.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePart_transpose", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_transpose.GetDB()

	// Get model if exist
	var part_transposeDB orm.Part_transposeDB
	if err := db.First(&part_transposeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&part_transposeDB)

	// get an instance (not staged) from DB instance, and call callback function
	part_transposeDeleted := new(models.Part_transpose)
	part_transposeDB.CopyBasicFieldsToPart_transpose(part_transposeDeleted)

	// get stage instance from DB instance, and call callback function
	part_transposeStaged := backRepo.BackRepoPart_transpose.Map_Part_transposeDBID_Part_transposePtr[part_transposeDB.ID]
	if part_transposeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), part_transposeStaged, part_transposeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
