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
var __Arrow__dummysDeclaration__ models.Arrow
var __Arrow_time__dummyDeclaration time.Duration

var mutexArrow sync.Mutex

// An ArrowID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getArrow updateArrow deleteArrow
type ArrowID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ArrowInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postArrow updateArrow
type ArrowInput struct {
	// The Arrow to submit or modify
	// in: body
	Arrow *orm.ArrowAPI
}

// GetArrows
//
// swagger:route GET /arrows arrows getArrows
//
// # Get all arrows
//
// Responses:
// default: genericError
//
//	200: arrowDBResponse
func (controller *Controller) GetArrows(c *gin.Context) {

	// source slice
	var arrowDBs []orm.ArrowDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetArrows", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArrow.GetDB()

	query := db.Find(&arrowDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	arrowAPIs := make([]orm.ArrowAPI, 0)

	// for each arrow, update fields from the database nullable fields
	for idx := range arrowDBs {
		arrowDB := &arrowDBs[idx]
		_ = arrowDB
		var arrowAPI orm.ArrowAPI

		// insertion point for updating fields
		arrowAPI.ID = arrowDB.ID
		arrowDB.CopyBasicFieldsToArrow_WOP(&arrowAPI.Arrow_WOP)
		arrowAPI.ArrowPointersEncoding = arrowDB.ArrowPointersEncoding
		arrowAPIs = append(arrowAPIs, arrowAPI)
	}

	c.JSON(http.StatusOK, arrowAPIs)
}

// PostArrow
//
// swagger:route POST /arrows arrows postArrow
//
// Creates a arrow
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostArrow(c *gin.Context) {

	mutexArrow.Lock()
	defer mutexArrow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostArrows", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArrow.GetDB()

	// Validate input
	var input orm.ArrowAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create arrow
	arrowDB := orm.ArrowDB{}
	arrowDB.ArrowPointersEncoding = input.ArrowPointersEncoding
	arrowDB.CopyBasicFieldsFromArrow_WOP(&input.Arrow_WOP)

	query := db.Create(&arrowDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoArrow.CheckoutPhaseOneInstance(&arrowDB)
	arrow := backRepo.BackRepoArrow.Map_ArrowDBID_ArrowPtr[arrowDB.ID]

	if arrow != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), arrow)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, arrowDB)
}

// GetArrow
//
// swagger:route GET /arrows/{ID} arrows getArrow
//
// Gets the details for a arrow.
//
// Responses:
// default: genericError
//
//	200: arrowDBResponse
func (controller *Controller) GetArrow(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetArrow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArrow.GetDB()

	// Get arrowDB in DB
	var arrowDB orm.ArrowDB
	if err := db.First(&arrowDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var arrowAPI orm.ArrowAPI
	arrowAPI.ID = arrowDB.ID
	arrowAPI.ArrowPointersEncoding = arrowDB.ArrowPointersEncoding
	arrowDB.CopyBasicFieldsToArrow_WOP(&arrowAPI.Arrow_WOP)

	c.JSON(http.StatusOK, arrowAPI)
}

// UpdateArrow
//
// swagger:route PATCH /arrows/{ID} arrows updateArrow
//
// # Update a arrow
//
// Responses:
// default: genericError
//
//	200: arrowDBResponse
func (controller *Controller) UpdateArrow(c *gin.Context) {

	mutexArrow.Lock()
	defer mutexArrow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateArrow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArrow.GetDB()

	// Validate input
	var input orm.ArrowAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var arrowDB orm.ArrowDB

	// fetch the arrow
	query := db.First(&arrowDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	arrowDB.CopyBasicFieldsFromArrow_WOP(&input.Arrow_WOP)
	arrowDB.ArrowPointersEncoding = input.ArrowPointersEncoding

	query = db.Model(&arrowDB).Updates(arrowDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	arrowNew := new(models.Arrow)
	arrowDB.CopyBasicFieldsToArrow(arrowNew)

	// redeem pointers
	arrowDB.DecodePointers(backRepo, arrowNew)

	// get stage instance from DB instance, and call callback function
	arrowOld := backRepo.BackRepoArrow.Map_ArrowDBID_ArrowPtr[arrowDB.ID]
	if arrowOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), arrowOld, arrowNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the arrowDB
	c.JSON(http.StatusOK, arrowDB)
}

// DeleteArrow
//
// swagger:route DELETE /arrows/{ID} arrows deleteArrow
//
// # Delete a arrow
//
// default: genericError
//
//	200: arrowDBResponse
func (controller *Controller) DeleteArrow(c *gin.Context) {

	mutexArrow.Lock()
	defer mutexArrow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteArrow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArrow.GetDB()

	// Get model if exist
	var arrowDB orm.ArrowDB
	if err := db.First(&arrowDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&arrowDB)

	// get an instance (not staged) from DB instance, and call callback function
	arrowDeleted := new(models.Arrow)
	arrowDB.CopyBasicFieldsToArrow(arrowDeleted)

	// get stage instance from DB instance, and call callback function
	arrowStaged := backRepo.BackRepoArrow.Map_ArrowDBID_ArrowPtr[arrowDB.ID]
	if arrowStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), arrowStaged, arrowDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
