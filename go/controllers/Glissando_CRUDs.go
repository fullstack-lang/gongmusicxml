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
var __Glissando__dummysDeclaration__ models.Glissando
var __Glissando_time__dummyDeclaration time.Duration

var mutexGlissando sync.Mutex

// An GlissandoID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGlissando updateGlissando deleteGlissando
type GlissandoID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GlissandoInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGlissando updateGlissando
type GlissandoInput struct {
	// The Glissando to submit or modify
	// in: body
	Glissando *orm.GlissandoAPI
}

// GetGlissandos
//
// swagger:route GET /glissandos glissandos getGlissandos
//
// # Get all glissandos
//
// Responses:
// default: genericError
//
//	200: glissandoDBResponse
func (controller *Controller) GetGlissandos(c *gin.Context) {

	// source slice
	var glissandoDBs []orm.GlissandoDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGlissandos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlissando.GetDB()

	query := db.Find(&glissandoDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	glissandoAPIs := make([]orm.GlissandoAPI, 0)

	// for each glissando, update fields from the database nullable fields
	for idx := range glissandoDBs {
		glissandoDB := &glissandoDBs[idx]
		_ = glissandoDB
		var glissandoAPI orm.GlissandoAPI

		// insertion point for updating fields
		glissandoAPI.ID = glissandoDB.ID
		glissandoDB.CopyBasicFieldsToGlissando_WOP(&glissandoAPI.Glissando_WOP)
		glissandoAPI.GlissandoPointersEncoding = glissandoDB.GlissandoPointersEncoding
		glissandoAPIs = append(glissandoAPIs, glissandoAPI)
	}

	c.JSON(http.StatusOK, glissandoAPIs)
}

// PostGlissando
//
// swagger:route POST /glissandos glissandos postGlissando
//
// Creates a glissando
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGlissando(c *gin.Context) {

	mutexGlissando.Lock()
	defer mutexGlissando.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGlissandos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlissando.GetDB()

	// Validate input
	var input orm.GlissandoAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create glissando
	glissandoDB := orm.GlissandoDB{}
	glissandoDB.GlissandoPointersEncoding = input.GlissandoPointersEncoding
	glissandoDB.CopyBasicFieldsFromGlissando_WOP(&input.Glissando_WOP)

	query := db.Create(&glissandoDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGlissando.CheckoutPhaseOneInstance(&glissandoDB)
	glissando := backRepo.BackRepoGlissando.Map_GlissandoDBID_GlissandoPtr[glissandoDB.ID]

	if glissando != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), glissando)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, glissandoDB)
}

// GetGlissando
//
// swagger:route GET /glissandos/{ID} glissandos getGlissando
//
// Gets the details for a glissando.
//
// Responses:
// default: genericError
//
//	200: glissandoDBResponse
func (controller *Controller) GetGlissando(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGlissando", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlissando.GetDB()

	// Get glissandoDB in DB
	var glissandoDB orm.GlissandoDB
	if err := db.First(&glissandoDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var glissandoAPI orm.GlissandoAPI
	glissandoAPI.ID = glissandoDB.ID
	glissandoAPI.GlissandoPointersEncoding = glissandoDB.GlissandoPointersEncoding
	glissandoDB.CopyBasicFieldsToGlissando_WOP(&glissandoAPI.Glissando_WOP)

	c.JSON(http.StatusOK, glissandoAPI)
}

// UpdateGlissando
//
// swagger:route PATCH /glissandos/{ID} glissandos updateGlissando
//
// # Update a glissando
//
// Responses:
// default: genericError
//
//	200: glissandoDBResponse
func (controller *Controller) UpdateGlissando(c *gin.Context) {

	mutexGlissando.Lock()
	defer mutexGlissando.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGlissando", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlissando.GetDB()

	// Validate input
	var input orm.GlissandoAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var glissandoDB orm.GlissandoDB

	// fetch the glissando
	query := db.First(&glissandoDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	glissandoDB.CopyBasicFieldsFromGlissando_WOP(&input.Glissando_WOP)
	glissandoDB.GlissandoPointersEncoding = input.GlissandoPointersEncoding

	query = db.Model(&glissandoDB).Updates(glissandoDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	glissandoNew := new(models.Glissando)
	glissandoDB.CopyBasicFieldsToGlissando(glissandoNew)

	// redeem pointers
	glissandoDB.DecodePointers(backRepo, glissandoNew)

	// get stage instance from DB instance, and call callback function
	glissandoOld := backRepo.BackRepoGlissando.Map_GlissandoDBID_GlissandoPtr[glissandoDB.ID]
	if glissandoOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), glissandoOld, glissandoNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the glissandoDB
	c.JSON(http.StatusOK, glissandoDB)
}

// DeleteGlissando
//
// swagger:route DELETE /glissandos/{ID} glissandos deleteGlissando
//
// # Delete a glissando
//
// default: genericError
//
//	200: glissandoDBResponse
func (controller *Controller) DeleteGlissando(c *gin.Context) {

	mutexGlissando.Lock()
	defer mutexGlissando.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGlissando", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlissando.GetDB()

	// Get model if exist
	var glissandoDB orm.GlissandoDB
	if err := db.First(&glissandoDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&glissandoDB)

	// get an instance (not staged) from DB instance, and call callback function
	glissandoDeleted := new(models.Glissando)
	glissandoDB.CopyBasicFieldsToGlissando(glissandoDeleted)

	// get stage instance from DB instance, and call callback function
	glissandoStaged := backRepo.BackRepoGlissando.Map_GlissandoDBID_GlissandoPtr[glissandoDB.ID]
	if glissandoStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), glissandoStaged, glissandoDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
