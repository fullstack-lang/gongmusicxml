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
var dummy_Accord_sql sql.NullBool
var dummy_Accord_time time.Duration
var dummy_Accord_sort sort.Float64Slice

// AccordAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model accordAPI
type AccordAPI struct {
	gorm.Model

	models.Accord_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	AccordPointersEncoding AccordPointersEncoding
}

// AccordPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type AccordPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// AccordDB describes a accord in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model accordDB
type AccordDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field accordDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	AccordPointersEncoding
}

// AccordDBs arrays accordDBs
// swagger:response accordDBsResponse
type AccordDBs []AccordDB

// AccordDBResponse provides response
// swagger:response accordDBResponse
type AccordDBResponse struct {
	AccordDB
}

// AccordWOP is a Accord without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type AccordWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Accord_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoAccordStruct struct {
	// stores AccordDB according to their gorm ID
	Map_AccordDBID_AccordDB map[uint]*AccordDB

	// stores AccordDB ID according to Accord address
	Map_AccordPtr_AccordDBID map[*models.Accord]uint

	// stores Accord according to their gorm ID
	Map_AccordDBID_AccordPtr map[uint]*models.Accord

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoAccord *BackRepoAccordStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoAccord.stage
	return
}

func (backRepoAccord *BackRepoAccordStruct) GetDB() *gorm.DB {
	return backRepoAccord.db
}

// GetAccordDBFromAccordPtr is a handy function to access the back repo instance from the stage instance
func (backRepoAccord *BackRepoAccordStruct) GetAccordDBFromAccordPtr(accord *models.Accord) (accordDB *AccordDB) {
	id := backRepoAccord.Map_AccordPtr_AccordDBID[accord]
	accordDB = backRepoAccord.Map_AccordDBID_AccordDB[id]
	return
}

// BackRepoAccord.CommitPhaseOne commits all staged instances of Accord to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAccord *BackRepoAccordStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for accord := range stage.Accords {
		backRepoAccord.CommitPhaseOneInstance(accord)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, accord := range backRepoAccord.Map_AccordDBID_AccordPtr {
		if _, ok := stage.Accords[accord]; !ok {
			backRepoAccord.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoAccord.CommitDeleteInstance commits deletion of Accord to the BackRepo
func (backRepoAccord *BackRepoAccordStruct) CommitDeleteInstance(id uint) (Error error) {

	accord := backRepoAccord.Map_AccordDBID_AccordPtr[id]

	// accord is not staged anymore, remove accordDB
	accordDB := backRepoAccord.Map_AccordDBID_AccordDB[id]
	query := backRepoAccord.db.Unscoped().Delete(&accordDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoAccord.Map_AccordPtr_AccordDBID, accord)
	delete(backRepoAccord.Map_AccordDBID_AccordPtr, id)
	delete(backRepoAccord.Map_AccordDBID_AccordDB, id)

	return
}

// BackRepoAccord.CommitPhaseOneInstance commits accord staged instances of Accord to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAccord *BackRepoAccordStruct) CommitPhaseOneInstance(accord *models.Accord) (Error error) {

	// check if the accord is not commited yet
	if _, ok := backRepoAccord.Map_AccordPtr_AccordDBID[accord]; ok {
		return
	}

	// initiate accord
	var accordDB AccordDB
	accordDB.CopyBasicFieldsFromAccord(accord)

	query := backRepoAccord.db.Create(&accordDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoAccord.Map_AccordPtr_AccordDBID[accord] = accordDB.ID
	backRepoAccord.Map_AccordDBID_AccordPtr[accordDB.ID] = accord
	backRepoAccord.Map_AccordDBID_AccordDB[accordDB.ID] = &accordDB

	return
}

// BackRepoAccord.CommitPhaseTwo commits all staged instances of Accord to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccord *BackRepoAccordStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, accord := range backRepoAccord.Map_AccordDBID_AccordPtr {
		backRepoAccord.CommitPhaseTwoInstance(backRepo, idx, accord)
	}

	return
}

// BackRepoAccord.CommitPhaseTwoInstance commits {{structname }} of models.Accord to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccord *BackRepoAccordStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, accord *models.Accord) (Error error) {

	// fetch matching accordDB
	if accordDB, ok := backRepoAccord.Map_AccordDBID_AccordDB[idx]; ok {

		accordDB.CopyBasicFieldsFromAccord(accord)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoAccord.db.Save(&accordDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Accord intance %s", accord.Name))
		return err
	}

	return
}

