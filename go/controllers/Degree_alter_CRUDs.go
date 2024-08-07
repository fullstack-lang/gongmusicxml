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
var __Degree_alter__dummysDeclaration__ models.Degree_alter
var __Degree_alter_time__dummyDeclaration time.Duration

var mutexDegree_alter sync.Mutex

// An Degree_alterID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDegree_alter updateDegree_alter deleteDegree_alter
type Degree_alterID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Degree_alterInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDegree_alter updateDegree_alter
type Degree_alterInput struct {
	// The Degree_alter to submit or modify
	// in: body
	Degree_alter *orm.Degree_alterAPI
}

// GetDegree_alters
//
// swagger:route GET /degree_alters degree_alters getDegree_alters
//
// # Get all degree_alters
//
// Responses:
// default: genericError
//
//	200: degree_alterDBResponse
func (controller *Controller) GetDegree_alters(c *gin.Context) {

	// source slice
	var degree_alterDBs []orm.Degree_alterDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDegree_alters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_alter.GetDB()

	query := db.Find(&degree_alterDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	degree_alterAPIs := make([]orm.Degree_alterAPI, 0)

	// for each degree_alter, update fields from the database nullable fields
	for idx := range degree_alterDBs {
		degree_alterDB := &degree_alterDBs[idx]
		_ = degree_alterDB
		var degree_alterAPI orm.Degree_alterAPI

		// insertion point for updating fields
		degree_alterAPI.ID = degree_alterDB.ID
		degree_alterDB.CopyBasicFieldsToDegree_alter_WOP(&degree_alterAPI.Degree_alter_WOP)
		degree_alterAPI.Degree_alterPointersEncoding = degree_alterDB.Degree_alterPointersEncoding
		degree_alterAPIs = append(degree_alterAPIs, degree_alterAPI)
	}

	c.JSON(http.StatusOK, degree_alterAPIs)
}

// PostDegree_alter
//
// swagger:route POST /degree_alters degree_alters postDegree_alter
//
// Creates a degree_alter
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDegree_alter(c *gin.Context) {

	mutexDegree_alter.Lock()
	defer mutexDegree_alter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDegree_alters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_alter.GetDB()

	// Validate input
	var input orm.Degree_alterAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create degree_alter
	degree_alterDB := orm.Degree_alterDB{}
	degree_alterDB.Degree_alterPointersEncoding = input.Degree_alterPointersEncoding
	degree_alterDB.CopyBasicFieldsFromDegree_alter_WOP(&input.Degree_alter_WOP)

	query := db.Create(&degree_alterDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDegree_alter.CheckoutPhaseOneInstance(&degree_alterDB)
	degree_alter := backRepo.BackRepoDegree_alter.Map_Degree_alterDBID_Degree_alterPtr[degree_alterDB.ID]

	if degree_alter != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), degree_alter)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, degree_alterDB)
}

// GetDegree_alter
//
// swagger:route GET /degree_alters/{ID} degree_alters getDegree_alter
//
// Gets the details for a degree_alter.
//
// Responses:
// default: genericError
//
//	200: degree_alterDBResponse
func (controller *Controller) GetDegree_alter(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDegree_alter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_alter.GetDB()

	// Get degree_alterDB in DB
	var degree_alterDB orm.Degree_alterDB
	if err := db.First(&degree_alterDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var degree_alterAPI orm.Degree_alterAPI
	degree_alterAPI.ID = degree_alterDB.ID
	degree_alterAPI.Degree_alterPointersEncoding = degree_alterDB.Degree_alterPointersEncoding
	degree_alterDB.CopyBasicFieldsToDegree_alter_WOP(&degree_alterAPI.Degree_alter_WOP)

	c.JSON(http.StatusOK, degree_alterAPI)
}

// UpdateDegree_alter
//
// swagger:route PATCH /degree_alters/{ID} degree_alters updateDegree_alter
//
// # Update a degree_alter
//
// Responses:
// default: genericError
//
//	200: degree_alterDBResponse
func (controller *Controller) UpdateDegree_alter(c *gin.Context) {

	mutexDegree_alter.Lock()
	defer mutexDegree_alter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDegree_alter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_alter.GetDB()

	// Validate input
	var input orm.Degree_alterAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var degree_alterDB orm.Degree_alterDB

	// fetch the degree_alter
	query := db.First(&degree_alterDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	degree_alterDB.CopyBasicFieldsFromDegree_alter_WOP(&input.Degree_alter_WOP)
	degree_alterDB.Degree_alterPointersEncoding = input.Degree_alterPointersEncoding

	query = db.Model(&degree_alterDB).Updates(degree_alterDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	degree_alterNew := new(models.Degree_alter)
	degree_alterDB.CopyBasicFieldsToDegree_alter(degree_alterNew)

	// redeem pointers
	degree_alterDB.DecodePointers(backRepo, degree_alterNew)

	// get stage instance from DB instance, and call callback function
	degree_alterOld := backRepo.BackRepoDegree_alter.Map_Degree_alterDBID_Degree_alterPtr[degree_alterDB.ID]
	if degree_alterOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), degree_alterOld, degree_alterNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the degree_alterDB
	c.JSON(http.StatusOK, degree_alterDB)
}

// DeleteDegree_alter
//
// swagger:route DELETE /degree_alters/{ID} degree_alters deleteDegree_alter
//
// # Delete a degree_alter
//
// default: genericError
//
//	200: degree_alterDBResponse
func (controller *Controller) DeleteDegree_alter(c *gin.Context) {

	mutexDegree_alter.Lock()
	defer mutexDegree_alter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDegree_alter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_alter.GetDB()

	// Get model if exist
	var degree_alterDB orm.Degree_alterDB
	if err := db.First(&degree_alterDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&degree_alterDB)

	// get an instance (not staged) from DB instance, and call callback function
	degree_alterDeleted := new(models.Degree_alter)
	degree_alterDB.CopyBasicFieldsToDegree_alter(degree_alterDeleted)

	// get stage instance from DB instance, and call callback function
	degree_alterStaged := backRepo.BackRepoDegree_alter.Map_Degree_alterDBID_Degree_alterPtr[degree_alterDB.ID]
	if degree_alterStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), degree_alterStaged, degree_alterDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
