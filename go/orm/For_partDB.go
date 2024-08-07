// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongmusicxml/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_For_part_sql sql.NullBool
var dummy_For_part_time time.Duration
var dummy_For_part_sort sort.Float64Slice

// For_partAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model for_partAPI
type For_partAPI struct {
	gorm.Model

	models.For_part_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	For_partPointersEncoding For_partPointersEncoding
}

// For_partPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type For_partPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Part_clef is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	Part_clefID sql.NullInt64

	// field Part_transpose is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	Part_transposeID sql.NullInt64
}

// For_partDB describes a for_part in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model for_partDB
type For_partDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field for_partDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	For_partPointersEncoding
}

// For_partDBs arrays for_partDBs
// swagger:response for_partDBsResponse
type For_partDBs []For_partDB

// For_partDBResponse provides response
// swagger:response for_partDBResponse
type For_partDBResponse struct {
	For_partDB
}

// For_partWOP is a For_part without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type For_partWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var For_part_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoFor_partStruct struct {
	// stores For_partDB according to their gorm ID
	Map_For_partDBID_For_partDB map[uint]*For_partDB

	// stores For_partDB ID according to For_part address
	Map_For_partPtr_For_partDBID map[*models.For_part]uint

	// stores For_part according to their gorm ID
	Map_For_partDBID_For_partPtr map[uint]*models.For_part

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFor_part *BackRepoFor_partStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFor_part.stage
	return
}

func (backRepoFor_part *BackRepoFor_partStruct) GetDB() *gorm.DB {
	return backRepoFor_part.db
}

// GetFor_partDBFromFor_partPtr is a handy function to access the back repo instance from the stage instance
func (backRepoFor_part *BackRepoFor_partStruct) GetFor_partDBFromFor_partPtr(for_part *models.For_part) (for_partDB *For_partDB) {
	id := backRepoFor_part.Map_For_partPtr_For_partDBID[for_part]
	for_partDB = backRepoFor_part.Map_For_partDBID_For_partDB[id]
	return
}

// BackRepoFor_part.CommitPhaseOne commits all staged instances of For_part to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFor_part *BackRepoFor_partStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for for_part := range stage.For_parts {
		backRepoFor_part.CommitPhaseOneInstance(for_part)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, for_part := range backRepoFor_part.Map_For_partDBID_For_partPtr {
		if _, ok := stage.For_parts[for_part]; !ok {
			backRepoFor_part.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFor_part.CommitDeleteInstance commits deletion of For_part to the BackRepo
func (backRepoFor_part *BackRepoFor_partStruct) CommitDeleteInstance(id uint) (Error error) {

	for_part := backRepoFor_part.Map_For_partDBID_For_partPtr[id]

	// for_part is not staged anymore, remove for_partDB
	for_partDB := backRepoFor_part.Map_For_partDBID_For_partDB[id]
	query := backRepoFor_part.db.Unscoped().Delete(&for_partDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoFor_part.Map_For_partPtr_For_partDBID, for_part)
	delete(backRepoFor_part.Map_For_partDBID_For_partPtr, id)
	delete(backRepoFor_part.Map_For_partDBID_For_partDB, id)

	return
}

// BackRepoFor_part.CommitPhaseOneInstance commits for_part staged instances of For_part to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFor_part *BackRepoFor_partStruct) CommitPhaseOneInstance(for_part *models.For_part) (Error error) {

	// check if the for_part is not commited yet
	if _, ok := backRepoFor_part.Map_For_partPtr_For_partDBID[for_part]; ok {
		return
	}

	// initiate for_part
	var for_partDB For_partDB
	for_partDB.CopyBasicFieldsFromFor_part(for_part)

	query := backRepoFor_part.db.Create(&for_partDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoFor_part.Map_For_partPtr_For_partDBID[for_part] = for_partDB.ID
	backRepoFor_part.Map_For_partDBID_For_partPtr[for_partDB.ID] = for_part
	backRepoFor_part.Map_For_partDBID_For_partDB[for_partDB.ID] = &for_partDB

	return
}

// BackRepoFor_part.CommitPhaseTwo commits all staged instances of For_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFor_part *BackRepoFor_partStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, for_part := range backRepoFor_part.Map_For_partDBID_For_partPtr {
		backRepoFor_part.CommitPhaseTwoInstance(backRepo, idx, for_part)
	}

	return
}

