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
var __Segno__dummysDeclaration__ models.Segno
var __Segno_time__dummyDeclaration time.Duration

var mutexSegno sync.Mutex

// An SegnoID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSegno updateSegno deleteSegno
type SegnoID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SegnoInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSegno updateSegno
type SegnoInput struct {
	// The Segno to submit or modify
	// in: body
	Segno *orm.SegnoAPI
}

// GetSegnos
//
// swagger:route GET /segnos segnos getSegnos
//
// # Get all segnos
//
// Responses:
// default: genericError
//
//	200: segnoDBResponse
func (controller *Controller) GetSegnos(c *gin.Context) {

	// source slice
	var segnoDBs []orm.SegnoDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSegnos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSegno.GetDB()

	query := db.Find(&segnoDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	segnoAPIs := make([]orm.SegnoAPI, 0)

	// for each segno, update fields from the database nullable fields
	for idx := range segnoDBs {
		segnoDB := &segnoDBs[idx]
		_ = segnoDB
		var segnoAPI orm.SegnoAPI

		// insertion point for updating fields
		segnoAPI.ID = segnoDB.ID
		segnoDB.CopyBasicFieldsToSegno_WOP(&segnoAPI.Segno_WOP)
		segnoAPI.SegnoPointersEncoding = segnoDB.SegnoPointersEncoding
		segnoAPIs = append(segnoAPIs, segnoAPI)
	}

	c.JSON(http.StatusOK, segnoAPIs)
}

// PostSegno
//
// swagger:route POST /segnos segnos postSegno
//
// Creates a segno
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSegno(c *gin.Context) {

	mutexSegno.Lock()
	defer mutexSegno.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSegnos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSegno.GetDB()

	// Validate input
	var input orm.SegnoAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create segno
	segnoDB := orm.SegnoDB{}
	segnoDB.SegnoPointersEncoding = input.SegnoPointersEncoding
	segnoDB.CopyBasicFieldsFromSegno_WOP(&input.Segno_WOP)

	query := db.Create(&segnoDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSegno.CheckoutPhaseOneInstance(&segnoDB)
	segno := backRepo.BackRepoSegno.Map_SegnoDBID_SegnoPtr[segnoDB.ID]

	if segno != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), segno)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, segnoDB)
}

// GetSegno
//
// swagger:route GET /segnos/{ID} segnos getSegno
//
// Gets the details for a segno.
//
// Responses:
// default: genericError
//
//	200: segnoDBResponse
func (controller *Controller) GetSegno(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSegno", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSegno.GetDB()

	// Get segnoDB in DB
	var segnoDB orm.SegnoDB
	if err := db.First(&segnoDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var segnoAPI orm.SegnoAPI
	segnoAPI.ID = segnoDB.ID
	segnoAPI.SegnoPointersEncoding = segnoDB.SegnoPointersEncoding
	segnoDB.CopyBasicFieldsToSegno_WOP(&segnoAPI.Segno_WOP)

	c.JSON(http.StatusOK, segnoAPI)
}

// UpdateSegno
//
// swagger:route PATCH /segnos/{ID} segnos updateSegno
//
// # Update a segno
//
// Responses:
// default: genericError
//
//	200: segnoDBResponse
func (controller *Controller) UpdateSegno(c *gin.Context) {

	mutexSegno.Lock()
	defer mutexSegno.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSegno", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSegno.GetDB()

	// Validate input
	var input orm.SegnoAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var segnoDB orm.SegnoDB

	// fetch the segno
	query := db.First(&segnoDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	segnoDB.CopyBasicFieldsFromSegno_WOP(&input.Segno_WOP)
	segnoDB.SegnoPointersEncoding = input.SegnoPointersEncoding

	query = db.Model(&segnoDB).Updates(segnoDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	segnoNew := new(models.Segno)
	segnoDB.CopyBasicFieldsToSegno(segnoNew)

	// redeem pointers
	segnoDB.DecodePointers(backRepo, segnoNew)

	// get stage instance from DB instance, and call callback function
	segnoOld := backRepo.BackRepoSegno.Map_SegnoDBID_SegnoPtr[segnoDB.ID]
	if segnoOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), segnoOld, segnoNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the segnoDB
	c.JSON(http.StatusOK, segnoDB)
}

// DeleteSegno
//
// swagger:route DELETE /segnos/{ID} segnos deleteSegno
//
// # Delete a segno
//
// default: genericError
//
//	200: segnoDBResponse
func (controller *Controller) DeleteSegno(c *gin.Context) {

	mutexSegno.Lock()
	defer mutexSegno.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSegno", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSegno.GetDB()

	// Get model if exist
	var segnoDB orm.SegnoDB
	if err := db.First(&segnoDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&segnoDB)

	// get an instance (not staged) from DB instance, and call callback function
	segnoDeleted := new(models.Segno)
	segnoDB.CopyBasicFieldsToSegno(segnoDeleted)

	// get stage instance from DB instance, and call callback function
	segnoStaged := backRepo.BackRepoSegno.Map_SegnoDBID_SegnoPtr[segnoDB.ID]
	if segnoStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), segnoStaged, segnoDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
