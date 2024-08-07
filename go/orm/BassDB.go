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
var dummy_Bass_sql sql.NullBool
var dummy_Bass_time time.Duration
var dummy_Bass_sort sort.Float64Slice

// BassAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model bassAPI
type BassAPI struct {
	gorm.Model

	models.Bass_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	BassPointersEncoding BassPointersEncoding
}

// BassPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type BassPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Bass_step is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	Bass_stepID sql.NullInt64

	// field Bass_alter is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	Bass_alterID sql.NullInt64
}

// BassDB describes a bass in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model bassDB
type BassDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field bassDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	BassPointersEncoding
}

// BassDBs arrays bassDBs
// swagger:response bassDBsResponse
type BassDBs []BassDB

// BassDBResponse provides response
// swagger:response bassDBResponse
type BassDBResponse struct {
	BassDB
}

// BassWOP is a Bass without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type BassWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Bass_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoBassStruct struct {
	// stores BassDB according to their gorm ID
	Map_BassDBID_BassDB map[uint]*BassDB

	// stores BassDB ID according to Bass address
	Map_BassPtr_BassDBID map[*models.Bass]uint

	// stores Bass according to their gorm ID
	Map_BassDBID_BassPtr map[uint]*models.Bass

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoBass *BackRepoBassStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoBass.stage
	return
}

func (backRepoBass *BackRepoBassStruct) GetDB() *gorm.DB {
	return backRepoBass.db
}

// GetBassDBFromBassPtr is a handy function to access the back repo instance from the stage instance
func (backRepoBass *BackRepoBassStruct) GetBassDBFromBassPtr(bass *models.Bass) (bassDB *BassDB) {
	id := backRepoBass.Map_BassPtr_BassDBID[bass]
	bassDB = backRepoBass.Map_BassDBID_BassDB[id]
	return
}

// BackRepoBass.CommitPhaseOne commits all staged instances of Bass to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoBass *BackRepoBassStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for bass := range stage.Basss {
		backRepoBass.CommitPhaseOneInstance(bass)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, bass := range backRepoBass.Map_BassDBID_BassPtr {
		if _, ok := stage.Basss[bass]; !ok {
			backRepoBass.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoBass.CommitDeleteInstance commits deletion of Bass to the BackRepo
func (backRepoBass *BackRepoBassStruct) CommitDeleteInstance(id uint) (Error error) {

	bass := backRepoBass.Map_BassDBID_BassPtr[id]

	// bass is not staged anymore, remove bassDB
	bassDB := backRepoBass.Map_BassDBID_BassDB[id]
	query := backRepoBass.db.Unscoped().Delete(&bassDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoBass.Map_BassPtr_BassDBID, bass)
	delete(backRepoBass.Map_BassDBID_BassPtr, id)
	delete(backRepoBass.Map_BassDBID_BassDB, id)

	return
}

// BackRepoBass.CommitPhaseOneInstance commits bass staged instances of Bass to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoBass *BackRepoBassStruct) CommitPhaseOneInstance(bass *models.Bass) (Error error) {

	// check if the bass is not commited yet
	if _, ok := backRepoBass.Map_BassPtr_BassDBID[bass]; ok {
		return
	}

	// initiate bass
	var bassDB BassDB
	bassDB.CopyBasicFieldsFromBass(bass)

	query := backRepoBass.db.Create(&bassDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoBass.Map_BassPtr_BassDBID[bass] = bassDB.ID
	backRepoBass.Map_BassDBID_BassPtr[bassDB.ID] = bass
	backRepoBass.Map_BassDBID_BassDB[bassDB.ID] = &bassDB

	return
}

// BackRepoBass.CommitPhaseTwo commits all staged instances of Bass to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBass *BackRepoBassStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, bass := range backRepoBass.Map_BassDBID_BassPtr {
		backRepoBass.CommitPhaseTwoInstance(backRepo, idx, bass)
	}

	return
}

