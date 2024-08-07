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
var __Barline__dummysDeclaration__ models.Barline
var __Barline_time__dummyDeclaration time.Duration

var mutexBarline sync.Mutex

// An BarlineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBarline updateBarline deleteBarline
type BarlineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BarlineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBarline updateBarline
type BarlineInput struct {
	// The Barline to submit or modify
	// in: body
	Barline *orm.BarlineAPI
}

// GetBarlines
//
// swagger:route GET /barlines barlines getBarlines
//
// # Get all barlines
//
// Responses:
// default: genericError
//
//	200: barlineDBResponse
func (controller *Controller) GetBarlines(c *gin.Context) {

	// source slice
	var barlineDBs []orm.BarlineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBarlines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarline.GetDB()

	query := db.Find(&barlineDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	barlineAPIs := make([]orm.BarlineAPI, 0)

	// for each barline, update fields from the database nullable fields
	for idx := range barlineDBs {
		barlineDB := &barlineDBs[idx]
		_ = barlineDB
		var barlineAPI orm.BarlineAPI

		// insertion point for updating fields
		barlineAPI.ID = barlineDB.ID
		barlineDB.CopyBasicFieldsToBarline_WOP(&barlineAPI.Barline_WOP)
		barlineAPI.BarlinePointersEncoding = barlineDB.BarlinePointersEncoding
		barlineAPIs = append(barlineAPIs, barlineAPI)
	}

	c.JSON(http.StatusOK, barlineAPIs)
}

// PostBarline
//
// swagger:route POST /barlines barlines postBarline
//
// Creates a barline
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBarline(c *gin.Context) {

	mutexBarline.Lock()
	defer mutexBarline.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBarlines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarline.GetDB()

	// Validate input
	var input orm.BarlineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create barline
	barlineDB := orm.BarlineDB{}
	barlineDB.BarlinePointersEncoding = input.BarlinePointersEncoding
	barlineDB.CopyBasicFieldsFromBarline_WOP(&input.Barline_WOP)

	query := db.Create(&barlineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBarline.CheckoutPhaseOneInstance(&barlineDB)
	barline := backRepo.BackRepoBarline.Map_BarlineDBID_BarlinePtr[barlineDB.ID]

	if barline != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), barline)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, barlineDB)
}

// GetBarline
//
// swagger:route GET /barlines/{ID} barlines getBarline
//
// Gets the details for a barline.
//
// Responses:
// default: genericError
//
//	200: barlineDBResponse
func (controller *Controller) GetBarline(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBarline", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarline.GetDB()

	// Get barlineDB in DB
	var barlineDB orm.BarlineDB
	if err := db.First(&barlineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var barlineAPI orm.BarlineAPI
	barlineAPI.ID = barlineDB.ID
	barlineAPI.BarlinePointersEncoding = barlineDB.BarlinePointersEncoding
	barlineDB.CopyBasicFieldsToBarline_WOP(&barlineAPI.Barline_WOP)

	c.JSON(http.StatusOK, barlineAPI)
}

// UpdateBarline
//
// swagger:route PATCH /barlines/{ID} barlines updateBarline
//
// # Update a barline
//
// Responses:
// default: genericError
//
//	200: barlineDBResponse
func (controller *Controller) UpdateBarline(c *gin.Context) {

	mutexBarline.Lock()
	defer mutexBarline.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBarline", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarline.GetDB()

	// Validate input
	var input orm.BarlineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var barlineDB orm.BarlineDB

	// fetch the barline
	query := db.First(&barlineDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	barlineDB.CopyBasicFieldsFromBarline_WOP(&input.Barline_WOP)
	barlineDB.BarlinePointersEncoding = input.BarlinePointersEncoding

	query = db.Model(&barlineDB).Updates(barlineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	barlineNew := new(models.Barline)
	barlineDB.CopyBasicFieldsToBarline(barlineNew)

	// redeem pointers
	barlineDB.DecodePointers(backRepo, barlineNew)

	// get stage instance from DB instance, and call callback function
	barlineOld := backRepo.BackRepoBarline.Map_BarlineDBID_BarlinePtr[barlineDB.ID]
	if barlineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), barlineOld, barlineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the barlineDB
	c.JSON(http.StatusOK, barlineDB)
}

// DeleteBarline
//
// swagger:route DELETE /barlines/{ID} barlines deleteBarline
//
// # Delete a barline
//
// default: genericError
//
//	200: barlineDBResponse
func (controller *Controller) DeleteBarline(c *gin.Context) {

	mutexBarline.Lock()
	defer mutexBarline.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBarline", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBarline.GetDB()

	// Get model if exist
	var barlineDB orm.BarlineDB
	if err := db.First(&barlineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&barlineDB)

	// get an instance (not staged) from DB instance, and call callback function
	barlineDeleted := new(models.Barline)
	barlineDB.CopyBasicFieldsToBarline(barlineDeleted)

	// get stage instance from DB instance, and call callback function
	barlineStaged := backRepo.BackRepoBarline.Map_BarlineDBID_BarlinePtr[barlineDB.ID]
	if barlineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), barlineStaged, barlineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