// BackRepoFor_part.CommitPhaseTwoInstance commits {{structname }} of models.For_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFor_part *BackRepoFor_partStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, for_part *models.For_part) (Error error) {

	// fetch matching for_partDB
	if for_partDB, ok := backRepoFor_part.Map_For_partDBID_For_partDB[idx]; ok {

		for_partDB.CopyBasicFieldsFromFor_part(for_part)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value for_part.Part_clef translates to updating the for_part.Part_clefID
		for_partDB.Part_clefID.Valid = true // allow for a 0 value (nil association)
		if for_part.Part_clef != nil {
			if Part_clefId, ok := backRepo.BackRepoPart_clef.Map_Part_clefPtr_Part_clefDBID[for_part.Part_clef]; ok {
				for_partDB.Part_clefID.Int64 = int64(Part_clefId)
				for_partDB.Part_clefID.Valid = true
			}
		} else {
			for_partDB.Part_clefID.Int64 = 0
			for_partDB.Part_clefID.Valid = true
		}

		// commit pointer value for_part.Part_transpose translates to updating the for_part.Part_transposeID
		for_partDB.Part_transposeID.Valid = true // allow for a 0 value (nil association)
		if for_part.Part_transpose != nil {
			if Part_transposeId, ok := backRepo.BackRepoPart_transpose.Map_Part_transposePtr_Part_transposeDBID[for_part.Part_transpose]; ok {
				for_partDB.Part_transposeID.Int64 = int64(Part_transposeId)
				for_partDB.Part_transposeID.Valid = true
			}
		} else {
			for_partDB.Part_transposeID.Int64 = 0
			for_partDB.Part_transposeID.Valid = true
		}

		query := backRepoFor_part.db.Save(&for_partDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown For_part intance %s", for_part.Name))
		return err
	}

	return
}

// BackRepoFor_part.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFor_part *BackRepoFor_partStruct) CheckoutPhaseOne() (Error error) {

	for_partDBArray := make([]For_partDB, 0)
	query := backRepoFor_part.db.Find(&for_partDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	for_partInstancesToBeRemovedFromTheStage := make(map[*models.For_part]any)
	for key, value := range backRepoFor_part.stage.For_parts {
		for_partInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, for_partDB := range for_partDBArray {
		backRepoFor_part.CheckoutPhaseOneInstance(&for_partDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		for_part, ok := backRepoFor_part.Map_For_partDBID_For_partPtr[for_partDB.ID]
		if ok {
			delete(for_partInstancesToBeRemovedFromTheStage, for_part)
		}
	}

	// remove from stage and back repo's 3 maps all for_parts that are not in the checkout
	for for_part := range for_partInstancesToBeRemovedFromTheStage {
		for_part.Unstage(backRepoFor_part.GetStage())

		// remove instance from the back repo 3 maps
		for_partID := backRepoFor_part.Map_For_partPtr_For_partDBID[for_part]
		delete(backRepoFor_part.Map_For_partPtr_For_partDBID, for_part)
		delete(backRepoFor_part.Map_For_partDBID_For_partDB, for_partID)
		delete(backRepoFor_part.Map_For_partDBID_For_partPtr, for_partID)
	}

	return
}

// CheckoutPhaseOneInstance takes a for_partDB that has been found in the DB, updates the backRepo and stages the
// models version of the for_partDB
func (backRepoFor_part *BackRepoFor_partStruct) CheckoutPhaseOneInstance(for_partDB *For_partDB) (Error error) {

	for_part, ok := backRepoFor_part.Map_For_partDBID_For_partPtr[for_partDB.ID]
	if !ok {
		for_part = new(models.For_part)

		backRepoFor_part.Map_For_partDBID_For_partPtr[for_partDB.ID] = for_part
		backRepoFor_part.Map_For_partPtr_For_partDBID[for_part] = for_partDB.ID

		// append model store with the new element
		for_part.Name = for_partDB.Name_Data.String
		for_part.Stage(backRepoFor_part.GetStage())
	}
	for_partDB.CopyBasicFieldsToFor_part(for_part)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	for_part.Stage(backRepoFor_part.GetStage())

	// preserve pointer to for_partDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_For_partDBID_For_partDB)[for_partDB hold variable pointers
	for_partDB_Data := *for_partDB
	preservedPtrToFor_part := &for_partDB_Data
	backRepoFor_part.Map_For_partDBID_For_partDB[for_partDB.ID] = preservedPtrToFor_part

	return
}

// BackRepoFor_part.CheckoutPhaseTwo Checkouts all staged instances of For_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFor_part *BackRepoFor_partStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, for_partDB := range backRepoFor_part.Map_For_partDBID_For_partDB {
		backRepoFor_part.CheckoutPhaseTwoInstance(backRepo, for_partDB)
	}
	return
}

