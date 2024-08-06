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
var __Lyric_font__dummysDeclaration__ models.Lyric_font
var __Lyric_font_time__dummyDeclaration time.Duration

var mutexLyric_font sync.Mutex

// An Lyric_fontID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLyric_font updateLyric_font deleteLyric_font
type Lyric_fontID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Lyric_fontInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLyric_font updateLyric_font
type Lyric_fontInput struct {
	// The Lyric_font to submit or modify
	// in: body
	Lyric_font *orm.Lyric_fontAPI
}

// GetLyric_fonts
//
// swagger:route GET /lyric_fonts lyric_fonts getLyric_fonts
//
// # Get all lyric_fonts
//
// Responses:
// default: genericError
//
//	200: lyric_fontDBResponse
func (controller *Controller) GetLyric_fonts(c *gin.Context) {

	// source slice
	var lyric_fontDBs []orm.Lyric_fontDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLyric_fonts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_font.GetDB()

	query := db.Find(&lyric_fontDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lyric_fontAPIs := make([]orm.Lyric_fontAPI, 0)

	// for each lyric_font, update fields from the database nullable fields
	for idx := range lyric_fontDBs {
		lyric_fontDB := &lyric_fontDBs[idx]
		_ = lyric_fontDB
		var lyric_fontAPI orm.Lyric_fontAPI

		// insertion point for updating fields
		lyric_fontAPI.ID = lyric_fontDB.ID
		lyric_fontDB.CopyBasicFieldsToLyric_font_WOP(&lyric_fontAPI.Lyric_font_WOP)
		lyric_fontAPI.Lyric_fontPointersEncoding = lyric_fontDB.Lyric_fontPointersEncoding
		lyric_fontAPIs = append(lyric_fontAPIs, lyric_fontAPI)
	}

	c.JSON(http.StatusOK, lyric_fontAPIs)
}

// PostLyric_font
//
// swagger:route POST /lyric_fonts lyric_fonts postLyric_font
//
// Creates a lyric_font
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLyric_font(c *gin.Context) {

	mutexLyric_font.Lock()
	defer mutexLyric_font.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLyric_fonts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_font.GetDB()

	// Validate input
	var input orm.Lyric_fontAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create lyric_font
	lyric_fontDB := orm.Lyric_fontDB{}
	lyric_fontDB.Lyric_fontPointersEncoding = input.Lyric_fontPointersEncoding
	lyric_fontDB.CopyBasicFieldsFromLyric_font_WOP(&input.Lyric_font_WOP)

	query := db.Create(&lyric_fontDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLyric_font.CheckoutPhaseOneInstance(&lyric_fontDB)
	lyric_font := backRepo.BackRepoLyric_font.Map_Lyric_fontDBID_Lyric_fontPtr[lyric_fontDB.ID]

	if lyric_font != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), lyric_font)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lyric_fontDB)
}

// GetLyric_font
//
// swagger:route GET /lyric_fonts/{ID} lyric_fonts getLyric_font
//
// Gets the details for a lyric_font.
//
// Responses:
// default: genericError
//
//	200: lyric_fontDBResponse
func (controller *Controller) GetLyric_font(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLyric_font", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_font.GetDB()

	// Get lyric_fontDB in DB
	var lyric_fontDB orm.Lyric_fontDB
	if err := db.First(&lyric_fontDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lyric_fontAPI orm.Lyric_fontAPI
	lyric_fontAPI.ID = lyric_fontDB.ID
	lyric_fontAPI.Lyric_fontPointersEncoding = lyric_fontDB.Lyric_fontPointersEncoding
	lyric_fontDB.CopyBasicFieldsToLyric_font_WOP(&lyric_fontAPI.Lyric_font_WOP)

	c.JSON(http.StatusOK, lyric_fontAPI)
}

// UpdateLyric_font
//
// swagger:route PATCH /lyric_fonts/{ID} lyric_fonts updateLyric_font
//
// # Update a lyric_font
//
// Responses:
// default: genericError
//
//	200: lyric_fontDBResponse
func (controller *Controller) UpdateLyric_font(c *gin.Context) {

	mutexLyric_font.Lock()
	defer mutexLyric_font.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLyric_font", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_font.GetDB()

	// Validate input
	var input orm.Lyric_fontAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lyric_fontDB orm.Lyric_fontDB

	// fetch the lyric_font
	query := db.First(&lyric_fontDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lyric_fontDB.CopyBasicFieldsFromLyric_font_WOP(&input.Lyric_font_WOP)
	lyric_fontDB.Lyric_fontPointersEncoding = input.Lyric_fontPointersEncoding

	query = db.Model(&lyric_fontDB).Updates(lyric_fontDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lyric_fontNew := new(models.Lyric_font)
	lyric_fontDB.CopyBasicFieldsToLyric_font(lyric_fontNew)

	// redeem pointers
	lyric_fontDB.DecodePointers(backRepo, lyric_fontNew)

	// get stage instance from DB instance, and call callback function
	lyric_fontOld := backRepo.BackRepoLyric_font.Map_Lyric_fontDBID_Lyric_fontPtr[lyric_fontDB.ID]
	if lyric_fontOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lyric_fontOld, lyric_fontNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lyric_fontDB
	c.JSON(http.StatusOK, lyric_fontDB)
}

// DeleteLyric_font
//
// swagger:route DELETE /lyric_fonts/{ID} lyric_fonts deleteLyric_font
//
// # Delete a lyric_font
//
// default: genericError
//
//	200: lyric_fontDBResponse
func (controller *Controller) DeleteLyric_font(c *gin.Context) {

	mutexLyric_font.Lock()
	defer mutexLyric_font.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLyric_font", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_font.GetDB()

	// Get model if exist
	var lyric_fontDB orm.Lyric_fontDB
	if err := db.First(&lyric_fontDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&lyric_fontDB)

	// get an instance (not staged) from DB instance, and call callback function
	lyric_fontDeleted := new(models.Lyric_font)
	lyric_fontDB.CopyBasicFieldsToLyric_font(lyric_fontDeleted)

	// get stage instance from DB instance, and call callback function
	lyric_fontStaged := backRepo.BackRepoLyric_font.Map_Lyric_fontDBID_Lyric_fontPtr[lyric_fontDB.ID]
	if lyric_fontStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lyric_fontStaged, lyric_fontDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
