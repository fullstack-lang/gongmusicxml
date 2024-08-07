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
var __Glass__dummysDeclaration__ models.Glass
var __Glass_time__dummyDeclaration time.Duration

var mutexGlass sync.Mutex

// An GlassID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGlass updateGlass deleteGlass
type GlassID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GlassInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGlass updateGlass
type GlassInput struct {
	// The Glass to submit or modify
	// in: body
	Glass *orm.GlassAPI
}

// GetGlasss
//
// swagger:route GET /glasss glasss getGlasss
//
// # Get all glasss
//
// Responses:
// default: genericError
//
//	200: glassDBResponse
func (controller *Controller) GetGlasss(c *gin.Context) {

	// source slice
	var glassDBs []orm.GlassDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGlasss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlass.GetDB()

	query := db.Find(&glassDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	glassAPIs := make([]orm.GlassAPI, 0)

	// for each glass, update fields from the database nullable fields
	for idx := range glassDBs {
		glassDB := &glassDBs[idx]
		_ = glassDB
		var glassAPI orm.GlassAPI

		// insertion point for updating fields
		glassAPI.ID = glassDB.ID
		glassDB.CopyBasicFieldsToGlass_WOP(&glassAPI.Glass_WOP)
		glassAPI.GlassPointersEncoding = glassDB.GlassPointersEncoding
		glassAPIs = append(glassAPIs, glassAPI)
	}

	c.JSON(http.StatusOK, glassAPIs)
}

// PostGlass
//
// swagger:route POST /glasss glasss postGlass
//
// Creates a glass
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGlass(c *gin.Context) {

	mutexGlass.Lock()
	defer mutexGlass.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGlasss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlass.GetDB()

	// Validate input
	var input orm.GlassAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create glass
	glassDB := orm.GlassDB{}
	glassDB.GlassPointersEncoding = input.GlassPointersEncoding
	glassDB.CopyBasicFieldsFromGlass_WOP(&input.Glass_WOP)

	query := db.Create(&glassDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGlass.CheckoutPhaseOneInstance(&glassDB)
	glass := backRepo.BackRepoGlass.Map_GlassDBID_GlassPtr[glassDB.ID]

	if glass != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), glass)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, glassDB)
}

// GetGlass
//
// swagger:route GET /glasss/{ID} glasss getGlass
//
// Gets the details for a glass.
//
// Responses:
// default: genericError
//
//	200: glassDBResponse
func (controller *Controller) GetGlass(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGlass", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlass.GetDB()

	// Get glassDB in DB
	var glassDB orm.GlassDB
	if err := db.First(&glassDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var glassAPI orm.GlassAPI
	glassAPI.ID = glassDB.ID
	glassAPI.GlassPointersEncoding = glassDB.GlassPointersEncoding
	glassDB.CopyBasicFieldsToGlass_WOP(&glassAPI.Glass_WOP)

	c.JSON(http.StatusOK, glassAPI)
}

// UpdateGlass
//
// swagger:route PATCH /glasss/{ID} glasss updateGlass
//
// # Update a glass
//
// Responses:
// default: genericError
//
//	200: glassDBResponse
func (controller *Controller) UpdateGlass(c *gin.Context) {

	mutexGlass.Lock()
	defer mutexGlass.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGlass", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlass.GetDB()

	// Validate input
	var input orm.GlassAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var glassDB orm.GlassDB

	// fetch the glass
	query := db.First(&glassDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	glassDB.CopyBasicFieldsFromGlass_WOP(&input.Glass_WOP)
	glassDB.GlassPointersEncoding = input.GlassPointersEncoding

	query = db.Model(&glassDB).Updates(glassDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	glassNew := new(models.Glass)
	glassDB.CopyBasicFieldsToGlass(glassNew)

	// redeem pointers
	glassDB.DecodePointers(backRepo, glassNew)

	// get stage instance from DB instance, and call callback function
	glassOld := backRepo.BackRepoGlass.Map_GlassDBID_GlassPtr[glassDB.ID]
	if glassOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), glassOld, glassNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the glassDB
	c.JSON(http.StatusOK, glassDB)
}

// DeleteGlass
//
// swagger:route DELETE /glasss/{ID} glasss deleteGlass
//
// # Delete a glass
//
// default: genericError
//
//	200: glassDBResponse
func (controller *Controller) DeleteGlass(c *gin.Context) {

	mutexGlass.Lock()
	defer mutexGlass.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGlass", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlass.GetDB()

	// Get model if exist
	var glassDB orm.GlassDB
	if err := db.First(&glassDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&glassDB)

	// get an instance (not staged) from DB instance, and call callback function
	glassDeleted := new(models.Glass)
	glassDB.CopyBasicFieldsToGlass(glassDeleted)

	// get stage instance from DB instance, and call callback function
	glassStaged := backRepo.BackRepoGlass.Map_GlassDBID_GlassPtr[glassDB.ID]
	if glassStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), glassStaged, glassDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
