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
var __Tuplet_dot__dummysDeclaration__ models.Tuplet_dot
var __Tuplet_dot_time__dummyDeclaration time.Duration

var mutexTuplet_dot sync.Mutex

// An Tuplet_dotID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTuplet_dot updateTuplet_dot deleteTuplet_dot
type Tuplet_dotID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Tuplet_dotInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTuplet_dot updateTuplet_dot
type Tuplet_dotInput struct {
	// The Tuplet_dot to submit or modify
	// in: body
	Tuplet_dot *orm.Tuplet_dotAPI
}

// GetTuplet_dots
//
// swagger:route GET /tuplet_dots tuplet_dots getTuplet_dots
//
// # Get all tuplet_dots
//
// Responses:
// default: genericError
//
//	200: tuplet_dotDBResponse
func (controller *Controller) GetTuplet_dots(c *gin.Context) {

	// source slice
	var tuplet_dotDBs []orm.Tuplet_dotDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplet_dots", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_dot.GetDB()

	query := db.Find(&tuplet_dotDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tuplet_dotAPIs := make([]orm.Tuplet_dotAPI, 0)

	// for each tuplet_dot, update fields from the database nullable fields
	for idx := range tuplet_dotDBs {
		tuplet_dotDB := &tuplet_dotDBs[idx]
		_ = tuplet_dotDB
		var tuplet_dotAPI orm.Tuplet_dotAPI

		// insertion point for updating fields
		tuplet_dotAPI.ID = tuplet_dotDB.ID
		tuplet_dotDB.CopyBasicFieldsToTuplet_dot_WOP(&tuplet_dotAPI.Tuplet_dot_WOP)
		tuplet_dotAPI.Tuplet_dotPointersEncoding = tuplet_dotDB.Tuplet_dotPointersEncoding
		tuplet_dotAPIs = append(tuplet_dotAPIs, tuplet_dotAPI)
	}

	c.JSON(http.StatusOK, tuplet_dotAPIs)
}

// PostTuplet_dot
//
// swagger:route POST /tuplet_dots tuplet_dots postTuplet_dot
//
// Creates a tuplet_dot
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTuplet_dot(c *gin.Context) {

	mutexTuplet_dot.Lock()
	defer mutexTuplet_dot.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTuplet_dots", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_dot.GetDB()

	// Validate input
	var input orm.Tuplet_dotAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tuplet_dot
	tuplet_dotDB := orm.Tuplet_dotDB{}
	tuplet_dotDB.Tuplet_dotPointersEncoding = input.Tuplet_dotPointersEncoding
	tuplet_dotDB.CopyBasicFieldsFromTuplet_dot_WOP(&input.Tuplet_dot_WOP)

	query := db.Create(&tuplet_dotDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTuplet_dot.CheckoutPhaseOneInstance(&tuplet_dotDB)
	tuplet_dot := backRepo.BackRepoTuplet_dot.Map_Tuplet_dotDBID_Tuplet_dotPtr[tuplet_dotDB.ID]

	if tuplet_dot != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tuplet_dot)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tuplet_dotDB)
}

// GetTuplet_dot
//
// swagger:route GET /tuplet_dots/{ID} tuplet_dots getTuplet_dot
//
// Gets the details for a tuplet_dot.
//
// Responses:
// default: genericError
//
//	200: tuplet_dotDBResponse
func (controller *Controller) GetTuplet_dot(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplet_dot", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_dot.GetDB()

	// Get tuplet_dotDB in DB
	var tuplet_dotDB orm.Tuplet_dotDB
	if err := db.First(&tuplet_dotDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tuplet_dotAPI orm.Tuplet_dotAPI
	tuplet_dotAPI.ID = tuplet_dotDB.ID
	tuplet_dotAPI.Tuplet_dotPointersEncoding = tuplet_dotDB.Tuplet_dotPointersEncoding
	tuplet_dotDB.CopyBasicFieldsToTuplet_dot_WOP(&tuplet_dotAPI.Tuplet_dot_WOP)

	c.JSON(http.StatusOK, tuplet_dotAPI)
}

// UpdateTuplet_dot
//
// swagger:route PATCH /tuplet_dots/{ID} tuplet_dots updateTuplet_dot
//
// # Update a tuplet_dot
//
// Responses:
// default: genericError
//
//	200: tuplet_dotDBResponse
func (controller *Controller) UpdateTuplet_dot(c *gin.Context) {

	mutexTuplet_dot.Lock()
	defer mutexTuplet_dot.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTuplet_dot", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_dot.GetDB()

	// Validate input
	var input orm.Tuplet_dotAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tuplet_dotDB orm.Tuplet_dotDB

	// fetch the tuplet_dot
	query := db.First(&tuplet_dotDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tuplet_dotDB.CopyBasicFieldsFromTuplet_dot_WOP(&input.Tuplet_dot_WOP)
	tuplet_dotDB.Tuplet_dotPointersEncoding = input.Tuplet_dotPointersEncoding

	query = db.Model(&tuplet_dotDB).Updates(tuplet_dotDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tuplet_dotNew := new(models.Tuplet_dot)
	tuplet_dotDB.CopyBasicFieldsToTuplet_dot(tuplet_dotNew)

	// redeem pointers
	tuplet_dotDB.DecodePointers(backRepo, tuplet_dotNew)

	// get stage instance from DB instance, and call callback function
	tuplet_dotOld := backRepo.BackRepoTuplet_dot.Map_Tuplet_dotDBID_Tuplet_dotPtr[tuplet_dotDB.ID]
	if tuplet_dotOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tuplet_dotOld, tuplet_dotNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tuplet_dotDB
	c.JSON(http.StatusOK, tuplet_dotDB)
}

// DeleteTuplet_dot
//
// swagger:route DELETE /tuplet_dots/{ID} tuplet_dots deleteTuplet_dot
//
// # Delete a tuplet_dot
//
// default: genericError
//
//	200: tuplet_dotDBResponse
func (controller *Controller) DeleteTuplet_dot(c *gin.Context) {

	mutexTuplet_dot.Lock()
	defer mutexTuplet_dot.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTuplet_dot", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_dot.GetDB()

	// Get model if exist
	var tuplet_dotDB orm.Tuplet_dotDB
	if err := db.First(&tuplet_dotDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tuplet_dotDB)

	// get an instance (not staged) from DB instance, and call callback function
	tuplet_dotDeleted := new(models.Tuplet_dot)
	tuplet_dotDB.CopyBasicFieldsToTuplet_dot(tuplet_dotDeleted)

	// get stage instance from DB instance, and call callback function
	tuplet_dotStaged := backRepo.BackRepoTuplet_dot.Map_Tuplet_dotDBID_Tuplet_dotPtr[tuplet_dotDB.ID]
	if tuplet_dotStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tuplet_dotStaged, tuplet_dotDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
