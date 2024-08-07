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
var __Fermata__dummysDeclaration__ models.Fermata
var __Fermata_time__dummyDeclaration time.Duration

var mutexFermata sync.Mutex

// An FermataID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFermata updateFermata deleteFermata
type FermataID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FermataInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFermata updateFermata
type FermataInput struct {
	// The Fermata to submit or modify
	// in: body
	Fermata *orm.FermataAPI
}

// GetFermatas
//
// swagger:route GET /fermatas fermatas getFermatas
//
// # Get all fermatas
//
// Responses:
// default: genericError
//
//	200: fermataDBResponse
func (controller *Controller) GetFermatas(c *gin.Context) {

	// source slice
	var fermataDBs []orm.FermataDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFermatas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFermata.GetDB()

	query := db.Find(&fermataDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	fermataAPIs := make([]orm.FermataAPI, 0)

	// for each fermata, update fields from the database nullable fields
	for idx := range fermataDBs {
		fermataDB := &fermataDBs[idx]
		_ = fermataDB
		var fermataAPI orm.FermataAPI

		// insertion point for updating fields
		fermataAPI.ID = fermataDB.ID
		fermataDB.CopyBasicFieldsToFermata_WOP(&fermataAPI.Fermata_WOP)
		fermataAPI.FermataPointersEncoding = fermataDB.FermataPointersEncoding
		fermataAPIs = append(fermataAPIs, fermataAPI)
	}

	c.JSON(http.StatusOK, fermataAPIs)
}

// PostFermata
//
// swagger:route POST /fermatas fermatas postFermata
//
// Creates a fermata
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFermata(c *gin.Context) {

	mutexFermata.Lock()
	defer mutexFermata.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFermatas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFermata.GetDB()

	// Validate input
	var input orm.FermataAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create fermata
	fermataDB := orm.FermataDB{}
	fermataDB.FermataPointersEncoding = input.FermataPointersEncoding
	fermataDB.CopyBasicFieldsFromFermata_WOP(&input.Fermata_WOP)

	query := db.Create(&fermataDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFermata.CheckoutPhaseOneInstance(&fermataDB)
	fermata := backRepo.BackRepoFermata.Map_FermataDBID_FermataPtr[fermataDB.ID]

	if fermata != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), fermata)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, fermataDB)
}

// GetFermata
//
// swagger:route GET /fermatas/{ID} fermatas getFermata
//
// Gets the details for a fermata.
//
// Responses:
// default: genericError
//
//	200: fermataDBResponse
func (controller *Controller) GetFermata(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFermata", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFermata.GetDB()

	// Get fermataDB in DB
	var fermataDB orm.FermataDB
	if err := db.First(&fermataDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var fermataAPI orm.FermataAPI
	fermataAPI.ID = fermataDB.ID
	fermataAPI.FermataPointersEncoding = fermataDB.FermataPointersEncoding
	fermataDB.CopyBasicFieldsToFermata_WOP(&fermataAPI.Fermata_WOP)

	c.JSON(http.StatusOK, fermataAPI)
}

// UpdateFermata
//
// swagger:route PATCH /fermatas/{ID} fermatas updateFermata
//
// # Update a fermata
//
// Responses:
// default: genericError
//
//	200: fermataDBResponse
func (controller *Controller) UpdateFermata(c *gin.Context) {

	mutexFermata.Lock()
	defer mutexFermata.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFermata", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFermata.GetDB()

	// Validate input
	var input orm.FermataAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var fermataDB orm.FermataDB

	// fetch the fermata
	query := db.First(&fermataDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	fermataDB.CopyBasicFieldsFromFermata_WOP(&input.Fermata_WOP)
	fermataDB.FermataPointersEncoding = input.FermataPointersEncoding

	query = db.Model(&fermataDB).Updates(fermataDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	fermataNew := new(models.Fermata)
	fermataDB.CopyBasicFieldsToFermata(fermataNew)

	// redeem pointers
	fermataDB.DecodePointers(backRepo, fermataNew)

	// get stage instance from DB instance, and call callback function
	fermataOld := backRepo.BackRepoFermata.Map_FermataDBID_FermataPtr[fermataDB.ID]
	if fermataOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), fermataOld, fermataNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the fermataDB
	c.JSON(http.StatusOK, fermataDB)
}

// DeleteFermata
//
// swagger:route DELETE /fermatas/{ID} fermatas deleteFermata
//
// # Delete a fermata
//
// default: genericError
//
//	200: fermataDBResponse
func (controller *Controller) DeleteFermata(c *gin.Context) {

	mutexFermata.Lock()
	defer mutexFermata.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFermata", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFermata.GetDB()

	// Get model if exist
	var fermataDB orm.FermataDB
	if err := db.First(&fermataDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&fermataDB)

	// get an instance (not staged) from DB instance, and call callback function
	fermataDeleted := new(models.Fermata)
	fermataDB.CopyBasicFieldsToFermata(fermataDeleted)

	// get stage instance from DB instance, and call callback function
	fermataStaged := backRepo.BackRepoFermata.Map_FermataDBID_FermataPtr[fermataDB.ID]
	if fermataStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), fermataStaged, fermataDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
