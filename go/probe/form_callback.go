// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmusicxml/go/models"
	"github.com/fullstack-lang/gongmusicxml/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__BookmarkFormCallback(
	bookmark *models.Bookmark,
	probe *Probe,
	formGroup *table.FormGroup,
) (bookmarkFormCallback *BookmarkFormCallback) {
	bookmarkFormCallback = new(BookmarkFormCallback)
	bookmarkFormCallback.probe = probe
	bookmarkFormCallback.bookmark = bookmark
	bookmarkFormCallback.formGroup = formGroup

	bookmarkFormCallback.CreationMode = (bookmark == nil)

	return
}

type BookmarkFormCallback struct {
	bookmark *models.Bookmark

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bookmarkFormCallback *BookmarkFormCallback) OnSave() {

	log.Println("BookmarkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bookmarkFormCallback.probe.formStage.Checkout()

	if bookmarkFormCallback.bookmark == nil {
		bookmarkFormCallback.bookmark = new(models.Bookmark).Stage(bookmarkFormCallback.probe.stageOfInterest)
	}
	bookmark_ := bookmarkFormCallback.bookmark
	_ = bookmark_

	for _, formDiv := range bookmarkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bookmark_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if bookmarkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bookmark_.Unstage(bookmarkFormCallback.probe.stageOfInterest)
	}

	bookmarkFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bookmark](
		bookmarkFormCallback.probe,
	)
	bookmarkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bookmarkFormCallback.CreationMode || bookmarkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bookmarkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(bookmarkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BookmarkFormCallback(
			nil,
			bookmarkFormCallback.probe,
			newFormGroup,
		)
		bookmark := new(models.Bookmark)
		FillUpForm(bookmark, newFormGroup, bookmarkFormCallback.probe)
		bookmarkFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bookmarkFormCallback.probe)
}
func __gong__New__FooFormCallback(
	foo *models.Foo,
	probe *Probe,
	formGroup *table.FormGroup,
) (fooFormCallback *FooFormCallback) {
	fooFormCallback = new(FooFormCallback)
	fooFormCallback.probe = probe
	fooFormCallback.foo = foo
	fooFormCallback.formGroup = formGroup

	fooFormCallback.CreationMode = (foo == nil)

	return
}

type FooFormCallback struct {
	foo *models.Foo

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (fooFormCallback *FooFormCallback) OnSave() {

	log.Println("FooFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fooFormCallback.probe.formStage.Checkout()

	if fooFormCallback.foo == nil {
		fooFormCallback.foo = new(models.Foo).Stage(fooFormCallback.probe.stageOfInterest)
	}
	foo_ := fooFormCallback.foo
	_ = foo_

	for _, formDiv := range fooFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(foo_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if fooFormCallback.formGroup.HasSuppressButtonBeenPressed {
		foo_.Unstage(fooFormCallback.probe.stageOfInterest)
	}

	fooFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Foo](
		fooFormCallback.probe,
	)
	fooFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fooFormCallback.CreationMode || fooFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fooFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(fooFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FooFormCallback(
			nil,
			fooFormCallback.probe,
			newFormGroup,
		)
		foo := new(models.Foo)
		FillUpForm(foo, newFormGroup, fooFormCallback.probe)
		fooFormCallback.probe.formStage.Commit()
	}

	fillUpTree(fooFormCallback.probe)
}
func __gong__New__LinkFormCallback(
	link *models.Link,
	probe *Probe,
	formGroup *table.FormGroup,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.probe = probe
	linkFormCallback.link = link
	linkFormCallback.formGroup = formGroup

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.probe.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.probe.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	for _, formDiv := range linkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		link_.Unstage(linkFormCallback.probe.stageOfInterest)
	}

	linkFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Link](
		linkFormCallback.probe,
	)
	linkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode || linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(linkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			linkFormCallback.probe,
			newFormGroup,
		)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.probe)
		linkFormCallback.probe.formStage.Commit()
	}

	fillUpTree(linkFormCallback.probe)
}
func __gong__New__LyricFormCallback(
	lyric *models.Lyric,
	probe *Probe,
	formGroup *table.FormGroup,
) (lyricFormCallback *LyricFormCallback) {
	lyricFormCallback = new(LyricFormCallback)
	lyricFormCallback.probe = probe
	lyricFormCallback.lyric = lyric
	lyricFormCallback.formGroup = formGroup

	lyricFormCallback.CreationMode = (lyric == nil)

	return
}