// BackRepoBass.CommitPhaseTwoInstance commits {{structname }} of models.Bass to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBass *BackRepoBassStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, bass *models.Bass) (Error error) {

	// fetch matching bassDB
	if bassDB, ok := backRepoBass.Map_BassDBID_BassDB[idx]; ok {

		bassDB.CopyBasicFieldsFromBass(bass)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value bass.Bass_step translates to updating the bass.Bass_stepID
		bassDB.Bass_stepID.Valid = true // allow for a 0 value (nil association)
		if bass.Bass_step != nil {
			if Bass_stepId, ok := backRepo.BackRepoBass_step.Map_Bass_stepPtr_Bass_stepDBID[bass.Bass_step]; ok {
				bassDB.Bass_stepID.Int64 = int64(Bass_stepId)
				bassDB.Bass_stepID.Valid = true
			}
		} else {
			bassDB.Bass_stepID.Int64 = 0
			bassDB.Bass_stepID.Valid = true
		}

		// commit pointer value bass.Bass_alter translates to updating the bass.Bass_alterID
		bassDB.Bass_alterID.Valid = true // allow for a 0 value (nil association)
		if bass.Bass_alter != nil {
			if Bass_alterId, ok := backRepo.BackRepoHarmony_alter.Map_Harmony_alterPtr_Harmony_alterDBID[bass.Bass_alter]; ok {
				bassDB.Bass_alterID.Int64 = int64(Bass_alterId)
				bassDB.Bass_alterID.Valid = true
			}
		} else {
			bassDB.Bass_alterID.Int64 = 0
			bassDB.Bass_alterID.Valid = true
		}

		query := backRepoBass.db.Save(&bassDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Bass intance %s", bass.Name))
		return err
	}

	return
}

// BackRepoBass.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoBass *BackRepoBassStruct) CheckoutPhaseOne() (Error error) {

	bassDBArray := make([]BassDB, 0)
	query := backRepoBass.db.Find(&bassDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	bassInstancesToBeRemovedFromTheStage := make(map[*models.Bass]any)
	for key, value := range backRepoBass.stage.Basss {
		bassInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, bassDB := range bassDBArray {
		backRepoBass.CheckoutPhaseOneInstance(&bassDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		bass, ok := backRepoBass.Map_BassDBID_BassPtr[bassDB.ID]
		if ok {
			delete(bassInstancesToBeRemovedFromTheStage, bass)
		}
	}

	// remove from stage and back repo's 3 maps all basss that are not in the checkout
	for bass := range bassInstancesToBeRemovedFromTheStage {
		bass.Unstage(backRepoBass.GetStage())

		// remove instance from the back repo 3 maps
		bassID := backRepoBass.Map_BassPtr_BassDBID[bass]
		delete(backRepoBass.Map_BassPtr_BassDBID, bass)
		delete(backRepoBass.Map_BassDBID_BassDB, bassID)
		delete(backRepoBass.Map_BassDBID_BassPtr, bassID)
	}

	return
}

// CheckoutPhaseOneInstance takes a bassDB that has been found in the DB, updates the backRepo and stages the
// models version of the bassDB
func (backRepoBass *BackRepoBassStruct) CheckoutPhaseOneInstance(bassDB *BassDB) (Error error) {

	bass, ok := backRepoBass.Map_BassDBID_BassPtr[bassDB.ID]
	if !ok {
		bass = new(models.Bass)

		backRepoBass.Map_BassDBID_BassPtr[bassDB.ID] = bass
		backRepoBass.Map_BassPtr_BassDBID[bass] = bassDB.ID

		// append model store with the new element
		bass.Name = bassDB.Name_Data.String
		bass.Stage(backRepoBass.GetStage())
	}
	bassDB.CopyBasicFieldsToBass(bass)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	bass.Stage(backRepoBass.GetStage())

	// preserve pointer to bassDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_BassDBID_BassDB)[bassDB hold variable pointers
	bassDB_Data := *bassDB
	preservedPtrToBass := &bassDB_Data
	backRepoBass.Map_BassDBID_BassDB[bassDB.ID] = preservedPtrToBass

	return
}

// BackRepoBass.CheckoutPhaseTwo Checkouts all staged instances of Bass to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBass *BackRepoBassStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, bassDB := range backRepoBass.Map_BassDBID_BassDB {
		backRepoBass.CheckoutPhaseTwoInstance(backRepo, bassDB)
	}
	return
}

