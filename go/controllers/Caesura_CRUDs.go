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
var __Caesura__dummysDeclaration__ models.Caesura
var __Caesura_time__dummyDeclaration time.Duration

var mutexCaesura sync.Mutex

// An CaesuraID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCaesura updateCaesura deleteCaesura
type CaesuraID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CaesuraInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCaesura updateCaesura
type CaesuraInput struct {
	// The Caesura to submit or modify
	// in: body
	Caesura *orm.CaesuraAPI
}

// GetCaesuras
//
// swagger:route GET /caesuras caesuras getCaesuras
//
// # Get all caesuras
//
// Responses:
// default: genericError
//
//	200: caesuraDBResponse
func (controller *Controller) GetCaesuras(c *gin.Context) {

	// source slice
	var caesuraDBs []orm.CaesuraDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCaesuras", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCaesura.GetDB()

	query := db.Find(&caesuraDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	caesuraAPIs := make([]orm.CaesuraAPI, 0)

	// for each caesura, update fields from the database nullable fields
	for idx := range caesuraDBs {
		caesuraDB := &caesuraDBs[idx]
		_ = caesuraDB
		var caesuraAPI orm.CaesuraAPI

		// insertion point for updating fields
		caesuraAPI.ID = caesuraDB.ID
		caesuraDB.CopyBasicFieldsToCaesura_WOP(&caesuraAPI.Caesura_WOP)
		caesuraAPI.CaesuraPointersEncoding = caesuraDB.CaesuraPointersEncoding
		caesuraAPIs = append(caesuraAPIs, caesuraAPI)
	}

	c.JSON(http.StatusOK, caesuraAPIs)
}

// PostCaesura
//
// swagger:route POST /caesuras caesuras postCaesura
//
// Creates a caesura
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCaesura(c *gin.Context) {

	mutexCaesura.Lock()
	defer mutexCaesura.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCaesuras", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCaesura.GetDB()

	// Validate input
	var input orm.CaesuraAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create caesura
	caesuraDB := orm.CaesuraDB{}
	caesuraDB.CaesuraPointersEncoding = input.CaesuraPointersEncoding
	caesuraDB.CopyBasicFieldsFromCaesura_WOP(&input.Caesura_WOP)

	query := db.Create(&caesuraDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCaesura.CheckoutPhaseOneInstance(&caesuraDB)
	caesura := backRepo.BackRepoCaesura.Map_CaesuraDBID_CaesuraPtr[caesuraDB.ID]

	if caesura != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), caesura)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, caesuraDB)
}

// GetCaesura
//
// swagger:route GET /caesuras/{ID} caesuras getCaesura
//
// Gets the details for a caesura.
//
// Responses:
// default: genericError
//
//	200: caesuraDBResponse
func (controller *Controller) GetCaesura(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCaesura", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCaesura.GetDB()

	// Get caesuraDB in DB
	var caesuraDB orm.CaesuraDB
	if err := db.First(&caesuraDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var caesuraAPI orm.CaesuraAPI
	caesuraAPI.ID = caesuraDB.ID
	caesuraAPI.CaesuraPointersEncoding = caesuraDB.CaesuraPointersEncoding
	caesuraDB.CopyBasicFieldsToCaesura_WOP(&caesuraAPI.Caesura_WOP)

	c.JSON(http.StatusOK, caesuraAPI)
}

// UpdateCaesura
//
// swagger:route PATCH /caesuras/{ID} caesuras updateCaesura
//
// # Update a caesura
//
// Responses:
// default: genericError
//
//	200: caesuraDBResponse
func (controller *Controller) UpdateCaesura(c *gin.Context) {

	mutexCaesura.Lock()
	defer mutexCaesura.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCaesura", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCaesura.GetDB()

	// Validate input
	var input orm.CaesuraAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var caesuraDB orm.CaesuraDB

	// fetch the caesura
	query := db.First(&caesuraDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	caesuraDB.CopyBasicFieldsFromCaesura_WOP(&input.Caesura_WOP)
	caesuraDB.CaesuraPointersEncoding = input.CaesuraPointersEncoding

	query = db.Model(&caesuraDB).Updates(caesuraDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	caesuraNew := new(models.Caesura)
	caesuraDB.CopyBasicFieldsToCaesura(caesuraNew)

	// redeem pointers
	caesuraDB.DecodePointers(backRepo, caesuraNew)

	// get stage instance from DB instance, and call callback function
	caesuraOld := backRepo.BackRepoCaesura.Map_CaesuraDBID_CaesuraPtr[caesuraDB.ID]
	if caesuraOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), caesuraOld, caesuraNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the caesuraDB
	c.JSON(http.StatusOK, caesuraDB)
}

// DeleteCaesura
//
// swagger:route DELETE /caesuras/{ID} caesuras deleteCaesura
//
// # Delete a caesura
//
// default: genericError
//
//	200: caesuraDBResponse
func (controller *Controller) DeleteCaesura(c *gin.Context) {

	mutexCaesura.Lock()
	defer mutexCaesura.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCaesura", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCaesura.GetDB()

	// Get model if exist
	var caesuraDB orm.CaesuraDB
	if err := db.First(&caesuraDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&caesuraDB)

	// get an instance (not staged) from DB instance, and call callback function
	caesuraDeleted := new(models.Caesura)
	caesuraDB.CopyBasicFieldsToCaesura(caesuraDeleted)

	// get stage instance from DB instance, and call callback function
	caesuraStaged := backRepo.BackRepoCaesura.Map_CaesuraDBID_CaesuraPtr[caesuraDB.ID]
	if caesuraStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), caesuraStaged, caesuraDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
