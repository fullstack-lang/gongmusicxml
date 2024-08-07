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
var dummy_Accidental_sql sql.NullBool
var dummy_Accidental_time time.Duration
var dummy_Accidental_sort sort.Float64Slice

// AccidentalAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model accidentalAPI
type AccidentalAPI struct {
	gorm.Model

	models.Accidental_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	AccidentalPointersEncoding AccidentalPointersEncoding
}

// AccidentalPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type AccidentalPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// AccidentalDB describes a accidental in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model accidentalDB
type AccidentalDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field accidentalDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	AccidentalPointersEncoding
}

// AccidentalDBs arrays accidentalDBs
// swagger:response accidentalDBsResponse
type AccidentalDBs []AccidentalDB

// AccidentalDBResponse provides response
// swagger:response accidentalDBResponse
type AccidentalDBResponse struct {
	AccidentalDB
}

// AccidentalWOP is a Accidental without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type AccidentalWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Accidental_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoAccidentalStruct struct {
	// stores AccidentalDB according to their gorm ID
	Map_AccidentalDBID_AccidentalDB map[uint]*AccidentalDB

	// stores AccidentalDB ID according to Accidental address
	Map_AccidentalPtr_AccidentalDBID map[*models.Accidental]uint

	// stores Accidental according to their gorm ID
	Map_AccidentalDBID_AccidentalPtr map[uint]*models.Accidental

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoAccidental *BackRepoAccidentalStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoAccidental.stage
	return
}

func (backRepoAccidental *BackRepoAccidentalStruct) GetDB() *gorm.DB {
	return backRepoAccidental.db
}

// GetAccidentalDBFromAccidentalPtr is a handy function to access the back repo instance from the stage instance
func (backRepoAccidental *BackRepoAccidentalStruct) GetAccidentalDBFromAccidentalPtr(accidental *models.Accidental) (accidentalDB *AccidentalDB) {
	id := backRepoAccidental.Map_AccidentalPtr_AccidentalDBID[accidental]
	accidentalDB = backRepoAccidental.Map_AccidentalDBID_AccidentalDB[id]
	return
}

// BackRepoAccidental.CommitPhaseOne commits all staged instances of Accidental to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAccidental *BackRepoAccidentalStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for accidental := range stage.Accidentals {
		backRepoAccidental.CommitPhaseOneInstance(accidental)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, accidental := range backRepoAccidental.Map_AccidentalDBID_AccidentalPtr {
		if _, ok := stage.Accidentals[accidental]; !ok {
			backRepoAccidental.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoAccidental.CommitDeleteInstance commits deletion of Accidental to the BackRepo
func (backRepoAccidental *BackRepoAccidentalStruct) CommitDeleteInstance(id uint) (Error error) {

	accidental := backRepoAccidental.Map_AccidentalDBID_AccidentalPtr[id]

	// accidental is not staged anymore, remove accidentalDB
	accidentalDB := backRepoAccidental.Map_AccidentalDBID_AccidentalDB[id]
	query := backRepoAccidental.db.Unscoped().Delete(&accidentalDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoAccidental.Map_AccidentalPtr_AccidentalDBID, accidental)
	delete(backRepoAccidental.Map_AccidentalDBID_AccidentalPtr, id)
	delete(backRepoAccidental.Map_AccidentalDBID_AccidentalDB, id)

	return
}

// BackRepoAccidental.CommitPhaseOneInstance commits accidental staged instances of Accidental to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAccidental *BackRepoAccidentalStruct) CommitPhaseOneInstance(accidental *models.Accidental) (Error error) {

	// check if the accidental is not commited yet
	if _, ok := backRepoAccidental.Map_AccidentalPtr_AccidentalDBID[accidental]; ok {
		return
	}

	// initiate accidental
	var accidentalDB AccidentalDB
	accidentalDB.CopyBasicFieldsFromAccidental(accidental)

	query := backRepoAccidental.db.Create(&accidentalDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoAccidental.Map_AccidentalPtr_AccidentalDBID[accidental] = accidentalDB.ID
	backRepoAccidental.Map_AccidentalDBID_AccidentalPtr[accidentalDB.ID] = accidental
	backRepoAccidental.Map_AccidentalDBID_AccidentalDB[accidentalDB.ID] = &accidentalDB

	return
}

// BackRepoAccidental.CommitPhaseTwo commits all staged instances of Accidental to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccidental *BackRepoAccidentalStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, accidental := range backRepoAccidental.Map_AccidentalDBID_AccidentalPtr {
		backRepoAccidental.CommitPhaseTwoInstance(backRepo, idx, accidental)
	}

	return
}

