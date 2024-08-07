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
var __Barre__dummysDeclaration__ models.Barre
var __Barre_time__dummyDeclaration time.Duration

var mutexBarre sync.Mutex

// An BarreID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBarre updateBarre deleteBarre
type BarreID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BarreInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBarre updateBarre
type BarreInput struct {
	// The Barre to submit or modify
	// in: body
	Barre *orm.BarreAPI
}

// GetBarres
//
// swagger:route GET /barres barres getBarres
//
// # Get all barres
//
// Responses:
// default: genericError
//
//	200: barreDBResponse
func (controller *Controller) GetBarres(c *gin.Context) {

	// source slice
	var barreDBs []orm.BarreDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBarres", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarre.GetDB()

	query := db.Find(&barreDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	barreAPIs := make([]orm.BarreAPI, 0)

	// for each barre, update fields from the database nullable fields
	for idx := range barreDBs {
		barreDB := &barreDBs[idx]
		_ = barreDB
		var barreAPI orm.BarreAPI

		// insertion point for updating fields
		barreAPI.ID = barreDB.ID
		barreDB.CopyBasicFieldsToBarre_WOP(&barreAPI.Barre_WOP)
		barreAPI.BarrePointersEncoding = barreDB.BarrePointersEncoding
		barreAPIs = append(barreAPIs, barreAPI)
	}

	c.JSON(http.StatusOK, barreAPIs)
}

// PostBarre
//
// swagger:route POST /barres barres postBarre
//
// Creates a barre
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBarre(c *gin.Context) {

	mutexBarre.Lock()
	defer mutexBarre.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBarres", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarre.GetDB()

	// Validate input
	var input orm.BarreAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create barre
	barreDB := orm.BarreDB{}
	barreDB.BarrePointersEncoding = input.BarrePointersEncoding
	barreDB.CopyBasicFieldsFromBarre_WOP(&input.Barre_WOP)

	query := db.Create(&barreDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBarre.CheckoutPhaseOneInstance(&barreDB)
	barre := backRepo.BackRepoBarre.Map_BarreDBID_BarrePtr[barreDB.ID]

	if barre != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), barre)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, barreDB)
}

// GetBarre
//
// swagger:route GET /barres/{ID} barres getBarre
//
// Gets the details for a barre.
//
// Responses:
// default: genericError
//
//	200: barreDBResponse
func (controller *Controller) GetBarre(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBarre", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarre.GetDB()

	// Get barreDB in DB
	var barreDB orm.BarreDB
	if err := db.First(&barreDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var barreAPI orm.BarreAPI
	barreAPI.ID = barreDB.ID
	barreAPI.BarrePointersEncoding = barreDB.BarrePointersEncoding
	barreDB.CopyBasicFieldsToBarre_WOP(&barreAPI.Barre_WOP)

	c.JSON(http.StatusOK, barreAPI)
}

// UpdateBarre
//
// swagger:route PATCH /barres/{ID} barres updateBarre
//
// # Update a barre
//
// Responses:
// default: genericError
//
//	200: barreDBResponse
func (controller *Controller) UpdateBarre(c *gin.Context) {

	mutexBarre.Lock()
	defer mutexBarre.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBarre", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarre.GetDB()

	// Validate input
	var input orm.BarreAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var barreDB orm.BarreDB

	// fetch the barre
	query := db.First(&barreDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	barreDB.CopyBasicFieldsFromBarre_WOP(&input.Barre_WOP)
	barreDB.BarrePointersEncoding = input.BarrePointersEncoding

	query = db.Model(&barreDB).Updates(barreDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	barreNew := new(models.Barre)
	barreDB.CopyBasicFieldsToBarre(barreNew)

	// redeem pointers
	barreDB.DecodePointers(backRepo, barreNew)

	// get stage instance from DB instance, and call callback function
	barreOld := backRepo.BackRepoBarre.Map_BarreDBID_BarrePtr[barreDB.ID]
	if barreOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), barreOld, barreNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the barreDB
	c.JSON(http.StatusOK, barreDB)
}

// DeleteBarre
//
// swagger:route DELETE /barres/{ID} barres deleteBarre
//
// # Delete a barre
//
// default: genericError
//
//	200: barreDBResponse
func (controller *Controller) DeleteBarre(c *gin.Context) {

	mutexBarre.Lock()
	defer mutexBarre.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBarre", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarre.GetDB()

	// Get model if exist
	var barreDB orm.BarreDB
	if err := db.First(&barreDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&barreDB)

	// get an instance (not staged) from DB instance, and call callback function
	barreDeleted := new(models.Barre)
	barreDB.CopyBasicFieldsToBarre(barreDeleted)

	// get stage instance from DB instance, and call callback function
	barreStaged := backRepo.BackRepoBarre.Map_BarreDBID_BarrePtr[barreDB.ID]
	if barreStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), barreStaged, barreDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
