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
var __Kind__dummysDeclaration__ models.Kind
var __Kind_time__dummyDeclaration time.Duration

var mutexKind sync.Mutex

// An KindID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getKind updateKind deleteKind
type KindID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// KindInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postKind updateKind
type KindInput struct {
	// The Kind to submit or modify
	// in: body
	Kind *orm.KindAPI
}

// GetKinds
//
// swagger:route GET /kinds kinds getKinds
//
// # Get all kinds
//
// Responses:
// default: genericError
//
//	200: kindDBResponse
func (controller *Controller) GetKinds(c *gin.Context) {

	// source slice
	var kindDBs []orm.KindDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKinds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKind.GetDB()

	query := db.Find(&kindDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	kindAPIs := make([]orm.KindAPI, 0)

	// for each kind, update fields from the database nullable fields
	for idx := range kindDBs {
		kindDB := &kindDBs[idx]
		_ = kindDB
		var kindAPI orm.KindAPI

		// insertion point for updating fields
		kindAPI.ID = kindDB.ID
		kindDB.CopyBasicFieldsToKind_WOP(&kindAPI.Kind_WOP)
		kindAPI.KindPointersEncoding = kindDB.KindPointersEncoding
		kindAPIs = append(kindAPIs, kindAPI)
	}

	c.JSON(http.StatusOK, kindAPIs)
}

// PostKind
//
// swagger:route POST /kinds kinds postKind
//
// Creates a kind
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostKind(c *gin.Context) {

	mutexKind.Lock()
	defer mutexKind.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostKinds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKind.GetDB()

	// Validate input
	var input orm.KindAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create kind
	kindDB := orm.KindDB{}
	kindDB.KindPointersEncoding = input.KindPointersEncoding
	kindDB.CopyBasicFieldsFromKind_WOP(&input.Kind_WOP)

	query := db.Create(&kindDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoKind.CheckoutPhaseOneInstance(&kindDB)
	kind := backRepo.BackRepoKind.Map_KindDBID_KindPtr[kindDB.ID]

	if kind != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), kind)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, kindDB)
}

// GetKind
//
// swagger:route GET /kinds/{ID} kinds getKind
//
// Gets the details for a kind.
//
// Responses:
// default: genericError
//
//	200: kindDBResponse
func (controller *Controller) GetKind(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKind", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKind.GetDB()

	// Get kindDB in DB
	var kindDB orm.KindDB
	if err := db.First(&kindDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var kindAPI orm.KindAPI
	kindAPI.ID = kindDB.ID
	kindAPI.KindPointersEncoding = kindDB.KindPointersEncoding
	kindDB.CopyBasicFieldsToKind_WOP(&kindAPI.Kind_WOP)

	c.JSON(http.StatusOK, kindAPI)
}

// UpdateKind
//
// swagger:route PATCH /kinds/{ID} kinds updateKind
//
// # Update a kind
//
// Responses:
// default: genericError
//
//	200: kindDBResponse
func (controller *Controller) UpdateKind(c *gin.Context) {

	mutexKind.Lock()
	defer mutexKind.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateKind", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKind.GetDB()

	// Validate input
	var input orm.KindAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var kindDB orm.KindDB

	// fetch the kind
	query := db.First(&kindDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	kindDB.CopyBasicFieldsFromKind_WOP(&input.Kind_WOP)
	kindDB.KindPointersEncoding = input.KindPointersEncoding

	query = db.Model(&kindDB).Updates(kindDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	kindNew := new(models.Kind)
	kindDB.CopyBasicFieldsToKind(kindNew)

	// redeem pointers
	kindDB.DecodePointers(backRepo, kindNew)

	// get stage instance from DB instance, and call callback function
	kindOld := backRepo.BackRepoKind.Map_KindDBID_KindPtr[kindDB.ID]
	if kindOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), kindOld, kindNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the kindDB
	c.JSON(http.StatusOK, kindDB)
}

// DeleteKind
//
// swagger:route DELETE /kinds/{ID} kinds deleteKind
//
// # Delete a kind
//
// default: genericError
//
//	200: kindDBResponse
func (controller *Controller) DeleteKind(c *gin.Context) {

	mutexKind.Lock()
	defer mutexKind.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteKind", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKind.GetDB()

	// Get model if exist
	var kindDB orm.KindDB
	if err := db.First(&kindDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&kindDB)

	// get an instance (not staged) from DB instance, and call callback function
	kindDeleted := new(models.Kind)
	kindDB.CopyBasicFieldsToKind(kindDeleted)

	// get stage instance from DB instance, and call callback function
	kindStaged := backRepo.BackRepoKind.Map_KindDBID_KindPtr[kindDB.ID]
	if kindStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), kindStaged, kindDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
