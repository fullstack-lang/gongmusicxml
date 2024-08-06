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
var __Fo__dummysDeclaration__ models.Fo
var __Fo_time__dummyDeclaration time.Duration

var mutexFo sync.Mutex

// An FoID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFo updateFo deleteFo
type FoID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FoInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFo updateFo
type FoInput struct {
	// The Fo to submit or modify
	// in: body
	Fo *orm.FoAPI
}

// GetFos
//
// swagger:route GET /fos fos getFos
//
// # Get all fos
//
// Responses:
// default: genericError
//
//	200: foDBResponse
func (controller *Controller) GetFos(c *gin.Context) {

	// source slice
	var foDBs []orm.FoDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFo.GetDB()

	query := db.Find(&foDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	foAPIs := make([]orm.FoAPI, 0)

	// for each fo, update fields from the database nullable fields
	for idx := range foDBs {
		foDB := &foDBs[idx]
		_ = foDB
		var foAPI orm.FoAPI

		// insertion point for updating fields
		foAPI.ID = foDB.ID
		foDB.CopyBasicFieldsToFo_WOP(&foAPI.Fo_WOP)
		foAPI.FoPointersEncoding = foDB.FoPointersEncoding
		foAPIs = append(foAPIs, foAPI)
	}

	c.JSON(http.StatusOK, foAPIs)
}

// PostFo
//
// swagger:route POST /fos fos postFo
//
// Creates a fo
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFo(c *gin.Context) {

	mutexFo.Lock()
	defer mutexFo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFo.GetDB()

	// Validate input
	var input orm.FoAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create fo
	foDB := orm.FoDB{}
	foDB.FoPointersEncoding = input.FoPointersEncoding
	foDB.CopyBasicFieldsFromFo_WOP(&input.Fo_WOP)

	query := db.Create(&foDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFo.CheckoutPhaseOneInstance(&foDB)
	fo := backRepo.BackRepoFo.Map_FoDBID_FoPtr[foDB.ID]

	if fo != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), fo)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, foDB)
}

// GetFo
//
// swagger:route GET /fos/{ID} fos getFo
//
// Gets the details for a fo.
//
// Responses:
// default: genericError
//
//	200: foDBResponse
func (controller *Controller) GetFo(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFo.GetDB()

	// Get foDB in DB
	var foDB orm.FoDB
	if err := db.First(&foDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var foAPI orm.FoAPI
	foAPI.ID = foDB.ID
	foAPI.FoPointersEncoding = foDB.FoPointersEncoding
	foDB.CopyBasicFieldsToFo_WOP(&foAPI.Fo_WOP)

	c.JSON(http.StatusOK, foAPI)
}

// UpdateFo
//
// swagger:route PATCH /fos/{ID} fos updateFo
//
// # Update a fo
//
// Responses:
// default: genericError
//
//	200: foDBResponse
func (controller *Controller) UpdateFo(c *gin.Context) {

	mutexFo.Lock()
	defer mutexFo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFo.GetDB()

	// Validate input
	var input orm.FoAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var foDB orm.FoDB

	// fetch the fo
	query := db.First(&foDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	foDB.CopyBasicFieldsFromFo_WOP(&input.Fo_WOP)
	foDB.FoPointersEncoding = input.FoPointersEncoding

	query = db.Model(&foDB).Updates(foDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	foNew := new(models.Fo)
	foDB.CopyBasicFieldsToFo(foNew)

	// redeem pointers
	foDB.DecodePointers(backRepo, foNew)

	// get stage instance from DB instance, and call callback function
	foOld := backRepo.BackRepoFo.Map_FoDBID_FoPtr[foDB.ID]
	if foOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), foOld, foNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the foDB
	c.JSON(http.StatusOK, foDB)
}

// DeleteFo
//
// swagger:route DELETE /fos/{ID} fos deleteFo
//
// # Delete a fo
//
// default: genericError
//
//	200: foDBResponse
func (controller *Controller) DeleteFo(c *gin.Context) {

	mutexFo.Lock()
	defer mutexFo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFo.GetDB()

	// Get model if exist
	var foDB orm.FoDB
	if err := db.First(&foDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&foDB)

	// get an instance (not staged) from DB instance, and call callback function
	foDeleted := new(models.Fo)
	foDB.CopyBasicFieldsToFo(foDeleted)

	// get stage instance from DB instance, and call callback function
	foStaged := backRepo.BackRepoFo.Map_FoDBID_FoPtr[foDB.ID]
	if foStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), foStaged, foDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
