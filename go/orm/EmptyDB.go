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
var dummy_Empty_sql sql.NullBool
var dummy_Empty_time time.Duration
var dummy_Empty_sort sort.Float64Slice

// EmptyAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model emptyAPI
type EmptyAPI struct {
	gorm.Model

	models.Empty_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	EmptyPointersEncoding EmptyPointersEncoding
}

// EmptyPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type EmptyPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// EmptyDB describes a empty in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model emptyDB
type EmptyDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field emptyDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	EmptyPointersEncoding
}

// EmptyDBs arrays emptyDBs
// swagger:response emptyDBsResponse
type EmptyDBs []EmptyDB

// EmptyDBResponse provides response
// swagger:response emptyDBResponse
type EmptyDBResponse struct {
	EmptyDB
}

// EmptyWOP is a Empty without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type EmptyWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Empty_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoEmptyStruct struct {
	// stores EmptyDB according to their gorm ID
	Map_EmptyDBID_EmptyDB map[uint]*EmptyDB

	// stores EmptyDB ID according to Empty address
	Map_EmptyPtr_EmptyDBID map[*models.Empty]uint

	// stores Empty according to their gorm ID
	Map_EmptyDBID_EmptyPtr map[uint]*models.Empty

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoEmpty *BackRepoEmptyStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoEmpty.stage
	return
}

func (backRepoEmpty *BackRepoEmptyStruct) GetDB() *gorm.DB {
	return backRepoEmpty.db
}

// GetEmptyDBFromEmptyPtr is a handy function to access the back repo instance from the stage instance
func (backRepoEmpty *BackRepoEmptyStruct) GetEmptyDBFromEmptyPtr(empty *models.Empty) (emptyDB *EmptyDB) {
	id := backRepoEmpty.Map_EmptyPtr_EmptyDBID[empty]
	emptyDB = backRepoEmpty.Map_EmptyDBID_EmptyDB[id]
	return
}

// BackRepoEmpty.CommitPhaseOne commits all staged instances of Empty to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoEmpty *BackRepoEmptyStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for empty := range stage.Emptys {
		backRepoEmpty.CommitPhaseOneInstance(empty)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, empty := range backRepoEmpty.Map_EmptyDBID_EmptyPtr {
		if _, ok := stage.Emptys[empty]; !ok {
			backRepoEmpty.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoEmpty.CommitDeleteInstance commits deletion of Empty to the BackRepo
func (backRepoEmpty *BackRepoEmptyStruct) CommitDeleteInstance(id uint) (Error error) {

	empty := backRepoEmpty.Map_EmptyDBID_EmptyPtr[id]

	// empty is not staged anymore, remove emptyDB
	emptyDB := backRepoEmpty.Map_EmptyDBID_EmptyDB[id]
	query := backRepoEmpty.db.Unscoped().Delete(&emptyDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoEmpty.Map_EmptyPtr_EmptyDBID, empty)
	delete(backRepoEmpty.Map_EmptyDBID_EmptyPtr, id)
	delete(backRepoEmpty.Map_EmptyDBID_EmptyDB, id)

	return
}

// BackRepoEmpty.CommitPhaseOneInstance commits empty staged instances of Empty to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoEmpty *BackRepoEmptyStruct) CommitPhaseOneInstance(empty *models.Empty) (Error error) {

	// check if the empty is not commited yet
	if _, ok := backRepoEmpty.Map_EmptyPtr_EmptyDBID[empty]; ok {
		return
	}

	// initiate empty
	var emptyDB EmptyDB
	emptyDB.CopyBasicFieldsFromEmpty(empty)

	query := backRepoEmpty.db.Create(&emptyDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoEmpty.Map_EmptyPtr_EmptyDBID[empty] = emptyDB.ID
	backRepoEmpty.Map_EmptyDBID_EmptyPtr[emptyDB.ID] = empty
	backRepoEmpty.Map_EmptyDBID_EmptyDB[emptyDB.ID] = &emptyDB

	return
}

// BackRepoEmpty.CommitPhaseTwo commits all staged instances of Empty to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty *BackRepoEmptyStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, empty := range backRepoEmpty.Map_EmptyDBID_EmptyPtr {
		backRepoEmpty.CommitPhaseTwoInstance(backRepo, idx, empty)
	}

	return
}

// BackRepoEmpty.CommitPhaseTwoInstance commits {{structname }} of models.Empty to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty *BackRepoEmptyStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, empty *models.Empty) (Error error) {

	// fetch matching emptyDB
	if emptyDB, ok := backRepoEmpty.Map_EmptyDBID_EmptyDB[idx]; ok {

		emptyDB.CopyBasicFieldsFromEmpty(empty)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoEmpty.db.Save(&emptyDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Empty intance %s", empty.Name))
		return err
	}

	return
}

