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
var __Non_arpeggiate__dummysDeclaration__ models.Non_arpeggiate
var __Non_arpeggiate_time__dummyDeclaration time.Duration

var mutexNon_arpeggiate sync.Mutex

// An Non_arpeggiateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNon_arpeggiate updateNon_arpeggiate deleteNon_arpeggiate
type Non_arpeggiateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Non_arpeggiateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNon_arpeggiate updateNon_arpeggiate
type Non_arpeggiateInput struct {
	// The Non_arpeggiate to submit or modify
	// in: body
	Non_arpeggiate *orm.Non_arpeggiateAPI
}

// GetNon_arpeggiates
//
// swagger:route GET /non_arpeggiates non_arpeggiates getNon_arpeggiates
//
// # Get all non_arpeggiates
//
// Responses:
// default: genericError
//
//	200: non_arpeggiateDBResponse
func (controller *Controller) GetNon_arpeggiates(c *gin.Context) {

	// source slice
	var non_arpeggiateDBs []orm.Non_arpeggiateDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNon_arpeggiates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNon_arpeggiate.GetDB()

	query := db.Find(&non_arpeggiateDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	non_arpeggiateAPIs := make([]orm.Non_arpeggiateAPI, 0)

	// for each non_arpeggiate, update fields from the database nullable fields
	for idx := range non_arpeggiateDBs {
		non_arpeggiateDB := &non_arpeggiateDBs[idx]
		_ = non_arpeggiateDB
		var non_arpeggiateAPI orm.Non_arpeggiateAPI

		// insertion point for updating fields
		non_arpeggiateAPI.ID = non_arpeggiateDB.ID
		non_arpeggiateDB.CopyBasicFieldsToNon_arpeggiate_WOP(&non_arpeggiateAPI.Non_arpeggiate_WOP)
		non_arpeggiateAPI.Non_arpeggiatePointersEncoding = non_arpeggiateDB.Non_arpeggiatePointersEncoding
		non_arpeggiateAPIs = append(non_arpeggiateAPIs, non_arpeggiateAPI)
	}

	c.JSON(http.StatusOK, non_arpeggiateAPIs)
}

// PostNon_arpeggiate
//
// swagger:route POST /non_arpeggiates non_arpeggiates postNon_arpeggiate
//
// Creates a non_arpeggiate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNon_arpeggiate(c *gin.Context) {

	mutexNon_arpeggiate.Lock()
	defer mutexNon_arpeggiate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNon_arpeggiates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNon_arpeggiate.GetDB()

	// Validate input
	var input orm.Non_arpeggiateAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create non_arpeggiate
	non_arpeggiateDB := orm.Non_arpeggiateDB{}
	non_arpeggiateDB.Non_arpeggiatePointersEncoding = input.Non_arpeggiatePointersEncoding
	non_arpeggiateDB.CopyBasicFieldsFromNon_arpeggiate_WOP(&input.Non_arpeggiate_WOP)

	query := db.Create(&non_arpeggiateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNon_arpeggiate.CheckoutPhaseOneInstance(&non_arpeggiateDB)
	non_arpeggiate := backRepo.BackRepoNon_arpeggiate.Map_Non_arpeggiateDBID_Non_arpeggiatePtr[non_arpeggiateDB.ID]

	if non_arpeggiate != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), non_arpeggiate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, non_arpeggiateDB)
}

// GetNon_arpeggiate
//
// swagger:route GET /non_arpeggiates/{ID} non_arpeggiates getNon_arpeggiate
//
// Gets the details for a non_arpeggiate.
//
// Responses:
// default: genericError
//
//	200: non_arpeggiateDBResponse
func (controller *Controller) GetNon_arpeggiate(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNon_arpeggiate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNon_arpeggiate.GetDB()

	// Get non_arpeggiateDB in DB
	var non_arpeggiateDB orm.Non_arpeggiateDB
	if err := db.First(&non_arpeggiateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var non_arpeggiateAPI orm.Non_arpeggiateAPI
	non_arpeggiateAPI.ID = non_arpeggiateDB.ID
	non_arpeggiateAPI.Non_arpeggiatePointersEncoding = non_arpeggiateDB.Non_arpeggiatePointersEncoding
	non_arpeggiateDB.CopyBasicFieldsToNon_arpeggiate_WOP(&non_arpeggiateAPI.Non_arpeggiate_WOP)

	c.JSON(http.StatusOK, non_arpeggiateAPI)
}

// UpdateNon_arpeggiate
//
// swagger:route PATCH /non_arpeggiates/{ID} non_arpeggiates updateNon_arpeggiate
//
// # Update a non_arpeggiate
//
// Responses:
// default: genericError
//
//	200: non_arpeggiateDBResponse
func (controller *Controller) UpdateNon_arpeggiate(c *gin.Context) {

	mutexNon_arpeggiate.Lock()
	defer mutexNon_arpeggiate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNon_arpeggiate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNon_arpeggiate.GetDB()

	// Validate input
	var input orm.Non_arpeggiateAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var non_arpeggiateDB orm.Non_arpeggiateDB

	// fetch the non_arpeggiate
	query := db.First(&non_arpeggiateDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	non_arpeggiateDB.CopyBasicFieldsFromNon_arpeggiate_WOP(&input.Non_arpeggiate_WOP)
	non_arpeggiateDB.Non_arpeggiatePointersEncoding = input.Non_arpeggiatePointersEncoding

	query = db.Model(&non_arpeggiateDB).Updates(non_arpeggiateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	non_arpeggiateNew := new(models.Non_arpeggiate)
	non_arpeggiateDB.CopyBasicFieldsToNon_arpeggiate(non_arpeggiateNew)

	// redeem pointers
	non_arpeggiateDB.DecodePointers(backRepo, non_arpeggiateNew)

	// get stage instance from DB instance, and call callback function
	non_arpeggiateOld := backRepo.BackRepoNon_arpeggiate.Map_Non_arpeggiateDBID_Non_arpeggiatePtr[non_arpeggiateDB.ID]
	if non_arpeggiateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), non_arpeggiateOld, non_arpeggiateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the non_arpeggiateDB
	c.JSON(http.StatusOK, non_arpeggiateDB)
}

// DeleteNon_arpeggiate
//
// swagger:route DELETE /non_arpeggiates/{ID} non_arpeggiates deleteNon_arpeggiate
//
// # Delete a non_arpeggiate
//
// default: genericError
//
//	200: non_arpeggiateDBResponse
func (controller *Controller) DeleteNon_arpeggiate(c *gin.Context) {

	mutexNon_arpeggiate.Lock()
	defer mutexNon_arpeggiate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNon_arpeggiate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNon_arpeggiate.GetDB()

	// Get model if exist
	var non_arpeggiateDB orm.Non_arpeggiateDB
	if err := db.First(&non_arpeggiateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&non_arpeggiateDB)

	// get an instance (not staged) from DB instance, and call callback function
	non_arpeggiateDeleted := new(models.Non_arpeggiate)
	non_arpeggiateDB.CopyBasicFieldsToNon_arpeggiate(non_arpeggiateDeleted)

	// get stage instance from DB instance, and call callback function
	non_arpeggiateStaged := backRepo.BackRepoNon_arpeggiate.Map_Non_arpeggiateDBID_Non_arpeggiatePtr[non_arpeggiateDB.ID]
	if non_arpeggiateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), non_arpeggiateStaged, non_arpeggiateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
