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
var __Accidental_text__dummysDeclaration__ models.Accidental_text
var __Accidental_text_time__dummyDeclaration time.Duration

var mutexAccidental_text sync.Mutex

// An Accidental_textID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAccidental_text updateAccidental_text deleteAccidental_text
type Accidental_textID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Accidental_textInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAccidental_text updateAccidental_text
type Accidental_textInput struct {
	// The Accidental_text to submit or modify
	// in: body
	Accidental_text *orm.Accidental_textAPI
}

// GetAccidental_texts
//
// swagger:route GET /accidental_texts accidental_texts getAccidental_texts
//
// # Get all accidental_texts
//
// Responses:
// default: genericError
//
//	200: accidental_textDBResponse
func (controller *Controller) GetAccidental_texts(c *gin.Context) {

	// source slice
	var accidental_textDBs []orm.Accidental_textDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccidental_texts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_text.GetDB()

	query := db.Find(&accidental_textDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	accidental_textAPIs := make([]orm.Accidental_textAPI, 0)

	// for each accidental_text, update fields from the database nullable fields
	for idx := range accidental_textDBs {
		accidental_textDB := &accidental_textDBs[idx]
		_ = accidental_textDB
		var accidental_textAPI orm.Accidental_textAPI

		// insertion point for updating fields
		accidental_textAPI.ID = accidental_textDB.ID
		accidental_textDB.CopyBasicFieldsToAccidental_text_WOP(&accidental_textAPI.Accidental_text_WOP)
		accidental_textAPI.Accidental_textPointersEncoding = accidental_textDB.Accidental_textPointersEncoding
		accidental_textAPIs = append(accidental_textAPIs, accidental_textAPI)
	}

	c.JSON(http.StatusOK, accidental_textAPIs)
}

// PostAccidental_text
//
// swagger:route POST /accidental_texts accidental_texts postAccidental_text
//
// Creates a accidental_text
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAccidental_text(c *gin.Context) {

	mutexAccidental_text.Lock()
	defer mutexAccidental_text.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAccidental_texts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_text.GetDB()

	// Validate input
	var input orm.Accidental_textAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create accidental_text
	accidental_textDB := orm.Accidental_textDB{}
	accidental_textDB.Accidental_textPointersEncoding = input.Accidental_textPointersEncoding
	accidental_textDB.CopyBasicFieldsFromAccidental_text_WOP(&input.Accidental_text_WOP)

	query := db.Create(&accidental_textDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAccidental_text.CheckoutPhaseOneInstance(&accidental_textDB)
	accidental_text := backRepo.BackRepoAccidental_text.Map_Accidental_textDBID_Accidental_textPtr[accidental_textDB.ID]

	if accidental_text != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), accidental_text)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, accidental_textDB)
}

// GetAccidental_text
//
// swagger:route GET /accidental_texts/{ID} accidental_texts getAccidental_text
//
// Gets the details for a accidental_text.
//
// Responses:
// default: genericError
//
//	200: accidental_textDBResponse
func (controller *Controller) GetAccidental_text(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccidental_text", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_text.GetDB()

	// Get accidental_textDB in DB
	var accidental_textDB orm.Accidental_textDB
	if err := db.First(&accidental_textDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var accidental_textAPI orm.Accidental_textAPI
	accidental_textAPI.ID = accidental_textDB.ID
	accidental_textAPI.Accidental_textPointersEncoding = accidental_textDB.Accidental_textPointersEncoding
	accidental_textDB.CopyBasicFieldsToAccidental_text_WOP(&accidental_textAPI.Accidental_text_WOP)

	c.JSON(http.StatusOK, accidental_textAPI)
}

// UpdateAccidental_text
//
// swagger:route PATCH /accidental_texts/{ID} accidental_texts updateAccidental_text
//
// # Update a accidental_text
//
// Responses:
// default: genericError
//
//	200: accidental_textDBResponse
func (controller *Controller) UpdateAccidental_text(c *gin.Context) {

	mutexAccidental_text.Lock()
	defer mutexAccidental_text.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAccidental_text", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_text.GetDB()

	// Validate input
	var input orm.Accidental_textAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var accidental_textDB orm.Accidental_textDB

	// fetch the accidental_text
	query := db.First(&accidental_textDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	accidental_textDB.CopyBasicFieldsFromAccidental_text_WOP(&input.Accidental_text_WOP)
	accidental_textDB.Accidental_textPointersEncoding = input.Accidental_textPointersEncoding

	query = db.Model(&accidental_textDB).Updates(accidental_textDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	accidental_textNew := new(models.Accidental_text)
	accidental_textDB.CopyBasicFieldsToAccidental_text(accidental_textNew)

	// redeem pointers
	accidental_textDB.DecodePointers(backRepo, accidental_textNew)

	// get stage instance from DB instance, and call callback function
	accidental_textOld := backRepo.BackRepoAccidental_text.Map_Accidental_textDBID_Accidental_textPtr[accidental_textDB.ID]
	if accidental_textOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), accidental_textOld, accidental_textNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the accidental_textDB
	c.JSON(http.StatusOK, accidental_textDB)
}

// DeleteAccidental_text
//
// swagger:route DELETE /accidental_texts/{ID} accidental_texts deleteAccidental_text
//
// # Delete a accidental_text
//
// default: genericError
//
//	200: accidental_textDBResponse
func (controller *Controller) DeleteAccidental_text(c *gin.Context) {

	mutexAccidental_text.Lock()
	defer mutexAccidental_text.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAccidental_text", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccidental_text.GetDB()

	// Get model if exist
	var accidental_textDB orm.Accidental_textDB
	if err := db.First(&accidental_textDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&accidental_textDB)

	// get an instance (not staged) from DB instance, and call callback function
	accidental_textDeleted := new(models.Accidental_text)
	accidental_textDB.CopyBasicFieldsToAccidental_text(accidental_textDeleted)

	// get stage instance from DB instance, and call callback function
	accidental_textStaged := backRepo.BackRepoAccidental_text.Map_Accidental_textDBID_Accidental_textPtr[accidental_textDB.ID]
	if accidental_textStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), accidental_textStaged, accidental_textDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