// BackRepoEmpty.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoEmpty *BackRepoEmptyStruct) CheckoutPhaseOne() (Error error) {

	emptyDBArray := make([]EmptyDB, 0)
	query := backRepoEmpty.db.Find(&emptyDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	emptyInstancesToBeRemovedFromTheStage := make(map[*models.Empty]any)
	for key, value := range backRepoEmpty.stage.Emptys {
		emptyInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, emptyDB := range emptyDBArray {
		backRepoEmpty.CheckoutPhaseOneInstance(&emptyDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		empty, ok := backRepoEmpty.Map_EmptyDBID_EmptyPtr[emptyDB.ID]
		if ok {
			delete(emptyInstancesToBeRemovedFromTheStage, empty)
		}
	}

	// remove from stage and back repo's 3 maps all emptys that are not in the checkout
	for empty := range emptyInstancesToBeRemovedFromTheStage {
		empty.Unstage(backRepoEmpty.GetStage())

		// remove instance from the back repo 3 maps
		emptyID := backRepoEmpty.Map_EmptyPtr_EmptyDBID[empty]
		delete(backRepoEmpty.Map_EmptyPtr_EmptyDBID, empty)
		delete(backRepoEmpty.Map_EmptyDBID_EmptyDB, emptyID)
		delete(backRepoEmpty.Map_EmptyDBID_EmptyPtr, emptyID)
	}

	return
}

// CheckoutPhaseOneInstance takes a emptyDB that has been found in the DB, updates the backRepo and stages the
// models version of the emptyDB
func (backRepoEmpty *BackRepoEmptyStruct) CheckoutPhaseOneInstance(emptyDB *EmptyDB) (Error error) {

	empty, ok := backRepoEmpty.Map_EmptyDBID_EmptyPtr[emptyDB.ID]
	if !ok {
		empty = new(models.Empty)

		backRepoEmpty.Map_EmptyDBID_EmptyPtr[emptyDB.ID] = empty
		backRepoEmpty.Map_EmptyPtr_EmptyDBID[empty] = emptyDB.ID

		// append model store with the new element
		empty.Name = emptyDB.Name_Data.String
		empty.Stage(backRepoEmpty.GetStage())
	}
	emptyDB.CopyBasicFieldsToEmpty(empty)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	empty.Stage(backRepoEmpty.GetStage())

	// preserve pointer to emptyDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_EmptyDBID_EmptyDB)[emptyDB hold variable pointers
	emptyDB_Data := *emptyDB
	preservedPtrToEmpty := &emptyDB_Data
	backRepoEmpty.Map_EmptyDBID_EmptyDB[emptyDB.ID] = preservedPtrToEmpty

	return
}

// BackRepoEmpty.CheckoutPhaseTwo Checkouts all staged instances of Empty to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty *BackRepoEmptyStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, emptyDB := range backRepoEmpty.Map_EmptyDBID_EmptyDB {
		backRepoEmpty.CheckoutPhaseTwoInstance(backRepo, emptyDB)
	}
	return
}

// BackRepoEmpty.CheckoutPhaseTwoInstance Checkouts staged instances of Empty to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty *BackRepoEmptyStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, emptyDB *EmptyDB) (Error error) {

	empty := backRepoEmpty.Map_EmptyDBID_EmptyPtr[emptyDB.ID]

	emptyDB.DecodePointers(backRepo, empty)

	return
}

