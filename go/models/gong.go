// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Bookmarks           map[*Bookmark]any
	Bookmarks_mapString map[string]*Bookmark

	// insertion point for slice of pointers maps

	OnAfterBookmarkCreateCallback OnAfterCreateInterface[Bookmark]
	OnAfterBookmarkUpdateCallback OnAfterUpdateInterface[Bookmark]
	OnAfterBookmarkDeleteCallback OnAfterDeleteInterface[Bookmark]
	OnAfterBookmarkReadCallback   OnAfterReadInterface[Bookmark]

	Foos           map[*Foo]any
	Foos_mapString map[string]*Foo

	// insertion point for slice of pointers maps

	OnAfterFooCreateCallback OnAfterCreateInterface[Foo]
	OnAfterFooUpdateCallback OnAfterUpdateInterface[Foo]
	OnAfterFooDeleteCallback OnAfterDeleteInterface[Foo]
	OnAfterFooReadCallback   OnAfterReadInterface[Foo]

	Links           map[*Link]any
	Links_mapString map[string]*Link

	// insertion point for slice of pointers maps

	OnAfterLinkCreateCallback OnAfterCreateInterface[Link]
	OnAfterLinkUpdateCallback OnAfterUpdateInterface[Link]
	OnAfterLinkDeleteCallback OnAfterDeleteInterface[Link]
	OnAfterLinkReadCallback   OnAfterReadInterface[Link]

	Lyrics           map[*Lyric]any
	Lyrics_mapString map[string]*Lyric

	// insertion point for slice of pointers maps

	OnAfterLyricCreateCallback OnAfterCreateInterface[Lyric]
	OnAfterLyricUpdateCallback OnAfterUpdateInterface[Lyric]
	OnAfterLyricDeleteCallback OnAfterDeleteInterface[Lyric]
	OnAfterLyricReadCallback   OnAfterReadInterface[Lyric]

	Lyric_fonts           map[*Lyric_font]any
	Lyric_fonts_mapString map[string]*Lyric_font

	// insertion point for slice of pointers maps

	OnAfterLyric_fontCreateCallback OnAfterCreateInterface[Lyric_font]
	OnAfterLyric_fontUpdateCallback OnAfterUpdateInterface[Lyric_font]
	OnAfterLyric_fontDeleteCallback OnAfterDeleteInterface[Lyric_font]
	OnAfterLyric_fontReadCallback   OnAfterReadInterface[Lyric_font]

	Lyric_languages           map[*Lyric_language]any
	Lyric_languages_mapString map[string]*Lyric_language

	// insertion point for slice of pointers maps

	OnAfterLyric_languageCreateCallback OnAfterCreateInterface[Lyric_language]
	OnAfterLyric_languageUpdateCallback OnAfterUpdateInterface[Lyric_language]
	OnAfterLyric_languageDeleteCallback OnAfterDeleteInterface[Lyric_language]
	OnAfterLyric_languageReadCallback   OnAfterReadInterface[Lyric_language]

	Miscellaneous_fields           map[*Miscellaneous_field]any
	Miscellaneous_fields_mapString map[string]*Miscellaneous_field

	// insertion point for slice of pointers maps

	OnAfterMiscellaneous_fieldCreateCallback OnAfterCreateInterface[Miscellaneous_field]
	OnAfterMiscellaneous_fieldUpdateCallback OnAfterUpdateInterface[Miscellaneous_field]
	OnAfterMiscellaneous_fieldDeleteCallback OnAfterDeleteInterface[Miscellaneous_field]
	OnAfterMiscellaneous_fieldReadCallback   OnAfterReadInterface[Miscellaneous_field]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongmusicxml/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitBookmark(bookmark *Bookmark)
	CheckoutBookmark(bookmark *Bookmark)
	CommitFoo(foo *Foo)
	CheckoutFoo(foo *Foo)
	CommitLink(link *Link)
	CheckoutLink(link *Link)
	CommitLyric(lyric *Lyric)
	CheckoutLyric(lyric *Lyric)
	CommitLyric_font(lyric_font *Lyric_font)
	CheckoutLyric_font(lyric_font *Lyric_font)
	CommitLyric_language(lyric_language *Lyric_language)
	CheckoutLyric_language(lyric_language *Lyric_language)
	CommitMiscellaneous_field(miscellaneous_field *Miscellaneous_field)
	CheckoutMiscellaneous_field(miscellaneous_field *Miscellaneous_field)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Bookmarks:           make(map[*Bookmark]any),
		Bookmarks_mapString: make(map[string]*Bookmark),

		Foos:           make(map[*Foo]any),
		Foos_mapString: make(map[string]*Foo),

		Links:           make(map[*Link]any),
		Links_mapString: make(map[string]*Link),

		Lyrics:           make(map[*Lyric]any),
		Lyrics_mapString: make(map[string]*Lyric),

		Lyric_fonts:           make(map[*Lyric_font]any),
		Lyric_fonts_mapString: make(map[string]*Lyric_font),

		Lyric_languages:           make(map[*Lyric_language]any),
		Lyric_languages_mapString: make(map[string]*Lyric_language),

		Miscellaneous_fields:           make(map[*Miscellaneous_field]any),
		Miscellaneous_fields_mapString: make(map[string]*Miscellaneous_field),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Bookmark"] = len(stage.Bookmarks)
	stage.Map_GongStructName_InstancesNb["Foo"] = len(stage.Foos)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
	stage.Map_GongStructName_InstancesNb["Lyric"] = len(stage.Lyrics)
	stage.Map_GongStructName_InstancesNb["Lyric_font"] = len(stage.Lyric_fonts)
	stage.Map_GongStructName_InstancesNb["Lyric_language"] = len(stage.Lyric_languages)
	stage.Map_GongStructName_InstancesNb["Miscellaneous_field"] = len(stage.Miscellaneous_fields)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Bookmark"] = len(stage.Bookmarks)
	stage.Map_GongStructName_InstancesNb["Foo"] = len(stage.Foos)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
	stage.Map_GongStructName_InstancesNb["Lyric"] = len(stage.Lyrics)
	stage.Map_GongStructName_InstancesNb["Lyric_font"] = len(stage.Lyric_fonts)
	stage.Map_GongStructName_InstancesNb["Lyric_language"] = len(stage.Lyric_languages)
	stage.Map_GongStructName_InstancesNb["Miscellaneous_field"] = len(stage.Miscellaneous_fields)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts bookmark to the model stage
