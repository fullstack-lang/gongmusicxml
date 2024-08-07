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
var __Stem__dummysDeclaration__ models.Stem
var __Stem_time__dummyDeclaration time.Duration

var mutexStem sync.Mutex

// An StemID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStem updateStem deleteStem
type StemID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StemInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStem updateStem
type StemInput struct {
	// The Stem to submit or modify
	// in: body
	Stem *orm.StemAPI
}

// GetStems
//
// swagger:route GET /stems stems getStems
//
// # Get all stems
//
// Responses:
// default: genericError
//
//	200: stemDBResponse
func (controller *Controller) GetStems(c *gin.Context) {

	// source slice
	var stemDBs []orm.StemDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStems", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStem.GetDB()

	query := db.Find(&stemDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	stemAPIs := make([]orm.StemAPI, 0)

	// for each stem, update fields from the database nullable fields
	for idx := range stemDBs {
		stemDB := &stemDBs[idx]
		_ = stemDB
		var stemAPI orm.StemAPI

		// insertion point for updating fields
		stemAPI.ID = stemDB.ID
		stemDB.CopyBasicFieldsToStem_WOP(&stemAPI.Stem_WOP)
		stemAPI.StemPointersEncoding = stemDB.StemPointersEncoding
		stemAPIs = append(stemAPIs, stemAPI)
	}

	c.JSON(http.StatusOK, stemAPIs)
}

// PostStem
//
// swagger:route POST /stems stems postStem
//
// Creates a stem
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStem(c *gin.Context) {

	mutexStem.Lock()
	defer mutexStem.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStems", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStem.GetDB()

	// Validate input
	var input orm.StemAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create stem
	stemDB := orm.StemDB{}
	stemDB.StemPointersEncoding = input.StemPointersEncoding
	stemDB.CopyBasicFieldsFromStem_WOP(&input.Stem_WOP)

	query := db.Create(&stemDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStem.CheckoutPhaseOneInstance(&stemDB)
	stem := backRepo.BackRepoStem.Map_StemDBID_StemPtr[stemDB.ID]

	if stem != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), stem)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, stemDB)
}

// GetStem
//
// swagger:route GET /stems/{ID} stems getStem
//
// Gets the details for a stem.
//
// Responses:
// default: genericError
//
//	200: stemDBResponse
func (controller *Controller) GetStem(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStem", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStem.GetDB()

	// Get stemDB in DB
	var stemDB orm.StemDB
	if err := db.First(&stemDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var stemAPI orm.StemAPI
	stemAPI.ID = stemDB.ID
	stemAPI.StemPointersEncoding = stemDB.StemPointersEncoding
	stemDB.CopyBasicFieldsToStem_WOP(&stemAPI.Stem_WOP)

	c.JSON(http.StatusOK, stemAPI)
}

// UpdateStem
//
// swagger:route PATCH /stems/{ID} stems updateStem
//
// # Update a stem
//
// Responses:
// default: genericError
//
//	200: stemDBResponse
func (controller *Controller) UpdateStem(c *gin.Context) {

	mutexStem.Lock()
	defer mutexStem.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStem", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStem.GetDB()

	// Validate input
	var input orm.StemAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var stemDB orm.StemDB

	// fetch the stem
	query := db.First(&stemDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	stemDB.CopyBasicFieldsFromStem_WOP(&input.Stem_WOP)
	stemDB.StemPointersEncoding = input.StemPointersEncoding

	query = db.Model(&stemDB).Updates(stemDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	stemNew := new(models.Stem)
	stemDB.CopyBasicFieldsToStem(stemNew)

	// redeem pointers
	stemDB.DecodePointers(backRepo, stemNew)

	// get stage instance from DB instance, and call callback function
	stemOld := backRepo.BackRepoStem.Map_StemDBID_StemPtr[stemDB.ID]
	if stemOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), stemOld, stemNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the stemDB
	c.JSON(http.StatusOK, stemDB)
}

// DeleteStem
//
// swagger:route DELETE /stems/{ID} stems deleteStem
//
// # Delete a stem
//
// default: genericError
//
//	200: stemDBResponse
func (controller *Controller) DeleteStem(c *gin.Context) {

	mutexStem.Lock()
	defer mutexStem.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStem", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStem.GetDB()

	// Get model if exist
	var stemDB orm.StemDB
	if err := db.First(&stemDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&stemDB)

	// get an instance (not staged) from DB instance, and call callback function
	stemDeleted := new(models.Stem)
	stemDB.CopyBasicFieldsToStem(stemDeleted)

	// get stage instance from DB instance, and call callback function
	stemStaged := backRepo.BackRepoStem.Map_StemDBID_StemPtr[stemDB.ID]
	if stemStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), stemStaged, stemDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
