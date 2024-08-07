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
var dummy_Staff_tuning_sql sql.NullBool
var dummy_Staff_tuning_time time.Duration
var dummy_Staff_tuning_sort sort.Float64Slice

// Staff_tuningAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model staff_tuningAPI
type Staff_tuningAPI struct {
	gorm.Model

	models.Staff_tuning_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Staff_tuningPointersEncoding Staff_tuningPointersEncoding
}

// Staff_tuningPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Staff_tuningPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Staff_tuningDB describes a staff_tuning in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model staff_tuningDB
type Staff_tuningDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field staff_tuningDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Staff_tuningPointersEncoding
}

// Staff_tuningDBs arrays staff_tuningDBs
// swagger:response staff_tuningDBsResponse
type Staff_tuningDBs []Staff_tuningDB

// Staff_tuningDBResponse provides response
// swagger:response staff_tuningDBResponse
type Staff_tuningDBResponse struct {
	Staff_tuningDB
}

// Staff_tuningWOP is a Staff_tuning without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Staff_tuningWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Staff_tuning_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoStaff_tuningStruct struct {
	// stores Staff_tuningDB according to their gorm ID
	Map_Staff_tuningDBID_Staff_tuningDB map[uint]*Staff_tuningDB

	// stores Staff_tuningDB ID according to Staff_tuning address
	Map_Staff_tuningPtr_Staff_tuningDBID map[*models.Staff_tuning]uint

	// stores Staff_tuning according to their gorm ID
	Map_Staff_tuningDBID_Staff_tuningPtr map[uint]*models.Staff_tuning

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoStaff_tuning.stage
	return
}

func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) GetDB() *gorm.DB {
	return backRepoStaff_tuning.db
}

// GetStaff_tuningDBFromStaff_tuningPtr is a handy function to access the back repo instance from the stage instance
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) GetStaff_tuningDBFromStaff_tuningPtr(staff_tuning *models.Staff_tuning) (staff_tuningDB *Staff_tuningDB) {
	id := backRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID[staff_tuning]
	staff_tuningDB = backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB[id]
	return
}

// BackRepoStaff_tuning.CommitPhaseOne commits all staged instances of Staff_tuning to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for staff_tuning := range stage.Staff_tunings {
		backRepoStaff_tuning.CommitPhaseOneInstance(staff_tuning)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, staff_tuning := range backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr {
		if _, ok := stage.Staff_tunings[staff_tuning]; !ok {
			backRepoStaff_tuning.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoStaff_tuning.CommitDeleteInstance commits deletion of Staff_tuning to the BackRepo
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) CommitDeleteInstance(id uint) (Error error) {

	staff_tuning := backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr[id]

	// staff_tuning is not staged anymore, remove staff_tuningDB
	staff_tuningDB := backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB[id]
	query := backRepoStaff_tuning.db.Unscoped().Delete(&staff_tuningDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID, staff_tuning)
	delete(backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr, id)
	delete(backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB, id)

	return
}

// BackRepoStaff_tuning.CommitPhaseOneInstance commits staff_tuning staged instances of Staff_tuning to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) CommitPhaseOneInstance(staff_tuning *models.Staff_tuning) (Error error) {

	// check if the staff_tuning is not commited yet
	if _, ok := backRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID[staff_tuning]; ok {
		return
	}

	// initiate staff_tuning
	var staff_tuningDB Staff_tuningDB
	staff_tuningDB.CopyBasicFieldsFromStaff_tuning(staff_tuning)

	query := backRepoStaff_tuning.db.Create(&staff_tuningDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID[staff_tuning] = staff_tuningDB.ID
	backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr[staff_tuningDB.ID] = staff_tuning
	backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB[staff_tuningDB.ID] = &staff_tuningDB

	return
}

// BackRepoStaff_tuning.CommitPhaseTwo commits all staged instances of Staff_tuning to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, staff_tuning := range backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr {
		backRepoStaff_tuning.CommitPhaseTwoInstance(backRepo, idx, staff_tuning)
	}

	return
}

// BackRepoStaff_tuning.CommitPhaseTwoInstance commits {{structname }} of models.Staff_tuning to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, staff_tuning *models.Staff_tuning) (Error error) {

	// fetch matching staff_tuningDB
	if staff_tuningDB, ok := backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB[idx]; ok {

		staff_tuningDB.CopyBasicFieldsFromStaff_tuning(staff_tuning)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoStaff_tuning.db.Save(&staff_tuningDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Staff_tuning intance %s", staff_tuning.Name))
		return err
	}

	return
}

