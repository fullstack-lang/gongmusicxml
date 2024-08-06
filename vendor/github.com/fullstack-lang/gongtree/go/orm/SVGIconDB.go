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

	"github.com/fullstack-lang/gongtree/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_SVGIcon_sql sql.NullBool
var dummy_SVGIcon_time time.Duration
var dummy_SVGIcon_sort sort.Float64Slice

// SVGIconAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model svgiconAPI
type SVGIconAPI struct {
	gorm.Model

	models.SVGIcon_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SVGIconPointersEncoding SVGIconPointersEncoding
}

// SVGIconPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SVGIconPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// SVGIconDB describes a svgicon in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model svgiconDB
type SVGIconDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field svgiconDB.Name
	Name_Data sql.NullString

	// Declation for basic field svgiconDB.SVG
	SVG_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SVGIconPointersEncoding
}

// SVGIconDBs arrays svgiconDBs
// swagger:response svgiconDBsResponse
type SVGIconDBs []SVGIconDB

// SVGIconDBResponse provides response
// swagger:response svgiconDBResponse
type SVGIconDBResponse struct {
	SVGIconDB
}

// SVGIconWOP is a SVGIcon without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SVGIconWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	SVG string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var SVGIcon_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"SVG",
}

type BackRepoSVGIconStruct struct {
	// stores SVGIconDB according to their gorm ID
	Map_SVGIconDBID_SVGIconDB map[uint]*SVGIconDB

	// stores SVGIconDB ID according to SVGIcon address
	Map_SVGIconPtr_SVGIconDBID map[*models.SVGIcon]uint

	// stores SVGIcon according to their gorm ID
	Map_SVGIconDBID_SVGIconPtr map[uint]*models.SVGIcon

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSVGIcon *BackRepoSVGIconStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSVGIcon.stage
	return
}

func (backRepoSVGIcon *BackRepoSVGIconStruct) GetDB() *gorm.DB {
	return backRepoSVGIcon.db
}

// GetSVGIconDBFromSVGIconPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSVGIcon *BackRepoSVGIconStruct) GetSVGIconDBFromSVGIconPtr(svgicon *models.SVGIcon) (svgiconDB *SVGIconDB) {
	id := backRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID[svgicon]
	svgiconDB = backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB[id]
	return
}

// BackRepoSVGIcon.CommitPhaseOne commits all staged instances of SVGIcon to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSVGIcon *BackRepoSVGIconStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for svgicon := range stage.SVGIcons {
		backRepoSVGIcon.CommitPhaseOneInstance(svgicon)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, svgicon := range backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr {
		if _, ok := stage.SVGIcons[svgicon]; !ok {
			backRepoSVGIcon.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSVGIcon.CommitDeleteInstance commits deletion of SVGIcon to the BackRepo
func (backRepoSVGIcon *BackRepoSVGIconStruct) CommitDeleteInstance(id uint) (Error error) {

	svgicon := backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[id]

	// svgicon is not staged anymore, remove svgiconDB
	svgiconDB := backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB[id]
	query := backRepoSVGIcon.db.Unscoped().Delete(&svgiconDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID, svgicon)
	delete(backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr, id)
	delete(backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB, id)

	return
}

// BackRepoSVGIcon.CommitPhaseOneInstance commits svgicon staged instances of SVGIcon to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSVGIcon *BackRepoSVGIconStruct) CommitPhaseOneInstance(svgicon *models.SVGIcon) (Error error) {

	// check if the svgicon is not commited yet
	if _, ok := backRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID[svgicon]; ok {
		return
	}

	// initiate svgicon
	var svgiconDB SVGIconDB
	svgiconDB.CopyBasicFieldsFromSVGIcon(svgicon)

	query := backRepoSVGIcon.db.Create(&svgiconDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID[svgicon] = svgiconDB.ID
	backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[svgiconDB.ID] = svgicon
	backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB[svgiconDB.ID] = &svgiconDB

	return
}

// BackRepoSVGIcon.CommitPhaseTwo commits all staged instances of SVGIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVGIcon *BackRepoSVGIconStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, svgicon := range backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr {
		backRepoSVGIcon.CommitPhaseTwoInstance(backRepo, idx, svgicon)
	}

	return
}

// BackRepoSVGIcon.CommitPhaseTwoInstance commits {{structname }} of models.SVGIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVGIcon *BackRepoSVGIconStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, svgicon *models.SVGIcon) (Error error) {

	// fetch matching svgiconDB
	if svgiconDB, ok := backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB[idx]; ok {

		svgiconDB.CopyBasicFieldsFromSVGIcon(svgicon)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoSVGIcon.db.Save(&svgiconDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SVGIcon intance %s", svgicon.Name))
		return err
	}

	return
}