func (bookmark *Bookmark) Stage(stage *StageStruct) *Bookmark {
	stage.Bookmarks[bookmark] = __member
	stage.Bookmarks_mapString[bookmark.Name] = bookmark

	return bookmark
}

// Unstage removes bookmark off the model stage
func (bookmark *Bookmark) Unstage(stage *StageStruct) *Bookmark {
	delete(stage.Bookmarks, bookmark)
	delete(stage.Bookmarks_mapString, bookmark.Name)
	return bookmark
}

// UnstageVoid removes bookmark off the model stage
func (bookmark *Bookmark) UnstageVoid(stage *StageStruct) {
	delete(stage.Bookmarks, bookmark)
	delete(stage.Bookmarks_mapString, bookmark.Name)
}

// commit bookmark to the back repo (if it is already staged)
func (bookmark *Bookmark) Commit(stage *StageStruct) *Bookmark {
	if _, ok := stage.Bookmarks[bookmark]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBookmark(bookmark)
		}
	}
	return bookmark
}

func (bookmark *Bookmark) CommitVoid(stage *StageStruct) {
	bookmark.Commit(stage)
}

// Checkout bookmark to the back repo (if it is already staged)
func (bookmark *Bookmark) Checkout(stage *StageStruct) *Bookmark {
	if _, ok := stage.Bookmarks[bookmark]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBookmark(bookmark)
		}
	}
	return bookmark
}