func (emptyDB *EmptyDB) DecodePointers(backRepo *BackRepoStruct, empty *models.Empty) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitEmpty allows commit of a single empty (if already staged)
func (backRepo *BackRepoStruct) CommitEmpty(empty *models.Empty) {
	backRepo.BackRepoEmpty.CommitPhaseOneInstance(empty)
	if id, ok := backRepo.BackRepoEmpty.Map_EmptyPtr_EmptyDBID[empty]; ok {
		backRepo.BackRepoEmpty.CommitPhaseTwoInstance(backRepo, id, empty)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitEmpty allows checkout of a single empty (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutEmpty(empty *models.Empty) {
	// check if the empty is staged
	if _, ok := backRepo.BackRepoEmpty.Map_EmptyPtr_EmptyDBID[empty]; ok {

		if id, ok := backRepo.BackRepoEmpty.Map_EmptyPtr_EmptyDBID[empty]; ok {
			var emptyDB EmptyDB
			emptyDB.ID = id

			if err := backRepo.BackRepoEmpty.db.First(&emptyDB, id).Error; err != nil {
				log.Fatalln("CheckoutEmpty : Problem with getting object with id:", id)
			}
			backRepo.BackRepoEmpty.CheckoutPhaseOneInstance(&emptyDB)
			backRepo.BackRepoEmpty.CheckoutPhaseTwoInstance(backRepo, &emptyDB)
		}
	}
}

// CopyBasicFieldsFromEmpty
func (emptyDB *EmptyDB) CopyBasicFieldsFromEmpty(empty *models.Empty) {
	// insertion point for fields commit

	emptyDB.Name_Data.String = empty.Name
	emptyDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromEmpty_WOP
func (emptyDB *EmptyDB) CopyBasicFieldsFromEmpty_WOP(empty *models.Empty_WOP) {
	// insertion point for fields commit

	emptyDB.Name_Data.String = empty.Name
	emptyDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromEmptyWOP
func (emptyDB *EmptyDB) CopyBasicFieldsFromEmptyWOP(empty *EmptyWOP) {
	// insertion point for fields commit

	emptyDB.Name_Data.String = empty.Name
	emptyDB.Name_Data.Valid = true
}

// CopyBasicFieldsToEmpty
func (emptyDB *EmptyDB) CopyBasicFieldsToEmpty(empty *models.Empty) {
	// insertion point for checkout of basic fields (back repo to stage)
	empty.Name = emptyDB.Name_Data.String
}

// CopyBasicFieldsToEmpty_WOP
func (emptyDB *EmptyDB) CopyBasicFieldsToEmpty_WOP(empty *models.Empty_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	empty.Name = emptyDB.Name_Data.String
}

// CopyBasicFieldsToEmptyWOP
func (emptyDB *EmptyDB) CopyBasicFieldsToEmptyWOP(empty *EmptyWOP) {
	empty.ID = int(emptyDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	empty.Name = emptyDB.Name_Data.String
}

// Backup generates a json file from a slice of all EmptyDB instances in the backrepo
func (backRepoEmpty *BackRepoEmptyStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "EmptyDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*EmptyDB, 0)
	for _, emptyDB := range backRepoEmpty.Map_EmptyDBID_EmptyDB {
		forBackup = append(forBackup, emptyDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Empty ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Empty file", err.Error())
	}
}

// Backup generates a json file from a slice of all EmptyDB instances in the backrepo
func (backRepoEmpty *BackRepoEmptyStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*EmptyDB, 0)
	for _, emptyDB := range backRepoEmpty.Map_EmptyDBID_EmptyDB {
		forBackup = append(forBackup, emptyDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Empty")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Empty_Fields, -1)
	for _, emptyDB := range forBackup {

		var emptyWOP EmptyWOP
		emptyDB.CopyBasicFieldsToEmptyWOP(&emptyWOP)

		row := sh.AddRow()
		row.WriteStruct(&emptyWOP, -1)
	}
}

// RestoreXL from the "Empty" sheet all EmptyDB instances
func (backRepoEmpty *BackRepoEmptyStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoEmptyid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Empty"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoEmpty.rowVisitorEmpty)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoEmpty *BackRepoEmptyStruct) rowVisitorEmpty(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var emptyWOP EmptyWOP
		row.ReadStruct(&emptyWOP)

		// add the unmarshalled struct to the stage
		emptyDB := new(EmptyDB)
		emptyDB.CopyBasicFieldsFromEmptyWOP(&emptyWOP)

		emptyDB_ID_atBackupTime := emptyDB.ID
		emptyDB.ID = 0
		query := backRepoEmpty.db.Create(emptyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoEmpty.Map_EmptyDBID_EmptyDB[emptyDB.ID] = emptyDB
		BackRepoEmptyid_atBckpTime_newID[emptyDB_ID_atBackupTime] = emptyDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "EmptyDB.json" in dirPath that stores an array
// of EmptyDB and stores it in the database
// the map BackRepoEmptyid_atBckpTime_newID is updated accordingly
func (backRepoEmpty *BackRepoEmptyStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoEmptyid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "EmptyDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Empty file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*EmptyDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_EmptyDBID_EmptyDB
	for _, emptyDB := range forRestore {

		emptyDB_ID_atBackupTime := emptyDB.ID
		emptyDB.ID = 0
		query := backRepoEmpty.db.Create(emptyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoEmpty.Map_EmptyDBID_EmptyDB[emptyDB.ID] = emptyDB
		BackRepoEmptyid_atBckpTime_newID[emptyDB_ID_atBackupTime] = emptyDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Empty file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Empty>id_atBckpTime_newID
// to compute new index
func (backRepoEmpty *BackRepoEmptyStruct) RestorePhaseTwo() {

	for _, emptyDB := range backRepoEmpty.Map_EmptyDBID_EmptyDB {

		// next line of code is to avert unused variable compilation error
		_ = emptyDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoEmpty.db.Model(emptyDB).Updates(*emptyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoEmpty.ResetReversePointers commits all staged instances of Empty to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty *BackRepoEmptyStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, empty := range backRepoEmpty.Map_EmptyDBID_EmptyPtr {
		backRepoEmpty.ResetReversePointersInstance(backRepo, idx, empty)
	}

	return
}

func (backRepoEmpty *BackRepoEmptyStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, empty *models.Empty) (Error error) {

	// fetch matching emptyDB
	if emptyDB, ok := backRepoEmpty.Map_EmptyDBID_EmptyDB[idx]; ok {
		_ = emptyDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoEmptyid_atBckpTime_newID map[uint]uint
