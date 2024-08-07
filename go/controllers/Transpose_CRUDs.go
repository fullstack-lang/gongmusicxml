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
var __Transpose__dummysDeclaration__ models.Transpose
var __Transpose_time__dummyDeclaration time.Duration

var mutexTranspose sync.Mutex

// An TransposeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTranspose updateTranspose deleteTranspose
type TransposeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TransposeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTranspose updateTranspose
type TransposeInput struct {
	// The Transpose to submit or modify
	// in: body
	Transpose *orm.TransposeAPI
}

// GetTransposes
//
// swagger:route GET /transposes transposes getTransposes
//
// # Get all transposes
//
// Responses:
// default: genericError
//
//	200: transposeDBResponse
func (controller *Controller) GetTransposes(c *gin.Context) {

	// source slice
	var transposeDBs []orm.TransposeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTransposes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTranspose.GetDB()

	query := db.Find(&transposeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	transposeAPIs := make([]orm.TransposeAPI, 0)

	// for each transpose, update fields from the database nullable fields
	for idx := range transposeDBs {
		transposeDB := &transposeDBs[idx]
		_ = transposeDB
		var transposeAPI orm.TransposeAPI

		// insertion point for updating fields
		transposeAPI.ID = transposeDB.ID
		transposeDB.CopyBasicFieldsToTranspose_WOP(&transposeAPI.Transpose_WOP)
		transposeAPI.TransposePointersEncoding = transposeDB.TransposePointersEncoding
		transposeAPIs = append(transposeAPIs, transposeAPI)
	}

	c.JSON(http.StatusOK, transposeAPIs)
}

// PostTranspose
//
// swagger:route POST /transposes transposes postTranspose
//
// Creates a transpose
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTranspose(c *gin.Context) {

	mutexTranspose.Lock()
	defer mutexTranspose.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTransposes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTranspose.GetDB()

	// Validate input
	var input orm.TransposeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create transpose
	transposeDB := orm.TransposeDB{}
	transposeDB.TransposePointersEncoding = input.TransposePointersEncoding
	transposeDB.CopyBasicFieldsFromTranspose_WOP(&input.Transpose_WOP)

	query := db.Create(&transposeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTranspose.CheckoutPhaseOneInstance(&transposeDB)
	transpose := backRepo.BackRepoTranspose.Map_TransposeDBID_TransposePtr[transposeDB.ID]

	if transpose != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), transpose)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, transposeDB)
}

// GetTranspose
//
// swagger:route GET /transposes/{ID} transposes getTranspose
//
// Gets the details for a transpose.
//
// Responses:
// default: genericError
//
//	200: transposeDBResponse
func (controller *Controller) GetTranspose(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTranspose", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTranspose.GetDB()

	// Get transposeDB in DB
	var transposeDB orm.TransposeDB
	if err := db.First(&transposeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var transposeAPI orm.TransposeAPI
	transposeAPI.ID = transposeDB.ID
	transposeAPI.TransposePointersEncoding = transposeDB.TransposePointersEncoding
	transposeDB.CopyBasicFieldsToTranspose_WOP(&transposeAPI.Transpose_WOP)

	c.JSON(http.StatusOK, transposeAPI)
}

// UpdateTranspose
//
// swagger:route PATCH /transposes/{ID} transposes updateTranspose
//
// # Update a transpose
//
// Responses:
// default: genericError
//
//	200: transposeDBResponse
func (controller *Controller) UpdateTranspose(c *gin.Context) {

	mutexTranspose.Lock()
	defer mutexTranspose.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTranspose", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTranspose.GetDB()

	// Validate input
	var input orm.TransposeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var transposeDB orm.TransposeDB

	// fetch the transpose
	query := db.First(&transposeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	transposeDB.CopyBasicFieldsFromTranspose_WOP(&input.Transpose_WOP)
	transposeDB.TransposePointersEncoding = input.TransposePointersEncoding

	query = db.Model(&transposeDB).Updates(transposeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	transposeNew := new(models.Transpose)
	transposeDB.CopyBasicFieldsToTranspose(transposeNew)

	// redeem pointers
	transposeDB.DecodePointers(backRepo, transposeNew)

	// get stage instance from DB instance, and call callback function
	transposeOld := backRepo.BackRepoTranspose.Map_TransposeDBID_TransposePtr[transposeDB.ID]
	if transposeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), transposeOld, transposeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the transposeDB
	c.JSON(http.StatusOK, transposeDB)
}

// DeleteTranspose
//
// swagger:route DELETE /transposes/{ID} transposes deleteTranspose
//
// # Delete a transpose
//
// default: genericError
//
//	200: transposeDBResponse
func (controller *Controller) DeleteTranspose(c *gin.Context) {

	mutexTranspose.Lock()
	defer mutexTranspose.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTranspose", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTranspose.GetDB()

	// Get model if exist
	var transposeDB orm.TransposeDB
	if err := db.First(&transposeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&transposeDB)

	// get an instance (not staged) from DB instance, and call callback function
	transposeDeleted := new(models.Transpose)
	transposeDB.CopyBasicFieldsToTranspose(transposeDeleted)

	// get stage instance from DB instance, and call callback function
	transposeStaged := backRepo.BackRepoTranspose.Map_TransposeDBID_TransposePtr[transposeDB.ID]
	if transposeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), transposeStaged, transposeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