// BackRepoAccord.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoAccord *BackRepoAccordStruct) CheckoutPhaseOne() (Error error) {

	accordDBArray := make([]AccordDB, 0)
	query := backRepoAccord.db.Find(&accordDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	accordInstancesToBeRemovedFromTheStage := make(map[*models.Accord]any)
	for key, value := range backRepoAccord.stage.Accords {
		accordInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, accordDB := range accordDBArray {
		backRepoAccord.CheckoutPhaseOneInstance(&accordDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		accord, ok := backRepoAccord.Map_AccordDBID_AccordPtr[accordDB.ID]
		if ok {
			delete(accordInstancesToBeRemovedFromTheStage, accord)
		}
	}

	// remove from stage and back repo's 3 maps all accords that are not in the checkout
	for accord := range accordInstancesToBeRemovedFromTheStage {
		accord.Unstage(backRepoAccord.GetStage())

		// remove instance from the back repo 3 maps
		accordID := backRepoAccord.Map_AccordPtr_AccordDBID[accord]
		delete(backRepoAccord.Map_AccordPtr_AccordDBID, accord)
		delete(backRepoAccord.Map_AccordDBID_AccordDB, accordID)
		delete(backRepoAccord.Map_AccordDBID_AccordPtr, accordID)
	}

	return
}

// CheckoutPhaseOneInstance takes a accordDB that has been found in the DB, updates the backRepo and stages the
// models version of the accordDB
func (backRepoAccord *BackRepoAccordStruct) CheckoutPhaseOneInstance(accordDB *AccordDB) (Error error) {

	accord, ok := backRepoAccord.Map_AccordDBID_AccordPtr[accordDB.ID]
	if !ok {
		accord = new(models.Accord)

		backRepoAccord.Map_AccordDBID_AccordPtr[accordDB.ID] = accord
		backRepoAccord.Map_AccordPtr_AccordDBID[accord] = accordDB.ID

		// append model store with the new element
		accord.Name = accordDB.Name_Data.String
		accord.Stage(backRepoAccord.GetStage())
	}
	accordDB.CopyBasicFieldsToAccord(accord)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	accord.Stage(backRepoAccord.GetStage())

	// preserve pointer to accordDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_AccordDBID_AccordDB)[accordDB hold variable pointers
	accordDB_Data := *accordDB
	preservedPtrToAccord := &accordDB_Data
	backRepoAccord.Map_AccordDBID_AccordDB[accordDB.ID] = preservedPtrToAccord

	return
}

// BackRepoAccord.CheckoutPhaseTwo Checkouts all staged instances of Accord to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccord *BackRepoAccordStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, accordDB := range backRepoAccord.Map_AccordDBID_AccordDB {
		backRepoAccord.CheckoutPhaseTwoInstance(backRepo, accordDB)
	}
	return
}

// BackRepoAccord.CheckoutPhaseTwoInstance Checkouts staged instances of Accord to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccord *BackRepoAccordStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, accordDB *AccordDB) (Error error) {

	accord := backRepoAccord.Map_AccordDBID_AccordPtr[accordDB.ID]

	accordDB.DecodePointers(backRepo, accord)

	return
}

