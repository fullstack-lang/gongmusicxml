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
var __Lyric__dummysDeclaration__ models.Lyric
var __Lyric_time__dummyDeclaration time.Duration

var mutexLyric sync.Mutex

// An LyricID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLyric updateLyric deleteLyric
type LyricID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LyricInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLyric updateLyric
type LyricInput struct {
	// The Lyric to submit or modify
	// in: body
	Lyric *orm.LyricAPI
}

// GetLyrics
//
// swagger:route GET /lyrics lyrics getLyrics
//
// # Get all lyrics
//
// Responses:
// default: genericError
//
//	200: lyricDBResponse
func (controller *Controller) GetLyrics(c *gin.Context) {

	// source slice
	var lyricDBs []orm.LyricDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLyrics", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric.GetDB()

	query := db.Find(&lyricDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lyricAPIs := make([]orm.LyricAPI, 0)

	// for each lyric, update fields from the database nullable fields
	for idx := range lyricDBs {
		lyricDB := &lyricDBs[idx]
		_ = lyricDB
		var lyricAPI orm.LyricAPI

		// insertion point for updating fields
		lyricAPI.ID = lyricDB.ID
		lyricDB.CopyBasicFieldsToLyric_WOP(&lyricAPI.Lyric_WOP)
		lyricAPI.LyricPointersEncoding = lyricDB.LyricPointersEncoding
		lyricAPIs = append(lyricAPIs, lyricAPI)
	}

	c.JSON(http.StatusOK, lyricAPIs)
}

// PostLyric
//
// swagger:route POST /lyrics lyrics postLyric
//
// Creates a lyric
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLyric(c *gin.Context) {

	mutexLyric.Lock()
	defer mutexLyric.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLyrics", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric.GetDB()

	// Validate input
	var input orm.LyricAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create lyric
	lyricDB := orm.LyricDB{}
	lyricDB.LyricPointersEncoding = input.LyricPointersEncoding
	lyricDB.CopyBasicFieldsFromLyric_WOP(&input.Lyric_WOP)

	query := db.Create(&lyricDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLyric.CheckoutPhaseOneInstance(&lyricDB)
	lyric := backRepo.BackRepoLyric.Map_LyricDBID_LyricPtr[lyricDB.ID]

	if lyric != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), lyric)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lyricDB)
}

// GetLyric
//
// swagger:route GET /lyrics/{ID} lyrics getLyric
//
// Gets the details for a lyric.
//
// Responses:
// default: genericError
//
//	200: lyricDBResponse
func (controller *Controller) GetLyric(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLyric", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric.GetDB()

	// Get lyricDB in DB
	var lyricDB orm.LyricDB
	if err := db.First(&lyricDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lyricAPI orm.LyricAPI
	lyricAPI.ID = lyricDB.ID
	lyricAPI.LyricPointersEncoding = lyricDB.LyricPointersEncoding
	lyricDB.CopyBasicFieldsToLyric_WOP(&lyricAPI.Lyric_WOP)

	c.JSON(http.StatusOK, lyricAPI)
}

// UpdateLyric
//
// swagger:route PATCH /lyrics/{ID} lyrics updateLyric
//
// # Update a lyric
//
// Responses:
// default: genericError
//
//	200: lyricDBResponse
func (controller *Controller) UpdateLyric(c *gin.Context) {

	mutexLyric.Lock()
	defer mutexLyric.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLyric", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric.GetDB()

	// Validate input
	var input orm.LyricAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lyricDB orm.LyricDB

	// fetch the lyric
	query := db.First(&lyricDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lyricDB.CopyBasicFieldsFromLyric_WOP(&input.Lyric_WOP)
	lyricDB.LyricPointersEncoding = input.LyricPointersEncoding

	query = db.Model(&lyricDB).Updates(lyricDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lyricNew := new(models.Lyric)
	lyricDB.CopyBasicFieldsToLyric(lyricNew)

	// redeem pointers
	lyricDB.DecodePointers(backRepo, lyricNew)

	// get stage instance from DB instance, and call callback function
	lyricOld := backRepo.BackRepoLyric.Map_LyricDBID_LyricPtr[lyricDB.ID]
	if lyricOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lyricOld, lyricNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lyricDB
	c.JSON(http.StatusOK, lyricDB)
}

// DeleteLyric
//
// swagger:route DELETE /lyrics/{ID} lyrics deleteLyric
//
// # Delete a lyric
//
// default: genericError
//
//	200: lyricDBResponse
func (controller *Controller) DeleteLyric(c *gin.Context) {

	mutexLyric.Lock()
	defer mutexLyric.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLyric", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLyric.GetDB()

	// Get model if exist
	var lyricDB orm.LyricDB
	if err := db.First(&lyricDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&lyricDB)

	// get an instance (not staged) from DB instance, and call callback function
	lyricDeleted := new(models.Lyric)
	lyricDB.CopyBasicFieldsToLyric(lyricDeleted)

	// get stage instance from DB instance, and call callback function
	lyricStaged := backRepo.BackRepoLyric.Map_LyricDBID_LyricPtr[lyricDB.ID]
	if lyricStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lyricStaged, lyricDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
