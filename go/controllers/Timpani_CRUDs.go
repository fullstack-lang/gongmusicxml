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
var __Timpani__dummysDeclaration__ models.Timpani
var __Timpani_time__dummyDeclaration time.Duration

var mutexTimpani sync.Mutex

// An TimpaniID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTimpani updateTimpani deleteTimpani
type TimpaniID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TimpaniInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTimpani updateTimpani
type TimpaniInput struct {
	// The Timpani to submit or modify
	// in: body
	Timpani *orm.TimpaniAPI
}

// GetTimpanis
//
// swagger:route GET /timpanis timpanis getTimpanis
//
// # Get all timpanis
//
// Responses:
// default: genericError
//
//	200: timpaniDBResponse
func (controller *Controller) GetTimpanis(c *gin.Context) {

	// source slice
	var timpaniDBs []orm.TimpaniDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTimpanis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTimpani.GetDB()

	query := db.Find(&timpaniDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	timpaniAPIs := make([]orm.TimpaniAPI, 0)

	// for each timpani, update fields from the database nullable fields
	for idx := range timpaniDBs {
		timpaniDB := &timpaniDBs[idx]
		_ = timpaniDB
		var timpaniAPI orm.TimpaniAPI

		// insertion point for updating fields
		timpaniAPI.ID = timpaniDB.ID
		timpaniDB.CopyBasicFieldsToTimpani_WOP(&timpaniAPI.Timpani_WOP)
		timpaniAPI.TimpaniPointersEncoding = timpaniDB.TimpaniPointersEncoding
		timpaniAPIs = append(timpaniAPIs, timpaniAPI)
	}

	c.JSON(http.StatusOK, timpaniAPIs)
}

// PostTimpani
//
// swagger:route POST /timpanis timpanis postTimpani
//
// Creates a timpani
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTimpani(c *gin.Context) {

	mutexTimpani.Lock()
	defer mutexTimpani.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTimpanis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTimpani.GetDB()

	// Validate input
	var input orm.TimpaniAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create timpani
	timpaniDB := orm.TimpaniDB{}
	timpaniDB.TimpaniPointersEncoding = input.TimpaniPointersEncoding
	timpaniDB.CopyBasicFieldsFromTimpani_WOP(&input.Timpani_WOP)

	query := db.Create(&timpaniDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTimpani.CheckoutPhaseOneInstance(&timpaniDB)
	timpani := backRepo.BackRepoTimpani.Map_TimpaniDBID_TimpaniPtr[timpaniDB.ID]

	if timpani != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), timpani)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, timpaniDB)
}

// GetTimpani
//
// swagger:route GET /timpanis/{ID} timpanis getTimpani
//
// Gets the details for a timpani.
//
// Responses:
// default: genericError
//
//	200: timpaniDBResponse
func (controller *Controller) GetTimpani(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTimpani", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTimpani.GetDB()

	// Get timpaniDB in DB
	var timpaniDB orm.TimpaniDB
	if err := db.First(&timpaniDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var timpaniAPI orm.TimpaniAPI
	timpaniAPI.ID = timpaniDB.ID
	timpaniAPI.TimpaniPointersEncoding = timpaniDB.TimpaniPointersEncoding
	timpaniDB.CopyBasicFieldsToTimpani_WOP(&timpaniAPI.Timpani_WOP)

	c.JSON(http.StatusOK, timpaniAPI)
}

// UpdateTimpani
//
// swagger:route PATCH /timpanis/{ID} timpanis updateTimpani
//
// # Update a timpani
//
// Responses:
// default: genericError
//
//	200: timpaniDBResponse
func (controller *Controller) UpdateTimpani(c *gin.Context) {

	mutexTimpani.Lock()
	defer mutexTimpani.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTimpani", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTimpani.GetDB()

	// Validate input
	var input orm.TimpaniAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var timpaniDB orm.TimpaniDB

	// fetch the timpani
	query := db.First(&timpaniDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	timpaniDB.CopyBasicFieldsFromTimpani_WOP(&input.Timpani_WOP)
	timpaniDB.TimpaniPointersEncoding = input.TimpaniPointersEncoding

	query = db.Model(&timpaniDB).Updates(timpaniDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	timpaniNew := new(models.Timpani)
	timpaniDB.CopyBasicFieldsToTimpani(timpaniNew)

	// redeem pointers
	timpaniDB.DecodePointers(backRepo, timpaniNew)

	// get stage instance from DB instance, and call callback function
	timpaniOld := backRepo.BackRepoTimpani.Map_TimpaniDBID_TimpaniPtr[timpaniDB.ID]
	if timpaniOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), timpaniOld, timpaniNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the timpaniDB
	c.JSON(http.StatusOK, timpaniDB)
}

// DeleteTimpani
//
// swagger:route DELETE /timpanis/{ID} timpanis deleteTimpani
//
// # Delete a timpani
//
// default: genericError
//
//	200: timpaniDBResponse
func (controller *Controller) DeleteTimpani(c *gin.Context) {

	mutexTimpani.Lock()
	defer mutexTimpani.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTimpani", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTimpani.GetDB()

	// Get model if exist
	var timpaniDB orm.TimpaniDB
	if err := db.First(&timpaniDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&timpaniDB)

	// get an instance (not staged) from DB instance, and call callback function
	timpaniDeleted := new(models.Timpani)
	timpaniDB.CopyBasicFieldsToTimpani(timpaniDeleted)

	// get stage instance from DB instance, and call callback function
	timpaniStaged := backRepo.BackRepoTimpani.Map_TimpaniDBID_TimpaniPtr[timpaniDB.ID]
	if timpaniStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), timpaniStaged, timpaniDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