// BackRepoSVGIcon.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSVGIcon *BackRepoSVGIconStruct) CheckoutPhaseOne() (Error error) {

	svgiconDBArray := make([]SVGIconDB, 0)
	query := backRepoSVGIcon.db.Find(&svgiconDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	svgiconInstancesToBeRemovedFromTheStage := make(map[*models.SVGIcon]any)
	for key, value := range backRepoSVGIcon.stage.SVGIcons {
		svgiconInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, svgiconDB := range svgiconDBArray {
		backRepoSVGIcon.CheckoutPhaseOneInstance(&svgiconDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		svgicon, ok := backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[svgiconDB.ID]
		if ok {
			delete(svgiconInstancesToBeRemovedFromTheStage, svgicon)
		}
	}

	// remove from stage and back repo's 3 maps all svgicons that are not in the checkout
	for svgicon := range svgiconInstancesToBeRemovedFromTheStage {
		svgicon.Unstage(backRepoSVGIcon.GetStage())

		// remove instance from the back repo 3 maps
		svgiconID := backRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID[svgicon]
		delete(backRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID, svgicon)
		delete(backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB, svgiconID)
		delete(backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr, svgiconID)
	}

	return
}

// CheckoutPhaseOneInstance takes a svgiconDB that has been found in the DB, updates the backRepo and stages the
// models version of the svgiconDB
func (backRepoSVGIcon *BackRepoSVGIconStruct) CheckoutPhaseOneInstance(svgiconDB *SVGIconDB) (Error error) {

	svgicon, ok := backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[svgiconDB.ID]
	if !ok {
		svgicon = new(models.SVGIcon)

		backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[svgiconDB.ID] = svgicon
		backRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID[svgicon] = svgiconDB.ID

		// append model store with the new element
		svgicon.Name = svgiconDB.Name_Data.String
		svgicon.Stage(backRepoSVGIcon.GetStage())
	}
	svgiconDB.CopyBasicFieldsToSVGIcon(svgicon)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	svgicon.Stage(backRepoSVGIcon.GetStage())

	// preserve pointer to svgiconDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SVGIconDBID_SVGIconDB)[svgiconDB hold variable pointers
	svgiconDB_Data := *svgiconDB
	preservedPtrToSVGIcon := &svgiconDB_Data
	backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB[svgiconDB.ID] = preservedPtrToSVGIcon

	return
}

// BackRepoSVGIcon.CheckoutPhaseTwo Checkouts all staged instances of SVGIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVGIcon *BackRepoSVGIconStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, svgiconDB := range backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB {
		backRepoSVGIcon.CheckoutPhaseTwoInstance(backRepo, svgiconDB)
	}
	return
}

// BackRepoSVGIcon.CheckoutPhaseTwoInstance Checkouts staged instances of SVGIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVGIcon *BackRepoSVGIconStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, svgiconDB *SVGIconDB) (Error error) {

	svgicon := backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr[svgiconDB.ID]

	svgiconDB.DecodePointers(backRepo, svgicon)

	return
}