// BackRepoAccidental.CommitPhaseTwoInstance commits {{structname }} of models.Accidental to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccidental *BackRepoAccidentalStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, accidental *models.Accidental) (Error error) {

	// fetch matching accidentalDB
	if accidentalDB, ok := backRepoAccidental.Map_AccidentalDBID_AccidentalDB[idx]; ok {

		accidentalDB.CopyBasicFieldsFromAccidental(accidental)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoAccidental.db.Save(&accidentalDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Accidental intance %s", accidental.Name))
		return err
	}

	return
}

// BackRepoAccidental.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoAccidental *BackRepoAccidentalStruct) CheckoutPhaseOne() (Error error) {

	accidentalDBArray := make([]AccidentalDB, 0)
	query := backRepoAccidental.db.Find(&accidentalDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	accidentalInstancesToBeRemovedFromTheStage := make(map[*models.Accidental]any)
	for key, value := range backRepoAccidental.stage.Accidentals {
		accidentalInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, accidentalDB := range accidentalDBArray {
		backRepoAccidental.CheckoutPhaseOneInstance(&accidentalDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		accidental, ok := backRepoAccidental.Map_AccidentalDBID_AccidentalPtr[accidentalDB.ID]
		if ok {
			delete(accidentalInstancesToBeRemovedFromTheStage, accidental)
		}
	}

	// remove from stage and back repo's 3 maps all accidentals that are not in the checkout
	for accidental := range accidentalInstancesToBeRemovedFromTheStage {
		accidental.Unstage(backRepoAccidental.GetStage())

		// remove instance from the back repo 3 maps
		accidentalID := backRepoAccidental.Map_AccidentalPtr_AccidentalDBID[accidental]
		delete(backRepoAccidental.Map_AccidentalPtr_AccidentalDBID, accidental)
		delete(backRepoAccidental.Map_AccidentalDBID_AccidentalDB, accidentalID)
		delete(backRepoAccidental.Map_AccidentalDBID_AccidentalPtr, accidentalID)
	}

	return
}

// CheckoutPhaseOneInstance takes a accidentalDB that has been found in the DB, updates the backRepo and stages the
// models version of the accidentalDB
func (backRepoAccidental *BackRepoAccidentalStruct) CheckoutPhaseOneInstance(accidentalDB *AccidentalDB) (Error error) {

	accidental, ok := backRepoAccidental.Map_AccidentalDBID_AccidentalPtr[accidentalDB.ID]
	if !ok {
		accidental = new(models.Accidental)

		backRepoAccidental.Map_AccidentalDBID_AccidentalPtr[accidentalDB.ID] = accidental
		backRepoAccidental.Map_AccidentalPtr_AccidentalDBID[accidental] = accidentalDB.ID

		// append model store with the new element
		accidental.Name = accidentalDB.Name_Data.String
		accidental.Stage(backRepoAccidental.GetStage())
	}
	accidentalDB.CopyBasicFieldsToAccidental(accidental)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	accidental.Stage(backRepoAccidental.GetStage())

	// preserve pointer to accidentalDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_AccidentalDBID_AccidentalDB)[accidentalDB hold variable pointers
	accidentalDB_Data := *accidentalDB
	preservedPtrToAccidental := &accidentalDB_Data
	backRepoAccidental.Map_AccidentalDBID_AccidentalDB[accidentalDB.ID] = preservedPtrToAccidental

	return
}

// BackRepoAccidental.CheckoutPhaseTwo Checkouts all staged instances of Accidental to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccidental *BackRepoAccidentalStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, accidentalDB := range backRepoAccidental.Map_AccidentalDBID_AccidentalDB {
		backRepoAccidental.CheckoutPhaseTwoInstance(backRepo, accidentalDB)
	}
	return
}