// BackRepoFor_part.CheckoutPhaseTwoInstance Checkouts staged instances of For_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFor_part *BackRepoFor_partStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, for_partDB *For_partDB) (Error error) {

	for_part := backRepoFor_part.Map_For_partDBID_For_partPtr[for_partDB.ID]

	for_partDB.DecodePointers(backRepo, for_part)

	return
}

func (for_partDB *For_partDB) DecodePointers(backRepo *BackRepoStruct, for_part *models.For_part) {

	// insertion point for checkout of pointer encoding
	// Part_clef field
	for_part.Part_clef = nil
	if for_partDB.Part_clefID.Int64 != 0 {
		for_part.Part_clef = backRepo.BackRepoPart_clef.Map_Part_clefDBID_Part_clefPtr[uint(for_partDB.Part_clefID.Int64)]
	}
	// Part_transpose field
	for_part.Part_transpose = nil
	if for_partDB.Part_transposeID.Int64 != 0 {
		for_part.Part_transpose = backRepo.BackRepoPart_transpose.Map_Part_transposeDBID_Part_transposePtr[uint(for_partDB.Part_transposeID.Int64)]
	}
	return
}

// CommitFor_part allows commit of a single for_part (if already staged)
func (backRepo *BackRepoStruct) CommitFor_part(for_part *models.For_part) {
	backRepo.BackRepoFor_part.CommitPhaseOneInstance(for_part)
	if id, ok := backRepo.BackRepoFor_part.Map_For_partPtr_For_partDBID[for_part]; ok {
		backRepo.BackRepoFor_part.CommitPhaseTwoInstance(backRepo, id, for_part)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFor_part allows checkout of a single for_part (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFor_part(for_part *models.For_part) {
	// check if the for_part is staged
	if _, ok := backRepo.BackRepoFor_part.Map_For_partPtr_For_partDBID[for_part]; ok {

		if id, ok := backRepo.BackRepoFor_part.Map_For_partPtr_For_partDBID[for_part]; ok {
			var for_partDB For_partDB
			for_partDB.ID = id

			if err := backRepo.BackRepoFor_part.db.First(&for_partDB, id).Error; err != nil {
				log.Fatalln("CheckoutFor_part : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFor_part.CheckoutPhaseOneInstance(&for_partDB)
			backRepo.BackRepoFor_part.CheckoutPhaseTwoInstance(backRepo, &for_partDB)
		}
	}
}

// CopyBasicFieldsFromFor_part
func (for_partDB *For_partDB) CopyBasicFieldsFromFor_part(for_part *models.For_part) {
	// insertion point for fields commit

	for_partDB.Name_Data.String = for_part.Name
	for_partDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromFor_part_WOP
func (for_partDB *For_partDB) CopyBasicFieldsFromFor_part_WOP(for_part *models.For_part_WOP) {
	// insertion point for fields commit

	for_partDB.Name_Data.String = for_part.Name
	for_partDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromFor_partWOP
func (for_partDB *For_partDB) CopyBasicFieldsFromFor_partWOP(for_part *For_partWOP) {
	// insertion point for fields commit

	for_partDB.Name_Data.String = for_part.Name
	for_partDB.Name_Data.Valid = true
}

// CopyBasicFieldsToFor_part
func (for_partDB *For_partDB) CopyBasicFieldsToFor_part(for_part *models.For_part) {
	// insertion point for checkout of basic fields (back repo to stage)
	for_part.Name = for_partDB.Name_Data.String
}

// CopyBasicFieldsToFor_part_WOP
func (for_partDB *For_partDB) CopyBasicFieldsToFor_part_WOP(for_part *models.For_part_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	for_part.Name = for_partDB.Name_Data.String
}

// CopyBasicFieldsToFor_partWOP
func (for_partDB *For_partDB) CopyBasicFieldsToFor_partWOP(for_part *For_partWOP) {
	for_part.ID = int(for_partDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	for_part.Name = for_partDB.Name_Data.String
}

// Backup generates a json file from a slice of all For_partDB instances in the backrepo
func (backRepoFor_part *BackRepoFor_partStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "For_partDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*For_partDB, 0)
	for _, for_partDB := range backRepoFor_part.Map_For_partDBID_For_partDB {
		forBackup = append(forBackup, for_partDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json For_part ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json For_part file", err.Error())
	}
}

// Backup generates a json file from a slice of all For_partDB instances in the backrepo
func (backRepoFor_part *BackRepoFor_partStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*For_partDB, 0)
	for _, for_partDB := range backRepoFor_part.Map_For_partDBID_For_partDB {
		forBackup = append(forBackup, for_partDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("For_part")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&For_part_Fields, -1)
	for _, for_partDB := range forBackup {

		var for_partWOP For_partWOP
		for_partDB.CopyBasicFieldsToFor_partWOP(&for_partWOP)

		row := sh.AddRow()
		row.WriteStruct(&for_partWOP, -1)
	}
}

// RestoreXL from the "For_part" sheet all For_partDB instances
func (backRepoFor_part *BackRepoFor_partStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFor_partid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["For_part"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFor_part.rowVisitorFor_part)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoFor_part *BackRepoFor_partStruct) rowVisitorFor_part(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var for_partWOP For_partWOP
		row.ReadStruct(&for_partWOP)

		// add the unmarshalled struct to the stage
		for_partDB := new(For_partDB)
		for_partDB.CopyBasicFieldsFromFor_partWOP(&for_partWOP)

		for_partDB_ID_atBackupTime := for_partDB.ID
		for_partDB.ID = 0
		query := backRepoFor_part.db.Create(for_partDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFor_part.Map_For_partDBID_For_partDB[for_partDB.ID] = for_partDB
		BackRepoFor_partid_atBckpTime_newID[for_partDB_ID_atBackupTime] = for_partDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "For_partDB.json" in dirPath that stores an array
// of For_partDB and stores it in the database
// the map BackRepoFor_partid_atBckpTime_newID is updated accordingly
func (backRepoFor_part *BackRepoFor_partStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFor_partid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "For_partDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json For_part file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*For_partDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_For_partDBID_For_partDB
	for _, for_partDB := range forRestore {

		for_partDB_ID_atBackupTime := for_partDB.ID
		for_partDB.ID = 0
		query := backRepoFor_part.db.Create(for_partDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFor_part.Map_For_partDBID_For_partDB[for_partDB.ID] = for_partDB
		BackRepoFor_partid_atBckpTime_newID[for_partDB_ID_atBackupTime] = for_partDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json For_part file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<For_part>id_atBckpTime_newID
// to compute new index
func (backRepoFor_part *BackRepoFor_partStruct) RestorePhaseTwo() {

	for _, for_partDB := range backRepoFor_part.Map_For_partDBID_For_partDB {

		// next line of code is to avert unused variable compilation error
		_ = for_partDB

		// insertion point for reindexing pointers encoding
		// reindexing Part_clef field
		if for_partDB.Part_clefID.Int64 != 0 {
			for_partDB.Part_clefID.Int64 = int64(BackRepoPart_clefid_atBckpTime_newID[uint(for_partDB.Part_clefID.Int64)])
			for_partDB.Part_clefID.Valid = true
		}

		// reindexing Part_transpose field
		if for_partDB.Part_transposeID.Int64 != 0 {
			for_partDB.Part_transposeID.Int64 = int64(BackRepoPart_transposeid_atBckpTime_newID[uint(for_partDB.Part_transposeID.Int64)])
			for_partDB.Part_transposeID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoFor_part.db.Model(for_partDB).Updates(*for_partDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoFor_part.ResetReversePointers commits all staged instances of For_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFor_part *BackRepoFor_partStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, for_part := range backRepoFor_part.Map_For_partDBID_For_partPtr {
		backRepoFor_part.ResetReversePointersInstance(backRepo, idx, for_part)
	}

	return
}

func (backRepoFor_part *BackRepoFor_partStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, for_part *models.For_part) (Error error) {

	// fetch matching for_partDB
	if for_partDB, ok := backRepoFor_part.Map_For_partDBID_For_partDB[idx]; ok {
		_ = for_partDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFor_partid_atBckpTime_newID map[uint]uint