// BackRepoBass.CheckoutPhaseTwoInstance Checkouts staged instances of Bass to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBass *BackRepoBassStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, bassDB *BassDB) (Error error) {

	bass := backRepoBass.Map_BassDBID_BassPtr[bassDB.ID]

	bassDB.DecodePointers(backRepo, bass)

	return
}

func (bassDB *BassDB) DecodePointers(backRepo *BackRepoStruct, bass *models.Bass) {

	// insertion point for checkout of pointer encoding
	// Bass_step field
	bass.Bass_step = nil
	if bassDB.Bass_stepID.Int64 != 0 {
		bass.Bass_step = backRepo.BackRepoBass_step.Map_Bass_stepDBID_Bass_stepPtr[uint(bassDB.Bass_stepID.Int64)]
	}
	// Bass_alter field
	bass.Bass_alter = nil
	if bassDB.Bass_alterID.Int64 != 0 {
		bass.Bass_alter = backRepo.BackRepoHarmony_alter.Map_Harmony_alterDBID_Harmony_alterPtr[uint(bassDB.Bass_alterID.Int64)]
	}
	return
}

// CommitBass allows commit of a single bass (if already staged)
func (backRepo *BackRepoStruct) CommitBass(bass *models.Bass) {
	backRepo.BackRepoBass.CommitPhaseOneInstance(bass)
	if id, ok := backRepo.BackRepoBass.Map_BassPtr_BassDBID[bass]; ok {
		backRepo.BackRepoBass.CommitPhaseTwoInstance(backRepo, id, bass)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitBass allows checkout of a single bass (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutBass(bass *models.Bass) {
	// check if the bass is staged
	if _, ok := backRepo.BackRepoBass.Map_BassPtr_BassDBID[bass]; ok {

		if id, ok := backRepo.BackRepoBass.Map_BassPtr_BassDBID[bass]; ok {
			var bassDB BassDB
			bassDB.ID = id

			if err := backRepo.BackRepoBass.db.First(&bassDB, id).Error; err != nil {
				log.Fatalln("CheckoutBass : Problem with getting object with id:", id)
			}
			backRepo.BackRepoBass.CheckoutPhaseOneInstance(&bassDB)
			backRepo.BackRepoBass.CheckoutPhaseTwoInstance(backRepo, &bassDB)
		}
	}
}

// CopyBasicFieldsFromBass
func (bassDB *BassDB) CopyBasicFieldsFromBass(bass *models.Bass) {
	// insertion point for fields commit

	bassDB.Name_Data.String = bass.Name
	bassDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromBass_WOP
func (bassDB *BassDB) CopyBasicFieldsFromBass_WOP(bass *models.Bass_WOP) {
	// insertion point for fields commit

	bassDB.Name_Data.String = bass.Name
	bassDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromBassWOP
func (bassDB *BassDB) CopyBasicFieldsFromBassWOP(bass *BassWOP) {
	// insertion point for fields commit

	bassDB.Name_Data.String = bass.Name
	bassDB.Name_Data.Valid = true
}

// CopyBasicFieldsToBass
func (bassDB *BassDB) CopyBasicFieldsToBass(bass *models.Bass) {
	// insertion point for checkout of basic fields (back repo to stage)
	bass.Name = bassDB.Name_Data.String
}

// CopyBasicFieldsToBass_WOP
func (bassDB *BassDB) CopyBasicFieldsToBass_WOP(bass *models.Bass_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	bass.Name = bassDB.Name_Data.String
}

// CopyBasicFieldsToBassWOP
func (bassDB *BassDB) CopyBasicFieldsToBassWOP(bass *BassWOP) {
	bass.ID = int(bassDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	bass.Name = bassDB.Name_Data.String
}

// Backup generates a json file from a slice of all BassDB instances in the backrepo
func (backRepoBass *BackRepoBassStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "BassDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*BassDB, 0)
	for _, bassDB := range backRepoBass.Map_BassDBID_BassDB {
		forBackup = append(forBackup, bassDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Bass ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Bass file", err.Error())
	}
}

// Backup generates a json file from a slice of all BassDB instances in the backrepo
func (backRepoBass *BackRepoBassStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*BassDB, 0)
	for _, bassDB := range backRepoBass.Map_BassDBID_BassDB {
		forBackup = append(forBackup, bassDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Bass")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Bass_Fields, -1)
	for _, bassDB := range forBackup {

		var bassWOP BassWOP
		bassDB.CopyBasicFieldsToBassWOP(&bassWOP)

		row := sh.AddRow()
		row.WriteStruct(&bassWOP, -1)
	}
}

// RestoreXL from the "Bass" sheet all BassDB instances
func (backRepoBass *BackRepoBassStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoBassid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Bass"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoBass.rowVisitorBass)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoBass *BackRepoBassStruct) rowVisitorBass(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var bassWOP BassWOP
		row.ReadStruct(&bassWOP)

		// add the unmarshalled struct to the stage
		bassDB := new(BassDB)
		bassDB.CopyBasicFieldsFromBassWOP(&bassWOP)

		bassDB_ID_atBackupTime := bassDB.ID
		bassDB.ID = 0
		query := backRepoBass.db.Create(bassDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoBass.Map_BassDBID_BassDB[bassDB.ID] = bassDB
		BackRepoBassid_atBckpTime_newID[bassDB_ID_atBackupTime] = bassDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "BassDB.json" in dirPath that stores an array
// of BassDB and stores it in the database
// the map BackRepoBassid_atBckpTime_newID is updated accordingly
func (backRepoBass *BackRepoBassStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoBassid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "BassDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Bass file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*BassDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_BassDBID_BassDB
	for _, bassDB := range forRestore {

		bassDB_ID_atBackupTime := bassDB.ID
		bassDB.ID = 0
		query := backRepoBass.db.Create(bassDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoBass.Map_BassDBID_BassDB[bassDB.ID] = bassDB
		BackRepoBassid_atBckpTime_newID[bassDB_ID_atBackupTime] = bassDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Bass file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Bass>id_atBckpTime_newID
// to compute new index
func (backRepoBass *BackRepoBassStruct) RestorePhaseTwo() {

	for _, bassDB := range backRepoBass.Map_BassDBID_BassDB {

		// next line of code is to avert unused variable compilation error
		_ = bassDB

		// insertion point for reindexing pointers encoding
		// reindexing Bass_step field
		if bassDB.Bass_stepID.Int64 != 0 {
			bassDB.Bass_stepID.Int64 = int64(BackRepoBass_stepid_atBckpTime_newID[uint(bassDB.Bass_stepID.Int64)])
			bassDB.Bass_stepID.Valid = true
		}

		// reindexing Bass_alter field
		if bassDB.Bass_alterID.Int64 != 0 {
			bassDB.Bass_alterID.Int64 = int64(BackRepoHarmony_alterid_atBckpTime_newID[uint(bassDB.Bass_alterID.Int64)])
			bassDB.Bass_alterID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoBass.db.Model(bassDB).Updates(*bassDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoBass.ResetReversePointers commits all staged instances of Bass to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBass *BackRepoBassStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, bass := range backRepoBass.Map_BassDBID_BassPtr {
		backRepoBass.ResetReversePointersInstance(backRepo, idx, bass)
	}

	return
}

func (backRepoBass *BackRepoBassStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, bass *models.Bass) (Error error) {

	// fetch matching bassDB
	if bassDB, ok := backRepoBass.Map_BassDBID_BassDB[idx]; ok {
		_ = bassDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoBassid_atBckpTime_newID map[uint]uint
