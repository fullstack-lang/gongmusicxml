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
var __Figure__dummysDeclaration__ models.Figure
var __Figure_time__dummyDeclaration time.Duration

var mutexFigure sync.Mutex

// An FigureID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFigure updateFigure deleteFigure
type FigureID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FigureInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFigure updateFigure
type FigureInput struct {
	// The Figure to submit or modify
	// in: body
	Figure *orm.FigureAPI
}

// GetFigures
//
// swagger:route GET /figures figures getFigures
//
// # Get all figures
//
// Responses:
// default: genericError
//
//	200: figureDBResponse
func (controller *Controller) GetFigures(c *gin.Context) {

	// source slice
	var figureDBs []orm.FigureDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFigures", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigure.GetDB()

	query := db.Find(&figureDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	figureAPIs := make([]orm.FigureAPI, 0)

	// for each figure, update fields from the database nullable fields
	for idx := range figureDBs {
		figureDB := &figureDBs[idx]
		_ = figureDB
		var figureAPI orm.FigureAPI

		// insertion point for updating fields
		figureAPI.ID = figureDB.ID
		figureDB.CopyBasicFieldsToFigure_WOP(&figureAPI.Figure_WOP)
		figureAPI.FigurePointersEncoding = figureDB.FigurePointersEncoding
		figureAPIs = append(figureAPIs, figureAPI)
	}

	c.JSON(http.StatusOK, figureAPIs)
}

// PostFigure
//
// swagger:route POST /figures figures postFigure
//
// Creates a figure
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFigure(c *gin.Context) {

	mutexFigure.Lock()
	defer mutexFigure.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFigures", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigure.GetDB()

	// Validate input
	var input orm.FigureAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create figure
	figureDB := orm.FigureDB{}
	figureDB.FigurePointersEncoding = input.FigurePointersEncoding
	figureDB.CopyBasicFieldsFromFigure_WOP(&input.Figure_WOP)

	query := db.Create(&figureDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFigure.CheckoutPhaseOneInstance(&figureDB)
	figure := backRepo.BackRepoFigure.Map_FigureDBID_FigurePtr[figureDB.ID]

	if figure != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), figure)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, figureDB)
}

// GetFigure
//
// swagger:route GET /figures/{ID} figures getFigure
//
// Gets the details for a figure.
//
// Responses:
// default: genericError
//
//	200: figureDBResponse
func (controller *Controller) GetFigure(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFigure", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigure.GetDB()

	// Get figureDB in DB
	var figureDB orm.FigureDB
	if err := db.First(&figureDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var figureAPI orm.FigureAPI
	figureAPI.ID = figureDB.ID
	figureAPI.FigurePointersEncoding = figureDB.FigurePointersEncoding
	figureDB.CopyBasicFieldsToFigure_WOP(&figureAPI.Figure_WOP)

	c.JSON(http.StatusOK, figureAPI)
}

// UpdateFigure
//
// swagger:route PATCH /figures/{ID} figures updateFigure
//
// # Update a figure
//
// Responses:
// default: genericError
//
//	200: figureDBResponse
func (controller *Controller) UpdateFigure(c *gin.Context) {

	mutexFigure.Lock()
	defer mutexFigure.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFigure", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigure.GetDB()

	// Validate input
	var input orm.FigureAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var figureDB orm.FigureDB

	// fetch the figure
	query := db.First(&figureDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	figureDB.CopyBasicFieldsFromFigure_WOP(&input.Figure_WOP)
	figureDB.FigurePointersEncoding = input.FigurePointersEncoding

	query = db.Model(&figureDB).Updates(figureDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	figureNew := new(models.Figure)
	figureDB.CopyBasicFieldsToFigure(figureNew)

	// redeem pointers
	figureDB.DecodePointers(backRepo, figureNew)

	// get stage instance from DB instance, and call callback function
	figureOld := backRepo.BackRepoFigure.Map_FigureDBID_FigurePtr[figureDB.ID]
	if figureOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), figureOld, figureNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the figureDB
	c.JSON(http.StatusOK, figureDB)
}

// DeleteFigure
//
// swagger:route DELETE /figures/{ID} figures deleteFigure
//
// # Delete a figure
//
// default: genericError
//
//	200: figureDBResponse
func (controller *Controller) DeleteFigure(c *gin.Context) {

	mutexFigure.Lock()
	defer mutexFigure.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFigure", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigure.GetDB()

	// Get model if exist
	var figureDB orm.FigureDB
	if err := db.First(&figureDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&figureDB)

	// get an instance (not staged) from DB instance, and call callback function
	figureDeleted := new(models.Figure)
	figureDB.CopyBasicFieldsToFigure(figureDeleted)

	// get stage instance from DB instance, and call callback function
	figureStaged := backRepo.BackRepoFigure.Map_FigureDBID_FigurePtr[figureDB.ID]
	if figureStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), figureStaged, figureDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
