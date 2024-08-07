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
var __Credit__dummysDeclaration__ models.Credit
var __Credit_time__dummyDeclaration time.Duration

var mutexCredit sync.Mutex

// An CreditID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCredit updateCredit deleteCredit
type CreditID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CreditInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCredit updateCredit
type CreditInput struct {
	// The Credit to submit or modify
	// in: body
	Credit *orm.CreditAPI
}

// GetCredits
//
// swagger:route GET /credits credits getCredits
//
// # Get all credits
//
// Responses:
// default: genericError
//
//	200: creditDBResponse
func (controller *Controller) GetCredits(c *gin.Context) {

	// source slice
	var creditDBs []orm.CreditDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCredits", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCredit.GetDB()

	query := db.Find(&creditDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	creditAPIs := make([]orm.CreditAPI, 0)

	// for each credit, update fields from the database nullable fields
	for idx := range creditDBs {
		creditDB := &creditDBs[idx]
		_ = creditDB
		var creditAPI orm.CreditAPI

		// insertion point for updating fields
		creditAPI.ID = creditDB.ID
		creditDB.CopyBasicFieldsToCredit_WOP(&creditAPI.Credit_WOP)
		creditAPI.CreditPointersEncoding = creditDB.CreditPointersEncoding
		creditAPIs = append(creditAPIs, creditAPI)
	}

	c.JSON(http.StatusOK, creditAPIs)
}

// PostCredit
//
// swagger:route POST /credits credits postCredit
//
// Creates a credit
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCredit(c *gin.Context) {

	mutexCredit.Lock()
	defer mutexCredit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCredits", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCredit.GetDB()

	// Validate input
	var input orm.CreditAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create credit
	creditDB := orm.CreditDB{}
	creditDB.CreditPointersEncoding = input.CreditPointersEncoding
	creditDB.CopyBasicFieldsFromCredit_WOP(&input.Credit_WOP)

	query := db.Create(&creditDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCredit.CheckoutPhaseOneInstance(&creditDB)
	credit := backRepo.BackRepoCredit.Map_CreditDBID_CreditPtr[creditDB.ID]

	if credit != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), credit)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, creditDB)
}

// GetCredit
//
// swagger:route GET /credits/{ID} credits getCredit
//
// Gets the details for a credit.
//
// Responses:
// default: genericError
//
//	200: creditDBResponse
func (controller *Controller) GetCredit(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCredit", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCredit.GetDB()

	// Get creditDB in DB
	var creditDB orm.CreditDB
	if err := db.First(&creditDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var creditAPI orm.CreditAPI
	creditAPI.ID = creditDB.ID
	creditAPI.CreditPointersEncoding = creditDB.CreditPointersEncoding
	creditDB.CopyBasicFieldsToCredit_WOP(&creditAPI.Credit_WOP)

	c.JSON(http.StatusOK, creditAPI)
}

// UpdateCredit
//
// swagger:route PATCH /credits/{ID} credits updateCredit
//
// # Update a credit
//
// Responses:
// default: genericError
//
//	200: creditDBResponse
func (controller *Controller) UpdateCredit(c *gin.Context) {

	mutexCredit.Lock()
	defer mutexCredit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCredit", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCredit.GetDB()

	// Validate input
	var input orm.CreditAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var creditDB orm.CreditDB

	// fetch the credit
	query := db.First(&creditDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	creditDB.CopyBasicFieldsFromCredit_WOP(&input.Credit_WOP)
	creditDB.CreditPointersEncoding = input.CreditPointersEncoding

	query = db.Model(&creditDB).Updates(creditDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	creditNew := new(models.Credit)
	creditDB.CopyBasicFieldsToCredit(creditNew)

	// redeem pointers
	creditDB.DecodePointers(backRepo, creditNew)

	// get stage instance from DB instance, and call callback function
	creditOld := backRepo.BackRepoCredit.Map_CreditDBID_CreditPtr[creditDB.ID]
	if creditOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), creditOld, creditNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the creditDB
	c.JSON(http.StatusOK, creditDB)
}

// DeleteCredit
//
// swagger:route DELETE /credits/{ID} credits deleteCredit
//
// # Delete a credit
//
// default: genericError
//
//	200: creditDBResponse
func (controller *Controller) DeleteCredit(c *gin.Context) {

	mutexCredit.Lock()
	defer mutexCredit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCredit", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCredit.GetDB()

	// Get model if exist
	var creditDB orm.CreditDB
	if err := db.First(&creditDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&creditDB)

	// get an instance (not staged) from DB instance, and call callback function
	creditDeleted := new(models.Credit)
	creditDB.CopyBasicFieldsToCredit(creditDeleted)

	// get stage instance from DB instance, and call callback function
	creditStaged := backRepo.BackRepoCredit.Map_CreditDBID_CreditPtr[creditDB.ID]
	if creditStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), creditStaged, creditDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
