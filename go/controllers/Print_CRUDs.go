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
var __Print__dummysDeclaration__ models.Print
var __Print_time__dummyDeclaration time.Duration

var mutexPrint sync.Mutex

// An PrintID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPrint updatePrint deletePrint
type PrintID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PrintInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPrint updatePrint
type PrintInput struct {
	// The Print to submit or modify
	// in: body
	Print *orm.PrintAPI
}

// GetPrints
//
// swagger:route GET /prints prints getPrints
//
// # Get all prints
//
// Responses:
// default: genericError
//
//	200: printDBResponse
func (controller *Controller) GetPrints(c *gin.Context) {

	// source slice
	var printDBs []orm.PrintDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPrints", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrint.GetDB()

	query := db.Find(&printDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	printAPIs := make([]orm.PrintAPI, 0)

	// for each print, update fields from the database nullable fields
	for idx := range printDBs {
		printDB := &printDBs[idx]
		_ = printDB
		var printAPI orm.PrintAPI

		// insertion point for updating fields
		printAPI.ID = printDB.ID
		printDB.CopyBasicFieldsToPrint_WOP(&printAPI.Print_WOP)
		printAPI.PrintPointersEncoding = printDB.PrintPointersEncoding
		printAPIs = append(printAPIs, printAPI)
	}

	c.JSON(http.StatusOK, printAPIs)
}

// PostPrint
//
// swagger:route POST /prints prints postPrint
//
// Creates a print
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPrint(c *gin.Context) {

	mutexPrint.Lock()
	defer mutexPrint.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPrints", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrint.GetDB()

	// Validate input
	var input orm.PrintAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create print
	printDB := orm.PrintDB{}
	printDB.PrintPointersEncoding = input.PrintPointersEncoding
	printDB.CopyBasicFieldsFromPrint_WOP(&input.Print_WOP)

	query := db.Create(&printDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPrint.CheckoutPhaseOneInstance(&printDB)
	print := backRepo.BackRepoPrint.Map_PrintDBID_PrintPtr[printDB.ID]

	if print != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), print)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, printDB)
}

// GetPrint
//
// swagger:route GET /prints/{ID} prints getPrint
//
// Gets the details for a print.
//
// Responses:
// default: genericError
//
//	200: printDBResponse
func (controller *Controller) GetPrint(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPrint", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrint.GetDB()

	// Get printDB in DB
	var printDB orm.PrintDB
	if err := db.First(&printDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var printAPI orm.PrintAPI
	printAPI.ID = printDB.ID
	printAPI.PrintPointersEncoding = printDB.PrintPointersEncoding
	printDB.CopyBasicFieldsToPrint_WOP(&printAPI.Print_WOP)

	c.JSON(http.StatusOK, printAPI)
}

// UpdatePrint
//
// swagger:route PATCH /prints/{ID} prints updatePrint
//
// # Update a print
//
// Responses:
// default: genericError
//
//	200: printDBResponse
func (controller *Controller) UpdatePrint(c *gin.Context) {

	mutexPrint.Lock()
	defer mutexPrint.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePrint", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrint.GetDB()

	// Validate input
	var input orm.PrintAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var printDB orm.PrintDB

	// fetch the print
	query := db.First(&printDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	printDB.CopyBasicFieldsFromPrint_WOP(&input.Print_WOP)
	printDB.PrintPointersEncoding = input.PrintPointersEncoding

	query = db.Model(&printDB).Updates(printDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	printNew := new(models.Print)
	printDB.CopyBasicFieldsToPrint(printNew)

	// redeem pointers
	printDB.DecodePointers(backRepo, printNew)

	// get stage instance from DB instance, and call callback function
	printOld := backRepo.BackRepoPrint.Map_PrintDBID_PrintPtr[printDB.ID]
	if printOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), printOld, printNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the printDB
	c.JSON(http.StatusOK, printDB)
}

// DeletePrint
//
// swagger:route DELETE /prints/{ID} prints deletePrint
//
// # Delete a print
//
// default: genericError
//
//	200: printDBResponse
func (controller *Controller) DeletePrint(c *gin.Context) {

	mutexPrint.Lock()
	defer mutexPrint.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePrint", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrint.GetDB()

	// Get model if exist
	var printDB orm.PrintDB
	if err := db.First(&printDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&printDB)

	// get an instance (not staged) from DB instance, and call callback function
	printDeleted := new(models.Print)
	printDB.CopyBasicFieldsToPrint(printDeleted)

	// get stage instance from DB instance, and call callback function
	printStaged := backRepo.BackRepoPrint.Map_PrintDBID_PrintPtr[printDB.ID]
	if printStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), printStaged, printDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
