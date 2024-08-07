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
var __First_fret__dummysDeclaration__ models.First_fret
var __First_fret_time__dummyDeclaration time.Duration

var mutexFirst_fret sync.Mutex

// An First_fretID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFirst_fret updateFirst_fret deleteFirst_fret
type First_fretID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// First_fretInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFirst_fret updateFirst_fret
type First_fretInput struct {
	// The First_fret to submit or modify
	// in: body
	First_fret *orm.First_fretAPI
}

// GetFirst_frets
//
// swagger:route GET /first_frets first_frets getFirst_frets
//
// # Get all first_frets
//
// Responses:
// default: genericError
//
//	200: first_fretDBResponse
func (controller *Controller) GetFirst_frets(c *gin.Context) {

	// source slice
	var first_fretDBs []orm.First_fretDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFirst_frets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFirst_fret.GetDB()

	query := db.Find(&first_fretDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	first_fretAPIs := make([]orm.First_fretAPI, 0)

	// for each first_fret, update fields from the database nullable fields
	for idx := range first_fretDBs {
		first_fretDB := &first_fretDBs[idx]
		_ = first_fretDB
		var first_fretAPI orm.First_fretAPI

		// insertion point for updating fields
		first_fretAPI.ID = first_fretDB.ID
		first_fretDB.CopyBasicFieldsToFirst_fret_WOP(&first_fretAPI.First_fret_WOP)
		first_fretAPI.First_fretPointersEncoding = first_fretDB.First_fretPointersEncoding
		first_fretAPIs = append(first_fretAPIs, first_fretAPI)
	}

	c.JSON(http.StatusOK, first_fretAPIs)
}

// PostFirst_fret
//
// swagger:route POST /first_frets first_frets postFirst_fret
//
// Creates a first_fret
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFirst_fret(c *gin.Context) {

	mutexFirst_fret.Lock()
	defer mutexFirst_fret.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFirst_frets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFirst_fret.GetDB()

	// Validate input
	var input orm.First_fretAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create first_fret
	first_fretDB := orm.First_fretDB{}
	first_fretDB.First_fretPointersEncoding = input.First_fretPointersEncoding
	first_fretDB.CopyBasicFieldsFromFirst_fret_WOP(&input.First_fret_WOP)

	query := db.Create(&first_fretDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFirst_fret.CheckoutPhaseOneInstance(&first_fretDB)
	first_fret := backRepo.BackRepoFirst_fret.Map_First_fretDBID_First_fretPtr[first_fretDB.ID]

	if first_fret != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), first_fret)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, first_fretDB)
}

// GetFirst_fret
//
// swagger:route GET /first_frets/{ID} first_frets getFirst_fret
//
// Gets the details for a first_fret.
//
// Responses:
// default: genericError
//
//	200: first_fretDBResponse
func (controller *Controller) GetFirst_fret(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFirst_fret", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFirst_fret.GetDB()

	// Get first_fretDB in DB
	var first_fretDB orm.First_fretDB
	if err := db.First(&first_fretDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var first_fretAPI orm.First_fretAPI
	first_fretAPI.ID = first_fretDB.ID
	first_fretAPI.First_fretPointersEncoding = first_fretDB.First_fretPointersEncoding
	first_fretDB.CopyBasicFieldsToFirst_fret_WOP(&first_fretAPI.First_fret_WOP)

	c.JSON(http.StatusOK, first_fretAPI)
}

// UpdateFirst_fret
//
// swagger:route PATCH /first_frets/{ID} first_frets updateFirst_fret
//
// # Update a first_fret
//
// Responses:
// default: genericError
//
//	200: first_fretDBResponse
func (controller *Controller) UpdateFirst_fret(c *gin.Context) {

	mutexFirst_fret.Lock()
	defer mutexFirst_fret.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFirst_fret", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFirst_fret.GetDB()

	// Validate input
	var input orm.First_fretAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var first_fretDB orm.First_fretDB

	// fetch the first_fret
	query := db.First(&first_fretDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	first_fretDB.CopyBasicFieldsFromFirst_fret_WOP(&input.First_fret_WOP)
	first_fretDB.First_fretPointersEncoding = input.First_fretPointersEncoding

	query = db.Model(&first_fretDB).Updates(first_fretDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	first_fretNew := new(models.First_fret)
	first_fretDB.CopyBasicFieldsToFirst_fret(first_fretNew)

	// redeem pointers
	first_fretDB.DecodePointers(backRepo, first_fretNew)

	// get stage instance from DB instance, and call callback function
	first_fretOld := backRepo.BackRepoFirst_fret.Map_First_fretDBID_First_fretPtr[first_fretDB.ID]
	if first_fretOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), first_fretOld, first_fretNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the first_fretDB
	c.JSON(http.StatusOK, first_fretDB)
}

// DeleteFirst_fret
//
// swagger:route DELETE /first_frets/{ID} first_frets deleteFirst_fret
//
// # Delete a first_fret
//
// default: genericError
//
//	200: first_fretDBResponse
func (controller *Controller) DeleteFirst_fret(c *gin.Context) {

	mutexFirst_fret.Lock()
	defer mutexFirst_fret.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFirst_fret", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFirst_fret.GetDB()

	// Get model if exist
	var first_fretDB orm.First_fretDB
	if err := db.First(&first_fretDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&first_fretDB)

	// get an instance (not staged) from DB instance, and call callback function
	first_fretDeleted := new(models.First_fret)
	first_fretDB.CopyBasicFieldsToFirst_fret(first_fretDeleted)

	// get stage instance from DB instance, and call callback function
	first_fretStaged := backRepo.BackRepoFirst_fret.Map_First_fretDBID_First_fretPtr[first_fretDB.ID]
	if first_fretStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), first_fretStaged, first_fretDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
