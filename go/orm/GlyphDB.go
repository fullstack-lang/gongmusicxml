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
var dummy_Glyph_sql sql.NullBool
var dummy_Glyph_time time.Duration
var dummy_Glyph_sort sort.Float64Slice

// GlyphAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model glyphAPI
type GlyphAPI struct {
	gorm.Model

	models.Glyph_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	GlyphPointersEncoding GlyphPointersEncoding
}

// GlyphPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GlyphPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// GlyphDB describes a glyph in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model glyphDB
type GlyphDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field glyphDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	GlyphPointersEncoding
}

// GlyphDBs arrays glyphDBs
// swagger:response glyphDBsResponse
type GlyphDBs []GlyphDB

// GlyphDBResponse provides response
// swagger:response glyphDBResponse
type GlyphDBResponse struct {
	GlyphDB
}

// GlyphWOP is a Glyph without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GlyphWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Glyph_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoGlyphStruct struct {
	// stores GlyphDB according to their gorm ID
	Map_GlyphDBID_GlyphDB map[uint]*GlyphDB

	// stores GlyphDB ID according to Glyph address
	Map_GlyphPtr_GlyphDBID map[*models.Glyph]uint

	// stores Glyph according to their gorm ID
	Map_GlyphDBID_GlyphPtr map[uint]*models.Glyph

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoGlyph *BackRepoGlyphStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoGlyph.stage
	return
}

func (backRepoGlyph *BackRepoGlyphStruct) GetDB() *gorm.DB {
	return backRepoGlyph.db
}

// GetGlyphDBFromGlyphPtr is a handy function to access the back repo instance from the stage instance
func (backRepoGlyph *BackRepoGlyphStruct) GetGlyphDBFromGlyphPtr(glyph *models.Glyph) (glyphDB *GlyphDB) {
	id := backRepoGlyph.Map_GlyphPtr_GlyphDBID[glyph]
	glyphDB = backRepoGlyph.Map_GlyphDBID_GlyphDB[id]
	return
}

// BackRepoGlyph.CommitPhaseOne commits all staged instances of Glyph to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGlyph *BackRepoGlyphStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for glyph := range stage.Glyphs {
		backRepoGlyph.CommitPhaseOneInstance(glyph)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, glyph := range backRepoGlyph.Map_GlyphDBID_GlyphPtr {
		if _, ok := stage.Glyphs[glyph]; !ok {
			backRepoGlyph.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGlyph.CommitDeleteInstance commits deletion of Glyph to the BackRepo
func (backRepoGlyph *BackRepoGlyphStruct) CommitDeleteInstance(id uint) (Error error) {

	glyph := backRepoGlyph.Map_GlyphDBID_GlyphPtr[id]

	// glyph is not staged anymore, remove glyphDB
	glyphDB := backRepoGlyph.Map_GlyphDBID_GlyphDB[id]
	query := backRepoGlyph.db.Unscoped().Delete(&glyphDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoGlyph.Map_GlyphPtr_GlyphDBID, glyph)
	delete(backRepoGlyph.Map_GlyphDBID_GlyphPtr, id)
	delete(backRepoGlyph.Map_GlyphDBID_GlyphDB, id)

	return
}

// BackRepoGlyph.CommitPhaseOneInstance commits glyph staged instances of Glyph to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGlyph *BackRepoGlyphStruct) CommitPhaseOneInstance(glyph *models.Glyph) (Error error) {

	// check if the glyph is not commited yet
	if _, ok := backRepoGlyph.Map_GlyphPtr_GlyphDBID[glyph]; ok {
		return
	}

	// initiate glyph
	var glyphDB GlyphDB
	glyphDB.CopyBasicFieldsFromGlyph(glyph)

	query := backRepoGlyph.db.Create(&glyphDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoGlyph.Map_GlyphPtr_GlyphDBID[glyph] = glyphDB.ID
	backRepoGlyph.Map_GlyphDBID_GlyphPtr[glyphDB.ID] = glyph
	backRepoGlyph.Map_GlyphDBID_GlyphDB[glyphDB.ID] = &glyphDB

	return
}

// BackRepoGlyph.CommitPhaseTwo commits all staged instances of Glyph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGlyph *BackRepoGlyphStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, glyph := range backRepoGlyph.Map_GlyphDBID_GlyphPtr {
		backRepoGlyph.CommitPhaseTwoInstance(backRepo, idx, glyph)
	}

	return
}

// BackRepoGlyph.CommitPhaseTwoInstance commits {{structname }} of models.Glyph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGlyph *BackRepoGlyphStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, glyph *models.Glyph) (Error error) {

	// fetch matching glyphDB
	if glyphDB, ok := backRepoGlyph.Map_GlyphDBID_GlyphDB[idx]; ok {

		glyphDB.CopyBasicFieldsFromGlyph(glyph)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoGlyph.db.Save(&glyphDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Glyph intance %s", glyph.Name))
		return err
	}

	return
}