type LyricFormCallback struct {
	lyric *models.Lyric

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lyricFormCallback *LyricFormCallback) OnSave() {

	log.Println("LyricFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lyricFormCallback.probe.formStage.Checkout()

	if lyricFormCallback.lyric == nil {
		lyricFormCallback.lyric = new(models.Lyric).Stage(lyricFormCallback.probe.stageOfInterest)
	}
	lyric_ := lyricFormCallback.lyric
	_ = lyric_

	for _, formDiv := range lyricFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lyric_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if lyricFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lyric_.Unstage(lyricFormCallback.probe.stageOfInterest)
	}

	lyricFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Lyric](
		lyricFormCallback.probe,
	)
	lyricFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lyricFormCallback.CreationMode || lyricFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lyricFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lyricFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LyricFormCallback(
			nil,
			lyricFormCallback.probe,
			newFormGroup,
		)
		lyric := new(models.Lyric)
		FillUpForm(lyric, newFormGroup, lyricFormCallback.probe)
		lyricFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lyricFormCallback.probe)
}
func __gong__New__Lyric_fontFormCallback(
	lyric_font *models.Lyric_font,
	probe *Probe,
	formGroup *table.FormGroup,
) (lyric_fontFormCallback *Lyric_fontFormCallback) {
	lyric_fontFormCallback = new(Lyric_fontFormCallback)
	lyric_fontFormCallback.probe = probe
	lyric_fontFormCallback.lyric_font = lyric_font
	lyric_fontFormCallback.formGroup = formGroup

	lyric_fontFormCallback.CreationMode = (lyric_font == nil)

	return
}

type Lyric_fontFormCallback struct {
	lyric_font *models.Lyric_font

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lyric_fontFormCallback *Lyric_fontFormCallback) OnSave() {

	log.Println("Lyric_fontFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lyric_fontFormCallback.probe.formStage.Checkout()

	if lyric_fontFormCallback.lyric_font == nil {
		lyric_fontFormCallback.lyric_font = new(models.Lyric_font).Stage(lyric_fontFormCallback.probe.stageOfInterest)
	}
	lyric_font_ := lyric_fontFormCallback.lyric_font
	_ = lyric_font_

	for _, formDiv := range lyric_fontFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lyric_font_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if lyric_fontFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lyric_font_.Unstage(lyric_fontFormCallback.probe.stageOfInterest)
	}

	lyric_fontFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Lyric_font](
		lyric_fontFormCallback.probe,
	)
	lyric_fontFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lyric_fontFormCallback.CreationMode || lyric_fontFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lyric_fontFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lyric_fontFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Lyric_fontFormCallback(
			nil,
			lyric_fontFormCallback.probe,
			newFormGroup,
		)
		lyric_font := new(models.Lyric_font)
		FillUpForm(lyric_font, newFormGroup, lyric_fontFormCallback.probe)
		lyric_fontFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lyric_fontFormCallback.probe)
}
func __gong__New__Lyric_languageFormCallback(
	lyric_language *models.Lyric_language,
	probe *Probe,
	formGroup *table.FormGroup,
) (lyric_languageFormCallback *Lyric_languageFormCallback) {
	lyric_languageFormCallback = new(Lyric_languageFormCallback)
	lyric_languageFormCallback.probe = probe
	lyric_languageFormCallback.lyric_language = lyric_language
	lyric_languageFormCallback.formGroup = formGroup

	lyric_languageFormCallback.CreationMode = (lyric_language == nil)

	return
}

