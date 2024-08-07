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
var __Empty_font__dummysDeclaration__ models.Empty_font
var __Empty_font_time__dummyDeclaration time.Duration

var mutexEmpty_font sync.Mutex

// An Empty_fontID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmpty_font updateEmpty_font deleteEmpty_font
type Empty_fontID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Empty_fontInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmpty_font updateEmpty_font
type Empty_fontInput struct {
	// The Empty_font to submit or modify
	// in: body
	Empty_font *orm.Empty_fontAPI
}

// GetEmpty_fonts
//
// swagger:route GET /empty_fonts empty_fonts getEmpty_fonts
//
// # Get all empty_fonts
//
// Responses:
// default: genericError
//
//	200: empty_fontDBResponse
func (controller *Controller) GetEmpty_fonts(c *gin.Context) {

	// source slice
	var empty_fontDBs []orm.Empty_fontDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_fonts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_font.GetDB()

	query := db.Find(&empty_fontDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	empty_fontAPIs := make([]orm.Empty_fontAPI, 0)

	// for each empty_font, update fields from the database nullable fields
	for idx := range empty_fontDBs {
		empty_fontDB := &empty_fontDBs[idx]
		_ = empty_fontDB
		var empty_fontAPI orm.Empty_fontAPI

		// insertion point for updating fields
		empty_fontAPI.ID = empty_fontDB.ID
		empty_fontDB.CopyBasicFieldsToEmpty_font_WOP(&empty_fontAPI.Empty_font_WOP)
		empty_fontAPI.Empty_fontPointersEncoding = empty_fontDB.Empty_fontPointersEncoding
		empty_fontAPIs = append(empty_fontAPIs, empty_fontAPI)
	}

	c.JSON(http.StatusOK, empty_fontAPIs)
}

// PostEmpty_font
//
// swagger:route POST /empty_fonts empty_fonts postEmpty_font
//
// Creates a empty_font
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmpty_font(c *gin.Context) {

	mutexEmpty_font.Lock()
	defer mutexEmpty_font.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmpty_fonts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_font.GetDB()

	// Validate input
	var input orm.Empty_fontAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create empty_font
	empty_fontDB := orm.Empty_fontDB{}
	empty_fontDB.Empty_fontPointersEncoding = input.Empty_fontPointersEncoding
	empty_fontDB.CopyBasicFieldsFromEmpty_font_WOP(&input.Empty_font_WOP)

	query := db.Create(&empty_fontDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmpty_font.CheckoutPhaseOneInstance(&empty_fontDB)
	empty_font := backRepo.BackRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr[empty_fontDB.ID]

	if empty_font != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), empty_font)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, empty_fontDB)
}

// GetEmpty_font
//
// swagger:route GET /empty_fonts/{ID} empty_fonts getEmpty_font
//
// Gets the details for a empty_font.
//
// Responses:
// default: genericError
//
//	200: empty_fontDBResponse
func (controller *Controller) GetEmpty_font(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_font", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_font.GetDB()

	// Get empty_fontDB in DB
	var empty_fontDB orm.Empty_fontDB
	if err := db.First(&empty_fontDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var empty_fontAPI orm.Empty_fontAPI
	empty_fontAPI.ID = empty_fontDB.ID
	empty_fontAPI.Empty_fontPointersEncoding = empty_fontDB.Empty_fontPointersEncoding
	empty_fontDB.CopyBasicFieldsToEmpty_font_WOP(&empty_fontAPI.Empty_font_WOP)

	c.JSON(http.StatusOK, empty_fontAPI)
}

// UpdateEmpty_font
//
// swagger:route PATCH /empty_fonts/{ID} empty_fonts updateEmpty_font
//
// # Update a empty_font
//
// Responses:
// default: genericError
//
//	200: empty_fontDBResponse
func (controller *Controller) UpdateEmpty_font(c *gin.Context) {

	mutexEmpty_font.Lock()
	defer mutexEmpty_font.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEmpty_font", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_font.GetDB()

	// Validate input
	var input orm.Empty_fontAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var empty_fontDB orm.Empty_fontDB

	// fetch the empty_font
	query := db.First(&empty_fontDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	empty_fontDB.CopyBasicFieldsFromEmpty_font_WOP(&input.Empty_font_WOP)
	empty_fontDB.Empty_fontPointersEncoding = input.Empty_fontPointersEncoding

	query = db.Model(&empty_fontDB).Updates(empty_fontDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	empty_fontNew := new(models.Empty_font)
	empty_fontDB.CopyBasicFieldsToEmpty_font(empty_fontNew)

	// redeem pointers
	empty_fontDB.DecodePointers(backRepo, empty_fontNew)

	// get stage instance from DB instance, and call callback function
	empty_fontOld := backRepo.BackRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr[empty_fontDB.ID]
	if empty_fontOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), empty_fontOld, empty_fontNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the empty_fontDB
	c.JSON(http.StatusOK, empty_fontDB)
}

// DeleteEmpty_font
//
// swagger:route DELETE /empty_fonts/{ID} empty_fonts deleteEmpty_font
//
// # Delete a empty_font
//
// default: genericError
//
//	200: empty_fontDBResponse
func (controller *Controller) DeleteEmpty_font(c *gin.Context) {

	mutexEmpty_font.Lock()
	defer mutexEmpty_font.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmpty_font", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_font.GetDB()

	// Get model if exist
	var empty_fontDB orm.Empty_fontDB
	if err := db.First(&empty_fontDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&empty_fontDB)

	// get an instance (not staged) from DB instance, and call callback function
	empty_fontDeleted := new(models.Empty_font)
	empty_fontDB.CopyBasicFieldsToEmpty_font(empty_fontDeleted)

	// get stage instance from DB instance, and call callback function
	empty_fontStaged := backRepo.BackRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr[empty_fontDB.ID]
	if empty_fontStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), empty_fontStaged, empty_fontDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