// BackRepoGlyph.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoGlyph *BackRepoGlyphStruct) CheckoutPhaseOne() (Error error) {

	glyphDBArray := make([]GlyphDB, 0)
	query := backRepoGlyph.db.Find(&glyphDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	glyphInstancesToBeRemovedFromTheStage := make(map[*models.Glyph]any)
	for key, value := range backRepoGlyph.stage.Glyphs {
		glyphInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, glyphDB := range glyphDBArray {
		backRepoGlyph.CheckoutPhaseOneInstance(&glyphDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		glyph, ok := backRepoGlyph.Map_GlyphDBID_GlyphPtr[glyphDB.ID]
		if ok {
			delete(glyphInstancesToBeRemovedFromTheStage, glyph)
		}
	}

	// remove from stage and back repo's 3 maps all glyphs that are not in the checkout
	for glyph := range glyphInstancesToBeRemovedFromTheStage {
		glyph.Unstage(backRepoGlyph.GetStage())

		// remove instance from the back repo 3 maps
		glyphID := backRepoGlyph.Map_GlyphPtr_GlyphDBID[glyph]
		delete(backRepoGlyph.Map_GlyphPtr_GlyphDBID, glyph)
		delete(backRepoGlyph.Map_GlyphDBID_GlyphDB, glyphID)
		delete(backRepoGlyph.Map_GlyphDBID_GlyphPtr, glyphID)
	}

	return
}

// CheckoutPhaseOneInstance takes a glyphDB that has been found in the DB, updates the backRepo and stages the
// models version of the glyphDB
func (backRepoGlyph *BackRepoGlyphStruct) CheckoutPhaseOneInstance(glyphDB *GlyphDB) (Error error) {

	glyph, ok := backRepoGlyph.Map_GlyphDBID_GlyphPtr[glyphDB.ID]
	if !ok {
		glyph = new(models.Glyph)

		backRepoGlyph.Map_GlyphDBID_GlyphPtr[glyphDB.ID] = glyph
		backRepoGlyph.Map_GlyphPtr_GlyphDBID[glyph] = glyphDB.ID

		// append model store with the new element
		glyph.Name = glyphDB.Name_Data.String
		glyph.Stage(backRepoGlyph.GetStage())
	}
	glyphDB.CopyBasicFieldsToGlyph(glyph)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	glyph.Stage(backRepoGlyph.GetStage())

	// preserve pointer to glyphDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GlyphDBID_GlyphDB)[glyphDB hold variable pointers
	glyphDB_Data := *glyphDB
	preservedPtrToGlyph := &glyphDB_Data
	backRepoGlyph.Map_GlyphDBID_GlyphDB[glyphDB.ID] = preservedPtrToGlyph

	return
}

// BackRepoGlyph.CheckoutPhaseTwo Checkouts all staged instances of Glyph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGlyph *BackRepoGlyphStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, glyphDB := range backRepoGlyph.Map_GlyphDBID_GlyphDB {
		backRepoGlyph.CheckoutPhaseTwoInstance(backRepo, glyphDB)
	}
	return
}

// BackRepoGlyph.CheckoutPhaseTwoInstance Checkouts staged instances of Glyph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGlyph *BackRepoGlyphStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, glyphDB *GlyphDB) (Error error) {

	glyph := backRepoGlyph.Map_GlyphDBID_GlyphPtr[glyphDB.ID]

	glyphDB.DecodePointers(backRepo, glyph)

	return
}