type Lyric_languageFormCallback struct {
	lyric_language *models.Lyric_language

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lyric_languageFormCallback *Lyric_languageFormCallback) OnSave() {

	log.Println("Lyric_languageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lyric_languageFormCallback.probe.formStage.Checkout()

	if lyric_languageFormCallback.lyric_language == nil {
		lyric_languageFormCallback.lyric_language = new(models.Lyric_language).Stage(lyric_languageFormCallback.probe.stageOfInterest)
	}
	lyric_language_ := lyric_languageFormCallback.lyric_language
	_ = lyric_language_

	for _, formDiv := range lyric_languageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lyric_language_.Name), formDiv)
		case "EmptyString":
			FormDivBasicFieldToField(&(lyric_language_.EmptyString), formDiv)
		}
	}

	// manage the suppress operation
	if lyric_languageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lyric_language_.Unstage(lyric_languageFormCallback.probe.stageOfInterest)
	}

	lyric_languageFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Lyric_language](
		lyric_languageFormCallback.probe,
	)
	lyric_languageFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lyric_languageFormCallback.CreationMode || lyric_languageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lyric_languageFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lyric_languageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Lyric_languageFormCallback(
			nil,
			lyric_languageFormCallback.probe,
			newFormGroup,
		)
		lyric_language := new(models.Lyric_language)
		FillUpForm(lyric_language, newFormGroup, lyric_languageFormCallback.probe)
		lyric_languageFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lyric_languageFormCallback.probe)
}
func __gong__New__Miscellaneous_fieldFormCallback(
	miscellaneous_field *models.Miscellaneous_field,
	probe *Probe,
	formGroup *table.FormGroup,
) (miscellaneous_fieldFormCallback *Miscellaneous_fieldFormCallback) {
	miscellaneous_fieldFormCallback = new(Miscellaneous_fieldFormCallback)
	miscellaneous_fieldFormCallback.probe = probe
	miscellaneous_fieldFormCallback.miscellaneous_field = miscellaneous_field
	miscellaneous_fieldFormCallback.formGroup = formGroup

	miscellaneous_fieldFormCallback.CreationMode = (miscellaneous_field == nil)

	return
}

type Miscellaneous_fieldFormCallback struct {
	miscellaneous_field *models.Miscellaneous_field

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (miscellaneous_fieldFormCallback *Miscellaneous_fieldFormCallback) OnSave() {

	log.Println("Miscellaneous_fieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	miscellaneous_fieldFormCallback.probe.formStage.Checkout()

	if miscellaneous_fieldFormCallback.miscellaneous_field == nil {
		miscellaneous_fieldFormCallback.miscellaneous_field = new(models.Miscellaneous_field).Stage(miscellaneous_fieldFormCallback.probe.stageOfInterest)
	}
	miscellaneous_field_ := miscellaneous_fieldFormCallback.miscellaneous_field
	_ = miscellaneous_field_

	for _, formDiv := range miscellaneous_fieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Value":
			FormDivBasicFieldToField(&(miscellaneous_field_.Value), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(miscellaneous_field_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if miscellaneous_fieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		miscellaneous_field_.Unstage(miscellaneous_fieldFormCallback.probe.stageOfInterest)
	}

	miscellaneous_fieldFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Miscellaneous_field](
		miscellaneous_fieldFormCallback.probe,
	)
	miscellaneous_fieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if miscellaneous_fieldFormCallback.CreationMode || miscellaneous_fieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		miscellaneous_fieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(miscellaneous_fieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Miscellaneous_fieldFormCallback(
			nil,
			miscellaneous_fieldFormCallback.probe,
			newFormGroup,
		)
		miscellaneous_field := new(models.Miscellaneous_field)
		FillUpForm(miscellaneous_field, newFormGroup, miscellaneous_fieldFormCallback.probe)
		miscellaneous_fieldFormCallback.probe.formStage.Commit()
	}

	fillUpTree(miscellaneous_fieldFormCallback.probe)
}
