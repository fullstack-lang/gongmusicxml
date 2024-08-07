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
var dummy_Beam_sql sql.NullBool
var dummy_Beam_time time.Duration
var dummy_Beam_sort sort.Float64Slice

// BeamAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model beamAPI
type BeamAPI struct {
	gorm.Model

	models.Beam_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	BeamPointersEncoding BeamPointersEncoding
}

// BeamPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type BeamPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// BeamDB describes a beam in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model beamDB
type BeamDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field beamDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	BeamPointersEncoding
}

// BeamDBs arrays beamDBs
// swagger:response beamDBsResponse
type BeamDBs []BeamDB

// BeamDBResponse provides response
// swagger:response beamDBResponse
type BeamDBResponse struct {
	BeamDB
}

// BeamWOP is a Beam without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type BeamWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Beam_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoBeamStruct struct {
	// stores BeamDB according to their gorm ID
	Map_BeamDBID_BeamDB map[uint]*BeamDB

	// stores BeamDB ID according to Beam address
	Map_BeamPtr_BeamDBID map[*models.Beam]uint

	// stores Beam according to their gorm ID
	Map_BeamDBID_BeamPtr map[uint]*models.Beam

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoBeam *BackRepoBeamStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoBeam.stage
	return
}

func (backRepoBeam *BackRepoBeamStruct) GetDB() *gorm.DB {
	return backRepoBeam.db
}

// GetBeamDBFromBeamPtr is a handy function to access the back repo instance from the stage instance
func (backRepoBeam *BackRepoBeamStruct) GetBeamDBFromBeamPtr(beam *models.Beam) (beamDB *BeamDB) {
	id := backRepoBeam.Map_BeamPtr_BeamDBID[beam]
	beamDB = backRepoBeam.Map_BeamDBID_BeamDB[id]
	return
}

// BackRepoBeam.CommitPhaseOne commits all staged instances of Beam to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoBeam *BackRepoBeamStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for beam := range stage.Beams {
		backRepoBeam.CommitPhaseOneInstance(beam)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, beam := range backRepoBeam.Map_BeamDBID_BeamPtr {
		if _, ok := stage.Beams[beam]; !ok {
			backRepoBeam.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoBeam.CommitDeleteInstance commits deletion of Beam to the BackRepo
func (backRepoBeam *BackRepoBeamStruct) CommitDeleteInstance(id uint) (Error error) {

	beam := backRepoBeam.Map_BeamDBID_BeamPtr[id]

	// beam is not staged anymore, remove beamDB
	beamDB := backRepoBeam.Map_BeamDBID_BeamDB[id]
	query := backRepoBeam.db.Unscoped().Delete(&beamDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoBeam.Map_BeamPtr_BeamDBID, beam)
	delete(backRepoBeam.Map_BeamDBID_BeamPtr, id)
	delete(backRepoBeam.Map_BeamDBID_BeamDB, id)

	return
}

// BackRepoBeam.CommitPhaseOneInstance commits beam staged instances of Beam to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoBeam *BackRepoBeamStruct) CommitPhaseOneInstance(beam *models.Beam) (Error error) {

	// check if the beam is not commited yet
	if _, ok := backRepoBeam.Map_BeamPtr_BeamDBID[beam]; ok {
		return
	}

	// initiate beam
	var beamDB BeamDB
	beamDB.CopyBasicFieldsFromBeam(beam)

	query := backRepoBeam.db.Create(&beamDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoBeam.Map_BeamPtr_BeamDBID[beam] = beamDB.ID
	backRepoBeam.Map_BeamDBID_BeamPtr[beamDB.ID] = beam
	backRepoBeam.Map_BeamDBID_BeamDB[beamDB.ID] = &beamDB

	return
}

// BackRepoBeam.CommitPhaseTwo commits all staged instances of Beam to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBeam *BackRepoBeamStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, beam := range backRepoBeam.Map_BeamDBID_BeamPtr {
		backRepoBeam.CommitPhaseTwoInstance(backRepo, idx, beam)
	}

	return
}

// BackRepoBeam.CommitPhaseTwoInstance commits {{structname }} of models.Beam to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBeam *BackRepoBeamStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, beam *models.Beam) (Error error) {

	// fetch matching beamDB
	if beamDB, ok := backRepoBeam.Map_BeamDBID_BeamDB[idx]; ok {

		beamDB.CopyBasicFieldsFromBeam(beam)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoBeam.db.Save(&beamDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Beam intance %s", beam.Name))
		return err
	}

	return
}