// BackRepoStaff_tuning.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) CheckoutPhaseOne() (Error error) {

	staff_tuningDBArray := make([]Staff_tuningDB, 0)
	query := backRepoStaff_tuning.db.Find(&staff_tuningDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	staff_tuningInstancesToBeRemovedFromTheStage := make(map[*models.Staff_tuning]any)
	for key, value := range backRepoStaff_tuning.stage.Staff_tunings {
		staff_tuningInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, staff_tuningDB := range staff_tuningDBArray {
		backRepoStaff_tuning.CheckoutPhaseOneInstance(&staff_tuningDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		staff_tuning, ok := backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr[staff_tuningDB.ID]
		if ok {
			delete(staff_tuningInstancesToBeRemovedFromTheStage, staff_tuning)
		}
	}

	// remove from stage and back repo's 3 maps all staff_tunings that are not in the checkout
	for staff_tuning := range staff_tuningInstancesToBeRemovedFromTheStage {
		staff_tuning.Unstage(backRepoStaff_tuning.GetStage())

		// remove instance from the back repo 3 maps
		staff_tuningID := backRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID[staff_tuning]
		delete(backRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID, staff_tuning)
		delete(backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB, staff_tuningID)
		delete(backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr, staff_tuningID)
	}

	return
}

// CheckoutPhaseOneInstance takes a staff_tuningDB that has been found in the DB, updates the backRepo and stages the
// models version of the staff_tuningDB
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) CheckoutPhaseOneInstance(staff_tuningDB *Staff_tuningDB) (Error error) {

	staff_tuning, ok := backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr[staff_tuningDB.ID]
	if !ok {
		staff_tuning = new(models.Staff_tuning)

		backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr[staff_tuningDB.ID] = staff_tuning
		backRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID[staff_tuning] = staff_tuningDB.ID

		// append model store with the new element
		staff_tuning.Name = staff_tuningDB.Name_Data.String
		staff_tuning.Stage(backRepoStaff_tuning.GetStage())
	}
	staff_tuningDB.CopyBasicFieldsToStaff_tuning(staff_tuning)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	staff_tuning.Stage(backRepoStaff_tuning.GetStage())

	// preserve pointer to staff_tuningDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Staff_tuningDBID_Staff_tuningDB)[staff_tuningDB hold variable pointers
	staff_tuningDB_Data := *staff_tuningDB
	preservedPtrToStaff_tuning := &staff_tuningDB_Data
	backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB[staff_tuningDB.ID] = preservedPtrToStaff_tuning

	return
}

// BackRepoStaff_tuning.CheckoutPhaseTwo Checkouts all staged instances of Staff_tuning to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, staff_tuningDB := range backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB {
		backRepoStaff_tuning.CheckoutPhaseTwoInstance(backRepo, staff_tuningDB)
	}
	return
}

// BackRepoStaff_tuning.CheckoutPhaseTwoInstance Checkouts staged instances of Staff_tuning to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, staff_tuningDB *Staff_tuningDB) (Error error) {

	staff_tuning := backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr[staff_tuningDB.ID]

	staff_tuningDB.DecodePointers(backRepo, staff_tuning)

	return
}