func (accordDB *AccordDB) DecodePointers(backRepo *BackRepoStruct, accord *models.Accord) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitAccord allows commit of a single accord (if already staged)
func (backRepo *BackRepoStruct) CommitAccord(accord *models.Accord) {
	backRepo.BackRepoAccord.CommitPhaseOneInstance(accord)
	if id, ok := backRepo.BackRepoAccord.Map_AccordPtr_AccordDBID[accord]; ok {
		backRepo.BackRepoAccord.CommitPhaseTwoInstance(backRepo, id, accord)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitAccord allows checkout of a single accord (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutAccord(accord *models.Accord) {
	// check if the accord is staged
	if _, ok := backRepo.BackRepoAccord.Map_AccordPtr_AccordDBID[accord]; ok {

		if id, ok := backRepo.BackRepoAccord.Map_AccordPtr_AccordDBID[accord]; ok {
			var accordDB AccordDB
			accordDB.ID = id

			if err := backRepo.BackRepoAccord.db.First(&accordDB, id).Error; err != nil {
				log.Fatalln("CheckoutAccord : Problem with getting object with id:", id)
			}
			backRepo.BackRepoAccord.CheckoutPhaseOneInstance(&accordDB)
			backRepo.BackRepoAccord.CheckoutPhaseTwoInstance(backRepo, &accordDB)
		}
	}
}

// CopyBasicFieldsFromAccord
func (accordDB *AccordDB) CopyBasicFieldsFromAccord(accord *models.Accord) {
	// insertion point for fields commit

	accordDB.Name_Data.String = accord.Name
	accordDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromAccord_WOP
func (accordDB *AccordDB) CopyBasicFieldsFromAccord_WOP(accord *models.Accord_WOP) {
	// insertion point for fields commit

	accordDB.Name_Data.String = accord.Name
	accordDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromAccordWOP
func (accordDB *AccordDB) CopyBasicFieldsFromAccordWOP(accord *AccordWOP) {
	// insertion point for fields commit

	accordDB.Name_Data.String = accord.Name
	accordDB.Name_Data.Valid = true
}

// CopyBasicFieldsToAccord
func (accordDB *AccordDB) CopyBasicFieldsToAccord(accord *models.Accord) {
	// insertion point for checkout of basic fields (back repo to stage)
	accord.Name = accordDB.Name_Data.String
}

// CopyBasicFieldsToAccord_WOP
func (accordDB *AccordDB) CopyBasicFieldsToAccord_WOP(accord *models.Accord_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	accord.Name = accordDB.Name_Data.String
}

// CopyBasicFieldsToAccordWOP
func (accordDB *AccordDB) CopyBasicFieldsToAccordWOP(accord *AccordWOP) {
	accord.ID = int(accordDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	accord.Name = accordDB.Name_Data.String
}

// Backup generates a json file from a slice of all AccordDB instances in the backrepo
func (backRepoAccord *BackRepoAccordStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "AccordDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AccordDB, 0)
	for _, accordDB := range backRepoAccord.Map_AccordDBID_AccordDB {
		forBackup = append(forBackup, accordDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Accord ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Accord file", err.Error())
	}
}

// Backup generates a json file from a slice of all AccordDB instances in the backrepo
func (backRepoAccord *BackRepoAccordStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AccordDB, 0)
	for _, accordDB := range backRepoAccord.Map_AccordDBID_AccordDB {
		forBackup = append(forBackup, accordDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Accord")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Accord_Fields, -1)
	for _, accordDB := range forBackup {

		var accordWOP AccordWOP
		accordDB.CopyBasicFieldsToAccordWOP(&accordWOP)

		row := sh.AddRow()
		row.WriteStruct(&accordWOP, -1)
	}
}

// RestoreXL from the "Accord" sheet all AccordDB instances
func (backRepoAccord *BackRepoAccordStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoAccordid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Accord"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoAccord.rowVisitorAccord)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoAccord *BackRepoAccordStruct) rowVisitorAccord(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var accordWOP AccordWOP
		row.ReadStruct(&accordWOP)

		// add the unmarshalled struct to the stage
		accordDB := new(AccordDB)
		accordDB.CopyBasicFieldsFromAccordWOP(&accordWOP)

		accordDB_ID_atBackupTime := accordDB.ID
		accordDB.ID = 0
		query := backRepoAccord.db.Create(accordDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoAccord.Map_AccordDBID_AccordDB[accordDB.ID] = accordDB
		BackRepoAccordid_atBckpTime_newID[accordDB_ID_atBackupTime] = accordDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "AccordDB.json" in dirPath that stores an array
// of AccordDB and stores it in the database
// the map BackRepoAccordid_atBckpTime_newID is updated accordingly
func (backRepoAccord *BackRepoAccordStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoAccordid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "AccordDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Accord file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*AccordDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_AccordDBID_AccordDB
	for _, accordDB := range forRestore {

		accordDB_ID_atBackupTime := accordDB.ID
		accordDB.ID = 0
		query := backRepoAccord.db.Create(accordDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoAccord.Map_AccordDBID_AccordDB[accordDB.ID] = accordDB
		BackRepoAccordid_atBckpTime_newID[accordDB_ID_atBackupTime] = accordDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Accord file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Accord>id_atBckpTime_newID
// to compute new index
func (backRepoAccord *BackRepoAccordStruct) RestorePhaseTwo() {

	for _, accordDB := range backRepoAccord.Map_AccordDBID_AccordDB {

		// next line of code is to avert unused variable compilation error
		_ = accordDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoAccord.db.Model(accordDB).Updates(*accordDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoAccord.ResetReversePointers commits all staged instances of Accord to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAccord *BackRepoAccordStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, accord := range backRepoAccord.Map_AccordDBID_AccordPtr {
		backRepoAccord.ResetReversePointersInstance(backRepo, idx, accord)
	}

	return
}

func (backRepoAccord *BackRepoAccordStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, accord *models.Accord) (Error error) {

	// fetch matching accordDB
	if accordDB, ok := backRepoAccord.Map_AccordDBID_AccordDB[idx]; ok {
		_ = accordDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoAccordid_atBckpTime_newID map[uint]uint
