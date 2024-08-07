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
var __Typed_text__dummysDeclaration__ models.Typed_text
var __Typed_text_time__dummyDeclaration time.Duration

var mutexTyped_text sync.Mutex

// An Typed_textID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTyped_text updateTyped_text deleteTyped_text
type Typed_textID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Typed_textInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTyped_text updateTyped_text
type Typed_textInput struct {
	// The Typed_text to submit or modify
	// in: body
	Typed_text *orm.Typed_textAPI
}

// GetTyped_texts
//
// swagger:route GET /typed_texts typed_texts getTyped_texts
//
// # Get all typed_texts
//
// Responses:
// default: genericError
//
//	200: typed_textDBResponse
func (controller *Controller) GetTyped_texts(c *gin.Context) {

	// source slice
	var typed_textDBs []orm.Typed_textDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTyped_texts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTyped_text.GetDB()

	query := db.Find(&typed_textDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	typed_textAPIs := make([]orm.Typed_textAPI, 0)

	// for each typed_text, update fields from the database nullable fields
	for idx := range typed_textDBs {
		typed_textDB := &typed_textDBs[idx]
		_ = typed_textDB
		var typed_textAPI orm.Typed_textAPI

		// insertion point for updating fields
		typed_textAPI.ID = typed_textDB.ID
		typed_textDB.CopyBasicFieldsToTyped_text_WOP(&typed_textAPI.Typed_text_WOP)
		typed_textAPI.Typed_textPointersEncoding = typed_textDB.Typed_textPointersEncoding
		typed_textAPIs = append(typed_textAPIs, typed_textAPI)
	}

	c.JSON(http.StatusOK, typed_textAPIs)
}

// PostTyped_text
//
// swagger:route POST /typed_texts typed_texts postTyped_text
//
// Creates a typed_text
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTyped_text(c *gin.Context) {

	mutexTyped_text.Lock()
	defer mutexTyped_text.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTyped_texts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTyped_text.GetDB()

	// Validate input
	var input orm.Typed_textAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create typed_text
	typed_textDB := orm.Typed_textDB{}
	typed_textDB.Typed_textPointersEncoding = input.Typed_textPointersEncoding
	typed_textDB.CopyBasicFieldsFromTyped_text_WOP(&input.Typed_text_WOP)

	query := db.Create(&typed_textDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTyped_text.CheckoutPhaseOneInstance(&typed_textDB)
	typed_text := backRepo.BackRepoTyped_text.Map_Typed_textDBID_Typed_textPtr[typed_textDB.ID]

	if typed_text != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), typed_text)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, typed_textDB)
}

// GetTyped_text
//
// swagger:route GET /typed_texts/{ID} typed_texts getTyped_text
//
// Gets the details for a typed_text.
//
// Responses:
// default: genericError
//
//	200: typed_textDBResponse
func (controller *Controller) GetTyped_text(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTyped_text", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTyped_text.GetDB()

	// Get typed_textDB in DB
	var typed_textDB orm.Typed_textDB
	if err := db.First(&typed_textDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var typed_textAPI orm.Typed_textAPI
	typed_textAPI.ID = typed_textDB.ID
	typed_textAPI.Typed_textPointersEncoding = typed_textDB.Typed_textPointersEncoding
	typed_textDB.CopyBasicFieldsToTyped_text_WOP(&typed_textAPI.Typed_text_WOP)

	c.JSON(http.StatusOK, typed_textAPI)
}

// UpdateTyped_text
//
// swagger:route PATCH /typed_texts/{ID} typed_texts updateTyped_text
//
// # Update a typed_text
//
// Responses:
// default: genericError
//
//	200: typed_textDBResponse
func (controller *Controller) UpdateTyped_text(c *gin.Context) {

	mutexTyped_text.Lock()
	defer mutexTyped_text.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTyped_text", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTyped_text.GetDB()

	// Validate input
	var input orm.Typed_textAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var typed_textDB orm.Typed_textDB

	// fetch the typed_text
	query := db.First(&typed_textDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	typed_textDB.CopyBasicFieldsFromTyped_text_WOP(&input.Typed_text_WOP)
	typed_textDB.Typed_textPointersEncoding = input.Typed_textPointersEncoding

	query = db.Model(&typed_textDB).Updates(typed_textDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	typed_textNew := new(models.Typed_text)
	typed_textDB.CopyBasicFieldsToTyped_text(typed_textNew)

	// redeem pointers
	typed_textDB.DecodePointers(backRepo, typed_textNew)

	// get stage instance from DB instance, and call callback function
	typed_textOld := backRepo.BackRepoTyped_text.Map_Typed_textDBID_Typed_textPtr[typed_textDB.ID]
	if typed_textOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), typed_textOld, typed_textNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the typed_textDB
	c.JSON(http.StatusOK, typed_textDB)
}

// DeleteTyped_text
//
// swagger:route DELETE /typed_texts/{ID} typed_texts deleteTyped_text
//
// # Delete a typed_text
//
// default: genericError
//
//	200: typed_textDBResponse
func (controller *Controller) DeleteTyped_text(c *gin.Context) {

	mutexTyped_text.Lock()
	defer mutexTyped_text.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTyped_text", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTyped_text.GetDB()

	// Get model if exist
	var typed_textDB orm.Typed_textDB
	if err := db.First(&typed_textDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&typed_textDB)

	// get an instance (not staged) from DB instance, and call callback function
	typed_textDeleted := new(models.Typed_text)
	typed_textDB.CopyBasicFieldsToTyped_text(typed_textDeleted)

	// get stage instance from DB instance, and call callback function
	typed_textStaged := backRepo.BackRepoTyped_text.Map_Typed_textDBID_Typed_textPtr[typed_textDB.ID]
	if typed_textStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), typed_textStaged, typed_textDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
