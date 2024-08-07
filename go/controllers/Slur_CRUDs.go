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
var __Slur__dummysDeclaration__ models.Slur
var __Slur_time__dummyDeclaration time.Duration

var mutexSlur sync.Mutex

// An SlurID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSlur updateSlur deleteSlur
type SlurID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SlurInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSlur updateSlur
type SlurInput struct {
	// The Slur to submit or modify
	// in: body
	Slur *orm.SlurAPI
}

// GetSlurs
//
// swagger:route GET /slurs slurs getSlurs
//
// # Get all slurs
//
// Responses:
// default: genericError
//
//	200: slurDBResponse
func (controller *Controller) GetSlurs(c *gin.Context) {

	// source slice
	var slurDBs []orm.SlurDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSlurs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlur.GetDB()

	query := db.Find(&slurDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	slurAPIs := make([]orm.SlurAPI, 0)

	// for each slur, update fields from the database nullable fields
	for idx := range slurDBs {
		slurDB := &slurDBs[idx]
		_ = slurDB
		var slurAPI orm.SlurAPI

		// insertion point for updating fields
		slurAPI.ID = slurDB.ID
		slurDB.CopyBasicFieldsToSlur_WOP(&slurAPI.Slur_WOP)
		slurAPI.SlurPointersEncoding = slurDB.SlurPointersEncoding
		slurAPIs = append(slurAPIs, slurAPI)
	}

	c.JSON(http.StatusOK, slurAPIs)
}

// PostSlur
//
// swagger:route POST /slurs slurs postSlur
//
// Creates a slur
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSlur(c *gin.Context) {

	mutexSlur.Lock()
	defer mutexSlur.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSlurs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlur.GetDB()

	// Validate input
	var input orm.SlurAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create slur
	slurDB := orm.SlurDB{}
	slurDB.SlurPointersEncoding = input.SlurPointersEncoding
	slurDB.CopyBasicFieldsFromSlur_WOP(&input.Slur_WOP)

	query := db.Create(&slurDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSlur.CheckoutPhaseOneInstance(&slurDB)
	slur := backRepo.BackRepoSlur.Map_SlurDBID_SlurPtr[slurDB.ID]

	if slur != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), slur)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, slurDB)
}

// GetSlur
//
// swagger:route GET /slurs/{ID} slurs getSlur
//
// Gets the details for a slur.
//
// Responses:
// default: genericError
//
//	200: slurDBResponse
func (controller *Controller) GetSlur(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSlur", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlur.GetDB()

	// Get slurDB in DB
	var slurDB orm.SlurDB
	if err := db.First(&slurDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var slurAPI orm.SlurAPI
	slurAPI.ID = slurDB.ID
	slurAPI.SlurPointersEncoding = slurDB.SlurPointersEncoding
	slurDB.CopyBasicFieldsToSlur_WOP(&slurAPI.Slur_WOP)

	c.JSON(http.StatusOK, slurAPI)
}

// UpdateSlur
//
// swagger:route PATCH /slurs/{ID} slurs updateSlur
//
// # Update a slur
//
// Responses:
// default: genericError
//
//	200: slurDBResponse
func (controller *Controller) UpdateSlur(c *gin.Context) {

	mutexSlur.Lock()
	defer mutexSlur.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSlur", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlur.GetDB()

	// Validate input
	var input orm.SlurAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var slurDB orm.SlurDB

	// fetch the slur
	query := db.First(&slurDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	slurDB.CopyBasicFieldsFromSlur_WOP(&input.Slur_WOP)
	slurDB.SlurPointersEncoding = input.SlurPointersEncoding

	query = db.Model(&slurDB).Updates(slurDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	slurNew := new(models.Slur)
	slurDB.CopyBasicFieldsToSlur(slurNew)

	// redeem pointers
	slurDB.DecodePointers(backRepo, slurNew)

	// get stage instance from DB instance, and call callback function
	slurOld := backRepo.BackRepoSlur.Map_SlurDBID_SlurPtr[slurDB.ID]
	if slurOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), slurOld, slurNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the slurDB
	c.JSON(http.StatusOK, slurDB)
}

// DeleteSlur
//
// swagger:route DELETE /slurs/{ID} slurs deleteSlur
//
// # Delete a slur
//
// default: genericError
//
//	200: slurDBResponse
func (controller *Controller) DeleteSlur(c *gin.Context) {

	mutexSlur.Lock()
	defer mutexSlur.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSlur", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlur.GetDB()

	// Get model if exist
	var slurDB orm.SlurDB
	if err := db.First(&slurDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&slurDB)

	// get an instance (not staged) from DB instance, and call callback function
	slurDeleted := new(models.Slur)
	slurDB.CopyBasicFieldsToSlur(slurDeleted)

	// get stage instance from DB instance, and call callback function
	slurStaged := backRepo.BackRepoSlur.Map_SlurDBID_SlurPtr[slurDB.ID]
	if slurStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), slurStaged, slurDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