func (glyphDB *GlyphDB) DecodePointers(backRepo *BackRepoStruct, glyph *models.Glyph) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitGlyph allows commit of a single glyph (if already staged)
func (backRepo *BackRepoStruct) CommitGlyph(glyph *models.Glyph) {
	backRepo.BackRepoGlyph.CommitPhaseOneInstance(glyph)
	if id, ok := backRepo.BackRepoGlyph.Map_GlyphPtr_GlyphDBID[glyph]; ok {
		backRepo.BackRepoGlyph.CommitPhaseTwoInstance(backRepo, id, glyph)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitGlyph allows checkout of a single glyph (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGlyph(glyph *models.Glyph) {
	// check if the glyph is staged
	if _, ok := backRepo.BackRepoGlyph.Map_GlyphPtr_GlyphDBID[glyph]; ok {

		if id, ok := backRepo.BackRepoGlyph.Map_GlyphPtr_GlyphDBID[glyph]; ok {
			var glyphDB GlyphDB
			glyphDB.ID = id

			if err := backRepo.BackRepoGlyph.db.First(&glyphDB, id).Error; err != nil {
				log.Fatalln("CheckoutGlyph : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGlyph.CheckoutPhaseOneInstance(&glyphDB)
			backRepo.BackRepoGlyph.CheckoutPhaseTwoInstance(backRepo, &glyphDB)
		}
	}
}

// CopyBasicFieldsFromGlyph
func (glyphDB *GlyphDB) CopyBasicFieldsFromGlyph(glyph *models.Glyph) {
	// insertion point for fields commit

	glyphDB.Name_Data.String = glyph.Name
	glyphDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromGlyph_WOP
func (glyphDB *GlyphDB) CopyBasicFieldsFromGlyph_WOP(glyph *models.Glyph_WOP) {
	// insertion point for fields commit

	glyphDB.Name_Data.String = glyph.Name
	glyphDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromGlyphWOP
func (glyphDB *GlyphDB) CopyBasicFieldsFromGlyphWOP(glyph *GlyphWOP) {
	// insertion point for fields commit

	glyphDB.Name_Data.String = glyph.Name
	glyphDB.Name_Data.Valid = true
}

// CopyBasicFieldsToGlyph
func (glyphDB *GlyphDB) CopyBasicFieldsToGlyph(glyph *models.Glyph) {
	// insertion point for checkout of basic fields (back repo to stage)
	glyph.Name = glyphDB.Name_Data.String
}

// CopyBasicFieldsToGlyph_WOP
func (glyphDB *GlyphDB) CopyBasicFieldsToGlyph_WOP(glyph *models.Glyph_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	glyph.Name = glyphDB.Name_Data.String
}

// CopyBasicFieldsToGlyphWOP
func (glyphDB *GlyphDB) CopyBasicFieldsToGlyphWOP(glyph *GlyphWOP) {
	glyph.ID = int(glyphDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	glyph.Name = glyphDB.Name_Data.String
}

// Backup generates a json file from a slice of all GlyphDB instances in the backrepo
func (backRepoGlyph *BackRepoGlyphStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GlyphDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GlyphDB, 0)
	for _, glyphDB := range backRepoGlyph.Map_GlyphDBID_GlyphDB {
		forBackup = append(forBackup, glyphDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Glyph ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Glyph file", err.Error())
	}
}

// Backup generates a json file from a slice of all GlyphDB instances in the backrepo
func (backRepoGlyph *BackRepoGlyphStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GlyphDB, 0)
	for _, glyphDB := range backRepoGlyph.Map_GlyphDBID_GlyphDB {
		forBackup = append(forBackup, glyphDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Glyph")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Glyph_Fields, -1)
	for _, glyphDB := range forBackup {

		var glyphWOP GlyphWOP
		glyphDB.CopyBasicFieldsToGlyphWOP(&glyphWOP)

		row := sh.AddRow()
		row.WriteStruct(&glyphWOP, -1)
	}
}

// RestoreXL from the "Glyph" sheet all GlyphDB instances
func (backRepoGlyph *BackRepoGlyphStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGlyphid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Glyph"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGlyph.rowVisitorGlyph)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoGlyph *BackRepoGlyphStruct) rowVisitorGlyph(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var glyphWOP GlyphWOP
		row.ReadStruct(&glyphWOP)

		// add the unmarshalled struct to the stage
		glyphDB := new(GlyphDB)
		glyphDB.CopyBasicFieldsFromGlyphWOP(&glyphWOP)

		glyphDB_ID_atBackupTime := glyphDB.ID
		glyphDB.ID = 0
		query := backRepoGlyph.db.Create(glyphDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoGlyph.Map_GlyphDBID_GlyphDB[glyphDB.ID] = glyphDB
		BackRepoGlyphid_atBckpTime_newID[glyphDB_ID_atBackupTime] = glyphDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GlyphDB.json" in dirPath that stores an array
// of GlyphDB and stores it in the database
// the map BackRepoGlyphid_atBckpTime_newID is updated accordingly
func (backRepoGlyph *BackRepoGlyphStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGlyphid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GlyphDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Glyph file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GlyphDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GlyphDBID_GlyphDB
	for _, glyphDB := range forRestore {

		glyphDB_ID_atBackupTime := glyphDB.ID
		glyphDB.ID = 0
		query := backRepoGlyph.db.Create(glyphDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoGlyph.Map_GlyphDBID_GlyphDB[glyphDB.ID] = glyphDB
		BackRepoGlyphid_atBckpTime_newID[glyphDB_ID_atBackupTime] = glyphDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Glyph file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Glyph>id_atBckpTime_newID
// to compute new index
func (backRepoGlyph *BackRepoGlyphStruct) RestorePhaseTwo() {

	for _, glyphDB := range backRepoGlyph.Map_GlyphDBID_GlyphDB {

		// next line of code is to avert unused variable compilation error
		_ = glyphDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoGlyph.db.Model(glyphDB).Updates(*glyphDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoGlyph.ResetReversePointers commits all staged instances of Glyph to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGlyph *BackRepoGlyphStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, glyph := range backRepoGlyph.Map_GlyphDBID_GlyphPtr {
		backRepoGlyph.ResetReversePointersInstance(backRepo, idx, glyph)
	}

	return
}

func (backRepoGlyph *BackRepoGlyphStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, glyph *models.Glyph) (Error error) {

	// fetch matching glyphDB
	if glyphDB, ok := backRepoGlyph.Map_GlyphDBID_GlyphDB[idx]; ok {
		_ = glyphDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGlyphid_atBckpTime_newID map[uint]uint
