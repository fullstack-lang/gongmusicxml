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
var __Forward__dummysDeclaration__ models.Forward
var __Forward_time__dummyDeclaration time.Duration

var mutexForward sync.Mutex

// An ForwardID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getForward updateForward deleteForward
type ForwardID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ForwardInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postForward updateForward
type ForwardInput struct {
	// The Forward to submit or modify
	// in: body
	Forward *orm.ForwardAPI
}

// GetForwards
//
// swagger:route GET /forwards forwards getForwards
//
// # Get all forwards
//
// Responses:
// default: genericError
//
//	200: forwardDBResponse
func (controller *Controller) GetForwards(c *gin.Context) {

	// source slice
	var forwardDBs []orm.ForwardDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetForwards", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoForward.GetDB()

	query := db.Find(&forwardDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	forwardAPIs := make([]orm.ForwardAPI, 0)

	// for each forward, update fields from the database nullable fields
	for idx := range forwardDBs {
		forwardDB := &forwardDBs[idx]
		_ = forwardDB
		var forwardAPI orm.ForwardAPI

		// insertion point for updating fields
		forwardAPI.ID = forwardDB.ID
		forwardDB.CopyBasicFieldsToForward_WOP(&forwardAPI.Forward_WOP)
		forwardAPI.ForwardPointersEncoding = forwardDB.ForwardPointersEncoding
		forwardAPIs = append(forwardAPIs, forwardAPI)
	}

	c.JSON(http.StatusOK, forwardAPIs)
}

// PostForward
//
// swagger:route POST /forwards forwards postForward
//
// Creates a forward
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostForward(c *gin.Context) {

	mutexForward.Lock()
	defer mutexForward.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostForwards", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoForward.GetDB()

	// Validate input
	var input orm.ForwardAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create forward
	forwardDB := orm.ForwardDB{}
	forwardDB.ForwardPointersEncoding = input.ForwardPointersEncoding
	forwardDB.CopyBasicFieldsFromForward_WOP(&input.Forward_WOP)

	query := db.Create(&forwardDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoForward.CheckoutPhaseOneInstance(&forwardDB)
	forward := backRepo.BackRepoForward.Map_ForwardDBID_ForwardPtr[forwardDB.ID]

	if forward != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), forward)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, forwardDB)
}

// GetForward
//
// swagger:route GET /forwards/{ID} forwards getForward
//
// Gets the details for a forward.
//
// Responses:
// default: genericError
//
//	200: forwardDBResponse
func (controller *Controller) GetForward(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetForward", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoForward.GetDB()

	// Get forwardDB in DB
	var forwardDB orm.ForwardDB
	if err := db.First(&forwardDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var forwardAPI orm.ForwardAPI
	forwardAPI.ID = forwardDB.ID
	forwardAPI.ForwardPointersEncoding = forwardDB.ForwardPointersEncoding
	forwardDB.CopyBasicFieldsToForward_WOP(&forwardAPI.Forward_WOP)

	c.JSON(http.StatusOK, forwardAPI)
}

// UpdateForward
//
// swagger:route PATCH /forwards/{ID} forwards updateForward
//
// # Update a forward
//
// Responses:
// default: genericError
//
//	200: forwardDBResponse
func (controller *Controller) UpdateForward(c *gin.Context) {

	mutexForward.Lock()
	defer mutexForward.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateForward", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoForward.GetDB()

	// Validate input
	var input orm.ForwardAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var forwardDB orm.ForwardDB

	// fetch the forward
	query := db.First(&forwardDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	forwardDB.CopyBasicFieldsFromForward_WOP(&input.Forward_WOP)
	forwardDB.ForwardPointersEncoding = input.ForwardPointersEncoding

	query = db.Model(&forwardDB).Updates(forwardDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	forwardNew := new(models.Forward)
	forwardDB.CopyBasicFieldsToForward(forwardNew)

	// redeem pointers
	forwardDB.DecodePointers(backRepo, forwardNew)

	// get stage instance from DB instance, and call callback function
	forwardOld := backRepo.BackRepoForward.Map_ForwardDBID_ForwardPtr[forwardDB.ID]
	if forwardOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), forwardOld, forwardNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the forwardDB
	c.JSON(http.StatusOK, forwardDB)
}

// DeleteForward
//
// swagger:route DELETE /forwards/{ID} forwards deleteForward
//
// # Delete a forward
//
// default: genericError
//
//	200: forwardDBResponse
func (controller *Controller) DeleteForward(c *gin.Context) {

	mutexForward.Lock()
	defer mutexForward.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteForward", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoForward.GetDB()

	// Get model if exist
	var forwardDB orm.ForwardDB
	if err := db.First(&forwardDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&forwardDB)

	// get an instance (not staged) from DB instance, and call callback function
	forwardDeleted := new(models.Forward)
	forwardDB.CopyBasicFieldsToForward(forwardDeleted)

	// get stage instance from DB instance, and call callback function
	forwardStaged := backRepo.BackRepoForward.Map_ForwardDBID_ForwardPtr[forwardDB.ID]
	if forwardStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), forwardStaged, forwardDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
