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
var __Clef__dummysDeclaration__ models.Clef
var __Clef_time__dummyDeclaration time.Duration

var mutexClef sync.Mutex

// An ClefID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getClef updateClef deleteClef
type ClefID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ClefInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postClef updateClef
type ClefInput struct {
	// The Clef to submit or modify
	// in: body
	Clef *orm.ClefAPI
}

// GetClefs
//
// swagger:route GET /clefs clefs getClefs
//
// # Get all clefs
//
// Responses:
// default: genericError
//
//	200: clefDBResponse
func (controller *Controller) GetClefs(c *gin.Context) {

	// source slice
	var clefDBs []orm.ClefDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetClefs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoClef.GetDB()

	query := db.Find(&clefDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	clefAPIs := make([]orm.ClefAPI, 0)

	// for each clef, update fields from the database nullable fields
	for idx := range clefDBs {
		clefDB := &clefDBs[idx]
		_ = clefDB
		var clefAPI orm.ClefAPI

		// insertion point for updating fields
		clefAPI.ID = clefDB.ID
		clefDB.CopyBasicFieldsToClef_WOP(&clefAPI.Clef_WOP)
		clefAPI.ClefPointersEncoding = clefDB.ClefPointersEncoding
		clefAPIs = append(clefAPIs, clefAPI)
	}

	c.JSON(http.StatusOK, clefAPIs)
}

// PostClef
//
// swagger:route POST /clefs clefs postClef
//
// Creates a clef
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostClef(c *gin.Context) {

	mutexClef.Lock()
	defer mutexClef.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostClefs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoClef.GetDB()

	// Validate input
	var input orm.ClefAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create clef
	clefDB := orm.ClefDB{}
	clefDB.ClefPointersEncoding = input.ClefPointersEncoding
	clefDB.CopyBasicFieldsFromClef_WOP(&input.Clef_WOP)

	query := db.Create(&clefDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoClef.CheckoutPhaseOneInstance(&clefDB)
	clef := backRepo.BackRepoClef.Map_ClefDBID_ClefPtr[clefDB.ID]

	if clef != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), clef)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, clefDB)
}

// GetClef
//
// swagger:route GET /clefs/{ID} clefs getClef
//
// Gets the details for a clef.
//
// Responses:
// default: genericError
//
//	200: clefDBResponse
func (controller *Controller) GetClef(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetClef", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoClef.GetDB()

	// Get clefDB in DB
	var clefDB orm.ClefDB
	if err := db.First(&clefDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var clefAPI orm.ClefAPI
	clefAPI.ID = clefDB.ID
	clefAPI.ClefPointersEncoding = clefDB.ClefPointersEncoding
	clefDB.CopyBasicFieldsToClef_WOP(&clefAPI.Clef_WOP)

	c.JSON(http.StatusOK, clefAPI)
}

// UpdateClef
//
// swagger:route PATCH /clefs/{ID} clefs updateClef
//
// # Update a clef
//
// Responses:
// default: genericError
//
//	200: clefDBResponse
func (controller *Controller) UpdateClef(c *gin.Context) {

	mutexClef.Lock()
	defer mutexClef.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateClef", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoClef.GetDB()

	// Validate input
	var input orm.ClefAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var clefDB orm.ClefDB

	// fetch the clef
	query := db.First(&clefDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	clefDB.CopyBasicFieldsFromClef_WOP(&input.Clef_WOP)
	clefDB.ClefPointersEncoding = input.ClefPointersEncoding

	query = db.Model(&clefDB).Updates(clefDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	clefNew := new(models.Clef)
	clefDB.CopyBasicFieldsToClef(clefNew)

	// redeem pointers
	clefDB.DecodePointers(backRepo, clefNew)

	// get stage instance from DB instance, and call callback function
	clefOld := backRepo.BackRepoClef.Map_ClefDBID_ClefPtr[clefDB.ID]
	if clefOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), clefOld, clefNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the clefDB
	c.JSON(http.StatusOK, clefDB)
}

// DeleteClef
//
// swagger:route DELETE /clefs/{ID} clefs deleteClef
//
// # Delete a clef
//
// default: genericError
//
//	200: clefDBResponse
func (controller *Controller) DeleteClef(c *gin.Context) {

	mutexClef.Lock()
	defer mutexClef.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteClef", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoClef.GetDB()

	// Get model if exist
	var clefDB orm.ClefDB
	if err := db.First(&clefDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&clefDB)

	// get an instance (not staged) from DB instance, and call callback function
	clefDeleted := new(models.Clef)
	clefDB.CopyBasicFieldsToClef(clefDeleted)

	// get stage instance from DB instance, and call callback function
	clefStaged := backRepo.BackRepoClef.Map_ClefDBID_ClefPtr[clefDB.ID]
	if clefStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), clefStaged, clefDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