func (staff_tuningDB *Staff_tuningDB) DecodePointers(backRepo *BackRepoStruct, staff_tuning *models.Staff_tuning) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitStaff_tuning allows commit of a single staff_tuning (if already staged)
func (backRepo *BackRepoStruct) CommitStaff_tuning(staff_tuning *models.Staff_tuning) {
	backRepo.BackRepoStaff_tuning.CommitPhaseOneInstance(staff_tuning)
	if id, ok := backRepo.BackRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID[staff_tuning]; ok {
		backRepo.BackRepoStaff_tuning.CommitPhaseTwoInstance(backRepo, id, staff_tuning)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitStaff_tuning allows checkout of a single staff_tuning (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutStaff_tuning(staff_tuning *models.Staff_tuning) {
	// check if the staff_tuning is staged
	if _, ok := backRepo.BackRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID[staff_tuning]; ok {

		if id, ok := backRepo.BackRepoStaff_tuning.Map_Staff_tuningPtr_Staff_tuningDBID[staff_tuning]; ok {
			var staff_tuningDB Staff_tuningDB
			staff_tuningDB.ID = id

			if err := backRepo.BackRepoStaff_tuning.db.First(&staff_tuningDB, id).Error; err != nil {
				log.Fatalln("CheckoutStaff_tuning : Problem with getting object with id:", id)
			}
			backRepo.BackRepoStaff_tuning.CheckoutPhaseOneInstance(&staff_tuningDB)
			backRepo.BackRepoStaff_tuning.CheckoutPhaseTwoInstance(backRepo, &staff_tuningDB)
		}
	}
}

// CopyBasicFieldsFromStaff_tuning
func (staff_tuningDB *Staff_tuningDB) CopyBasicFieldsFromStaff_tuning(staff_tuning *models.Staff_tuning) {
	// insertion point for fields commit

	staff_tuningDB.Name_Data.String = staff_tuning.Name
	staff_tuningDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromStaff_tuning_WOP
func (staff_tuningDB *Staff_tuningDB) CopyBasicFieldsFromStaff_tuning_WOP(staff_tuning *models.Staff_tuning_WOP) {
	// insertion point for fields commit

	staff_tuningDB.Name_Data.String = staff_tuning.Name
	staff_tuningDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromStaff_tuningWOP
func (staff_tuningDB *Staff_tuningDB) CopyBasicFieldsFromStaff_tuningWOP(staff_tuning *Staff_tuningWOP) {
	// insertion point for fields commit

	staff_tuningDB.Name_Data.String = staff_tuning.Name
	staff_tuningDB.Name_Data.Valid = true
}

// CopyBasicFieldsToStaff_tuning
func (staff_tuningDB *Staff_tuningDB) CopyBasicFieldsToStaff_tuning(staff_tuning *models.Staff_tuning) {
	// insertion point for checkout of basic fields (back repo to stage)
	staff_tuning.Name = staff_tuningDB.Name_Data.String
}

// CopyBasicFieldsToStaff_tuning_WOP
func (staff_tuningDB *Staff_tuningDB) CopyBasicFieldsToStaff_tuning_WOP(staff_tuning *models.Staff_tuning_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	staff_tuning.Name = staff_tuningDB.Name_Data.String
}

// CopyBasicFieldsToStaff_tuningWOP
func (staff_tuningDB *Staff_tuningDB) CopyBasicFieldsToStaff_tuningWOP(staff_tuning *Staff_tuningWOP) {
	staff_tuning.ID = int(staff_tuningDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	staff_tuning.Name = staff_tuningDB.Name_Data.String
}

// Backup generates a json file from a slice of all Staff_tuningDB instances in the backrepo
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Staff_tuningDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Staff_tuningDB, 0)
	for _, staff_tuningDB := range backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB {
		forBackup = append(forBackup, staff_tuningDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Staff_tuning ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Staff_tuning file", err.Error())
	}
}

// Backup generates a json file from a slice of all Staff_tuningDB instances in the backrepo
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Staff_tuningDB, 0)
	for _, staff_tuningDB := range backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB {
		forBackup = append(forBackup, staff_tuningDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Staff_tuning")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Staff_tuning_Fields, -1)
	for _, staff_tuningDB := range forBackup {

		var staff_tuningWOP Staff_tuningWOP
		staff_tuningDB.CopyBasicFieldsToStaff_tuningWOP(&staff_tuningWOP)

		row := sh.AddRow()
		row.WriteStruct(&staff_tuningWOP, -1)
	}
}

// RestoreXL from the "Staff_tuning" sheet all Staff_tuningDB instances
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoStaff_tuningid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Staff_tuning"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoStaff_tuning.rowVisitorStaff_tuning)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) rowVisitorStaff_tuning(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var staff_tuningWOP Staff_tuningWOP
		row.ReadStruct(&staff_tuningWOP)

		// add the unmarshalled struct to the stage
		staff_tuningDB := new(Staff_tuningDB)
		staff_tuningDB.CopyBasicFieldsFromStaff_tuningWOP(&staff_tuningWOP)

		staff_tuningDB_ID_atBackupTime := staff_tuningDB.ID
		staff_tuningDB.ID = 0
		query := backRepoStaff_tuning.db.Create(staff_tuningDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB[staff_tuningDB.ID] = staff_tuningDB
		BackRepoStaff_tuningid_atBckpTime_newID[staff_tuningDB_ID_atBackupTime] = staff_tuningDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Staff_tuningDB.json" in dirPath that stores an array
// of Staff_tuningDB and stores it in the database
// the map BackRepoStaff_tuningid_atBckpTime_newID is updated accordingly
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoStaff_tuningid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Staff_tuningDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Staff_tuning file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Staff_tuningDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Staff_tuningDBID_Staff_tuningDB
	for _, staff_tuningDB := range forRestore {

		staff_tuningDB_ID_atBackupTime := staff_tuningDB.ID
		staff_tuningDB.ID = 0
		query := backRepoStaff_tuning.db.Create(staff_tuningDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB[staff_tuningDB.ID] = staff_tuningDB
		BackRepoStaff_tuningid_atBckpTime_newID[staff_tuningDB_ID_atBackupTime] = staff_tuningDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Staff_tuning file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Staff_tuning>id_atBckpTime_newID
// to compute new index
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) RestorePhaseTwo() {

	for _, staff_tuningDB := range backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB {

		// next line of code is to avert unused variable compilation error
		_ = staff_tuningDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoStaff_tuning.db.Model(staff_tuningDB).Updates(*staff_tuningDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoStaff_tuning.ResetReversePointers commits all staged instances of Staff_tuning to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, staff_tuning := range backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr {
		backRepoStaff_tuning.ResetReversePointersInstance(backRepo, idx, staff_tuning)
	}

	return
}

func (backRepoStaff_tuning *BackRepoStaff_tuningStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, staff_tuning *models.Staff_tuning) (Error error) {

	// fetch matching staff_tuningDB
	if staff_tuningDB, ok := backRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningDB[idx]; ok {
		_ = staff_tuningDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoStaff_tuningid_atBckpTime_newID map[uint]uint