func (svgiconDB *SVGIconDB) DecodePointers(backRepo *BackRepoStruct, svgicon *models.SVGIcon) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitSVGIcon allows commit of a single svgicon (if already staged)
func (backRepo *BackRepoStruct) CommitSVGIcon(svgicon *models.SVGIcon) {
	backRepo.BackRepoSVGIcon.CommitPhaseOneInstance(svgicon)
	if id, ok := backRepo.BackRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID[svgicon]; ok {
		backRepo.BackRepoSVGIcon.CommitPhaseTwoInstance(backRepo, id, svgicon)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSVGIcon allows checkout of a single svgicon (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSVGIcon(svgicon *models.SVGIcon) {
	// check if the svgicon is staged
	if _, ok := backRepo.BackRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID[svgicon]; ok {

		if id, ok := backRepo.BackRepoSVGIcon.Map_SVGIconPtr_SVGIconDBID[svgicon]; ok {
			var svgiconDB SVGIconDB
			svgiconDB.ID = id

			if err := backRepo.BackRepoSVGIcon.db.First(&svgiconDB, id).Error; err != nil {
				log.Fatalln("CheckoutSVGIcon : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSVGIcon.CheckoutPhaseOneInstance(&svgiconDB)
			backRepo.BackRepoSVGIcon.CheckoutPhaseTwoInstance(backRepo, &svgiconDB)
		}
	}
}

// CopyBasicFieldsFromSVGIcon
func (svgiconDB *SVGIconDB) CopyBasicFieldsFromSVGIcon(svgicon *models.SVGIcon) {
	// insertion point for fields commit

	svgiconDB.Name_Data.String = svgicon.Name
	svgiconDB.Name_Data.Valid = true

	svgiconDB.SVG_Data.String = svgicon.SVG
	svgiconDB.SVG_Data.Valid = true
}

// CopyBasicFieldsFromSVGIcon_WOP
func (svgiconDB *SVGIconDB) CopyBasicFieldsFromSVGIcon_WOP(svgicon *models.SVGIcon_WOP) {
	// insertion point for fields commit

	svgiconDB.Name_Data.String = svgicon.Name
	svgiconDB.Name_Data.Valid = true

	svgiconDB.SVG_Data.String = svgicon.SVG
	svgiconDB.SVG_Data.Valid = true
}

// CopyBasicFieldsFromSVGIconWOP
func (svgiconDB *SVGIconDB) CopyBasicFieldsFromSVGIconWOP(svgicon *SVGIconWOP) {
	// insertion point for fields commit

	svgiconDB.Name_Data.String = svgicon.Name
	svgiconDB.Name_Data.Valid = true

	svgiconDB.SVG_Data.String = svgicon.SVG
	svgiconDB.SVG_Data.Valid = true
}

// CopyBasicFieldsToSVGIcon
func (svgiconDB *SVGIconDB) CopyBasicFieldsToSVGIcon(svgicon *models.SVGIcon) {
	// insertion point for checkout of basic fields (back repo to stage)
	svgicon.Name = svgiconDB.Name_Data.String
	svgicon.SVG = svgiconDB.SVG_Data.String
}

// CopyBasicFieldsToSVGIcon_WOP
func (svgiconDB *SVGIconDB) CopyBasicFieldsToSVGIcon_WOP(svgicon *models.SVGIcon_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	svgicon.Name = svgiconDB.Name_Data.String
	svgicon.SVG = svgiconDB.SVG_Data.String
}

// CopyBasicFieldsToSVGIconWOP
func (svgiconDB *SVGIconDB) CopyBasicFieldsToSVGIconWOP(svgicon *SVGIconWOP) {
	svgicon.ID = int(svgiconDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	svgicon.Name = svgiconDB.Name_Data.String
	svgicon.SVG = svgiconDB.SVG_Data.String
}

// Backup generates a json file from a slice of all SVGIconDB instances in the backrepo
func (backRepoSVGIcon *BackRepoSVGIconStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SVGIconDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SVGIconDB, 0)
	for _, svgiconDB := range backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB {
		forBackup = append(forBackup, svgiconDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SVGIcon ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SVGIcon file", err.Error())
	}
}

// Backup generates a json file from a slice of all SVGIconDB instances in the backrepo
func (backRepoSVGIcon *BackRepoSVGIconStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SVGIconDB, 0)
	for _, svgiconDB := range backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB {
		forBackup = append(forBackup, svgiconDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SVGIcon")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SVGIcon_Fields, -1)
	for _, svgiconDB := range forBackup {

		var svgiconWOP SVGIconWOP
		svgiconDB.CopyBasicFieldsToSVGIconWOP(&svgiconWOP)

		row := sh.AddRow()
		row.WriteStruct(&svgiconWOP, -1)
	}
}

// RestoreXL from the "SVGIcon" sheet all SVGIconDB instances
func (backRepoSVGIcon *BackRepoSVGIconStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSVGIconid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SVGIcon"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSVGIcon.rowVisitorSVGIcon)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSVGIcon *BackRepoSVGIconStruct) rowVisitorSVGIcon(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var svgiconWOP SVGIconWOP
		row.ReadStruct(&svgiconWOP)

		// add the unmarshalled struct to the stage
		svgiconDB := new(SVGIconDB)
		svgiconDB.CopyBasicFieldsFromSVGIconWOP(&svgiconWOP)

		svgiconDB_ID_atBackupTime := svgiconDB.ID
		svgiconDB.ID = 0
		query := backRepoSVGIcon.db.Create(svgiconDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB[svgiconDB.ID] = svgiconDB
		BackRepoSVGIconid_atBckpTime_newID[svgiconDB_ID_atBackupTime] = svgiconDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SVGIconDB.json" in dirPath that stores an array
// of SVGIconDB and stores it in the database
// the map BackRepoSVGIconid_atBckpTime_newID is updated accordingly
func (backRepoSVGIcon *BackRepoSVGIconStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSVGIconid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SVGIconDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SVGIcon file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SVGIconDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SVGIconDBID_SVGIconDB
	for _, svgiconDB := range forRestore {

		svgiconDB_ID_atBackupTime := svgiconDB.ID
		svgiconDB.ID = 0
		query := backRepoSVGIcon.db.Create(svgiconDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB[svgiconDB.ID] = svgiconDB
		BackRepoSVGIconid_atBckpTime_newID[svgiconDB_ID_atBackupTime] = svgiconDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SVGIcon file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SVGIcon>id_atBckpTime_newID
// to compute new index
func (backRepoSVGIcon *BackRepoSVGIconStruct) RestorePhaseTwo() {

	for _, svgiconDB := range backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB {

		// next line of code is to avert unused variable compilation error
		_ = svgiconDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoSVGIcon.db.Model(svgiconDB).Updates(*svgiconDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSVGIcon.ResetReversePointers commits all staged instances of SVGIcon to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSVGIcon *BackRepoSVGIconStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, svgicon := range backRepoSVGIcon.Map_SVGIconDBID_SVGIconPtr {
		backRepoSVGIcon.ResetReversePointersInstance(backRepo, idx, svgicon)
	}

	return
}

func (backRepoSVGIcon *BackRepoSVGIconStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, svgicon *models.SVGIcon) (Error error) {

	// fetch matching svgiconDB
	if svgiconDB, ok := backRepoSVGIcon.Map_SVGIconDBID_SVGIconDB[idx]; ok {
		_ = svgiconDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSVGIconid_atBckpTime_newID map[uint]uint