// BackRepoAccidental.CheckoutPhaseTwoInstance Checkouts staged instances of Accidental to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccidental *BackRepoAccidentalStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, accidentalDB *AccidentalDB) (Error error) {

	accidental := backRepoAccidental.Map_AccidentalDBID_AccidentalPtr[accidentalDB.ID]

	accidentalDB.DecodePointers(backRepo, accidental)

	return
}

func (accidentalDB *AccidentalDB) DecodePointers(backRepo *BackRepoStruct, accidental *models.Accidental) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitAccidental allows commit of a single accidental (if already staged)
func (backRepo *BackRepoStruct) CommitAccidental(accidental *models.Accidental) {
	backRepo.BackRepoAccidental.CommitPhaseOneInstance(accidental)
	if id, ok := backRepo.BackRepoAccidental.Map_AccidentalPtr_AccidentalDBID[accidental]; ok {
		backRepo.BackRepoAccidental.CommitPhaseTwoInstance(backRepo, id, accidental)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitAccidental allows checkout of a single accidental (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutAccidental(accidental *models.Accidental) {
	// check if the accidental is staged
	if _, ok := backRepo.BackRepoAccidental.Map_AccidentalPtr_AccidentalDBID[accidental]; ok {

		if id, ok := backRepo.BackRepoAccidental.Map_AccidentalPtr_AccidentalDBID[accidental]; ok {
			var accidentalDB AccidentalDB
			accidentalDB.ID = id

			if err := backRepo.BackRepoAccidental.db.First(&accidentalDB, id).Error; err != nil {
				log.Fatalln("CheckoutAccidental : Problem with getting object with id:", id)
			}
			backRepo.BackRepoAccidental.CheckoutPhaseOneInstance(&accidentalDB)
			backRepo.BackRepoAccidental.CheckoutPhaseTwoInstance(backRepo, &accidentalDB)
		}
	}
}

// CopyBasicFieldsFromAccidental
func (accidentalDB *AccidentalDB) CopyBasicFieldsFromAccidental(accidental *models.Accidental) {
	// insertion point for fields commit

	accidentalDB.Name_Data.String = accidental.Name
	accidentalDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromAccidental_WOP
func (accidentalDB *AccidentalDB) CopyBasicFieldsFromAccidental_WOP(accidental *models.Accidental_WOP) {
	// insertion point for fields commit

	accidentalDB.Name_Data.String = accidental.Name
	accidentalDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromAccidentalWOP
func (accidentalDB *AccidentalDB) CopyBasicFieldsFromAccidentalWOP(accidental *AccidentalWOP) {
	// insertion point for fields commit

	accidentalDB.Name_Data.String = accidental.Name
	accidentalDB.Name_Data.Valid = true
}

// CopyBasicFieldsToAccidental
func (accidentalDB *AccidentalDB) CopyBasicFieldsToAccidental(accidental *models.Accidental) {
	// insertion point for checkout of basic fields (back repo to stage)
	accidental.Name = accidentalDB.Name_Data.String
}

// CopyBasicFieldsToAccidental_WOP
func (accidentalDB *AccidentalDB) CopyBasicFieldsToAccidental_WOP(accidental *models.Accidental_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	accidental.Name = accidentalDB.Name_Data.String
}

// CopyBasicFieldsToAccidentalWOP
func (accidentalDB *AccidentalDB) CopyBasicFieldsToAccidentalWOP(accidental *AccidentalWOP) {
	accidental.ID = int(accidentalDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	accidental.Name = accidentalDB.Name_Data.String
}

// Backup generates a json file from a slice of all AccidentalDB instances in the backrepo
func (backRepoAccidental *BackRepoAccidentalStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "AccidentalDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AccidentalDB, 0)
	for _, accidentalDB := range backRepoAccidental.Map_AccidentalDBID_AccidentalDB {
		forBackup = append(forBackup, accidentalDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Accidental ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Accidental file", err.Error())
	}
}

// Backup generates a json file from a slice of all AccidentalDB instances in the backrepo
func (backRepoAccidental *BackRepoAccidentalStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AccidentalDB, 0)
	for _, accidentalDB := range backRepoAccidental.Map_AccidentalDBID_AccidentalDB {
		forBackup = append(forBackup, accidentalDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Accidental")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Accidental_Fields, -1)
	for _, accidentalDB := range forBackup {

		var accidentalWOP AccidentalWOP
		accidentalDB.CopyBasicFieldsToAccidentalWOP(&accidentalWOP)

		row := sh.AddRow()
		row.WriteStruct(&accidentalWOP, -1)
	}
}

// RestoreXL from the "Accidental" sheet all AccidentalDB instances
func (backRepoAccidental *BackRepoAccidentalStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoAccidentalid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Accidental"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoAccidental.rowVisitorAccidental)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoAccidental *BackRepoAccidentalStruct) rowVisitorAccidental(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var accidentalWOP AccidentalWOP
		row.ReadStruct(&accidentalWOP)

		// add the unmarshalled struct to the stage
		accidentalDB := new(AccidentalDB)
		accidentalDB.CopyBasicFieldsFromAccidentalWOP(&accidentalWOP)

		accidentalDB_ID_atBackupTime := accidentalDB.ID
		accidentalDB.ID = 0
		query := backRepoAccidental.db.Create(accidentalDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoAccidental.Map_AccidentalDBID_AccidentalDB[accidentalDB.ID] = accidentalDB
		BackRepoAccidentalid_atBckpTime_newID[accidentalDB_ID_atBackupTime] = accidentalDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "AccidentalDB.json" in dirPath that stores an array
// of AccidentalDB and stores it in the database
// the map BackRepoAccidentalid_atBckpTime_newID is updated accordingly
func (backRepoAccidental *BackRepoAccidentalStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoAccidentalid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "AccidentalDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Accidental file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*AccidentalDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_AccidentalDBID_AccidentalDB
	for _, accidentalDB := range forRestore {

		accidentalDB_ID_atBackupTime := accidentalDB.ID
		accidentalDB.ID = 0
		query := backRepoAccidental.db.Create(accidentalDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoAccidental.Map_AccidentalDBID_AccidentalDB[accidentalDB.ID] = accidentalDB
		BackRepoAccidentalid_atBckpTime_newID[accidentalDB_ID_atBackupTime] = accidentalDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Accidental file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Accidental>id_atBckpTime_newID
// to compute new index
func (backRepoAccidental *BackRepoAccidentalStruct) RestorePhaseTwo() {

	for _, accidentalDB := range backRepoAccidental.Map_AccidentalDBID_AccidentalDB {

		// next line of code is to avert unused variable compilation error
		_ = accidentalDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoAccidental.db.Model(accidentalDB).Updates(*accidentalDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoAccidental.ResetReversePointers commits all staged instances of Accidental to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccidental *BackRepoAccidentalStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, accidental := range backRepoAccidental.Map_AccidentalDBID_AccidentalPtr {
		backRepoAccidental.ResetReversePointersInstance(backRepo, idx, accidental)
	}

	return
}

func (backRepoAccidental *BackRepoAccidentalStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, accidental *models.Accidental) (Error error) {

	// fetch matching accidentalDB
	if accidentalDB, ok := backRepoAccidental.Map_AccidentalDBID_AccidentalDB[idx]; ok {
		_ = accidentalDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoAccidentalid_atBckpTime_newID map[uint]uint
