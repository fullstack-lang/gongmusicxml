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
var __Strong_accent__dummysDeclaration__ models.Strong_accent
var __Strong_accent_time__dummyDeclaration time.Duration

var mutexStrong_accent sync.Mutex

// An Strong_accentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStrong_accent updateStrong_accent deleteStrong_accent
type Strong_accentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Strong_accentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStrong_accent updateStrong_accent
type Strong_accentInput struct {
	// The Strong_accent to submit or modify
	// in: body
	Strong_accent *orm.Strong_accentAPI
}

// GetStrong_accents
//
// swagger:route GET /strong_accents strong_accents getStrong_accents
//
// # Get all strong_accents
//
// Responses:
// default: genericError
//
//	200: strong_accentDBResponse
func (controller *Controller) GetStrong_accents(c *gin.Context) {

	// source slice
	var strong_accentDBs []orm.Strong_accentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStrong_accents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStrong_accent.GetDB()

	query := db.Find(&strong_accentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	strong_accentAPIs := make([]orm.Strong_accentAPI, 0)

	// for each strong_accent, update fields from the database nullable fields
	for idx := range strong_accentDBs {
		strong_accentDB := &strong_accentDBs[idx]
		_ = strong_accentDB
		var strong_accentAPI orm.Strong_accentAPI

		// insertion point for updating fields
		strong_accentAPI.ID = strong_accentDB.ID
		strong_accentDB.CopyBasicFieldsToStrong_accent_WOP(&strong_accentAPI.Strong_accent_WOP)
		strong_accentAPI.Strong_accentPointersEncoding = strong_accentDB.Strong_accentPointersEncoding
		strong_accentAPIs = append(strong_accentAPIs, strong_accentAPI)
	}

	c.JSON(http.StatusOK, strong_accentAPIs)
}

// PostStrong_accent
//
// swagger:route POST /strong_accents strong_accents postStrong_accent
//
// Creates a strong_accent
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStrong_accent(c *gin.Context) {

	mutexStrong_accent.Lock()
	defer mutexStrong_accent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStrong_accents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStrong_accent.GetDB()

	// Validate input
	var input orm.Strong_accentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create strong_accent
	strong_accentDB := orm.Strong_accentDB{}
	strong_accentDB.Strong_accentPointersEncoding = input.Strong_accentPointersEncoding
	strong_accentDB.CopyBasicFieldsFromStrong_accent_WOP(&input.Strong_accent_WOP)

	query := db.Create(&strong_accentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStrong_accent.CheckoutPhaseOneInstance(&strong_accentDB)
	strong_accent := backRepo.BackRepoStrong_accent.Map_Strong_accentDBID_Strong_accentPtr[strong_accentDB.ID]

	if strong_accent != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), strong_accent)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, strong_accentDB)
}

// GetStrong_accent
//
// swagger:route GET /strong_accents/{ID} strong_accents getStrong_accent
//
// Gets the details for a strong_accent.
//
// Responses:
// default: genericError
//
//	200: strong_accentDBResponse
func (controller *Controller) GetStrong_accent(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStrong_accent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStrong_accent.GetDB()

	// Get strong_accentDB in DB
	var strong_accentDB orm.Strong_accentDB
	if err := db.First(&strong_accentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var strong_accentAPI orm.Strong_accentAPI
	strong_accentAPI.ID = strong_accentDB.ID
	strong_accentAPI.Strong_accentPointersEncoding = strong_accentDB.Strong_accentPointersEncoding
	strong_accentDB.CopyBasicFieldsToStrong_accent_WOP(&strong_accentAPI.Strong_accent_WOP)

	c.JSON(http.StatusOK, strong_accentAPI)
}

// UpdateStrong_accent
//
// swagger:route PATCH /strong_accents/{ID} strong_accents updateStrong_accent
//
// # Update a strong_accent
//
// Responses:
// default: genericError
//
//	200: strong_accentDBResponse
func (controller *Controller) UpdateStrong_accent(c *gin.Context) {

	mutexStrong_accent.Lock()
	defer mutexStrong_accent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStrong_accent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStrong_accent.GetDB()

	// Validate input
	var input orm.Strong_accentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var strong_accentDB orm.Strong_accentDB

	// fetch the strong_accent
	query := db.First(&strong_accentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	strong_accentDB.CopyBasicFieldsFromStrong_accent_WOP(&input.Strong_accent_WOP)
	strong_accentDB.Strong_accentPointersEncoding = input.Strong_accentPointersEncoding

	query = db.Model(&strong_accentDB).Updates(strong_accentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	strong_accentNew := new(models.Strong_accent)
	strong_accentDB.CopyBasicFieldsToStrong_accent(strong_accentNew)

	// redeem pointers
	strong_accentDB.DecodePointers(backRepo, strong_accentNew)

	// get stage instance from DB instance, and call callback function
	strong_accentOld := backRepo.BackRepoStrong_accent.Map_Strong_accentDBID_Strong_accentPtr[strong_accentDB.ID]
	if strong_accentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), strong_accentOld, strong_accentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the strong_accentDB
	c.JSON(http.StatusOK, strong_accentDB)
}

// DeleteStrong_accent
//
// swagger:route DELETE /strong_accents/{ID} strong_accents deleteStrong_accent
//
// # Delete a strong_accent
//
// default: genericError
//
//	200: strong_accentDBResponse
func (controller *Controller) DeleteStrong_accent(c *gin.Context) {

	mutexStrong_accent.Lock()
	defer mutexStrong_accent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStrong_accent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStrong_accent.GetDB()

	// Get model if exist
	var strong_accentDB orm.Strong_accentDB
	if err := db.First(&strong_accentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&strong_accentDB)

	// get an instance (not staged) from DB instance, and call callback function
	strong_accentDeleted := new(models.Strong_accent)
	strong_accentDB.CopyBasicFieldsToStrong_accent(strong_accentDeleted)

	// get stage instance from DB instance, and call callback function
	strong_accentStaged := backRepo.BackRepoStrong_accent.Map_Strong_accentDBID_Strong_accentPtr[strong_accentDB.ID]
	if strong_accentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), strong_accentStaged, strong_accentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