// BackRepoBeam.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoBeam *BackRepoBeamStruct) CheckoutPhaseOne() (Error error) {

	beamDBArray := make([]BeamDB, 0)
	query := backRepoBeam.db.Find(&beamDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	beamInstancesToBeRemovedFromTheStage := make(map[*models.Beam]any)
	for key, value := range backRepoBeam.stage.Beams {
		beamInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, beamDB := range beamDBArray {
		backRepoBeam.CheckoutPhaseOneInstance(&beamDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		beam, ok := backRepoBeam.Map_BeamDBID_BeamPtr[beamDB.ID]
		if ok {
			delete(beamInstancesToBeRemovedFromTheStage, beam)
		}
	}

	// remove from stage and back repo's 3 maps all beams that are not in the checkout
	for beam := range beamInstancesToBeRemovedFromTheStage {
		beam.Unstage(backRepoBeam.GetStage())

		// remove instance from the back repo 3 maps
		beamID := backRepoBeam.Map_BeamPtr_BeamDBID[beam]
		delete(backRepoBeam.Map_BeamPtr_BeamDBID, beam)
		delete(backRepoBeam.Map_BeamDBID_BeamDB, beamID)
		delete(backRepoBeam.Map_BeamDBID_BeamPtr, beamID)
	}

	return
}

// CheckoutPhaseOneInstance takes a beamDB that has been found in the DB, updates the backRepo and stages the
// models version of the beamDB
func (backRepoBeam *BackRepoBeamStruct) CheckoutPhaseOneInstance(beamDB *BeamDB) (Error error) {

	beam, ok := backRepoBeam.Map_BeamDBID_BeamPtr[beamDB.ID]
	if !ok {
		beam = new(models.Beam)

		backRepoBeam.Map_BeamDBID_BeamPtr[beamDB.ID] = beam
		backRepoBeam.Map_BeamPtr_BeamDBID[beam] = beamDB.ID

		// append model store with the new element
		beam.Name = beamDB.Name_Data.String
		beam.Stage(backRepoBeam.GetStage())
	}
	beamDB.CopyBasicFieldsToBeam(beam)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	beam.Stage(backRepoBeam.GetStage())

	// preserve pointer to beamDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_BeamDBID_BeamDB)[beamDB hold variable pointers
	beamDB_Data := *beamDB
	preservedPtrToBeam := &beamDB_Data
	backRepoBeam.Map_BeamDBID_BeamDB[beamDB.ID] = preservedPtrToBeam

	return
}

// BackRepoBeam.CheckoutPhaseTwo Checkouts all staged instances of Beam to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBeam *BackRepoBeamStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, beamDB := range backRepoBeam.Map_BeamDBID_BeamDB {
		backRepoBeam.CheckoutPhaseTwoInstance(backRepo, beamDB)
	}
	return
}

// BackRepoBeam.CheckoutPhaseTwoInstance Checkouts staged instances of Beam to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBeam *BackRepoBeamStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, beamDB *BeamDB) (Error error) {

	beam := backRepoBeam.Map_BeamDBID_BeamPtr[beamDB.ID]

	beamDB.DecodePointers(backRepo, beam)

	return
}

