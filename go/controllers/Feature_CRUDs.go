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
var __Feature__dummysDeclaration__ models.Feature
var __Feature_time__dummyDeclaration time.Duration

var mutexFeature sync.Mutex

// An FeatureID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFeature updateFeature deleteFeature
type FeatureID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FeatureInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFeature updateFeature
type FeatureInput struct {
	// The Feature to submit or modify
	// in: body
	Feature *orm.FeatureAPI
}

// GetFeatures
//
// swagger:route GET /features features getFeatures
//
// # Get all features
//
// Responses:
// default: genericError
//
//	200: featureDBResponse
func (controller *Controller) GetFeatures(c *gin.Context) {

	// source slice
	var featureDBs []orm.FeatureDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFeatures", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFeature.GetDB()

	query := db.Find(&featureDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	featureAPIs := make([]orm.FeatureAPI, 0)

	// for each feature, update fields from the database nullable fields
	for idx := range featureDBs {
		featureDB := &featureDBs[idx]
		_ = featureDB
		var featureAPI orm.FeatureAPI

		// insertion point for updating fields
		featureAPI.ID = featureDB.ID
		featureDB.CopyBasicFieldsToFeature_WOP(&featureAPI.Feature_WOP)
		featureAPI.FeaturePointersEncoding = featureDB.FeaturePointersEncoding
		featureAPIs = append(featureAPIs, featureAPI)
	}

	c.JSON(http.StatusOK, featureAPIs)
}

// PostFeature
//
// swagger:route POST /features features postFeature
//
// Creates a feature
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFeature(c *gin.Context) {

	mutexFeature.Lock()
	defer mutexFeature.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFeatures", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFeature.GetDB()

	// Validate input
	var input orm.FeatureAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create feature
	featureDB := orm.FeatureDB{}
	featureDB.FeaturePointersEncoding = input.FeaturePointersEncoding
	featureDB.CopyBasicFieldsFromFeature_WOP(&input.Feature_WOP)

	query := db.Create(&featureDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFeature.CheckoutPhaseOneInstance(&featureDB)
	feature := backRepo.BackRepoFeature.Map_FeatureDBID_FeaturePtr[featureDB.ID]

	if feature != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), feature)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, featureDB)
}

// GetFeature
//
// swagger:route GET /features/{ID} features getFeature
//
// Gets the details for a feature.
//
// Responses:
// default: genericError
//
//	200: featureDBResponse
func (controller *Controller) GetFeature(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFeature", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFeature.GetDB()

	// Get featureDB in DB
	var featureDB orm.FeatureDB
	if err := db.First(&featureDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var featureAPI orm.FeatureAPI
	featureAPI.ID = featureDB.ID
	featureAPI.FeaturePointersEncoding = featureDB.FeaturePointersEncoding
	featureDB.CopyBasicFieldsToFeature_WOP(&featureAPI.Feature_WOP)

	c.JSON(http.StatusOK, featureAPI)
}

// UpdateFeature
//
// swagger:route PATCH /features/{ID} features updateFeature
//
// # Update a feature
//
// Responses:
// default: genericError
//
//	200: featureDBResponse
func (controller *Controller) UpdateFeature(c *gin.Context) {

	mutexFeature.Lock()
	defer mutexFeature.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFeature", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFeature.GetDB()

	// Validate input
	var input orm.FeatureAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var featureDB orm.FeatureDB

	// fetch the feature
	query := db.First(&featureDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	featureDB.CopyBasicFieldsFromFeature_WOP(&input.Feature_WOP)
	featureDB.FeaturePointersEncoding = input.FeaturePointersEncoding

	query = db.Model(&featureDB).Updates(featureDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	featureNew := new(models.Feature)
	featureDB.CopyBasicFieldsToFeature(featureNew)

	// redeem pointers
	featureDB.DecodePointers(backRepo, featureNew)

	// get stage instance from DB instance, and call callback function
	featureOld := backRepo.BackRepoFeature.Map_FeatureDBID_FeaturePtr[featureDB.ID]
	if featureOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), featureOld, featureNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the featureDB
	c.JSON(http.StatusOK, featureDB)
}

// DeleteFeature
//
// swagger:route DELETE /features/{ID} features deleteFeature
//
// # Delete a feature
//
// default: genericError
//
//	200: featureDBResponse
func (controller *Controller) DeleteFeature(c *gin.Context) {

	mutexFeature.Lock()
	defer mutexFeature.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFeature", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFeature.GetDB()

	// Get model if exist
	var featureDB orm.FeatureDB
	if err := db.First(&featureDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&featureDB)

	// get an instance (not staged) from DB instance, and call callback function
	featureDeleted := new(models.Feature)
	featureDB.CopyBasicFieldsToFeature(featureDeleted)

	// get stage instance from DB instance, and call callback function
	featureStaged := backRepo.BackRepoFeature.Map_FeatureDBID_FeaturePtr[featureDB.ID]
	if featureStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), featureStaged, featureDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
