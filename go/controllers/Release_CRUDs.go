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
var __Release__dummysDeclaration__ models.Release
var __Release_time__dummyDeclaration time.Duration

var mutexRelease sync.Mutex

// An ReleaseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRelease updateRelease deleteRelease
type ReleaseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ReleaseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRelease updateRelease
type ReleaseInput struct {
	// The Release to submit or modify
	// in: body
	Release *orm.ReleaseAPI
}

// GetReleases
//
// swagger:route GET /releases releases getReleases
//
// # Get all releases
//
// Responses:
// default: genericError
//
//	200: releaseDBResponse
func (controller *Controller) GetReleases(c *gin.Context) {

	// source slice
	var releaseDBs []orm.ReleaseDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetReleases", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRelease.GetDB()

	query := db.Find(&releaseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	releaseAPIs := make([]orm.ReleaseAPI, 0)

	// for each release, update fields from the database nullable fields
	for idx := range releaseDBs {
		releaseDB := &releaseDBs[idx]
		_ = releaseDB
		var releaseAPI orm.ReleaseAPI

		// insertion point for updating fields
		releaseAPI.ID = releaseDB.ID
		releaseDB.CopyBasicFieldsToRelease_WOP(&releaseAPI.Release_WOP)
		releaseAPI.ReleasePointersEncoding = releaseDB.ReleasePointersEncoding
		releaseAPIs = append(releaseAPIs, releaseAPI)
	}

	c.JSON(http.StatusOK, releaseAPIs)
}

// PostRelease
//
// swagger:route POST /releases releases postRelease
//
// Creates a release
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRelease(c *gin.Context) {

	mutexRelease.Lock()
	defer mutexRelease.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostReleases", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRelease.GetDB()

	// Validate input
	var input orm.ReleaseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create release
	releaseDB := orm.ReleaseDB{}
	releaseDB.ReleasePointersEncoding = input.ReleasePointersEncoding
	releaseDB.CopyBasicFieldsFromRelease_WOP(&input.Release_WOP)

	query := db.Create(&releaseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRelease.CheckoutPhaseOneInstance(&releaseDB)
	release := backRepo.BackRepoRelease.Map_ReleaseDBID_ReleasePtr[releaseDB.ID]

	if release != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), release)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, releaseDB)
}

// GetRelease
//
// swagger:route GET /releases/{ID} releases getRelease
//
// Gets the details for a release.
//
// Responses:
// default: genericError
//
//	200: releaseDBResponse
func (controller *Controller) GetRelease(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRelease", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRelease.GetDB()

	// Get releaseDB in DB
	var releaseDB orm.ReleaseDB
	if err := db.First(&releaseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var releaseAPI orm.ReleaseAPI
	releaseAPI.ID = releaseDB.ID
	releaseAPI.ReleasePointersEncoding = releaseDB.ReleasePointersEncoding
	releaseDB.CopyBasicFieldsToRelease_WOP(&releaseAPI.Release_WOP)

	c.JSON(http.StatusOK, releaseAPI)
}

// UpdateRelease
//
// swagger:route PATCH /releases/{ID} releases updateRelease
//
// # Update a release
//
// Responses:
// default: genericError
//
//	200: releaseDBResponse
func (controller *Controller) UpdateRelease(c *gin.Context) {

	mutexRelease.Lock()
	defer mutexRelease.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRelease", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRelease.GetDB()

	// Validate input
	var input orm.ReleaseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var releaseDB orm.ReleaseDB

	// fetch the release
	query := db.First(&releaseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	releaseDB.CopyBasicFieldsFromRelease_WOP(&input.Release_WOP)
	releaseDB.ReleasePointersEncoding = input.ReleasePointersEncoding

	query = db.Model(&releaseDB).Updates(releaseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	releaseNew := new(models.Release)
	releaseDB.CopyBasicFieldsToRelease(releaseNew)

	// redeem pointers
	releaseDB.DecodePointers(backRepo, releaseNew)

	// get stage instance from DB instance, and call callback function
	releaseOld := backRepo.BackRepoRelease.Map_ReleaseDBID_ReleasePtr[releaseDB.ID]
	if releaseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), releaseOld, releaseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the releaseDB
	c.JSON(http.StatusOK, releaseDB)
}

// DeleteRelease
//
// swagger:route DELETE /releases/{ID} releases deleteRelease
//
// # Delete a release
//
// default: genericError
//
//	200: releaseDBResponse
func (controller *Controller) DeleteRelease(c *gin.Context) {

	mutexRelease.Lock()
	defer mutexRelease.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRelease", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRelease.GetDB()

	// Get model if exist
	var releaseDB orm.ReleaseDB
	if err := db.First(&releaseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&releaseDB)

	// get an instance (not staged) from DB instance, and call callback function
	releaseDeleted := new(models.Release)
	releaseDB.CopyBasicFieldsToRelease(releaseDeleted)

	// get stage instance from DB instance, and call callback function
	releaseStaged := backRepo.BackRepoRelease.Map_ReleaseDBID_ReleasePtr[releaseDB.ID]
	if releaseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), releaseStaged, releaseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
