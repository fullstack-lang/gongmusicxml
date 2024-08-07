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
var __Wedge__dummysDeclaration__ models.Wedge
var __Wedge_time__dummyDeclaration time.Duration

var mutexWedge sync.Mutex

// An WedgeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getWedge updateWedge deleteWedge
type WedgeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// WedgeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postWedge updateWedge
type WedgeInput struct {
	// The Wedge to submit or modify
	// in: body
	Wedge *orm.WedgeAPI
}

// GetWedges
//
// swagger:route GET /wedges wedges getWedges
//
// # Get all wedges
//
// Responses:
// default: genericError
//
//	200: wedgeDBResponse
func (controller *Controller) GetWedges(c *gin.Context) {

	// source slice
	var wedgeDBs []orm.WedgeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWedges", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWedge.GetDB()

	query := db.Find(&wedgeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	wedgeAPIs := make([]orm.WedgeAPI, 0)

	// for each wedge, update fields from the database nullable fields
	for idx := range wedgeDBs {
		wedgeDB := &wedgeDBs[idx]
		_ = wedgeDB
		var wedgeAPI orm.WedgeAPI

		// insertion point for updating fields
		wedgeAPI.ID = wedgeDB.ID
		wedgeDB.CopyBasicFieldsToWedge_WOP(&wedgeAPI.Wedge_WOP)
		wedgeAPI.WedgePointersEncoding = wedgeDB.WedgePointersEncoding
		wedgeAPIs = append(wedgeAPIs, wedgeAPI)
	}

	c.JSON(http.StatusOK, wedgeAPIs)
}

// PostWedge
//
// swagger:route POST /wedges wedges postWedge
//
// Creates a wedge
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostWedge(c *gin.Context) {

	mutexWedge.Lock()
	defer mutexWedge.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostWedges", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWedge.GetDB()

	// Validate input
	var input orm.WedgeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create wedge
	wedgeDB := orm.WedgeDB{}
	wedgeDB.WedgePointersEncoding = input.WedgePointersEncoding
	wedgeDB.CopyBasicFieldsFromWedge_WOP(&input.Wedge_WOP)

	query := db.Create(&wedgeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoWedge.CheckoutPhaseOneInstance(&wedgeDB)
	wedge := backRepo.BackRepoWedge.Map_WedgeDBID_WedgePtr[wedgeDB.ID]

	if wedge != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), wedge)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, wedgeDB)
}

// GetWedge
//
// swagger:route GET /wedges/{ID} wedges getWedge
//
// Gets the details for a wedge.
//
// Responses:
// default: genericError
//
//	200: wedgeDBResponse
func (controller *Controller) GetWedge(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWedge", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWedge.GetDB()

	// Get wedgeDB in DB
	var wedgeDB orm.WedgeDB
	if err := db.First(&wedgeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var wedgeAPI orm.WedgeAPI
	wedgeAPI.ID = wedgeDB.ID
	wedgeAPI.WedgePointersEncoding = wedgeDB.WedgePointersEncoding
	wedgeDB.CopyBasicFieldsToWedge_WOP(&wedgeAPI.Wedge_WOP)

	c.JSON(http.StatusOK, wedgeAPI)
}

// UpdateWedge
//
// swagger:route PATCH /wedges/{ID} wedges updateWedge
//
// # Update a wedge
//
// Responses:
// default: genericError
//
//	200: wedgeDBResponse
func (controller *Controller) UpdateWedge(c *gin.Context) {

	mutexWedge.Lock()
	defer mutexWedge.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateWedge", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWedge.GetDB()

	// Validate input
	var input orm.WedgeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var wedgeDB orm.WedgeDB

	// fetch the wedge
	query := db.First(&wedgeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	wedgeDB.CopyBasicFieldsFromWedge_WOP(&input.Wedge_WOP)
	wedgeDB.WedgePointersEncoding = input.WedgePointersEncoding

	query = db.Model(&wedgeDB).Updates(wedgeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	wedgeNew := new(models.Wedge)
	wedgeDB.CopyBasicFieldsToWedge(wedgeNew)

	// redeem pointers
	wedgeDB.DecodePointers(backRepo, wedgeNew)

	// get stage instance from DB instance, and call callback function
	wedgeOld := backRepo.BackRepoWedge.Map_WedgeDBID_WedgePtr[wedgeDB.ID]
	if wedgeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), wedgeOld, wedgeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the wedgeDB
	c.JSON(http.StatusOK, wedgeDB)
}

// DeleteWedge
//
// swagger:route DELETE /wedges/{ID} wedges deleteWedge
//
// # Delete a wedge
//
// default: genericError
//
//	200: wedgeDBResponse
func (controller *Controller) DeleteWedge(c *gin.Context) {

	mutexWedge.Lock()
	defer mutexWedge.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteWedge", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWedge.GetDB()

	// Get model if exist
	var wedgeDB orm.WedgeDB
	if err := db.First(&wedgeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&wedgeDB)

	// get an instance (not staged) from DB instance, and call callback function
	wedgeDeleted := new(models.Wedge)
	wedgeDB.CopyBasicFieldsToWedge(wedgeDeleted)

	// get stage instance from DB instance, and call callback function
	wedgeStaged := backRepo.BackRepoWedge.Map_WedgeDBID_WedgePtr[wedgeDB.ID]
	if wedgeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), wedgeStaged, wedgeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
