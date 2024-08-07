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
var __Glyph__dummysDeclaration__ models.Glyph
var __Glyph_time__dummyDeclaration time.Duration

var mutexGlyph sync.Mutex

// An GlyphID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGlyph updateGlyph deleteGlyph
type GlyphID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GlyphInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGlyph updateGlyph
type GlyphInput struct {
	// The Glyph to submit or modify
	// in: body
	Glyph *orm.GlyphAPI
}

// GetGlyphs
//
// swagger:route GET /glyphs glyphs getGlyphs
//
// # Get all glyphs
//
// Responses:
// default: genericError
//
//	200: glyphDBResponse
func (controller *Controller) GetGlyphs(c *gin.Context) {

	// source slice
	var glyphDBs []orm.GlyphDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGlyphs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlyph.GetDB()

	query := db.Find(&glyphDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	glyphAPIs := make([]orm.GlyphAPI, 0)

	// for each glyph, update fields from the database nullable fields
	for idx := range glyphDBs {
		glyphDB := &glyphDBs[idx]
		_ = glyphDB
		var glyphAPI orm.GlyphAPI

		// insertion point for updating fields
		glyphAPI.ID = glyphDB.ID
		glyphDB.CopyBasicFieldsToGlyph_WOP(&glyphAPI.Glyph_WOP)
		glyphAPI.GlyphPointersEncoding = glyphDB.GlyphPointersEncoding
		glyphAPIs = append(glyphAPIs, glyphAPI)
	}

	c.JSON(http.StatusOK, glyphAPIs)
}

// PostGlyph
//
// swagger:route POST /glyphs glyphs postGlyph
//
// Creates a glyph
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGlyph(c *gin.Context) {

	mutexGlyph.Lock()
	defer mutexGlyph.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGlyphs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlyph.GetDB()

	// Validate input
	var input orm.GlyphAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create glyph
	glyphDB := orm.GlyphDB{}
	glyphDB.GlyphPointersEncoding = input.GlyphPointersEncoding
	glyphDB.CopyBasicFieldsFromGlyph_WOP(&input.Glyph_WOP)

	query := db.Create(&glyphDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGlyph.CheckoutPhaseOneInstance(&glyphDB)
	glyph := backRepo.BackRepoGlyph.Map_GlyphDBID_GlyphPtr[glyphDB.ID]

	if glyph != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), glyph)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, glyphDB)
}

// GetGlyph
//
// swagger:route GET /glyphs/{ID} glyphs getGlyph
//
// Gets the details for a glyph.
//
// Responses:
// default: genericError
//
//	200: glyphDBResponse
func (controller *Controller) GetGlyph(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGlyph", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlyph.GetDB()

	// Get glyphDB in DB
	var glyphDB orm.GlyphDB
	if err := db.First(&glyphDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var glyphAPI orm.GlyphAPI
	glyphAPI.ID = glyphDB.ID
	glyphAPI.GlyphPointersEncoding = glyphDB.GlyphPointersEncoding
	glyphDB.CopyBasicFieldsToGlyph_WOP(&glyphAPI.Glyph_WOP)

	c.JSON(http.StatusOK, glyphAPI)
}

// UpdateGlyph
//
// swagger:route PATCH /glyphs/{ID} glyphs updateGlyph
//
// # Update a glyph
//
// Responses:
// default: genericError
//
//	200: glyphDBResponse
func (controller *Controller) UpdateGlyph(c *gin.Context) {

	mutexGlyph.Lock()
	defer mutexGlyph.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGlyph", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlyph.GetDB()

	// Validate input
	var input orm.GlyphAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var glyphDB orm.GlyphDB

	// fetch the glyph
	query := db.First(&glyphDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	glyphDB.CopyBasicFieldsFromGlyph_WOP(&input.Glyph_WOP)
	glyphDB.GlyphPointersEncoding = input.GlyphPointersEncoding

	query = db.Model(&glyphDB).Updates(glyphDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	glyphNew := new(models.Glyph)
	glyphDB.CopyBasicFieldsToGlyph(glyphNew)

	// redeem pointers
	glyphDB.DecodePointers(backRepo, glyphNew)

	// get stage instance from DB instance, and call callback function
	glyphOld := backRepo.BackRepoGlyph.Map_GlyphDBID_GlyphPtr[glyphDB.ID]
	if glyphOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), glyphOld, glyphNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the glyphDB
	c.JSON(http.StatusOK, glyphDB)
}

// DeleteGlyph
//
// swagger:route DELETE /glyphs/{ID} glyphs deleteGlyph
//
// # Delete a glyph
//
// default: genericError
//
//	200: glyphDBResponse
func (controller *Controller) DeleteGlyph(c *gin.Context) {

	mutexGlyph.Lock()
	defer mutexGlyph.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGlyph", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGlyph.GetDB()

	// Get model if exist
	var glyphDB orm.GlyphDB
	if err := db.First(&glyphDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&glyphDB)

	// get an instance (not staged) from DB instance, and call callback function
	glyphDeleted := new(models.Glyph)
	glyphDB.CopyBasicFieldsToGlyph(glyphDeleted)

	// get stage instance from DB instance, and call callback function
	glyphStaged := backRepo.BackRepoGlyph.Map_GlyphDBID_GlyphPtr[glyphDB.ID]
	if glyphStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), glyphStaged, glyphDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
