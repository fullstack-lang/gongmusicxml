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
var __Notehead_text__dummysDeclaration__ models.Notehead_text
var __Notehead_text_time__dummyDeclaration time.Duration

var mutexNotehead_text sync.Mutex

// An Notehead_textID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNotehead_text updateNotehead_text deleteNotehead_text
type Notehead_textID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Notehead_textInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNotehead_text updateNotehead_text
type Notehead_textInput struct {
	// The Notehead_text to submit or modify
	// in: body
	Notehead_text *orm.Notehead_textAPI
}

// GetNotehead_texts
//
// swagger:route GET /notehead_texts notehead_texts getNotehead_texts
//
// # Get all notehead_texts
//
// Responses:
// default: genericError
//
//	200: notehead_textDBResponse
func (controller *Controller) GetNotehead_texts(c *gin.Context) {

	// source slice
	var notehead_textDBs []orm.Notehead_textDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNotehead_texts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead_text.GetDB()

	query := db.Find(&notehead_textDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	notehead_textAPIs := make([]orm.Notehead_textAPI, 0)

	// for each notehead_text, update fields from the database nullable fields
	for idx := range notehead_textDBs {
		notehead_textDB := &notehead_textDBs[idx]
		_ = notehead_textDB
		var notehead_textAPI orm.Notehead_textAPI

		// insertion point for updating fields
		notehead_textAPI.ID = notehead_textDB.ID
		notehead_textDB.CopyBasicFieldsToNotehead_text_WOP(&notehead_textAPI.Notehead_text_WOP)
		notehead_textAPI.Notehead_textPointersEncoding = notehead_textDB.Notehead_textPointersEncoding
		notehead_textAPIs = append(notehead_textAPIs, notehead_textAPI)
	}

	c.JSON(http.StatusOK, notehead_textAPIs)
}

// PostNotehead_text
//
// swagger:route POST /notehead_texts notehead_texts postNotehead_text
//
// Creates a notehead_text
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNotehead_text(c *gin.Context) {

	mutexNotehead_text.Lock()
	defer mutexNotehead_text.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNotehead_texts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead_text.GetDB()

	// Validate input
	var input orm.Notehead_textAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create notehead_text
	notehead_textDB := orm.Notehead_textDB{}
	notehead_textDB.Notehead_textPointersEncoding = input.Notehead_textPointersEncoding
	notehead_textDB.CopyBasicFieldsFromNotehead_text_WOP(&input.Notehead_text_WOP)

	query := db.Create(&notehead_textDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNotehead_text.CheckoutPhaseOneInstance(&notehead_textDB)
	notehead_text := backRepo.BackRepoNotehead_text.Map_Notehead_textDBID_Notehead_textPtr[notehead_textDB.ID]

	if notehead_text != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), notehead_text)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, notehead_textDB)
}

// GetNotehead_text
//
// swagger:route GET /notehead_texts/{ID} notehead_texts getNotehead_text
//
// Gets the details for a notehead_text.
//
// Responses:
// default: genericError
//
//	200: notehead_textDBResponse
func (controller *Controller) GetNotehead_text(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNotehead_text", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead_text.GetDB()

	// Get notehead_textDB in DB
	var notehead_textDB orm.Notehead_textDB
	if err := db.First(&notehead_textDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var notehead_textAPI orm.Notehead_textAPI
	notehead_textAPI.ID = notehead_textDB.ID
	notehead_textAPI.Notehead_textPointersEncoding = notehead_textDB.Notehead_textPointersEncoding
	notehead_textDB.CopyBasicFieldsToNotehead_text_WOP(&notehead_textAPI.Notehead_text_WOP)

	c.JSON(http.StatusOK, notehead_textAPI)
}

// UpdateNotehead_text
//
// swagger:route PATCH /notehead_texts/{ID} notehead_texts updateNotehead_text
//
// # Update a notehead_text
//
// Responses:
// default: genericError
//
//	200: notehead_textDBResponse
func (controller *Controller) UpdateNotehead_text(c *gin.Context) {

	mutexNotehead_text.Lock()
	defer mutexNotehead_text.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNotehead_text", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead_text.GetDB()

	// Validate input
	var input orm.Notehead_textAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var notehead_textDB orm.Notehead_textDB

	// fetch the notehead_text
	query := db.First(&notehead_textDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	notehead_textDB.CopyBasicFieldsFromNotehead_text_WOP(&input.Notehead_text_WOP)
	notehead_textDB.Notehead_textPointersEncoding = input.Notehead_textPointersEncoding

	query = db.Model(&notehead_textDB).Updates(notehead_textDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	notehead_textNew := new(models.Notehead_text)
	notehead_textDB.CopyBasicFieldsToNotehead_text(notehead_textNew)

	// redeem pointers
	notehead_textDB.DecodePointers(backRepo, notehead_textNew)

	// get stage instance from DB instance, and call callback function
	notehead_textOld := backRepo.BackRepoNotehead_text.Map_Notehead_textDBID_Notehead_textPtr[notehead_textDB.ID]
	if notehead_textOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), notehead_textOld, notehead_textNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the notehead_textDB
	c.JSON(http.StatusOK, notehead_textDB)
}

// DeleteNotehead_text
//
// swagger:route DELETE /notehead_texts/{ID} notehead_texts deleteNotehead_text
//
// # Delete a notehead_text
//
// default: genericError
//
//	200: notehead_textDBResponse
func (controller *Controller) DeleteNotehead_text(c *gin.Context) {

	mutexNotehead_text.Lock()
	defer mutexNotehead_text.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNotehead_text", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotehead_text.GetDB()

	// Get model if exist
	var notehead_textDB orm.Notehead_textDB
	if err := db.First(&notehead_textDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&notehead_textDB)

	// get an instance (not staged) from DB instance, and call callback function
	notehead_textDeleted := new(models.Notehead_text)
	notehead_textDB.CopyBasicFieldsToNotehead_text(notehead_textDeleted)

	// get stage instance from DB instance, and call callback function
	notehead_textStaged := backRepo.BackRepoNotehead_text.Map_Notehead_textDBID_Notehead_textPtr[notehead_textDB.ID]
	if notehead_textStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), notehead_textStaged, notehead_textDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
