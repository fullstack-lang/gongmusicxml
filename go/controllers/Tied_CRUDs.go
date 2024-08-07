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
var __Tied__dummysDeclaration__ models.Tied
var __Tied_time__dummyDeclaration time.Duration

var mutexTied sync.Mutex

// An TiedID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTied updateTied deleteTied
type TiedID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TiedInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTied updateTied
type TiedInput struct {
	// The Tied to submit or modify
	// in: body
	Tied *orm.TiedAPI
}

// GetTieds
//
// swagger:route GET /tieds tieds getTieds
//
// # Get all tieds
//
// Responses:
// default: genericError
//
//	200: tiedDBResponse
func (controller *Controller) GetTieds(c *gin.Context) {

	// source slice
	var tiedDBs []orm.TiedDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTieds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTied.GetDB()

	query := db.Find(&tiedDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tiedAPIs := make([]orm.TiedAPI, 0)

	// for each tied, update fields from the database nullable fields
	for idx := range tiedDBs {
		tiedDB := &tiedDBs[idx]
		_ = tiedDB
		var tiedAPI orm.TiedAPI

		// insertion point for updating fields
		tiedAPI.ID = tiedDB.ID
		tiedDB.CopyBasicFieldsToTied_WOP(&tiedAPI.Tied_WOP)
		tiedAPI.TiedPointersEncoding = tiedDB.TiedPointersEncoding
		tiedAPIs = append(tiedAPIs, tiedAPI)
	}

	c.JSON(http.StatusOK, tiedAPIs)
}

// PostTied
//
// swagger:route POST /tieds tieds postTied
//
// Creates a tied
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTied(c *gin.Context) {

	mutexTied.Lock()
	defer mutexTied.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTieds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTied.GetDB()

	// Validate input
	var input orm.TiedAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tied
	tiedDB := orm.TiedDB{}
	tiedDB.TiedPointersEncoding = input.TiedPointersEncoding
	tiedDB.CopyBasicFieldsFromTied_WOP(&input.Tied_WOP)

	query := db.Create(&tiedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTied.CheckoutPhaseOneInstance(&tiedDB)
	tied := backRepo.BackRepoTied.Map_TiedDBID_TiedPtr[tiedDB.ID]

	if tied != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tied)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tiedDB)
}

// GetTied
//
// swagger:route GET /tieds/{ID} tieds getTied
//
// Gets the details for a tied.
//
// Responses:
// default: genericError
//
//	200: tiedDBResponse
func (controller *Controller) GetTied(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTied", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTied.GetDB()

	// Get tiedDB in DB
	var tiedDB orm.TiedDB
	if err := db.First(&tiedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tiedAPI orm.TiedAPI
	tiedAPI.ID = tiedDB.ID
	tiedAPI.TiedPointersEncoding = tiedDB.TiedPointersEncoding
	tiedDB.CopyBasicFieldsToTied_WOP(&tiedAPI.Tied_WOP)

	c.JSON(http.StatusOK, tiedAPI)
}

// UpdateTied
//
// swagger:route PATCH /tieds/{ID} tieds updateTied
//
// # Update a tied
//
// Responses:
// default: genericError
//
//	200: tiedDBResponse
func (controller *Controller) UpdateTied(c *gin.Context) {

	mutexTied.Lock()
	defer mutexTied.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTied", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTied.GetDB()

	// Validate input
	var input orm.TiedAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tiedDB orm.TiedDB

	// fetch the tied
	query := db.First(&tiedDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tiedDB.CopyBasicFieldsFromTied_WOP(&input.Tied_WOP)
	tiedDB.TiedPointersEncoding = input.TiedPointersEncoding

	query = db.Model(&tiedDB).Updates(tiedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tiedNew := new(models.Tied)
	tiedDB.CopyBasicFieldsToTied(tiedNew)

	// redeem pointers
	tiedDB.DecodePointers(backRepo, tiedNew)

	// get stage instance from DB instance, and call callback function
	tiedOld := backRepo.BackRepoTied.Map_TiedDBID_TiedPtr[tiedDB.ID]
	if tiedOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tiedOld, tiedNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tiedDB
	c.JSON(http.StatusOK, tiedDB)
}

// DeleteTied
//
// swagger:route DELETE /tieds/{ID} tieds deleteTied
//
// # Delete a tied
//
// default: genericError
//
//	200: tiedDBResponse
func (controller *Controller) DeleteTied(c *gin.Context) {

	mutexTied.Lock()
	defer mutexTied.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTied", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTied.GetDB()

	// Get model if exist
	var tiedDB orm.TiedDB
	if err := db.First(&tiedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tiedDB)

	// get an instance (not staged) from DB instance, and call callback function
	tiedDeleted := new(models.Tied)
	tiedDB.CopyBasicFieldsToTied(tiedDeleted)

	// get stage instance from DB instance, and call callback function
	tiedStaged := backRepo.BackRepoTied.Map_TiedDBID_TiedPtr[tiedDB.ID]
	if tiedStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tiedStaged, tiedDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
