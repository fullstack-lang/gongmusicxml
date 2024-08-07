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
var __Bracket__dummysDeclaration__ models.Bracket
var __Bracket_time__dummyDeclaration time.Duration

var mutexBracket sync.Mutex

// An BracketID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBracket updateBracket deleteBracket
type BracketID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BracketInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBracket updateBracket
type BracketInput struct {
	// The Bracket to submit or modify
	// in: body
	Bracket *orm.BracketAPI
}

// GetBrackets
//
// swagger:route GET /brackets brackets getBrackets
//
// # Get all brackets
//
// Responses:
// default: genericError
//
//	200: bracketDBResponse
func (controller *Controller) GetBrackets(c *gin.Context) {

	// source slice
	var bracketDBs []orm.BracketDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBrackets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBracket.GetDB()

	query := db.Find(&bracketDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bracketAPIs := make([]orm.BracketAPI, 0)

	// for each bracket, update fields from the database nullable fields
	for idx := range bracketDBs {
		bracketDB := &bracketDBs[idx]
		_ = bracketDB
		var bracketAPI orm.BracketAPI

		// insertion point for updating fields
		bracketAPI.ID = bracketDB.ID
		bracketDB.CopyBasicFieldsToBracket_WOP(&bracketAPI.Bracket_WOP)
		bracketAPI.BracketPointersEncoding = bracketDB.BracketPointersEncoding
		bracketAPIs = append(bracketAPIs, bracketAPI)
	}

	c.JSON(http.StatusOK, bracketAPIs)
}

// PostBracket
//
// swagger:route POST /brackets brackets postBracket
//
// Creates a bracket
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBracket(c *gin.Context) {

	mutexBracket.Lock()
	defer mutexBracket.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBrackets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBracket.GetDB()

	// Validate input
	var input orm.BracketAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bracket
	bracketDB := orm.BracketDB{}
	bracketDB.BracketPointersEncoding = input.BracketPointersEncoding
	bracketDB.CopyBasicFieldsFromBracket_WOP(&input.Bracket_WOP)

	query := db.Create(&bracketDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBracket.CheckoutPhaseOneInstance(&bracketDB)
	bracket := backRepo.BackRepoBracket.Map_BracketDBID_BracketPtr[bracketDB.ID]

	if bracket != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bracket)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bracketDB)
}

// GetBracket
//
// swagger:route GET /brackets/{ID} brackets getBracket
//
// Gets the details for a bracket.
//
// Responses:
// default: genericError
//
//	200: bracketDBResponse
func (controller *Controller) GetBracket(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBracket", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBracket.GetDB()

	// Get bracketDB in DB
	var bracketDB orm.BracketDB
	if err := db.First(&bracketDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bracketAPI orm.BracketAPI
	bracketAPI.ID = bracketDB.ID
	bracketAPI.BracketPointersEncoding = bracketDB.BracketPointersEncoding
	bracketDB.CopyBasicFieldsToBracket_WOP(&bracketAPI.Bracket_WOP)

	c.JSON(http.StatusOK, bracketAPI)
}

// UpdateBracket
//
// swagger:route PATCH /brackets/{ID} brackets updateBracket
//
// # Update a bracket
//
// Responses:
// default: genericError
//
//	200: bracketDBResponse
func (controller *Controller) UpdateBracket(c *gin.Context) {

	mutexBracket.Lock()
	defer mutexBracket.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBracket", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBracket.GetDB()

	// Validate input
	var input orm.BracketAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bracketDB orm.BracketDB

	// fetch the bracket
	query := db.First(&bracketDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bracketDB.CopyBasicFieldsFromBracket_WOP(&input.Bracket_WOP)
	bracketDB.BracketPointersEncoding = input.BracketPointersEncoding

	query = db.Model(&bracketDB).Updates(bracketDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bracketNew := new(models.Bracket)
	bracketDB.CopyBasicFieldsToBracket(bracketNew)

	// redeem pointers
	bracketDB.DecodePointers(backRepo, bracketNew)

	// get stage instance from DB instance, and call callback function
	bracketOld := backRepo.BackRepoBracket.Map_BracketDBID_BracketPtr[bracketDB.ID]
	if bracketOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bracketOld, bracketNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bracketDB
	c.JSON(http.StatusOK, bracketDB)
}

// DeleteBracket
//
// swagger:route DELETE /brackets/{ID} brackets deleteBracket
//
// # Delete a bracket
//
// default: genericError
//
//	200: bracketDBResponse
func (controller *Controller) DeleteBracket(c *gin.Context) {

	mutexBracket.Lock()
	defer mutexBracket.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBracket", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBracket.GetDB()

	// Get model if exist
	var bracketDB orm.BracketDB
	if err := db.First(&bracketDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bracketDB)

	// get an instance (not staged) from DB instance, and call callback function
	bracketDeleted := new(models.Bracket)
	bracketDB.CopyBasicFieldsToBracket(bracketDeleted)

	// get stage instance from DB instance, and call callback function
	bracketStaged := backRepo.BackRepoBracket.Map_BracketDBID_BracketPtr[bracketDB.ID]
	if bracketStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bracketStaged, bracketDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