func (beamDB *BeamDB) DecodePointers(backRepo *BackRepoStruct, beam *models.Beam) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitBeam allows commit of a single beam (if already staged)
func (backRepo *BackRepoStruct) CommitBeam(beam *models.Beam) {
	backRepo.BackRepoBeam.CommitPhaseOneInstance(beam)
	if id, ok := backRepo.BackRepoBeam.Map_BeamPtr_BeamDBID[beam]; ok {
		backRepo.BackRepoBeam.CommitPhaseTwoInstance(backRepo, id, beam)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitBeam allows checkout of a single beam (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutBeam(beam *models.Beam) {
	// check if the beam is staged
	if _, ok := backRepo.BackRepoBeam.Map_BeamPtr_BeamDBID[beam]; ok {

		if id, ok := backRepo.BackRepoBeam.Map_BeamPtr_BeamDBID[beam]; ok {
			var beamDB BeamDB
			beamDB.ID = id

			if err := backRepo.BackRepoBeam.db.First(&beamDB, id).Error; err != nil {
				log.Fatalln("CheckoutBeam : Problem with getting object with id:", id)
			}
			backRepo.BackRepoBeam.CheckoutPhaseOneInstance(&beamDB)
			backRepo.BackRepoBeam.CheckoutPhaseTwoInstance(backRepo, &beamDB)
		}
	}
}

// CopyBasicFieldsFromBeam
func (beamDB *BeamDB) CopyBasicFieldsFromBeam(beam *models.Beam) {
	// insertion point for fields commit

	beamDB.Name_Data.String = beam.Name
	beamDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromBeam_WOP
func (beamDB *BeamDB) CopyBasicFieldsFromBeam_WOP(beam *models.Beam_WOP) {
	// insertion point for fields commit

	beamDB.Name_Data.String = beam.Name
	beamDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromBeamWOP
func (beamDB *BeamDB) CopyBasicFieldsFromBeamWOP(beam *BeamWOP) {
	// insertion point for fields commit

	beamDB.Name_Data.String = beam.Name
	beamDB.Name_Data.Valid = true
}

// CopyBasicFieldsToBeam
func (beamDB *BeamDB) CopyBasicFieldsToBeam(beam *models.Beam) {
	// insertion point for checkout of basic fields (back repo to stage)
	beam.Name = beamDB.Name_Data.String
}

// CopyBasicFieldsToBeam_WOP
func (beamDB *BeamDB) CopyBasicFieldsToBeam_WOP(beam *models.Beam_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	beam.Name = beamDB.Name_Data.String
}

// CopyBasicFieldsToBeamWOP
func (beamDB *BeamDB) CopyBasicFieldsToBeamWOP(beam *BeamWOP) {
	beam.ID = int(beamDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	beam.Name = beamDB.Name_Data.String
}

// Backup generates a json file from a slice of all BeamDB instances in the backrepo
func (backRepoBeam *BackRepoBeamStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "BeamDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*BeamDB, 0)
	for _, beamDB := range backRepoBeam.Map_BeamDBID_BeamDB {
		forBackup = append(forBackup, beamDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Beam ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Beam file", err.Error())
	}
}

// Backup generates a json file from a slice of all BeamDB instances in the backrepo
func (backRepoBeam *BackRepoBeamStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*BeamDB, 0)
	for _, beamDB := range backRepoBeam.Map_BeamDBID_BeamDB {
		forBackup = append(forBackup, beamDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Beam")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Beam_Fields, -1)
	for _, beamDB := range forBackup {

		var beamWOP BeamWOP
		beamDB.CopyBasicFieldsToBeamWOP(&beamWOP)

		row := sh.AddRow()
		row.WriteStruct(&beamWOP, -1)
	}
}

// RestoreXL from the "Beam" sheet all BeamDB instances
func (backRepoBeam *BackRepoBeamStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoBeamid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Beam"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoBeam.rowVisitorBeam)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoBeam *BackRepoBeamStruct) rowVisitorBeam(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var beamWOP BeamWOP
		row.ReadStruct(&beamWOP)

		// add the unmarshalled struct to the stage
		beamDB := new(BeamDB)
		beamDB.CopyBasicFieldsFromBeamWOP(&beamWOP)

		beamDB_ID_atBackupTime := beamDB.ID
		beamDB.ID = 0
		query := backRepoBeam.db.Create(beamDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoBeam.Map_BeamDBID_BeamDB[beamDB.ID] = beamDB
		BackRepoBeamid_atBckpTime_newID[beamDB_ID_atBackupTime] = beamDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "BeamDB.json" in dirPath that stores an array
// of BeamDB and stores it in the database
// the map BackRepoBeamid_atBckpTime_newID is updated accordingly
func (backRepoBeam *BackRepoBeamStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoBeamid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "BeamDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Beam file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*BeamDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_BeamDBID_BeamDB
	for _, beamDB := range forRestore {

		beamDB_ID_atBackupTime := beamDB.ID
		beamDB.ID = 0
		query := backRepoBeam.db.Create(beamDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoBeam.Map_BeamDBID_BeamDB[beamDB.ID] = beamDB
		BackRepoBeamid_atBckpTime_newID[beamDB_ID_atBackupTime] = beamDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Beam file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Beam>id_atBckpTime_newID
// to compute new index
func (backRepoBeam *BackRepoBeamStruct) RestorePhaseTwo() {

	for _, beamDB := range backRepoBeam.Map_BeamDBID_BeamDB {

		// next line of code is to avert unused variable compilation error
		_ = beamDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoBeam.db.Model(beamDB).Updates(*beamDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoBeam.ResetReversePointers commits all staged instances of Beam to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoBeam *BackRepoBeamStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, beam := range backRepoBeam.Map_BeamDBID_BeamPtr {
		backRepoBeam.ResetReversePointersInstance(backRepo, idx, beam)
	}

	return
}

func (backRepoBeam *BackRepoBeamStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, beam *models.Beam) (Error error) {

	// fetch matching beamDB
	if beamDB, ok := backRepoBeam.Map_BeamDBID_BeamDB[idx]; ok {
		_ = beamDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoBeamid_atBckpTime_newID map[uint]uint