// for satisfaction of GongStruct interface
func (bookmark *Bookmark) GetName() (res string) {
	return bookmark.Name
}

// Stage puts foo to the model stage
func (foo *Foo) Stage(stage *StageStruct) *Foo {
	stage.Foos[foo] = __member
	stage.Foos_mapString[foo.Name] = foo

	return foo
}

// Unstage removes foo off the model stage
func (foo *Foo) Unstage(stage *StageStruct) *Foo {
	delete(stage.Foos, foo)
	delete(stage.Foos_mapString, foo.Name)
	return foo
}

// UnstageVoid removes foo off the model stage
func (foo *Foo) UnstageVoid(stage *StageStruct) {
	delete(stage.Foos, foo)
	delete(stage.Foos_mapString, foo.Name)
}

// commit foo to the back repo (if it is already staged)
func (foo *Foo) Commit(stage *StageStruct) *Foo {
	if _, ok := stage.Foos[foo]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFoo(foo)
		}
	}
	return foo
}

func (foo *Foo) CommitVoid(stage *StageStruct) {
	foo.Commit(stage)
}

// Checkout foo to the back repo (if it is already staged)
func (foo *Foo) Checkout(stage *StageStruct) *Foo {
	if _, ok := stage.Foos[foo]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFoo(foo)
		}
	}
	return foo
}

// for satisfaction of GongStruct interface
func (foo *Foo) GetName() (res string) {
	return foo.Name
}

// Stage puts link to the model stage
func (link *Link) Stage(stage *StageStruct) *Link {
	stage.Links[link] = __member
	stage.Links_mapString[link.Name] = link

	return link
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *StageStruct) *Link {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *StageStruct) {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
}

// commit link to the back repo (if it is already staged)
func (link *Link) Commit(stage *StageStruct) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLink(link)
		}
	}
	return link
}

func (link *Link) CommitVoid(stage *StageStruct) {
	link.Commit(stage)
}

// Checkout link to the back repo (if it is already staged)
func (link *Link) Checkout(stage *StageStruct) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLink(link)
		}
	}
	return link
}

// for satisfaction of GongStruct interface
func (link *Link) GetName() (res string) {
	return link.Name
}

// Stage puts lyric to the model stage
func (lyric *Lyric) Stage(stage *StageStruct) *Lyric {
	stage.Lyrics[lyric] = __member
	stage.Lyrics_mapString[lyric.Name] = lyric

	return lyric
}

// Unstage removes lyric off the model stage
func (lyric *Lyric) Unstage(stage *StageStruct) *Lyric {
	delete(stage.Lyrics, lyric)
	delete(stage.Lyrics_mapString, lyric.Name)
	return lyric
}

// UnstageVoid removes lyric off the model stage
func (lyric *Lyric) UnstageVoid(stage *StageStruct) {
	delete(stage.Lyrics, lyric)
	delete(stage.Lyrics_mapString, lyric.Name)
}

// commit lyric to the back repo (if it is already staged)
func (lyric *Lyric) Commit(stage *StageStruct) *Lyric {
	if _, ok := stage.Lyrics[lyric]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLyric(lyric)
		}
	}
	return lyric
}

func (lyric *Lyric) CommitVoid(stage *StageStruct) {
	lyric.Commit(stage)
}

// Checkout lyric to the back repo (if it is already staged)
func (lyric *Lyric) Checkout(stage *StageStruct) *Lyric {
	if _, ok := stage.Lyrics[lyric]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLyric(lyric)
		}
	}
	return lyric
}

// for satisfaction of GongStruct interface
func (lyric *Lyric) GetName() (res string) {
	return lyric.Name
}

// Stage puts lyric_font to the model stage
func (lyric_font *Lyric_font) Stage(stage *StageStruct) *Lyric_font {
	stage.Lyric_fonts[lyric_font] = __member
	stage.Lyric_fonts_mapString[lyric_font.Name] = lyric_font

	return lyric_font
}

