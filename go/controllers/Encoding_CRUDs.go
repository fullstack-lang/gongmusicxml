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
var __Encoding__dummysDeclaration__ models.Encoding
var __Encoding_time__dummyDeclaration time.Duration

var mutexEncoding sync.Mutex

// An EncodingID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEncoding updateEncoding deleteEncoding
type EncodingID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EncodingInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEncoding updateEncoding
type EncodingInput struct {
	// The Encoding to submit or modify
	// in: body
	Encoding *orm.EncodingAPI
}

// GetEncodings
//
// swagger:route GET /encodings encodings getEncodings
//
// # Get all encodings
//
// Responses:
// default: genericError
//
//	200: encodingDBResponse
func (controller *Controller) GetEncodings(c *gin.Context) {

	// source slice
	var encodingDBs []orm.EncodingDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEncodings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEncoding.GetDB()

	query := db.Find(&encodingDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	encodingAPIs := make([]orm.EncodingAPI, 0)

	// for each encoding, update fields from the database nullable fields
	for idx := range encodingDBs {
		encodingDB := &encodingDBs[idx]
		_ = encodingDB
		var encodingAPI orm.EncodingAPI

		// insertion point for updating fields
		encodingAPI.ID = encodingDB.ID
		encodingDB.CopyBasicFieldsToEncoding_WOP(&encodingAPI.Encoding_WOP)
		encodingAPI.EncodingPointersEncoding = encodingDB.EncodingPointersEncoding
		encodingAPIs = append(encodingAPIs, encodingAPI)
	}

	c.JSON(http.StatusOK, encodingAPIs)
}

// PostEncoding
//
// swagger:route POST /encodings encodings postEncoding
//
// Creates a encoding
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEncoding(c *gin.Context) {

	mutexEncoding.Lock()
	defer mutexEncoding.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEncodings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEncoding.GetDB()

	// Validate input
	var input orm.EncodingAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create encoding
	encodingDB := orm.EncodingDB{}
	encodingDB.EncodingPointersEncoding = input.EncodingPointersEncoding
	encodingDB.CopyBasicFieldsFromEncoding_WOP(&input.Encoding_WOP)

	query := db.Create(&encodingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEncoding.CheckoutPhaseOneInstance(&encodingDB)
	encoding := backRepo.BackRepoEncoding.Map_EncodingDBID_EncodingPtr[encodingDB.ID]

	if encoding != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), encoding)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, encodingDB)
}

// GetEncoding
//
// swagger:route GET /encodings/{ID} encodings getEncoding
//
// Gets the details for a encoding.
//
// Responses:
// default: genericError
//
//	200: encodingDBResponse
func (controller *Controller) GetEncoding(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEncoding", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEncoding.GetDB()

	// Get encodingDB in DB
	var encodingDB orm.EncodingDB
	if err := db.First(&encodingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var encodingAPI orm.EncodingAPI
	encodingAPI.ID = encodingDB.ID
	encodingAPI.EncodingPointersEncoding = encodingDB.EncodingPointersEncoding
	encodingDB.CopyBasicFieldsToEncoding_WOP(&encodingAPI.Encoding_WOP)

	c.JSON(http.StatusOK, encodingAPI)
}

// UpdateEncoding
//
// swagger:route PATCH /encodings/{ID} encodings updateEncoding
//
// # Update a encoding
//
// Responses:
// default: genericError
//
//	200: encodingDBResponse
func (controller *Controller) UpdateEncoding(c *gin.Context) {

	mutexEncoding.Lock()
	defer mutexEncoding.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEncoding", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEncoding.GetDB()

	// Validate input
	var input orm.EncodingAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var encodingDB orm.EncodingDB

	// fetch the encoding
	query := db.First(&encodingDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	encodingDB.CopyBasicFieldsFromEncoding_WOP(&input.Encoding_WOP)
	encodingDB.EncodingPointersEncoding = input.EncodingPointersEncoding

	query = db.Model(&encodingDB).Updates(encodingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	encodingNew := new(models.Encoding)
	encodingDB.CopyBasicFieldsToEncoding(encodingNew)

	// redeem pointers
	encodingDB.DecodePointers(backRepo, encodingNew)

	// get stage instance from DB instance, and call callback function
	encodingOld := backRepo.BackRepoEncoding.Map_EncodingDBID_EncodingPtr[encodingDB.ID]
	if encodingOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), encodingOld, encodingNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the encodingDB
	c.JSON(http.StatusOK, encodingDB)
}

// DeleteEncoding
//
// swagger:route DELETE /encodings/{ID} encodings deleteEncoding
//
// # Delete a encoding
//
// default: genericError
//
//	200: encodingDBResponse
func (controller *Controller) DeleteEncoding(c *gin.Context) {

	mutexEncoding.Lock()
	defer mutexEncoding.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEncoding", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEncoding.GetDB()

	// Get model if exist
	var encodingDB orm.EncodingDB
	if err := db.First(&encodingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&encodingDB)

	// get an instance (not staged) from DB instance, and call callback function
	encodingDeleted := new(models.Encoding)
	encodingDB.CopyBasicFieldsToEncoding(encodingDeleted)

	// get stage instance from DB instance, and call callback function
	encodingStaged := backRepo.BackRepoEncoding.Map_EncodingDBID_EncodingPtr[encodingDB.ID]
	if encodingStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), encodingStaged, encodingDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
