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
var __Lyric_language__dummysDeclaration__ models.Lyric_language
var __Lyric_language_time__dummyDeclaration time.Duration

var mutexLyric_language sync.Mutex

// An Lyric_languageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLyric_language updateLyric_language deleteLyric_language
type Lyric_languageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Lyric_languageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLyric_language updateLyric_language
type Lyric_languageInput struct {
	// The Lyric_language to submit or modify
	// in: body
	Lyric_language *orm.Lyric_languageAPI
}

// GetLyric_languages
//
// swagger:route GET /lyric_languages lyric_languages getLyric_languages
//
// # Get all lyric_languages
//
// Responses:
// default: genericError
//
//	200: lyric_languageDBResponse
func (controller *Controller) GetLyric_languages(c *gin.Context) {

	// source slice
	var lyric_languageDBs []orm.Lyric_languageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLyric_languages", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_language.GetDB()

	query := db.Find(&lyric_languageDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lyric_languageAPIs := make([]orm.Lyric_languageAPI, 0)

	// for each lyric_language, update fields from the database nullable fields
	for idx := range lyric_languageDBs {
		lyric_languageDB := &lyric_languageDBs[idx]
		_ = lyric_languageDB
		var lyric_languageAPI orm.Lyric_languageAPI

		// insertion point for updating fields
		lyric_languageAPI.ID = lyric_languageDB.ID
		lyric_languageDB.CopyBasicFieldsToLyric_language_WOP(&lyric_languageAPI.Lyric_language_WOP)
		lyric_languageAPI.Lyric_languagePointersEncoding = lyric_languageDB.Lyric_languagePointersEncoding
		lyric_languageAPIs = append(lyric_languageAPIs, lyric_languageAPI)
	}

	c.JSON(http.StatusOK, lyric_languageAPIs)
}

// PostLyric_language
//
// swagger:route POST /lyric_languages lyric_languages postLyric_language
//
// Creates a lyric_language
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLyric_language(c *gin.Context) {

	mutexLyric_language.Lock()
	defer mutexLyric_language.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLyric_languages", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_language.GetDB()

	// Validate input
	var input orm.Lyric_languageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create lyric_language
	lyric_languageDB := orm.Lyric_languageDB{}
	lyric_languageDB.Lyric_languagePointersEncoding = input.Lyric_languagePointersEncoding
	lyric_languageDB.CopyBasicFieldsFromLyric_language_WOP(&input.Lyric_language_WOP)

	query := db.Create(&lyric_languageDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLyric_language.CheckoutPhaseOneInstance(&lyric_languageDB)
	lyric_language := backRepo.BackRepoLyric_language.Map_Lyric_languageDBID_Lyric_languagePtr[lyric_languageDB.ID]

	if lyric_language != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), lyric_language)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lyric_languageDB)
}

// GetLyric_language
//
// swagger:route GET /lyric_languages/{ID} lyric_languages getLyric_language
//
// Gets the details for a lyric_language.
//
// Responses:
// default: genericError
//
//	200: lyric_languageDBResponse
func (controller *Controller) GetLyric_language(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLyric_language", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_language.GetDB()

	// Get lyric_languageDB in DB
	var lyric_languageDB orm.Lyric_languageDB
	if err := db.First(&lyric_languageDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lyric_languageAPI orm.Lyric_languageAPI
	lyric_languageAPI.ID = lyric_languageDB.ID
	lyric_languageAPI.Lyric_languagePointersEncoding = lyric_languageDB.Lyric_languagePointersEncoding
	lyric_languageDB.CopyBasicFieldsToLyric_language_WOP(&lyric_languageAPI.Lyric_language_WOP)

	c.JSON(http.StatusOK, lyric_languageAPI)
}

// UpdateLyric_language
//
// swagger:route PATCH /lyric_languages/{ID} lyric_languages updateLyric_language
//
// # Update a lyric_language
//
// Responses:
// default: genericError
//
//	200: lyric_languageDBResponse
func (controller *Controller) UpdateLyric_language(c *gin.Context) {

	mutexLyric_language.Lock()
	defer mutexLyric_language.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLyric_language", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_language.GetDB()

	// Validate input
	var input orm.Lyric_languageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lyric_languageDB orm.Lyric_languageDB

	// fetch the lyric_language
	query := db.First(&lyric_languageDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lyric_languageDB.CopyBasicFieldsFromLyric_language_WOP(&input.Lyric_language_WOP)
	lyric_languageDB.Lyric_languagePointersEncoding = input.Lyric_languagePointersEncoding

	query = db.Model(&lyric_languageDB).Updates(lyric_languageDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lyric_languageNew := new(models.Lyric_language)
	lyric_languageDB.CopyBasicFieldsToLyric_language(lyric_languageNew)

	// redeem pointers
	lyric_languageDB.DecodePointers(backRepo, lyric_languageNew)

	// get stage instance from DB instance, and call callback function
	lyric_languageOld := backRepo.BackRepoLyric_language.Map_Lyric_languageDBID_Lyric_languagePtr[lyric_languageDB.ID]
	if lyric_languageOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lyric_languageOld, lyric_languageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lyric_languageDB
	c.JSON(http.StatusOK, lyric_languageDB)
}

// DeleteLyric_language
//
// swagger:route DELETE /lyric_languages/{ID} lyric_languages deleteLyric_language
//
// # Delete a lyric_language
//
// default: genericError
//
//	200: lyric_languageDBResponse
func (controller *Controller) DeleteLyric_language(c *gin.Context) {

	mutexLyric_language.Lock()
	defer mutexLyric_language.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLyric_language", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric_language.GetDB()

	// Get model if exist
	var lyric_languageDB orm.Lyric_languageDB
	if err := db.First(&lyric_languageDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&lyric_languageDB)

	// get an instance (not staged) from DB instance, and call callback function
	lyric_languageDeleted := new(models.Lyric_language)
	lyric_languageDB.CopyBasicFieldsToLyric_language(lyric_languageDeleted)

	// get stage instance from DB instance, and call callback function
	lyric_languageStaged := backRepo.BackRepoLyric_language.Map_Lyric_languageDBID_Lyric_languagePtr[lyric_languageDB.ID]
	if lyric_languageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lyric_languageStaged, lyric_languageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
