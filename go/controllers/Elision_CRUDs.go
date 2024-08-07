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
var __Elision__dummysDeclaration__ models.Elision
var __Elision_time__dummyDeclaration time.Duration

var mutexElision sync.Mutex

// An ElisionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getElision updateElision deleteElision
type ElisionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ElisionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postElision updateElision
type ElisionInput struct {
	// The Elision to submit or modify
	// in: body
	Elision *orm.ElisionAPI
}

// GetElisions
//
// swagger:route GET /elisions elisions getElisions
//
// # Get all elisions
//
// Responses:
// default: genericError
//
//	200: elisionDBResponse
func (controller *Controller) GetElisions(c *gin.Context) {

	// source slice
	var elisionDBs []orm.ElisionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetElisions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElision.GetDB()

	query := db.Find(&elisionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	elisionAPIs := make([]orm.ElisionAPI, 0)

	// for each elision, update fields from the database nullable fields
	for idx := range elisionDBs {
		elisionDB := &elisionDBs[idx]
		_ = elisionDB
		var elisionAPI orm.ElisionAPI

		// insertion point for updating fields
		elisionAPI.ID = elisionDB.ID
		elisionDB.CopyBasicFieldsToElision_WOP(&elisionAPI.Elision_WOP)
		elisionAPI.ElisionPointersEncoding = elisionDB.ElisionPointersEncoding
		elisionAPIs = append(elisionAPIs, elisionAPI)
	}

	c.JSON(http.StatusOK, elisionAPIs)
}

// PostElision
//
// swagger:route POST /elisions elisions postElision
//
// Creates a elision
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostElision(c *gin.Context) {

	mutexElision.Lock()
	defer mutexElision.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostElisions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElision.GetDB()

	// Validate input
	var input orm.ElisionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create elision
	elisionDB := orm.ElisionDB{}
	elisionDB.ElisionPointersEncoding = input.ElisionPointersEncoding
	elisionDB.CopyBasicFieldsFromElision_WOP(&input.Elision_WOP)

	query := db.Create(&elisionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoElision.CheckoutPhaseOneInstance(&elisionDB)
	elision := backRepo.BackRepoElision.Map_ElisionDBID_ElisionPtr[elisionDB.ID]

	if elision != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), elision)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, elisionDB)
}

// GetElision
//
// swagger:route GET /elisions/{ID} elisions getElision
//
// Gets the details for a elision.
//
// Responses:
// default: genericError
//
//	200: elisionDBResponse
func (controller *Controller) GetElision(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetElision", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElision.GetDB()

	// Get elisionDB in DB
	var elisionDB orm.ElisionDB
	if err := db.First(&elisionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var elisionAPI orm.ElisionAPI
	elisionAPI.ID = elisionDB.ID
	elisionAPI.ElisionPointersEncoding = elisionDB.ElisionPointersEncoding
	elisionDB.CopyBasicFieldsToElision_WOP(&elisionAPI.Elision_WOP)

	c.JSON(http.StatusOK, elisionAPI)
}

// UpdateElision
//
// swagger:route PATCH /elisions/{ID} elisions updateElision
//
// # Update a elision
//
// Responses:
// default: genericError
//
//	200: elisionDBResponse
func (controller *Controller) UpdateElision(c *gin.Context) {

	mutexElision.Lock()
	defer mutexElision.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateElision", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElision.GetDB()

	// Validate input
	var input orm.ElisionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var elisionDB orm.ElisionDB

	// fetch the elision
	query := db.First(&elisionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	elisionDB.CopyBasicFieldsFromElision_WOP(&input.Elision_WOP)
	elisionDB.ElisionPointersEncoding = input.ElisionPointersEncoding

	query = db.Model(&elisionDB).Updates(elisionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	elisionNew := new(models.Elision)
	elisionDB.CopyBasicFieldsToElision(elisionNew)

	// redeem pointers
	elisionDB.DecodePointers(backRepo, elisionNew)

	// get stage instance from DB instance, and call callback function
	elisionOld := backRepo.BackRepoElision.Map_ElisionDBID_ElisionPtr[elisionDB.ID]
	if elisionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), elisionOld, elisionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the elisionDB
	c.JSON(http.StatusOK, elisionDB)
}

// DeleteElision
//
// swagger:route DELETE /elisions/{ID} elisions deleteElision
//
// # Delete a elision
//
// default: genericError
//
//	200: elisionDBResponse
func (controller *Controller) DeleteElision(c *gin.Context) {

	mutexElision.Lock()
	defer mutexElision.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteElision", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElision.GetDB()

	// Get model if exist
	var elisionDB orm.ElisionDB
	if err := db.First(&elisionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&elisionDB)

	// get an instance (not staged) from DB instance, and call callback function
	elisionDeleted := new(models.Elision)
	elisionDB.CopyBasicFieldsToElision(elisionDeleted)

	// get stage instance from DB instance, and call callback function
	elisionStaged := backRepo.BackRepoElision.Map_ElisionDBID_ElisionPtr[elisionDB.ID]
	if elisionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), elisionStaged, elisionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
