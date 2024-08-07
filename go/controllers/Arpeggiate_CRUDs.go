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
var __Arpeggiate__dummysDeclaration__ models.Arpeggiate
var __Arpeggiate_time__dummyDeclaration time.Duration

var mutexArpeggiate sync.Mutex

// An ArpeggiateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getArpeggiate updateArpeggiate deleteArpeggiate
type ArpeggiateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ArpeggiateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postArpeggiate updateArpeggiate
type ArpeggiateInput struct {
	// The Arpeggiate to submit or modify
	// in: body
	Arpeggiate *orm.ArpeggiateAPI
}

// GetArpeggiates
//
// swagger:route GET /arpeggiates arpeggiates getArpeggiates
//
// # Get all arpeggiates
//
// Responses:
// default: genericError
//
//	200: arpeggiateDBResponse
func (controller *Controller) GetArpeggiates(c *gin.Context) {

	// source slice
	var arpeggiateDBs []orm.ArpeggiateDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetArpeggiates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArpeggiate.GetDB()

	query := db.Find(&arpeggiateDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	arpeggiateAPIs := make([]orm.ArpeggiateAPI, 0)

	// for each arpeggiate, update fields from the database nullable fields
	for idx := range arpeggiateDBs {
		arpeggiateDB := &arpeggiateDBs[idx]
		_ = arpeggiateDB
		var arpeggiateAPI orm.ArpeggiateAPI

		// insertion point for updating fields
		arpeggiateAPI.ID = arpeggiateDB.ID
		arpeggiateDB.CopyBasicFieldsToArpeggiate_WOP(&arpeggiateAPI.Arpeggiate_WOP)
		arpeggiateAPI.ArpeggiatePointersEncoding = arpeggiateDB.ArpeggiatePointersEncoding
		arpeggiateAPIs = append(arpeggiateAPIs, arpeggiateAPI)
	}

	c.JSON(http.StatusOK, arpeggiateAPIs)
}

// PostArpeggiate
//
// swagger:route POST /arpeggiates arpeggiates postArpeggiate
//
// Creates a arpeggiate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostArpeggiate(c *gin.Context) {

	mutexArpeggiate.Lock()
	defer mutexArpeggiate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostArpeggiates", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArpeggiate.GetDB()

	// Validate input
	var input orm.ArpeggiateAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create arpeggiate
	arpeggiateDB := orm.ArpeggiateDB{}
	arpeggiateDB.ArpeggiatePointersEncoding = input.ArpeggiatePointersEncoding
	arpeggiateDB.CopyBasicFieldsFromArpeggiate_WOP(&input.Arpeggiate_WOP)

	query := db.Create(&arpeggiateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoArpeggiate.CheckoutPhaseOneInstance(&arpeggiateDB)
	arpeggiate := backRepo.BackRepoArpeggiate.Map_ArpeggiateDBID_ArpeggiatePtr[arpeggiateDB.ID]

	if arpeggiate != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), arpeggiate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, arpeggiateDB)
}

// GetArpeggiate
//
// swagger:route GET /arpeggiates/{ID} arpeggiates getArpeggiate
//
// Gets the details for a arpeggiate.
//
// Responses:
// default: genericError
//
//	200: arpeggiateDBResponse
func (controller *Controller) GetArpeggiate(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetArpeggiate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArpeggiate.GetDB()

	// Get arpeggiateDB in DB
	var arpeggiateDB orm.ArpeggiateDB
	if err := db.First(&arpeggiateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var arpeggiateAPI orm.ArpeggiateAPI
	arpeggiateAPI.ID = arpeggiateDB.ID
	arpeggiateAPI.ArpeggiatePointersEncoding = arpeggiateDB.ArpeggiatePointersEncoding
	arpeggiateDB.CopyBasicFieldsToArpeggiate_WOP(&arpeggiateAPI.Arpeggiate_WOP)

	c.JSON(http.StatusOK, arpeggiateAPI)
}

// UpdateArpeggiate
//
// swagger:route PATCH /arpeggiates/{ID} arpeggiates updateArpeggiate
//
// # Update a arpeggiate
//
// Responses:
// default: genericError
//
//	200: arpeggiateDBResponse
func (controller *Controller) UpdateArpeggiate(c *gin.Context) {

	mutexArpeggiate.Lock()
	defer mutexArpeggiate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateArpeggiate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArpeggiate.GetDB()

	// Validate input
	var input orm.ArpeggiateAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var arpeggiateDB orm.ArpeggiateDB

	// fetch the arpeggiate
	query := db.First(&arpeggiateDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	arpeggiateDB.CopyBasicFieldsFromArpeggiate_WOP(&input.Arpeggiate_WOP)
	arpeggiateDB.ArpeggiatePointersEncoding = input.ArpeggiatePointersEncoding

	query = db.Model(&arpeggiateDB).Updates(arpeggiateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	arpeggiateNew := new(models.Arpeggiate)
	arpeggiateDB.CopyBasicFieldsToArpeggiate(arpeggiateNew)

	// redeem pointers
	arpeggiateDB.DecodePointers(backRepo, arpeggiateNew)

	// get stage instance from DB instance, and call callback function
	arpeggiateOld := backRepo.BackRepoArpeggiate.Map_ArpeggiateDBID_ArpeggiatePtr[arpeggiateDB.ID]
	if arpeggiateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), arpeggiateOld, arpeggiateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the arpeggiateDB
	c.JSON(http.StatusOK, arpeggiateDB)
}

// DeleteArpeggiate
//
// swagger:route DELETE /arpeggiates/{ID} arpeggiates deleteArpeggiate
//
// # Delete a arpeggiate
//
// default: genericError
//
//	200: arpeggiateDBResponse
func (controller *Controller) DeleteArpeggiate(c *gin.Context) {

	mutexArpeggiate.Lock()
	defer mutexArpeggiate.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteArpeggiate", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArpeggiate.GetDB()

	// Get model if exist
	var arpeggiateDB orm.ArpeggiateDB
	if err := db.First(&arpeggiateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&arpeggiateDB)

	// get an instance (not staged) from DB instance, and call callback function
	arpeggiateDeleted := new(models.Arpeggiate)
	arpeggiateDB.CopyBasicFieldsToArpeggiate(arpeggiateDeleted)

	// get stage instance from DB instance, and call callback function
	arpeggiateStaged := backRepo.BackRepoArpeggiate.Map_ArpeggiateDBID_ArpeggiatePtr[arpeggiateDB.ID]
	if arpeggiateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), arpeggiateStaged, arpeggiateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