// Unstage removes lyric_font off the model stage
func (lyric_font *Lyric_font) Unstage(stage *StageStruct) *Lyric_font {
	delete(stage.Lyric_fonts, lyric_font)
	delete(stage.Lyric_fonts_mapString, lyric_font.Name)
	return lyric_font
}

// UnstageVoid removes lyric_font off the model stage
func (lyric_font *Lyric_font) UnstageVoid(stage *StageStruct) {
	delete(stage.Lyric_fonts, lyric_font)
	delete(stage.Lyric_fonts_mapString, lyric_font.Name)
}

// commit lyric_font to the back repo (if it is already staged)
func (lyric_font *Lyric_font) Commit(stage *StageStruct) *Lyric_font {
	if _, ok := stage.Lyric_fonts[lyric_font]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLyric_font(lyric_font)
		}
	}
	return lyric_font
}

func (lyric_font *Lyric_font) CommitVoid(stage *StageStruct) {
	lyric_font.Commit(stage)
}

// Checkout lyric_font to the back repo (if it is already staged)
func (lyric_font *Lyric_font) Checkout(stage *StageStruct) *Lyric_font {
	if _, ok := stage.Lyric_fonts[lyric_font]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLyric_font(lyric_font)
		}
	}
	return lyric_font
}

// for satisfaction of GongStruct interface
func (lyric_font *Lyric_font) GetName() (res string) {
	return lyric_font.Name
}

// Stage puts lyric_language to the model stage
func (lyric_language *Lyric_language) Stage(stage *StageStruct) *Lyric_language {
	stage.Lyric_languages[lyric_language] = __member
	stage.Lyric_languages_mapString[lyric_language.Name] = lyric_language

	return lyric_language
}

// Unstage removes lyric_language off the model stage
func (lyric_language *Lyric_language) Unstage(stage *StageStruct) *Lyric_language {
	delete(stage.Lyric_languages, lyric_language)
	delete(stage.Lyric_languages_mapString, lyric_language.Name)
	return lyric_language
}

// UnstageVoid removes lyric_language off the model stage
func (lyric_language *Lyric_language) UnstageVoid(stage *StageStruct) {
	delete(stage.Lyric_languages, lyric_language)
	delete(stage.Lyric_languages_mapString, lyric_language.Name)
}

// commit lyric_language to the back repo (if it is already staged)
func (lyric_language *Lyric_language) Commit(stage *StageStruct) *Lyric_language {
	if _, ok := stage.Lyric_languages[lyric_language]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLyric_language(lyric_language)
		}
	}
	return lyric_language
}

func (lyric_language *Lyric_language) CommitVoid(stage *StageStruct) {
	lyric_language.Commit(stage)
}

// Checkout lyric_language to the back repo (if it is already staged)
func (lyric_language *Lyric_language) Checkout(stage *StageStruct) *Lyric_language {
	if _, ok := stage.Lyric_languages[lyric_language]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLyric_language(lyric_language)
		}
	}
	return lyric_language
}

// for satisfaction of GongStruct interface
func (lyric_language *Lyric_language) GetName() (res string) {
	return lyric_language.Name
}

// Stage puts miscellaneous_field to the model stage
func (miscellaneous_field *Miscellaneous_field) Stage(stage *StageStruct) *Miscellaneous_field {
	stage.Miscellaneous_fields[miscellaneous_field] = __member
	stage.Miscellaneous_fields_mapString[miscellaneous_field.Name] = miscellaneous_field

	return miscellaneous_field
}

// Unstage removes miscellaneous_field off the model stage
func (miscellaneous_field *Miscellaneous_field) Unstage(stage *StageStruct) *Miscellaneous_field {
	delete(stage.Miscellaneous_fields, miscellaneous_field)
	delete(stage.Miscellaneous_fields_mapString, miscellaneous_field.Name)
	return miscellaneous_field
}

