// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongmusicxml/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoBookmark BackRepoBookmarkStruct

	BackRepoFoo BackRepoFooStruct

	BackRepoLink BackRepoLinkStruct

	BackRepoLyric BackRepoLyricStruct

	BackRepoLyric_font BackRepoLyric_fontStruct

	BackRepoLyric_language BackRepoLyric_languageStruct

	BackRepoMiscellaneous_field BackRepoMiscellaneous_fieldStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&BookmarkDB{},
		&FooDB{},
		&LinkDB{},
		&LyricDB{},
		&Lyric_fontDB{},
		&Lyric_languageDB{},
		&Miscellaneous_fieldDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoBookmark = BackRepoBookmarkStruct{
		Map_BookmarkDBID_BookmarkPtr: make(map[uint]*models.Bookmark, 0),
		Map_BookmarkDBID_BookmarkDB:  make(map[uint]*BookmarkDB, 0),
		Map_BookmarkPtr_BookmarkDBID: make(map[*models.Bookmark]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFoo = BackRepoFooStruct{
		Map_FooDBID_FooPtr: make(map[uint]*models.Foo, 0),
		Map_FooDBID_FooDB:  make(map[uint]*FooDB, 0),
		Map_FooPtr_FooDBID: make(map[*models.Foo]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLink = BackRepoLinkStruct{
		Map_LinkDBID_LinkPtr: make(map[uint]*models.Link, 0),
		Map_LinkDBID_LinkDB:  make(map[uint]*LinkDB, 0),
		Map_LinkPtr_LinkDBID: make(map[*models.Link]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLyric = BackRepoLyricStruct{
		Map_LyricDBID_LyricPtr: make(map[uint]*models.Lyric, 0),
		Map_LyricDBID_LyricDB:  make(map[uint]*LyricDB, 0),
		Map_LyricPtr_LyricDBID: make(map[*models.Lyric]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLyric_font = BackRepoLyric_fontStruct{
		Map_Lyric_fontDBID_Lyric_fontPtr: make(map[uint]*models.Lyric_font, 0),
		Map_Lyric_fontDBID_Lyric_fontDB:  make(map[uint]*Lyric_fontDB, 0),
		Map_Lyric_fontPtr_Lyric_fontDBID: make(map[*models.Lyric_font]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLyric_language = BackRepoLyric_languageStruct{
		Map_Lyric_languageDBID_Lyric_languagePtr: make(map[uint]*models.Lyric_language, 0),
		Map_Lyric_languageDBID_Lyric_languageDB:  make(map[uint]*Lyric_languageDB, 0),
		Map_Lyric_languagePtr_Lyric_languageDBID: make(map[*models.Lyric_language]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMiscellaneous_field = BackRepoMiscellaneous_fieldStruct{
		Map_Miscellaneous_fieldDBID_Miscellaneous_fieldPtr: make(map[uint]*models.Miscellaneous_field, 0),
		Map_Miscellaneous_fieldDBID_Miscellaneous_fieldDB:  make(map[uint]*Miscellaneous_fieldDB, 0),
		Map_Miscellaneous_fieldPtr_Miscellaneous_fieldDBID: make(map[*models.Miscellaneous_field]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()
	
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoBookmark.CommitPhaseOne(stage)
	backRepo.BackRepoFoo.CommitPhaseOne(stage)
	backRepo.BackRepoLink.CommitPhaseOne(stage)
	backRepo.BackRepoLyric.CommitPhaseOne(stage)
	backRepo.BackRepoLyric_font.CommitPhaseOne(stage)
	backRepo.BackRepoLyric_language.CommitPhaseOne(stage)
	backRepo.BackRepoMiscellaneous_field.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoBookmark.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFoo.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLink.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLyric.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLyric_font.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLyric_language.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMiscellaneous_field.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoBookmark.CheckoutPhaseOne()
	backRepo.BackRepoFoo.CheckoutPhaseOne()
	backRepo.BackRepoLink.CheckoutPhaseOne()
	backRepo.BackRepoLyric.CheckoutPhaseOne()
	backRepo.BackRepoLyric_font.CheckoutPhaseOne()
	backRepo.BackRepoLyric_language.CheckoutPhaseOne()
	backRepo.BackRepoMiscellaneous_field.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoBookmark.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFoo.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLink.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLyric.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLyric_font.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLyric_language.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMiscellaneous_field.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoBookmark.Backup(dirPath)
	backRepo.BackRepoFoo.Backup(dirPath)
	backRepo.BackRepoLink.Backup(dirPath)
	backRepo.BackRepoLyric.Backup(dirPath)
	backRepo.BackRepoLyric_font.Backup(dirPath)
	backRepo.BackRepoLyric_language.Backup(dirPath)
	backRepo.BackRepoMiscellaneous_field.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoBookmark.BackupXL(file)
	backRepo.BackRepoFoo.BackupXL(file)
	backRepo.BackRepoLink.BackupXL(file)
	backRepo.BackRepoLyric.BackupXL(file)
	backRepo.BackRepoLyric_font.BackupXL(file)
	backRepo.BackRepoLyric_language.BackupXL(file)
	backRepo.BackRepoMiscellaneous_field.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoBookmark.RestorePhaseOne(dirPath)
	backRepo.BackRepoFoo.RestorePhaseOne(dirPath)
	backRepo.BackRepoLink.RestorePhaseOne(dirPath)
	backRepo.BackRepoLyric.RestorePhaseOne(dirPath)
	backRepo.BackRepoLyric_font.RestorePhaseOne(dirPath)
	backRepo.BackRepoLyric_language.RestorePhaseOne(dirPath)
	backRepo.BackRepoMiscellaneous_field.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoBookmark.RestorePhaseTwo()
	backRepo.BackRepoFoo.RestorePhaseTwo()
	backRepo.BackRepoLink.RestorePhaseTwo()
	backRepo.BackRepoLyric.RestorePhaseTwo()
	backRepo.BackRepoLyric_font.RestorePhaseTwo()
	backRepo.BackRepoLyric_language.RestorePhaseTwo()
	backRepo.BackRepoMiscellaneous_field.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoBookmark.RestoreXLPhaseOne(file)
	backRepo.BackRepoFoo.RestoreXLPhaseOne(file)
	backRepo.BackRepoLink.RestoreXLPhaseOne(file)
	backRepo.BackRepoLyric.RestoreXLPhaseOne(file)
	backRepo.BackRepoLyric_font.RestoreXLPhaseOne(file)
	backRepo.BackRepoLyric_language.RestoreXLPhaseOne(file)
	backRepo.BackRepoMiscellaneous_field.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb() <-chan int {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()

	ch := make(chan int)
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	return ch
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.rwMutex.RLock()
	defer backRepoStruct.rwMutex.RUnlock()

	activeChannels := make([]chan int, 0)

	for _, ch := range backRepoStruct.subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			activeChannels = append(activeChannels, ch)
		default:
			// Assume channel is no longer active; don't add to activeChannels
			log.Println("Channel no longer active")
			close(ch)
		}
	}
	backRepoStruct.subscribers = activeChannels
}