// UnstageVoid removes miscellaneous_field off the model stage
func (miscellaneous_field *Miscellaneous_field) UnstageVoid(stage *StageStruct) {
	delete(stage.Miscellaneous_fields, miscellaneous_field)
	delete(stage.Miscellaneous_fields_mapString, miscellaneous_field.Name)
}

// commit miscellaneous_field to the back repo (if it is already staged)
func (miscellaneous_field *Miscellaneous_field) Commit(stage *StageStruct) *Miscellaneous_field {
	if _, ok := stage.Miscellaneous_fields[miscellaneous_field]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMiscellaneous_field(miscellaneous_field)
		}
	}
	return miscellaneous_field
}

func (miscellaneous_field *Miscellaneous_field) CommitVoid(stage *StageStruct) {
	miscellaneous_field.Commit(stage)
}

// Checkout miscellaneous_field to the back repo (if it is already staged)
func (miscellaneous_field *Miscellaneous_field) Checkout(stage *StageStruct) *Miscellaneous_field {
	if _, ok := stage.Miscellaneous_fields[miscellaneous_field]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMiscellaneous_field(miscellaneous_field)
		}
	}
	return miscellaneous_field
}

// for satisfaction of GongStruct interface
func (miscellaneous_field *Miscellaneous_field) GetName() (res string) {
	return miscellaneous_field.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMBookmark(Bookmark *Bookmark)
	CreateORMFoo(Foo *Foo)
	CreateORMLink(Link *Link)
	CreateORMLyric(Lyric *Lyric)
	CreateORMLyric_font(Lyric_font *Lyric_font)
	CreateORMLyric_language(Lyric_language *Lyric_language)
	CreateORMMiscellaneous_field(Miscellaneous_field *Miscellaneous_field)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMBookmark(Bookmark *Bookmark)
	DeleteORMFoo(Foo *Foo)
	DeleteORMLink(Link *Link)
	DeleteORMLyric(Lyric *Lyric)
	DeleteORMLyric_font(Lyric_font *Lyric_font)
	DeleteORMLyric_language(Lyric_language *Lyric_language)
	DeleteORMMiscellaneous_field(Miscellaneous_field *Miscellaneous_field)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Bookmarks = make(map[*Bookmark]any)
	stage.Bookmarks_mapString = make(map[string]*Bookmark)

	stage.Foos = make(map[*Foo]any)
	stage.Foos_mapString = make(map[string]*Foo)

	stage.Links = make(map[*Link]any)
	stage.Links_mapString = make(map[string]*Link)

	stage.Lyrics = make(map[*Lyric]any)
	stage.Lyrics_mapString = make(map[string]*Lyric)

	stage.Lyric_fonts = make(map[*Lyric_font]any)
	stage.Lyric_fonts_mapString = make(map[string]*Lyric_font)

	stage.Lyric_languages = make(map[*Lyric_language]any)
	stage.Lyric_languages_mapString = make(map[string]*Lyric_language)

	stage.Miscellaneous_fields = make(map[*Miscellaneous_field]any)
	stage.Miscellaneous_fields_mapString = make(map[string]*Miscellaneous_field)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Bookmarks = nil
	stage.Bookmarks_mapString = nil

	stage.Foos = nil
	stage.Foos_mapString = nil

	stage.Links = nil
	stage.Links_mapString = nil

	stage.Lyrics = nil
	stage.Lyrics_mapString = nil

	stage.Lyric_fonts = nil
	stage.Lyric_fonts_mapString = nil

	stage.Lyric_languages = nil
	stage.Lyric_languages_mapString = nil

	stage.Miscellaneous_fields = nil
	stage.Miscellaneous_fields_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for bookmark := range stage.Bookmarks {
		bookmark.Unstage(stage)
	}

	for foo := range stage.Foos {
		foo.Unstage(stage)
	}

	for link := range stage.Links {
		link.Unstage(stage)
	}

	for lyric := range stage.Lyrics {
		lyric.Unstage(stage)
	}

	for lyric_font := range stage.Lyric_fonts {
		lyric_font.Unstage(stage)
	}

	for lyric_language := range stage.Lyric_languages {
		lyric_language.Unstage(stage)
	}

	for miscellaneous_field := range stage.Miscellaneous_fields {
		miscellaneous_field.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	Bookmark | Foo | Link | Lyric | Lyric_font | Lyric_language | Miscellaneous_field
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*Bookmark | *Foo | *Link | *Lyric | *Lyric_font | *Lyric_language | *Miscellaneous_field
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*Bookmark]any |
		map[*Foo]any |
		map[*Link]any |
		map[*Lyric]any |
		map[*Lyric_font]any |
		map[*Lyric_language]any |
		map[*Miscellaneous_field]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*Bookmark |
		map[string]*Foo |
		map[string]*Link |
		map[string]*Lyric |
		map[string]*Lyric_font |
		map[string]*Lyric_language |
		map[string]*Miscellaneous_field |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Bookmark]any:
		return any(&stage.Bookmarks).(*Type)
	case map[*Foo]any:
		return any(&stage.Foos).(*Type)
	case map[*Link]any:
		return any(&stage.Links).(*Type)
	case map[*Lyric]any:
		return any(&stage.Lyrics).(*Type)
	case map[*Lyric_font]any:
		return any(&stage.Lyric_fonts).(*Type)
	case map[*Lyric_language]any:
		return any(&stage.Lyric_languages).(*Type)
	case map[*Miscellaneous_field]any:
		return any(&stage.Miscellaneous_fields).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Bookmark:
		return any(&stage.Bookmarks_mapString).(*Type)
	case map[string]*Foo:
		return any(&stage.Foos_mapString).(*Type)
	case map[string]*Link:
		return any(&stage.Links_mapString).(*Type)
	case map[string]*Lyric:
		return any(&stage.Lyrics_mapString).(*Type)
	case map[string]*Lyric_font:
		return any(&stage.Lyric_fonts_mapString).(*Type)
	case map[string]*Lyric_language:
		return any(&stage.Lyric_languages_mapString).(*Type)
	case map[string]*Miscellaneous_field:
		return any(&stage.Miscellaneous_fields_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Bookmark:
		return any(&stage.Bookmarks).(*map[*Type]any)
	case Foo:
		return any(&stage.Foos).(*map[*Type]any)
	case Link:
		return any(&stage.Links).(*map[*Type]any)
	case Lyric:
		return any(&stage.Lyrics).(*map[*Type]any)
	case Lyric_font:
		return any(&stage.Lyric_fonts).(*map[*Type]any)
	case Lyric_language:
		return any(&stage.Lyric_languages).(*map[*Type]any)
	case Miscellaneous_field:
		return any(&stage.Miscellaneous_fields).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Bookmark:
		return any(&stage.Bookmarks).(*map[Type]any)
	case *Foo:
		return any(&stage.Foos).(*map[Type]any)
	case *Link:
		return any(&stage.Links).(*map[Type]any)
	case *Lyric:
		return any(&stage.Lyrics).(*map[Type]any)
	case *Lyric_font:
		return any(&stage.Lyric_fonts).(*map[Type]any)
	case *Lyric_language:
		return any(&stage.Lyric_languages).(*map[Type]any)
	case *Miscellaneous_field:
		return any(&stage.Miscellaneous_fields).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Bookmark:
		return any(&stage.Bookmarks_mapString).(*map[string]*Type)
	case Foo:
		return any(&stage.Foos_mapString).(*map[string]*Type)
	case Link:
		return any(&stage.Links_mapString).(*map[string]*Type)
	case Lyric:
		return any(&stage.Lyrics_mapString).(*map[string]*Type)
	case Lyric_font:
		return any(&stage.Lyric_fonts_mapString).(*map[string]*Type)
	case Lyric_language:
		return any(&stage.Lyric_languages_mapString).(*map[string]*Type)
	case Miscellaneous_field:
		return any(&stage.Miscellaneous_fields_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Bookmark:
		return any(&Bookmark{
			// Initialisation of associations
		}).(*Type)
	case Foo:
		return any(&Foo{
			// Initialisation of associations
		}).(*Type)
	case Link:
		return any(&Link{
			// Initialisation of associations
		}).(*Type)
	case Lyric:
		return any(&Lyric{
			// Initialisation of associations
		}).(*Type)
	case Lyric_font:
		return any(&Lyric_font{
			// Initialisation of associations
		}).(*Type)
	case Lyric_language:
		return any(&Lyric_language{
			// Initialisation of associations
		}).(*Type)
	case Miscellaneous_field:
		return any(&Miscellaneous_field{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Bookmark
	case Bookmark:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Foo
	case Foo:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lyric
	case Lyric:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lyric_font
	case Lyric_font:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lyric_language
	case Lyric_language:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Miscellaneous_field
	case Miscellaneous_field:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Bookmark
	case Bookmark:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Foo
	case Foo:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lyric
	case Lyric:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lyric_font
	case Lyric_font:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lyric_language
	case Lyric_language:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Miscellaneous_field
	case Miscellaneous_field:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Bookmark:
		res = "Bookmark"
	case Foo:
		res = "Foo"
	case Link:
		res = "Link"
	case Lyric:
		res = "Lyric"
	case Lyric_font:
		res = "Lyric_font"
	case Lyric_language:
		res = "Lyric_language"
	case Miscellaneous_field:
		res = "Miscellaneous_field"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Bookmark:
		res = "Bookmark"
	case *Foo:
		res = "Foo"
	case *Link:
		res = "Link"
	case *Lyric:
		res = "Lyric"
	case *Lyric_font:
		res = "Lyric_font"
	case *Lyric_language:
		res = "Lyric_language"
	case *Miscellaneous_field:
		res = "Miscellaneous_field"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Bookmark:
		res = []string{"Name"}
	case Foo:
		res = []string{"Name"}
	case Link:
		res = []string{"Name"}
	case Lyric:
		res = []string{"Name"}
	case Lyric_font:
		res = []string{"Name"}
	case Lyric_language:
		res = []string{"Name", "EmptyString"}
	case Miscellaneous_field:
		res = []string{"Value", "Name"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Bookmark:
		var rf ReverseField
		_ = rf
	case Foo:
		var rf ReverseField
		_ = rf
	case Link:
		var rf ReverseField
		_ = rf
	case Lyric:
		var rf ReverseField
		_ = rf
	case Lyric_font:
		var rf ReverseField
		_ = rf
	case Lyric_language:
		var rf ReverseField
		_ = rf
	case Miscellaneous_field:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Bookmark:
		res = []string{"Name"}
	case *Foo:
		res = []string{"Name"}
	case *Link:
		res = []string{"Name"}
	case *Lyric:
		res = []string{"Name"}
	case *Lyric_font:
		res = []string{"Name"}
	case *Lyric_language:
		res = []string{"Name", "EmptyString"}
	case *Miscellaneous_field:
		res = []string{"Value", "Name"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Bookmark:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Foo:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Lyric:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Lyric_font:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Lyric_language:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EmptyString":
			res = inferedInstance.EmptyString
		}
	case *Miscellaneous_field:
		switch fieldName {
		// string value of fields
		case "Value":
			res = inferedInstance.Value
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Bookmark:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Foo:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Lyric:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Lyric_font:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Lyric_language:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EmptyString":
			res = inferedInstance.EmptyString
		}
	case Miscellaneous_field:
		switch fieldName {
		// string value of fields
		case "Value":
			res = inferedInstance.Value
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
