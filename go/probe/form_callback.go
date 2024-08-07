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
func __gong__New__AccidentalFormCallback(
	accidental *models.Accidental,
	probe *Probe,
	formGroup *table.FormGroup,
) (accidentalFormCallback *AccidentalFormCallback) {
	accidentalFormCallback = new(AccidentalFormCallback)
	accidentalFormCallback.probe = probe
	accidentalFormCallback.accidental = accidental
	accidentalFormCallback.formGroup = formGroup

	accidentalFormCallback.CreationMode = (accidental == nil)

	return
}

type AccidentalFormCallback struct {
	accidental *models.Accidental

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (accidentalFormCallback *AccidentalFormCallback) OnSave() {

	log.Println("AccidentalFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	accidentalFormCallback.probe.formStage.Checkout()

	if accidentalFormCallback.accidental == nil {
		accidentalFormCallback.accidental = new(models.Accidental).Stage(accidentalFormCallback.probe.stageOfInterest)
	}
	accidental_ := accidentalFormCallback.accidental
	_ = accidental_

	for _, formDiv := range accidentalFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(accidental_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if accidentalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accidental_.Unstage(accidentalFormCallback.probe.stageOfInterest)
	}

	accidentalFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Accidental](
		accidentalFormCallback.probe,
	)
	accidentalFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if accidentalFormCallback.CreationMode || accidentalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accidentalFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(accidentalFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AccidentalFormCallback(
			nil,
			accidentalFormCallback.probe,
			newFormGroup,
		)
		accidental := new(models.Accidental)
		FillUpForm(accidental, newFormGroup, accidentalFormCallback.probe)
		accidentalFormCallback.probe.formStage.Commit()
	}

	fillUpTree(accidentalFormCallback.probe)
}
func __gong__New__Accidental_markFormCallback(
	accidental_mark *models.Accidental_mark,
	probe *Probe,
	formGroup *table.FormGroup,
) (accidental_markFormCallback *Accidental_markFormCallback) {
	accidental_markFormCallback = new(Accidental_markFormCallback)
	accidental_markFormCallback.probe = probe
	accidental_markFormCallback.accidental_mark = accidental_mark
	accidental_markFormCallback.formGroup = formGroup

	accidental_markFormCallback.CreationMode = (accidental_mark == nil)

	return
}

type Accidental_markFormCallback struct {
	accidental_mark *models.Accidental_mark

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (accidental_markFormCallback *Accidental_markFormCallback) OnSave() {

	log.Println("Accidental_markFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	accidental_markFormCallback.probe.formStage.Checkout()

	if accidental_markFormCallback.accidental_mark == nil {
		accidental_markFormCallback.accidental_mark = new(models.Accidental_mark).Stage(accidental_markFormCallback.probe.stageOfInterest)
	}
	accidental_mark_ := accidental_markFormCallback.accidental_mark
	_ = accidental_mark_

	for _, formDiv := range accidental_markFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(accidental_mark_.Name), formDiv)
		case "Ornaments:Accidental_mark":
			// we need to retrieve the field owner before the change
			var pastOrnamentsOwner *models.Ornaments
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Accidental_mark"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				accidental_markFormCallback.probe.stageOfInterest,
				accidental_markFormCallback.probe.backRepoOfInterest,
				accidental_mark_,
				&rf)

			if reverseFieldOwner != nil {
				pastOrnamentsOwner = reverseFieldOwner.(*models.Ornaments)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastOrnamentsOwner != nil {
					idx := slices.Index(pastOrnamentsOwner.Accidental_mark, accidental_mark_)
					pastOrnamentsOwner.Accidental_mark = slices.Delete(pastOrnamentsOwner.Accidental_mark, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _ornaments := range *models.GetGongstructInstancesSet[models.Ornaments](accidental_markFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _ornaments.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newOrnamentsOwner := _ornaments // we have a match
						if pastOrnamentsOwner != nil {
							if newOrnamentsOwner != pastOrnamentsOwner {
								idx := slices.Index(pastOrnamentsOwner.Accidental_mark, accidental_mark_)
								pastOrnamentsOwner.Accidental_mark = slices.Delete(pastOrnamentsOwner.Accidental_mark, idx, idx+1)
								newOrnamentsOwner.Accidental_mark = append(newOrnamentsOwner.Accidental_mark, accidental_mark_)
							}
						} else {
							newOrnamentsOwner.Accidental_mark = append(newOrnamentsOwner.Accidental_mark, accidental_mark_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if accidental_markFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accidental_mark_.Unstage(accidental_markFormCallback.probe.stageOfInterest)
	}

	accidental_markFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Accidental_mark](
		accidental_markFormCallback.probe,
	)
	accidental_markFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if accidental_markFormCallback.CreationMode || accidental_markFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accidental_markFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(accidental_markFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Accidental_markFormCallback(
			nil,
			accidental_markFormCallback.probe,
			newFormGroup,
		)
		accidental_mark := new(models.Accidental_mark)
		FillUpForm(accidental_mark, newFormGroup, accidental_markFormCallback.probe)
		accidental_markFormCallback.probe.formStage.Commit()
	}

	fillUpTree(accidental_markFormCallback.probe)
}
func __gong__New__Accidental_textFormCallback(
	accidental_text *models.Accidental_text,
	probe *Probe,
	formGroup *table.FormGroup,
) (accidental_textFormCallback *Accidental_textFormCallback) {
	accidental_textFormCallback = new(Accidental_textFormCallback)
	accidental_textFormCallback.probe = probe
	accidental_textFormCallback.accidental_text = accidental_text
	accidental_textFormCallback.formGroup = formGroup

	accidental_textFormCallback.CreationMode = (accidental_text == nil)

	return
}

type Accidental_textFormCallback struct {
	accidental_text *models.Accidental_text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (accidental_textFormCallback *Accidental_textFormCallback) OnSave() {

	log.Println("Accidental_textFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	accidental_textFormCallback.probe.formStage.Checkout()

	if accidental_textFormCallback.accidental_text == nil {
		accidental_textFormCallback.accidental_text = new(models.Accidental_text).Stage(accidental_textFormCallback.probe.stageOfInterest)
	}
	accidental_text_ := accidental_textFormCallback.accidental_text
	_ = accidental_text_

	for _, formDiv := range accidental_textFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(accidental_text_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if accidental_textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accidental_text_.Unstage(accidental_textFormCallback.probe.stageOfInterest)
	}

	accidental_textFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Accidental_text](
		accidental_textFormCallback.probe,
	)
	accidental_textFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if accidental_textFormCallback.CreationMode || accidental_textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accidental_textFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(accidental_textFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Accidental_textFormCallback(
			nil,
			accidental_textFormCallback.probe,
			newFormGroup,
		)
		accidental_text := new(models.Accidental_text)
		FillUpForm(accidental_text, newFormGroup, accidental_textFormCallback.probe)
		accidental_textFormCallback.probe.formStage.Commit()
	}

	fillUpTree(accidental_textFormCallback.probe)
}
func __gong__New__AccordFormCallback(
	accord *models.Accord,
	probe *Probe,
	formGroup *table.FormGroup,
) (accordFormCallback *AccordFormCallback) {
	accordFormCallback = new(AccordFormCallback)
	accordFormCallback.probe = probe
	accordFormCallback.accord = accord
	accordFormCallback.formGroup = formGroup

	accordFormCallback.CreationMode = (accord == nil)

	return
}

type AccordFormCallback struct {
	accord *models.Accord

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (accordFormCallback *AccordFormCallback) OnSave() {

	log.Println("AccordFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	accordFormCallback.probe.formStage.Checkout()

	if accordFormCallback.accord == nil {
		accordFormCallback.accord = new(models.Accord).Stage(accordFormCallback.probe.stageOfInterest)
	}
	accord_ := accordFormCallback.accord
	_ = accord_

	for _, formDiv := range accordFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(accord_.Name), formDiv)
		case "Scordatura:Accord":
			// we need to retrieve the field owner before the change
			var pastScordaturaOwner *models.Scordatura
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Scordatura"
			rf.Fieldname = "Accord"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				accordFormCallback.probe.stageOfInterest,
				accordFormCallback.probe.backRepoOfInterest,
				accord_,
				&rf)

			if reverseFieldOwner != nil {
				pastScordaturaOwner = reverseFieldOwner.(*models.Scordatura)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastScordaturaOwner != nil {
					idx := slices.Index(pastScordaturaOwner.Accord, accord_)
					pastScordaturaOwner.Accord = slices.Delete(pastScordaturaOwner.Accord, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _scordatura := range *models.GetGongstructInstancesSet[models.Scordatura](accordFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _scordatura.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newScordaturaOwner := _scordatura // we have a match
						if pastScordaturaOwner != nil {
							if newScordaturaOwner != pastScordaturaOwner {
								idx := slices.Index(pastScordaturaOwner.Accord, accord_)
								pastScordaturaOwner.Accord = slices.Delete(pastScordaturaOwner.Accord, idx, idx+1)
								newScordaturaOwner.Accord = append(newScordaturaOwner.Accord, accord_)
							}
						} else {
							newScordaturaOwner.Accord = append(newScordaturaOwner.Accord, accord_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if accordFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accord_.Unstage(accordFormCallback.probe.stageOfInterest)
	}

	accordFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Accord](
		accordFormCallback.probe,
	)
	accordFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if accordFormCallback.CreationMode || accordFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accordFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(accordFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AccordFormCallback(
			nil,
			accordFormCallback.probe,
			newFormGroup,
		)
		accord := new(models.Accord)
		FillUpForm(accord, newFormGroup, accordFormCallback.probe)
		accordFormCallback.probe.formStage.Commit()
	}

	fillUpTree(accordFormCallback.probe)
}
func __gong__New__Accordion_registrationFormCallback(
	accordion_registration *models.Accordion_registration,
	probe *Probe,
	formGroup *table.FormGroup,
) (accordion_registrationFormCallback *Accordion_registrationFormCallback) {
	accordion_registrationFormCallback = new(Accordion_registrationFormCallback)
	accordion_registrationFormCallback.probe = probe
	accordion_registrationFormCallback.accordion_registration = accordion_registration
	accordion_registrationFormCallback.formGroup = formGroup

	accordion_registrationFormCallback.CreationMode = (accordion_registration == nil)

	return
}

type Accordion_registrationFormCallback struct {
	accordion_registration *models.Accordion_registration

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (accordion_registrationFormCallback *Accordion_registrationFormCallback) OnSave() {

	log.Println("Accordion_registrationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	accordion_registrationFormCallback.probe.formStage.Checkout()

	if accordion_registrationFormCallback.accordion_registration == nil {
		accordion_registrationFormCallback.accordion_registration = new(models.Accordion_registration).Stage(accordion_registrationFormCallback.probe.stageOfInterest)
	}
	accordion_registration_ := accordion_registrationFormCallback.accordion_registration
	_ = accordion_registration_

	for _, formDiv := range accordion_registrationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(accordion_registration_.Name), formDiv)
		case "Accordion_high":
			FormDivSelectFieldToField(&(accordion_registration_.Accordion_high), accordion_registrationFormCallback.probe.stageOfInterest, formDiv)
		case "Accordion_low":
			FormDivSelectFieldToField(&(accordion_registration_.Accordion_low), accordion_registrationFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if accordion_registrationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accordion_registration_.Unstage(accordion_registrationFormCallback.probe.stageOfInterest)
	}

	accordion_registrationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Accordion_registration](
		accordion_registrationFormCallback.probe,
	)
	accordion_registrationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if accordion_registrationFormCallback.CreationMode || accordion_registrationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		accordion_registrationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(accordion_registrationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Accordion_registrationFormCallback(
			nil,
			accordion_registrationFormCallback.probe,
			newFormGroup,
		)
		accordion_registration := new(models.Accordion_registration)
		FillUpForm(accordion_registration, newFormGroup, accordion_registrationFormCallback.probe)
		accordion_registrationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(accordion_registrationFormCallback.probe)
}
func __gong__New__AnyTypeFormCallback(
	anytype *models.AnyType,
	probe *Probe,
	formGroup *table.FormGroup,
) (anytypeFormCallback *AnyTypeFormCallback) {
	anytypeFormCallback = new(AnyTypeFormCallback)
	anytypeFormCallback.probe = probe
	anytypeFormCallback.anytype = anytype
	anytypeFormCallback.formGroup = formGroup

	anytypeFormCallback.CreationMode = (anytype == nil)

	return
}

type AnyTypeFormCallback struct {
	anytype *models.AnyType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (anytypeFormCallback *AnyTypeFormCallback) OnSave() {

	log.Println("AnyTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	anytypeFormCallback.probe.formStage.Checkout()

	if anytypeFormCallback.anytype == nil {
		anytypeFormCallback.anytype = new(models.AnyType).Stage(anytypeFormCallback.probe.stageOfInterest)
	}
	anytype_ := anytypeFormCallback.anytype
	_ = anytype_

	for _, formDiv := range anytypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(anytype_.Name), formDiv)
		case "InnerXML":
			FormDivBasicFieldToField(&(anytype_.InnerXML), formDiv)
		}
	}

	// manage the suppress operation
	if anytypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		anytype_.Unstage(anytypeFormCallback.probe.stageOfInterest)
	}

	anytypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AnyType](
		anytypeFormCallback.probe,
	)
	anytypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if anytypeFormCallback.CreationMode || anytypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		anytypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(anytypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnyTypeFormCallback(
			nil,
			anytypeFormCallback.probe,
			newFormGroup,
		)
		anytype := new(models.AnyType)
		FillUpForm(anytype, newFormGroup, anytypeFormCallback.probe)
		anytypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(anytypeFormCallback.probe)
}
func __gong__New__AppearanceFormCallback(
	appearance *models.Appearance,
	probe *Probe,
	formGroup *table.FormGroup,
) (appearanceFormCallback *AppearanceFormCallback) {
	appearanceFormCallback = new(AppearanceFormCallback)
	appearanceFormCallback.probe = probe
	appearanceFormCallback.appearance = appearance
	appearanceFormCallback.formGroup = formGroup

	appearanceFormCallback.CreationMode = (appearance == nil)

	return
}

type AppearanceFormCallback struct {
	appearance *models.Appearance

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (appearanceFormCallback *AppearanceFormCallback) OnSave() {

	log.Println("AppearanceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	appearanceFormCallback.probe.formStage.Checkout()

	if appearanceFormCallback.appearance == nil {
		appearanceFormCallback.appearance = new(models.Appearance).Stage(appearanceFormCallback.probe.stageOfInterest)
	}
	appearance_ := appearanceFormCallback.appearance
	_ = appearance_

	for _, formDiv := range appearanceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(appearance_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if appearanceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		appearance_.Unstage(appearanceFormCallback.probe.stageOfInterest)
	}

	appearanceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Appearance](
		appearanceFormCallback.probe,
	)
	appearanceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if appearanceFormCallback.CreationMode || appearanceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		appearanceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(appearanceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AppearanceFormCallback(
			nil,
			appearanceFormCallback.probe,
			newFormGroup,
		)
		appearance := new(models.Appearance)
		FillUpForm(appearance, newFormGroup, appearanceFormCallback.probe)
		appearanceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(appearanceFormCallback.probe)
}
func __gong__New__ArpeggiateFormCallback(
	arpeggiate *models.Arpeggiate,
	probe *Probe,
	formGroup *table.FormGroup,
) (arpeggiateFormCallback *ArpeggiateFormCallback) {
	arpeggiateFormCallback = new(ArpeggiateFormCallback)
	arpeggiateFormCallback.probe = probe
	arpeggiateFormCallback.arpeggiate = arpeggiate
	arpeggiateFormCallback.formGroup = formGroup

	arpeggiateFormCallback.CreationMode = (arpeggiate == nil)

	return
}

type ArpeggiateFormCallback struct {
	arpeggiate *models.Arpeggiate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (arpeggiateFormCallback *ArpeggiateFormCallback) OnSave() {

	log.Println("ArpeggiateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	arpeggiateFormCallback.probe.formStage.Checkout()

	if arpeggiateFormCallback.arpeggiate == nil {
		arpeggiateFormCallback.arpeggiate = new(models.Arpeggiate).Stage(arpeggiateFormCallback.probe.stageOfInterest)
	}
	arpeggiate_ := arpeggiateFormCallback.arpeggiate
	_ = arpeggiate_

	for _, formDiv := range arpeggiateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(arpeggiate_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if arpeggiateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arpeggiate_.Unstage(arpeggiateFormCallback.probe.stageOfInterest)
	}

	arpeggiateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Arpeggiate](
		arpeggiateFormCallback.probe,
	)
	arpeggiateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if arpeggiateFormCallback.CreationMode || arpeggiateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arpeggiateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(arpeggiateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArpeggiateFormCallback(
			nil,
			arpeggiateFormCallback.probe,
			newFormGroup,
		)
		arpeggiate := new(models.Arpeggiate)
		FillUpForm(arpeggiate, newFormGroup, arpeggiateFormCallback.probe)
		arpeggiateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(arpeggiateFormCallback.probe)
}
func __gong__New__ArrowFormCallback(
	arrow *models.Arrow,
	probe *Probe,
	formGroup *table.FormGroup,
) (arrowFormCallback *ArrowFormCallback) {
	arrowFormCallback = new(ArrowFormCallback)
	arrowFormCallback.probe = probe
	arrowFormCallback.arrow = arrow
	arrowFormCallback.formGroup = formGroup

	arrowFormCallback.CreationMode = (arrow == nil)

	return
}

type ArrowFormCallback struct {
	arrow *models.Arrow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (arrowFormCallback *ArrowFormCallback) OnSave() {

	log.Println("ArrowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	arrowFormCallback.probe.formStage.Checkout()

	if arrowFormCallback.arrow == nil {
		arrowFormCallback.arrow = new(models.Arrow).Stage(arrowFormCallback.probe.stageOfInterest)
	}
	arrow_ := arrowFormCallback.arrow
	_ = arrow_

	for _, formDiv := range arrowFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(arrow_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if arrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arrow_.Unstage(arrowFormCallback.probe.stageOfInterest)
	}

	arrowFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Arrow](
		arrowFormCallback.probe,
	)
	arrowFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if arrowFormCallback.CreationMode || arrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arrowFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(arrowFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArrowFormCallback(
			nil,
			arrowFormCallback.probe,
			newFormGroup,
		)
		arrow := new(models.Arrow)
		FillUpForm(arrow, newFormGroup, arrowFormCallback.probe)
		arrowFormCallback.probe.formStage.Commit()
	}

	fillUpTree(arrowFormCallback.probe)
}
func __gong__New__ArticulationsFormCallback(
	articulations *models.Articulations,
	probe *Probe,
	formGroup *table.FormGroup,
) (articulationsFormCallback *ArticulationsFormCallback) {
	articulationsFormCallback = new(ArticulationsFormCallback)
	articulationsFormCallback.probe = probe
	articulationsFormCallback.articulations = articulations
	articulationsFormCallback.formGroup = formGroup

	articulationsFormCallback.CreationMode = (articulations == nil)

	return
}

type ArticulationsFormCallback struct {
	articulations *models.Articulations

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (articulationsFormCallback *ArticulationsFormCallback) OnSave() {

	log.Println("ArticulationsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	articulationsFormCallback.probe.formStage.Checkout()

	if articulationsFormCallback.articulations == nil {
		articulationsFormCallback.articulations = new(models.Articulations).Stage(articulationsFormCallback.probe.stageOfInterest)
	}
	articulations_ := articulationsFormCallback.articulations
	_ = articulations_

	for _, formDiv := range articulationsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(articulations_.Name), formDiv)
		case "Accent":
			FormDivSelectFieldToField(&(articulations_.Accent), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Strong_accent":
			FormDivSelectFieldToField(&(articulations_.Strong_accent), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Staccato":
			FormDivSelectFieldToField(&(articulations_.Staccato), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Tenuto":
			FormDivSelectFieldToField(&(articulations_.Tenuto), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Detached_legato":
			FormDivSelectFieldToField(&(articulations_.Detached_legato), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Staccatissimo":
			FormDivSelectFieldToField(&(articulations_.Staccatissimo), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Spiccato":
			FormDivSelectFieldToField(&(articulations_.Spiccato), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Scoop":
			FormDivSelectFieldToField(&(articulations_.Scoop), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Plop":
			FormDivSelectFieldToField(&(articulations_.Plop), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Doit":
			FormDivSelectFieldToField(&(articulations_.Doit), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Falloff":
			FormDivSelectFieldToField(&(articulations_.Falloff), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Breath_mark":
			FormDivSelectFieldToField(&(articulations_.Breath_mark), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Caesura":
			FormDivSelectFieldToField(&(articulations_.Caesura), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Stress":
			FormDivSelectFieldToField(&(articulations_.Stress), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Unstress":
			FormDivSelectFieldToField(&(articulations_.Unstress), articulationsFormCallback.probe.stageOfInterest, formDiv)
		case "Soft_accent":
			FormDivSelectFieldToField(&(articulations_.Soft_accent), articulationsFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if articulationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		articulations_.Unstage(articulationsFormCallback.probe.stageOfInterest)
	}

	articulationsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Articulations](
		articulationsFormCallback.probe,
	)
	articulationsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if articulationsFormCallback.CreationMode || articulationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		articulationsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(articulationsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArticulationsFormCallback(
			nil,
			articulationsFormCallback.probe,
			newFormGroup,
		)
		articulations := new(models.Articulations)
		FillUpForm(articulations, newFormGroup, articulationsFormCallback.probe)
		articulationsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(articulationsFormCallback.probe)
}
func __gong__New__AssessFormCallback(
	assess *models.Assess,
	probe *Probe,
	formGroup *table.FormGroup,
) (assessFormCallback *AssessFormCallback) {
	assessFormCallback = new(AssessFormCallback)
	assessFormCallback.probe = probe
	assessFormCallback.assess = assess
	assessFormCallback.formGroup = formGroup

	assessFormCallback.CreationMode = (assess == nil)

	return
}

type AssessFormCallback struct {
	assess *models.Assess

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (assessFormCallback *AssessFormCallback) OnSave() {

	log.Println("AssessFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	assessFormCallback.probe.formStage.Checkout()

	if assessFormCallback.assess == nil {
		assessFormCallback.assess = new(models.Assess).Stage(assessFormCallback.probe.stageOfInterest)
	}
	assess_ := assessFormCallback.assess
	_ = assess_

	for _, formDiv := range assessFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(assess_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if assessFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assess_.Unstage(assessFormCallback.probe.stageOfInterest)
	}

	assessFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Assess](
		assessFormCallback.probe,
	)
	assessFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if assessFormCallback.CreationMode || assessFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assessFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(assessFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AssessFormCallback(
			nil,
			assessFormCallback.probe,
			newFormGroup,
		)
		assess := new(models.Assess)
		FillUpForm(assess, newFormGroup, assessFormCallback.probe)
		assessFormCallback.probe.formStage.Commit()
	}

	fillUpTree(assessFormCallback.probe)
}
func __gong__New__AttributesFormCallback(
	attributes *models.Attributes,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributesFormCallback *AttributesFormCallback) {
	attributesFormCallback = new(AttributesFormCallback)
	attributesFormCallback.probe = probe
	attributesFormCallback.attributes = attributes
	attributesFormCallback.formGroup = formGroup

	attributesFormCallback.CreationMode = (attributes == nil)

	return
}

type AttributesFormCallback struct {
	attributes *models.Attributes

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributesFormCallback *AttributesFormCallback) OnSave() {

	log.Println("AttributesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributesFormCallback.probe.formStage.Checkout()

	if attributesFormCallback.attributes == nil {
		attributesFormCallback.attributes = new(models.Attributes).Stage(attributesFormCallback.probe.stageOfInterest)
	}
	attributes_ := attributesFormCallback.attributes
	_ = attributes_

	for _, formDiv := range attributesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributes_.Name), formDiv)
		case "Part_symbol":
			FormDivSelectFieldToField(&(attributes_.Part_symbol), attributesFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if attributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributes_.Unstage(attributesFormCallback.probe.stageOfInterest)
	}

	attributesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Attributes](
		attributesFormCallback.probe,
	)
	attributesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributesFormCallback.CreationMode || attributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AttributesFormCallback(
			nil,
			attributesFormCallback.probe,
			newFormGroup,
		)
		attributes := new(models.Attributes)
		FillUpForm(attributes, newFormGroup, attributesFormCallback.probe)
		attributesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributesFormCallback.probe)
}
func __gong__New__BackupFormCallback(
	backup *models.Backup,
	probe *Probe,
	formGroup *table.FormGroup,
) (backupFormCallback *BackupFormCallback) {
	backupFormCallback = new(BackupFormCallback)
	backupFormCallback.probe = probe
	backupFormCallback.backup = backup
	backupFormCallback.formGroup = formGroup

	backupFormCallback.CreationMode = (backup == nil)

	return
}

type BackupFormCallback struct {
	backup *models.Backup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (backupFormCallback *BackupFormCallback) OnSave() {

	log.Println("BackupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	backupFormCallback.probe.formStage.Checkout()

	if backupFormCallback.backup == nil {
		backupFormCallback.backup = new(models.Backup).Stage(backupFormCallback.probe.stageOfInterest)
	}
	backup_ := backupFormCallback.backup
	_ = backup_

	for _, formDiv := range backupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(backup_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if backupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		backup_.Unstage(backupFormCallback.probe.stageOfInterest)
	}

	backupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Backup](
		backupFormCallback.probe,
	)
	backupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if backupFormCallback.CreationMode || backupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		backupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(backupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BackupFormCallback(
			nil,
			backupFormCallback.probe,
			newFormGroup,
		)
		backup := new(models.Backup)
		FillUpForm(backup, newFormGroup, backupFormCallback.probe)
		backupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(backupFormCallback.probe)
}
func __gong__New__Bar_style_colorFormCallback(
	bar_style_color *models.Bar_style_color,
	probe *Probe,
	formGroup *table.FormGroup,
) (bar_style_colorFormCallback *Bar_style_colorFormCallback) {
	bar_style_colorFormCallback = new(Bar_style_colorFormCallback)
	bar_style_colorFormCallback.probe = probe
	bar_style_colorFormCallback.bar_style_color = bar_style_color
	bar_style_colorFormCallback.formGroup = formGroup

	bar_style_colorFormCallback.CreationMode = (bar_style_color == nil)

	return
}

type Bar_style_colorFormCallback struct {
	bar_style_color *models.Bar_style_color

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bar_style_colorFormCallback *Bar_style_colorFormCallback) OnSave() {

	log.Println("Bar_style_colorFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bar_style_colorFormCallback.probe.formStage.Checkout()

	if bar_style_colorFormCallback.bar_style_color == nil {
		bar_style_colorFormCallback.bar_style_color = new(models.Bar_style_color).Stage(bar_style_colorFormCallback.probe.stageOfInterest)
	}
	bar_style_color_ := bar_style_colorFormCallback.bar_style_color
	_ = bar_style_color_

	for _, formDiv := range bar_style_colorFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bar_style_color_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if bar_style_colorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bar_style_color_.Unstage(bar_style_colorFormCallback.probe.stageOfInterest)
	}

	bar_style_colorFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bar_style_color](
		bar_style_colorFormCallback.probe,
	)
	bar_style_colorFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bar_style_colorFormCallback.CreationMode || bar_style_colorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bar_style_colorFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(bar_style_colorFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Bar_style_colorFormCallback(
			nil,
			bar_style_colorFormCallback.probe,
			newFormGroup,
		)
		bar_style_color := new(models.Bar_style_color)
		FillUpForm(bar_style_color, newFormGroup, bar_style_colorFormCallback.probe)
		bar_style_colorFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bar_style_colorFormCallback.probe)
}
func __gong__New__BarlineFormCallback(
	barline *models.Barline,
	probe *Probe,
	formGroup *table.FormGroup,
) (barlineFormCallback *BarlineFormCallback) {
	barlineFormCallback = new(BarlineFormCallback)
	barlineFormCallback.probe = probe
	barlineFormCallback.barline = barline
	barlineFormCallback.formGroup = formGroup

	barlineFormCallback.CreationMode = (barline == nil)

	return
}

type BarlineFormCallback struct {
	barline *models.Barline

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (barlineFormCallback *BarlineFormCallback) OnSave() {

	log.Println("BarlineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	barlineFormCallback.probe.formStage.Checkout()

	if barlineFormCallback.barline == nil {
		barlineFormCallback.barline = new(models.Barline).Stage(barlineFormCallback.probe.stageOfInterest)
	}
	barline_ := barlineFormCallback.barline
	_ = barline_

	for _, formDiv := range barlineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(barline_.Name), formDiv)
		case "Bar_style":
			FormDivSelectFieldToField(&(barline_.Bar_style), barlineFormCallback.probe.stageOfInterest, formDiv)
		case "Wavy_line":
			FormDivSelectFieldToField(&(barline_.Wavy_line), barlineFormCallback.probe.stageOfInterest, formDiv)
		case "Fermata":
			FormDivSelectFieldToField(&(barline_.Fermata), barlineFormCallback.probe.stageOfInterest, formDiv)
		case "Ending":
			FormDivSelectFieldToField(&(barline_.Ending), barlineFormCallback.probe.stageOfInterest, formDiv)
		case "Repeat":
			FormDivSelectFieldToField(&(barline_.Repeat), barlineFormCallback.probe.stageOfInterest, formDiv)
		case "Segno":
			FormDivBasicFieldToField(&(barline_.Segno), formDiv)
		case "Coda":
			FormDivBasicFieldToField(&(barline_.Coda), formDiv)
		}
	}

	// manage the suppress operation
	if barlineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		barline_.Unstage(barlineFormCallback.probe.stageOfInterest)
	}

	barlineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Barline](
		barlineFormCallback.probe,
	)
	barlineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if barlineFormCallback.CreationMode || barlineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		barlineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(barlineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BarlineFormCallback(
			nil,
			barlineFormCallback.probe,
			newFormGroup,
		)
		barline := new(models.Barline)
		FillUpForm(barline, newFormGroup, barlineFormCallback.probe)
		barlineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(barlineFormCallback.probe)
}
func __gong__New__BarreFormCallback(
	barre *models.Barre,
	probe *Probe,
	formGroup *table.FormGroup,
) (barreFormCallback *BarreFormCallback) {
	barreFormCallback = new(BarreFormCallback)
	barreFormCallback.probe = probe
	barreFormCallback.barre = barre
	barreFormCallback.formGroup = formGroup

	barreFormCallback.CreationMode = (barre == nil)

	return
}

type BarreFormCallback struct {
	barre *models.Barre

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (barreFormCallback *BarreFormCallback) OnSave() {

	log.Println("BarreFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	barreFormCallback.probe.formStage.Checkout()

	if barreFormCallback.barre == nil {
		barreFormCallback.barre = new(models.Barre).Stage(barreFormCallback.probe.stageOfInterest)
	}
	barre_ := barreFormCallback.barre
	_ = barre_

	for _, formDiv := range barreFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(barre_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if barreFormCallback.formGroup.HasSuppressButtonBeenPressed {
		barre_.Unstage(barreFormCallback.probe.stageOfInterest)
	}

	barreFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Barre](
		barreFormCallback.probe,
	)
	barreFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if barreFormCallback.CreationMode || barreFormCallback.formGroup.HasSuppressButtonBeenPressed {
		barreFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(barreFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BarreFormCallback(
			nil,
			barreFormCallback.probe,
			newFormGroup,
		)
		barre := new(models.Barre)
		FillUpForm(barre, newFormGroup, barreFormCallback.probe)
		barreFormCallback.probe.formStage.Commit()
	}

	fillUpTree(barreFormCallback.probe)
}
func __gong__New__BassFormCallback(
	bass *models.Bass,
	probe *Probe,
	formGroup *table.FormGroup,
) (bassFormCallback *BassFormCallback) {
	bassFormCallback = new(BassFormCallback)
	bassFormCallback.probe = probe
	bassFormCallback.bass = bass
	bassFormCallback.formGroup = formGroup

	bassFormCallback.CreationMode = (bass == nil)

	return
}

type BassFormCallback struct {
	bass *models.Bass

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bassFormCallback *BassFormCallback) OnSave() {

	log.Println("BassFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bassFormCallback.probe.formStage.Checkout()

	if bassFormCallback.bass == nil {
		bassFormCallback.bass = new(models.Bass).Stage(bassFormCallback.probe.stageOfInterest)
	}
	bass_ := bassFormCallback.bass
	_ = bass_

	for _, formDiv := range bassFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bass_.Name), formDiv)
		case "Bass_step":
			FormDivSelectFieldToField(&(bass_.Bass_step), bassFormCallback.probe.stageOfInterest, formDiv)
		case "Bass_alter":
			FormDivSelectFieldToField(&(bass_.Bass_alter), bassFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if bassFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bass_.Unstage(bassFormCallback.probe.stageOfInterest)
	}

	bassFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bass](
		bassFormCallback.probe,
	)
	bassFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bassFormCallback.CreationMode || bassFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bassFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(bassFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BassFormCallback(
			nil,
			bassFormCallback.probe,
			newFormGroup,
		)
		bass := new(models.Bass)
		FillUpForm(bass, newFormGroup, bassFormCallback.probe)
		bassFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bassFormCallback.probe)
}
func __gong__New__Bass_stepFormCallback(
	bass_step *models.Bass_step,
	probe *Probe,
	formGroup *table.FormGroup,
) (bass_stepFormCallback *Bass_stepFormCallback) {
	bass_stepFormCallback = new(Bass_stepFormCallback)
	bass_stepFormCallback.probe = probe
	bass_stepFormCallback.bass_step = bass_step
	bass_stepFormCallback.formGroup = formGroup

	bass_stepFormCallback.CreationMode = (bass_step == nil)

	return
}

type Bass_stepFormCallback struct {
	bass_step *models.Bass_step

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bass_stepFormCallback *Bass_stepFormCallback) OnSave() {

	log.Println("Bass_stepFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bass_stepFormCallback.probe.formStage.Checkout()

	if bass_stepFormCallback.bass_step == nil {
		bass_stepFormCallback.bass_step = new(models.Bass_step).Stage(bass_stepFormCallback.probe.stageOfInterest)
	}
	bass_step_ := bass_stepFormCallback.bass_step
	_ = bass_step_

	for _, formDiv := range bass_stepFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bass_step_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(bass_step_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if bass_stepFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bass_step_.Unstage(bass_stepFormCallback.probe.stageOfInterest)
	}

	bass_stepFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bass_step](
		bass_stepFormCallback.probe,
	)
	bass_stepFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bass_stepFormCallback.CreationMode || bass_stepFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bass_stepFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(bass_stepFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Bass_stepFormCallback(
			nil,
			bass_stepFormCallback.probe,
			newFormGroup,
		)
		bass_step := new(models.Bass_step)
		FillUpForm(bass_step, newFormGroup, bass_stepFormCallback.probe)
		bass_stepFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bass_stepFormCallback.probe)
}
func __gong__New__BeamFormCallback(
	beam *models.Beam,
	probe *Probe,
	formGroup *table.FormGroup,
) (beamFormCallback *BeamFormCallback) {
	beamFormCallback = new(BeamFormCallback)
	beamFormCallback.probe = probe
	beamFormCallback.beam = beam
	beamFormCallback.formGroup = formGroup

	beamFormCallback.CreationMode = (beam == nil)

	return
}

type BeamFormCallback struct {
	beam *models.Beam

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (beamFormCallback *BeamFormCallback) OnSave() {

	log.Println("BeamFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	beamFormCallback.probe.formStage.Checkout()

	if beamFormCallback.beam == nil {
		beamFormCallback.beam = new(models.Beam).Stage(beamFormCallback.probe.stageOfInterest)
	}
	beam_ := beamFormCallback.beam
	_ = beam_

	for _, formDiv := range beamFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(beam_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if beamFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beam_.Unstage(beamFormCallback.probe.stageOfInterest)
	}

	beamFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Beam](
		beamFormCallback.probe,
	)
	beamFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beamFormCallback.CreationMode || beamFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beamFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(beamFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BeamFormCallback(
			nil,
			beamFormCallback.probe,
			newFormGroup,
		)
		beam := new(models.Beam)
		FillUpForm(beam, newFormGroup, beamFormCallback.probe)
		beamFormCallback.probe.formStage.Commit()
	}

	fillUpTree(beamFormCallback.probe)
}
func __gong__New__Beat_repeatFormCallback(
	beat_repeat *models.Beat_repeat,
	probe *Probe,
	formGroup *table.FormGroup,
) (beat_repeatFormCallback *Beat_repeatFormCallback) {
	beat_repeatFormCallback = new(Beat_repeatFormCallback)
	beat_repeatFormCallback.probe = probe
	beat_repeatFormCallback.beat_repeat = beat_repeat
	beat_repeatFormCallback.formGroup = formGroup

	beat_repeatFormCallback.CreationMode = (beat_repeat == nil)

	return
}

type Beat_repeatFormCallback struct {
	beat_repeat *models.Beat_repeat

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (beat_repeatFormCallback *Beat_repeatFormCallback) OnSave() {

	log.Println("Beat_repeatFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	beat_repeatFormCallback.probe.formStage.Checkout()

	if beat_repeatFormCallback.beat_repeat == nil {
		beat_repeatFormCallback.beat_repeat = new(models.Beat_repeat).Stage(beat_repeatFormCallback.probe.stageOfInterest)
	}
	beat_repeat_ := beat_repeatFormCallback.beat_repeat
	_ = beat_repeat_

	for _, formDiv := range beat_repeatFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(beat_repeat_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if beat_repeatFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beat_repeat_.Unstage(beat_repeatFormCallback.probe.stageOfInterest)
	}

	beat_repeatFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Beat_repeat](
		beat_repeatFormCallback.probe,
	)
	beat_repeatFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beat_repeatFormCallback.CreationMode || beat_repeatFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beat_repeatFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(beat_repeatFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Beat_repeatFormCallback(
			nil,
			beat_repeatFormCallback.probe,
			newFormGroup,
		)
		beat_repeat := new(models.Beat_repeat)
		FillUpForm(beat_repeat, newFormGroup, beat_repeatFormCallback.probe)
		beat_repeatFormCallback.probe.formStage.Commit()
	}

	fillUpTree(beat_repeatFormCallback.probe)
}
func __gong__New__Beat_unit_tiedFormCallback(
	beat_unit_tied *models.Beat_unit_tied,
	probe *Probe,
	formGroup *table.FormGroup,
) (beat_unit_tiedFormCallback *Beat_unit_tiedFormCallback) {
	beat_unit_tiedFormCallback = new(Beat_unit_tiedFormCallback)
	beat_unit_tiedFormCallback.probe = probe
	beat_unit_tiedFormCallback.beat_unit_tied = beat_unit_tied
	beat_unit_tiedFormCallback.formGroup = formGroup

	beat_unit_tiedFormCallback.CreationMode = (beat_unit_tied == nil)

	return
}

type Beat_unit_tiedFormCallback struct {
	beat_unit_tied *models.Beat_unit_tied

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (beat_unit_tiedFormCallback *Beat_unit_tiedFormCallback) OnSave() {

	log.Println("Beat_unit_tiedFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	beat_unit_tiedFormCallback.probe.formStage.Checkout()

	if beat_unit_tiedFormCallback.beat_unit_tied == nil {
		beat_unit_tiedFormCallback.beat_unit_tied = new(models.Beat_unit_tied).Stage(beat_unit_tiedFormCallback.probe.stageOfInterest)
	}
	beat_unit_tied_ := beat_unit_tiedFormCallback.beat_unit_tied
	_ = beat_unit_tied_

	for _, formDiv := range beat_unit_tiedFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(beat_unit_tied_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if beat_unit_tiedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beat_unit_tied_.Unstage(beat_unit_tiedFormCallback.probe.stageOfInterest)
	}

	beat_unit_tiedFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Beat_unit_tied](
		beat_unit_tiedFormCallback.probe,
	)
	beat_unit_tiedFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beat_unit_tiedFormCallback.CreationMode || beat_unit_tiedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beat_unit_tiedFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(beat_unit_tiedFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Beat_unit_tiedFormCallback(
			nil,
			beat_unit_tiedFormCallback.probe,
			newFormGroup,
		)
		beat_unit_tied := new(models.Beat_unit_tied)
		FillUpForm(beat_unit_tied, newFormGroup, beat_unit_tiedFormCallback.probe)
		beat_unit_tiedFormCallback.probe.formStage.Commit()
	}

	fillUpTree(beat_unit_tiedFormCallback.probe)
}
func __gong__New__BeaterFormCallback(
	beater *models.Beater,
	probe *Probe,
	formGroup *table.FormGroup,
) (beaterFormCallback *BeaterFormCallback) {
	beaterFormCallback = new(BeaterFormCallback)
	beaterFormCallback.probe = probe
	beaterFormCallback.beater = beater
	beaterFormCallback.formGroup = formGroup

	beaterFormCallback.CreationMode = (beater == nil)

	return
}

type BeaterFormCallback struct {
	beater *models.Beater

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (beaterFormCallback *BeaterFormCallback) OnSave() {

	log.Println("BeaterFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	beaterFormCallback.probe.formStage.Checkout()

	if beaterFormCallback.beater == nil {
		beaterFormCallback.beater = new(models.Beater).Stage(beaterFormCallback.probe.stageOfInterest)
	}
	beater_ := beaterFormCallback.beater
	_ = beater_

	for _, formDiv := range beaterFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(beater_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if beaterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beater_.Unstage(beaterFormCallback.probe.stageOfInterest)
	}

	beaterFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Beater](
		beaterFormCallback.probe,
	)
	beaterFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if beaterFormCallback.CreationMode || beaterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		beaterFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(beaterFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BeaterFormCallback(
			nil,
			beaterFormCallback.probe,
			newFormGroup,
		)
		beater := new(models.Beater)
		FillUpForm(beater, newFormGroup, beaterFormCallback.probe)
		beaterFormCallback.probe.formStage.Commit()
	}

	fillUpTree(beaterFormCallback.probe)
}
func __gong__New__BendFormCallback(
	bend *models.Bend,
	probe *Probe,
	formGroup *table.FormGroup,
) (bendFormCallback *BendFormCallback) {
	bendFormCallback = new(BendFormCallback)
	bendFormCallback.probe = probe
	bendFormCallback.bend = bend
	bendFormCallback.formGroup = formGroup

	bendFormCallback.CreationMode = (bend == nil)

	return
}

type BendFormCallback struct {
	bend *models.Bend

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bendFormCallback *BendFormCallback) OnSave() {

	log.Println("BendFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bendFormCallback.probe.formStage.Checkout()

	if bendFormCallback.bend == nil {
		bendFormCallback.bend = new(models.Bend).Stage(bendFormCallback.probe.stageOfInterest)
	}
	bend_ := bendFormCallback.bend
	_ = bend_

	for _, formDiv := range bendFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bend_.Name), formDiv)
		case "Pre_bend":
			FormDivSelectFieldToField(&(bend_.Pre_bend), bendFormCallback.probe.stageOfInterest, formDiv)
		case "Release":
			FormDivSelectFieldToField(&(bend_.Release), bendFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if bendFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bend_.Unstage(bendFormCallback.probe.stageOfInterest)
	}

	bendFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bend](
		bendFormCallback.probe,
	)
	bendFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bendFormCallback.CreationMode || bendFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bendFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(bendFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BendFormCallback(
			nil,
			bendFormCallback.probe,
			newFormGroup,
		)
		bend := new(models.Bend)
		FillUpForm(bend, newFormGroup, bendFormCallback.probe)
		bendFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bendFormCallback.probe)
}
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
		case "Credit:Bookmark":
			// we need to retrieve the field owner before the change
			var pastCreditOwner *models.Credit
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Credit"
			rf.Fieldname = "Bookmark"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				bookmarkFormCallback.probe.stageOfInterest,
				bookmarkFormCallback.probe.backRepoOfInterest,
				bookmark_,
				&rf)

			if reverseFieldOwner != nil {
				pastCreditOwner = reverseFieldOwner.(*models.Credit)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastCreditOwner != nil {
					idx := slices.Index(pastCreditOwner.Bookmark, bookmark_)
					pastCreditOwner.Bookmark = slices.Delete(pastCreditOwner.Bookmark, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _credit := range *models.GetGongstructInstancesSet[models.Credit](bookmarkFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _credit.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newCreditOwner := _credit // we have a match
						if pastCreditOwner != nil {
							if newCreditOwner != pastCreditOwner {
								idx := slices.Index(pastCreditOwner.Bookmark, bookmark_)
								pastCreditOwner.Bookmark = slices.Delete(pastCreditOwner.Bookmark, idx, idx+1)
								newCreditOwner.Bookmark = append(newCreditOwner.Bookmark, bookmark_)
							}
						} else {
							newCreditOwner.Bookmark = append(newCreditOwner.Bookmark, bookmark_)
						}
					}
				}
			}
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
func __gong__New__BracketFormCallback(
	bracket *models.Bracket,
	probe *Probe,
	formGroup *table.FormGroup,
) (bracketFormCallback *BracketFormCallback) {
	bracketFormCallback = new(BracketFormCallback)
	bracketFormCallback.probe = probe
	bracketFormCallback.bracket = bracket
	bracketFormCallback.formGroup = formGroup

	bracketFormCallback.CreationMode = (bracket == nil)

	return
}

type BracketFormCallback struct {
	bracket *models.Bracket

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bracketFormCallback *BracketFormCallback) OnSave() {

	log.Println("BracketFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bracketFormCallback.probe.formStage.Checkout()

	if bracketFormCallback.bracket == nil {
		bracketFormCallback.bracket = new(models.Bracket).Stage(bracketFormCallback.probe.stageOfInterest)
	}
	bracket_ := bracketFormCallback.bracket
	_ = bracket_

	for _, formDiv := range bracketFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bracket_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if bracketFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bracket_.Unstage(bracketFormCallback.probe.stageOfInterest)
	}

	bracketFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bracket](
		bracketFormCallback.probe,
	)
	bracketFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bracketFormCallback.CreationMode || bracketFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bracketFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(bracketFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BracketFormCallback(
			nil,
			bracketFormCallback.probe,
			newFormGroup,
		)
		bracket := new(models.Bracket)
		FillUpForm(bracket, newFormGroup, bracketFormCallback.probe)
		bracketFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bracketFormCallback.probe)
}
func __gong__New__Breath_markFormCallback(
	breath_mark *models.Breath_mark,
	probe *Probe,
	formGroup *table.FormGroup,
) (breath_markFormCallback *Breath_markFormCallback) {
	breath_markFormCallback = new(Breath_markFormCallback)
	breath_markFormCallback.probe = probe
	breath_markFormCallback.breath_mark = breath_mark
	breath_markFormCallback.formGroup = formGroup

	breath_markFormCallback.CreationMode = (breath_mark == nil)

	return
}

type Breath_markFormCallback struct {
	breath_mark *models.Breath_mark

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (breath_markFormCallback *Breath_markFormCallback) OnSave() {

	log.Println("Breath_markFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	breath_markFormCallback.probe.formStage.Checkout()

	if breath_markFormCallback.breath_mark == nil {
		breath_markFormCallback.breath_mark = new(models.Breath_mark).Stage(breath_markFormCallback.probe.stageOfInterest)
	}
	breath_mark_ := breath_markFormCallback.breath_mark
	_ = breath_mark_

	for _, formDiv := range breath_markFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(breath_mark_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if breath_markFormCallback.formGroup.HasSuppressButtonBeenPressed {
		breath_mark_.Unstage(breath_markFormCallback.probe.stageOfInterest)
	}

	breath_markFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Breath_mark](
		breath_markFormCallback.probe,
	)
	breath_markFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if breath_markFormCallback.CreationMode || breath_markFormCallback.formGroup.HasSuppressButtonBeenPressed {
		breath_markFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(breath_markFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Breath_markFormCallback(
			nil,
			breath_markFormCallback.probe,
			newFormGroup,
		)
		breath_mark := new(models.Breath_mark)
		FillUpForm(breath_mark, newFormGroup, breath_markFormCallback.probe)
		breath_markFormCallback.probe.formStage.Commit()
	}

	fillUpTree(breath_markFormCallback.probe)
}
func __gong__New__CaesuraFormCallback(
	caesura *models.Caesura,
	probe *Probe,
	formGroup *table.FormGroup,
) (caesuraFormCallback *CaesuraFormCallback) {
	caesuraFormCallback = new(CaesuraFormCallback)
	caesuraFormCallback.probe = probe
	caesuraFormCallback.caesura = caesura
	caesuraFormCallback.formGroup = formGroup

	caesuraFormCallback.CreationMode = (caesura == nil)

	return
}

type CaesuraFormCallback struct {
	caesura *models.Caesura

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (caesuraFormCallback *CaesuraFormCallback) OnSave() {

	log.Println("CaesuraFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	caesuraFormCallback.probe.formStage.Checkout()

	if caesuraFormCallback.caesura == nil {
		caesuraFormCallback.caesura = new(models.Caesura).Stage(caesuraFormCallback.probe.stageOfInterest)
	}
	caesura_ := caesuraFormCallback.caesura
	_ = caesura_

	for _, formDiv := range caesuraFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(caesura_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if caesuraFormCallback.formGroup.HasSuppressButtonBeenPressed {
		caesura_.Unstage(caesuraFormCallback.probe.stageOfInterest)
	}

	caesuraFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Caesura](
		caesuraFormCallback.probe,
	)
	caesuraFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if caesuraFormCallback.CreationMode || caesuraFormCallback.formGroup.HasSuppressButtonBeenPressed {
		caesuraFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(caesuraFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CaesuraFormCallback(
			nil,
			caesuraFormCallback.probe,
			newFormGroup,
		)
		caesura := new(models.Caesura)
		FillUpForm(caesura, newFormGroup, caesuraFormCallback.probe)
		caesuraFormCallback.probe.formStage.Commit()
	}

	fillUpTree(caesuraFormCallback.probe)
}
func __gong__New__CancelFormCallback(
	cancel *models.Cancel,
	probe *Probe,
	formGroup *table.FormGroup,
) (cancelFormCallback *CancelFormCallback) {
	cancelFormCallback = new(CancelFormCallback)
	cancelFormCallback.probe = probe
	cancelFormCallback.cancel = cancel
	cancelFormCallback.formGroup = formGroup

	cancelFormCallback.CreationMode = (cancel == nil)

	return
}

type CancelFormCallback struct {
	cancel *models.Cancel

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (cancelFormCallback *CancelFormCallback) OnSave() {

	log.Println("CancelFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cancelFormCallback.probe.formStage.Checkout()

	if cancelFormCallback.cancel == nil {
		cancelFormCallback.cancel = new(models.Cancel).Stage(cancelFormCallback.probe.stageOfInterest)
	}
	cancel_ := cancelFormCallback.cancel
	_ = cancel_

	for _, formDiv := range cancelFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cancel_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if cancelFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cancel_.Unstage(cancelFormCallback.probe.stageOfInterest)
	}

	cancelFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Cancel](
		cancelFormCallback.probe,
	)
	cancelFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if cancelFormCallback.CreationMode || cancelFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cancelFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(cancelFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CancelFormCallback(
			nil,
			cancelFormCallback.probe,
			newFormGroup,
		)
		cancel := new(models.Cancel)
		FillUpForm(cancel, newFormGroup, cancelFormCallback.probe)
		cancelFormCallback.probe.formStage.Commit()
	}

	fillUpTree(cancelFormCallback.probe)
}
func __gong__New__ClefFormCallback(
	clef *models.Clef,
	probe *Probe,
	formGroup *table.FormGroup,
) (clefFormCallback *ClefFormCallback) {
	clefFormCallback = new(ClefFormCallback)
	clefFormCallback.probe = probe
	clefFormCallback.clef = clef
	clefFormCallback.formGroup = formGroup

	clefFormCallback.CreationMode = (clef == nil)

	return
}

type ClefFormCallback struct {
	clef *models.Clef

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (clefFormCallback *ClefFormCallback) OnSave() {

	log.Println("ClefFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	clefFormCallback.probe.formStage.Checkout()

	if clefFormCallback.clef == nil {
		clefFormCallback.clef = new(models.Clef).Stage(clefFormCallback.probe.stageOfInterest)
	}
	clef_ := clefFormCallback.clef
	_ = clef_

	for _, formDiv := range clefFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(clef_.Name), formDiv)
		case "Attributes:Clef":
			// we need to retrieve the field owner before the change
			var pastAttributesOwner *models.Attributes
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Clef"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				clefFormCallback.probe.stageOfInterest,
				clefFormCallback.probe.backRepoOfInterest,
				clef_,
				&rf)

			if reverseFieldOwner != nil {
				pastAttributesOwner = reverseFieldOwner.(*models.Attributes)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAttributesOwner != nil {
					idx := slices.Index(pastAttributesOwner.Clef, clef_)
					pastAttributesOwner.Clef = slices.Delete(pastAttributesOwner.Clef, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attributes := range *models.GetGongstructInstancesSet[models.Attributes](clefFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAttributesOwner := _attributes // we have a match
						if pastAttributesOwner != nil {
							if newAttributesOwner != pastAttributesOwner {
								idx := slices.Index(pastAttributesOwner.Clef, clef_)
								pastAttributesOwner.Clef = slices.Delete(pastAttributesOwner.Clef, idx, idx+1)
								newAttributesOwner.Clef = append(newAttributesOwner.Clef, clef_)
							}
						} else {
							newAttributesOwner.Clef = append(newAttributesOwner.Clef, clef_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if clefFormCallback.formGroup.HasSuppressButtonBeenPressed {
		clef_.Unstage(clefFormCallback.probe.stageOfInterest)
	}

	clefFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Clef](
		clefFormCallback.probe,
	)
	clefFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if clefFormCallback.CreationMode || clefFormCallback.formGroup.HasSuppressButtonBeenPressed {
		clefFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(clefFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ClefFormCallback(
			nil,
			clefFormCallback.probe,
			newFormGroup,
		)
		clef := new(models.Clef)
		FillUpForm(clef, newFormGroup, clefFormCallback.probe)
		clefFormCallback.probe.formStage.Commit()
	}

	fillUpTree(clefFormCallback.probe)
}
func __gong__New__CodaFormCallback(
	coda *models.Coda,
	probe *Probe,
	formGroup *table.FormGroup,
) (codaFormCallback *CodaFormCallback) {
	codaFormCallback = new(CodaFormCallback)
	codaFormCallback.probe = probe
	codaFormCallback.coda = coda
	codaFormCallback.formGroup = formGroup

	codaFormCallback.CreationMode = (coda == nil)

	return
}

type CodaFormCallback struct {
	coda *models.Coda

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (codaFormCallback *CodaFormCallback) OnSave() {

	log.Println("CodaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	codaFormCallback.probe.formStage.Checkout()

	if codaFormCallback.coda == nil {
		codaFormCallback.coda = new(models.Coda).Stage(codaFormCallback.probe.stageOfInterest)
	}
	coda_ := codaFormCallback.coda
	_ = coda_

	for _, formDiv := range codaFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(coda_.Name), formDiv)
		case "Direction_type:Coda":
			// we need to retrieve the field owner before the change
			var pastDirection_typeOwner *models.Direction_type
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Coda"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				codaFormCallback.probe.stageOfInterest,
				codaFormCallback.probe.backRepoOfInterest,
				coda_,
				&rf)

			if reverseFieldOwner != nil {
				pastDirection_typeOwner = reverseFieldOwner.(*models.Direction_type)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDirection_typeOwner != nil {
					idx := slices.Index(pastDirection_typeOwner.Coda, coda_)
					pastDirection_typeOwner.Coda = slices.Delete(pastDirection_typeOwner.Coda, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _direction_type := range *models.GetGongstructInstancesSet[models.Direction_type](codaFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _direction_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDirection_typeOwner := _direction_type // we have a match
						if pastDirection_typeOwner != nil {
							if newDirection_typeOwner != pastDirection_typeOwner {
								idx := slices.Index(pastDirection_typeOwner.Coda, coda_)
								pastDirection_typeOwner.Coda = slices.Delete(pastDirection_typeOwner.Coda, idx, idx+1)
								newDirection_typeOwner.Coda = append(newDirection_typeOwner.Coda, coda_)
							}
						} else {
							newDirection_typeOwner.Coda = append(newDirection_typeOwner.Coda, coda_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if codaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		coda_.Unstage(codaFormCallback.probe.stageOfInterest)
	}

	codaFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Coda](
		codaFormCallback.probe,
	)
	codaFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if codaFormCallback.CreationMode || codaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		codaFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(codaFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CodaFormCallback(
			nil,
			codaFormCallback.probe,
			newFormGroup,
		)
		coda := new(models.Coda)
		FillUpForm(coda, newFormGroup, codaFormCallback.probe)
		codaFormCallback.probe.formStage.Commit()
	}

	fillUpTree(codaFormCallback.probe)
}
func __gong__New__CreditFormCallback(
	credit *models.Credit,
	probe *Probe,
	formGroup *table.FormGroup,
) (creditFormCallback *CreditFormCallback) {
	creditFormCallback = new(CreditFormCallback)
	creditFormCallback.probe = probe
	creditFormCallback.credit = credit
	creditFormCallback.formGroup = formGroup

	creditFormCallback.CreationMode = (credit == nil)

	return
}

type CreditFormCallback struct {
	credit *models.Credit

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (creditFormCallback *CreditFormCallback) OnSave() {

	log.Println("CreditFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	creditFormCallback.probe.formStage.Checkout()

	if creditFormCallback.credit == nil {
		creditFormCallback.credit = new(models.Credit).Stage(creditFormCallback.probe.stageOfInterest)
	}
	credit_ := creditFormCallback.credit
	_ = credit_

	for _, formDiv := range creditFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(credit_.Name), formDiv)
		case "Credit_image":
			FormDivSelectFieldToField(&(credit_.Credit_image), creditFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if creditFormCallback.formGroup.HasSuppressButtonBeenPressed {
		credit_.Unstage(creditFormCallback.probe.stageOfInterest)
	}

	creditFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Credit](
		creditFormCallback.probe,
	)
	creditFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if creditFormCallback.CreationMode || creditFormCallback.formGroup.HasSuppressButtonBeenPressed {
		creditFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(creditFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CreditFormCallback(
			nil,
			creditFormCallback.probe,
			newFormGroup,
		)
		credit := new(models.Credit)
		FillUpForm(credit, newFormGroup, creditFormCallback.probe)
		creditFormCallback.probe.formStage.Commit()
	}

	fillUpTree(creditFormCallback.probe)
}
func __gong__New__DashesFormCallback(
	dashes *models.Dashes,
	probe *Probe,
	formGroup *table.FormGroup,
) (dashesFormCallback *DashesFormCallback) {
	dashesFormCallback = new(DashesFormCallback)
	dashesFormCallback.probe = probe
	dashesFormCallback.dashes = dashes
	dashesFormCallback.formGroup = formGroup

	dashesFormCallback.CreationMode = (dashes == nil)

	return
}

type DashesFormCallback struct {
	dashes *models.Dashes

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (dashesFormCallback *DashesFormCallback) OnSave() {

	log.Println("DashesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dashesFormCallback.probe.formStage.Checkout()

	if dashesFormCallback.dashes == nil {
		dashesFormCallback.dashes = new(models.Dashes).Stage(dashesFormCallback.probe.stageOfInterest)
	}
	dashes_ := dashesFormCallback.dashes
	_ = dashes_

	for _, formDiv := range dashesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(dashes_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if dashesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dashes_.Unstage(dashesFormCallback.probe.stageOfInterest)
	}

	dashesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Dashes](
		dashesFormCallback.probe,
	)
	dashesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if dashesFormCallback.CreationMode || dashesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dashesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(dashesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DashesFormCallback(
			nil,
			dashesFormCallback.probe,
			newFormGroup,
		)
		dashes := new(models.Dashes)
		FillUpForm(dashes, newFormGroup, dashesFormCallback.probe)
		dashesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(dashesFormCallback.probe)
}
func __gong__New__DefaultsFormCallback(
	defaults *models.Defaults,
	probe *Probe,
	formGroup *table.FormGroup,
) (defaultsFormCallback *DefaultsFormCallback) {
	defaultsFormCallback = new(DefaultsFormCallback)
	defaultsFormCallback.probe = probe
	defaultsFormCallback.defaults = defaults
	defaultsFormCallback.formGroup = formGroup

	defaultsFormCallback.CreationMode = (defaults == nil)

	return
}

type DefaultsFormCallback struct {
	defaults *models.Defaults

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (defaultsFormCallback *DefaultsFormCallback) OnSave() {

	log.Println("DefaultsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	defaultsFormCallback.probe.formStage.Checkout()

	if defaultsFormCallback.defaults == nil {
		defaultsFormCallback.defaults = new(models.Defaults).Stage(defaultsFormCallback.probe.stageOfInterest)
	}
	defaults_ := defaultsFormCallback.defaults
	_ = defaults_

	for _, formDiv := range defaultsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(defaults_.Name), formDiv)
		case "Scaling":
			FormDivSelectFieldToField(&(defaults_.Scaling), defaultsFormCallback.probe.stageOfInterest, formDiv)
		case "Concert_score":
			FormDivSelectFieldToField(&(defaults_.Concert_score), defaultsFormCallback.probe.stageOfInterest, formDiv)
		case "Appearance":
			FormDivSelectFieldToField(&(defaults_.Appearance), defaultsFormCallback.probe.stageOfInterest, formDiv)
		case "Music_font":
			FormDivSelectFieldToField(&(defaults_.Music_font), defaultsFormCallback.probe.stageOfInterest, formDiv)
		case "Word_font":
			FormDivSelectFieldToField(&(defaults_.Word_font), defaultsFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if defaultsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		defaults_.Unstage(defaultsFormCallback.probe.stageOfInterest)
	}

	defaultsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Defaults](
		defaultsFormCallback.probe,
	)
	defaultsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if defaultsFormCallback.CreationMode || defaultsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		defaultsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(defaultsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DefaultsFormCallback(
			nil,
			defaultsFormCallback.probe,
			newFormGroup,
		)
		defaults := new(models.Defaults)
		FillUpForm(defaults, newFormGroup, defaultsFormCallback.probe)
		defaultsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(defaultsFormCallback.probe)
}
func __gong__New__DegreeFormCallback(
	degree *models.Degree,
	probe *Probe,
	formGroup *table.FormGroup,
) (degreeFormCallback *DegreeFormCallback) {
	degreeFormCallback = new(DegreeFormCallback)
	degreeFormCallback.probe = probe
	degreeFormCallback.degree = degree
	degreeFormCallback.formGroup = formGroup

	degreeFormCallback.CreationMode = (degree == nil)

	return
}

type DegreeFormCallback struct {
	degree *models.Degree

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (degreeFormCallback *DegreeFormCallback) OnSave() {

	log.Println("DegreeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	degreeFormCallback.probe.formStage.Checkout()

	if degreeFormCallback.degree == nil {
		degreeFormCallback.degree = new(models.Degree).Stage(degreeFormCallback.probe.stageOfInterest)
	}
	degree_ := degreeFormCallback.degree
	_ = degree_

	for _, formDiv := range degreeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(degree_.Name), formDiv)
		case "Degree_value":
			FormDivSelectFieldToField(&(degree_.Degree_value), degreeFormCallback.probe.stageOfInterest, formDiv)
		case "Degree_alter":
			FormDivSelectFieldToField(&(degree_.Degree_alter), degreeFormCallback.probe.stageOfInterest, formDiv)
		case "Degree_type":
			FormDivSelectFieldToField(&(degree_.Degree_type), degreeFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if degreeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		degree_.Unstage(degreeFormCallback.probe.stageOfInterest)
	}

	degreeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Degree](
		degreeFormCallback.probe,
	)
	degreeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if degreeFormCallback.CreationMode || degreeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		degreeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(degreeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DegreeFormCallback(
			nil,
			degreeFormCallback.probe,
			newFormGroup,
		)
		degree := new(models.Degree)
		FillUpForm(degree, newFormGroup, degreeFormCallback.probe)
		degreeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(degreeFormCallback.probe)
}
func __gong__New__Degree_alterFormCallback(
	degree_alter *models.Degree_alter,
	probe *Probe,
	formGroup *table.FormGroup,
) (degree_alterFormCallback *Degree_alterFormCallback) {
	degree_alterFormCallback = new(Degree_alterFormCallback)
	degree_alterFormCallback.probe = probe
	degree_alterFormCallback.degree_alter = degree_alter
	degree_alterFormCallback.formGroup = formGroup

	degree_alterFormCallback.CreationMode = (degree_alter == nil)

	return
}

type Degree_alterFormCallback struct {
	degree_alter *models.Degree_alter

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (degree_alterFormCallback *Degree_alterFormCallback) OnSave() {

	log.Println("Degree_alterFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	degree_alterFormCallback.probe.formStage.Checkout()

	if degree_alterFormCallback.degree_alter == nil {
		degree_alterFormCallback.degree_alter = new(models.Degree_alter).Stage(degree_alterFormCallback.probe.stageOfInterest)
	}
	degree_alter_ := degree_alterFormCallback.degree_alter
	_ = degree_alter_

	for _, formDiv := range degree_alterFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(degree_alter_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if degree_alterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		degree_alter_.Unstage(degree_alterFormCallback.probe.stageOfInterest)
	}

	degree_alterFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Degree_alter](
		degree_alterFormCallback.probe,
	)
	degree_alterFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if degree_alterFormCallback.CreationMode || degree_alterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		degree_alterFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(degree_alterFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Degree_alterFormCallback(
			nil,
			degree_alterFormCallback.probe,
			newFormGroup,
		)
		degree_alter := new(models.Degree_alter)
		FillUpForm(degree_alter, newFormGroup, degree_alterFormCallback.probe)
		degree_alterFormCallback.probe.formStage.Commit()
	}

	fillUpTree(degree_alterFormCallback.probe)
}
func __gong__New__Degree_typeFormCallback(
	degree_type *models.Degree_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (degree_typeFormCallback *Degree_typeFormCallback) {
	degree_typeFormCallback = new(Degree_typeFormCallback)
	degree_typeFormCallback.probe = probe
	degree_typeFormCallback.degree_type = degree_type
	degree_typeFormCallback.formGroup = formGroup

	degree_typeFormCallback.CreationMode = (degree_type == nil)

	return
}

type Degree_typeFormCallback struct {
	degree_type *models.Degree_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (degree_typeFormCallback *Degree_typeFormCallback) OnSave() {

	log.Println("Degree_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	degree_typeFormCallback.probe.formStage.Checkout()

	if degree_typeFormCallback.degree_type == nil {
		degree_typeFormCallback.degree_type = new(models.Degree_type).Stage(degree_typeFormCallback.probe.stageOfInterest)
	}
	degree_type_ := degree_typeFormCallback.degree_type
	_ = degree_type_

	for _, formDiv := range degree_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(degree_type_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(degree_type_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if degree_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		degree_type_.Unstage(degree_typeFormCallback.probe.stageOfInterest)
	}

	degree_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Degree_type](
		degree_typeFormCallback.probe,
	)
	degree_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if degree_typeFormCallback.CreationMode || degree_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		degree_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(degree_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Degree_typeFormCallback(
			nil,
			degree_typeFormCallback.probe,
			newFormGroup,
		)
		degree_type := new(models.Degree_type)
		FillUpForm(degree_type, newFormGroup, degree_typeFormCallback.probe)
		degree_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(degree_typeFormCallback.probe)
}
func __gong__New__Degree_valueFormCallback(
	degree_value *models.Degree_value,
	probe *Probe,
	formGroup *table.FormGroup,
) (degree_valueFormCallback *Degree_valueFormCallback) {
	degree_valueFormCallback = new(Degree_valueFormCallback)
	degree_valueFormCallback.probe = probe
	degree_valueFormCallback.degree_value = degree_value
	degree_valueFormCallback.formGroup = formGroup

	degree_valueFormCallback.CreationMode = (degree_value == nil)

	return
}

type Degree_valueFormCallback struct {
	degree_value *models.Degree_value

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (degree_valueFormCallback *Degree_valueFormCallback) OnSave() {

	log.Println("Degree_valueFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	degree_valueFormCallback.probe.formStage.Checkout()

	if degree_valueFormCallback.degree_value == nil {
		degree_valueFormCallback.degree_value = new(models.Degree_value).Stage(degree_valueFormCallback.probe.stageOfInterest)
	}
	degree_value_ := degree_valueFormCallback.degree_value
	_ = degree_value_

	for _, formDiv := range degree_valueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(degree_value_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(degree_value_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if degree_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		degree_value_.Unstage(degree_valueFormCallback.probe.stageOfInterest)
	}

	degree_valueFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Degree_value](
		degree_valueFormCallback.probe,
	)
	degree_valueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if degree_valueFormCallback.CreationMode || degree_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		degree_valueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(degree_valueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Degree_valueFormCallback(
			nil,
			degree_valueFormCallback.probe,
			newFormGroup,
		)
		degree_value := new(models.Degree_value)
		FillUpForm(degree_value, newFormGroup, degree_valueFormCallback.probe)
		degree_valueFormCallback.probe.formStage.Commit()
	}

	fillUpTree(degree_valueFormCallback.probe)
}
func __gong__New__DirectionFormCallback(
	direction *models.Direction,
	probe *Probe,
	formGroup *table.FormGroup,
) (directionFormCallback *DirectionFormCallback) {
	directionFormCallback = new(DirectionFormCallback)
	directionFormCallback.probe = probe
	directionFormCallback.direction = direction
	directionFormCallback.formGroup = formGroup

	directionFormCallback.CreationMode = (direction == nil)

	return
}

type DirectionFormCallback struct {
	direction *models.Direction

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (directionFormCallback *DirectionFormCallback) OnSave() {

	log.Println("DirectionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	directionFormCallback.probe.formStage.Checkout()

	if directionFormCallback.direction == nil {
		directionFormCallback.direction = new(models.Direction).Stage(directionFormCallback.probe.stageOfInterest)
	}
	direction_ := directionFormCallback.direction
	_ = direction_

	for _, formDiv := range directionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(direction_.Name), formDiv)
		case "Offset":
			FormDivSelectFieldToField(&(direction_.Offset), directionFormCallback.probe.stageOfInterest, formDiv)
		case "Sound":
			FormDivSelectFieldToField(&(direction_.Sound), directionFormCallback.probe.stageOfInterest, formDiv)
		case "Listening":
			FormDivSelectFieldToField(&(direction_.Listening), directionFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if directionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		direction_.Unstage(directionFormCallback.probe.stageOfInterest)
	}

	directionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Direction](
		directionFormCallback.probe,
	)
	directionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if directionFormCallback.CreationMode || directionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		directionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(directionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DirectionFormCallback(
			nil,
			directionFormCallback.probe,
			newFormGroup,
		)
		direction := new(models.Direction)
		FillUpForm(direction, newFormGroup, directionFormCallback.probe)
		directionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(directionFormCallback.probe)
}
func __gong__New__Direction_typeFormCallback(
	direction_type *models.Direction_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (direction_typeFormCallback *Direction_typeFormCallback) {
	direction_typeFormCallback = new(Direction_typeFormCallback)
	direction_typeFormCallback.probe = probe
	direction_typeFormCallback.direction_type = direction_type
	direction_typeFormCallback.formGroup = formGroup

	direction_typeFormCallback.CreationMode = (direction_type == nil)

	return
}

type Direction_typeFormCallback struct {
	direction_type *models.Direction_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (direction_typeFormCallback *Direction_typeFormCallback) OnSave() {

	log.Println("Direction_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	direction_typeFormCallback.probe.formStage.Checkout()

	if direction_typeFormCallback.direction_type == nil {
		direction_typeFormCallback.direction_type = new(models.Direction_type).Stage(direction_typeFormCallback.probe.stageOfInterest)
	}
	direction_type_ := direction_typeFormCallback.direction_type
	_ = direction_type_

	for _, formDiv := range direction_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(direction_type_.Name), formDiv)
		case "Wedge":
			FormDivSelectFieldToField(&(direction_type_.Wedge), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Dashes":
			FormDivSelectFieldToField(&(direction_type_.Dashes), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Bracket":
			FormDivSelectFieldToField(&(direction_type_.Bracket), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Pedal":
			FormDivSelectFieldToField(&(direction_type_.Pedal), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Metronome":
			FormDivSelectFieldToField(&(direction_type_.Metronome), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Octave_shift":
			FormDivSelectFieldToField(&(direction_type_.Octave_shift), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Harp_pedals":
			FormDivSelectFieldToField(&(direction_type_.Harp_pedals), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Damp":
			FormDivSelectFieldToField(&(direction_type_.Damp), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Damp_all":
			FormDivSelectFieldToField(&(direction_type_.Damp_all), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Eyeglasses":
			FormDivSelectFieldToField(&(direction_type_.Eyeglasses), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "String_mute":
			FormDivSelectFieldToField(&(direction_type_.String_mute), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Scordatura":
			FormDivSelectFieldToField(&(direction_type_.Scordatura), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Image":
			FormDivSelectFieldToField(&(direction_type_.Image), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Principal_voice":
			FormDivSelectFieldToField(&(direction_type_.Principal_voice), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Accordion_registration":
			FormDivSelectFieldToField(&(direction_type_.Accordion_registration), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Staff_divide":
			FormDivSelectFieldToField(&(direction_type_.Staff_divide), direction_typeFormCallback.probe.stageOfInterest, formDiv)
		case "Direction:Direction_type":
			// we need to retrieve the field owner before the change
			var pastDirectionOwner *models.Direction
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction"
			rf.Fieldname = "Direction_type"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				direction_typeFormCallback.probe.stageOfInterest,
				direction_typeFormCallback.probe.backRepoOfInterest,
				direction_type_,
				&rf)

			if reverseFieldOwner != nil {
				pastDirectionOwner = reverseFieldOwner.(*models.Direction)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDirectionOwner != nil {
					idx := slices.Index(pastDirectionOwner.Direction_type, direction_type_)
					pastDirectionOwner.Direction_type = slices.Delete(pastDirectionOwner.Direction_type, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _direction := range *models.GetGongstructInstancesSet[models.Direction](direction_typeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _direction.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDirectionOwner := _direction // we have a match
						if pastDirectionOwner != nil {
							if newDirectionOwner != pastDirectionOwner {
								idx := slices.Index(pastDirectionOwner.Direction_type, direction_type_)
								pastDirectionOwner.Direction_type = slices.Delete(pastDirectionOwner.Direction_type, idx, idx+1)
								newDirectionOwner.Direction_type = append(newDirectionOwner.Direction_type, direction_type_)
							}
						} else {
							newDirectionOwner.Direction_type = append(newDirectionOwner.Direction_type, direction_type_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if direction_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		direction_type_.Unstage(direction_typeFormCallback.probe.stageOfInterest)
	}

	direction_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Direction_type](
		direction_typeFormCallback.probe,
	)
	direction_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if direction_typeFormCallback.CreationMode || direction_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		direction_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(direction_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Direction_typeFormCallback(
			nil,
			direction_typeFormCallback.probe,
			newFormGroup,
		)
		direction_type := new(models.Direction_type)
		FillUpForm(direction_type, newFormGroup, direction_typeFormCallback.probe)
		direction_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(direction_typeFormCallback.probe)
}
func __gong__New__DistanceFormCallback(
	distance *models.Distance,
	probe *Probe,
	formGroup *table.FormGroup,
) (distanceFormCallback *DistanceFormCallback) {
	distanceFormCallback = new(DistanceFormCallback)
	distanceFormCallback.probe = probe
	distanceFormCallback.distance = distance
	distanceFormCallback.formGroup = formGroup

	distanceFormCallback.CreationMode = (distance == nil)

	return
}

type DistanceFormCallback struct {
	distance *models.Distance

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (distanceFormCallback *DistanceFormCallback) OnSave() {

	log.Println("DistanceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	distanceFormCallback.probe.formStage.Checkout()

	if distanceFormCallback.distance == nil {
		distanceFormCallback.distance = new(models.Distance).Stage(distanceFormCallback.probe.stageOfInterest)
	}
	distance_ := distanceFormCallback.distance
	_ = distance_

	for _, formDiv := range distanceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(distance_.Name), formDiv)
		case "Appearance:Distance":
			// we need to retrieve the field owner before the change
			var pastAppearanceOwner *models.Appearance
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Distance"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				distanceFormCallback.probe.stageOfInterest,
				distanceFormCallback.probe.backRepoOfInterest,
				distance_,
				&rf)

			if reverseFieldOwner != nil {
				pastAppearanceOwner = reverseFieldOwner.(*models.Appearance)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAppearanceOwner != nil {
					idx := slices.Index(pastAppearanceOwner.Distance, distance_)
					pastAppearanceOwner.Distance = slices.Delete(pastAppearanceOwner.Distance, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _appearance := range *models.GetGongstructInstancesSet[models.Appearance](distanceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _appearance.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAppearanceOwner := _appearance // we have a match
						if pastAppearanceOwner != nil {
							if newAppearanceOwner != pastAppearanceOwner {
								idx := slices.Index(pastAppearanceOwner.Distance, distance_)
								pastAppearanceOwner.Distance = slices.Delete(pastAppearanceOwner.Distance, idx, idx+1)
								newAppearanceOwner.Distance = append(newAppearanceOwner.Distance, distance_)
							}
						} else {
							newAppearanceOwner.Distance = append(newAppearanceOwner.Distance, distance_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if distanceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		distance_.Unstage(distanceFormCallback.probe.stageOfInterest)
	}

	distanceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Distance](
		distanceFormCallback.probe,
	)
	distanceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if distanceFormCallback.CreationMode || distanceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		distanceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(distanceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DistanceFormCallback(
			nil,
			distanceFormCallback.probe,
			newFormGroup,
		)
		distance := new(models.Distance)
		FillUpForm(distance, newFormGroup, distanceFormCallback.probe)
		distanceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(distanceFormCallback.probe)
}
func __gong__New__DoubleFormCallback(
	double *models.Double,
	probe *Probe,
	formGroup *table.FormGroup,
) (doubleFormCallback *DoubleFormCallback) {
	doubleFormCallback = new(DoubleFormCallback)
	doubleFormCallback.probe = probe
	doubleFormCallback.double = double
	doubleFormCallback.formGroup = formGroup

	doubleFormCallback.CreationMode = (double == nil)

	return
}

type DoubleFormCallback struct {
	double *models.Double

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (doubleFormCallback *DoubleFormCallback) OnSave() {

	log.Println("DoubleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	doubleFormCallback.probe.formStage.Checkout()

	if doubleFormCallback.double == nil {
		doubleFormCallback.double = new(models.Double).Stage(doubleFormCallback.probe.stageOfInterest)
	}
	double_ := doubleFormCallback.double
	_ = double_

	for _, formDiv := range doubleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(double_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if doubleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		double_.Unstage(doubleFormCallback.probe.stageOfInterest)
	}

	doubleFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Double](
		doubleFormCallback.probe,
	)
	doubleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if doubleFormCallback.CreationMode || doubleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		doubleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(doubleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DoubleFormCallback(
			nil,
			doubleFormCallback.probe,
			newFormGroup,
		)
		double := new(models.Double)
		FillUpForm(double, newFormGroup, doubleFormCallback.probe)
		doubleFormCallback.probe.formStage.Commit()
	}

	fillUpTree(doubleFormCallback.probe)
}
func __gong__New__DynamicsFormCallback(
	dynamics *models.Dynamics,
	probe *Probe,
	formGroup *table.FormGroup,
) (dynamicsFormCallback *DynamicsFormCallback) {
	dynamicsFormCallback = new(DynamicsFormCallback)
	dynamicsFormCallback.probe = probe
	dynamicsFormCallback.dynamics = dynamics
	dynamicsFormCallback.formGroup = formGroup

	dynamicsFormCallback.CreationMode = (dynamics == nil)

	return
}

type DynamicsFormCallback struct {
	dynamics *models.Dynamics

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (dynamicsFormCallback *DynamicsFormCallback) OnSave() {

	log.Println("DynamicsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dynamicsFormCallback.probe.formStage.Checkout()

	if dynamicsFormCallback.dynamics == nil {
		dynamicsFormCallback.dynamics = new(models.Dynamics).Stage(dynamicsFormCallback.probe.stageOfInterest)
	}
	dynamics_ := dynamicsFormCallback.dynamics
	_ = dynamics_

	for _, formDiv := range dynamicsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(dynamics_.Name), formDiv)
		case "P":
			FormDivSelectFieldToField(&(dynamics_.P), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Pp":
			FormDivSelectFieldToField(&(dynamics_.Pp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Ppp":
			FormDivSelectFieldToField(&(dynamics_.Ppp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Pppp":
			FormDivSelectFieldToField(&(dynamics_.Pppp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Ppppp":
			FormDivSelectFieldToField(&(dynamics_.Ppppp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Pppppp":
			FormDivSelectFieldToField(&(dynamics_.Pppppp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "F":
			FormDivSelectFieldToField(&(dynamics_.F), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Ff":
			FormDivSelectFieldToField(&(dynamics_.Ff), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Fff":
			FormDivSelectFieldToField(&(dynamics_.Fff), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Ffff":
			FormDivSelectFieldToField(&(dynamics_.Ffff), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Fffff":
			FormDivSelectFieldToField(&(dynamics_.Fffff), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Ffffff":
			FormDivSelectFieldToField(&(dynamics_.Ffffff), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Mp":
			FormDivSelectFieldToField(&(dynamics_.Mp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Mf":
			FormDivSelectFieldToField(&(dynamics_.Mf), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Sf":
			FormDivSelectFieldToField(&(dynamics_.Sf), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Sfp":
			FormDivSelectFieldToField(&(dynamics_.Sfp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Sfpp":
			FormDivSelectFieldToField(&(dynamics_.Sfpp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Fp":
			FormDivSelectFieldToField(&(dynamics_.Fp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Rf":
			FormDivSelectFieldToField(&(dynamics_.Rf), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Rfz":
			FormDivSelectFieldToField(&(dynamics_.Rfz), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Sfz":
			FormDivSelectFieldToField(&(dynamics_.Sfz), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Sffz":
			FormDivSelectFieldToField(&(dynamics_.Sffz), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Fz":
			FormDivSelectFieldToField(&(dynamics_.Fz), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "N":
			FormDivSelectFieldToField(&(dynamics_.N), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Pf":
			FormDivSelectFieldToField(&(dynamics_.Pf), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Sfzp":
			FormDivSelectFieldToField(&(dynamics_.Sfzp), dynamicsFormCallback.probe.stageOfInterest, formDiv)
		case "Direction_type:Dynamics":
			// we need to retrieve the field owner before the change
			var pastDirection_typeOwner *models.Direction_type
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Dynamics"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				dynamicsFormCallback.probe.stageOfInterest,
				dynamicsFormCallback.probe.backRepoOfInterest,
				dynamics_,
				&rf)

			if reverseFieldOwner != nil {
				pastDirection_typeOwner = reverseFieldOwner.(*models.Direction_type)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDirection_typeOwner != nil {
					idx := slices.Index(pastDirection_typeOwner.Dynamics, dynamics_)
					pastDirection_typeOwner.Dynamics = slices.Delete(pastDirection_typeOwner.Dynamics, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _direction_type := range *models.GetGongstructInstancesSet[models.Direction_type](dynamicsFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _direction_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDirection_typeOwner := _direction_type // we have a match
						if pastDirection_typeOwner != nil {
							if newDirection_typeOwner != pastDirection_typeOwner {
								idx := slices.Index(pastDirection_typeOwner.Dynamics, dynamics_)
								pastDirection_typeOwner.Dynamics = slices.Delete(pastDirection_typeOwner.Dynamics, idx, idx+1)
								newDirection_typeOwner.Dynamics = append(newDirection_typeOwner.Dynamics, dynamics_)
							}
						} else {
							newDirection_typeOwner.Dynamics = append(newDirection_typeOwner.Dynamics, dynamics_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if dynamicsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dynamics_.Unstage(dynamicsFormCallback.probe.stageOfInterest)
	}

	dynamicsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Dynamics](
		dynamicsFormCallback.probe,
	)
	dynamicsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if dynamicsFormCallback.CreationMode || dynamicsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dynamicsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(dynamicsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DynamicsFormCallback(
			nil,
			dynamicsFormCallback.probe,
			newFormGroup,
		)
		dynamics := new(models.Dynamics)
		FillUpForm(dynamics, newFormGroup, dynamicsFormCallback.probe)
		dynamicsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(dynamicsFormCallback.probe)
}
func __gong__New__EffectFormCallback(
	effect *models.Effect,
	probe *Probe,
	formGroup *table.FormGroup,
) (effectFormCallback *EffectFormCallback) {
	effectFormCallback = new(EffectFormCallback)
	effectFormCallback.probe = probe
	effectFormCallback.effect = effect
	effectFormCallback.formGroup = formGroup

	effectFormCallback.CreationMode = (effect == nil)

	return
}

type EffectFormCallback struct {
	effect *models.Effect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (effectFormCallback *EffectFormCallback) OnSave() {

	log.Println("EffectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	effectFormCallback.probe.formStage.Checkout()

	if effectFormCallback.effect == nil {
		effectFormCallback.effect = new(models.Effect).Stage(effectFormCallback.probe.stageOfInterest)
	}
	effect_ := effectFormCallback.effect
	_ = effect_

	for _, formDiv := range effectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(effect_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if effectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		effect_.Unstage(effectFormCallback.probe.stageOfInterest)
	}

	effectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Effect](
		effectFormCallback.probe,
	)
	effectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if effectFormCallback.CreationMode || effectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		effectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(effectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EffectFormCallback(
			nil,
			effectFormCallback.probe,
			newFormGroup,
		)
		effect := new(models.Effect)
		FillUpForm(effect, newFormGroup, effectFormCallback.probe)
		effectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(effectFormCallback.probe)
}
func __gong__New__ElisionFormCallback(
	elision *models.Elision,
	probe *Probe,
	formGroup *table.FormGroup,
) (elisionFormCallback *ElisionFormCallback) {
	elisionFormCallback = new(ElisionFormCallback)
	elisionFormCallback.probe = probe
	elisionFormCallback.elision = elision
	elisionFormCallback.formGroup = formGroup

	elisionFormCallback.CreationMode = (elision == nil)

	return
}

type ElisionFormCallback struct {
	elision *models.Elision

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (elisionFormCallback *ElisionFormCallback) OnSave() {

	log.Println("ElisionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	elisionFormCallback.probe.formStage.Checkout()

	if elisionFormCallback.elision == nil {
		elisionFormCallback.elision = new(models.Elision).Stage(elisionFormCallback.probe.stageOfInterest)
	}
	elision_ := elisionFormCallback.elision
	_ = elision_

	for _, formDiv := range elisionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(elision_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(elision_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if elisionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		elision_.Unstage(elisionFormCallback.probe.stageOfInterest)
	}

	elisionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Elision](
		elisionFormCallback.probe,
	)
	elisionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if elisionFormCallback.CreationMode || elisionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		elisionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(elisionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ElisionFormCallback(
			nil,
			elisionFormCallback.probe,
			newFormGroup,
		)
		elision := new(models.Elision)
		FillUpForm(elision, newFormGroup, elisionFormCallback.probe)
		elisionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(elisionFormCallback.probe)
}
func __gong__New__EmptyFormCallback(
	empty *models.Empty,
	probe *Probe,
	formGroup *table.FormGroup,
) (emptyFormCallback *EmptyFormCallback) {
	emptyFormCallback = new(EmptyFormCallback)
	emptyFormCallback.probe = probe
	emptyFormCallback.empty = empty
	emptyFormCallback.formGroup = formGroup

	emptyFormCallback.CreationMode = (empty == nil)

	return
}

type EmptyFormCallback struct {
	empty *models.Empty

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (emptyFormCallback *EmptyFormCallback) OnSave() {

	log.Println("EmptyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	emptyFormCallback.probe.formStage.Checkout()

	if emptyFormCallback.empty == nil {
		emptyFormCallback.empty = new(models.Empty).Stage(emptyFormCallback.probe.stageOfInterest)
	}
	empty_ := emptyFormCallback.empty
	_ = empty_

	for _, formDiv := range emptyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_.Name), formDiv)
		case "Metronome_note:Metronome_dot":
			// we need to retrieve the field owner before the change
			var pastMetronome_noteOwner *models.Metronome_note
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Metronome_note"
			rf.Fieldname = "Metronome_dot"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				emptyFormCallback.probe.stageOfInterest,
				emptyFormCallback.probe.backRepoOfInterest,
				empty_,
				&rf)

			if reverseFieldOwner != nil {
				pastMetronome_noteOwner = reverseFieldOwner.(*models.Metronome_note)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastMetronome_noteOwner != nil {
					idx := slices.Index(pastMetronome_noteOwner.Metronome_dot, empty_)
					pastMetronome_noteOwner.Metronome_dot = slices.Delete(pastMetronome_noteOwner.Metronome_dot, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _metronome_note := range *models.GetGongstructInstancesSet[models.Metronome_note](emptyFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _metronome_note.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newMetronome_noteOwner := _metronome_note // we have a match
						if pastMetronome_noteOwner != nil {
							if newMetronome_noteOwner != pastMetronome_noteOwner {
								idx := slices.Index(pastMetronome_noteOwner.Metronome_dot, empty_)
								pastMetronome_noteOwner.Metronome_dot = slices.Delete(pastMetronome_noteOwner.Metronome_dot, idx, idx+1)
								newMetronome_noteOwner.Metronome_dot = append(newMetronome_noteOwner.Metronome_dot, empty_)
							}
						} else {
							newMetronome_noteOwner.Metronome_dot = append(newMetronome_noteOwner.Metronome_dot, empty_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if emptyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_.Unstage(emptyFormCallback.probe.stageOfInterest)
	}

	emptyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty](
		emptyFormCallback.probe,
	)
	emptyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if emptyFormCallback.CreationMode || emptyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		emptyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(emptyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EmptyFormCallback(
			nil,
			emptyFormCallback.probe,
			newFormGroup,
		)
		empty := new(models.Empty)
		FillUpForm(empty, newFormGroup, emptyFormCallback.probe)
		emptyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(emptyFormCallback.probe)
}
func __gong__New__Empty_fontFormCallback(
	empty_font *models.Empty_font,
	probe *Probe,
	formGroup *table.FormGroup,
) (empty_fontFormCallback *Empty_fontFormCallback) {
	empty_fontFormCallback = new(Empty_fontFormCallback)
	empty_fontFormCallback.probe = probe
	empty_fontFormCallback.empty_font = empty_font
	empty_fontFormCallback.formGroup = formGroup

	empty_fontFormCallback.CreationMode = (empty_font == nil)

	return
}

type Empty_fontFormCallback struct {
	empty_font *models.Empty_font

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (empty_fontFormCallback *Empty_fontFormCallback) OnSave() {

	log.Println("Empty_fontFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	empty_fontFormCallback.probe.formStage.Checkout()

	if empty_fontFormCallback.empty_font == nil {
		empty_fontFormCallback.empty_font = new(models.Empty_font).Stage(empty_fontFormCallback.probe.stageOfInterest)
	}
	empty_font_ := empty_fontFormCallback.empty_font
	_ = empty_font_

	for _, formDiv := range empty_fontFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_font_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if empty_fontFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_font_.Unstage(empty_fontFormCallback.probe.stageOfInterest)
	}

	empty_fontFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty_font](
		empty_fontFormCallback.probe,
	)
	empty_fontFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if empty_fontFormCallback.CreationMode || empty_fontFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_fontFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(empty_fontFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Empty_fontFormCallback(
			nil,
			empty_fontFormCallback.probe,
			newFormGroup,
		)
		empty_font := new(models.Empty_font)
		FillUpForm(empty_font, newFormGroup, empty_fontFormCallback.probe)
		empty_fontFormCallback.probe.formStage.Commit()
	}

	fillUpTree(empty_fontFormCallback.probe)
}
func __gong__New__Empty_lineFormCallback(
	empty_line *models.Empty_line,
	probe *Probe,
	formGroup *table.FormGroup,
) (empty_lineFormCallback *Empty_lineFormCallback) {
	empty_lineFormCallback = new(Empty_lineFormCallback)
	empty_lineFormCallback.probe = probe
	empty_lineFormCallback.empty_line = empty_line
	empty_lineFormCallback.formGroup = formGroup

	empty_lineFormCallback.CreationMode = (empty_line == nil)

	return
}

type Empty_lineFormCallback struct {
	empty_line *models.Empty_line

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (empty_lineFormCallback *Empty_lineFormCallback) OnSave() {

	log.Println("Empty_lineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	empty_lineFormCallback.probe.formStage.Checkout()

	if empty_lineFormCallback.empty_line == nil {
		empty_lineFormCallback.empty_line = new(models.Empty_line).Stage(empty_lineFormCallback.probe.stageOfInterest)
	}
	empty_line_ := empty_lineFormCallback.empty_line
	_ = empty_line_

	for _, formDiv := range empty_lineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_line_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if empty_lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_line_.Unstage(empty_lineFormCallback.probe.stageOfInterest)
	}

	empty_lineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty_line](
		empty_lineFormCallback.probe,
	)
	empty_lineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if empty_lineFormCallback.CreationMode || empty_lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_lineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(empty_lineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Empty_lineFormCallback(
			nil,
			empty_lineFormCallback.probe,
			newFormGroup,
		)
		empty_line := new(models.Empty_line)
		FillUpForm(empty_line, newFormGroup, empty_lineFormCallback.probe)
		empty_lineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(empty_lineFormCallback.probe)
}
func __gong__New__Empty_placementFormCallback(
	empty_placement *models.Empty_placement,
	probe *Probe,
	formGroup *table.FormGroup,
) (empty_placementFormCallback *Empty_placementFormCallback) {
	empty_placementFormCallback = new(Empty_placementFormCallback)
	empty_placementFormCallback.probe = probe
	empty_placementFormCallback.empty_placement = empty_placement
	empty_placementFormCallback.formGroup = formGroup

	empty_placementFormCallback.CreationMode = (empty_placement == nil)

	return
}

type Empty_placementFormCallback struct {
	empty_placement *models.Empty_placement

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (empty_placementFormCallback *Empty_placementFormCallback) OnSave() {

	log.Println("Empty_placementFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	empty_placementFormCallback.probe.formStage.Checkout()

	if empty_placementFormCallback.empty_placement == nil {
		empty_placementFormCallback.empty_placement = new(models.Empty_placement).Stage(empty_placementFormCallback.probe.stageOfInterest)
	}
	empty_placement_ := empty_placementFormCallback.empty_placement
	_ = empty_placement_

	for _, formDiv := range empty_placementFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_placement_.Name), formDiv)
		case "Note:Dot":
			// we need to retrieve the field owner before the change
			var pastNoteOwner *models.Note
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Dot"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				empty_placementFormCallback.probe.stageOfInterest,
				empty_placementFormCallback.probe.backRepoOfInterest,
				empty_placement_,
				&rf)

			if reverseFieldOwner != nil {
				pastNoteOwner = reverseFieldOwner.(*models.Note)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNoteOwner != nil {
					idx := slices.Index(pastNoteOwner.Dot, empty_placement_)
					pastNoteOwner.Dot = slices.Delete(pastNoteOwner.Dot, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _note := range *models.GetGongstructInstancesSet[models.Note](empty_placementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _note.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNoteOwner := _note // we have a match
						if pastNoteOwner != nil {
							if newNoteOwner != pastNoteOwner {
								idx := slices.Index(pastNoteOwner.Dot, empty_placement_)
								pastNoteOwner.Dot = slices.Delete(pastNoteOwner.Dot, idx, idx+1)
								newNoteOwner.Dot = append(newNoteOwner.Dot, empty_placement_)
							}
						} else {
							newNoteOwner.Dot = append(newNoteOwner.Dot, empty_placement_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if empty_placementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_placement_.Unstage(empty_placementFormCallback.probe.stageOfInterest)
	}

	empty_placementFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty_placement](
		empty_placementFormCallback.probe,
	)
	empty_placementFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if empty_placementFormCallback.CreationMode || empty_placementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_placementFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(empty_placementFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Empty_placementFormCallback(
			nil,
			empty_placementFormCallback.probe,
			newFormGroup,
		)
		empty_placement := new(models.Empty_placement)
		FillUpForm(empty_placement, newFormGroup, empty_placementFormCallback.probe)
		empty_placementFormCallback.probe.formStage.Commit()
	}

	fillUpTree(empty_placementFormCallback.probe)
}
func __gong__New__Empty_placement_smuflFormCallback(
	empty_placement_smufl *models.Empty_placement_smufl,
	probe *Probe,
	formGroup *table.FormGroup,
) (empty_placement_smuflFormCallback *Empty_placement_smuflFormCallback) {
	empty_placement_smuflFormCallback = new(Empty_placement_smuflFormCallback)
	empty_placement_smuflFormCallback.probe = probe
	empty_placement_smuflFormCallback.empty_placement_smufl = empty_placement_smufl
	empty_placement_smuflFormCallback.formGroup = formGroup

	empty_placement_smuflFormCallback.CreationMode = (empty_placement_smufl == nil)

	return
}

type Empty_placement_smuflFormCallback struct {
	empty_placement_smufl *models.Empty_placement_smufl

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (empty_placement_smuflFormCallback *Empty_placement_smuflFormCallback) OnSave() {

	log.Println("Empty_placement_smuflFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	empty_placement_smuflFormCallback.probe.formStage.Checkout()

	if empty_placement_smuflFormCallback.empty_placement_smufl == nil {
		empty_placement_smuflFormCallback.empty_placement_smufl = new(models.Empty_placement_smufl).Stage(empty_placement_smuflFormCallback.probe.stageOfInterest)
	}
	empty_placement_smufl_ := empty_placement_smuflFormCallback.empty_placement_smufl
	_ = empty_placement_smufl_

	for _, formDiv := range empty_placement_smuflFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_placement_smufl_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if empty_placement_smuflFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_placement_smufl_.Unstage(empty_placement_smuflFormCallback.probe.stageOfInterest)
	}

	empty_placement_smuflFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty_placement_smufl](
		empty_placement_smuflFormCallback.probe,
	)
	empty_placement_smuflFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if empty_placement_smuflFormCallback.CreationMode || empty_placement_smuflFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_placement_smuflFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(empty_placement_smuflFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Empty_placement_smuflFormCallback(
			nil,
			empty_placement_smuflFormCallback.probe,
			newFormGroup,
		)
		empty_placement_smufl := new(models.Empty_placement_smufl)
		FillUpForm(empty_placement_smufl, newFormGroup, empty_placement_smuflFormCallback.probe)
		empty_placement_smuflFormCallback.probe.formStage.Commit()
	}

	fillUpTree(empty_placement_smuflFormCallback.probe)
}
func __gong__New__Empty_print_object_style_alignFormCallback(
	empty_print_object_style_align *models.Empty_print_object_style_align,
	probe *Probe,
	formGroup *table.FormGroup,
) (empty_print_object_style_alignFormCallback *Empty_print_object_style_alignFormCallback) {
	empty_print_object_style_alignFormCallback = new(Empty_print_object_style_alignFormCallback)
	empty_print_object_style_alignFormCallback.probe = probe
	empty_print_object_style_alignFormCallback.empty_print_object_style_align = empty_print_object_style_align
	empty_print_object_style_alignFormCallback.formGroup = formGroup

	empty_print_object_style_alignFormCallback.CreationMode = (empty_print_object_style_align == nil)

	return
}

type Empty_print_object_style_alignFormCallback struct {
	empty_print_object_style_align *models.Empty_print_object_style_align

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (empty_print_object_style_alignFormCallback *Empty_print_object_style_alignFormCallback) OnSave() {

	log.Println("Empty_print_object_style_alignFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	empty_print_object_style_alignFormCallback.probe.formStage.Checkout()

	if empty_print_object_style_alignFormCallback.empty_print_object_style_align == nil {
		empty_print_object_style_alignFormCallback.empty_print_object_style_align = new(models.Empty_print_object_style_align).Stage(empty_print_object_style_alignFormCallback.probe.stageOfInterest)
	}
	empty_print_object_style_align_ := empty_print_object_style_alignFormCallback.empty_print_object_style_align
	_ = empty_print_object_style_align_

	for _, formDiv := range empty_print_object_style_alignFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_print_object_style_align_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if empty_print_object_style_alignFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_print_object_style_align_.Unstage(empty_print_object_style_alignFormCallback.probe.stageOfInterest)
	}

	empty_print_object_style_alignFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty_print_object_style_align](
		empty_print_object_style_alignFormCallback.probe,
	)
	empty_print_object_style_alignFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if empty_print_object_style_alignFormCallback.CreationMode || empty_print_object_style_alignFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_print_object_style_alignFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(empty_print_object_style_alignFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Empty_print_object_style_alignFormCallback(
			nil,
			empty_print_object_style_alignFormCallback.probe,
			newFormGroup,
		)
		empty_print_object_style_align := new(models.Empty_print_object_style_align)
		FillUpForm(empty_print_object_style_align, newFormGroup, empty_print_object_style_alignFormCallback.probe)
		empty_print_object_style_alignFormCallback.probe.formStage.Commit()
	}

	fillUpTree(empty_print_object_style_alignFormCallback.probe)
}
func __gong__New__Empty_print_styleFormCallback(
	empty_print_style *models.Empty_print_style,
	probe *Probe,
	formGroup *table.FormGroup,
) (empty_print_styleFormCallback *Empty_print_styleFormCallback) {
	empty_print_styleFormCallback = new(Empty_print_styleFormCallback)
	empty_print_styleFormCallback.probe = probe
	empty_print_styleFormCallback.empty_print_style = empty_print_style
	empty_print_styleFormCallback.formGroup = formGroup

	empty_print_styleFormCallback.CreationMode = (empty_print_style == nil)

	return
}

type Empty_print_styleFormCallback struct {
	empty_print_style *models.Empty_print_style

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (empty_print_styleFormCallback *Empty_print_styleFormCallback) OnSave() {

	log.Println("Empty_print_styleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	empty_print_styleFormCallback.probe.formStage.Checkout()

	if empty_print_styleFormCallback.empty_print_style == nil {
		empty_print_styleFormCallback.empty_print_style = new(models.Empty_print_style).Stage(empty_print_styleFormCallback.probe.stageOfInterest)
	}
	empty_print_style_ := empty_print_styleFormCallback.empty_print_style
	_ = empty_print_style_

	for _, formDiv := range empty_print_styleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_print_style_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if empty_print_styleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_print_style_.Unstage(empty_print_styleFormCallback.probe.stageOfInterest)
	}

	empty_print_styleFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty_print_style](
		empty_print_styleFormCallback.probe,
	)
	empty_print_styleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if empty_print_styleFormCallback.CreationMode || empty_print_styleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_print_styleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(empty_print_styleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Empty_print_styleFormCallback(
			nil,
			empty_print_styleFormCallback.probe,
			newFormGroup,
		)
		empty_print_style := new(models.Empty_print_style)
		FillUpForm(empty_print_style, newFormGroup, empty_print_styleFormCallback.probe)
		empty_print_styleFormCallback.probe.formStage.Commit()
	}

	fillUpTree(empty_print_styleFormCallback.probe)
}
func __gong__New__Empty_print_style_alignFormCallback(
	empty_print_style_align *models.Empty_print_style_align,
	probe *Probe,
	formGroup *table.FormGroup,
) (empty_print_style_alignFormCallback *Empty_print_style_alignFormCallback) {
	empty_print_style_alignFormCallback = new(Empty_print_style_alignFormCallback)
	empty_print_style_alignFormCallback.probe = probe
	empty_print_style_alignFormCallback.empty_print_style_align = empty_print_style_align
	empty_print_style_alignFormCallback.formGroup = formGroup

	empty_print_style_alignFormCallback.CreationMode = (empty_print_style_align == nil)

	return
}

type Empty_print_style_alignFormCallback struct {
	empty_print_style_align *models.Empty_print_style_align

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (empty_print_style_alignFormCallback *Empty_print_style_alignFormCallback) OnSave() {

	log.Println("Empty_print_style_alignFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	empty_print_style_alignFormCallback.probe.formStage.Checkout()

	if empty_print_style_alignFormCallback.empty_print_style_align == nil {
		empty_print_style_alignFormCallback.empty_print_style_align = new(models.Empty_print_style_align).Stage(empty_print_style_alignFormCallback.probe.stageOfInterest)
	}
	empty_print_style_align_ := empty_print_style_alignFormCallback.empty_print_style_align
	_ = empty_print_style_align_

	for _, formDiv := range empty_print_style_alignFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_print_style_align_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if empty_print_style_alignFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_print_style_align_.Unstage(empty_print_style_alignFormCallback.probe.stageOfInterest)
	}

	empty_print_style_alignFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty_print_style_align](
		empty_print_style_alignFormCallback.probe,
	)
	empty_print_style_alignFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if empty_print_style_alignFormCallback.CreationMode || empty_print_style_alignFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_print_style_alignFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(empty_print_style_alignFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Empty_print_style_alignFormCallback(
			nil,
			empty_print_style_alignFormCallback.probe,
			newFormGroup,
		)
		empty_print_style_align := new(models.Empty_print_style_align)
		FillUpForm(empty_print_style_align, newFormGroup, empty_print_style_alignFormCallback.probe)
		empty_print_style_alignFormCallback.probe.formStage.Commit()
	}

	fillUpTree(empty_print_style_alignFormCallback.probe)
}
func __gong__New__Empty_print_style_align_idFormCallback(
	empty_print_style_align_id *models.Empty_print_style_align_id,
	probe *Probe,
	formGroup *table.FormGroup,
) (empty_print_style_align_idFormCallback *Empty_print_style_align_idFormCallback) {
	empty_print_style_align_idFormCallback = new(Empty_print_style_align_idFormCallback)
	empty_print_style_align_idFormCallback.probe = probe
	empty_print_style_align_idFormCallback.empty_print_style_align_id = empty_print_style_align_id
	empty_print_style_align_idFormCallback.formGroup = formGroup

	empty_print_style_align_idFormCallback.CreationMode = (empty_print_style_align_id == nil)

	return
}

type Empty_print_style_align_idFormCallback struct {
	empty_print_style_align_id *models.Empty_print_style_align_id

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (empty_print_style_align_idFormCallback *Empty_print_style_align_idFormCallback) OnSave() {

	log.Println("Empty_print_style_align_idFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	empty_print_style_align_idFormCallback.probe.formStage.Checkout()

	if empty_print_style_align_idFormCallback.empty_print_style_align_id == nil {
		empty_print_style_align_idFormCallback.empty_print_style_align_id = new(models.Empty_print_style_align_id).Stage(empty_print_style_align_idFormCallback.probe.stageOfInterest)
	}
	empty_print_style_align_id_ := empty_print_style_align_idFormCallback.empty_print_style_align_id
	_ = empty_print_style_align_id_

	for _, formDiv := range empty_print_style_align_idFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_print_style_align_id_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if empty_print_style_align_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_print_style_align_id_.Unstage(empty_print_style_align_idFormCallback.probe.stageOfInterest)
	}

	empty_print_style_align_idFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty_print_style_align_id](
		empty_print_style_align_idFormCallback.probe,
	)
	empty_print_style_align_idFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if empty_print_style_align_idFormCallback.CreationMode || empty_print_style_align_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_print_style_align_idFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(empty_print_style_align_idFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Empty_print_style_align_idFormCallback(
			nil,
			empty_print_style_align_idFormCallback.probe,
			newFormGroup,
		)
		empty_print_style_align_id := new(models.Empty_print_style_align_id)
		FillUpForm(empty_print_style_align_id, newFormGroup, empty_print_style_align_idFormCallback.probe)
		empty_print_style_align_idFormCallback.probe.formStage.Commit()
	}

	fillUpTree(empty_print_style_align_idFormCallback.probe)
}
func __gong__New__Empty_trill_soundFormCallback(
	empty_trill_sound *models.Empty_trill_sound,
	probe *Probe,
	formGroup *table.FormGroup,
) (empty_trill_soundFormCallback *Empty_trill_soundFormCallback) {
	empty_trill_soundFormCallback = new(Empty_trill_soundFormCallback)
	empty_trill_soundFormCallback.probe = probe
	empty_trill_soundFormCallback.empty_trill_sound = empty_trill_sound
	empty_trill_soundFormCallback.formGroup = formGroup

	empty_trill_soundFormCallback.CreationMode = (empty_trill_sound == nil)

	return
}

type Empty_trill_soundFormCallback struct {
	empty_trill_sound *models.Empty_trill_sound

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (empty_trill_soundFormCallback *Empty_trill_soundFormCallback) OnSave() {

	log.Println("Empty_trill_soundFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	empty_trill_soundFormCallback.probe.formStage.Checkout()

	if empty_trill_soundFormCallback.empty_trill_sound == nil {
		empty_trill_soundFormCallback.empty_trill_sound = new(models.Empty_trill_sound).Stage(empty_trill_soundFormCallback.probe.stageOfInterest)
	}
	empty_trill_sound_ := empty_trill_soundFormCallback.empty_trill_sound
	_ = empty_trill_sound_

	for _, formDiv := range empty_trill_soundFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(empty_trill_sound_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if empty_trill_soundFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_trill_sound_.Unstage(empty_trill_soundFormCallback.probe.stageOfInterest)
	}

	empty_trill_soundFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Empty_trill_sound](
		empty_trill_soundFormCallback.probe,
	)
	empty_trill_soundFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if empty_trill_soundFormCallback.CreationMode || empty_trill_soundFormCallback.formGroup.HasSuppressButtonBeenPressed {
		empty_trill_soundFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(empty_trill_soundFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Empty_trill_soundFormCallback(
			nil,
			empty_trill_soundFormCallback.probe,
			newFormGroup,
		)
		empty_trill_sound := new(models.Empty_trill_sound)
		FillUpForm(empty_trill_sound, newFormGroup, empty_trill_soundFormCallback.probe)
		empty_trill_soundFormCallback.probe.formStage.Commit()
	}

	fillUpTree(empty_trill_soundFormCallback.probe)
}
func __gong__New__EncodingFormCallback(
	encoding *models.Encoding,
	probe *Probe,
	formGroup *table.FormGroup,
) (encodingFormCallback *EncodingFormCallback) {
	encodingFormCallback = new(EncodingFormCallback)
	encodingFormCallback.probe = probe
	encodingFormCallback.encoding = encoding
	encodingFormCallback.formGroup = formGroup

	encodingFormCallback.CreationMode = (encoding == nil)

	return
}

type EncodingFormCallback struct {
	encoding *models.Encoding

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (encodingFormCallback *EncodingFormCallback) OnSave() {

	log.Println("EncodingFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	encodingFormCallback.probe.formStage.Checkout()

	if encodingFormCallback.encoding == nil {
		encodingFormCallback.encoding = new(models.Encoding).Stage(encodingFormCallback.probe.stageOfInterest)
	}
	encoding_ := encodingFormCallback.encoding
	_ = encoding_

	for _, formDiv := range encodingFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(encoding_.Name), formDiv)
		case "Encoder":
			FormDivSelectFieldToField(&(encoding_.Encoder), encodingFormCallback.probe.stageOfInterest, formDiv)
		case "Software":
			FormDivBasicFieldToField(&(encoding_.Software), formDiv)
		case "Encoding_description":
			FormDivBasicFieldToField(&(encoding_.Encoding_description), formDiv)
		case "Supports":
			FormDivSelectFieldToField(&(encoding_.Supports), encodingFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if encodingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		encoding_.Unstage(encodingFormCallback.probe.stageOfInterest)
	}

	encodingFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Encoding](
		encodingFormCallback.probe,
	)
	encodingFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if encodingFormCallback.CreationMode || encodingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		encodingFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(encodingFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EncodingFormCallback(
			nil,
			encodingFormCallback.probe,
			newFormGroup,
		)
		encoding := new(models.Encoding)
		FillUpForm(encoding, newFormGroup, encodingFormCallback.probe)
		encodingFormCallback.probe.formStage.Commit()
	}

	fillUpTree(encodingFormCallback.probe)
}
func __gong__New__EndingFormCallback(
	ending *models.Ending,
	probe *Probe,
	formGroup *table.FormGroup,
) (endingFormCallback *EndingFormCallback) {
	endingFormCallback = new(EndingFormCallback)
	endingFormCallback.probe = probe
	endingFormCallback.ending = ending
	endingFormCallback.formGroup = formGroup

	endingFormCallback.CreationMode = (ending == nil)

	return
}

type EndingFormCallback struct {
	ending *models.Ending

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (endingFormCallback *EndingFormCallback) OnSave() {

	log.Println("EndingFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	endingFormCallback.probe.formStage.Checkout()

	if endingFormCallback.ending == nil {
		endingFormCallback.ending = new(models.Ending).Stage(endingFormCallback.probe.stageOfInterest)
	}
	ending_ := endingFormCallback.ending
	_ = ending_

	for _, formDiv := range endingFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(ending_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(ending_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if endingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ending_.Unstage(endingFormCallback.probe.stageOfInterest)
	}

	endingFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Ending](
		endingFormCallback.probe,
	)
	endingFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if endingFormCallback.CreationMode || endingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endingFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(endingFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EndingFormCallback(
			nil,
			endingFormCallback.probe,
			newFormGroup,
		)
		ending := new(models.Ending)
		FillUpForm(ending, newFormGroup, endingFormCallback.probe)
		endingFormCallback.probe.formStage.Commit()
	}

	fillUpTree(endingFormCallback.probe)
}
func __gong__New__ExtendFormCallback(
	extend *models.Extend,
	probe *Probe,
	formGroup *table.FormGroup,
) (extendFormCallback *ExtendFormCallback) {
	extendFormCallback = new(ExtendFormCallback)
	extendFormCallback.probe = probe
	extendFormCallback.extend = extend
	extendFormCallback.formGroup = formGroup

	extendFormCallback.CreationMode = (extend == nil)

	return
}

type ExtendFormCallback struct {
	extend *models.Extend

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (extendFormCallback *ExtendFormCallback) OnSave() {

	log.Println("ExtendFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	extendFormCallback.probe.formStage.Checkout()

	if extendFormCallback.extend == nil {
		extendFormCallback.extend = new(models.Extend).Stage(extendFormCallback.probe.stageOfInterest)
	}
	extend_ := extendFormCallback.extend
	_ = extend_

	for _, formDiv := range extendFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(extend_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if extendFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extend_.Unstage(extendFormCallback.probe.stageOfInterest)
	}

	extendFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Extend](
		extendFormCallback.probe,
	)
	extendFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if extendFormCallback.CreationMode || extendFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extendFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(extendFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ExtendFormCallback(
			nil,
			extendFormCallback.probe,
			newFormGroup,
		)
		extend := new(models.Extend)
		FillUpForm(extend, newFormGroup, extendFormCallback.probe)
		extendFormCallback.probe.formStage.Commit()
	}

	fillUpTree(extendFormCallback.probe)
}
func __gong__New__FeatureFormCallback(
	feature *models.Feature,
	probe *Probe,
	formGroup *table.FormGroup,
) (featureFormCallback *FeatureFormCallback) {
	featureFormCallback = new(FeatureFormCallback)
	featureFormCallback.probe = probe
	featureFormCallback.feature = feature
	featureFormCallback.formGroup = formGroup

	featureFormCallback.CreationMode = (feature == nil)

	return
}

type FeatureFormCallback struct {
	feature *models.Feature

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (featureFormCallback *FeatureFormCallback) OnSave() {

	log.Println("FeatureFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	featureFormCallback.probe.formStage.Checkout()

	if featureFormCallback.feature == nil {
		featureFormCallback.feature = new(models.Feature).Stage(featureFormCallback.probe.stageOfInterest)
	}
	feature_ := featureFormCallback.feature
	_ = feature_

	for _, formDiv := range featureFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(feature_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(feature_.Value), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(feature_.Type), formDiv)
		case "Grouping:Feature":
			// we need to retrieve the field owner before the change
			var pastGroupingOwner *models.Grouping
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Grouping"
			rf.Fieldname = "Feature"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				featureFormCallback.probe.stageOfInterest,
				featureFormCallback.probe.backRepoOfInterest,
				feature_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupingOwner = reverseFieldOwner.(*models.Grouping)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupingOwner != nil {
					idx := slices.Index(pastGroupingOwner.Feature, feature_)
					pastGroupingOwner.Feature = slices.Delete(pastGroupingOwner.Feature, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _grouping := range *models.GetGongstructInstancesSet[models.Grouping](featureFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _grouping.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupingOwner := _grouping // we have a match
						if pastGroupingOwner != nil {
							if newGroupingOwner != pastGroupingOwner {
								idx := slices.Index(pastGroupingOwner.Feature, feature_)
								pastGroupingOwner.Feature = slices.Delete(pastGroupingOwner.Feature, idx, idx+1)
								newGroupingOwner.Feature = append(newGroupingOwner.Feature, feature_)
							}
						} else {
							newGroupingOwner.Feature = append(newGroupingOwner.Feature, feature_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if featureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		feature_.Unstage(featureFormCallback.probe.stageOfInterest)
	}

	featureFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Feature](
		featureFormCallback.probe,
	)
	featureFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if featureFormCallback.CreationMode || featureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		featureFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(featureFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FeatureFormCallback(
			nil,
			featureFormCallback.probe,
			newFormGroup,
		)
		feature := new(models.Feature)
		FillUpForm(feature, newFormGroup, featureFormCallback.probe)
		featureFormCallback.probe.formStage.Commit()
	}

	fillUpTree(featureFormCallback.probe)
}
func __gong__New__FermataFormCallback(
	fermata *models.Fermata,
	probe *Probe,
	formGroup *table.FormGroup,
) (fermataFormCallback *FermataFormCallback) {
	fermataFormCallback = new(FermataFormCallback)
	fermataFormCallback.probe = probe
	fermataFormCallback.fermata = fermata
	fermataFormCallback.formGroup = formGroup

	fermataFormCallback.CreationMode = (fermata == nil)

	return
}

type FermataFormCallback struct {
	fermata *models.Fermata

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (fermataFormCallback *FermataFormCallback) OnSave() {

	log.Println("FermataFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fermataFormCallback.probe.formStage.Checkout()

	if fermataFormCallback.fermata == nil {
		fermataFormCallback.fermata = new(models.Fermata).Stage(fermataFormCallback.probe.stageOfInterest)
	}
	fermata_ := fermataFormCallback.fermata
	_ = fermata_

	for _, formDiv := range fermataFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(fermata_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if fermataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fermata_.Unstage(fermataFormCallback.probe.stageOfInterest)
	}

	fermataFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Fermata](
		fermataFormCallback.probe,
	)
	fermataFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fermataFormCallback.CreationMode || fermataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fermataFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(fermataFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FermataFormCallback(
			nil,
			fermataFormCallback.probe,
			newFormGroup,
		)
		fermata := new(models.Fermata)
		FillUpForm(fermata, newFormGroup, fermataFormCallback.probe)
		fermataFormCallback.probe.formStage.Commit()
	}

	fillUpTree(fermataFormCallback.probe)
}
func __gong__New__FigureFormCallback(
	figure *models.Figure,
	probe *Probe,
	formGroup *table.FormGroup,
) (figureFormCallback *FigureFormCallback) {
	figureFormCallback = new(FigureFormCallback)
	figureFormCallback.probe = probe
	figureFormCallback.figure = figure
	figureFormCallback.formGroup = formGroup

	figureFormCallback.CreationMode = (figure == nil)

	return
}

type FigureFormCallback struct {
	figure *models.Figure

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (figureFormCallback *FigureFormCallback) OnSave() {

	log.Println("FigureFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	figureFormCallback.probe.formStage.Checkout()

	if figureFormCallback.figure == nil {
		figureFormCallback.figure = new(models.Figure).Stage(figureFormCallback.probe.stageOfInterest)
	}
	figure_ := figureFormCallback.figure
	_ = figure_

	for _, formDiv := range figureFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(figure_.Name), formDiv)
		case "Extend":
			FormDivSelectFieldToField(&(figure_.Extend), figureFormCallback.probe.stageOfInterest, formDiv)
		case "Figured_bass:Figure":
			// we need to retrieve the field owner before the change
			var pastFigured_bassOwner *models.Figured_bass
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Figured_bass"
			rf.Fieldname = "Figure"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				figureFormCallback.probe.stageOfInterest,
				figureFormCallback.probe.backRepoOfInterest,
				figure_,
				&rf)

			if reverseFieldOwner != nil {
				pastFigured_bassOwner = reverseFieldOwner.(*models.Figured_bass)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFigured_bassOwner != nil {
					idx := slices.Index(pastFigured_bassOwner.Figure, figure_)
					pastFigured_bassOwner.Figure = slices.Delete(pastFigured_bassOwner.Figure, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _figured_bass := range *models.GetGongstructInstancesSet[models.Figured_bass](figureFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _figured_bass.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFigured_bassOwner := _figured_bass // we have a match
						if pastFigured_bassOwner != nil {
							if newFigured_bassOwner != pastFigured_bassOwner {
								idx := slices.Index(pastFigured_bassOwner.Figure, figure_)
								pastFigured_bassOwner.Figure = slices.Delete(pastFigured_bassOwner.Figure, idx, idx+1)
								newFigured_bassOwner.Figure = append(newFigured_bassOwner.Figure, figure_)
							}
						} else {
							newFigured_bassOwner.Figure = append(newFigured_bassOwner.Figure, figure_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if figureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		figure_.Unstage(figureFormCallback.probe.stageOfInterest)
	}

	figureFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Figure](
		figureFormCallback.probe,
	)
	figureFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if figureFormCallback.CreationMode || figureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		figureFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(figureFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FigureFormCallback(
			nil,
			figureFormCallback.probe,
			newFormGroup,
		)
		figure := new(models.Figure)
		FillUpForm(figure, newFormGroup, figureFormCallback.probe)
		figureFormCallback.probe.formStage.Commit()
	}

	fillUpTree(figureFormCallback.probe)
}
func __gong__New__Figured_bassFormCallback(
	figured_bass *models.Figured_bass,
	probe *Probe,
	formGroup *table.FormGroup,
) (figured_bassFormCallback *Figured_bassFormCallback) {
	figured_bassFormCallback = new(Figured_bassFormCallback)
	figured_bassFormCallback.probe = probe
	figured_bassFormCallback.figured_bass = figured_bass
	figured_bassFormCallback.formGroup = formGroup

	figured_bassFormCallback.CreationMode = (figured_bass == nil)

	return
}

type Figured_bassFormCallback struct {
	figured_bass *models.Figured_bass

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (figured_bassFormCallback *Figured_bassFormCallback) OnSave() {

	log.Println("Figured_bassFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	figured_bassFormCallback.probe.formStage.Checkout()

	if figured_bassFormCallback.figured_bass == nil {
		figured_bassFormCallback.figured_bass = new(models.Figured_bass).Stage(figured_bassFormCallback.probe.stageOfInterest)
	}
	figured_bass_ := figured_bassFormCallback.figured_bass
	_ = figured_bass_

	for _, formDiv := range figured_bassFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(figured_bass_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if figured_bassFormCallback.formGroup.HasSuppressButtonBeenPressed {
		figured_bass_.Unstage(figured_bassFormCallback.probe.stageOfInterest)
	}

	figured_bassFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Figured_bass](
		figured_bassFormCallback.probe,
	)
	figured_bassFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if figured_bassFormCallback.CreationMode || figured_bassFormCallback.formGroup.HasSuppressButtonBeenPressed {
		figured_bassFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(figured_bassFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Figured_bassFormCallback(
			nil,
			figured_bassFormCallback.probe,
			newFormGroup,
		)
		figured_bass := new(models.Figured_bass)
		FillUpForm(figured_bass, newFormGroup, figured_bassFormCallback.probe)
		figured_bassFormCallback.probe.formStage.Commit()
	}

	fillUpTree(figured_bassFormCallback.probe)
}
func __gong__New__FingeringFormCallback(
	fingering *models.Fingering,
	probe *Probe,
	formGroup *table.FormGroup,
) (fingeringFormCallback *FingeringFormCallback) {
	fingeringFormCallback = new(FingeringFormCallback)
	fingeringFormCallback.probe = probe
	fingeringFormCallback.fingering = fingering
	fingeringFormCallback.formGroup = formGroup

	fingeringFormCallback.CreationMode = (fingering == nil)

	return
}

type FingeringFormCallback struct {
	fingering *models.Fingering

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (fingeringFormCallback *FingeringFormCallback) OnSave() {

	log.Println("FingeringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fingeringFormCallback.probe.formStage.Checkout()

	if fingeringFormCallback.fingering == nil {
		fingeringFormCallback.fingering = new(models.Fingering).Stage(fingeringFormCallback.probe.stageOfInterest)
	}
	fingering_ := fingeringFormCallback.fingering
	_ = fingering_

	for _, formDiv := range fingeringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(fingering_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(fingering_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if fingeringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fingering_.Unstage(fingeringFormCallback.probe.stageOfInterest)
	}

	fingeringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Fingering](
		fingeringFormCallback.probe,
	)
	fingeringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fingeringFormCallback.CreationMode || fingeringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fingeringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(fingeringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FingeringFormCallback(
			nil,
			fingeringFormCallback.probe,
			newFormGroup,
		)
		fingering := new(models.Fingering)
		FillUpForm(fingering, newFormGroup, fingeringFormCallback.probe)
		fingeringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(fingeringFormCallback.probe)
}
func __gong__New__First_fretFormCallback(
	first_fret *models.First_fret,
	probe *Probe,
	formGroup *table.FormGroup,
) (first_fretFormCallback *First_fretFormCallback) {
	first_fretFormCallback = new(First_fretFormCallback)
	first_fretFormCallback.probe = probe
	first_fretFormCallback.first_fret = first_fret
	first_fretFormCallback.formGroup = formGroup

	first_fretFormCallback.CreationMode = (first_fret == nil)

	return
}

type First_fretFormCallback struct {
	first_fret *models.First_fret

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (first_fretFormCallback *First_fretFormCallback) OnSave() {

	log.Println("First_fretFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	first_fretFormCallback.probe.formStage.Checkout()

	if first_fretFormCallback.first_fret == nil {
		first_fretFormCallback.first_fret = new(models.First_fret).Stage(first_fretFormCallback.probe.stageOfInterest)
	}
	first_fret_ := first_fretFormCallback.first_fret
	_ = first_fret_

	for _, formDiv := range first_fretFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(first_fret_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(first_fret_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if first_fretFormCallback.formGroup.HasSuppressButtonBeenPressed {
		first_fret_.Unstage(first_fretFormCallback.probe.stageOfInterest)
	}

	first_fretFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.First_fret](
		first_fretFormCallback.probe,
	)
	first_fretFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if first_fretFormCallback.CreationMode || first_fretFormCallback.formGroup.HasSuppressButtonBeenPressed {
		first_fretFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(first_fretFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__First_fretFormCallback(
			nil,
			first_fretFormCallback.probe,
			newFormGroup,
		)
		first_fret := new(models.First_fret)
		FillUpForm(first_fret, newFormGroup, first_fretFormCallback.probe)
		first_fretFormCallback.probe.formStage.Commit()
	}

	fillUpTree(first_fretFormCallback.probe)
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
func __gong__New__For_partFormCallback(
	for_part *models.For_part,
	probe *Probe,
	formGroup *table.FormGroup,
) (for_partFormCallback *For_partFormCallback) {
	for_partFormCallback = new(For_partFormCallback)
	for_partFormCallback.probe = probe
	for_partFormCallback.for_part = for_part
	for_partFormCallback.formGroup = formGroup

	for_partFormCallback.CreationMode = (for_part == nil)

	return
}

type For_partFormCallback struct {
	for_part *models.For_part

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (for_partFormCallback *For_partFormCallback) OnSave() {

	log.Println("For_partFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	for_partFormCallback.probe.formStage.Checkout()

	if for_partFormCallback.for_part == nil {
		for_partFormCallback.for_part = new(models.For_part).Stage(for_partFormCallback.probe.stageOfInterest)
	}
	for_part_ := for_partFormCallback.for_part
	_ = for_part_

	for _, formDiv := range for_partFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(for_part_.Name), formDiv)
		case "Part_clef":
			FormDivSelectFieldToField(&(for_part_.Part_clef), for_partFormCallback.probe.stageOfInterest, formDiv)
		case "Part_transpose":
			FormDivSelectFieldToField(&(for_part_.Part_transpose), for_partFormCallback.probe.stageOfInterest, formDiv)
		case "Attributes:For_part":
			// we need to retrieve the field owner before the change
			var pastAttributesOwner *models.Attributes
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "For_part"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				for_partFormCallback.probe.stageOfInterest,
				for_partFormCallback.probe.backRepoOfInterest,
				for_part_,
				&rf)

			if reverseFieldOwner != nil {
				pastAttributesOwner = reverseFieldOwner.(*models.Attributes)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAttributesOwner != nil {
					idx := slices.Index(pastAttributesOwner.For_part, for_part_)
					pastAttributesOwner.For_part = slices.Delete(pastAttributesOwner.For_part, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attributes := range *models.GetGongstructInstancesSet[models.Attributes](for_partFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAttributesOwner := _attributes // we have a match
						if pastAttributesOwner != nil {
							if newAttributesOwner != pastAttributesOwner {
								idx := slices.Index(pastAttributesOwner.For_part, for_part_)
								pastAttributesOwner.For_part = slices.Delete(pastAttributesOwner.For_part, idx, idx+1)
								newAttributesOwner.For_part = append(newAttributesOwner.For_part, for_part_)
							}
						} else {
							newAttributesOwner.For_part = append(newAttributesOwner.For_part, for_part_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if for_partFormCallback.formGroup.HasSuppressButtonBeenPressed {
		for_part_.Unstage(for_partFormCallback.probe.stageOfInterest)
	}

	for_partFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.For_part](
		for_partFormCallback.probe,
	)
	for_partFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if for_partFormCallback.CreationMode || for_partFormCallback.formGroup.HasSuppressButtonBeenPressed {
		for_partFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(for_partFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__For_partFormCallback(
			nil,
			for_partFormCallback.probe,
			newFormGroup,
		)
		for_part := new(models.For_part)
		FillUpForm(for_part, newFormGroup, for_partFormCallback.probe)
		for_partFormCallback.probe.formStage.Commit()
	}

	fillUpTree(for_partFormCallback.probe)
}
func __gong__New__Formatted_symbolFormCallback(
	formatted_symbol *models.Formatted_symbol,
	probe *Probe,
	formGroup *table.FormGroup,
) (formatted_symbolFormCallback *Formatted_symbolFormCallback) {
	formatted_symbolFormCallback = new(Formatted_symbolFormCallback)
	formatted_symbolFormCallback.probe = probe
	formatted_symbolFormCallback.formatted_symbol = formatted_symbol
	formatted_symbolFormCallback.formGroup = formGroup

	formatted_symbolFormCallback.CreationMode = (formatted_symbol == nil)

	return
}

type Formatted_symbolFormCallback struct {
	formatted_symbol *models.Formatted_symbol

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formatted_symbolFormCallback *Formatted_symbolFormCallback) OnSave() {

	log.Println("Formatted_symbolFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formatted_symbolFormCallback.probe.formStage.Checkout()

	if formatted_symbolFormCallback.formatted_symbol == nil {
		formatted_symbolFormCallback.formatted_symbol = new(models.Formatted_symbol).Stage(formatted_symbolFormCallback.probe.stageOfInterest)
	}
	formatted_symbol_ := formatted_symbolFormCallback.formatted_symbol
	_ = formatted_symbol_

	for _, formDiv := range formatted_symbolFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formatted_symbol_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if formatted_symbolFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formatted_symbol_.Unstage(formatted_symbolFormCallback.probe.stageOfInterest)
	}

	formatted_symbolFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Formatted_symbol](
		formatted_symbolFormCallback.probe,
	)
	formatted_symbolFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formatted_symbolFormCallback.CreationMode || formatted_symbolFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formatted_symbolFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(formatted_symbolFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Formatted_symbolFormCallback(
			nil,
			formatted_symbolFormCallback.probe,
			newFormGroup,
		)
		formatted_symbol := new(models.Formatted_symbol)
		FillUpForm(formatted_symbol, newFormGroup, formatted_symbolFormCallback.probe)
		formatted_symbolFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formatted_symbolFormCallback.probe)
}
func __gong__New__Formatted_symbol_idFormCallback(
	formatted_symbol_id *models.Formatted_symbol_id,
	probe *Probe,
	formGroup *table.FormGroup,
) (formatted_symbol_idFormCallback *Formatted_symbol_idFormCallback) {
	formatted_symbol_idFormCallback = new(Formatted_symbol_idFormCallback)
	formatted_symbol_idFormCallback.probe = probe
	formatted_symbol_idFormCallback.formatted_symbol_id = formatted_symbol_id
	formatted_symbol_idFormCallback.formGroup = formGroup

	formatted_symbol_idFormCallback.CreationMode = (formatted_symbol_id == nil)

	return
}

type Formatted_symbol_idFormCallback struct {
	formatted_symbol_id *models.Formatted_symbol_id

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formatted_symbol_idFormCallback *Formatted_symbol_idFormCallback) OnSave() {

	log.Println("Formatted_symbol_idFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formatted_symbol_idFormCallback.probe.formStage.Checkout()

	if formatted_symbol_idFormCallback.formatted_symbol_id == nil {
		formatted_symbol_idFormCallback.formatted_symbol_id = new(models.Formatted_symbol_id).Stage(formatted_symbol_idFormCallback.probe.stageOfInterest)
	}
	formatted_symbol_id_ := formatted_symbol_idFormCallback.formatted_symbol_id
	_ = formatted_symbol_id_

	for _, formDiv := range formatted_symbol_idFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formatted_symbol_id_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if formatted_symbol_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formatted_symbol_id_.Unstage(formatted_symbol_idFormCallback.probe.stageOfInterest)
	}

	formatted_symbol_idFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Formatted_symbol_id](
		formatted_symbol_idFormCallback.probe,
	)
	formatted_symbol_idFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formatted_symbol_idFormCallback.CreationMode || formatted_symbol_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formatted_symbol_idFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(formatted_symbol_idFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Formatted_symbol_idFormCallback(
			nil,
			formatted_symbol_idFormCallback.probe,
			newFormGroup,
		)
		formatted_symbol_id := new(models.Formatted_symbol_id)
		FillUpForm(formatted_symbol_id, newFormGroup, formatted_symbol_idFormCallback.probe)
		formatted_symbol_idFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formatted_symbol_idFormCallback.probe)
}
func __gong__New__ForwardFormCallback(
	forward *models.Forward,
	probe *Probe,
	formGroup *table.FormGroup,
) (forwardFormCallback *ForwardFormCallback) {
	forwardFormCallback = new(ForwardFormCallback)
	forwardFormCallback.probe = probe
	forwardFormCallback.forward = forward
	forwardFormCallback.formGroup = formGroup

	forwardFormCallback.CreationMode = (forward == nil)

	return
}

type ForwardFormCallback struct {
	forward *models.Forward

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (forwardFormCallback *ForwardFormCallback) OnSave() {

	log.Println("ForwardFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	forwardFormCallback.probe.formStage.Checkout()

	if forwardFormCallback.forward == nil {
		forwardFormCallback.forward = new(models.Forward).Stage(forwardFormCallback.probe.stageOfInterest)
	}
	forward_ := forwardFormCallback.forward
	_ = forward_

	for _, formDiv := range forwardFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(forward_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if forwardFormCallback.formGroup.HasSuppressButtonBeenPressed {
		forward_.Unstage(forwardFormCallback.probe.stageOfInterest)
	}

	forwardFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Forward](
		forwardFormCallback.probe,
	)
	forwardFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if forwardFormCallback.CreationMode || forwardFormCallback.formGroup.HasSuppressButtonBeenPressed {
		forwardFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(forwardFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ForwardFormCallback(
			nil,
			forwardFormCallback.probe,
			newFormGroup,
		)
		forward := new(models.Forward)
		FillUpForm(forward, newFormGroup, forwardFormCallback.probe)
		forwardFormCallback.probe.formStage.Commit()
	}

	fillUpTree(forwardFormCallback.probe)
}
func __gong__New__FrameFormCallback(
	frame *models.Frame,
	probe *Probe,
	formGroup *table.FormGroup,
) (frameFormCallback *FrameFormCallback) {
	frameFormCallback = new(FrameFormCallback)
	frameFormCallback.probe = probe
	frameFormCallback.frame = frame
	frameFormCallback.formGroup = formGroup

	frameFormCallback.CreationMode = (frame == nil)

	return
}

type FrameFormCallback struct {
	frame *models.Frame

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (frameFormCallback *FrameFormCallback) OnSave() {

	log.Println("FrameFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	frameFormCallback.probe.formStage.Checkout()

	if frameFormCallback.frame == nil {
		frameFormCallback.frame = new(models.Frame).Stage(frameFormCallback.probe.stageOfInterest)
	}
	frame_ := frameFormCallback.frame
	_ = frame_

	for _, formDiv := range frameFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(frame_.Name), formDiv)
		case "First_fret":
			FormDivSelectFieldToField(&(frame_.First_fret), frameFormCallback.probe.stageOfInterest, formDiv)
		case "Unplayed":
			FormDivBasicFieldToField(&(frame_.Unplayed), formDiv)
		}
	}

	// manage the suppress operation
	if frameFormCallback.formGroup.HasSuppressButtonBeenPressed {
		frame_.Unstage(frameFormCallback.probe.stageOfInterest)
	}

	frameFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Frame](
		frameFormCallback.probe,
	)
	frameFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if frameFormCallback.CreationMode || frameFormCallback.formGroup.HasSuppressButtonBeenPressed {
		frameFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(frameFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FrameFormCallback(
			nil,
			frameFormCallback.probe,
			newFormGroup,
		)
		frame := new(models.Frame)
		FillUpForm(frame, newFormGroup, frameFormCallback.probe)
		frameFormCallback.probe.formStage.Commit()
	}

	fillUpTree(frameFormCallback.probe)
}
func __gong__New__Frame_noteFormCallback(
	frame_note *models.Frame_note,
	probe *Probe,
	formGroup *table.FormGroup,
) (frame_noteFormCallback *Frame_noteFormCallback) {
	frame_noteFormCallback = new(Frame_noteFormCallback)
	frame_noteFormCallback.probe = probe
	frame_noteFormCallback.frame_note = frame_note
	frame_noteFormCallback.formGroup = formGroup

	frame_noteFormCallback.CreationMode = (frame_note == nil)

	return
}

type Frame_noteFormCallback struct {
	frame_note *models.Frame_note

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (frame_noteFormCallback *Frame_noteFormCallback) OnSave() {

	log.Println("Frame_noteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	frame_noteFormCallback.probe.formStage.Checkout()

	if frame_noteFormCallback.frame_note == nil {
		frame_noteFormCallback.frame_note = new(models.Frame_note).Stage(frame_noteFormCallback.probe.stageOfInterest)
	}
	frame_note_ := frame_noteFormCallback.frame_note
	_ = frame_note_

	for _, formDiv := range frame_noteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(frame_note_.Name), formDiv)
		case "Astring":
			FormDivBasicFieldToField(&(frame_note_.Astring), formDiv)
		case "Fret":
			FormDivSelectFieldToField(&(frame_note_.Fret), frame_noteFormCallback.probe.stageOfInterest, formDiv)
		case "Fingering":
			FormDivSelectFieldToField(&(frame_note_.Fingering), frame_noteFormCallback.probe.stageOfInterest, formDiv)
		case "Barre":
			FormDivSelectFieldToField(&(frame_note_.Barre), frame_noteFormCallback.probe.stageOfInterest, formDiv)
		case "Frame:Frame_note":
			// we need to retrieve the field owner before the change
			var pastFrameOwner *models.Frame
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Frame"
			rf.Fieldname = "Frame_note"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				frame_noteFormCallback.probe.stageOfInterest,
				frame_noteFormCallback.probe.backRepoOfInterest,
				frame_note_,
				&rf)

			if reverseFieldOwner != nil {
				pastFrameOwner = reverseFieldOwner.(*models.Frame)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFrameOwner != nil {
					idx := slices.Index(pastFrameOwner.Frame_note, frame_note_)
					pastFrameOwner.Frame_note = slices.Delete(pastFrameOwner.Frame_note, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _frame := range *models.GetGongstructInstancesSet[models.Frame](frame_noteFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _frame.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFrameOwner := _frame // we have a match
						if pastFrameOwner != nil {
							if newFrameOwner != pastFrameOwner {
								idx := slices.Index(pastFrameOwner.Frame_note, frame_note_)
								pastFrameOwner.Frame_note = slices.Delete(pastFrameOwner.Frame_note, idx, idx+1)
								newFrameOwner.Frame_note = append(newFrameOwner.Frame_note, frame_note_)
							}
						} else {
							newFrameOwner.Frame_note = append(newFrameOwner.Frame_note, frame_note_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if frame_noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		frame_note_.Unstage(frame_noteFormCallback.probe.stageOfInterest)
	}

	frame_noteFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Frame_note](
		frame_noteFormCallback.probe,
	)
	frame_noteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if frame_noteFormCallback.CreationMode || frame_noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		frame_noteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(frame_noteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Frame_noteFormCallback(
			nil,
			frame_noteFormCallback.probe,
			newFormGroup,
		)
		frame_note := new(models.Frame_note)
		FillUpForm(frame_note, newFormGroup, frame_noteFormCallback.probe)
		frame_noteFormCallback.probe.formStage.Commit()
	}

	fillUpTree(frame_noteFormCallback.probe)
}
func __gong__New__FretFormCallback(
	fret *models.Fret,
	probe *Probe,
	formGroup *table.FormGroup,
) (fretFormCallback *FretFormCallback) {
	fretFormCallback = new(FretFormCallback)
	fretFormCallback.probe = probe
	fretFormCallback.fret = fret
	fretFormCallback.formGroup = formGroup

	fretFormCallback.CreationMode = (fret == nil)

	return
}

type FretFormCallback struct {
	fret *models.Fret

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (fretFormCallback *FretFormCallback) OnSave() {

	log.Println("FretFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fretFormCallback.probe.formStage.Checkout()

	if fretFormCallback.fret == nil {
		fretFormCallback.fret = new(models.Fret).Stage(fretFormCallback.probe.stageOfInterest)
	}
	fret_ := fretFormCallback.fret
	_ = fret_

	for _, formDiv := range fretFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(fret_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if fretFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fret_.Unstage(fretFormCallback.probe.stageOfInterest)
	}

	fretFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Fret](
		fretFormCallback.probe,
	)
	fretFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fretFormCallback.CreationMode || fretFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fretFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(fretFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FretFormCallback(
			nil,
			fretFormCallback.probe,
			newFormGroup,
		)
		fret := new(models.Fret)
		FillUpForm(fret, newFormGroup, fretFormCallback.probe)
		fretFormCallback.probe.formStage.Commit()
	}

	fillUpTree(fretFormCallback.probe)
}
func __gong__New__GlassFormCallback(
	glass *models.Glass,
	probe *Probe,
	formGroup *table.FormGroup,
) (glassFormCallback *GlassFormCallback) {
	glassFormCallback = new(GlassFormCallback)
	glassFormCallback.probe = probe
	glassFormCallback.glass = glass
	glassFormCallback.formGroup = formGroup

	glassFormCallback.CreationMode = (glass == nil)

	return
}

type GlassFormCallback struct {
	glass *models.Glass

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (glassFormCallback *GlassFormCallback) OnSave() {

	log.Println("GlassFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	glassFormCallback.probe.formStage.Checkout()

	if glassFormCallback.glass == nil {
		glassFormCallback.glass = new(models.Glass).Stage(glassFormCallback.probe.stageOfInterest)
	}
	glass_ := glassFormCallback.glass
	_ = glass_

	for _, formDiv := range glassFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(glass_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if glassFormCallback.formGroup.HasSuppressButtonBeenPressed {
		glass_.Unstage(glassFormCallback.probe.stageOfInterest)
	}

	glassFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Glass](
		glassFormCallback.probe,
	)
	glassFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if glassFormCallback.CreationMode || glassFormCallback.formGroup.HasSuppressButtonBeenPressed {
		glassFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(glassFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GlassFormCallback(
			nil,
			glassFormCallback.probe,
			newFormGroup,
		)
		glass := new(models.Glass)
		FillUpForm(glass, newFormGroup, glassFormCallback.probe)
		glassFormCallback.probe.formStage.Commit()
	}

	fillUpTree(glassFormCallback.probe)
}
func __gong__New__GlissandoFormCallback(
	glissando *models.Glissando,
	probe *Probe,
	formGroup *table.FormGroup,
) (glissandoFormCallback *GlissandoFormCallback) {
	glissandoFormCallback = new(GlissandoFormCallback)
	glissandoFormCallback.probe = probe
	glissandoFormCallback.glissando = glissando
	glissandoFormCallback.formGroup = formGroup

	glissandoFormCallback.CreationMode = (glissando == nil)

	return
}

type GlissandoFormCallback struct {
	glissando *models.Glissando

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (glissandoFormCallback *GlissandoFormCallback) OnSave() {

	log.Println("GlissandoFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	glissandoFormCallback.probe.formStage.Checkout()

	if glissandoFormCallback.glissando == nil {
		glissandoFormCallback.glissando = new(models.Glissando).Stage(glissandoFormCallback.probe.stageOfInterest)
	}
	glissando_ := glissandoFormCallback.glissando
	_ = glissando_

	for _, formDiv := range glissandoFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(glissando_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(glissando_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if glissandoFormCallback.formGroup.HasSuppressButtonBeenPressed {
		glissando_.Unstage(glissandoFormCallback.probe.stageOfInterest)
	}

	glissandoFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Glissando](
		glissandoFormCallback.probe,
	)
	glissandoFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if glissandoFormCallback.CreationMode || glissandoFormCallback.formGroup.HasSuppressButtonBeenPressed {
		glissandoFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(glissandoFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GlissandoFormCallback(
			nil,
			glissandoFormCallback.probe,
			newFormGroup,
		)
		glissando := new(models.Glissando)
		FillUpForm(glissando, newFormGroup, glissandoFormCallback.probe)
		glissandoFormCallback.probe.formStage.Commit()
	}

	fillUpTree(glissandoFormCallback.probe)
}
func __gong__New__GlyphFormCallback(
	glyph *models.Glyph,
	probe *Probe,
	formGroup *table.FormGroup,
) (glyphFormCallback *GlyphFormCallback) {
	glyphFormCallback = new(GlyphFormCallback)
	glyphFormCallback.probe = probe
	glyphFormCallback.glyph = glyph
	glyphFormCallback.formGroup = formGroup

	glyphFormCallback.CreationMode = (glyph == nil)

	return
}

type GlyphFormCallback struct {
	glyph *models.Glyph

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (glyphFormCallback *GlyphFormCallback) OnSave() {

	log.Println("GlyphFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	glyphFormCallback.probe.formStage.Checkout()

	if glyphFormCallback.glyph == nil {
		glyphFormCallback.glyph = new(models.Glyph).Stage(glyphFormCallback.probe.stageOfInterest)
	}
	glyph_ := glyphFormCallback.glyph
	_ = glyph_

	for _, formDiv := range glyphFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(glyph_.Name), formDiv)
		case "Appearance:Glyph":
			// we need to retrieve the field owner before the change
			var pastAppearanceOwner *models.Appearance
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Glyph"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				glyphFormCallback.probe.stageOfInterest,
				glyphFormCallback.probe.backRepoOfInterest,
				glyph_,
				&rf)

			if reverseFieldOwner != nil {
				pastAppearanceOwner = reverseFieldOwner.(*models.Appearance)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAppearanceOwner != nil {
					idx := slices.Index(pastAppearanceOwner.Glyph, glyph_)
					pastAppearanceOwner.Glyph = slices.Delete(pastAppearanceOwner.Glyph, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _appearance := range *models.GetGongstructInstancesSet[models.Appearance](glyphFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _appearance.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAppearanceOwner := _appearance // we have a match
						if pastAppearanceOwner != nil {
							if newAppearanceOwner != pastAppearanceOwner {
								idx := slices.Index(pastAppearanceOwner.Glyph, glyph_)
								pastAppearanceOwner.Glyph = slices.Delete(pastAppearanceOwner.Glyph, idx, idx+1)
								newAppearanceOwner.Glyph = append(newAppearanceOwner.Glyph, glyph_)
							}
						} else {
							newAppearanceOwner.Glyph = append(newAppearanceOwner.Glyph, glyph_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if glyphFormCallback.formGroup.HasSuppressButtonBeenPressed {
		glyph_.Unstage(glyphFormCallback.probe.stageOfInterest)
	}

	glyphFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Glyph](
		glyphFormCallback.probe,
	)
	glyphFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if glyphFormCallback.CreationMode || glyphFormCallback.formGroup.HasSuppressButtonBeenPressed {
		glyphFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(glyphFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GlyphFormCallback(
			nil,
			glyphFormCallback.probe,
			newFormGroup,
		)
		glyph := new(models.Glyph)
		FillUpForm(glyph, newFormGroup, glyphFormCallback.probe)
		glyphFormCallback.probe.formStage.Commit()
	}

	fillUpTree(glyphFormCallback.probe)
}
func __gong__New__GraceFormCallback(
	grace *models.Grace,
	probe *Probe,
	formGroup *table.FormGroup,
) (graceFormCallback *GraceFormCallback) {
	graceFormCallback = new(GraceFormCallback)
	graceFormCallback.probe = probe
	graceFormCallback.grace = grace
	graceFormCallback.formGroup = formGroup

	graceFormCallback.CreationMode = (grace == nil)

	return
}

type GraceFormCallback struct {
	grace *models.Grace

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (graceFormCallback *GraceFormCallback) OnSave() {

	log.Println("GraceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	graceFormCallback.probe.formStage.Checkout()

	if graceFormCallback.grace == nil {
		graceFormCallback.grace = new(models.Grace).Stage(graceFormCallback.probe.stageOfInterest)
	}
	grace_ := graceFormCallback.grace
	_ = grace_

	for _, formDiv := range graceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(grace_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if graceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		grace_.Unstage(graceFormCallback.probe.stageOfInterest)
	}

	graceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Grace](
		graceFormCallback.probe,
	)
	graceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if graceFormCallback.CreationMode || graceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		graceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(graceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GraceFormCallback(
			nil,
			graceFormCallback.probe,
			newFormGroup,
		)
		grace := new(models.Grace)
		FillUpForm(grace, newFormGroup, graceFormCallback.probe)
		graceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(graceFormCallback.probe)
}
func __gong__New__Group_barlineFormCallback(
	group_barline *models.Group_barline,
	probe *Probe,
	formGroup *table.FormGroup,
) (group_barlineFormCallback *Group_barlineFormCallback) {
	group_barlineFormCallback = new(Group_barlineFormCallback)
	group_barlineFormCallback.probe = probe
	group_barlineFormCallback.group_barline = group_barline
	group_barlineFormCallback.formGroup = formGroup

	group_barlineFormCallback.CreationMode = (group_barline == nil)

	return
}

type Group_barlineFormCallback struct {
	group_barline *models.Group_barline

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (group_barlineFormCallback *Group_barlineFormCallback) OnSave() {

	log.Println("Group_barlineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	group_barlineFormCallback.probe.formStage.Checkout()

	if group_barlineFormCallback.group_barline == nil {
		group_barlineFormCallback.group_barline = new(models.Group_barline).Stage(group_barlineFormCallback.probe.stageOfInterest)
	}
	group_barline_ := group_barlineFormCallback.group_barline
	_ = group_barline_

	for _, formDiv := range group_barlineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(group_barline_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if group_barlineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_barline_.Unstage(group_barlineFormCallback.probe.stageOfInterest)
	}

	group_barlineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Group_barline](
		group_barlineFormCallback.probe,
	)
	group_barlineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if group_barlineFormCallback.CreationMode || group_barlineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_barlineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(group_barlineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Group_barlineFormCallback(
			nil,
			group_barlineFormCallback.probe,
			newFormGroup,
		)
		group_barline := new(models.Group_barline)
		FillUpForm(group_barline, newFormGroup, group_barlineFormCallback.probe)
		group_barlineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(group_barlineFormCallback.probe)
}
func __gong__New__Group_symbolFormCallback(
	group_symbol *models.Group_symbol,
	probe *Probe,
	formGroup *table.FormGroup,
) (group_symbolFormCallback *Group_symbolFormCallback) {
	group_symbolFormCallback = new(Group_symbolFormCallback)
	group_symbolFormCallback.probe = probe
	group_symbolFormCallback.group_symbol = group_symbol
	group_symbolFormCallback.formGroup = formGroup

	group_symbolFormCallback.CreationMode = (group_symbol == nil)

	return
}

type Group_symbolFormCallback struct {
	group_symbol *models.Group_symbol

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (group_symbolFormCallback *Group_symbolFormCallback) OnSave() {

	log.Println("Group_symbolFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	group_symbolFormCallback.probe.formStage.Checkout()

	if group_symbolFormCallback.group_symbol == nil {
		group_symbolFormCallback.group_symbol = new(models.Group_symbol).Stage(group_symbolFormCallback.probe.stageOfInterest)
	}
	group_symbol_ := group_symbolFormCallback.group_symbol
	_ = group_symbol_

	for _, formDiv := range group_symbolFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(group_symbol_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if group_symbolFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_symbol_.Unstage(group_symbolFormCallback.probe.stageOfInterest)
	}

	group_symbolFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Group_symbol](
		group_symbolFormCallback.probe,
	)
	group_symbolFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if group_symbolFormCallback.CreationMode || group_symbolFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_symbolFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(group_symbolFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Group_symbolFormCallback(
			nil,
			group_symbolFormCallback.probe,
			newFormGroup,
		)
		group_symbol := new(models.Group_symbol)
		FillUpForm(group_symbol, newFormGroup, group_symbolFormCallback.probe)
		group_symbolFormCallback.probe.formStage.Commit()
	}

	fillUpTree(group_symbolFormCallback.probe)
}
func __gong__New__GroupingFormCallback(
	grouping *models.Grouping,
	probe *Probe,
	formGroup *table.FormGroup,
) (groupingFormCallback *GroupingFormCallback) {
	groupingFormCallback = new(GroupingFormCallback)
	groupingFormCallback.probe = probe
	groupingFormCallback.grouping = grouping
	groupingFormCallback.formGroup = formGroup

	groupingFormCallback.CreationMode = (grouping == nil)

	return
}

type GroupingFormCallback struct {
	grouping *models.Grouping

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (groupingFormCallback *GroupingFormCallback) OnSave() {

	log.Println("GroupingFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	groupingFormCallback.probe.formStage.Checkout()

	if groupingFormCallback.grouping == nil {
		groupingFormCallback.grouping = new(models.Grouping).Stage(groupingFormCallback.probe.stageOfInterest)
	}
	grouping_ := groupingFormCallback.grouping
	_ = grouping_

	for _, formDiv := range groupingFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(grouping_.Name), formDiv)
		case "Number":
			FormDivBasicFieldToField(&(grouping_.Number), formDiv)
		case "Member_of":
			FormDivBasicFieldToField(&(grouping_.Member_of), formDiv)
		}
	}

	// manage the suppress operation
	if groupingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		grouping_.Unstage(groupingFormCallback.probe.stageOfInterest)
	}

	groupingFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Grouping](
		groupingFormCallback.probe,
	)
	groupingFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if groupingFormCallback.CreationMode || groupingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupingFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(groupingFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GroupingFormCallback(
			nil,
			groupingFormCallback.probe,
			newFormGroup,
		)
		grouping := new(models.Grouping)
		FillUpForm(grouping, newFormGroup, groupingFormCallback.probe)
		groupingFormCallback.probe.formStage.Commit()
	}

	fillUpTree(groupingFormCallback.probe)
}
func __gong__New__Hammer_on_pull_offFormCallback(
	hammer_on_pull_off *models.Hammer_on_pull_off,
	probe *Probe,
	formGroup *table.FormGroup,
) (hammer_on_pull_offFormCallback *Hammer_on_pull_offFormCallback) {
	hammer_on_pull_offFormCallback = new(Hammer_on_pull_offFormCallback)
	hammer_on_pull_offFormCallback.probe = probe
	hammer_on_pull_offFormCallback.hammer_on_pull_off = hammer_on_pull_off
	hammer_on_pull_offFormCallback.formGroup = formGroup

	hammer_on_pull_offFormCallback.CreationMode = (hammer_on_pull_off == nil)

	return
}

type Hammer_on_pull_offFormCallback struct {
	hammer_on_pull_off *models.Hammer_on_pull_off

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (hammer_on_pull_offFormCallback *Hammer_on_pull_offFormCallback) OnSave() {

	log.Println("Hammer_on_pull_offFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	hammer_on_pull_offFormCallback.probe.formStage.Checkout()

	if hammer_on_pull_offFormCallback.hammer_on_pull_off == nil {
		hammer_on_pull_offFormCallback.hammer_on_pull_off = new(models.Hammer_on_pull_off).Stage(hammer_on_pull_offFormCallback.probe.stageOfInterest)
	}
	hammer_on_pull_off_ := hammer_on_pull_offFormCallback.hammer_on_pull_off
	_ = hammer_on_pull_off_

	for _, formDiv := range hammer_on_pull_offFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(hammer_on_pull_off_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(hammer_on_pull_off_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if hammer_on_pull_offFormCallback.formGroup.HasSuppressButtonBeenPressed {
		hammer_on_pull_off_.Unstage(hammer_on_pull_offFormCallback.probe.stageOfInterest)
	}

	hammer_on_pull_offFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Hammer_on_pull_off](
		hammer_on_pull_offFormCallback.probe,
	)
	hammer_on_pull_offFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if hammer_on_pull_offFormCallback.CreationMode || hammer_on_pull_offFormCallback.formGroup.HasSuppressButtonBeenPressed {
		hammer_on_pull_offFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(hammer_on_pull_offFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Hammer_on_pull_offFormCallback(
			nil,
			hammer_on_pull_offFormCallback.probe,
			newFormGroup,
		)
		hammer_on_pull_off := new(models.Hammer_on_pull_off)
		FillUpForm(hammer_on_pull_off, newFormGroup, hammer_on_pull_offFormCallback.probe)
		hammer_on_pull_offFormCallback.probe.formStage.Commit()
	}

	fillUpTree(hammer_on_pull_offFormCallback.probe)
}
func __gong__New__HandbellFormCallback(
	handbell *models.Handbell,
	probe *Probe,
	formGroup *table.FormGroup,
) (handbellFormCallback *HandbellFormCallback) {
	handbellFormCallback = new(HandbellFormCallback)
	handbellFormCallback.probe = probe
	handbellFormCallback.handbell = handbell
	handbellFormCallback.formGroup = formGroup

	handbellFormCallback.CreationMode = (handbell == nil)

	return
}

type HandbellFormCallback struct {
	handbell *models.Handbell

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (handbellFormCallback *HandbellFormCallback) OnSave() {

	log.Println("HandbellFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	handbellFormCallback.probe.formStage.Checkout()

	if handbellFormCallback.handbell == nil {
		handbellFormCallback.handbell = new(models.Handbell).Stage(handbellFormCallback.probe.stageOfInterest)
	}
	handbell_ := handbellFormCallback.handbell
	_ = handbell_

	for _, formDiv := range handbellFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(handbell_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if handbellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		handbell_.Unstage(handbellFormCallback.probe.stageOfInterest)
	}

	handbellFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Handbell](
		handbellFormCallback.probe,
	)
	handbellFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if handbellFormCallback.CreationMode || handbellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		handbellFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(handbellFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__HandbellFormCallback(
			nil,
			handbellFormCallback.probe,
			newFormGroup,
		)
		handbell := new(models.Handbell)
		FillUpForm(handbell, newFormGroup, handbellFormCallback.probe)
		handbellFormCallback.probe.formStage.Commit()
	}

	fillUpTree(handbellFormCallback.probe)
}
func __gong__New__Harmon_closedFormCallback(
	harmon_closed *models.Harmon_closed,
	probe *Probe,
	formGroup *table.FormGroup,
) (harmon_closedFormCallback *Harmon_closedFormCallback) {
	harmon_closedFormCallback = new(Harmon_closedFormCallback)
	harmon_closedFormCallback.probe = probe
	harmon_closedFormCallback.harmon_closed = harmon_closed
	harmon_closedFormCallback.formGroup = formGroup

	harmon_closedFormCallback.CreationMode = (harmon_closed == nil)

	return
}

type Harmon_closedFormCallback struct {
	harmon_closed *models.Harmon_closed

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (harmon_closedFormCallback *Harmon_closedFormCallback) OnSave() {

	log.Println("Harmon_closedFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	harmon_closedFormCallback.probe.formStage.Checkout()

	if harmon_closedFormCallback.harmon_closed == nil {
		harmon_closedFormCallback.harmon_closed = new(models.Harmon_closed).Stage(harmon_closedFormCallback.probe.stageOfInterest)
	}
	harmon_closed_ := harmon_closedFormCallback.harmon_closed
	_ = harmon_closed_

	for _, formDiv := range harmon_closedFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(harmon_closed_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if harmon_closedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmon_closed_.Unstage(harmon_closedFormCallback.probe.stageOfInterest)
	}

	harmon_closedFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Harmon_closed](
		harmon_closedFormCallback.probe,
	)
	harmon_closedFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if harmon_closedFormCallback.CreationMode || harmon_closedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmon_closedFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(harmon_closedFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Harmon_closedFormCallback(
			nil,
			harmon_closedFormCallback.probe,
			newFormGroup,
		)
		harmon_closed := new(models.Harmon_closed)
		FillUpForm(harmon_closed, newFormGroup, harmon_closedFormCallback.probe)
		harmon_closedFormCallback.probe.formStage.Commit()
	}

	fillUpTree(harmon_closedFormCallback.probe)
}
func __gong__New__Harmon_muteFormCallback(
	harmon_mute *models.Harmon_mute,
	probe *Probe,
	formGroup *table.FormGroup,
) (harmon_muteFormCallback *Harmon_muteFormCallback) {
	harmon_muteFormCallback = new(Harmon_muteFormCallback)
	harmon_muteFormCallback.probe = probe
	harmon_muteFormCallback.harmon_mute = harmon_mute
	harmon_muteFormCallback.formGroup = formGroup

	harmon_muteFormCallback.CreationMode = (harmon_mute == nil)

	return
}

type Harmon_muteFormCallback struct {
	harmon_mute *models.Harmon_mute

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (harmon_muteFormCallback *Harmon_muteFormCallback) OnSave() {

	log.Println("Harmon_muteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	harmon_muteFormCallback.probe.formStage.Checkout()

	if harmon_muteFormCallback.harmon_mute == nil {
		harmon_muteFormCallback.harmon_mute = new(models.Harmon_mute).Stage(harmon_muteFormCallback.probe.stageOfInterest)
	}
	harmon_mute_ := harmon_muteFormCallback.harmon_mute
	_ = harmon_mute_

	for _, formDiv := range harmon_muteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(harmon_mute_.Name), formDiv)
		case "Harmon_closed":
			FormDivSelectFieldToField(&(harmon_mute_.Harmon_closed), harmon_muteFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if harmon_muteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmon_mute_.Unstage(harmon_muteFormCallback.probe.stageOfInterest)
	}

	harmon_muteFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Harmon_mute](
		harmon_muteFormCallback.probe,
	)
	harmon_muteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if harmon_muteFormCallback.CreationMode || harmon_muteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmon_muteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(harmon_muteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Harmon_muteFormCallback(
			nil,
			harmon_muteFormCallback.probe,
			newFormGroup,
		)
		harmon_mute := new(models.Harmon_mute)
		FillUpForm(harmon_mute, newFormGroup, harmon_muteFormCallback.probe)
		harmon_muteFormCallback.probe.formStage.Commit()
	}

	fillUpTree(harmon_muteFormCallback.probe)
}
func __gong__New__HarmonicFormCallback(
	harmonic *models.Harmonic,
	probe *Probe,
	formGroup *table.FormGroup,
) (harmonicFormCallback *HarmonicFormCallback) {
	harmonicFormCallback = new(HarmonicFormCallback)
	harmonicFormCallback.probe = probe
	harmonicFormCallback.harmonic = harmonic
	harmonicFormCallback.formGroup = formGroup

	harmonicFormCallback.CreationMode = (harmonic == nil)

	return
}

type HarmonicFormCallback struct {
	harmonic *models.Harmonic

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (harmonicFormCallback *HarmonicFormCallback) OnSave() {

	log.Println("HarmonicFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	harmonicFormCallback.probe.formStage.Checkout()

	if harmonicFormCallback.harmonic == nil {
		harmonicFormCallback.harmonic = new(models.Harmonic).Stage(harmonicFormCallback.probe.stageOfInterest)
	}
	harmonic_ := harmonicFormCallback.harmonic
	_ = harmonic_

	for _, formDiv := range harmonicFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(harmonic_.Name), formDiv)
		case "Natural":
			FormDivSelectFieldToField(&(harmonic_.Natural), harmonicFormCallback.probe.stageOfInterest, formDiv)
		case "Artificial":
			FormDivSelectFieldToField(&(harmonic_.Artificial), harmonicFormCallback.probe.stageOfInterest, formDiv)
		case "Base_pitch":
			FormDivSelectFieldToField(&(harmonic_.Base_pitch), harmonicFormCallback.probe.stageOfInterest, formDiv)
		case "Touching_pitch":
			FormDivSelectFieldToField(&(harmonic_.Touching_pitch), harmonicFormCallback.probe.stageOfInterest, formDiv)
		case "Sounding_pitch":
			FormDivSelectFieldToField(&(harmonic_.Sounding_pitch), harmonicFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if harmonicFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmonic_.Unstage(harmonicFormCallback.probe.stageOfInterest)
	}

	harmonicFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Harmonic](
		harmonicFormCallback.probe,
	)
	harmonicFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if harmonicFormCallback.CreationMode || harmonicFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmonicFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(harmonicFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__HarmonicFormCallback(
			nil,
			harmonicFormCallback.probe,
			newFormGroup,
		)
		harmonic := new(models.Harmonic)
		FillUpForm(harmonic, newFormGroup, harmonicFormCallback.probe)
		harmonicFormCallback.probe.formStage.Commit()
	}

	fillUpTree(harmonicFormCallback.probe)
}
func __gong__New__HarmonyFormCallback(
	harmony *models.Harmony,
	probe *Probe,
	formGroup *table.FormGroup,
) (harmonyFormCallback *HarmonyFormCallback) {
	harmonyFormCallback = new(HarmonyFormCallback)
	harmonyFormCallback.probe = probe
	harmonyFormCallback.harmony = harmony
	harmonyFormCallback.formGroup = formGroup

	harmonyFormCallback.CreationMode = (harmony == nil)

	return
}

type HarmonyFormCallback struct {
	harmony *models.Harmony

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (harmonyFormCallback *HarmonyFormCallback) OnSave() {

	log.Println("HarmonyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	harmonyFormCallback.probe.formStage.Checkout()

	if harmonyFormCallback.harmony == nil {
		harmonyFormCallback.harmony = new(models.Harmony).Stage(harmonyFormCallback.probe.stageOfInterest)
	}
	harmony_ := harmonyFormCallback.harmony
	_ = harmony_

	for _, formDiv := range harmonyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(harmony_.Name), formDiv)
		case "Frame":
			FormDivSelectFieldToField(&(harmony_.Frame), harmonyFormCallback.probe.stageOfInterest, formDiv)
		case "Offset":
			FormDivSelectFieldToField(&(harmony_.Offset), harmonyFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if harmonyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmony_.Unstage(harmonyFormCallback.probe.stageOfInterest)
	}

	harmonyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Harmony](
		harmonyFormCallback.probe,
	)
	harmonyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if harmonyFormCallback.CreationMode || harmonyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmonyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(harmonyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__HarmonyFormCallback(
			nil,
			harmonyFormCallback.probe,
			newFormGroup,
		)
		harmony := new(models.Harmony)
		FillUpForm(harmony, newFormGroup, harmonyFormCallback.probe)
		harmonyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(harmonyFormCallback.probe)
}
func __gong__New__Harmony_alterFormCallback(
	harmony_alter *models.Harmony_alter,
	probe *Probe,
	formGroup *table.FormGroup,
) (harmony_alterFormCallback *Harmony_alterFormCallback) {
	harmony_alterFormCallback = new(Harmony_alterFormCallback)
	harmony_alterFormCallback.probe = probe
	harmony_alterFormCallback.harmony_alter = harmony_alter
	harmony_alterFormCallback.formGroup = formGroup

	harmony_alterFormCallback.CreationMode = (harmony_alter == nil)

	return
}

type Harmony_alterFormCallback struct {
	harmony_alter *models.Harmony_alter

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (harmony_alterFormCallback *Harmony_alterFormCallback) OnSave() {

	log.Println("Harmony_alterFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	harmony_alterFormCallback.probe.formStage.Checkout()

	if harmony_alterFormCallback.harmony_alter == nil {
		harmony_alterFormCallback.harmony_alter = new(models.Harmony_alter).Stage(harmony_alterFormCallback.probe.stageOfInterest)
	}
	harmony_alter_ := harmony_alterFormCallback.harmony_alter
	_ = harmony_alter_

	for _, formDiv := range harmony_alterFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(harmony_alter_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if harmony_alterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmony_alter_.Unstage(harmony_alterFormCallback.probe.stageOfInterest)
	}

	harmony_alterFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Harmony_alter](
		harmony_alterFormCallback.probe,
	)
	harmony_alterFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if harmony_alterFormCallback.CreationMode || harmony_alterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harmony_alterFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(harmony_alterFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Harmony_alterFormCallback(
			nil,
			harmony_alterFormCallback.probe,
			newFormGroup,
		)
		harmony_alter := new(models.Harmony_alter)
		FillUpForm(harmony_alter, newFormGroup, harmony_alterFormCallback.probe)
		harmony_alterFormCallback.probe.formStage.Commit()
	}

	fillUpTree(harmony_alterFormCallback.probe)
}
func __gong__New__Harp_pedalsFormCallback(
	harp_pedals *models.Harp_pedals,
	probe *Probe,
	formGroup *table.FormGroup,
) (harp_pedalsFormCallback *Harp_pedalsFormCallback) {
	harp_pedalsFormCallback = new(Harp_pedalsFormCallback)
	harp_pedalsFormCallback.probe = probe
	harp_pedalsFormCallback.harp_pedals = harp_pedals
	harp_pedalsFormCallback.formGroup = formGroup

	harp_pedalsFormCallback.CreationMode = (harp_pedals == nil)

	return
}

type Harp_pedalsFormCallback struct {
	harp_pedals *models.Harp_pedals

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (harp_pedalsFormCallback *Harp_pedalsFormCallback) OnSave() {

	log.Println("Harp_pedalsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	harp_pedalsFormCallback.probe.formStage.Checkout()

	if harp_pedalsFormCallback.harp_pedals == nil {
		harp_pedalsFormCallback.harp_pedals = new(models.Harp_pedals).Stage(harp_pedalsFormCallback.probe.stageOfInterest)
	}
	harp_pedals_ := harp_pedalsFormCallback.harp_pedals
	_ = harp_pedals_

	for _, formDiv := range harp_pedalsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(harp_pedals_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if harp_pedalsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harp_pedals_.Unstage(harp_pedalsFormCallback.probe.stageOfInterest)
	}

	harp_pedalsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Harp_pedals](
		harp_pedalsFormCallback.probe,
	)
	harp_pedalsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if harp_pedalsFormCallback.CreationMode || harp_pedalsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		harp_pedalsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(harp_pedalsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Harp_pedalsFormCallback(
			nil,
			harp_pedalsFormCallback.probe,
			newFormGroup,
		)
		harp_pedals := new(models.Harp_pedals)
		FillUpForm(harp_pedals, newFormGroup, harp_pedalsFormCallback.probe)
		harp_pedalsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(harp_pedalsFormCallback.probe)
}
func __gong__New__Heel_toeFormCallback(
	heel_toe *models.Heel_toe,
	probe *Probe,
	formGroup *table.FormGroup,
) (heel_toeFormCallback *Heel_toeFormCallback) {
	heel_toeFormCallback = new(Heel_toeFormCallback)
	heel_toeFormCallback.probe = probe
	heel_toeFormCallback.heel_toe = heel_toe
	heel_toeFormCallback.formGroup = formGroup

	heel_toeFormCallback.CreationMode = (heel_toe == nil)

	return
}

type Heel_toeFormCallback struct {
	heel_toe *models.Heel_toe

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (heel_toeFormCallback *Heel_toeFormCallback) OnSave() {

	log.Println("Heel_toeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	heel_toeFormCallback.probe.formStage.Checkout()

	if heel_toeFormCallback.heel_toe == nil {
		heel_toeFormCallback.heel_toe = new(models.Heel_toe).Stage(heel_toeFormCallback.probe.stageOfInterest)
	}
	heel_toe_ := heel_toeFormCallback.heel_toe
	_ = heel_toe_

	for _, formDiv := range heel_toeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(heel_toe_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if heel_toeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		heel_toe_.Unstage(heel_toeFormCallback.probe.stageOfInterest)
	}

	heel_toeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Heel_toe](
		heel_toeFormCallback.probe,
	)
	heel_toeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if heel_toeFormCallback.CreationMode || heel_toeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		heel_toeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(heel_toeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Heel_toeFormCallback(
			nil,
			heel_toeFormCallback.probe,
			newFormGroup,
		)
		heel_toe := new(models.Heel_toe)
		FillUpForm(heel_toe, newFormGroup, heel_toeFormCallback.probe)
		heel_toeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(heel_toeFormCallback.probe)
}
func __gong__New__HoleFormCallback(
	hole *models.Hole,
	probe *Probe,
	formGroup *table.FormGroup,
) (holeFormCallback *HoleFormCallback) {
	holeFormCallback = new(HoleFormCallback)
	holeFormCallback.probe = probe
	holeFormCallback.hole = hole
	holeFormCallback.formGroup = formGroup

	holeFormCallback.CreationMode = (hole == nil)

	return
}

type HoleFormCallback struct {
	hole *models.Hole

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (holeFormCallback *HoleFormCallback) OnSave() {

	log.Println("HoleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	holeFormCallback.probe.formStage.Checkout()

	if holeFormCallback.hole == nil {
		holeFormCallback.hole = new(models.Hole).Stage(holeFormCallback.probe.stageOfInterest)
	}
	hole_ := holeFormCallback.hole
	_ = hole_

	for _, formDiv := range holeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(hole_.Name), formDiv)
		case "Hole_type":
			FormDivBasicFieldToField(&(hole_.Hole_type), formDiv)
		case "Hole_closed":
			FormDivSelectFieldToField(&(hole_.Hole_closed), holeFormCallback.probe.stageOfInterest, formDiv)
		case "Hole_shape":
			FormDivBasicFieldToField(&(hole_.Hole_shape), formDiv)
		}
	}

	// manage the suppress operation
	if holeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		hole_.Unstage(holeFormCallback.probe.stageOfInterest)
	}

	holeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Hole](
		holeFormCallback.probe,
	)
	holeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if holeFormCallback.CreationMode || holeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		holeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(holeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__HoleFormCallback(
			nil,
			holeFormCallback.probe,
			newFormGroup,
		)
		hole := new(models.Hole)
		FillUpForm(hole, newFormGroup, holeFormCallback.probe)
		holeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(holeFormCallback.probe)
}
func __gong__New__Hole_closedFormCallback(
	hole_closed *models.Hole_closed,
	probe *Probe,
	formGroup *table.FormGroup,
) (hole_closedFormCallback *Hole_closedFormCallback) {
	hole_closedFormCallback = new(Hole_closedFormCallback)
	hole_closedFormCallback.probe = probe
	hole_closedFormCallback.hole_closed = hole_closed
	hole_closedFormCallback.formGroup = formGroup

	hole_closedFormCallback.CreationMode = (hole_closed == nil)

	return
}

type Hole_closedFormCallback struct {
	hole_closed *models.Hole_closed

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (hole_closedFormCallback *Hole_closedFormCallback) OnSave() {

	log.Println("Hole_closedFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	hole_closedFormCallback.probe.formStage.Checkout()

	if hole_closedFormCallback.hole_closed == nil {
		hole_closedFormCallback.hole_closed = new(models.Hole_closed).Stage(hole_closedFormCallback.probe.stageOfInterest)
	}
	hole_closed_ := hole_closedFormCallback.hole_closed
	_ = hole_closed_

	for _, formDiv := range hole_closedFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(hole_closed_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if hole_closedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		hole_closed_.Unstage(hole_closedFormCallback.probe.stageOfInterest)
	}

	hole_closedFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Hole_closed](
		hole_closedFormCallback.probe,
	)
	hole_closedFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if hole_closedFormCallback.CreationMode || hole_closedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		hole_closedFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(hole_closedFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Hole_closedFormCallback(
			nil,
			hole_closedFormCallback.probe,
			newFormGroup,
		)
		hole_closed := new(models.Hole_closed)
		FillUpForm(hole_closed, newFormGroup, hole_closedFormCallback.probe)
		hole_closedFormCallback.probe.formStage.Commit()
	}

	fillUpTree(hole_closedFormCallback.probe)
}
func __gong__New__Horizontal_turnFormCallback(
	horizontal_turn *models.Horizontal_turn,
	probe *Probe,
	formGroup *table.FormGroup,
) (horizontal_turnFormCallback *Horizontal_turnFormCallback) {
	horizontal_turnFormCallback = new(Horizontal_turnFormCallback)
	horizontal_turnFormCallback.probe = probe
	horizontal_turnFormCallback.horizontal_turn = horizontal_turn
	horizontal_turnFormCallback.formGroup = formGroup

	horizontal_turnFormCallback.CreationMode = (horizontal_turn == nil)

	return
}

type Horizontal_turnFormCallback struct {
	horizontal_turn *models.Horizontal_turn

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (horizontal_turnFormCallback *Horizontal_turnFormCallback) OnSave() {

	log.Println("Horizontal_turnFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	horizontal_turnFormCallback.probe.formStage.Checkout()

	if horizontal_turnFormCallback.horizontal_turn == nil {
		horizontal_turnFormCallback.horizontal_turn = new(models.Horizontal_turn).Stage(horizontal_turnFormCallback.probe.stageOfInterest)
	}
	horizontal_turn_ := horizontal_turnFormCallback.horizontal_turn
	_ = horizontal_turn_

	for _, formDiv := range horizontal_turnFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(horizontal_turn_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if horizontal_turnFormCallback.formGroup.HasSuppressButtonBeenPressed {
		horizontal_turn_.Unstage(horizontal_turnFormCallback.probe.stageOfInterest)
	}

	horizontal_turnFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Horizontal_turn](
		horizontal_turnFormCallback.probe,
	)
	horizontal_turnFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if horizontal_turnFormCallback.CreationMode || horizontal_turnFormCallback.formGroup.HasSuppressButtonBeenPressed {
		horizontal_turnFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(horizontal_turnFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Horizontal_turnFormCallback(
			nil,
			horizontal_turnFormCallback.probe,
			newFormGroup,
		)
		horizontal_turn := new(models.Horizontal_turn)
		FillUpForm(horizontal_turn, newFormGroup, horizontal_turnFormCallback.probe)
		horizontal_turnFormCallback.probe.formStage.Commit()
	}

	fillUpTree(horizontal_turnFormCallback.probe)
}
func __gong__New__IdentificationFormCallback(
	identification *models.Identification,
	probe *Probe,
	formGroup *table.FormGroup,
) (identificationFormCallback *IdentificationFormCallback) {
	identificationFormCallback = new(IdentificationFormCallback)
	identificationFormCallback.probe = probe
	identificationFormCallback.identification = identification
	identificationFormCallback.formGroup = formGroup

	identificationFormCallback.CreationMode = (identification == nil)

	return
}

type IdentificationFormCallback struct {
	identification *models.Identification

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (identificationFormCallback *IdentificationFormCallback) OnSave() {

	log.Println("IdentificationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	identificationFormCallback.probe.formStage.Checkout()

	if identificationFormCallback.identification == nil {
		identificationFormCallback.identification = new(models.Identification).Stage(identificationFormCallback.probe.stageOfInterest)
	}
	identification_ := identificationFormCallback.identification
	_ = identification_

	for _, formDiv := range identificationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(identification_.Name), formDiv)
		case "Encoding":
			FormDivSelectFieldToField(&(identification_.Encoding), identificationFormCallback.probe.stageOfInterest, formDiv)
		case "Source":
			FormDivBasicFieldToField(&(identification_.Source), formDiv)
		case "Miscellaneous":
			FormDivSelectFieldToField(&(identification_.Miscellaneous), identificationFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if identificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		identification_.Unstage(identificationFormCallback.probe.stageOfInterest)
	}

	identificationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Identification](
		identificationFormCallback.probe,
	)
	identificationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if identificationFormCallback.CreationMode || identificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		identificationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(identificationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__IdentificationFormCallback(
			nil,
			identificationFormCallback.probe,
			newFormGroup,
		)
		identification := new(models.Identification)
		FillUpForm(identification, newFormGroup, identificationFormCallback.probe)
		identificationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(identificationFormCallback.probe)
}
func __gong__New__ImageFormCallback(
	image *models.Image,
	probe *Probe,
	formGroup *table.FormGroup,
) (imageFormCallback *ImageFormCallback) {
	imageFormCallback = new(ImageFormCallback)
	imageFormCallback.probe = probe
	imageFormCallback.image = image
	imageFormCallback.formGroup = formGroup

	imageFormCallback.CreationMode = (image == nil)

	return
}

type ImageFormCallback struct {
	image *models.Image

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (imageFormCallback *ImageFormCallback) OnSave() {

	log.Println("ImageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	imageFormCallback.probe.formStage.Checkout()

	if imageFormCallback.image == nil {
		imageFormCallback.image = new(models.Image).Stage(imageFormCallback.probe.stageOfInterest)
	}
	image_ := imageFormCallback.image
	_ = image_

	for _, formDiv := range imageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(image_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if imageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		image_.Unstage(imageFormCallback.probe.stageOfInterest)
	}

	imageFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Image](
		imageFormCallback.probe,
	)
	imageFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if imageFormCallback.CreationMode || imageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		imageFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(imageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ImageFormCallback(
			nil,
			imageFormCallback.probe,
			newFormGroup,
		)
		image := new(models.Image)
		FillUpForm(image, newFormGroup, imageFormCallback.probe)
		imageFormCallback.probe.formStage.Commit()
	}

	fillUpTree(imageFormCallback.probe)
}
func __gong__New__InstrumentFormCallback(
	instrument *models.Instrument,
	probe *Probe,
	formGroup *table.FormGroup,
) (instrumentFormCallback *InstrumentFormCallback) {
	instrumentFormCallback = new(InstrumentFormCallback)
	instrumentFormCallback.probe = probe
	instrumentFormCallback.instrument = instrument
	instrumentFormCallback.formGroup = formGroup

	instrumentFormCallback.CreationMode = (instrument == nil)

	return
}

type InstrumentFormCallback struct {
	instrument *models.Instrument

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (instrumentFormCallback *InstrumentFormCallback) OnSave() {

	log.Println("InstrumentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	instrumentFormCallback.probe.formStage.Checkout()

	if instrumentFormCallback.instrument == nil {
		instrumentFormCallback.instrument = new(models.Instrument).Stage(instrumentFormCallback.probe.stageOfInterest)
	}
	instrument_ := instrumentFormCallback.instrument
	_ = instrument_

	for _, formDiv := range instrumentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(instrument_.Name), formDiv)
		case "Note:Instrument":
			// we need to retrieve the field owner before the change
			var pastNoteOwner *models.Note
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Instrument"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				instrumentFormCallback.probe.stageOfInterest,
				instrumentFormCallback.probe.backRepoOfInterest,
				instrument_,
				&rf)

			if reverseFieldOwner != nil {
				pastNoteOwner = reverseFieldOwner.(*models.Note)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNoteOwner != nil {
					idx := slices.Index(pastNoteOwner.Instrument, instrument_)
					pastNoteOwner.Instrument = slices.Delete(pastNoteOwner.Instrument, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _note := range *models.GetGongstructInstancesSet[models.Note](instrumentFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _note.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNoteOwner := _note // we have a match
						if pastNoteOwner != nil {
							if newNoteOwner != pastNoteOwner {
								idx := slices.Index(pastNoteOwner.Instrument, instrument_)
								pastNoteOwner.Instrument = slices.Delete(pastNoteOwner.Instrument, idx, idx+1)
								newNoteOwner.Instrument = append(newNoteOwner.Instrument, instrument_)
							}
						} else {
							newNoteOwner.Instrument = append(newNoteOwner.Instrument, instrument_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if instrumentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		instrument_.Unstage(instrumentFormCallback.probe.stageOfInterest)
	}

	instrumentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Instrument](
		instrumentFormCallback.probe,
	)
	instrumentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if instrumentFormCallback.CreationMode || instrumentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		instrumentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(instrumentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__InstrumentFormCallback(
			nil,
			instrumentFormCallback.probe,
			newFormGroup,
		)
		instrument := new(models.Instrument)
		FillUpForm(instrument, newFormGroup, instrumentFormCallback.probe)
		instrumentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(instrumentFormCallback.probe)
}
func __gong__New__Instrument_changeFormCallback(
	instrument_change *models.Instrument_change,
	probe *Probe,
	formGroup *table.FormGroup,
) (instrument_changeFormCallback *Instrument_changeFormCallback) {
	instrument_changeFormCallback = new(Instrument_changeFormCallback)
	instrument_changeFormCallback.probe = probe
	instrument_changeFormCallback.instrument_change = instrument_change
	instrument_changeFormCallback.formGroup = formGroup

	instrument_changeFormCallback.CreationMode = (instrument_change == nil)

	return
}

type Instrument_changeFormCallback struct {
	instrument_change *models.Instrument_change

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (instrument_changeFormCallback *Instrument_changeFormCallback) OnSave() {

	log.Println("Instrument_changeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	instrument_changeFormCallback.probe.formStage.Checkout()

	if instrument_changeFormCallback.instrument_change == nil {
		instrument_changeFormCallback.instrument_change = new(models.Instrument_change).Stage(instrument_changeFormCallback.probe.stageOfInterest)
	}
	instrument_change_ := instrument_changeFormCallback.instrument_change
	_ = instrument_change_

	for _, formDiv := range instrument_changeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(instrument_change_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if instrument_changeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		instrument_change_.Unstage(instrument_changeFormCallback.probe.stageOfInterest)
	}

	instrument_changeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Instrument_change](
		instrument_changeFormCallback.probe,
	)
	instrument_changeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if instrument_changeFormCallback.CreationMode || instrument_changeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		instrument_changeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(instrument_changeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Instrument_changeFormCallback(
			nil,
			instrument_changeFormCallback.probe,
			newFormGroup,
		)
		instrument_change := new(models.Instrument_change)
		FillUpForm(instrument_change, newFormGroup, instrument_changeFormCallback.probe)
		instrument_changeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(instrument_changeFormCallback.probe)
}
func __gong__New__Instrument_linkFormCallback(
	instrument_link *models.Instrument_link,
	probe *Probe,
	formGroup *table.FormGroup,
) (instrument_linkFormCallback *Instrument_linkFormCallback) {
	instrument_linkFormCallback = new(Instrument_linkFormCallback)
	instrument_linkFormCallback.probe = probe
	instrument_linkFormCallback.instrument_link = instrument_link
	instrument_linkFormCallback.formGroup = formGroup

	instrument_linkFormCallback.CreationMode = (instrument_link == nil)

	return
}

type Instrument_linkFormCallback struct {
	instrument_link *models.Instrument_link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (instrument_linkFormCallback *Instrument_linkFormCallback) OnSave() {

	log.Println("Instrument_linkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	instrument_linkFormCallback.probe.formStage.Checkout()

	if instrument_linkFormCallback.instrument_link == nil {
		instrument_linkFormCallback.instrument_link = new(models.Instrument_link).Stage(instrument_linkFormCallback.probe.stageOfInterest)
	}
	instrument_link_ := instrument_linkFormCallback.instrument_link
	_ = instrument_link_

	for _, formDiv := range instrument_linkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(instrument_link_.Name), formDiv)
		case "Part_link:Instrument_link":
			// we need to retrieve the field owner before the change
			var pastPart_linkOwner *models.Part_link
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Part_link"
			rf.Fieldname = "Instrument_link"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				instrument_linkFormCallback.probe.stageOfInterest,
				instrument_linkFormCallback.probe.backRepoOfInterest,
				instrument_link_,
				&rf)

			if reverseFieldOwner != nil {
				pastPart_linkOwner = reverseFieldOwner.(*models.Part_link)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastPart_linkOwner != nil {
					idx := slices.Index(pastPart_linkOwner.Instrument_link, instrument_link_)
					pastPart_linkOwner.Instrument_link = slices.Delete(pastPart_linkOwner.Instrument_link, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _part_link := range *models.GetGongstructInstancesSet[models.Part_link](instrument_linkFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _part_link.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newPart_linkOwner := _part_link // we have a match
						if pastPart_linkOwner != nil {
							if newPart_linkOwner != pastPart_linkOwner {
								idx := slices.Index(pastPart_linkOwner.Instrument_link, instrument_link_)
								pastPart_linkOwner.Instrument_link = slices.Delete(pastPart_linkOwner.Instrument_link, idx, idx+1)
								newPart_linkOwner.Instrument_link = append(newPart_linkOwner.Instrument_link, instrument_link_)
							}
						} else {
							newPart_linkOwner.Instrument_link = append(newPart_linkOwner.Instrument_link, instrument_link_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if instrument_linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		instrument_link_.Unstage(instrument_linkFormCallback.probe.stageOfInterest)
	}

	instrument_linkFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Instrument_link](
		instrument_linkFormCallback.probe,
	)
	instrument_linkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if instrument_linkFormCallback.CreationMode || instrument_linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		instrument_linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(instrument_linkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Instrument_linkFormCallback(
			nil,
			instrument_linkFormCallback.probe,
			newFormGroup,
		)
		instrument_link := new(models.Instrument_link)
		FillUpForm(instrument_link, newFormGroup, instrument_linkFormCallback.probe)
		instrument_linkFormCallback.probe.formStage.Commit()
	}

	fillUpTree(instrument_linkFormCallback.probe)
}
func __gong__New__InterchangeableFormCallback(
	interchangeable *models.Interchangeable,
	probe *Probe,
	formGroup *table.FormGroup,
) (interchangeableFormCallback *InterchangeableFormCallback) {
	interchangeableFormCallback = new(InterchangeableFormCallback)
	interchangeableFormCallback.probe = probe
	interchangeableFormCallback.interchangeable = interchangeable
	interchangeableFormCallback.formGroup = formGroup

	interchangeableFormCallback.CreationMode = (interchangeable == nil)

	return
}

type InterchangeableFormCallback struct {
	interchangeable *models.Interchangeable

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (interchangeableFormCallback *InterchangeableFormCallback) OnSave() {

	log.Println("InterchangeableFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	interchangeableFormCallback.probe.formStage.Checkout()

	if interchangeableFormCallback.interchangeable == nil {
		interchangeableFormCallback.interchangeable = new(models.Interchangeable).Stage(interchangeableFormCallback.probe.stageOfInterest)
	}
	interchangeable_ := interchangeableFormCallback.interchangeable
	_ = interchangeable_

	for _, formDiv := range interchangeableFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(interchangeable_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if interchangeableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		interchangeable_.Unstage(interchangeableFormCallback.probe.stageOfInterest)
	}

	interchangeableFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Interchangeable](
		interchangeableFormCallback.probe,
	)
	interchangeableFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if interchangeableFormCallback.CreationMode || interchangeableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		interchangeableFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(interchangeableFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__InterchangeableFormCallback(
			nil,
			interchangeableFormCallback.probe,
			newFormGroup,
		)
		interchangeable := new(models.Interchangeable)
		FillUpForm(interchangeable, newFormGroup, interchangeableFormCallback.probe)
		interchangeableFormCallback.probe.formStage.Commit()
	}

	fillUpTree(interchangeableFormCallback.probe)
}
func __gong__New__InversionFormCallback(
	inversion *models.Inversion,
	probe *Probe,
	formGroup *table.FormGroup,
) (inversionFormCallback *InversionFormCallback) {
	inversionFormCallback = new(InversionFormCallback)
	inversionFormCallback.probe = probe
	inversionFormCallback.inversion = inversion
	inversionFormCallback.formGroup = formGroup

	inversionFormCallback.CreationMode = (inversion == nil)

	return
}

type InversionFormCallback struct {
	inversion *models.Inversion

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (inversionFormCallback *InversionFormCallback) OnSave() {

	log.Println("InversionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	inversionFormCallback.probe.formStage.Checkout()

	if inversionFormCallback.inversion == nil {
		inversionFormCallback.inversion = new(models.Inversion).Stage(inversionFormCallback.probe.stageOfInterest)
	}
	inversion_ := inversionFormCallback.inversion
	_ = inversion_

	for _, formDiv := range inversionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(inversion_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(inversion_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if inversionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		inversion_.Unstage(inversionFormCallback.probe.stageOfInterest)
	}

	inversionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Inversion](
		inversionFormCallback.probe,
	)
	inversionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if inversionFormCallback.CreationMode || inversionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		inversionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(inversionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__InversionFormCallback(
			nil,
			inversionFormCallback.probe,
			newFormGroup,
		)
		inversion := new(models.Inversion)
		FillUpForm(inversion, newFormGroup, inversionFormCallback.probe)
		inversionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(inversionFormCallback.probe)
}
func __gong__New__KeyFormCallback(
	key *models.Key,
	probe *Probe,
	formGroup *table.FormGroup,
) (keyFormCallback *KeyFormCallback) {
	keyFormCallback = new(KeyFormCallback)
	keyFormCallback.probe = probe
	keyFormCallback.key = key
	keyFormCallback.formGroup = formGroup

	keyFormCallback.CreationMode = (key == nil)

	return
}

type KeyFormCallback struct {
	key *models.Key

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (keyFormCallback *KeyFormCallback) OnSave() {

	log.Println("KeyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	keyFormCallback.probe.formStage.Checkout()

	if keyFormCallback.key == nil {
		keyFormCallback.key = new(models.Key).Stage(keyFormCallback.probe.stageOfInterest)
	}
	key_ := keyFormCallback.key
	_ = key_

	for _, formDiv := range keyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(key_.Name), formDiv)
		case "Attributes:Key":
			// we need to retrieve the field owner before the change
			var pastAttributesOwner *models.Attributes
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Key"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				keyFormCallback.probe.stageOfInterest,
				keyFormCallback.probe.backRepoOfInterest,
				key_,
				&rf)

			if reverseFieldOwner != nil {
				pastAttributesOwner = reverseFieldOwner.(*models.Attributes)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAttributesOwner != nil {
					idx := slices.Index(pastAttributesOwner.Key, key_)
					pastAttributesOwner.Key = slices.Delete(pastAttributesOwner.Key, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attributes := range *models.GetGongstructInstancesSet[models.Attributes](keyFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAttributesOwner := _attributes // we have a match
						if pastAttributesOwner != nil {
							if newAttributesOwner != pastAttributesOwner {
								idx := slices.Index(pastAttributesOwner.Key, key_)
								pastAttributesOwner.Key = slices.Delete(pastAttributesOwner.Key, idx, idx+1)
								newAttributesOwner.Key = append(newAttributesOwner.Key, key_)
							}
						} else {
							newAttributesOwner.Key = append(newAttributesOwner.Key, key_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if keyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		key_.Unstage(keyFormCallback.probe.stageOfInterest)
	}

	keyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Key](
		keyFormCallback.probe,
	)
	keyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if keyFormCallback.CreationMode || keyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		keyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(keyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__KeyFormCallback(
			nil,
			keyFormCallback.probe,
			newFormGroup,
		)
		key := new(models.Key)
		FillUpForm(key, newFormGroup, keyFormCallback.probe)
		keyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(keyFormCallback.probe)
}
func __gong__New__Key_accidentalFormCallback(
	key_accidental *models.Key_accidental,
	probe *Probe,
	formGroup *table.FormGroup,
) (key_accidentalFormCallback *Key_accidentalFormCallback) {
	key_accidentalFormCallback = new(Key_accidentalFormCallback)
	key_accidentalFormCallback.probe = probe
	key_accidentalFormCallback.key_accidental = key_accidental
	key_accidentalFormCallback.formGroup = formGroup

	key_accidentalFormCallback.CreationMode = (key_accidental == nil)

	return
}

type Key_accidentalFormCallback struct {
	key_accidental *models.Key_accidental

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (key_accidentalFormCallback *Key_accidentalFormCallback) OnSave() {

	log.Println("Key_accidentalFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	key_accidentalFormCallback.probe.formStage.Checkout()

	if key_accidentalFormCallback.key_accidental == nil {
		key_accidentalFormCallback.key_accidental = new(models.Key_accidental).Stage(key_accidentalFormCallback.probe.stageOfInterest)
	}
	key_accidental_ := key_accidentalFormCallback.key_accidental
	_ = key_accidental_

	for _, formDiv := range key_accidentalFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(key_accidental_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if key_accidentalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		key_accidental_.Unstage(key_accidentalFormCallback.probe.stageOfInterest)
	}

	key_accidentalFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Key_accidental](
		key_accidentalFormCallback.probe,
	)
	key_accidentalFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if key_accidentalFormCallback.CreationMode || key_accidentalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		key_accidentalFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(key_accidentalFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Key_accidentalFormCallback(
			nil,
			key_accidentalFormCallback.probe,
			newFormGroup,
		)
		key_accidental := new(models.Key_accidental)
		FillUpForm(key_accidental, newFormGroup, key_accidentalFormCallback.probe)
		key_accidentalFormCallback.probe.formStage.Commit()
	}

	fillUpTree(key_accidentalFormCallback.probe)
}
func __gong__New__Key_octaveFormCallback(
	key_octave *models.Key_octave,
	probe *Probe,
	formGroup *table.FormGroup,
) (key_octaveFormCallback *Key_octaveFormCallback) {
	key_octaveFormCallback = new(Key_octaveFormCallback)
	key_octaveFormCallback.probe = probe
	key_octaveFormCallback.key_octave = key_octave
	key_octaveFormCallback.formGroup = formGroup

	key_octaveFormCallback.CreationMode = (key_octave == nil)

	return
}

type Key_octaveFormCallback struct {
	key_octave *models.Key_octave

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (key_octaveFormCallback *Key_octaveFormCallback) OnSave() {

	log.Println("Key_octaveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	key_octaveFormCallback.probe.formStage.Checkout()

	if key_octaveFormCallback.key_octave == nil {
		key_octaveFormCallback.key_octave = new(models.Key_octave).Stage(key_octaveFormCallback.probe.stageOfInterest)
	}
	key_octave_ := key_octaveFormCallback.key_octave
	_ = key_octave_

	for _, formDiv := range key_octaveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(key_octave_.Name), formDiv)
		case "Key:Key_octave":
			// we need to retrieve the field owner before the change
			var pastKeyOwner *models.Key
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Key"
			rf.Fieldname = "Key_octave"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				key_octaveFormCallback.probe.stageOfInterest,
				key_octaveFormCallback.probe.backRepoOfInterest,
				key_octave_,
				&rf)

			if reverseFieldOwner != nil {
				pastKeyOwner = reverseFieldOwner.(*models.Key)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastKeyOwner != nil {
					idx := slices.Index(pastKeyOwner.Key_octave, key_octave_)
					pastKeyOwner.Key_octave = slices.Delete(pastKeyOwner.Key_octave, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _key := range *models.GetGongstructInstancesSet[models.Key](key_octaveFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _key.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newKeyOwner := _key // we have a match
						if pastKeyOwner != nil {
							if newKeyOwner != pastKeyOwner {
								idx := slices.Index(pastKeyOwner.Key_octave, key_octave_)
								pastKeyOwner.Key_octave = slices.Delete(pastKeyOwner.Key_octave, idx, idx+1)
								newKeyOwner.Key_octave = append(newKeyOwner.Key_octave, key_octave_)
							}
						} else {
							newKeyOwner.Key_octave = append(newKeyOwner.Key_octave, key_octave_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if key_octaveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		key_octave_.Unstage(key_octaveFormCallback.probe.stageOfInterest)
	}

	key_octaveFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Key_octave](
		key_octaveFormCallback.probe,
	)
	key_octaveFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if key_octaveFormCallback.CreationMode || key_octaveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		key_octaveFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(key_octaveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Key_octaveFormCallback(
			nil,
			key_octaveFormCallback.probe,
			newFormGroup,
		)
		key_octave := new(models.Key_octave)
		FillUpForm(key_octave, newFormGroup, key_octaveFormCallback.probe)
		key_octaveFormCallback.probe.formStage.Commit()
	}

	fillUpTree(key_octaveFormCallback.probe)
}
func __gong__New__KindFormCallback(
	kind *models.Kind,
	probe *Probe,
	formGroup *table.FormGroup,
) (kindFormCallback *KindFormCallback) {
	kindFormCallback = new(KindFormCallback)
	kindFormCallback.probe = probe
	kindFormCallback.kind = kind
	kindFormCallback.formGroup = formGroup

	kindFormCallback.CreationMode = (kind == nil)

	return
}

type KindFormCallback struct {
	kind *models.Kind

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (kindFormCallback *KindFormCallback) OnSave() {

	log.Println("KindFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	kindFormCallback.probe.formStage.Checkout()

	if kindFormCallback.kind == nil {
		kindFormCallback.kind = new(models.Kind).Stage(kindFormCallback.probe.stageOfInterest)
	}
	kind_ := kindFormCallback.kind
	_ = kind_

	for _, formDiv := range kindFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(kind_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(kind_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if kindFormCallback.formGroup.HasSuppressButtonBeenPressed {
		kind_.Unstage(kindFormCallback.probe.stageOfInterest)
	}

	kindFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Kind](
		kindFormCallback.probe,
	)
	kindFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if kindFormCallback.CreationMode || kindFormCallback.formGroup.HasSuppressButtonBeenPressed {
		kindFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(kindFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__KindFormCallback(
			nil,
			kindFormCallback.probe,
			newFormGroup,
		)
		kind := new(models.Kind)
		FillUpForm(kind, newFormGroup, kindFormCallback.probe)
		kindFormCallback.probe.formStage.Commit()
	}

	fillUpTree(kindFormCallback.probe)
}
func __gong__New__LevelFormCallback(
	level *models.Level,
	probe *Probe,
	formGroup *table.FormGroup,
) (levelFormCallback *LevelFormCallback) {
	levelFormCallback = new(LevelFormCallback)
	levelFormCallback.probe = probe
	levelFormCallback.level = level
	levelFormCallback.formGroup = formGroup

	levelFormCallback.CreationMode = (level == nil)

	return
}

type LevelFormCallback struct {
	level *models.Level

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (levelFormCallback *LevelFormCallback) OnSave() {

	log.Println("LevelFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	levelFormCallback.probe.formStage.Checkout()

	if levelFormCallback.level == nil {
		levelFormCallback.level = new(models.Level).Stage(levelFormCallback.probe.stageOfInterest)
	}
	level_ := levelFormCallback.level
	_ = level_

	for _, formDiv := range levelFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(level_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(level_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if levelFormCallback.formGroup.HasSuppressButtonBeenPressed {
		level_.Unstage(levelFormCallback.probe.stageOfInterest)
	}

	levelFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Level](
		levelFormCallback.probe,
	)
	levelFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if levelFormCallback.CreationMode || levelFormCallback.formGroup.HasSuppressButtonBeenPressed {
		levelFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(levelFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LevelFormCallback(
			nil,
			levelFormCallback.probe,
			newFormGroup,
		)
		level := new(models.Level)
		FillUpForm(level, newFormGroup, levelFormCallback.probe)
		levelFormCallback.probe.formStage.Commit()
	}

	fillUpTree(levelFormCallback.probe)
}
func __gong__New__Line_detailFormCallback(
	line_detail *models.Line_detail,
	probe *Probe,
	formGroup *table.FormGroup,
) (line_detailFormCallback *Line_detailFormCallback) {
	line_detailFormCallback = new(Line_detailFormCallback)
	line_detailFormCallback.probe = probe
	line_detailFormCallback.line_detail = line_detail
	line_detailFormCallback.formGroup = formGroup

	line_detailFormCallback.CreationMode = (line_detail == nil)

	return
}

type Line_detailFormCallback struct {
	line_detail *models.Line_detail

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (line_detailFormCallback *Line_detailFormCallback) OnSave() {

	log.Println("Line_detailFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	line_detailFormCallback.probe.formStage.Checkout()

	if line_detailFormCallback.line_detail == nil {
		line_detailFormCallback.line_detail = new(models.Line_detail).Stage(line_detailFormCallback.probe.stageOfInterest)
	}
	line_detail_ := line_detailFormCallback.line_detail
	_ = line_detail_

	for _, formDiv := range line_detailFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(line_detail_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if line_detailFormCallback.formGroup.HasSuppressButtonBeenPressed {
		line_detail_.Unstage(line_detailFormCallback.probe.stageOfInterest)
	}

	line_detailFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Line_detail](
		line_detailFormCallback.probe,
	)
	line_detailFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if line_detailFormCallback.CreationMode || line_detailFormCallback.formGroup.HasSuppressButtonBeenPressed {
		line_detailFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(line_detailFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Line_detailFormCallback(
			nil,
			line_detailFormCallback.probe,
			newFormGroup,
		)
		line_detail := new(models.Line_detail)
		FillUpForm(line_detail, newFormGroup, line_detailFormCallback.probe)
		line_detailFormCallback.probe.formStage.Commit()
	}

	fillUpTree(line_detailFormCallback.probe)
}
func __gong__New__Line_widthFormCallback(
	line_width *models.Line_width,
	probe *Probe,
	formGroup *table.FormGroup,
) (line_widthFormCallback *Line_widthFormCallback) {
	line_widthFormCallback = new(Line_widthFormCallback)
	line_widthFormCallback.probe = probe
	line_widthFormCallback.line_width = line_width
	line_widthFormCallback.formGroup = formGroup

	line_widthFormCallback.CreationMode = (line_width == nil)

	return
}

type Line_widthFormCallback struct {
	line_width *models.Line_width

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (line_widthFormCallback *Line_widthFormCallback) OnSave() {

	log.Println("Line_widthFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	line_widthFormCallback.probe.formStage.Checkout()

	if line_widthFormCallback.line_width == nil {
		line_widthFormCallback.line_width = new(models.Line_width).Stage(line_widthFormCallback.probe.stageOfInterest)
	}
	line_width_ := line_widthFormCallback.line_width
	_ = line_width_

	for _, formDiv := range line_widthFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(line_width_.Name), formDiv)
		case "Appearance:Line_width":
			// we need to retrieve the field owner before the change
			var pastAppearanceOwner *models.Appearance
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Line_width"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				line_widthFormCallback.probe.stageOfInterest,
				line_widthFormCallback.probe.backRepoOfInterest,
				line_width_,
				&rf)

			if reverseFieldOwner != nil {
				pastAppearanceOwner = reverseFieldOwner.(*models.Appearance)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAppearanceOwner != nil {
					idx := slices.Index(pastAppearanceOwner.Line_width, line_width_)
					pastAppearanceOwner.Line_width = slices.Delete(pastAppearanceOwner.Line_width, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _appearance := range *models.GetGongstructInstancesSet[models.Appearance](line_widthFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _appearance.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAppearanceOwner := _appearance // we have a match
						if pastAppearanceOwner != nil {
							if newAppearanceOwner != pastAppearanceOwner {
								idx := slices.Index(pastAppearanceOwner.Line_width, line_width_)
								pastAppearanceOwner.Line_width = slices.Delete(pastAppearanceOwner.Line_width, idx, idx+1)
								newAppearanceOwner.Line_width = append(newAppearanceOwner.Line_width, line_width_)
							}
						} else {
							newAppearanceOwner.Line_width = append(newAppearanceOwner.Line_width, line_width_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if line_widthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		line_width_.Unstage(line_widthFormCallback.probe.stageOfInterest)
	}

	line_widthFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Line_width](
		line_widthFormCallback.probe,
	)
	line_widthFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if line_widthFormCallback.CreationMode || line_widthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		line_widthFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(line_widthFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Line_widthFormCallback(
			nil,
			line_widthFormCallback.probe,
			newFormGroup,
		)
		line_width := new(models.Line_width)
		FillUpForm(line_width, newFormGroup, line_widthFormCallback.probe)
		line_widthFormCallback.probe.formStage.Commit()
	}

	fillUpTree(line_widthFormCallback.probe)
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
		case "Credit:Link":
			// we need to retrieve the field owner before the change
			var pastCreditOwner *models.Credit
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Credit"
			rf.Fieldname = "Link"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				linkFormCallback.probe.stageOfInterest,
				linkFormCallback.probe.backRepoOfInterest,
				link_,
				&rf)

			if reverseFieldOwner != nil {
				pastCreditOwner = reverseFieldOwner.(*models.Credit)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastCreditOwner != nil {
					idx := slices.Index(pastCreditOwner.Link, link_)
					pastCreditOwner.Link = slices.Delete(pastCreditOwner.Link, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _credit := range *models.GetGongstructInstancesSet[models.Credit](linkFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _credit.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newCreditOwner := _credit // we have a match
						if pastCreditOwner != nil {
							if newCreditOwner != pastCreditOwner {
								idx := slices.Index(pastCreditOwner.Link, link_)
								pastCreditOwner.Link = slices.Delete(pastCreditOwner.Link, idx, idx+1)
								newCreditOwner.Link = append(newCreditOwner.Link, link_)
							}
						} else {
							newCreditOwner.Link = append(newCreditOwner.Link, link_)
						}
					}
				}
			}
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
func __gong__New__ListenFormCallback(
	listen *models.Listen,
	probe *Probe,
	formGroup *table.FormGroup,
) (listenFormCallback *ListenFormCallback) {
	listenFormCallback = new(ListenFormCallback)
	listenFormCallback.probe = probe
	listenFormCallback.listen = listen
	listenFormCallback.formGroup = formGroup

	listenFormCallback.CreationMode = (listen == nil)

	return
}

type ListenFormCallback struct {
	listen *models.Listen

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (listenFormCallback *ListenFormCallback) OnSave() {

	log.Println("ListenFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	listenFormCallback.probe.formStage.Checkout()

	if listenFormCallback.listen == nil {
		listenFormCallback.listen = new(models.Listen).Stage(listenFormCallback.probe.stageOfInterest)
	}
	listen_ := listenFormCallback.listen
	_ = listen_

	for _, formDiv := range listenFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(listen_.Name), formDiv)
		case "Assess":
			FormDivSelectFieldToField(&(listen_.Assess), listenFormCallback.probe.stageOfInterest, formDiv)
		case "Wait":
			FormDivSelectFieldToField(&(listen_.Wait), listenFormCallback.probe.stageOfInterest, formDiv)
		case "Other_listen":
			FormDivSelectFieldToField(&(listen_.Other_listen), listenFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if listenFormCallback.formGroup.HasSuppressButtonBeenPressed {
		listen_.Unstage(listenFormCallback.probe.stageOfInterest)
	}

	listenFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Listen](
		listenFormCallback.probe,
	)
	listenFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if listenFormCallback.CreationMode || listenFormCallback.formGroup.HasSuppressButtonBeenPressed {
		listenFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(listenFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ListenFormCallback(
			nil,
			listenFormCallback.probe,
			newFormGroup,
		)
		listen := new(models.Listen)
		FillUpForm(listen, newFormGroup, listenFormCallback.probe)
		listenFormCallback.probe.formStage.Commit()
	}

	fillUpTree(listenFormCallback.probe)
}
func __gong__New__ListeningFormCallback(
	listening *models.Listening,
	probe *Probe,
	formGroup *table.FormGroup,
) (listeningFormCallback *ListeningFormCallback) {
	listeningFormCallback = new(ListeningFormCallback)
	listeningFormCallback.probe = probe
	listeningFormCallback.listening = listening
	listeningFormCallback.formGroup = formGroup

	listeningFormCallback.CreationMode = (listening == nil)

	return
}

type ListeningFormCallback struct {
	listening *models.Listening

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (listeningFormCallback *ListeningFormCallback) OnSave() {

	log.Println("ListeningFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	listeningFormCallback.probe.formStage.Checkout()

	if listeningFormCallback.listening == nil {
		listeningFormCallback.listening = new(models.Listening).Stage(listeningFormCallback.probe.stageOfInterest)
	}
	listening_ := listeningFormCallback.listening
	_ = listening_

	for _, formDiv := range listeningFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(listening_.Name), formDiv)
		case "Offset":
			FormDivSelectFieldToField(&(listening_.Offset), listeningFormCallback.probe.stageOfInterest, formDiv)
		case "Sync":
			FormDivSelectFieldToField(&(listening_.Sync), listeningFormCallback.probe.stageOfInterest, formDiv)
		case "Other_listening":
			FormDivSelectFieldToField(&(listening_.Other_listening), listeningFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if listeningFormCallback.formGroup.HasSuppressButtonBeenPressed {
		listening_.Unstage(listeningFormCallback.probe.stageOfInterest)
	}

	listeningFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Listening](
		listeningFormCallback.probe,
	)
	listeningFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if listeningFormCallback.CreationMode || listeningFormCallback.formGroup.HasSuppressButtonBeenPressed {
		listeningFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(listeningFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ListeningFormCallback(
			nil,
			listeningFormCallback.probe,
			newFormGroup,
		)
		listening := new(models.Listening)
		FillUpForm(listening, newFormGroup, listeningFormCallback.probe)
		listeningFormCallback.probe.formStage.Commit()
	}

	fillUpTree(listeningFormCallback.probe)
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
		case "End_line":
			FormDivSelectFieldToField(&(lyric_.End_line), lyricFormCallback.probe.stageOfInterest, formDiv)
		case "End_paragraph":
			FormDivSelectFieldToField(&(lyric_.End_paragraph), lyricFormCallback.probe.stageOfInterest, formDiv)
		case "Extend":
			FormDivSelectFieldToField(&(lyric_.Extend), lyricFormCallback.probe.stageOfInterest, formDiv)
		case "Laughing":
			FormDivSelectFieldToField(&(lyric_.Laughing), lyricFormCallback.probe.stageOfInterest, formDiv)
		case "Humming":
			FormDivSelectFieldToField(&(lyric_.Humming), lyricFormCallback.probe.stageOfInterest, formDiv)
		case "Name":
			FormDivBasicFieldToField(&(lyric_.Name), formDiv)
		case "Note:Lyric":
			// we need to retrieve the field owner before the change
			var pastNoteOwner *models.Note
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Lyric"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				lyricFormCallback.probe.stageOfInterest,
				lyricFormCallback.probe.backRepoOfInterest,
				lyric_,
				&rf)

			if reverseFieldOwner != nil {
				pastNoteOwner = reverseFieldOwner.(*models.Note)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNoteOwner != nil {
					idx := slices.Index(pastNoteOwner.Lyric, lyric_)
					pastNoteOwner.Lyric = slices.Delete(pastNoteOwner.Lyric, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _note := range *models.GetGongstructInstancesSet[models.Note](lyricFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _note.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNoteOwner := _note // we have a match
						if pastNoteOwner != nil {
							if newNoteOwner != pastNoteOwner {
								idx := slices.Index(pastNoteOwner.Lyric, lyric_)
								pastNoteOwner.Lyric = slices.Delete(pastNoteOwner.Lyric, idx, idx+1)
								newNoteOwner.Lyric = append(newNoteOwner.Lyric, lyric_)
							}
						} else {
							newNoteOwner.Lyric = append(newNoteOwner.Lyric, lyric_)
						}
					}
				}
			}
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
		case "Defaults:Lyric_font":
			// we need to retrieve the field owner before the change
			var pastDefaultsOwner *models.Defaults
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Defaults"
			rf.Fieldname = "Lyric_font"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				lyric_fontFormCallback.probe.stageOfInterest,
				lyric_fontFormCallback.probe.backRepoOfInterest,
				lyric_font_,
				&rf)

			if reverseFieldOwner != nil {
				pastDefaultsOwner = reverseFieldOwner.(*models.Defaults)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDefaultsOwner != nil {
					idx := slices.Index(pastDefaultsOwner.Lyric_font, lyric_font_)
					pastDefaultsOwner.Lyric_font = slices.Delete(pastDefaultsOwner.Lyric_font, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _defaults := range *models.GetGongstructInstancesSet[models.Defaults](lyric_fontFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _defaults.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDefaultsOwner := _defaults // we have a match
						if pastDefaultsOwner != nil {
							if newDefaultsOwner != pastDefaultsOwner {
								idx := slices.Index(pastDefaultsOwner.Lyric_font, lyric_font_)
								pastDefaultsOwner.Lyric_font = slices.Delete(pastDefaultsOwner.Lyric_font, idx, idx+1)
								newDefaultsOwner.Lyric_font = append(newDefaultsOwner.Lyric_font, lyric_font_)
							}
						} else {
							newDefaultsOwner.Lyric_font = append(newDefaultsOwner.Lyric_font, lyric_font_)
						}
					}
				}
			}
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
		case "Defaults:Lyric_language":
			// we need to retrieve the field owner before the change
			var pastDefaultsOwner *models.Defaults
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Defaults"
			rf.Fieldname = "Lyric_language"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				lyric_languageFormCallback.probe.stageOfInterest,
				lyric_languageFormCallback.probe.backRepoOfInterest,
				lyric_language_,
				&rf)

			if reverseFieldOwner != nil {
				pastDefaultsOwner = reverseFieldOwner.(*models.Defaults)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDefaultsOwner != nil {
					idx := slices.Index(pastDefaultsOwner.Lyric_language, lyric_language_)
					pastDefaultsOwner.Lyric_language = slices.Delete(pastDefaultsOwner.Lyric_language, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _defaults := range *models.GetGongstructInstancesSet[models.Defaults](lyric_languageFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _defaults.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDefaultsOwner := _defaults // we have a match
						if pastDefaultsOwner != nil {
							if newDefaultsOwner != pastDefaultsOwner {
								idx := slices.Index(pastDefaultsOwner.Lyric_language, lyric_language_)
								pastDefaultsOwner.Lyric_language = slices.Delete(pastDefaultsOwner.Lyric_language, idx, idx+1)
								newDefaultsOwner.Lyric_language = append(newDefaultsOwner.Lyric_language, lyric_language_)
							}
						} else {
							newDefaultsOwner.Lyric_language = append(newDefaultsOwner.Lyric_language, lyric_language_)
						}
					}
				}
			}
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
func __gong__New__Measure_layoutFormCallback(
	measure_layout *models.Measure_layout,
	probe *Probe,
	formGroup *table.FormGroup,
) (measure_layoutFormCallback *Measure_layoutFormCallback) {
	measure_layoutFormCallback = new(Measure_layoutFormCallback)
	measure_layoutFormCallback.probe = probe
	measure_layoutFormCallback.measure_layout = measure_layout
	measure_layoutFormCallback.formGroup = formGroup

	measure_layoutFormCallback.CreationMode = (measure_layout == nil)

	return
}

type Measure_layoutFormCallback struct {
	measure_layout *models.Measure_layout

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (measure_layoutFormCallback *Measure_layoutFormCallback) OnSave() {

	log.Println("Measure_layoutFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	measure_layoutFormCallback.probe.formStage.Checkout()

	if measure_layoutFormCallback.measure_layout == nil {
		measure_layoutFormCallback.measure_layout = new(models.Measure_layout).Stage(measure_layoutFormCallback.probe.stageOfInterest)
	}
	measure_layout_ := measure_layoutFormCallback.measure_layout
	_ = measure_layout_

	for _, formDiv := range measure_layoutFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(measure_layout_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if measure_layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		measure_layout_.Unstage(measure_layoutFormCallback.probe.stageOfInterest)
	}

	measure_layoutFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Measure_layout](
		measure_layoutFormCallback.probe,
	)
	measure_layoutFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if measure_layoutFormCallback.CreationMode || measure_layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		measure_layoutFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(measure_layoutFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Measure_layoutFormCallback(
			nil,
			measure_layoutFormCallback.probe,
			newFormGroup,
		)
		measure_layout := new(models.Measure_layout)
		FillUpForm(measure_layout, newFormGroup, measure_layoutFormCallback.probe)
		measure_layoutFormCallback.probe.formStage.Commit()
	}

	fillUpTree(measure_layoutFormCallback.probe)
}
func __gong__New__Measure_numberingFormCallback(
	measure_numbering *models.Measure_numbering,
	probe *Probe,
	formGroup *table.FormGroup,
) (measure_numberingFormCallback *Measure_numberingFormCallback) {
	measure_numberingFormCallback = new(Measure_numberingFormCallback)
	measure_numberingFormCallback.probe = probe
	measure_numberingFormCallback.measure_numbering = measure_numbering
	measure_numberingFormCallback.formGroup = formGroup

	measure_numberingFormCallback.CreationMode = (measure_numbering == nil)

	return
}

type Measure_numberingFormCallback struct {
	measure_numbering *models.Measure_numbering

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (measure_numberingFormCallback *Measure_numberingFormCallback) OnSave() {

	log.Println("Measure_numberingFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	measure_numberingFormCallback.probe.formStage.Checkout()

	if measure_numberingFormCallback.measure_numbering == nil {
		measure_numberingFormCallback.measure_numbering = new(models.Measure_numbering).Stage(measure_numberingFormCallback.probe.stageOfInterest)
	}
	measure_numbering_ := measure_numberingFormCallback.measure_numbering
	_ = measure_numbering_

	for _, formDiv := range measure_numberingFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(measure_numbering_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if measure_numberingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		measure_numbering_.Unstage(measure_numberingFormCallback.probe.stageOfInterest)
	}

	measure_numberingFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Measure_numbering](
		measure_numberingFormCallback.probe,
	)
	measure_numberingFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if measure_numberingFormCallback.CreationMode || measure_numberingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		measure_numberingFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(measure_numberingFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Measure_numberingFormCallback(
			nil,
			measure_numberingFormCallback.probe,
			newFormGroup,
		)
		measure_numbering := new(models.Measure_numbering)
		FillUpForm(measure_numbering, newFormGroup, measure_numberingFormCallback.probe)
		measure_numberingFormCallback.probe.formStage.Commit()
	}

	fillUpTree(measure_numberingFormCallback.probe)
}
func __gong__New__Measure_repeatFormCallback(
	measure_repeat *models.Measure_repeat,
	probe *Probe,
	formGroup *table.FormGroup,
) (measure_repeatFormCallback *Measure_repeatFormCallback) {
	measure_repeatFormCallback = new(Measure_repeatFormCallback)
	measure_repeatFormCallback.probe = probe
	measure_repeatFormCallback.measure_repeat = measure_repeat
	measure_repeatFormCallback.formGroup = formGroup

	measure_repeatFormCallback.CreationMode = (measure_repeat == nil)

	return
}

type Measure_repeatFormCallback struct {
	measure_repeat *models.Measure_repeat

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (measure_repeatFormCallback *Measure_repeatFormCallback) OnSave() {

	log.Println("Measure_repeatFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	measure_repeatFormCallback.probe.formStage.Checkout()

	if measure_repeatFormCallback.measure_repeat == nil {
		measure_repeatFormCallback.measure_repeat = new(models.Measure_repeat).Stage(measure_repeatFormCallback.probe.stageOfInterest)
	}
	measure_repeat_ := measure_repeatFormCallback.measure_repeat
	_ = measure_repeat_

	for _, formDiv := range measure_repeatFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(measure_repeat_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if measure_repeatFormCallback.formGroup.HasSuppressButtonBeenPressed {
		measure_repeat_.Unstage(measure_repeatFormCallback.probe.stageOfInterest)
	}

	measure_repeatFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Measure_repeat](
		measure_repeatFormCallback.probe,
	)
	measure_repeatFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if measure_repeatFormCallback.CreationMode || measure_repeatFormCallback.formGroup.HasSuppressButtonBeenPressed {
		measure_repeatFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(measure_repeatFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Measure_repeatFormCallback(
			nil,
			measure_repeatFormCallback.probe,
			newFormGroup,
		)
		measure_repeat := new(models.Measure_repeat)
		FillUpForm(measure_repeat, newFormGroup, measure_repeatFormCallback.probe)
		measure_repeatFormCallback.probe.formStage.Commit()
	}

	fillUpTree(measure_repeatFormCallback.probe)
}
func __gong__New__Measure_styleFormCallback(
	measure_style *models.Measure_style,
	probe *Probe,
	formGroup *table.FormGroup,
) (measure_styleFormCallback *Measure_styleFormCallback) {
	measure_styleFormCallback = new(Measure_styleFormCallback)
	measure_styleFormCallback.probe = probe
	measure_styleFormCallback.measure_style = measure_style
	measure_styleFormCallback.formGroup = formGroup

	measure_styleFormCallback.CreationMode = (measure_style == nil)

	return
}

type Measure_styleFormCallback struct {
	measure_style *models.Measure_style

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (measure_styleFormCallback *Measure_styleFormCallback) OnSave() {

	log.Println("Measure_styleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	measure_styleFormCallback.probe.formStage.Checkout()

	if measure_styleFormCallback.measure_style == nil {
		measure_styleFormCallback.measure_style = new(models.Measure_style).Stage(measure_styleFormCallback.probe.stageOfInterest)
	}
	measure_style_ := measure_styleFormCallback.measure_style
	_ = measure_style_

	for _, formDiv := range measure_styleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(measure_style_.Name), formDiv)
		case "Multiple_rest":
			FormDivSelectFieldToField(&(measure_style_.Multiple_rest), measure_styleFormCallback.probe.stageOfInterest, formDiv)
		case "Measure_repeat":
			FormDivSelectFieldToField(&(measure_style_.Measure_repeat), measure_styleFormCallback.probe.stageOfInterest, formDiv)
		case "Beat_repeat":
			FormDivSelectFieldToField(&(measure_style_.Beat_repeat), measure_styleFormCallback.probe.stageOfInterest, formDiv)
		case "Slash":
			FormDivSelectFieldToField(&(measure_style_.Slash), measure_styleFormCallback.probe.stageOfInterest, formDiv)
		case "Attributes:Measure_style":
			// we need to retrieve the field owner before the change
			var pastAttributesOwner *models.Attributes
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Measure_style"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				measure_styleFormCallback.probe.stageOfInterest,
				measure_styleFormCallback.probe.backRepoOfInterest,
				measure_style_,
				&rf)

			if reverseFieldOwner != nil {
				pastAttributesOwner = reverseFieldOwner.(*models.Attributes)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAttributesOwner != nil {
					idx := slices.Index(pastAttributesOwner.Measure_style, measure_style_)
					pastAttributesOwner.Measure_style = slices.Delete(pastAttributesOwner.Measure_style, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attributes := range *models.GetGongstructInstancesSet[models.Attributes](measure_styleFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAttributesOwner := _attributes // we have a match
						if pastAttributesOwner != nil {
							if newAttributesOwner != pastAttributesOwner {
								idx := slices.Index(pastAttributesOwner.Measure_style, measure_style_)
								pastAttributesOwner.Measure_style = slices.Delete(pastAttributesOwner.Measure_style, idx, idx+1)
								newAttributesOwner.Measure_style = append(newAttributesOwner.Measure_style, measure_style_)
							}
						} else {
							newAttributesOwner.Measure_style = append(newAttributesOwner.Measure_style, measure_style_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if measure_styleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		measure_style_.Unstage(measure_styleFormCallback.probe.stageOfInterest)
	}

	measure_styleFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Measure_style](
		measure_styleFormCallback.probe,
	)
	measure_styleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if measure_styleFormCallback.CreationMode || measure_styleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		measure_styleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(measure_styleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Measure_styleFormCallback(
			nil,
			measure_styleFormCallback.probe,
			newFormGroup,
		)
		measure_style := new(models.Measure_style)
		FillUpForm(measure_style, newFormGroup, measure_styleFormCallback.probe)
		measure_styleFormCallback.probe.formStage.Commit()
	}

	fillUpTree(measure_styleFormCallback.probe)
}
func __gong__New__MembraneFormCallback(
	membrane *models.Membrane,
	probe *Probe,
	formGroup *table.FormGroup,
) (membraneFormCallback *MembraneFormCallback) {
	membraneFormCallback = new(MembraneFormCallback)
	membraneFormCallback.probe = probe
	membraneFormCallback.membrane = membrane
	membraneFormCallback.formGroup = formGroup

	membraneFormCallback.CreationMode = (membrane == nil)

	return
}

type MembraneFormCallback struct {
	membrane *models.Membrane

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (membraneFormCallback *MembraneFormCallback) OnSave() {

	log.Println("MembraneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	membraneFormCallback.probe.formStage.Checkout()

	if membraneFormCallback.membrane == nil {
		membraneFormCallback.membrane = new(models.Membrane).Stage(membraneFormCallback.probe.stageOfInterest)
	}
	membrane_ := membraneFormCallback.membrane
	_ = membrane_

	for _, formDiv := range membraneFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(membrane_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if membraneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		membrane_.Unstage(membraneFormCallback.probe.stageOfInterest)
	}

	membraneFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Membrane](
		membraneFormCallback.probe,
	)
	membraneFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if membraneFormCallback.CreationMode || membraneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		membraneFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(membraneFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MembraneFormCallback(
			nil,
			membraneFormCallback.probe,
			newFormGroup,
		)
		membrane := new(models.Membrane)
		FillUpForm(membrane, newFormGroup, membraneFormCallback.probe)
		membraneFormCallback.probe.formStage.Commit()
	}

	fillUpTree(membraneFormCallback.probe)
}
func __gong__New__MetalFormCallback(
	metal *models.Metal,
	probe *Probe,
	formGroup *table.FormGroup,
) (metalFormCallback *MetalFormCallback) {
	metalFormCallback = new(MetalFormCallback)
	metalFormCallback.probe = probe
	metalFormCallback.metal = metal
	metalFormCallback.formGroup = formGroup

	metalFormCallback.CreationMode = (metal == nil)

	return
}

type MetalFormCallback struct {
	metal *models.Metal

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (metalFormCallback *MetalFormCallback) OnSave() {

	log.Println("MetalFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metalFormCallback.probe.formStage.Checkout()

	if metalFormCallback.metal == nil {
		metalFormCallback.metal = new(models.Metal).Stage(metalFormCallback.probe.stageOfInterest)
	}
	metal_ := metalFormCallback.metal
	_ = metal_

	for _, formDiv := range metalFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metal_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if metalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metal_.Unstage(metalFormCallback.probe.stageOfInterest)
	}

	metalFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Metal](
		metalFormCallback.probe,
	)
	metalFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if metalFormCallback.CreationMode || metalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metalFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(metalFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MetalFormCallback(
			nil,
			metalFormCallback.probe,
			newFormGroup,
		)
		metal := new(models.Metal)
		FillUpForm(metal, newFormGroup, metalFormCallback.probe)
		metalFormCallback.probe.formStage.Commit()
	}

	fillUpTree(metalFormCallback.probe)
}
func __gong__New__MetronomeFormCallback(
	metronome *models.Metronome,
	probe *Probe,
	formGroup *table.FormGroup,
) (metronomeFormCallback *MetronomeFormCallback) {
	metronomeFormCallback = new(MetronomeFormCallback)
	metronomeFormCallback.probe = probe
	metronomeFormCallback.metronome = metronome
	metronomeFormCallback.formGroup = formGroup

	metronomeFormCallback.CreationMode = (metronome == nil)

	return
}

type MetronomeFormCallback struct {
	metronome *models.Metronome

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (metronomeFormCallback *MetronomeFormCallback) OnSave() {

	log.Println("MetronomeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metronomeFormCallback.probe.formStage.Checkout()

	if metronomeFormCallback.metronome == nil {
		metronomeFormCallback.metronome = new(models.Metronome).Stage(metronomeFormCallback.probe.stageOfInterest)
	}
	metronome_ := metronomeFormCallback.metronome
	_ = metronome_

	for _, formDiv := range metronomeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metronome_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if metronomeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronome_.Unstage(metronomeFormCallback.probe.stageOfInterest)
	}

	metronomeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Metronome](
		metronomeFormCallback.probe,
	)
	metronomeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if metronomeFormCallback.CreationMode || metronomeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronomeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(metronomeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MetronomeFormCallback(
			nil,
			metronomeFormCallback.probe,
			newFormGroup,
		)
		metronome := new(models.Metronome)
		FillUpForm(metronome, newFormGroup, metronomeFormCallback.probe)
		metronomeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(metronomeFormCallback.probe)
}
func __gong__New__Metronome_beamFormCallback(
	metronome_beam *models.Metronome_beam,
	probe *Probe,
	formGroup *table.FormGroup,
) (metronome_beamFormCallback *Metronome_beamFormCallback) {
	metronome_beamFormCallback = new(Metronome_beamFormCallback)
	metronome_beamFormCallback.probe = probe
	metronome_beamFormCallback.metronome_beam = metronome_beam
	metronome_beamFormCallback.formGroup = formGroup

	metronome_beamFormCallback.CreationMode = (metronome_beam == nil)

	return
}

type Metronome_beamFormCallback struct {
	metronome_beam *models.Metronome_beam

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (metronome_beamFormCallback *Metronome_beamFormCallback) OnSave() {

	log.Println("Metronome_beamFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metronome_beamFormCallback.probe.formStage.Checkout()

	if metronome_beamFormCallback.metronome_beam == nil {
		metronome_beamFormCallback.metronome_beam = new(models.Metronome_beam).Stage(metronome_beamFormCallback.probe.stageOfInterest)
	}
	metronome_beam_ := metronome_beamFormCallback.metronome_beam
	_ = metronome_beam_

	for _, formDiv := range metronome_beamFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metronome_beam_.Name), formDiv)
		case "Metronome_note:Metronome_beam":
			// we need to retrieve the field owner before the change
			var pastMetronome_noteOwner *models.Metronome_note
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Metronome_note"
			rf.Fieldname = "Metronome_beam"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				metronome_beamFormCallback.probe.stageOfInterest,
				metronome_beamFormCallback.probe.backRepoOfInterest,
				metronome_beam_,
				&rf)

			if reverseFieldOwner != nil {
				pastMetronome_noteOwner = reverseFieldOwner.(*models.Metronome_note)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastMetronome_noteOwner != nil {
					idx := slices.Index(pastMetronome_noteOwner.Metronome_beam, metronome_beam_)
					pastMetronome_noteOwner.Metronome_beam = slices.Delete(pastMetronome_noteOwner.Metronome_beam, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _metronome_note := range *models.GetGongstructInstancesSet[models.Metronome_note](metronome_beamFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _metronome_note.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newMetronome_noteOwner := _metronome_note // we have a match
						if pastMetronome_noteOwner != nil {
							if newMetronome_noteOwner != pastMetronome_noteOwner {
								idx := slices.Index(pastMetronome_noteOwner.Metronome_beam, metronome_beam_)
								pastMetronome_noteOwner.Metronome_beam = slices.Delete(pastMetronome_noteOwner.Metronome_beam, idx, idx+1)
								newMetronome_noteOwner.Metronome_beam = append(newMetronome_noteOwner.Metronome_beam, metronome_beam_)
							}
						} else {
							newMetronome_noteOwner.Metronome_beam = append(newMetronome_noteOwner.Metronome_beam, metronome_beam_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if metronome_beamFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronome_beam_.Unstage(metronome_beamFormCallback.probe.stageOfInterest)
	}

	metronome_beamFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Metronome_beam](
		metronome_beamFormCallback.probe,
	)
	metronome_beamFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if metronome_beamFormCallback.CreationMode || metronome_beamFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronome_beamFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(metronome_beamFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Metronome_beamFormCallback(
			nil,
			metronome_beamFormCallback.probe,
			newFormGroup,
		)
		metronome_beam := new(models.Metronome_beam)
		FillUpForm(metronome_beam, newFormGroup, metronome_beamFormCallback.probe)
		metronome_beamFormCallback.probe.formStage.Commit()
	}

	fillUpTree(metronome_beamFormCallback.probe)
}
func __gong__New__Metronome_noteFormCallback(
	metronome_note *models.Metronome_note,
	probe *Probe,
	formGroup *table.FormGroup,
) (metronome_noteFormCallback *Metronome_noteFormCallback) {
	metronome_noteFormCallback = new(Metronome_noteFormCallback)
	metronome_noteFormCallback.probe = probe
	metronome_noteFormCallback.metronome_note = metronome_note
	metronome_noteFormCallback.formGroup = formGroup

	metronome_noteFormCallback.CreationMode = (metronome_note == nil)

	return
}

type Metronome_noteFormCallback struct {
	metronome_note *models.Metronome_note

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (metronome_noteFormCallback *Metronome_noteFormCallback) OnSave() {

	log.Println("Metronome_noteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metronome_noteFormCallback.probe.formStage.Checkout()

	if metronome_noteFormCallback.metronome_note == nil {
		metronome_noteFormCallback.metronome_note = new(models.Metronome_note).Stage(metronome_noteFormCallback.probe.stageOfInterest)
	}
	metronome_note_ := metronome_noteFormCallback.metronome_note
	_ = metronome_note_

	for _, formDiv := range metronome_noteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metronome_note_.Name), formDiv)
		case "Metronome_tied":
			FormDivSelectFieldToField(&(metronome_note_.Metronome_tied), metronome_noteFormCallback.probe.stageOfInterest, formDiv)
		case "Metronome_tuplet":
			FormDivSelectFieldToField(&(metronome_note_.Metronome_tuplet), metronome_noteFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if metronome_noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronome_note_.Unstage(metronome_noteFormCallback.probe.stageOfInterest)
	}

	metronome_noteFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Metronome_note](
		metronome_noteFormCallback.probe,
	)
	metronome_noteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if metronome_noteFormCallback.CreationMode || metronome_noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronome_noteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(metronome_noteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Metronome_noteFormCallback(
			nil,
			metronome_noteFormCallback.probe,
			newFormGroup,
		)
		metronome_note := new(models.Metronome_note)
		FillUpForm(metronome_note, newFormGroup, metronome_noteFormCallback.probe)
		metronome_noteFormCallback.probe.formStage.Commit()
	}

	fillUpTree(metronome_noteFormCallback.probe)
}
func __gong__New__Metronome_tiedFormCallback(
	metronome_tied *models.Metronome_tied,
	probe *Probe,
	formGroup *table.FormGroup,
) (metronome_tiedFormCallback *Metronome_tiedFormCallback) {
	metronome_tiedFormCallback = new(Metronome_tiedFormCallback)
	metronome_tiedFormCallback.probe = probe
	metronome_tiedFormCallback.metronome_tied = metronome_tied
	metronome_tiedFormCallback.formGroup = formGroup

	metronome_tiedFormCallback.CreationMode = (metronome_tied == nil)

	return
}

type Metronome_tiedFormCallback struct {
	metronome_tied *models.Metronome_tied

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (metronome_tiedFormCallback *Metronome_tiedFormCallback) OnSave() {

	log.Println("Metronome_tiedFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metronome_tiedFormCallback.probe.formStage.Checkout()

	if metronome_tiedFormCallback.metronome_tied == nil {
		metronome_tiedFormCallback.metronome_tied = new(models.Metronome_tied).Stage(metronome_tiedFormCallback.probe.stageOfInterest)
	}
	metronome_tied_ := metronome_tiedFormCallback.metronome_tied
	_ = metronome_tied_

	for _, formDiv := range metronome_tiedFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metronome_tied_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if metronome_tiedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronome_tied_.Unstage(metronome_tiedFormCallback.probe.stageOfInterest)
	}

	metronome_tiedFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Metronome_tied](
		metronome_tiedFormCallback.probe,
	)
	metronome_tiedFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if metronome_tiedFormCallback.CreationMode || metronome_tiedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronome_tiedFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(metronome_tiedFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Metronome_tiedFormCallback(
			nil,
			metronome_tiedFormCallback.probe,
			newFormGroup,
		)
		metronome_tied := new(models.Metronome_tied)
		FillUpForm(metronome_tied, newFormGroup, metronome_tiedFormCallback.probe)
		metronome_tiedFormCallback.probe.formStage.Commit()
	}

	fillUpTree(metronome_tiedFormCallback.probe)
}
func __gong__New__Metronome_tupletFormCallback(
	metronome_tuplet *models.Metronome_tuplet,
	probe *Probe,
	formGroup *table.FormGroup,
) (metronome_tupletFormCallback *Metronome_tupletFormCallback) {
	metronome_tupletFormCallback = new(Metronome_tupletFormCallback)
	metronome_tupletFormCallback.probe = probe
	metronome_tupletFormCallback.metronome_tuplet = metronome_tuplet
	metronome_tupletFormCallback.formGroup = formGroup

	metronome_tupletFormCallback.CreationMode = (metronome_tuplet == nil)

	return
}

type Metronome_tupletFormCallback struct {
	metronome_tuplet *models.Metronome_tuplet

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (metronome_tupletFormCallback *Metronome_tupletFormCallback) OnSave() {

	log.Println("Metronome_tupletFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metronome_tupletFormCallback.probe.formStage.Checkout()

	if metronome_tupletFormCallback.metronome_tuplet == nil {
		metronome_tupletFormCallback.metronome_tuplet = new(models.Metronome_tuplet).Stage(metronome_tupletFormCallback.probe.stageOfInterest)
	}
	metronome_tuplet_ := metronome_tupletFormCallback.metronome_tuplet
	_ = metronome_tuplet_

	for _, formDiv := range metronome_tupletFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metronome_tuplet_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if metronome_tupletFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronome_tuplet_.Unstage(metronome_tupletFormCallback.probe.stageOfInterest)
	}

	metronome_tupletFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Metronome_tuplet](
		metronome_tupletFormCallback.probe,
	)
	metronome_tupletFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if metronome_tupletFormCallback.CreationMode || metronome_tupletFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metronome_tupletFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(metronome_tupletFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Metronome_tupletFormCallback(
			nil,
			metronome_tupletFormCallback.probe,
			newFormGroup,
		)
		metronome_tuplet := new(models.Metronome_tuplet)
		FillUpForm(metronome_tuplet, newFormGroup, metronome_tupletFormCallback.probe)
		metronome_tupletFormCallback.probe.formStage.Commit()
	}

	fillUpTree(metronome_tupletFormCallback.probe)
}
func __gong__New__Midi_deviceFormCallback(
	midi_device *models.Midi_device,
	probe *Probe,
	formGroup *table.FormGroup,
) (midi_deviceFormCallback *Midi_deviceFormCallback) {
	midi_deviceFormCallback = new(Midi_deviceFormCallback)
	midi_deviceFormCallback.probe = probe
	midi_deviceFormCallback.midi_device = midi_device
	midi_deviceFormCallback.formGroup = formGroup

	midi_deviceFormCallback.CreationMode = (midi_device == nil)

	return
}

type Midi_deviceFormCallback struct {
	midi_device *models.Midi_device

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (midi_deviceFormCallback *Midi_deviceFormCallback) OnSave() {

	log.Println("Midi_deviceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	midi_deviceFormCallback.probe.formStage.Checkout()

	if midi_deviceFormCallback.midi_device == nil {
		midi_deviceFormCallback.midi_device = new(models.Midi_device).Stage(midi_deviceFormCallback.probe.stageOfInterest)
	}
	midi_device_ := midi_deviceFormCallback.midi_device
	_ = midi_device_

	for _, formDiv := range midi_deviceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(midi_device_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(midi_device_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if midi_deviceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		midi_device_.Unstage(midi_deviceFormCallback.probe.stageOfInterest)
	}

	midi_deviceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Midi_device](
		midi_deviceFormCallback.probe,
	)
	midi_deviceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if midi_deviceFormCallback.CreationMode || midi_deviceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		midi_deviceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(midi_deviceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Midi_deviceFormCallback(
			nil,
			midi_deviceFormCallback.probe,
			newFormGroup,
		)
		midi_device := new(models.Midi_device)
		FillUpForm(midi_device, newFormGroup, midi_deviceFormCallback.probe)
		midi_deviceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(midi_deviceFormCallback.probe)
}
func __gong__New__Midi_instrumentFormCallback(
	midi_instrument *models.Midi_instrument,
	probe *Probe,
	formGroup *table.FormGroup,
) (midi_instrumentFormCallback *Midi_instrumentFormCallback) {
	midi_instrumentFormCallback = new(Midi_instrumentFormCallback)
	midi_instrumentFormCallback.probe = probe
	midi_instrumentFormCallback.midi_instrument = midi_instrument
	midi_instrumentFormCallback.formGroup = formGroup

	midi_instrumentFormCallback.CreationMode = (midi_instrument == nil)

	return
}

type Midi_instrumentFormCallback struct {
	midi_instrument *models.Midi_instrument

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (midi_instrumentFormCallback *Midi_instrumentFormCallback) OnSave() {

	log.Println("Midi_instrumentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	midi_instrumentFormCallback.probe.formStage.Checkout()

	if midi_instrumentFormCallback.midi_instrument == nil {
		midi_instrumentFormCallback.midi_instrument = new(models.Midi_instrument).Stage(midi_instrumentFormCallback.probe.stageOfInterest)
	}
	midi_instrument_ := midi_instrumentFormCallback.midi_instrument
	_ = midi_instrument_

	for _, formDiv := range midi_instrumentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(midi_instrument_.Name), formDiv)
		case "Midi_name":
			FormDivBasicFieldToField(&(midi_instrument_.Midi_name), formDiv)
		}
	}

	// manage the suppress operation
	if midi_instrumentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		midi_instrument_.Unstage(midi_instrumentFormCallback.probe.stageOfInterest)
	}

	midi_instrumentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Midi_instrument](
		midi_instrumentFormCallback.probe,
	)
	midi_instrumentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if midi_instrumentFormCallback.CreationMode || midi_instrumentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		midi_instrumentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(midi_instrumentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Midi_instrumentFormCallback(
			nil,
			midi_instrumentFormCallback.probe,
			newFormGroup,
		)
		midi_instrument := new(models.Midi_instrument)
		FillUpForm(midi_instrument, newFormGroup, midi_instrumentFormCallback.probe)
		midi_instrumentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(midi_instrumentFormCallback.probe)
}
func __gong__New__MiscellaneousFormCallback(
	miscellaneous *models.Miscellaneous,
	probe *Probe,
	formGroup *table.FormGroup,
) (miscellaneousFormCallback *MiscellaneousFormCallback) {
	miscellaneousFormCallback = new(MiscellaneousFormCallback)
	miscellaneousFormCallback.probe = probe
	miscellaneousFormCallback.miscellaneous = miscellaneous
	miscellaneousFormCallback.formGroup = formGroup

	miscellaneousFormCallback.CreationMode = (miscellaneous == nil)

	return
}

type MiscellaneousFormCallback struct {
	miscellaneous *models.Miscellaneous

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (miscellaneousFormCallback *MiscellaneousFormCallback) OnSave() {

	log.Println("MiscellaneousFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	miscellaneousFormCallback.probe.formStage.Checkout()

	if miscellaneousFormCallback.miscellaneous == nil {
		miscellaneousFormCallback.miscellaneous = new(models.Miscellaneous).Stage(miscellaneousFormCallback.probe.stageOfInterest)
	}
	miscellaneous_ := miscellaneousFormCallback.miscellaneous
	_ = miscellaneous_

	for _, formDiv := range miscellaneousFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(miscellaneous_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if miscellaneousFormCallback.formGroup.HasSuppressButtonBeenPressed {
		miscellaneous_.Unstage(miscellaneousFormCallback.probe.stageOfInterest)
	}

	miscellaneousFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Miscellaneous](
		miscellaneousFormCallback.probe,
	)
	miscellaneousFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if miscellaneousFormCallback.CreationMode || miscellaneousFormCallback.formGroup.HasSuppressButtonBeenPressed {
		miscellaneousFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(miscellaneousFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MiscellaneousFormCallback(
			nil,
			miscellaneousFormCallback.probe,
			newFormGroup,
		)
		miscellaneous := new(models.Miscellaneous)
		FillUpForm(miscellaneous, newFormGroup, miscellaneousFormCallback.probe)
		miscellaneousFormCallback.probe.formStage.Commit()
	}

	fillUpTree(miscellaneousFormCallback.probe)
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
		case "Miscellaneous:Miscellaneous_field":
			// we need to retrieve the field owner before the change
			var pastMiscellaneousOwner *models.Miscellaneous
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Miscellaneous"
			rf.Fieldname = "Miscellaneous_field"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				miscellaneous_fieldFormCallback.probe.stageOfInterest,
				miscellaneous_fieldFormCallback.probe.backRepoOfInterest,
				miscellaneous_field_,
				&rf)

			if reverseFieldOwner != nil {
				pastMiscellaneousOwner = reverseFieldOwner.(*models.Miscellaneous)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastMiscellaneousOwner != nil {
					idx := slices.Index(pastMiscellaneousOwner.Miscellaneous_field, miscellaneous_field_)
					pastMiscellaneousOwner.Miscellaneous_field = slices.Delete(pastMiscellaneousOwner.Miscellaneous_field, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _miscellaneous := range *models.GetGongstructInstancesSet[models.Miscellaneous](miscellaneous_fieldFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _miscellaneous.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newMiscellaneousOwner := _miscellaneous // we have a match
						if pastMiscellaneousOwner != nil {
							if newMiscellaneousOwner != pastMiscellaneousOwner {
								idx := slices.Index(pastMiscellaneousOwner.Miscellaneous_field, miscellaneous_field_)
								pastMiscellaneousOwner.Miscellaneous_field = slices.Delete(pastMiscellaneousOwner.Miscellaneous_field, idx, idx+1)
								newMiscellaneousOwner.Miscellaneous_field = append(newMiscellaneousOwner.Miscellaneous_field, miscellaneous_field_)
							}
						} else {
							newMiscellaneousOwner.Miscellaneous_field = append(newMiscellaneousOwner.Miscellaneous_field, miscellaneous_field_)
						}
					}
				}
			}
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
func __gong__New__MordentFormCallback(
	mordent *models.Mordent,
	probe *Probe,
	formGroup *table.FormGroup,
) (mordentFormCallback *MordentFormCallback) {
	mordentFormCallback = new(MordentFormCallback)
	mordentFormCallback.probe = probe
	mordentFormCallback.mordent = mordent
	mordentFormCallback.formGroup = formGroup

	mordentFormCallback.CreationMode = (mordent == nil)

	return
}

type MordentFormCallback struct {
	mordent *models.Mordent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (mordentFormCallback *MordentFormCallback) OnSave() {

	log.Println("MordentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mordentFormCallback.probe.formStage.Checkout()

	if mordentFormCallback.mordent == nil {
		mordentFormCallback.mordent = new(models.Mordent).Stage(mordentFormCallback.probe.stageOfInterest)
	}
	mordent_ := mordentFormCallback.mordent
	_ = mordent_

	for _, formDiv := range mordentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(mordent_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if mordentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mordent_.Unstage(mordentFormCallback.probe.stageOfInterest)
	}

	mordentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Mordent](
		mordentFormCallback.probe,
	)
	mordentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if mordentFormCallback.CreationMode || mordentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mordentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(mordentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MordentFormCallback(
			nil,
			mordentFormCallback.probe,
			newFormGroup,
		)
		mordent := new(models.Mordent)
		FillUpForm(mordent, newFormGroup, mordentFormCallback.probe)
		mordentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(mordentFormCallback.probe)
}
func __gong__New__Multiple_restFormCallback(
	multiple_rest *models.Multiple_rest,
	probe *Probe,
	formGroup *table.FormGroup,
) (multiple_restFormCallback *Multiple_restFormCallback) {
	multiple_restFormCallback = new(Multiple_restFormCallback)
	multiple_restFormCallback.probe = probe
	multiple_restFormCallback.multiple_rest = multiple_rest
	multiple_restFormCallback.formGroup = formGroup

	multiple_restFormCallback.CreationMode = (multiple_rest == nil)

	return
}

type Multiple_restFormCallback struct {
	multiple_rest *models.Multiple_rest

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (multiple_restFormCallback *Multiple_restFormCallback) OnSave() {

	log.Println("Multiple_restFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	multiple_restFormCallback.probe.formStage.Checkout()

	if multiple_restFormCallback.multiple_rest == nil {
		multiple_restFormCallback.multiple_rest = new(models.Multiple_rest).Stage(multiple_restFormCallback.probe.stageOfInterest)
	}
	multiple_rest_ := multiple_restFormCallback.multiple_rest
	_ = multiple_rest_

	for _, formDiv := range multiple_restFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(multiple_rest_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if multiple_restFormCallback.formGroup.HasSuppressButtonBeenPressed {
		multiple_rest_.Unstage(multiple_restFormCallback.probe.stageOfInterest)
	}

	multiple_restFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Multiple_rest](
		multiple_restFormCallback.probe,
	)
	multiple_restFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if multiple_restFormCallback.CreationMode || multiple_restFormCallback.formGroup.HasSuppressButtonBeenPressed {
		multiple_restFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(multiple_restFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Multiple_restFormCallback(
			nil,
			multiple_restFormCallback.probe,
			newFormGroup,
		)
		multiple_rest := new(models.Multiple_rest)
		FillUpForm(multiple_rest, newFormGroup, multiple_restFormCallback.probe)
		multiple_restFormCallback.probe.formStage.Commit()
	}

	fillUpTree(multiple_restFormCallback.probe)
}
func __gong__New__Name_displayFormCallback(
	name_display *models.Name_display,
	probe *Probe,
	formGroup *table.FormGroup,
) (name_displayFormCallback *Name_displayFormCallback) {
	name_displayFormCallback = new(Name_displayFormCallback)
	name_displayFormCallback.probe = probe
	name_displayFormCallback.name_display = name_display
	name_displayFormCallback.formGroup = formGroup

	name_displayFormCallback.CreationMode = (name_display == nil)

	return
}

type Name_displayFormCallback struct {
	name_display *models.Name_display

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (name_displayFormCallback *Name_displayFormCallback) OnSave() {

	log.Println("Name_displayFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	name_displayFormCallback.probe.formStage.Checkout()

	if name_displayFormCallback.name_display == nil {
		name_displayFormCallback.name_display = new(models.Name_display).Stage(name_displayFormCallback.probe.stageOfInterest)
	}
	name_display_ := name_displayFormCallback.name_display
	_ = name_display_

	for _, formDiv := range name_displayFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(name_display_.Name), formDiv)
		case "Accidental_text":
			FormDivSelectFieldToField(&(name_display_.Accidental_text), name_displayFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if name_displayFormCallback.formGroup.HasSuppressButtonBeenPressed {
		name_display_.Unstage(name_displayFormCallback.probe.stageOfInterest)
	}

	name_displayFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Name_display](
		name_displayFormCallback.probe,
	)
	name_displayFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if name_displayFormCallback.CreationMode || name_displayFormCallback.formGroup.HasSuppressButtonBeenPressed {
		name_displayFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(name_displayFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Name_displayFormCallback(
			nil,
			name_displayFormCallback.probe,
			newFormGroup,
		)
		name_display := new(models.Name_display)
		FillUpForm(name_display, newFormGroup, name_displayFormCallback.probe)
		name_displayFormCallback.probe.formStage.Commit()
	}

	fillUpTree(name_displayFormCallback.probe)
}
func __gong__New__Non_arpeggiateFormCallback(
	non_arpeggiate *models.Non_arpeggiate,
	probe *Probe,
	formGroup *table.FormGroup,
) (non_arpeggiateFormCallback *Non_arpeggiateFormCallback) {
	non_arpeggiateFormCallback = new(Non_arpeggiateFormCallback)
	non_arpeggiateFormCallback.probe = probe
	non_arpeggiateFormCallback.non_arpeggiate = non_arpeggiate
	non_arpeggiateFormCallback.formGroup = formGroup

	non_arpeggiateFormCallback.CreationMode = (non_arpeggiate == nil)

	return
}

type Non_arpeggiateFormCallback struct {
	non_arpeggiate *models.Non_arpeggiate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (non_arpeggiateFormCallback *Non_arpeggiateFormCallback) OnSave() {

	log.Println("Non_arpeggiateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	non_arpeggiateFormCallback.probe.formStage.Checkout()

	if non_arpeggiateFormCallback.non_arpeggiate == nil {
		non_arpeggiateFormCallback.non_arpeggiate = new(models.Non_arpeggiate).Stage(non_arpeggiateFormCallback.probe.stageOfInterest)
	}
	non_arpeggiate_ := non_arpeggiateFormCallback.non_arpeggiate
	_ = non_arpeggiate_

	for _, formDiv := range non_arpeggiateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(non_arpeggiate_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if non_arpeggiateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		non_arpeggiate_.Unstage(non_arpeggiateFormCallback.probe.stageOfInterest)
	}

	non_arpeggiateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Non_arpeggiate](
		non_arpeggiateFormCallback.probe,
	)
	non_arpeggiateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if non_arpeggiateFormCallback.CreationMode || non_arpeggiateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		non_arpeggiateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(non_arpeggiateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Non_arpeggiateFormCallback(
			nil,
			non_arpeggiateFormCallback.probe,
			newFormGroup,
		)
		non_arpeggiate := new(models.Non_arpeggiate)
		FillUpForm(non_arpeggiate, newFormGroup, non_arpeggiateFormCallback.probe)
		non_arpeggiateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(non_arpeggiateFormCallback.probe)
}
func __gong__New__NotationsFormCallback(
	notations *models.Notations,
	probe *Probe,
	formGroup *table.FormGroup,
) (notationsFormCallback *NotationsFormCallback) {
	notationsFormCallback = new(NotationsFormCallback)
	notationsFormCallback.probe = probe
	notationsFormCallback.notations = notations
	notationsFormCallback.formGroup = formGroup

	notationsFormCallback.CreationMode = (notations == nil)

	return
}

type NotationsFormCallback struct {
	notations *models.Notations

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (notationsFormCallback *NotationsFormCallback) OnSave() {

	log.Println("NotationsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notationsFormCallback.probe.formStage.Checkout()

	if notationsFormCallback.notations == nil {
		notationsFormCallback.notations = new(models.Notations).Stage(notationsFormCallback.probe.stageOfInterest)
	}
	notations_ := notationsFormCallback.notations
	_ = notations_

	for _, formDiv := range notationsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notations_.Name), formDiv)
		case "Tied":
			FormDivSelectFieldToField(&(notations_.Tied), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Slur":
			FormDivSelectFieldToField(&(notations_.Slur), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Tuplet":
			FormDivSelectFieldToField(&(notations_.Tuplet), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Glissando":
			FormDivSelectFieldToField(&(notations_.Glissando), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Slide":
			FormDivSelectFieldToField(&(notations_.Slide), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Ornaments":
			FormDivSelectFieldToField(&(notations_.Ornaments), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Technical":
			FormDivSelectFieldToField(&(notations_.Technical), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Articulations":
			FormDivSelectFieldToField(&(notations_.Articulations), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Dynamics":
			FormDivSelectFieldToField(&(notations_.Dynamics), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Fermata":
			FormDivSelectFieldToField(&(notations_.Fermata), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Arpeggiate":
			FormDivSelectFieldToField(&(notations_.Arpeggiate), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Non_arpeggiate":
			FormDivSelectFieldToField(&(notations_.Non_arpeggiate), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Accidental_mark":
			FormDivSelectFieldToField(&(notations_.Accidental_mark), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Other_notation":
			FormDivSelectFieldToField(&(notations_.Other_notation), notationsFormCallback.probe.stageOfInterest, formDiv)
		case "Note:Notations":
			// we need to retrieve the field owner before the change
			var pastNoteOwner *models.Note
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Notations"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				notationsFormCallback.probe.stageOfInterest,
				notationsFormCallback.probe.backRepoOfInterest,
				notations_,
				&rf)

			if reverseFieldOwner != nil {
				pastNoteOwner = reverseFieldOwner.(*models.Note)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNoteOwner != nil {
					idx := slices.Index(pastNoteOwner.Notations, notations_)
					pastNoteOwner.Notations = slices.Delete(pastNoteOwner.Notations, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _note := range *models.GetGongstructInstancesSet[models.Note](notationsFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _note.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNoteOwner := _note // we have a match
						if pastNoteOwner != nil {
							if newNoteOwner != pastNoteOwner {
								idx := slices.Index(pastNoteOwner.Notations, notations_)
								pastNoteOwner.Notations = slices.Delete(pastNoteOwner.Notations, idx, idx+1)
								newNoteOwner.Notations = append(newNoteOwner.Notations, notations_)
							}
						} else {
							newNoteOwner.Notations = append(newNoteOwner.Notations, notations_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notations_.Unstage(notationsFormCallback.probe.stageOfInterest)
	}

	notationsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Notations](
		notationsFormCallback.probe,
	)
	notationsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if notationsFormCallback.CreationMode || notationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notationsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(notationsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NotationsFormCallback(
			nil,
			notationsFormCallback.probe,
			newFormGroup,
		)
		notations := new(models.Notations)
		FillUpForm(notations, newFormGroup, notationsFormCallback.probe)
		notationsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(notationsFormCallback.probe)
}
func __gong__New__NoteFormCallback(
	note *models.Note,
	probe *Probe,
	formGroup *table.FormGroup,
) (noteFormCallback *NoteFormCallback) {
	noteFormCallback = new(NoteFormCallback)
	noteFormCallback.probe = probe
	noteFormCallback.note = note
	noteFormCallback.formGroup = formGroup

	noteFormCallback.CreationMode = (note == nil)

	return
}

type NoteFormCallback struct {
	note *models.Note

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (noteFormCallback *NoteFormCallback) OnSave() {

	log.Println("NoteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteFormCallback.probe.formStage.Checkout()

	if noteFormCallback.note == nil {
		noteFormCallback.note = new(models.Note).Stage(noteFormCallback.probe.stageOfInterest)
	}
	note_ := noteFormCallback.note
	_ = note_

	for _, formDiv := range noteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(note_.Name), formDiv)
		case "Type_":
			FormDivSelectFieldToField(&(note_.Type_), noteFormCallback.probe.stageOfInterest, formDiv)
		case "Accidental":
			FormDivSelectFieldToField(&(note_.Accidental), noteFormCallback.probe.stageOfInterest, formDiv)
		case "Time_modification":
			FormDivSelectFieldToField(&(note_.Time_modification), noteFormCallback.probe.stageOfInterest, formDiv)
		case "Stem":
			FormDivSelectFieldToField(&(note_.Stem), noteFormCallback.probe.stageOfInterest, formDiv)
		case "Notehead":
			FormDivSelectFieldToField(&(note_.Notehead), noteFormCallback.probe.stageOfInterest, formDiv)
		case "Notehead_text":
			FormDivSelectFieldToField(&(note_.Notehead_text), noteFormCallback.probe.stageOfInterest, formDiv)
		case "Beam":
			FormDivSelectFieldToField(&(note_.Beam), noteFormCallback.probe.stageOfInterest, formDiv)
		case "Play":
			FormDivSelectFieldToField(&(note_.Play), noteFormCallback.probe.stageOfInterest, formDiv)
		case "Listen":
			FormDivSelectFieldToField(&(note_.Listen), noteFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		note_.Unstage(noteFormCallback.probe.stageOfInterest)
	}

	noteFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Note](
		noteFormCallback.probe,
	)
	noteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if noteFormCallback.CreationMode || noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(noteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			noteFormCallback.probe,
			newFormGroup,
		)
		note := new(models.Note)
		FillUpForm(note, newFormGroup, noteFormCallback.probe)
		noteFormCallback.probe.formStage.Commit()
	}

	fillUpTree(noteFormCallback.probe)
}
func __gong__New__Note_sizeFormCallback(
	note_size *models.Note_size,
	probe *Probe,
	formGroup *table.FormGroup,
) (note_sizeFormCallback *Note_sizeFormCallback) {
	note_sizeFormCallback = new(Note_sizeFormCallback)
	note_sizeFormCallback.probe = probe
	note_sizeFormCallback.note_size = note_size
	note_sizeFormCallback.formGroup = formGroup

	note_sizeFormCallback.CreationMode = (note_size == nil)

	return
}

type Note_sizeFormCallback struct {
	note_size *models.Note_size

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (note_sizeFormCallback *Note_sizeFormCallback) OnSave() {

	log.Println("Note_sizeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	note_sizeFormCallback.probe.formStage.Checkout()

	if note_sizeFormCallback.note_size == nil {
		note_sizeFormCallback.note_size = new(models.Note_size).Stage(note_sizeFormCallback.probe.stageOfInterest)
	}
	note_size_ := note_sizeFormCallback.note_size
	_ = note_size_

	for _, formDiv := range note_sizeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(note_size_.Name), formDiv)
		case "Appearance:Note_size":
			// we need to retrieve the field owner before the change
			var pastAppearanceOwner *models.Appearance
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Note_size"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				note_sizeFormCallback.probe.stageOfInterest,
				note_sizeFormCallback.probe.backRepoOfInterest,
				note_size_,
				&rf)

			if reverseFieldOwner != nil {
				pastAppearanceOwner = reverseFieldOwner.(*models.Appearance)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAppearanceOwner != nil {
					idx := slices.Index(pastAppearanceOwner.Note_size, note_size_)
					pastAppearanceOwner.Note_size = slices.Delete(pastAppearanceOwner.Note_size, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _appearance := range *models.GetGongstructInstancesSet[models.Appearance](note_sizeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _appearance.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAppearanceOwner := _appearance // we have a match
						if pastAppearanceOwner != nil {
							if newAppearanceOwner != pastAppearanceOwner {
								idx := slices.Index(pastAppearanceOwner.Note_size, note_size_)
								pastAppearanceOwner.Note_size = slices.Delete(pastAppearanceOwner.Note_size, idx, idx+1)
								newAppearanceOwner.Note_size = append(newAppearanceOwner.Note_size, note_size_)
							}
						} else {
							newAppearanceOwner.Note_size = append(newAppearanceOwner.Note_size, note_size_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if note_sizeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		note_size_.Unstage(note_sizeFormCallback.probe.stageOfInterest)
	}

	note_sizeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Note_size](
		note_sizeFormCallback.probe,
	)
	note_sizeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if note_sizeFormCallback.CreationMode || note_sizeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		note_sizeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(note_sizeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Note_sizeFormCallback(
			nil,
			note_sizeFormCallback.probe,
			newFormGroup,
		)
		note_size := new(models.Note_size)
		FillUpForm(note_size, newFormGroup, note_sizeFormCallback.probe)
		note_sizeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(note_sizeFormCallback.probe)
}
func __gong__New__Note_typeFormCallback(
	note_type *models.Note_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (note_typeFormCallback *Note_typeFormCallback) {
	note_typeFormCallback = new(Note_typeFormCallback)
	note_typeFormCallback.probe = probe
	note_typeFormCallback.note_type = note_type
	note_typeFormCallback.formGroup = formGroup

	note_typeFormCallback.CreationMode = (note_type == nil)

	return
}

type Note_typeFormCallback struct {
	note_type *models.Note_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (note_typeFormCallback *Note_typeFormCallback) OnSave() {

	log.Println("Note_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	note_typeFormCallback.probe.formStage.Checkout()

	if note_typeFormCallback.note_type == nil {
		note_typeFormCallback.note_type = new(models.Note_type).Stage(note_typeFormCallback.probe.stageOfInterest)
	}
	note_type_ := note_typeFormCallback.note_type
	_ = note_type_

	for _, formDiv := range note_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(note_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if note_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		note_type_.Unstage(note_typeFormCallback.probe.stageOfInterest)
	}

	note_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Note_type](
		note_typeFormCallback.probe,
	)
	note_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if note_typeFormCallback.CreationMode || note_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		note_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(note_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Note_typeFormCallback(
			nil,
			note_typeFormCallback.probe,
			newFormGroup,
		)
		note_type := new(models.Note_type)
		FillUpForm(note_type, newFormGroup, note_typeFormCallback.probe)
		note_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(note_typeFormCallback.probe)
}
func __gong__New__NoteheadFormCallback(
	notehead *models.Notehead,
	probe *Probe,
	formGroup *table.FormGroup,
) (noteheadFormCallback *NoteheadFormCallback) {
	noteheadFormCallback = new(NoteheadFormCallback)
	noteheadFormCallback.probe = probe
	noteheadFormCallback.notehead = notehead
	noteheadFormCallback.formGroup = formGroup

	noteheadFormCallback.CreationMode = (notehead == nil)

	return
}

type NoteheadFormCallback struct {
	notehead *models.Notehead

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (noteheadFormCallback *NoteheadFormCallback) OnSave() {

	log.Println("NoteheadFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteheadFormCallback.probe.formStage.Checkout()

	if noteheadFormCallback.notehead == nil {
		noteheadFormCallback.notehead = new(models.Notehead).Stage(noteheadFormCallback.probe.stageOfInterest)
	}
	notehead_ := noteheadFormCallback.notehead
	_ = notehead_

	for _, formDiv := range noteheadFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notehead_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if noteheadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notehead_.Unstage(noteheadFormCallback.probe.stageOfInterest)
	}

	noteheadFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Notehead](
		noteheadFormCallback.probe,
	)
	noteheadFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if noteheadFormCallback.CreationMode || noteheadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteheadFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(noteheadFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteheadFormCallback(
			nil,
			noteheadFormCallback.probe,
			newFormGroup,
		)
		notehead := new(models.Notehead)
		FillUpForm(notehead, newFormGroup, noteheadFormCallback.probe)
		noteheadFormCallback.probe.formStage.Commit()
	}

	fillUpTree(noteheadFormCallback.probe)
}
func __gong__New__Notehead_textFormCallback(
	notehead_text *models.Notehead_text,
	probe *Probe,
	formGroup *table.FormGroup,
) (notehead_textFormCallback *Notehead_textFormCallback) {
	notehead_textFormCallback = new(Notehead_textFormCallback)
	notehead_textFormCallback.probe = probe
	notehead_textFormCallback.notehead_text = notehead_text
	notehead_textFormCallback.formGroup = formGroup

	notehead_textFormCallback.CreationMode = (notehead_text == nil)

	return
}

type Notehead_textFormCallback struct {
	notehead_text *models.Notehead_text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (notehead_textFormCallback *Notehead_textFormCallback) OnSave() {

	log.Println("Notehead_textFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notehead_textFormCallback.probe.formStage.Checkout()

	if notehead_textFormCallback.notehead_text == nil {
		notehead_textFormCallback.notehead_text = new(models.Notehead_text).Stage(notehead_textFormCallback.probe.stageOfInterest)
	}
	notehead_text_ := notehead_textFormCallback.notehead_text
	_ = notehead_text_

	for _, formDiv := range notehead_textFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notehead_text_.Name), formDiv)
		case "Accidental_text":
			FormDivSelectFieldToField(&(notehead_text_.Accidental_text), notehead_textFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if notehead_textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notehead_text_.Unstage(notehead_textFormCallback.probe.stageOfInterest)
	}

	notehead_textFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Notehead_text](
		notehead_textFormCallback.probe,
	)
	notehead_textFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if notehead_textFormCallback.CreationMode || notehead_textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notehead_textFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(notehead_textFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Notehead_textFormCallback(
			nil,
			notehead_textFormCallback.probe,
			newFormGroup,
		)
		notehead_text := new(models.Notehead_text)
		FillUpForm(notehead_text, newFormGroup, notehead_textFormCallback.probe)
		notehead_textFormCallback.probe.formStage.Commit()
	}

	fillUpTree(notehead_textFormCallback.probe)
}
func __gong__New__NumeralFormCallback(
	numeral *models.Numeral,
	probe *Probe,
	formGroup *table.FormGroup,
) (numeralFormCallback *NumeralFormCallback) {
	numeralFormCallback = new(NumeralFormCallback)
	numeralFormCallback.probe = probe
	numeralFormCallback.numeral = numeral
	numeralFormCallback.formGroup = formGroup

	numeralFormCallback.CreationMode = (numeral == nil)

	return
}

type NumeralFormCallback struct {
	numeral *models.Numeral

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (numeralFormCallback *NumeralFormCallback) OnSave() {

	log.Println("NumeralFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	numeralFormCallback.probe.formStage.Checkout()

	if numeralFormCallback.numeral == nil {
		numeralFormCallback.numeral = new(models.Numeral).Stage(numeralFormCallback.probe.stageOfInterest)
	}
	numeral_ := numeralFormCallback.numeral
	_ = numeral_

	for _, formDiv := range numeralFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(numeral_.Name), formDiv)
		case "Numeral_root":
			FormDivSelectFieldToField(&(numeral_.Numeral_root), numeralFormCallback.probe.stageOfInterest, formDiv)
		case "Numeral_alter":
			FormDivSelectFieldToField(&(numeral_.Numeral_alter), numeralFormCallback.probe.stageOfInterest, formDiv)
		case "Numeral_key":
			FormDivSelectFieldToField(&(numeral_.Numeral_key), numeralFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if numeralFormCallback.formGroup.HasSuppressButtonBeenPressed {
		numeral_.Unstage(numeralFormCallback.probe.stageOfInterest)
	}

	numeralFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Numeral](
		numeralFormCallback.probe,
	)
	numeralFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if numeralFormCallback.CreationMode || numeralFormCallback.formGroup.HasSuppressButtonBeenPressed {
		numeralFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(numeralFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NumeralFormCallback(
			nil,
			numeralFormCallback.probe,
			newFormGroup,
		)
		numeral := new(models.Numeral)
		FillUpForm(numeral, newFormGroup, numeralFormCallback.probe)
		numeralFormCallback.probe.formStage.Commit()
	}

	fillUpTree(numeralFormCallback.probe)
}
func __gong__New__Numeral_keyFormCallback(
	numeral_key *models.Numeral_key,
	probe *Probe,
	formGroup *table.FormGroup,
) (numeral_keyFormCallback *Numeral_keyFormCallback) {
	numeral_keyFormCallback = new(Numeral_keyFormCallback)
	numeral_keyFormCallback.probe = probe
	numeral_keyFormCallback.numeral_key = numeral_key
	numeral_keyFormCallback.formGroup = formGroup

	numeral_keyFormCallback.CreationMode = (numeral_key == nil)

	return
}

type Numeral_keyFormCallback struct {
	numeral_key *models.Numeral_key

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (numeral_keyFormCallback *Numeral_keyFormCallback) OnSave() {

	log.Println("Numeral_keyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	numeral_keyFormCallback.probe.formStage.Checkout()

	if numeral_keyFormCallback.numeral_key == nil {
		numeral_keyFormCallback.numeral_key = new(models.Numeral_key).Stage(numeral_keyFormCallback.probe.stageOfInterest)
	}
	numeral_key_ := numeral_keyFormCallback.numeral_key
	_ = numeral_key_

	for _, formDiv := range numeral_keyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(numeral_key_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if numeral_keyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		numeral_key_.Unstage(numeral_keyFormCallback.probe.stageOfInterest)
	}

	numeral_keyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Numeral_key](
		numeral_keyFormCallback.probe,
	)
	numeral_keyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if numeral_keyFormCallback.CreationMode || numeral_keyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		numeral_keyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(numeral_keyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Numeral_keyFormCallback(
			nil,
			numeral_keyFormCallback.probe,
			newFormGroup,
		)
		numeral_key := new(models.Numeral_key)
		FillUpForm(numeral_key, newFormGroup, numeral_keyFormCallback.probe)
		numeral_keyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(numeral_keyFormCallback.probe)
}
func __gong__New__Numeral_rootFormCallback(
	numeral_root *models.Numeral_root,
	probe *Probe,
	formGroup *table.FormGroup,
) (numeral_rootFormCallback *Numeral_rootFormCallback) {
	numeral_rootFormCallback = new(Numeral_rootFormCallback)
	numeral_rootFormCallback.probe = probe
	numeral_rootFormCallback.numeral_root = numeral_root
	numeral_rootFormCallback.formGroup = formGroup

	numeral_rootFormCallback.CreationMode = (numeral_root == nil)

	return
}

type Numeral_rootFormCallback struct {
	numeral_root *models.Numeral_root

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (numeral_rootFormCallback *Numeral_rootFormCallback) OnSave() {

	log.Println("Numeral_rootFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	numeral_rootFormCallback.probe.formStage.Checkout()

	if numeral_rootFormCallback.numeral_root == nil {
		numeral_rootFormCallback.numeral_root = new(models.Numeral_root).Stage(numeral_rootFormCallback.probe.stageOfInterest)
	}
	numeral_root_ := numeral_rootFormCallback.numeral_root
	_ = numeral_root_

	for _, formDiv := range numeral_rootFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(numeral_root_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(numeral_root_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if numeral_rootFormCallback.formGroup.HasSuppressButtonBeenPressed {
		numeral_root_.Unstage(numeral_rootFormCallback.probe.stageOfInterest)
	}

	numeral_rootFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Numeral_root](
		numeral_rootFormCallback.probe,
	)
	numeral_rootFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if numeral_rootFormCallback.CreationMode || numeral_rootFormCallback.formGroup.HasSuppressButtonBeenPressed {
		numeral_rootFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(numeral_rootFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Numeral_rootFormCallback(
			nil,
			numeral_rootFormCallback.probe,
			newFormGroup,
		)
		numeral_root := new(models.Numeral_root)
		FillUpForm(numeral_root, newFormGroup, numeral_rootFormCallback.probe)
		numeral_rootFormCallback.probe.formStage.Commit()
	}

	fillUpTree(numeral_rootFormCallback.probe)
}
func __gong__New__Octave_shiftFormCallback(
	octave_shift *models.Octave_shift,
	probe *Probe,
	formGroup *table.FormGroup,
) (octave_shiftFormCallback *Octave_shiftFormCallback) {
	octave_shiftFormCallback = new(Octave_shiftFormCallback)
	octave_shiftFormCallback.probe = probe
	octave_shiftFormCallback.octave_shift = octave_shift
	octave_shiftFormCallback.formGroup = formGroup

	octave_shiftFormCallback.CreationMode = (octave_shift == nil)

	return
}

type Octave_shiftFormCallback struct {
	octave_shift *models.Octave_shift

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (octave_shiftFormCallback *Octave_shiftFormCallback) OnSave() {

	log.Println("Octave_shiftFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	octave_shiftFormCallback.probe.formStage.Checkout()

	if octave_shiftFormCallback.octave_shift == nil {
		octave_shiftFormCallback.octave_shift = new(models.Octave_shift).Stage(octave_shiftFormCallback.probe.stageOfInterest)
	}
	octave_shift_ := octave_shiftFormCallback.octave_shift
	_ = octave_shift_

	for _, formDiv := range octave_shiftFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(octave_shift_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if octave_shiftFormCallback.formGroup.HasSuppressButtonBeenPressed {
		octave_shift_.Unstage(octave_shiftFormCallback.probe.stageOfInterest)
	}

	octave_shiftFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Octave_shift](
		octave_shiftFormCallback.probe,
	)
	octave_shiftFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if octave_shiftFormCallback.CreationMode || octave_shiftFormCallback.formGroup.HasSuppressButtonBeenPressed {
		octave_shiftFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(octave_shiftFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Octave_shiftFormCallback(
			nil,
			octave_shiftFormCallback.probe,
			newFormGroup,
		)
		octave_shift := new(models.Octave_shift)
		FillUpForm(octave_shift, newFormGroup, octave_shiftFormCallback.probe)
		octave_shiftFormCallback.probe.formStage.Commit()
	}

	fillUpTree(octave_shiftFormCallback.probe)
}
func __gong__New__OffsetFormCallback(
	offset *models.Offset,
	probe *Probe,
	formGroup *table.FormGroup,
) (offsetFormCallback *OffsetFormCallback) {
	offsetFormCallback = new(OffsetFormCallback)
	offsetFormCallback.probe = probe
	offsetFormCallback.offset = offset
	offsetFormCallback.formGroup = formGroup

	offsetFormCallback.CreationMode = (offset == nil)

	return
}

type OffsetFormCallback struct {
	offset *models.Offset

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (offsetFormCallback *OffsetFormCallback) OnSave() {

	log.Println("OffsetFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	offsetFormCallback.probe.formStage.Checkout()

	if offsetFormCallback.offset == nil {
		offsetFormCallback.offset = new(models.Offset).Stage(offsetFormCallback.probe.stageOfInterest)
	}
	offset_ := offsetFormCallback.offset
	_ = offset_

	for _, formDiv := range offsetFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(offset_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if offsetFormCallback.formGroup.HasSuppressButtonBeenPressed {
		offset_.Unstage(offsetFormCallback.probe.stageOfInterest)
	}

	offsetFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Offset](
		offsetFormCallback.probe,
	)
	offsetFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if offsetFormCallback.CreationMode || offsetFormCallback.formGroup.HasSuppressButtonBeenPressed {
		offsetFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(offsetFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__OffsetFormCallback(
			nil,
			offsetFormCallback.probe,
			newFormGroup,
		)
		offset := new(models.Offset)
		FillUpForm(offset, newFormGroup, offsetFormCallback.probe)
		offsetFormCallback.probe.formStage.Commit()
	}

	fillUpTree(offsetFormCallback.probe)
}
func __gong__New__OpusFormCallback(
	opus *models.Opus,
	probe *Probe,
	formGroup *table.FormGroup,
) (opusFormCallback *OpusFormCallback) {
	opusFormCallback = new(OpusFormCallback)
	opusFormCallback.probe = probe
	opusFormCallback.opus = opus
	opusFormCallback.formGroup = formGroup

	opusFormCallback.CreationMode = (opus == nil)

	return
}

type OpusFormCallback struct {
	opus *models.Opus

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (opusFormCallback *OpusFormCallback) OnSave() {

	log.Println("OpusFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	opusFormCallback.probe.formStage.Checkout()

	if opusFormCallback.opus == nil {
		opusFormCallback.opus = new(models.Opus).Stage(opusFormCallback.probe.stageOfInterest)
	}
	opus_ := opusFormCallback.opus
	_ = opus_

	for _, formDiv := range opusFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(opus_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if opusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		opus_.Unstage(opusFormCallback.probe.stageOfInterest)
	}

	opusFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Opus](
		opusFormCallback.probe,
	)
	opusFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if opusFormCallback.CreationMode || opusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		opusFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(opusFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__OpusFormCallback(
			nil,
			opusFormCallback.probe,
			newFormGroup,
		)
		opus := new(models.Opus)
		FillUpForm(opus, newFormGroup, opusFormCallback.probe)
		opusFormCallback.probe.formStage.Commit()
	}

	fillUpTree(opusFormCallback.probe)
}
func __gong__New__OrnamentsFormCallback(
	ornaments *models.Ornaments,
	probe *Probe,
	formGroup *table.FormGroup,
) (ornamentsFormCallback *OrnamentsFormCallback) {
	ornamentsFormCallback = new(OrnamentsFormCallback)
	ornamentsFormCallback.probe = probe
	ornamentsFormCallback.ornaments = ornaments
	ornamentsFormCallback.formGroup = formGroup

	ornamentsFormCallback.CreationMode = (ornaments == nil)

	return
}

type OrnamentsFormCallback struct {
	ornaments *models.Ornaments

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (ornamentsFormCallback *OrnamentsFormCallback) OnSave() {

	log.Println("OrnamentsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ornamentsFormCallback.probe.formStage.Checkout()

	if ornamentsFormCallback.ornaments == nil {
		ornamentsFormCallback.ornaments = new(models.Ornaments).Stage(ornamentsFormCallback.probe.stageOfInterest)
	}
	ornaments_ := ornamentsFormCallback.ornaments
	_ = ornaments_

	for _, formDiv := range ornamentsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(ornaments_.Name), formDiv)
		case "Trill_mark":
			FormDivSelectFieldToField(&(ornaments_.Trill_mark), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Turn":
			FormDivSelectFieldToField(&(ornaments_.Turn), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Delayed_turn":
			FormDivSelectFieldToField(&(ornaments_.Delayed_turn), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Inverted_turn":
			FormDivSelectFieldToField(&(ornaments_.Inverted_turn), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Delayed_inverted_turn":
			FormDivSelectFieldToField(&(ornaments_.Delayed_inverted_turn), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Vertical_turn":
			FormDivSelectFieldToField(&(ornaments_.Vertical_turn), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Inverted_vertical_turn":
			FormDivSelectFieldToField(&(ornaments_.Inverted_vertical_turn), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Shake":
			FormDivSelectFieldToField(&(ornaments_.Shake), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Wavy_line":
			FormDivSelectFieldToField(&(ornaments_.Wavy_line), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Mordent":
			FormDivSelectFieldToField(&(ornaments_.Mordent), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Inverted_mordent":
			FormDivSelectFieldToField(&(ornaments_.Inverted_mordent), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Schleifer":
			FormDivSelectFieldToField(&(ornaments_.Schleifer), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Tremolo":
			FormDivSelectFieldToField(&(ornaments_.Tremolo), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		case "Haydn":
			FormDivSelectFieldToField(&(ornaments_.Haydn), ornamentsFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if ornamentsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ornaments_.Unstage(ornamentsFormCallback.probe.stageOfInterest)
	}

	ornamentsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Ornaments](
		ornamentsFormCallback.probe,
	)
	ornamentsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if ornamentsFormCallback.CreationMode || ornamentsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ornamentsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(ornamentsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__OrnamentsFormCallback(
			nil,
			ornamentsFormCallback.probe,
			newFormGroup,
		)
		ornaments := new(models.Ornaments)
		FillUpForm(ornaments, newFormGroup, ornamentsFormCallback.probe)
		ornamentsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(ornamentsFormCallback.probe)
}
func __gong__New__Other_appearanceFormCallback(
	other_appearance *models.Other_appearance,
	probe *Probe,
	formGroup *table.FormGroup,
) (other_appearanceFormCallback *Other_appearanceFormCallback) {
	other_appearanceFormCallback = new(Other_appearanceFormCallback)
	other_appearanceFormCallback.probe = probe
	other_appearanceFormCallback.other_appearance = other_appearance
	other_appearanceFormCallback.formGroup = formGroup

	other_appearanceFormCallback.CreationMode = (other_appearance == nil)

	return
}

type Other_appearanceFormCallback struct {
	other_appearance *models.Other_appearance

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (other_appearanceFormCallback *Other_appearanceFormCallback) OnSave() {

	log.Println("Other_appearanceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	other_appearanceFormCallback.probe.formStage.Checkout()

	if other_appearanceFormCallback.other_appearance == nil {
		other_appearanceFormCallback.other_appearance = new(models.Other_appearance).Stage(other_appearanceFormCallback.probe.stageOfInterest)
	}
	other_appearance_ := other_appearanceFormCallback.other_appearance
	_ = other_appearance_

	for _, formDiv := range other_appearanceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(other_appearance_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(other_appearance_.Value), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(other_appearance_.Type), formDiv)
		case "Appearance:Other_appearance":
			// we need to retrieve the field owner before the change
			var pastAppearanceOwner *models.Appearance
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Other_appearance"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				other_appearanceFormCallback.probe.stageOfInterest,
				other_appearanceFormCallback.probe.backRepoOfInterest,
				other_appearance_,
				&rf)

			if reverseFieldOwner != nil {
				pastAppearanceOwner = reverseFieldOwner.(*models.Appearance)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAppearanceOwner != nil {
					idx := slices.Index(pastAppearanceOwner.Other_appearance, other_appearance_)
					pastAppearanceOwner.Other_appearance = slices.Delete(pastAppearanceOwner.Other_appearance, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _appearance := range *models.GetGongstructInstancesSet[models.Appearance](other_appearanceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _appearance.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAppearanceOwner := _appearance // we have a match
						if pastAppearanceOwner != nil {
							if newAppearanceOwner != pastAppearanceOwner {
								idx := slices.Index(pastAppearanceOwner.Other_appearance, other_appearance_)
								pastAppearanceOwner.Other_appearance = slices.Delete(pastAppearanceOwner.Other_appearance, idx, idx+1)
								newAppearanceOwner.Other_appearance = append(newAppearanceOwner.Other_appearance, other_appearance_)
							}
						} else {
							newAppearanceOwner.Other_appearance = append(newAppearanceOwner.Other_appearance, other_appearance_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if other_appearanceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		other_appearance_.Unstage(other_appearanceFormCallback.probe.stageOfInterest)
	}

	other_appearanceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Other_appearance](
		other_appearanceFormCallback.probe,
	)
	other_appearanceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if other_appearanceFormCallback.CreationMode || other_appearanceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		other_appearanceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(other_appearanceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Other_appearanceFormCallback(
			nil,
			other_appearanceFormCallback.probe,
			newFormGroup,
		)
		other_appearance := new(models.Other_appearance)
		FillUpForm(other_appearance, newFormGroup, other_appearanceFormCallback.probe)
		other_appearanceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(other_appearanceFormCallback.probe)
}
func __gong__New__Other_listeningFormCallback(
	other_listening *models.Other_listening,
	probe *Probe,
	formGroup *table.FormGroup,
) (other_listeningFormCallback *Other_listeningFormCallback) {
	other_listeningFormCallback = new(Other_listeningFormCallback)
	other_listeningFormCallback.probe = probe
	other_listeningFormCallback.other_listening = other_listening
	other_listeningFormCallback.formGroup = formGroup

	other_listeningFormCallback.CreationMode = (other_listening == nil)

	return
}

type Other_listeningFormCallback struct {
	other_listening *models.Other_listening

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (other_listeningFormCallback *Other_listeningFormCallback) OnSave() {

	log.Println("Other_listeningFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	other_listeningFormCallback.probe.formStage.Checkout()

	if other_listeningFormCallback.other_listening == nil {
		other_listeningFormCallback.other_listening = new(models.Other_listening).Stage(other_listeningFormCallback.probe.stageOfInterest)
	}
	other_listening_ := other_listeningFormCallback.other_listening
	_ = other_listening_

	for _, formDiv := range other_listeningFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(other_listening_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(other_listening_.Value), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(other_listening_.Type), formDiv)
		}
	}

	// manage the suppress operation
	if other_listeningFormCallback.formGroup.HasSuppressButtonBeenPressed {
		other_listening_.Unstage(other_listeningFormCallback.probe.stageOfInterest)
	}

	other_listeningFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Other_listening](
		other_listeningFormCallback.probe,
	)
	other_listeningFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if other_listeningFormCallback.CreationMode || other_listeningFormCallback.formGroup.HasSuppressButtonBeenPressed {
		other_listeningFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(other_listeningFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Other_listeningFormCallback(
			nil,
			other_listeningFormCallback.probe,
			newFormGroup,
		)
		other_listening := new(models.Other_listening)
		FillUpForm(other_listening, newFormGroup, other_listeningFormCallback.probe)
		other_listeningFormCallback.probe.formStage.Commit()
	}

	fillUpTree(other_listeningFormCallback.probe)
}
func __gong__New__Other_notationFormCallback(
	other_notation *models.Other_notation,
	probe *Probe,
	formGroup *table.FormGroup,
) (other_notationFormCallback *Other_notationFormCallback) {
	other_notationFormCallback = new(Other_notationFormCallback)
	other_notationFormCallback.probe = probe
	other_notationFormCallback.other_notation = other_notation
	other_notationFormCallback.formGroup = formGroup

	other_notationFormCallback.CreationMode = (other_notation == nil)

	return
}

type Other_notationFormCallback struct {
	other_notation *models.Other_notation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (other_notationFormCallback *Other_notationFormCallback) OnSave() {

	log.Println("Other_notationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	other_notationFormCallback.probe.formStage.Checkout()

	if other_notationFormCallback.other_notation == nil {
		other_notationFormCallback.other_notation = new(models.Other_notation).Stage(other_notationFormCallback.probe.stageOfInterest)
	}
	other_notation_ := other_notationFormCallback.other_notation
	_ = other_notation_

	for _, formDiv := range other_notationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(other_notation_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(other_notation_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if other_notationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		other_notation_.Unstage(other_notationFormCallback.probe.stageOfInterest)
	}

	other_notationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Other_notation](
		other_notationFormCallback.probe,
	)
	other_notationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if other_notationFormCallback.CreationMode || other_notationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		other_notationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(other_notationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Other_notationFormCallback(
			nil,
			other_notationFormCallback.probe,
			newFormGroup,
		)
		other_notation := new(models.Other_notation)
		FillUpForm(other_notation, newFormGroup, other_notationFormCallback.probe)
		other_notationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(other_notationFormCallback.probe)
}
func __gong__New__Other_playFormCallback(
	other_play *models.Other_play,
	probe *Probe,
	formGroup *table.FormGroup,
) (other_playFormCallback *Other_playFormCallback) {
	other_playFormCallback = new(Other_playFormCallback)
	other_playFormCallback.probe = probe
	other_playFormCallback.other_play = other_play
	other_playFormCallback.formGroup = formGroup

	other_playFormCallback.CreationMode = (other_play == nil)

	return
}

type Other_playFormCallback struct {
	other_play *models.Other_play

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (other_playFormCallback *Other_playFormCallback) OnSave() {

	log.Println("Other_playFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	other_playFormCallback.probe.formStage.Checkout()

	if other_playFormCallback.other_play == nil {
		other_playFormCallback.other_play = new(models.Other_play).Stage(other_playFormCallback.probe.stageOfInterest)
	}
	other_play_ := other_playFormCallback.other_play
	_ = other_play_

	for _, formDiv := range other_playFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(other_play_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(other_play_.Value), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(other_play_.Type), formDiv)
		}
	}

	// manage the suppress operation
	if other_playFormCallback.formGroup.HasSuppressButtonBeenPressed {
		other_play_.Unstage(other_playFormCallback.probe.stageOfInterest)
	}

	other_playFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Other_play](
		other_playFormCallback.probe,
	)
	other_playFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if other_playFormCallback.CreationMode || other_playFormCallback.formGroup.HasSuppressButtonBeenPressed {
		other_playFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(other_playFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Other_playFormCallback(
			nil,
			other_playFormCallback.probe,
			newFormGroup,
		)
		other_play := new(models.Other_play)
		FillUpForm(other_play, newFormGroup, other_playFormCallback.probe)
		other_playFormCallback.probe.formStage.Commit()
	}

	fillUpTree(other_playFormCallback.probe)
}
func __gong__New__Page_layoutFormCallback(
	page_layout *models.Page_layout,
	probe *Probe,
	formGroup *table.FormGroup,
) (page_layoutFormCallback *Page_layoutFormCallback) {
	page_layoutFormCallback = new(Page_layoutFormCallback)
	page_layoutFormCallback.probe = probe
	page_layoutFormCallback.page_layout = page_layout
	page_layoutFormCallback.formGroup = formGroup

	page_layoutFormCallback.CreationMode = (page_layout == nil)

	return
}

type Page_layoutFormCallback struct {
	page_layout *models.Page_layout

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (page_layoutFormCallback *Page_layoutFormCallback) OnSave() {

	log.Println("Page_layoutFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	page_layoutFormCallback.probe.formStage.Checkout()

	if page_layoutFormCallback.page_layout == nil {
		page_layoutFormCallback.page_layout = new(models.Page_layout).Stage(page_layoutFormCallback.probe.stageOfInterest)
	}
	page_layout_ := page_layoutFormCallback.page_layout
	_ = page_layout_

	for _, formDiv := range page_layoutFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(page_layout_.Name), formDiv)
		case "Page_margins":
			FormDivSelectFieldToField(&(page_layout_.Page_margins), page_layoutFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if page_layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		page_layout_.Unstage(page_layoutFormCallback.probe.stageOfInterest)
	}

	page_layoutFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Page_layout](
		page_layoutFormCallback.probe,
	)
	page_layoutFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if page_layoutFormCallback.CreationMode || page_layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		page_layoutFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(page_layoutFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Page_layoutFormCallback(
			nil,
			page_layoutFormCallback.probe,
			newFormGroup,
		)
		page_layout := new(models.Page_layout)
		FillUpForm(page_layout, newFormGroup, page_layoutFormCallback.probe)
		page_layoutFormCallback.probe.formStage.Commit()
	}

	fillUpTree(page_layoutFormCallback.probe)
}
func __gong__New__Page_marginsFormCallback(
	page_margins *models.Page_margins,
	probe *Probe,
	formGroup *table.FormGroup,
) (page_marginsFormCallback *Page_marginsFormCallback) {
	page_marginsFormCallback = new(Page_marginsFormCallback)
	page_marginsFormCallback.probe = probe
	page_marginsFormCallback.page_margins = page_margins
	page_marginsFormCallback.formGroup = formGroup

	page_marginsFormCallback.CreationMode = (page_margins == nil)

	return
}

type Page_marginsFormCallback struct {
	page_margins *models.Page_margins

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (page_marginsFormCallback *Page_marginsFormCallback) OnSave() {

	log.Println("Page_marginsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	page_marginsFormCallback.probe.formStage.Checkout()

	if page_marginsFormCallback.page_margins == nil {
		page_marginsFormCallback.page_margins = new(models.Page_margins).Stage(page_marginsFormCallback.probe.stageOfInterest)
	}
	page_margins_ := page_marginsFormCallback.page_margins
	_ = page_margins_

	for _, formDiv := range page_marginsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(page_margins_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if page_marginsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		page_margins_.Unstage(page_marginsFormCallback.probe.stageOfInterest)
	}

	page_marginsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Page_margins](
		page_marginsFormCallback.probe,
	)
	page_marginsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if page_marginsFormCallback.CreationMode || page_marginsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		page_marginsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(page_marginsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Page_marginsFormCallback(
			nil,
			page_marginsFormCallback.probe,
			newFormGroup,
		)
		page_margins := new(models.Page_margins)
		FillUpForm(page_margins, newFormGroup, page_marginsFormCallback.probe)
		page_marginsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(page_marginsFormCallback.probe)
}
func __gong__New__Part_clefFormCallback(
	part_clef *models.Part_clef,
	probe *Probe,
	formGroup *table.FormGroup,
) (part_clefFormCallback *Part_clefFormCallback) {
	part_clefFormCallback = new(Part_clefFormCallback)
	part_clefFormCallback.probe = probe
	part_clefFormCallback.part_clef = part_clef
	part_clefFormCallback.formGroup = formGroup

	part_clefFormCallback.CreationMode = (part_clef == nil)

	return
}

type Part_clefFormCallback struct {
	part_clef *models.Part_clef

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (part_clefFormCallback *Part_clefFormCallback) OnSave() {

	log.Println("Part_clefFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	part_clefFormCallback.probe.formStage.Checkout()

	if part_clefFormCallback.part_clef == nil {
		part_clefFormCallback.part_clef = new(models.Part_clef).Stage(part_clefFormCallback.probe.stageOfInterest)
	}
	part_clef_ := part_clefFormCallback.part_clef
	_ = part_clef_

	for _, formDiv := range part_clefFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(part_clef_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if part_clefFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_clef_.Unstage(part_clefFormCallback.probe.stageOfInterest)
	}

	part_clefFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Part_clef](
		part_clefFormCallback.probe,
	)
	part_clefFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if part_clefFormCallback.CreationMode || part_clefFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_clefFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(part_clefFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Part_clefFormCallback(
			nil,
			part_clefFormCallback.probe,
			newFormGroup,
		)
		part_clef := new(models.Part_clef)
		FillUpForm(part_clef, newFormGroup, part_clefFormCallback.probe)
		part_clefFormCallback.probe.formStage.Commit()
	}

	fillUpTree(part_clefFormCallback.probe)
}
func __gong__New__Part_groupFormCallback(
	part_group *models.Part_group,
	probe *Probe,
	formGroup *table.FormGroup,
) (part_groupFormCallback *Part_groupFormCallback) {
	part_groupFormCallback = new(Part_groupFormCallback)
	part_groupFormCallback.probe = probe
	part_groupFormCallback.part_group = part_group
	part_groupFormCallback.formGroup = formGroup

	part_groupFormCallback.CreationMode = (part_group == nil)

	return
}

type Part_groupFormCallback struct {
	part_group *models.Part_group

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (part_groupFormCallback *Part_groupFormCallback) OnSave() {

	log.Println("Part_groupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	part_groupFormCallback.probe.formStage.Checkout()

	if part_groupFormCallback.part_group == nil {
		part_groupFormCallback.part_group = new(models.Part_group).Stage(part_groupFormCallback.probe.stageOfInterest)
	}
	part_group_ := part_groupFormCallback.part_group
	_ = part_group_

	for _, formDiv := range part_groupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(part_group_.Name), formDiv)
		case "Group_name_display":
			FormDivSelectFieldToField(&(part_group_.Group_name_display), part_groupFormCallback.probe.stageOfInterest, formDiv)
		case "Group_abbreviation_display":
			FormDivSelectFieldToField(&(part_group_.Group_abbreviation_display), part_groupFormCallback.probe.stageOfInterest, formDiv)
		case "Group_symbol":
			FormDivSelectFieldToField(&(part_group_.Group_symbol), part_groupFormCallback.probe.stageOfInterest, formDiv)
		case "Group_barline":
			FormDivSelectFieldToField(&(part_group_.Group_barline), part_groupFormCallback.probe.stageOfInterest, formDiv)
		case "Group_time":
			FormDivSelectFieldToField(&(part_group_.Group_time), part_groupFormCallback.probe.stageOfInterest, formDiv)
		case "Number":
			FormDivBasicFieldToField(&(part_group_.Number), formDiv)
		}
	}

	// manage the suppress operation
	if part_groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_group_.Unstage(part_groupFormCallback.probe.stageOfInterest)
	}

	part_groupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Part_group](
		part_groupFormCallback.probe,
	)
	part_groupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if part_groupFormCallback.CreationMode || part_groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(part_groupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Part_groupFormCallback(
			nil,
			part_groupFormCallback.probe,
			newFormGroup,
		)
		part_group := new(models.Part_group)
		FillUpForm(part_group, newFormGroup, part_groupFormCallback.probe)
		part_groupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(part_groupFormCallback.probe)
}
func __gong__New__Part_linkFormCallback(
	part_link *models.Part_link,
	probe *Probe,
	formGroup *table.FormGroup,
) (part_linkFormCallback *Part_linkFormCallback) {
	part_linkFormCallback = new(Part_linkFormCallback)
	part_linkFormCallback.probe = probe
	part_linkFormCallback.part_link = part_link
	part_linkFormCallback.formGroup = formGroup

	part_linkFormCallback.CreationMode = (part_link == nil)

	return
}

type Part_linkFormCallback struct {
	part_link *models.Part_link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (part_linkFormCallback *Part_linkFormCallback) OnSave() {

	log.Println("Part_linkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	part_linkFormCallback.probe.formStage.Checkout()

	if part_linkFormCallback.part_link == nil {
		part_linkFormCallback.part_link = new(models.Part_link).Stage(part_linkFormCallback.probe.stageOfInterest)
	}
	part_link_ := part_linkFormCallback.part_link
	_ = part_link_

	for _, formDiv := range part_linkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(part_link_.Name), formDiv)
		case "Score_part:Part_link":
			// we need to retrieve the field owner before the change
			var pastScore_partOwner *models.Score_part
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Part_link"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				part_linkFormCallback.probe.stageOfInterest,
				part_linkFormCallback.probe.backRepoOfInterest,
				part_link_,
				&rf)

			if reverseFieldOwner != nil {
				pastScore_partOwner = reverseFieldOwner.(*models.Score_part)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastScore_partOwner != nil {
					idx := slices.Index(pastScore_partOwner.Part_link, part_link_)
					pastScore_partOwner.Part_link = slices.Delete(pastScore_partOwner.Part_link, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _score_part := range *models.GetGongstructInstancesSet[models.Score_part](part_linkFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _score_part.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newScore_partOwner := _score_part // we have a match
						if pastScore_partOwner != nil {
							if newScore_partOwner != pastScore_partOwner {
								idx := slices.Index(pastScore_partOwner.Part_link, part_link_)
								pastScore_partOwner.Part_link = slices.Delete(pastScore_partOwner.Part_link, idx, idx+1)
								newScore_partOwner.Part_link = append(newScore_partOwner.Part_link, part_link_)
							}
						} else {
							newScore_partOwner.Part_link = append(newScore_partOwner.Part_link, part_link_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if part_linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_link_.Unstage(part_linkFormCallback.probe.stageOfInterest)
	}

	part_linkFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Part_link](
		part_linkFormCallback.probe,
	)
	part_linkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if part_linkFormCallback.CreationMode || part_linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(part_linkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Part_linkFormCallback(
			nil,
			part_linkFormCallback.probe,
			newFormGroup,
		)
		part_link := new(models.Part_link)
		FillUpForm(part_link, newFormGroup, part_linkFormCallback.probe)
		part_linkFormCallback.probe.formStage.Commit()
	}

	fillUpTree(part_linkFormCallback.probe)
}
func __gong__New__Part_listFormCallback(
	part_list *models.Part_list,
	probe *Probe,
	formGroup *table.FormGroup,
) (part_listFormCallback *Part_listFormCallback) {
	part_listFormCallback = new(Part_listFormCallback)
	part_listFormCallback.probe = probe
	part_listFormCallback.part_list = part_list
	part_listFormCallback.formGroup = formGroup

	part_listFormCallback.CreationMode = (part_list == nil)

	return
}

type Part_listFormCallback struct {
	part_list *models.Part_list

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (part_listFormCallback *Part_listFormCallback) OnSave() {

	log.Println("Part_listFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	part_listFormCallback.probe.formStage.Checkout()

	if part_listFormCallback.part_list == nil {
		part_listFormCallback.part_list = new(models.Part_list).Stage(part_listFormCallback.probe.stageOfInterest)
	}
	part_list_ := part_listFormCallback.part_list
	_ = part_list_

	for _, formDiv := range part_listFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(part_list_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if part_listFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_list_.Unstage(part_listFormCallback.probe.stageOfInterest)
	}

	part_listFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Part_list](
		part_listFormCallback.probe,
	)
	part_listFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if part_listFormCallback.CreationMode || part_listFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_listFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(part_listFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Part_listFormCallback(
			nil,
			part_listFormCallback.probe,
			newFormGroup,
		)
		part_list := new(models.Part_list)
		FillUpForm(part_list, newFormGroup, part_listFormCallback.probe)
		part_listFormCallback.probe.formStage.Commit()
	}

	fillUpTree(part_listFormCallback.probe)
}
func __gong__New__Part_symbolFormCallback(
	part_symbol *models.Part_symbol,
	probe *Probe,
	formGroup *table.FormGroup,
) (part_symbolFormCallback *Part_symbolFormCallback) {
	part_symbolFormCallback = new(Part_symbolFormCallback)
	part_symbolFormCallback.probe = probe
	part_symbolFormCallback.part_symbol = part_symbol
	part_symbolFormCallback.formGroup = formGroup

	part_symbolFormCallback.CreationMode = (part_symbol == nil)

	return
}

type Part_symbolFormCallback struct {
	part_symbol *models.Part_symbol

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (part_symbolFormCallback *Part_symbolFormCallback) OnSave() {

	log.Println("Part_symbolFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	part_symbolFormCallback.probe.formStage.Checkout()

	if part_symbolFormCallback.part_symbol == nil {
		part_symbolFormCallback.part_symbol = new(models.Part_symbol).Stage(part_symbolFormCallback.probe.stageOfInterest)
	}
	part_symbol_ := part_symbolFormCallback.part_symbol
	_ = part_symbol_

	for _, formDiv := range part_symbolFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(part_symbol_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if part_symbolFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_symbol_.Unstage(part_symbolFormCallback.probe.stageOfInterest)
	}

	part_symbolFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Part_symbol](
		part_symbolFormCallback.probe,
	)
	part_symbolFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if part_symbolFormCallback.CreationMode || part_symbolFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_symbolFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(part_symbolFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Part_symbolFormCallback(
			nil,
			part_symbolFormCallback.probe,
			newFormGroup,
		)
		part_symbol := new(models.Part_symbol)
		FillUpForm(part_symbol, newFormGroup, part_symbolFormCallback.probe)
		part_symbolFormCallback.probe.formStage.Commit()
	}

	fillUpTree(part_symbolFormCallback.probe)
}
func __gong__New__Part_transposeFormCallback(
	part_transpose *models.Part_transpose,
	probe *Probe,
	formGroup *table.FormGroup,
) (part_transposeFormCallback *Part_transposeFormCallback) {
	part_transposeFormCallback = new(Part_transposeFormCallback)
	part_transposeFormCallback.probe = probe
	part_transposeFormCallback.part_transpose = part_transpose
	part_transposeFormCallback.formGroup = formGroup

	part_transposeFormCallback.CreationMode = (part_transpose == nil)

	return
}

type Part_transposeFormCallback struct {
	part_transpose *models.Part_transpose

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (part_transposeFormCallback *Part_transposeFormCallback) OnSave() {

	log.Println("Part_transposeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	part_transposeFormCallback.probe.formStage.Checkout()

	if part_transposeFormCallback.part_transpose == nil {
		part_transposeFormCallback.part_transpose = new(models.Part_transpose).Stage(part_transposeFormCallback.probe.stageOfInterest)
	}
	part_transpose_ := part_transposeFormCallback.part_transpose
	_ = part_transpose_

	for _, formDiv := range part_transposeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(part_transpose_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if part_transposeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_transpose_.Unstage(part_transposeFormCallback.probe.stageOfInterest)
	}

	part_transposeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Part_transpose](
		part_transposeFormCallback.probe,
	)
	part_transposeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if part_transposeFormCallback.CreationMode || part_transposeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_transposeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(part_transposeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Part_transposeFormCallback(
			nil,
			part_transposeFormCallback.probe,
			newFormGroup,
		)
		part_transpose := new(models.Part_transpose)
		FillUpForm(part_transpose, newFormGroup, part_transposeFormCallback.probe)
		part_transposeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(part_transposeFormCallback.probe)
}
func __gong__New__PedalFormCallback(
	pedal *models.Pedal,
	probe *Probe,
	formGroup *table.FormGroup,
) (pedalFormCallback *PedalFormCallback) {
	pedalFormCallback = new(PedalFormCallback)
	pedalFormCallback.probe = probe
	pedalFormCallback.pedal = pedal
	pedalFormCallback.formGroup = formGroup

	pedalFormCallback.CreationMode = (pedal == nil)

	return
}

type PedalFormCallback struct {
	pedal *models.Pedal

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (pedalFormCallback *PedalFormCallback) OnSave() {

	log.Println("PedalFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pedalFormCallback.probe.formStage.Checkout()

	if pedalFormCallback.pedal == nil {
		pedalFormCallback.pedal = new(models.Pedal).Stage(pedalFormCallback.probe.stageOfInterest)
	}
	pedal_ := pedalFormCallback.pedal
	_ = pedal_

	for _, formDiv := range pedalFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pedal_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if pedalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pedal_.Unstage(pedalFormCallback.probe.stageOfInterest)
	}

	pedalFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Pedal](
		pedalFormCallback.probe,
	)
	pedalFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pedalFormCallback.CreationMode || pedalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pedalFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(pedalFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PedalFormCallback(
			nil,
			pedalFormCallback.probe,
			newFormGroup,
		)
		pedal := new(models.Pedal)
		FillUpForm(pedal, newFormGroup, pedalFormCallback.probe)
		pedalFormCallback.probe.formStage.Commit()
	}

	fillUpTree(pedalFormCallback.probe)
}
func __gong__New__Pedal_tuningFormCallback(
	pedal_tuning *models.Pedal_tuning,
	probe *Probe,
	formGroup *table.FormGroup,
) (pedal_tuningFormCallback *Pedal_tuningFormCallback) {
	pedal_tuningFormCallback = new(Pedal_tuningFormCallback)
	pedal_tuningFormCallback.probe = probe
	pedal_tuningFormCallback.pedal_tuning = pedal_tuning
	pedal_tuningFormCallback.formGroup = formGroup

	pedal_tuningFormCallback.CreationMode = (pedal_tuning == nil)

	return
}

type Pedal_tuningFormCallback struct {
	pedal_tuning *models.Pedal_tuning

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (pedal_tuningFormCallback *Pedal_tuningFormCallback) OnSave() {

	log.Println("Pedal_tuningFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pedal_tuningFormCallback.probe.formStage.Checkout()

	if pedal_tuningFormCallback.pedal_tuning == nil {
		pedal_tuningFormCallback.pedal_tuning = new(models.Pedal_tuning).Stage(pedal_tuningFormCallback.probe.stageOfInterest)
	}
	pedal_tuning_ := pedal_tuningFormCallback.pedal_tuning
	_ = pedal_tuning_

	for _, formDiv := range pedal_tuningFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pedal_tuning_.Name), formDiv)
		case "Harp_pedals:Pedal_tuning":
			// we need to retrieve the field owner before the change
			var pastHarp_pedalsOwner *models.Harp_pedals
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Harp_pedals"
			rf.Fieldname = "Pedal_tuning"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				pedal_tuningFormCallback.probe.stageOfInterest,
				pedal_tuningFormCallback.probe.backRepoOfInterest,
				pedal_tuning_,
				&rf)

			if reverseFieldOwner != nil {
				pastHarp_pedalsOwner = reverseFieldOwner.(*models.Harp_pedals)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastHarp_pedalsOwner != nil {
					idx := slices.Index(pastHarp_pedalsOwner.Pedal_tuning, pedal_tuning_)
					pastHarp_pedalsOwner.Pedal_tuning = slices.Delete(pastHarp_pedalsOwner.Pedal_tuning, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _harp_pedals := range *models.GetGongstructInstancesSet[models.Harp_pedals](pedal_tuningFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _harp_pedals.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newHarp_pedalsOwner := _harp_pedals // we have a match
						if pastHarp_pedalsOwner != nil {
							if newHarp_pedalsOwner != pastHarp_pedalsOwner {
								idx := slices.Index(pastHarp_pedalsOwner.Pedal_tuning, pedal_tuning_)
								pastHarp_pedalsOwner.Pedal_tuning = slices.Delete(pastHarp_pedalsOwner.Pedal_tuning, idx, idx+1)
								newHarp_pedalsOwner.Pedal_tuning = append(newHarp_pedalsOwner.Pedal_tuning, pedal_tuning_)
							}
						} else {
							newHarp_pedalsOwner.Pedal_tuning = append(newHarp_pedalsOwner.Pedal_tuning, pedal_tuning_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if pedal_tuningFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pedal_tuning_.Unstage(pedal_tuningFormCallback.probe.stageOfInterest)
	}

	pedal_tuningFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Pedal_tuning](
		pedal_tuningFormCallback.probe,
	)
	pedal_tuningFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pedal_tuningFormCallback.CreationMode || pedal_tuningFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pedal_tuningFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(pedal_tuningFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Pedal_tuningFormCallback(
			nil,
			pedal_tuningFormCallback.probe,
			newFormGroup,
		)
		pedal_tuning := new(models.Pedal_tuning)
		FillUpForm(pedal_tuning, newFormGroup, pedal_tuningFormCallback.probe)
		pedal_tuningFormCallback.probe.formStage.Commit()
	}

	fillUpTree(pedal_tuningFormCallback.probe)
}
func __gong__New__PercussionFormCallback(
	percussion *models.Percussion,
	probe *Probe,
	formGroup *table.FormGroup,
) (percussionFormCallback *PercussionFormCallback) {
	percussionFormCallback = new(PercussionFormCallback)
	percussionFormCallback.probe = probe
	percussionFormCallback.percussion = percussion
	percussionFormCallback.formGroup = formGroup

	percussionFormCallback.CreationMode = (percussion == nil)

	return
}

type PercussionFormCallback struct {
	percussion *models.Percussion

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (percussionFormCallback *PercussionFormCallback) OnSave() {

	log.Println("PercussionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	percussionFormCallback.probe.formStage.Checkout()

	if percussionFormCallback.percussion == nil {
		percussionFormCallback.percussion = new(models.Percussion).Stage(percussionFormCallback.probe.stageOfInterest)
	}
	percussion_ := percussionFormCallback.percussion
	_ = percussion_

	for _, formDiv := range percussionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(percussion_.Name), formDiv)
		case "Glass":
			FormDivSelectFieldToField(&(percussion_.Glass), percussionFormCallback.probe.stageOfInterest, formDiv)
		case "Metal":
			FormDivSelectFieldToField(&(percussion_.Metal), percussionFormCallback.probe.stageOfInterest, formDiv)
		case "Wood":
			FormDivSelectFieldToField(&(percussion_.Wood), percussionFormCallback.probe.stageOfInterest, formDiv)
		case "Pitched":
			FormDivSelectFieldToField(&(percussion_.Pitched), percussionFormCallback.probe.stageOfInterest, formDiv)
		case "Membrane":
			FormDivSelectFieldToField(&(percussion_.Membrane), percussionFormCallback.probe.stageOfInterest, formDiv)
		case "Effect":
			FormDivSelectFieldToField(&(percussion_.Effect), percussionFormCallback.probe.stageOfInterest, formDiv)
		case "Timpani":
			FormDivSelectFieldToField(&(percussion_.Timpani), percussionFormCallback.probe.stageOfInterest, formDiv)
		case "Beater":
			FormDivSelectFieldToField(&(percussion_.Beater), percussionFormCallback.probe.stageOfInterest, formDiv)
		case "Stick":
			FormDivSelectFieldToField(&(percussion_.Stick), percussionFormCallback.probe.stageOfInterest, formDiv)
		case "Direction_type:Percussion":
			// we need to retrieve the field owner before the change
			var pastDirection_typeOwner *models.Direction_type
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Percussion"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				percussionFormCallback.probe.stageOfInterest,
				percussionFormCallback.probe.backRepoOfInterest,
				percussion_,
				&rf)

			if reverseFieldOwner != nil {
				pastDirection_typeOwner = reverseFieldOwner.(*models.Direction_type)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDirection_typeOwner != nil {
					idx := slices.Index(pastDirection_typeOwner.Percussion, percussion_)
					pastDirection_typeOwner.Percussion = slices.Delete(pastDirection_typeOwner.Percussion, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _direction_type := range *models.GetGongstructInstancesSet[models.Direction_type](percussionFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _direction_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDirection_typeOwner := _direction_type // we have a match
						if pastDirection_typeOwner != nil {
							if newDirection_typeOwner != pastDirection_typeOwner {
								idx := slices.Index(pastDirection_typeOwner.Percussion, percussion_)
								pastDirection_typeOwner.Percussion = slices.Delete(pastDirection_typeOwner.Percussion, idx, idx+1)
								newDirection_typeOwner.Percussion = append(newDirection_typeOwner.Percussion, percussion_)
							}
						} else {
							newDirection_typeOwner.Percussion = append(newDirection_typeOwner.Percussion, percussion_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if percussionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		percussion_.Unstage(percussionFormCallback.probe.stageOfInterest)
	}

	percussionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Percussion](
		percussionFormCallback.probe,
	)
	percussionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if percussionFormCallback.CreationMode || percussionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		percussionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(percussionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PercussionFormCallback(
			nil,
			percussionFormCallback.probe,
			newFormGroup,
		)
		percussion := new(models.Percussion)
		FillUpForm(percussion, newFormGroup, percussionFormCallback.probe)
		percussionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(percussionFormCallback.probe)
}
func __gong__New__PitchFormCallback(
	pitch *models.Pitch,
	probe *Probe,
	formGroup *table.FormGroup,
) (pitchFormCallback *PitchFormCallback) {
	pitchFormCallback = new(PitchFormCallback)
	pitchFormCallback.probe = probe
	pitchFormCallback.pitch = pitch
	pitchFormCallback.formGroup = formGroup

	pitchFormCallback.CreationMode = (pitch == nil)

	return
}

type PitchFormCallback struct {
	pitch *models.Pitch

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (pitchFormCallback *PitchFormCallback) OnSave() {

	log.Println("PitchFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pitchFormCallback.probe.formStage.Checkout()

	if pitchFormCallback.pitch == nil {
		pitchFormCallback.pitch = new(models.Pitch).Stage(pitchFormCallback.probe.stageOfInterest)
	}
	pitch_ := pitchFormCallback.pitch
	_ = pitch_

	for _, formDiv := range pitchFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pitch_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if pitchFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pitch_.Unstage(pitchFormCallback.probe.stageOfInterest)
	}

	pitchFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Pitch](
		pitchFormCallback.probe,
	)
	pitchFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pitchFormCallback.CreationMode || pitchFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pitchFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(pitchFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PitchFormCallback(
			nil,
			pitchFormCallback.probe,
			newFormGroup,
		)
		pitch := new(models.Pitch)
		FillUpForm(pitch, newFormGroup, pitchFormCallback.probe)
		pitchFormCallback.probe.formStage.Commit()
	}

	fillUpTree(pitchFormCallback.probe)
}
func __gong__New__PitchedFormCallback(
	pitched *models.Pitched,
	probe *Probe,
	formGroup *table.FormGroup,
) (pitchedFormCallback *PitchedFormCallback) {
	pitchedFormCallback = new(PitchedFormCallback)
	pitchedFormCallback.probe = probe
	pitchedFormCallback.pitched = pitched
	pitchedFormCallback.formGroup = formGroup

	pitchedFormCallback.CreationMode = (pitched == nil)

	return
}

type PitchedFormCallback struct {
	pitched *models.Pitched

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (pitchedFormCallback *PitchedFormCallback) OnSave() {

	log.Println("PitchedFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pitchedFormCallback.probe.formStage.Checkout()

	if pitchedFormCallback.pitched == nil {
		pitchedFormCallback.pitched = new(models.Pitched).Stage(pitchedFormCallback.probe.stageOfInterest)
	}
	pitched_ := pitchedFormCallback.pitched
	_ = pitched_

	for _, formDiv := range pitchedFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pitched_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if pitchedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pitched_.Unstage(pitchedFormCallback.probe.stageOfInterest)
	}

	pitchedFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Pitched](
		pitchedFormCallback.probe,
	)
	pitchedFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pitchedFormCallback.CreationMode || pitchedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pitchedFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(pitchedFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PitchedFormCallback(
			nil,
			pitchedFormCallback.probe,
			newFormGroup,
		)
		pitched := new(models.Pitched)
		FillUpForm(pitched, newFormGroup, pitchedFormCallback.probe)
		pitchedFormCallback.probe.formStage.Commit()
	}

	fillUpTree(pitchedFormCallback.probe)
}
func __gong__New__PlayFormCallback(
	play *models.Play,
	probe *Probe,
	formGroup *table.FormGroup,
) (playFormCallback *PlayFormCallback) {
	playFormCallback = new(PlayFormCallback)
	playFormCallback.probe = probe
	playFormCallback.play = play
	playFormCallback.formGroup = formGroup

	playFormCallback.CreationMode = (play == nil)

	return
}

type PlayFormCallback struct {
	play *models.Play

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (playFormCallback *PlayFormCallback) OnSave() {

	log.Println("PlayFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	playFormCallback.probe.formStage.Checkout()

	if playFormCallback.play == nil {
		playFormCallback.play = new(models.Play).Stage(playFormCallback.probe.stageOfInterest)
	}
	play_ := playFormCallback.play
	_ = play_

	for _, formDiv := range playFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(play_.Name), formDiv)
		case "Ipa":
			FormDivBasicFieldToField(&(play_.Ipa), formDiv)
		case "Other_play":
			FormDivSelectFieldToField(&(play_.Other_play), playFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if playFormCallback.formGroup.HasSuppressButtonBeenPressed {
		play_.Unstage(playFormCallback.probe.stageOfInterest)
	}

	playFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Play](
		playFormCallback.probe,
	)
	playFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if playFormCallback.CreationMode || playFormCallback.formGroup.HasSuppressButtonBeenPressed {
		playFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(playFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PlayFormCallback(
			nil,
			playFormCallback.probe,
			newFormGroup,
		)
		play := new(models.Play)
		FillUpForm(play, newFormGroup, playFormCallback.probe)
		playFormCallback.probe.formStage.Commit()
	}

	fillUpTree(playFormCallback.probe)
}
func __gong__New__PlayerFormCallback(
	player *models.Player,
	probe *Probe,
	formGroup *table.FormGroup,
) (playerFormCallback *PlayerFormCallback) {
	playerFormCallback = new(PlayerFormCallback)
	playerFormCallback.probe = probe
	playerFormCallback.player = player
	playerFormCallback.formGroup = formGroup

	playerFormCallback.CreationMode = (player == nil)

	return
}

type PlayerFormCallback struct {
	player *models.Player

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (playerFormCallback *PlayerFormCallback) OnSave() {

	log.Println("PlayerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	playerFormCallback.probe.formStage.Checkout()

	if playerFormCallback.player == nil {
		playerFormCallback.player = new(models.Player).Stage(playerFormCallback.probe.stageOfInterest)
	}
	player_ := playerFormCallback.player
	_ = player_

	for _, formDiv := range playerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(player_.Name), formDiv)
		case "Player_name":
			FormDivBasicFieldToField(&(player_.Player_name), formDiv)
		case "Score_part:Player":
			// we need to retrieve the field owner before the change
			var pastScore_partOwner *models.Score_part
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Player"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				playerFormCallback.probe.stageOfInterest,
				playerFormCallback.probe.backRepoOfInterest,
				player_,
				&rf)

			if reverseFieldOwner != nil {
				pastScore_partOwner = reverseFieldOwner.(*models.Score_part)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastScore_partOwner != nil {
					idx := slices.Index(pastScore_partOwner.Player, player_)
					pastScore_partOwner.Player = slices.Delete(pastScore_partOwner.Player, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _score_part := range *models.GetGongstructInstancesSet[models.Score_part](playerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _score_part.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newScore_partOwner := _score_part // we have a match
						if pastScore_partOwner != nil {
							if newScore_partOwner != pastScore_partOwner {
								idx := slices.Index(pastScore_partOwner.Player, player_)
								pastScore_partOwner.Player = slices.Delete(pastScore_partOwner.Player, idx, idx+1)
								newScore_partOwner.Player = append(newScore_partOwner.Player, player_)
							}
						} else {
							newScore_partOwner.Player = append(newScore_partOwner.Player, player_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if playerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		player_.Unstage(playerFormCallback.probe.stageOfInterest)
	}

	playerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Player](
		playerFormCallback.probe,
	)
	playerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if playerFormCallback.CreationMode || playerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		playerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(playerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PlayerFormCallback(
			nil,
			playerFormCallback.probe,
			newFormGroup,
		)
		player := new(models.Player)
		FillUpForm(player, newFormGroup, playerFormCallback.probe)
		playerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(playerFormCallback.probe)
}
func __gong__New__Principal_voiceFormCallback(
	principal_voice *models.Principal_voice,
	probe *Probe,
	formGroup *table.FormGroup,
) (principal_voiceFormCallback *Principal_voiceFormCallback) {
	principal_voiceFormCallback = new(Principal_voiceFormCallback)
	principal_voiceFormCallback.probe = probe
	principal_voiceFormCallback.principal_voice = principal_voice
	principal_voiceFormCallback.formGroup = formGroup

	principal_voiceFormCallback.CreationMode = (principal_voice == nil)

	return
}

type Principal_voiceFormCallback struct {
	principal_voice *models.Principal_voice

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (principal_voiceFormCallback *Principal_voiceFormCallback) OnSave() {

	log.Println("Principal_voiceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	principal_voiceFormCallback.probe.formStage.Checkout()

	if principal_voiceFormCallback.principal_voice == nil {
		principal_voiceFormCallback.principal_voice = new(models.Principal_voice).Stage(principal_voiceFormCallback.probe.stageOfInterest)
	}
	principal_voice_ := principal_voiceFormCallback.principal_voice
	_ = principal_voice_

	for _, formDiv := range principal_voiceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(principal_voice_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(principal_voice_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if principal_voiceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		principal_voice_.Unstage(principal_voiceFormCallback.probe.stageOfInterest)
	}

	principal_voiceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Principal_voice](
		principal_voiceFormCallback.probe,
	)
	principal_voiceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if principal_voiceFormCallback.CreationMode || principal_voiceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		principal_voiceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(principal_voiceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Principal_voiceFormCallback(
			nil,
			principal_voiceFormCallback.probe,
			newFormGroup,
		)
		principal_voice := new(models.Principal_voice)
		FillUpForm(principal_voice, newFormGroup, principal_voiceFormCallback.probe)
		principal_voiceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(principal_voiceFormCallback.probe)
}
func __gong__New__PrintFormCallback(
	print *models.Print,
	probe *Probe,
	formGroup *table.FormGroup,
) (printFormCallback *PrintFormCallback) {
	printFormCallback = new(PrintFormCallback)
	printFormCallback.probe = probe
	printFormCallback.print = print
	printFormCallback.formGroup = formGroup

	printFormCallback.CreationMode = (print == nil)

	return
}

type PrintFormCallback struct {
	print *models.Print

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (printFormCallback *PrintFormCallback) OnSave() {

	log.Println("PrintFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	printFormCallback.probe.formStage.Checkout()

	if printFormCallback.print == nil {
		printFormCallback.print = new(models.Print).Stage(printFormCallback.probe.stageOfInterest)
	}
	print_ := printFormCallback.print
	_ = print_

	for _, formDiv := range printFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(print_.Name), formDiv)
		case "Measure_layout":
			FormDivSelectFieldToField(&(print_.Measure_layout), printFormCallback.probe.stageOfInterest, formDiv)
		case "Measure_numbering":
			FormDivSelectFieldToField(&(print_.Measure_numbering), printFormCallback.probe.stageOfInterest, formDiv)
		case "Part_name_display":
			FormDivSelectFieldToField(&(print_.Part_name_display), printFormCallback.probe.stageOfInterest, formDiv)
		case "Part_abbreviation_display":
			FormDivSelectFieldToField(&(print_.Part_abbreviation_display), printFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if printFormCallback.formGroup.HasSuppressButtonBeenPressed {
		print_.Unstage(printFormCallback.probe.stageOfInterest)
	}

	printFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Print](
		printFormCallback.probe,
	)
	printFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if printFormCallback.CreationMode || printFormCallback.formGroup.HasSuppressButtonBeenPressed {
		printFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(printFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PrintFormCallback(
			nil,
			printFormCallback.probe,
			newFormGroup,
		)
		print := new(models.Print)
		FillUpForm(print, newFormGroup, printFormCallback.probe)
		printFormCallback.probe.formStage.Commit()
	}

	fillUpTree(printFormCallback.probe)
}
func __gong__New__ReleaseFormCallback(
	release *models.Release,
	probe *Probe,
	formGroup *table.FormGroup,
) (releaseFormCallback *ReleaseFormCallback) {
	releaseFormCallback = new(ReleaseFormCallback)
	releaseFormCallback.probe = probe
	releaseFormCallback.release = release
	releaseFormCallback.formGroup = formGroup

	releaseFormCallback.CreationMode = (release == nil)

	return
}

type ReleaseFormCallback struct {
	release *models.Release

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (releaseFormCallback *ReleaseFormCallback) OnSave() {

	log.Println("ReleaseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	releaseFormCallback.probe.formStage.Checkout()

	if releaseFormCallback.release == nil {
		releaseFormCallback.release = new(models.Release).Stage(releaseFormCallback.probe.stageOfInterest)
	}
	release_ := releaseFormCallback.release
	_ = release_

	for _, formDiv := range releaseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(release_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if releaseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		release_.Unstage(releaseFormCallback.probe.stageOfInterest)
	}

	releaseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Release](
		releaseFormCallback.probe,
	)
	releaseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if releaseFormCallback.CreationMode || releaseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		releaseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(releaseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ReleaseFormCallback(
			nil,
			releaseFormCallback.probe,
			newFormGroup,
		)
		release := new(models.Release)
		FillUpForm(release, newFormGroup, releaseFormCallback.probe)
		releaseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(releaseFormCallback.probe)
}
func __gong__New__RepeatFormCallback(
	repeat *models.Repeat,
	probe *Probe,
	formGroup *table.FormGroup,
) (repeatFormCallback *RepeatFormCallback) {
	repeatFormCallback = new(RepeatFormCallback)
	repeatFormCallback.probe = probe
	repeatFormCallback.repeat = repeat
	repeatFormCallback.formGroup = formGroup

	repeatFormCallback.CreationMode = (repeat == nil)

	return
}

type RepeatFormCallback struct {
	repeat *models.Repeat

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (repeatFormCallback *RepeatFormCallback) OnSave() {

	log.Println("RepeatFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	repeatFormCallback.probe.formStage.Checkout()

	if repeatFormCallback.repeat == nil {
		repeatFormCallback.repeat = new(models.Repeat).Stage(repeatFormCallback.probe.stageOfInterest)
	}
	repeat_ := repeatFormCallback.repeat
	_ = repeat_

	for _, formDiv := range repeatFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(repeat_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if repeatFormCallback.formGroup.HasSuppressButtonBeenPressed {
		repeat_.Unstage(repeatFormCallback.probe.stageOfInterest)
	}

	repeatFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Repeat](
		repeatFormCallback.probe,
	)
	repeatFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if repeatFormCallback.CreationMode || repeatFormCallback.formGroup.HasSuppressButtonBeenPressed {
		repeatFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(repeatFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RepeatFormCallback(
			nil,
			repeatFormCallback.probe,
			newFormGroup,
		)
		repeat := new(models.Repeat)
		FillUpForm(repeat, newFormGroup, repeatFormCallback.probe)
		repeatFormCallback.probe.formStage.Commit()
	}

	fillUpTree(repeatFormCallback.probe)
}
func __gong__New__RestFormCallback(
	rest *models.Rest,
	probe *Probe,
	formGroup *table.FormGroup,
) (restFormCallback *RestFormCallback) {
	restFormCallback = new(RestFormCallback)
	restFormCallback.probe = probe
	restFormCallback.rest = rest
	restFormCallback.formGroup = formGroup

	restFormCallback.CreationMode = (rest == nil)

	return
}

type RestFormCallback struct {
	rest *models.Rest

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (restFormCallback *RestFormCallback) OnSave() {

	log.Println("RestFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	restFormCallback.probe.formStage.Checkout()

	if restFormCallback.rest == nil {
		restFormCallback.rest = new(models.Rest).Stage(restFormCallback.probe.stageOfInterest)
	}
	rest_ := restFormCallback.rest
	_ = rest_

	for _, formDiv := range restFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rest_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if restFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rest_.Unstage(restFormCallback.probe.stageOfInterest)
	}

	restFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Rest](
		restFormCallback.probe,
	)
	restFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if restFormCallback.CreationMode || restFormCallback.formGroup.HasSuppressButtonBeenPressed {
		restFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(restFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RestFormCallback(
			nil,
			restFormCallback.probe,
			newFormGroup,
		)
		rest := new(models.Rest)
		FillUpForm(rest, newFormGroup, restFormCallback.probe)
		restFormCallback.probe.formStage.Commit()
	}

	fillUpTree(restFormCallback.probe)
}
func __gong__New__RootFormCallback(
	root *models.Root,
	probe *Probe,
	formGroup *table.FormGroup,
) (rootFormCallback *RootFormCallback) {
	rootFormCallback = new(RootFormCallback)
	rootFormCallback.probe = probe
	rootFormCallback.root = root
	rootFormCallback.formGroup = formGroup

	rootFormCallback.CreationMode = (root == nil)

	return
}

type RootFormCallback struct {
	root *models.Root

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rootFormCallback *RootFormCallback) OnSave() {

	log.Println("RootFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rootFormCallback.probe.formStage.Checkout()

	if rootFormCallback.root == nil {
		rootFormCallback.root = new(models.Root).Stage(rootFormCallback.probe.stageOfInterest)
	}
	root_ := rootFormCallback.root
	_ = root_

	for _, formDiv := range rootFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(root_.Name), formDiv)
		case "Root_step":
			FormDivSelectFieldToField(&(root_.Root_step), rootFormCallback.probe.stageOfInterest, formDiv)
		case "Root_alter":
			FormDivSelectFieldToField(&(root_.Root_alter), rootFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if rootFormCallback.formGroup.HasSuppressButtonBeenPressed {
		root_.Unstage(rootFormCallback.probe.stageOfInterest)
	}

	rootFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Root](
		rootFormCallback.probe,
	)
	rootFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rootFormCallback.CreationMode || rootFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rootFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(rootFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RootFormCallback(
			nil,
			rootFormCallback.probe,
			newFormGroup,
		)
		root := new(models.Root)
		FillUpForm(root, newFormGroup, rootFormCallback.probe)
		rootFormCallback.probe.formStage.Commit()
	}

	fillUpTree(rootFormCallback.probe)
}
func __gong__New__Root_stepFormCallback(
	root_step *models.Root_step,
	probe *Probe,
	formGroup *table.FormGroup,
) (root_stepFormCallback *Root_stepFormCallback) {
	root_stepFormCallback = new(Root_stepFormCallback)
	root_stepFormCallback.probe = probe
	root_stepFormCallback.root_step = root_step
	root_stepFormCallback.formGroup = formGroup

	root_stepFormCallback.CreationMode = (root_step == nil)

	return
}

type Root_stepFormCallback struct {
	root_step *models.Root_step

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (root_stepFormCallback *Root_stepFormCallback) OnSave() {

	log.Println("Root_stepFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	root_stepFormCallback.probe.formStage.Checkout()

	if root_stepFormCallback.root_step == nil {
		root_stepFormCallback.root_step = new(models.Root_step).Stage(root_stepFormCallback.probe.stageOfInterest)
	}
	root_step_ := root_stepFormCallback.root_step
	_ = root_step_

	for _, formDiv := range root_stepFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(root_step_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(root_step_.Text), formDiv)
		}
	}

	// manage the suppress operation
	if root_stepFormCallback.formGroup.HasSuppressButtonBeenPressed {
		root_step_.Unstage(root_stepFormCallback.probe.stageOfInterest)
	}

	root_stepFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Root_step](
		root_stepFormCallback.probe,
	)
	root_stepFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if root_stepFormCallback.CreationMode || root_stepFormCallback.formGroup.HasSuppressButtonBeenPressed {
		root_stepFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(root_stepFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Root_stepFormCallback(
			nil,
			root_stepFormCallback.probe,
			newFormGroup,
		)
		root_step := new(models.Root_step)
		FillUpForm(root_step, newFormGroup, root_stepFormCallback.probe)
		root_stepFormCallback.probe.formStage.Commit()
	}

	fillUpTree(root_stepFormCallback.probe)
}
func __gong__New__ScalingFormCallback(
	scaling *models.Scaling,
	probe *Probe,
	formGroup *table.FormGroup,
) (scalingFormCallback *ScalingFormCallback) {
	scalingFormCallback = new(ScalingFormCallback)
	scalingFormCallback.probe = probe
	scalingFormCallback.scaling = scaling
	scalingFormCallback.formGroup = formGroup

	scalingFormCallback.CreationMode = (scaling == nil)

	return
}

type ScalingFormCallback struct {
	scaling *models.Scaling

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (scalingFormCallback *ScalingFormCallback) OnSave() {

	log.Println("ScalingFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	scalingFormCallback.probe.formStage.Checkout()

	if scalingFormCallback.scaling == nil {
		scalingFormCallback.scaling = new(models.Scaling).Stage(scalingFormCallback.probe.stageOfInterest)
	}
	scaling_ := scalingFormCallback.scaling
	_ = scaling_

	for _, formDiv := range scalingFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(scaling_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if scalingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scaling_.Unstage(scalingFormCallback.probe.stageOfInterest)
	}

	scalingFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Scaling](
		scalingFormCallback.probe,
	)
	scalingFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if scalingFormCallback.CreationMode || scalingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scalingFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(scalingFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ScalingFormCallback(
			nil,
			scalingFormCallback.probe,
			newFormGroup,
		)
		scaling := new(models.Scaling)
		FillUpForm(scaling, newFormGroup, scalingFormCallback.probe)
		scalingFormCallback.probe.formStage.Commit()
	}

	fillUpTree(scalingFormCallback.probe)
}
func __gong__New__ScordaturaFormCallback(
	scordatura *models.Scordatura,
	probe *Probe,
	formGroup *table.FormGroup,
) (scordaturaFormCallback *ScordaturaFormCallback) {
	scordaturaFormCallback = new(ScordaturaFormCallback)
	scordaturaFormCallback.probe = probe
	scordaturaFormCallback.scordatura = scordatura
	scordaturaFormCallback.formGroup = formGroup

	scordaturaFormCallback.CreationMode = (scordatura == nil)

	return
}

type ScordaturaFormCallback struct {
	scordatura *models.Scordatura

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (scordaturaFormCallback *ScordaturaFormCallback) OnSave() {

	log.Println("ScordaturaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	scordaturaFormCallback.probe.formStage.Checkout()

	if scordaturaFormCallback.scordatura == nil {
		scordaturaFormCallback.scordatura = new(models.Scordatura).Stage(scordaturaFormCallback.probe.stageOfInterest)
	}
	scordatura_ := scordaturaFormCallback.scordatura
	_ = scordatura_

	for _, formDiv := range scordaturaFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(scordatura_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if scordaturaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scordatura_.Unstage(scordaturaFormCallback.probe.stageOfInterest)
	}

	scordaturaFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Scordatura](
		scordaturaFormCallback.probe,
	)
	scordaturaFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if scordaturaFormCallback.CreationMode || scordaturaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scordaturaFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(scordaturaFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ScordaturaFormCallback(
			nil,
			scordaturaFormCallback.probe,
			newFormGroup,
		)
		scordatura := new(models.Scordatura)
		FillUpForm(scordatura, newFormGroup, scordaturaFormCallback.probe)
		scordaturaFormCallback.probe.formStage.Commit()
	}

	fillUpTree(scordaturaFormCallback.probe)
}
func __gong__New__Score_instrumentFormCallback(
	score_instrument *models.Score_instrument,
	probe *Probe,
	formGroup *table.FormGroup,
) (score_instrumentFormCallback *Score_instrumentFormCallback) {
	score_instrumentFormCallback = new(Score_instrumentFormCallback)
	score_instrumentFormCallback.probe = probe
	score_instrumentFormCallback.score_instrument = score_instrument
	score_instrumentFormCallback.formGroup = formGroup

	score_instrumentFormCallback.CreationMode = (score_instrument == nil)

	return
}

type Score_instrumentFormCallback struct {
	score_instrument *models.Score_instrument

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (score_instrumentFormCallback *Score_instrumentFormCallback) OnSave() {

	log.Println("Score_instrumentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	score_instrumentFormCallback.probe.formStage.Checkout()

	if score_instrumentFormCallback.score_instrument == nil {
		score_instrumentFormCallback.score_instrument = new(models.Score_instrument).Stage(score_instrumentFormCallback.probe.stageOfInterest)
	}
	score_instrument_ := score_instrumentFormCallback.score_instrument
	_ = score_instrument_

	for _, formDiv := range score_instrumentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(score_instrument_.Name), formDiv)
		case "Instrument_name":
			FormDivBasicFieldToField(&(score_instrument_.Instrument_name), formDiv)
		case "Instrument_abbreviation":
			FormDivBasicFieldToField(&(score_instrument_.Instrument_abbreviation), formDiv)
		case "Score_part:Score_instrument":
			// we need to retrieve the field owner before the change
			var pastScore_partOwner *models.Score_part
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Score_instrument"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				score_instrumentFormCallback.probe.stageOfInterest,
				score_instrumentFormCallback.probe.backRepoOfInterest,
				score_instrument_,
				&rf)

			if reverseFieldOwner != nil {
				pastScore_partOwner = reverseFieldOwner.(*models.Score_part)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastScore_partOwner != nil {
					idx := slices.Index(pastScore_partOwner.Score_instrument, score_instrument_)
					pastScore_partOwner.Score_instrument = slices.Delete(pastScore_partOwner.Score_instrument, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _score_part := range *models.GetGongstructInstancesSet[models.Score_part](score_instrumentFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _score_part.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newScore_partOwner := _score_part // we have a match
						if pastScore_partOwner != nil {
							if newScore_partOwner != pastScore_partOwner {
								idx := slices.Index(pastScore_partOwner.Score_instrument, score_instrument_)
								pastScore_partOwner.Score_instrument = slices.Delete(pastScore_partOwner.Score_instrument, idx, idx+1)
								newScore_partOwner.Score_instrument = append(newScore_partOwner.Score_instrument, score_instrument_)
							}
						} else {
							newScore_partOwner.Score_instrument = append(newScore_partOwner.Score_instrument, score_instrument_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if score_instrumentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		score_instrument_.Unstage(score_instrumentFormCallback.probe.stageOfInterest)
	}

	score_instrumentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Score_instrument](
		score_instrumentFormCallback.probe,
	)
	score_instrumentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if score_instrumentFormCallback.CreationMode || score_instrumentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		score_instrumentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(score_instrumentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Score_instrumentFormCallback(
			nil,
			score_instrumentFormCallback.probe,
			newFormGroup,
		)
		score_instrument := new(models.Score_instrument)
		FillUpForm(score_instrument, newFormGroup, score_instrumentFormCallback.probe)
		score_instrumentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(score_instrumentFormCallback.probe)
}
func __gong__New__Score_partFormCallback(
	score_part *models.Score_part,
	probe *Probe,
	formGroup *table.FormGroup,
) (score_partFormCallback *Score_partFormCallback) {
	score_partFormCallback = new(Score_partFormCallback)
	score_partFormCallback.probe = probe
	score_partFormCallback.score_part = score_part
	score_partFormCallback.formGroup = formGroup

	score_partFormCallback.CreationMode = (score_part == nil)

	return
}

type Score_partFormCallback struct {
	score_part *models.Score_part

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (score_partFormCallback *Score_partFormCallback) OnSave() {

	log.Println("Score_partFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	score_partFormCallback.probe.formStage.Checkout()

	if score_partFormCallback.score_part == nil {
		score_partFormCallback.score_part = new(models.Score_part).Stage(score_partFormCallback.probe.stageOfInterest)
	}
	score_part_ := score_partFormCallback.score_part
	_ = score_part_

	for _, formDiv := range score_partFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(score_part_.Name), formDiv)
		case "Identification":
			FormDivSelectFieldToField(&(score_part_.Identification), score_partFormCallback.probe.stageOfInterest, formDiv)
		case "Part_name_display":
			FormDivSelectFieldToField(&(score_part_.Part_name_display), score_partFormCallback.probe.stageOfInterest, formDiv)
		case "Part_abbreviation_display":
			FormDivSelectFieldToField(&(score_part_.Part_abbreviation_display), score_partFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if score_partFormCallback.formGroup.HasSuppressButtonBeenPressed {
		score_part_.Unstage(score_partFormCallback.probe.stageOfInterest)
	}

	score_partFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Score_part](
		score_partFormCallback.probe,
	)
	score_partFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if score_partFormCallback.CreationMode || score_partFormCallback.formGroup.HasSuppressButtonBeenPressed {
		score_partFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(score_partFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Score_partFormCallback(
			nil,
			score_partFormCallback.probe,
			newFormGroup,
		)
		score_part := new(models.Score_part)
		FillUpForm(score_part, newFormGroup, score_partFormCallback.probe)
		score_partFormCallback.probe.formStage.Commit()
	}

	fillUpTree(score_partFormCallback.probe)
}
func __gong__New__Score_partwiseFormCallback(
	score_partwise *models.Score_partwise,
	probe *Probe,
	formGroup *table.FormGroup,
) (score_partwiseFormCallback *Score_partwiseFormCallback) {
	score_partwiseFormCallback = new(Score_partwiseFormCallback)
	score_partwiseFormCallback.probe = probe
	score_partwiseFormCallback.score_partwise = score_partwise
	score_partwiseFormCallback.formGroup = formGroup

	score_partwiseFormCallback.CreationMode = (score_partwise == nil)

	return
}

type Score_partwiseFormCallback struct {
	score_partwise *models.Score_partwise

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (score_partwiseFormCallback *Score_partwiseFormCallback) OnSave() {

	log.Println("Score_partwiseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	score_partwiseFormCallback.probe.formStage.Checkout()

	if score_partwiseFormCallback.score_partwise == nil {
		score_partwiseFormCallback.score_partwise = new(models.Score_partwise).Stage(score_partwiseFormCallback.probe.stageOfInterest)
	}
	score_partwise_ := score_partwiseFormCallback.score_partwise
	_ = score_partwise_

	for _, formDiv := range score_partwiseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(score_partwise_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if score_partwiseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		score_partwise_.Unstage(score_partwiseFormCallback.probe.stageOfInterest)
	}

	score_partwiseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Score_partwise](
		score_partwiseFormCallback.probe,
	)
	score_partwiseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if score_partwiseFormCallback.CreationMode || score_partwiseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		score_partwiseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(score_partwiseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Score_partwiseFormCallback(
			nil,
			score_partwiseFormCallback.probe,
			newFormGroup,
		)
		score_partwise := new(models.Score_partwise)
		FillUpForm(score_partwise, newFormGroup, score_partwiseFormCallback.probe)
		score_partwiseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(score_partwiseFormCallback.probe)
}
func __gong__New__Score_timewiseFormCallback(
	score_timewise *models.Score_timewise,
	probe *Probe,
	formGroup *table.FormGroup,
) (score_timewiseFormCallback *Score_timewiseFormCallback) {
	score_timewiseFormCallback = new(Score_timewiseFormCallback)
	score_timewiseFormCallback.probe = probe
	score_timewiseFormCallback.score_timewise = score_timewise
	score_timewiseFormCallback.formGroup = formGroup

	score_timewiseFormCallback.CreationMode = (score_timewise == nil)

	return
}

type Score_timewiseFormCallback struct {
	score_timewise *models.Score_timewise

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (score_timewiseFormCallback *Score_timewiseFormCallback) OnSave() {

	log.Println("Score_timewiseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	score_timewiseFormCallback.probe.formStage.Checkout()

	if score_timewiseFormCallback.score_timewise == nil {
		score_timewiseFormCallback.score_timewise = new(models.Score_timewise).Stage(score_timewiseFormCallback.probe.stageOfInterest)
	}
	score_timewise_ := score_timewiseFormCallback.score_timewise
	_ = score_timewise_

	for _, formDiv := range score_timewiseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(score_timewise_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if score_timewiseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		score_timewise_.Unstage(score_timewiseFormCallback.probe.stageOfInterest)
	}

	score_timewiseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Score_timewise](
		score_timewiseFormCallback.probe,
	)
	score_timewiseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if score_timewiseFormCallback.CreationMode || score_timewiseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		score_timewiseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(score_timewiseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Score_timewiseFormCallback(
			nil,
			score_timewiseFormCallback.probe,
			newFormGroup,
		)
		score_timewise := new(models.Score_timewise)
		FillUpForm(score_timewise, newFormGroup, score_timewiseFormCallback.probe)
		score_timewiseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(score_timewiseFormCallback.probe)
}
func __gong__New__SegnoFormCallback(
	segno *models.Segno,
	probe *Probe,
	formGroup *table.FormGroup,
) (segnoFormCallback *SegnoFormCallback) {
	segnoFormCallback = new(SegnoFormCallback)
	segnoFormCallback.probe = probe
	segnoFormCallback.segno = segno
	segnoFormCallback.formGroup = formGroup

	segnoFormCallback.CreationMode = (segno == nil)

	return
}

type SegnoFormCallback struct {
	segno *models.Segno

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (segnoFormCallback *SegnoFormCallback) OnSave() {

	log.Println("SegnoFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	segnoFormCallback.probe.formStage.Checkout()

	if segnoFormCallback.segno == nil {
		segnoFormCallback.segno = new(models.Segno).Stage(segnoFormCallback.probe.stageOfInterest)
	}
	segno_ := segnoFormCallback.segno
	_ = segno_

	for _, formDiv := range segnoFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(segno_.Name), formDiv)
		case "Direction_type:Segno":
			// we need to retrieve the field owner before the change
			var pastDirection_typeOwner *models.Direction_type
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Segno"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				segnoFormCallback.probe.stageOfInterest,
				segnoFormCallback.probe.backRepoOfInterest,
				segno_,
				&rf)

			if reverseFieldOwner != nil {
				pastDirection_typeOwner = reverseFieldOwner.(*models.Direction_type)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDirection_typeOwner != nil {
					idx := slices.Index(pastDirection_typeOwner.Segno, segno_)
					pastDirection_typeOwner.Segno = slices.Delete(pastDirection_typeOwner.Segno, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _direction_type := range *models.GetGongstructInstancesSet[models.Direction_type](segnoFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _direction_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDirection_typeOwner := _direction_type // we have a match
						if pastDirection_typeOwner != nil {
							if newDirection_typeOwner != pastDirection_typeOwner {
								idx := slices.Index(pastDirection_typeOwner.Segno, segno_)
								pastDirection_typeOwner.Segno = slices.Delete(pastDirection_typeOwner.Segno, idx, idx+1)
								newDirection_typeOwner.Segno = append(newDirection_typeOwner.Segno, segno_)
							}
						} else {
							newDirection_typeOwner.Segno = append(newDirection_typeOwner.Segno, segno_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if segnoFormCallback.formGroup.HasSuppressButtonBeenPressed {
		segno_.Unstage(segnoFormCallback.probe.stageOfInterest)
	}

	segnoFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Segno](
		segnoFormCallback.probe,
	)
	segnoFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if segnoFormCallback.CreationMode || segnoFormCallback.formGroup.HasSuppressButtonBeenPressed {
		segnoFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(segnoFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SegnoFormCallback(
			nil,
			segnoFormCallback.probe,
			newFormGroup,
		)
		segno := new(models.Segno)
		FillUpForm(segno, newFormGroup, segnoFormCallback.probe)
		segnoFormCallback.probe.formStage.Commit()
	}

	fillUpTree(segnoFormCallback.probe)
}
func __gong__New__SlashFormCallback(
	slash *models.Slash,
	probe *Probe,
	formGroup *table.FormGroup,
) (slashFormCallback *SlashFormCallback) {
	slashFormCallback = new(SlashFormCallback)
	slashFormCallback.probe = probe
	slashFormCallback.slash = slash
	slashFormCallback.formGroup = formGroup

	slashFormCallback.CreationMode = (slash == nil)

	return
}

type SlashFormCallback struct {
	slash *models.Slash

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (slashFormCallback *SlashFormCallback) OnSave() {

	log.Println("SlashFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	slashFormCallback.probe.formStage.Checkout()

	if slashFormCallback.slash == nil {
		slashFormCallback.slash = new(models.Slash).Stage(slashFormCallback.probe.stageOfInterest)
	}
	slash_ := slashFormCallback.slash
	_ = slash_

	for _, formDiv := range slashFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(slash_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if slashFormCallback.formGroup.HasSuppressButtonBeenPressed {
		slash_.Unstage(slashFormCallback.probe.stageOfInterest)
	}

	slashFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Slash](
		slashFormCallback.probe,
	)
	slashFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if slashFormCallback.CreationMode || slashFormCallback.formGroup.HasSuppressButtonBeenPressed {
		slashFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(slashFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SlashFormCallback(
			nil,
			slashFormCallback.probe,
			newFormGroup,
		)
		slash := new(models.Slash)
		FillUpForm(slash, newFormGroup, slashFormCallback.probe)
		slashFormCallback.probe.formStage.Commit()
	}

	fillUpTree(slashFormCallback.probe)
}
func __gong__New__SlideFormCallback(
	slide *models.Slide,
	probe *Probe,
	formGroup *table.FormGroup,
) (slideFormCallback *SlideFormCallback) {
	slideFormCallback = new(SlideFormCallback)
	slideFormCallback.probe = probe
	slideFormCallback.slide = slide
	slideFormCallback.formGroup = formGroup

	slideFormCallback.CreationMode = (slide == nil)

	return
}

type SlideFormCallback struct {
	slide *models.Slide

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (slideFormCallback *SlideFormCallback) OnSave() {

	log.Println("SlideFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	slideFormCallback.probe.formStage.Checkout()

	if slideFormCallback.slide == nil {
		slideFormCallback.slide = new(models.Slide).Stage(slideFormCallback.probe.stageOfInterest)
	}
	slide_ := slideFormCallback.slide
	_ = slide_

	for _, formDiv := range slideFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(slide_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(slide_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if slideFormCallback.formGroup.HasSuppressButtonBeenPressed {
		slide_.Unstage(slideFormCallback.probe.stageOfInterest)
	}

	slideFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Slide](
		slideFormCallback.probe,
	)
	slideFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if slideFormCallback.CreationMode || slideFormCallback.formGroup.HasSuppressButtonBeenPressed {
		slideFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(slideFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SlideFormCallback(
			nil,
			slideFormCallback.probe,
			newFormGroup,
		)
		slide := new(models.Slide)
		FillUpForm(slide, newFormGroup, slideFormCallback.probe)
		slideFormCallback.probe.formStage.Commit()
	}

	fillUpTree(slideFormCallback.probe)
}
func __gong__New__SlurFormCallback(
	slur *models.Slur,
	probe *Probe,
	formGroup *table.FormGroup,
) (slurFormCallback *SlurFormCallback) {
	slurFormCallback = new(SlurFormCallback)
	slurFormCallback.probe = probe
	slurFormCallback.slur = slur
	slurFormCallback.formGroup = formGroup

	slurFormCallback.CreationMode = (slur == nil)

	return
}

type SlurFormCallback struct {
	slur *models.Slur

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (slurFormCallback *SlurFormCallback) OnSave() {

	log.Println("SlurFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	slurFormCallback.probe.formStage.Checkout()

	if slurFormCallback.slur == nil {
		slurFormCallback.slur = new(models.Slur).Stage(slurFormCallback.probe.stageOfInterest)
	}
	slur_ := slurFormCallback.slur
	_ = slur_

	for _, formDiv := range slurFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(slur_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if slurFormCallback.formGroup.HasSuppressButtonBeenPressed {
		slur_.Unstage(slurFormCallback.probe.stageOfInterest)
	}

	slurFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Slur](
		slurFormCallback.probe,
	)
	slurFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if slurFormCallback.CreationMode || slurFormCallback.formGroup.HasSuppressButtonBeenPressed {
		slurFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(slurFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SlurFormCallback(
			nil,
			slurFormCallback.probe,
			newFormGroup,
		)
		slur := new(models.Slur)
		FillUpForm(slur, newFormGroup, slurFormCallback.probe)
		slurFormCallback.probe.formStage.Commit()
	}

	fillUpTree(slurFormCallback.probe)
}
func __gong__New__SoundFormCallback(
	sound *models.Sound,
	probe *Probe,
	formGroup *table.FormGroup,
) (soundFormCallback *SoundFormCallback) {
	soundFormCallback = new(SoundFormCallback)
	soundFormCallback.probe = probe
	soundFormCallback.sound = sound
	soundFormCallback.formGroup = formGroup

	soundFormCallback.CreationMode = (sound == nil)

	return
}

type SoundFormCallback struct {
	sound *models.Sound

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (soundFormCallback *SoundFormCallback) OnSave() {

	log.Println("SoundFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	soundFormCallback.probe.formStage.Checkout()

	if soundFormCallback.sound == nil {
		soundFormCallback.sound = new(models.Sound).Stage(soundFormCallback.probe.stageOfInterest)
	}
	sound_ := soundFormCallback.sound
	_ = sound_

	for _, formDiv := range soundFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sound_.Name), formDiv)
		case "Swing":
			FormDivSelectFieldToField(&(sound_.Swing), soundFormCallback.probe.stageOfInterest, formDiv)
		case "Offset":
			FormDivSelectFieldToField(&(sound_.Offset), soundFormCallback.probe.stageOfInterest, formDiv)
		case "Segno":
			FormDivBasicFieldToField(&(sound_.Segno), formDiv)
		case "Dalsegno":
			FormDivBasicFieldToField(&(sound_.Dalsegno), formDiv)
		case "Coda":
			FormDivBasicFieldToField(&(sound_.Coda), formDiv)
		case "Tocoda":
			FormDivBasicFieldToField(&(sound_.Tocoda), formDiv)
		case "Fine":
			FormDivBasicFieldToField(&(sound_.Fine), formDiv)
		}
	}

	// manage the suppress operation
	if soundFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sound_.Unstage(soundFormCallback.probe.stageOfInterest)
	}

	soundFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Sound](
		soundFormCallback.probe,
	)
	soundFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if soundFormCallback.CreationMode || soundFormCallback.formGroup.HasSuppressButtonBeenPressed {
		soundFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(soundFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SoundFormCallback(
			nil,
			soundFormCallback.probe,
			newFormGroup,
		)
		sound := new(models.Sound)
		FillUpForm(sound, newFormGroup, soundFormCallback.probe)
		soundFormCallback.probe.formStage.Commit()
	}

	fillUpTree(soundFormCallback.probe)
}
func __gong__New__Staff_detailsFormCallback(
	staff_details *models.Staff_details,
	probe *Probe,
	formGroup *table.FormGroup,
) (staff_detailsFormCallback *Staff_detailsFormCallback) {
	staff_detailsFormCallback = new(Staff_detailsFormCallback)
	staff_detailsFormCallback.probe = probe
	staff_detailsFormCallback.staff_details = staff_details
	staff_detailsFormCallback.formGroup = formGroup

	staff_detailsFormCallback.CreationMode = (staff_details == nil)

	return
}

type Staff_detailsFormCallback struct {
	staff_details *models.Staff_details

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (staff_detailsFormCallback *Staff_detailsFormCallback) OnSave() {

	log.Println("Staff_detailsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	staff_detailsFormCallback.probe.formStage.Checkout()

	if staff_detailsFormCallback.staff_details == nil {
		staff_detailsFormCallback.staff_details = new(models.Staff_details).Stage(staff_detailsFormCallback.probe.stageOfInterest)
	}
	staff_details_ := staff_detailsFormCallback.staff_details
	_ = staff_details_

	for _, formDiv := range staff_detailsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(staff_details_.Name), formDiv)
		case "Staff_size":
			FormDivSelectFieldToField(&(staff_details_.Staff_size), staff_detailsFormCallback.probe.stageOfInterest, formDiv)
		case "Attributes:Staff_details":
			// we need to retrieve the field owner before the change
			var pastAttributesOwner *models.Attributes
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Staff_details"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				staff_detailsFormCallback.probe.stageOfInterest,
				staff_detailsFormCallback.probe.backRepoOfInterest,
				staff_details_,
				&rf)

			if reverseFieldOwner != nil {
				pastAttributesOwner = reverseFieldOwner.(*models.Attributes)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAttributesOwner != nil {
					idx := slices.Index(pastAttributesOwner.Staff_details, staff_details_)
					pastAttributesOwner.Staff_details = slices.Delete(pastAttributesOwner.Staff_details, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attributes := range *models.GetGongstructInstancesSet[models.Attributes](staff_detailsFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAttributesOwner := _attributes // we have a match
						if pastAttributesOwner != nil {
							if newAttributesOwner != pastAttributesOwner {
								idx := slices.Index(pastAttributesOwner.Staff_details, staff_details_)
								pastAttributesOwner.Staff_details = slices.Delete(pastAttributesOwner.Staff_details, idx, idx+1)
								newAttributesOwner.Staff_details = append(newAttributesOwner.Staff_details, staff_details_)
							}
						} else {
							newAttributesOwner.Staff_details = append(newAttributesOwner.Staff_details, staff_details_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if staff_detailsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_details_.Unstage(staff_detailsFormCallback.probe.stageOfInterest)
	}

	staff_detailsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Staff_details](
		staff_detailsFormCallback.probe,
	)
	staff_detailsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if staff_detailsFormCallback.CreationMode || staff_detailsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_detailsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(staff_detailsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Staff_detailsFormCallback(
			nil,
			staff_detailsFormCallback.probe,
			newFormGroup,
		)
		staff_details := new(models.Staff_details)
		FillUpForm(staff_details, newFormGroup, staff_detailsFormCallback.probe)
		staff_detailsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(staff_detailsFormCallback.probe)
}
func __gong__New__Staff_divideFormCallback(
	staff_divide *models.Staff_divide,
	probe *Probe,
	formGroup *table.FormGroup,
) (staff_divideFormCallback *Staff_divideFormCallback) {
	staff_divideFormCallback = new(Staff_divideFormCallback)
	staff_divideFormCallback.probe = probe
	staff_divideFormCallback.staff_divide = staff_divide
	staff_divideFormCallback.formGroup = formGroup

	staff_divideFormCallback.CreationMode = (staff_divide == nil)

	return
}

type Staff_divideFormCallback struct {
	staff_divide *models.Staff_divide

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (staff_divideFormCallback *Staff_divideFormCallback) OnSave() {

	log.Println("Staff_divideFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	staff_divideFormCallback.probe.formStage.Checkout()

	if staff_divideFormCallback.staff_divide == nil {
		staff_divideFormCallback.staff_divide = new(models.Staff_divide).Stage(staff_divideFormCallback.probe.stageOfInterest)
	}
	staff_divide_ := staff_divideFormCallback.staff_divide
	_ = staff_divide_

	for _, formDiv := range staff_divideFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(staff_divide_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if staff_divideFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_divide_.Unstage(staff_divideFormCallback.probe.stageOfInterest)
	}

	staff_divideFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Staff_divide](
		staff_divideFormCallback.probe,
	)
	staff_divideFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if staff_divideFormCallback.CreationMode || staff_divideFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_divideFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(staff_divideFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Staff_divideFormCallback(
			nil,
			staff_divideFormCallback.probe,
			newFormGroup,
		)
		staff_divide := new(models.Staff_divide)
		FillUpForm(staff_divide, newFormGroup, staff_divideFormCallback.probe)
		staff_divideFormCallback.probe.formStage.Commit()
	}

	fillUpTree(staff_divideFormCallback.probe)
}
func __gong__New__Staff_layoutFormCallback(
	staff_layout *models.Staff_layout,
	probe *Probe,
	formGroup *table.FormGroup,
) (staff_layoutFormCallback *Staff_layoutFormCallback) {
	staff_layoutFormCallback = new(Staff_layoutFormCallback)
	staff_layoutFormCallback.probe = probe
	staff_layoutFormCallback.staff_layout = staff_layout
	staff_layoutFormCallback.formGroup = formGroup

	staff_layoutFormCallback.CreationMode = (staff_layout == nil)

	return
}

type Staff_layoutFormCallback struct {
	staff_layout *models.Staff_layout

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (staff_layoutFormCallback *Staff_layoutFormCallback) OnSave() {

	log.Println("Staff_layoutFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	staff_layoutFormCallback.probe.formStage.Checkout()

	if staff_layoutFormCallback.staff_layout == nil {
		staff_layoutFormCallback.staff_layout = new(models.Staff_layout).Stage(staff_layoutFormCallback.probe.stageOfInterest)
	}
	staff_layout_ := staff_layoutFormCallback.staff_layout
	_ = staff_layout_

	for _, formDiv := range staff_layoutFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(staff_layout_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if staff_layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_layout_.Unstage(staff_layoutFormCallback.probe.stageOfInterest)
	}

	staff_layoutFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Staff_layout](
		staff_layoutFormCallback.probe,
	)
	staff_layoutFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if staff_layoutFormCallback.CreationMode || staff_layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_layoutFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(staff_layoutFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Staff_layoutFormCallback(
			nil,
			staff_layoutFormCallback.probe,
			newFormGroup,
		)
		staff_layout := new(models.Staff_layout)
		FillUpForm(staff_layout, newFormGroup, staff_layoutFormCallback.probe)
		staff_layoutFormCallback.probe.formStage.Commit()
	}

	fillUpTree(staff_layoutFormCallback.probe)
}
func __gong__New__Staff_sizeFormCallback(
	staff_size *models.Staff_size,
	probe *Probe,
	formGroup *table.FormGroup,
) (staff_sizeFormCallback *Staff_sizeFormCallback) {
	staff_sizeFormCallback = new(Staff_sizeFormCallback)
	staff_sizeFormCallback.probe = probe
	staff_sizeFormCallback.staff_size = staff_size
	staff_sizeFormCallback.formGroup = formGroup

	staff_sizeFormCallback.CreationMode = (staff_size == nil)

	return
}

type Staff_sizeFormCallback struct {
	staff_size *models.Staff_size

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (staff_sizeFormCallback *Staff_sizeFormCallback) OnSave() {

	log.Println("Staff_sizeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	staff_sizeFormCallback.probe.formStage.Checkout()

	if staff_sizeFormCallback.staff_size == nil {
		staff_sizeFormCallback.staff_size = new(models.Staff_size).Stage(staff_sizeFormCallback.probe.stageOfInterest)
	}
	staff_size_ := staff_sizeFormCallback.staff_size
	_ = staff_size_

	for _, formDiv := range staff_sizeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(staff_size_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if staff_sizeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_size_.Unstage(staff_sizeFormCallback.probe.stageOfInterest)
	}

	staff_sizeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Staff_size](
		staff_sizeFormCallback.probe,
	)
	staff_sizeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if staff_sizeFormCallback.CreationMode || staff_sizeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_sizeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(staff_sizeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Staff_sizeFormCallback(
			nil,
			staff_sizeFormCallback.probe,
			newFormGroup,
		)
		staff_size := new(models.Staff_size)
		FillUpForm(staff_size, newFormGroup, staff_sizeFormCallback.probe)
		staff_sizeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(staff_sizeFormCallback.probe)
}
func __gong__New__Staff_tuningFormCallback(
	staff_tuning *models.Staff_tuning,
	probe *Probe,
	formGroup *table.FormGroup,
) (staff_tuningFormCallback *Staff_tuningFormCallback) {
	staff_tuningFormCallback = new(Staff_tuningFormCallback)
	staff_tuningFormCallback.probe = probe
	staff_tuningFormCallback.staff_tuning = staff_tuning
	staff_tuningFormCallback.formGroup = formGroup

	staff_tuningFormCallback.CreationMode = (staff_tuning == nil)

	return
}

type Staff_tuningFormCallback struct {
	staff_tuning *models.Staff_tuning

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (staff_tuningFormCallback *Staff_tuningFormCallback) OnSave() {

	log.Println("Staff_tuningFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	staff_tuningFormCallback.probe.formStage.Checkout()

	if staff_tuningFormCallback.staff_tuning == nil {
		staff_tuningFormCallback.staff_tuning = new(models.Staff_tuning).Stage(staff_tuningFormCallback.probe.stageOfInterest)
	}
	staff_tuning_ := staff_tuningFormCallback.staff_tuning
	_ = staff_tuning_

	for _, formDiv := range staff_tuningFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(staff_tuning_.Name), formDiv)
		case "Staff_details:Staff_tuning":
			// we need to retrieve the field owner before the change
			var pastStaff_detailsOwner *models.Staff_details
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Staff_details"
			rf.Fieldname = "Staff_tuning"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				staff_tuningFormCallback.probe.stageOfInterest,
				staff_tuningFormCallback.probe.backRepoOfInterest,
				staff_tuning_,
				&rf)

			if reverseFieldOwner != nil {
				pastStaff_detailsOwner = reverseFieldOwner.(*models.Staff_details)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastStaff_detailsOwner != nil {
					idx := slices.Index(pastStaff_detailsOwner.Staff_tuning, staff_tuning_)
					pastStaff_detailsOwner.Staff_tuning = slices.Delete(pastStaff_detailsOwner.Staff_tuning, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _staff_details := range *models.GetGongstructInstancesSet[models.Staff_details](staff_tuningFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _staff_details.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newStaff_detailsOwner := _staff_details // we have a match
						if pastStaff_detailsOwner != nil {
							if newStaff_detailsOwner != pastStaff_detailsOwner {
								idx := slices.Index(pastStaff_detailsOwner.Staff_tuning, staff_tuning_)
								pastStaff_detailsOwner.Staff_tuning = slices.Delete(pastStaff_detailsOwner.Staff_tuning, idx, idx+1)
								newStaff_detailsOwner.Staff_tuning = append(newStaff_detailsOwner.Staff_tuning, staff_tuning_)
							}
						} else {
							newStaff_detailsOwner.Staff_tuning = append(newStaff_detailsOwner.Staff_tuning, staff_tuning_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if staff_tuningFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_tuning_.Unstage(staff_tuningFormCallback.probe.stageOfInterest)
	}

	staff_tuningFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Staff_tuning](
		staff_tuningFormCallback.probe,
	)
	staff_tuningFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if staff_tuningFormCallback.CreationMode || staff_tuningFormCallback.formGroup.HasSuppressButtonBeenPressed {
		staff_tuningFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(staff_tuningFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Staff_tuningFormCallback(
			nil,
			staff_tuningFormCallback.probe,
			newFormGroup,
		)
		staff_tuning := new(models.Staff_tuning)
		FillUpForm(staff_tuning, newFormGroup, staff_tuningFormCallback.probe)
		staff_tuningFormCallback.probe.formStage.Commit()
	}

	fillUpTree(staff_tuningFormCallback.probe)
}
func __gong__New__StemFormCallback(
	stem *models.Stem,
	probe *Probe,
	formGroup *table.FormGroup,
) (stemFormCallback *StemFormCallback) {
	stemFormCallback = new(StemFormCallback)
	stemFormCallback.probe = probe
	stemFormCallback.stem = stem
	stemFormCallback.formGroup = formGroup

	stemFormCallback.CreationMode = (stem == nil)

	return
}

type StemFormCallback struct {
	stem *models.Stem

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (stemFormCallback *StemFormCallback) OnSave() {

	log.Println("StemFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stemFormCallback.probe.formStage.Checkout()

	if stemFormCallback.stem == nil {
		stemFormCallback.stem = new(models.Stem).Stage(stemFormCallback.probe.stageOfInterest)
	}
	stem_ := stemFormCallback.stem
	_ = stem_

	for _, formDiv := range stemFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stem_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if stemFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stem_.Unstage(stemFormCallback.probe.stageOfInterest)
	}

	stemFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Stem](
		stemFormCallback.probe,
	)
	stemFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if stemFormCallback.CreationMode || stemFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stemFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(stemFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StemFormCallback(
			nil,
			stemFormCallback.probe,
			newFormGroup,
		)
		stem := new(models.Stem)
		FillUpForm(stem, newFormGroup, stemFormCallback.probe)
		stemFormCallback.probe.formStage.Commit()
	}

	fillUpTree(stemFormCallback.probe)
}
func __gong__New__StickFormCallback(
	stick *models.Stick,
	probe *Probe,
	formGroup *table.FormGroup,
) (stickFormCallback *StickFormCallback) {
	stickFormCallback = new(StickFormCallback)
	stickFormCallback.probe = probe
	stickFormCallback.stick = stick
	stickFormCallback.formGroup = formGroup

	stickFormCallback.CreationMode = (stick == nil)

	return
}

type StickFormCallback struct {
	stick *models.Stick

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (stickFormCallback *StickFormCallback) OnSave() {

	log.Println("StickFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stickFormCallback.probe.formStage.Checkout()

	if stickFormCallback.stick == nil {
		stickFormCallback.stick = new(models.Stick).Stage(stickFormCallback.probe.stageOfInterest)
	}
	stick_ := stickFormCallback.stick
	_ = stick_

	for _, formDiv := range stickFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stick_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if stickFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stick_.Unstage(stickFormCallback.probe.stageOfInterest)
	}

	stickFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Stick](
		stickFormCallback.probe,
	)
	stickFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if stickFormCallback.CreationMode || stickFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stickFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(stickFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StickFormCallback(
			nil,
			stickFormCallback.probe,
			newFormGroup,
		)
		stick := new(models.Stick)
		FillUpForm(stick, newFormGroup, stickFormCallback.probe)
		stickFormCallback.probe.formStage.Commit()
	}

	fillUpTree(stickFormCallback.probe)
}
func __gong__New__String_muteFormCallback(
	string_mute *models.String_mute,
	probe *Probe,
	formGroup *table.FormGroup,
) (string_muteFormCallback *String_muteFormCallback) {
	string_muteFormCallback = new(String_muteFormCallback)
	string_muteFormCallback.probe = probe
	string_muteFormCallback.string_mute = string_mute
	string_muteFormCallback.formGroup = formGroup

	string_muteFormCallback.CreationMode = (string_mute == nil)

	return
}

type String_muteFormCallback struct {
	string_mute *models.String_mute

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (string_muteFormCallback *String_muteFormCallback) OnSave() {

	log.Println("String_muteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	string_muteFormCallback.probe.formStage.Checkout()

	if string_muteFormCallback.string_mute == nil {
		string_muteFormCallback.string_mute = new(models.String_mute).Stage(string_muteFormCallback.probe.stageOfInterest)
	}
	string_mute_ := string_muteFormCallback.string_mute
	_ = string_mute_

	for _, formDiv := range string_muteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(string_mute_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if string_muteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		string_mute_.Unstage(string_muteFormCallback.probe.stageOfInterest)
	}

	string_muteFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.String_mute](
		string_muteFormCallback.probe,
	)
	string_muteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if string_muteFormCallback.CreationMode || string_muteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		string_muteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(string_muteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__String_muteFormCallback(
			nil,
			string_muteFormCallback.probe,
			newFormGroup,
		)
		string_mute := new(models.String_mute)
		FillUpForm(string_mute, newFormGroup, string_muteFormCallback.probe)
		string_muteFormCallback.probe.formStage.Commit()
	}

	fillUpTree(string_muteFormCallback.probe)
}
func __gong__New__Strong_accentFormCallback(
	strong_accent *models.Strong_accent,
	probe *Probe,
	formGroup *table.FormGroup,
) (strong_accentFormCallback *Strong_accentFormCallback) {
	strong_accentFormCallback = new(Strong_accentFormCallback)
	strong_accentFormCallback.probe = probe
	strong_accentFormCallback.strong_accent = strong_accent
	strong_accentFormCallback.formGroup = formGroup

	strong_accentFormCallback.CreationMode = (strong_accent == nil)

	return
}

type Strong_accentFormCallback struct {
	strong_accent *models.Strong_accent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (strong_accentFormCallback *Strong_accentFormCallback) OnSave() {

	log.Println("Strong_accentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	strong_accentFormCallback.probe.formStage.Checkout()

	if strong_accentFormCallback.strong_accent == nil {
		strong_accentFormCallback.strong_accent = new(models.Strong_accent).Stage(strong_accentFormCallback.probe.stageOfInterest)
	}
	strong_accent_ := strong_accentFormCallback.strong_accent
	_ = strong_accent_

	for _, formDiv := range strong_accentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(strong_accent_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if strong_accentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		strong_accent_.Unstage(strong_accentFormCallback.probe.stageOfInterest)
	}

	strong_accentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Strong_accent](
		strong_accentFormCallback.probe,
	)
	strong_accentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if strong_accentFormCallback.CreationMode || strong_accentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		strong_accentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(strong_accentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Strong_accentFormCallback(
			nil,
			strong_accentFormCallback.probe,
			newFormGroup,
		)
		strong_accent := new(models.Strong_accent)
		FillUpForm(strong_accent, newFormGroup, strong_accentFormCallback.probe)
		strong_accentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(strong_accentFormCallback.probe)
}
func __gong__New__SupportsFormCallback(
	supports *models.Supports,
	probe *Probe,
	formGroup *table.FormGroup,
) (supportsFormCallback *SupportsFormCallback) {
	supportsFormCallback = new(SupportsFormCallback)
	supportsFormCallback.probe = probe
	supportsFormCallback.supports = supports
	supportsFormCallback.formGroup = formGroup

	supportsFormCallback.CreationMode = (supports == nil)

	return
}

type SupportsFormCallback struct {
	supports *models.Supports

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (supportsFormCallback *SupportsFormCallback) OnSave() {

	log.Println("SupportsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	supportsFormCallback.probe.formStage.Checkout()

	if supportsFormCallback.supports == nil {
		supportsFormCallback.supports = new(models.Supports).Stage(supportsFormCallback.probe.stageOfInterest)
	}
	supports_ := supportsFormCallback.supports
	_ = supports_

	for _, formDiv := range supportsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(supports_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(supports_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if supportsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		supports_.Unstage(supportsFormCallback.probe.stageOfInterest)
	}

	supportsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Supports](
		supportsFormCallback.probe,
	)
	supportsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if supportsFormCallback.CreationMode || supportsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		supportsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(supportsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SupportsFormCallback(
			nil,
			supportsFormCallback.probe,
			newFormGroup,
		)
		supports := new(models.Supports)
		FillUpForm(supports, newFormGroup, supportsFormCallback.probe)
		supportsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(supportsFormCallback.probe)
}
func __gong__New__SwingFormCallback(
	swing *models.Swing,
	probe *Probe,
	formGroup *table.FormGroup,
) (swingFormCallback *SwingFormCallback) {
	swingFormCallback = new(SwingFormCallback)
	swingFormCallback.probe = probe
	swingFormCallback.swing = swing
	swingFormCallback.formGroup = formGroup

	swingFormCallback.CreationMode = (swing == nil)

	return
}

type SwingFormCallback struct {
	swing *models.Swing

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (swingFormCallback *SwingFormCallback) OnSave() {

	log.Println("SwingFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	swingFormCallback.probe.formStage.Checkout()

	if swingFormCallback.swing == nil {
		swingFormCallback.swing = new(models.Swing).Stage(swingFormCallback.probe.stageOfInterest)
	}
	swing_ := swingFormCallback.swing
	_ = swing_

	for _, formDiv := range swingFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(swing_.Name), formDiv)
		case "Swing_style":
			FormDivBasicFieldToField(&(swing_.Swing_style), formDiv)
		case "Straight":
			FormDivSelectFieldToField(&(swing_.Straight), swingFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if swingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		swing_.Unstage(swingFormCallback.probe.stageOfInterest)
	}

	swingFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Swing](
		swingFormCallback.probe,
	)
	swingFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if swingFormCallback.CreationMode || swingFormCallback.formGroup.HasSuppressButtonBeenPressed {
		swingFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(swingFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SwingFormCallback(
			nil,
			swingFormCallback.probe,
			newFormGroup,
		)
		swing := new(models.Swing)
		FillUpForm(swing, newFormGroup, swingFormCallback.probe)
		swingFormCallback.probe.formStage.Commit()
	}

	fillUpTree(swingFormCallback.probe)
}
func __gong__New__SyncFormCallback(
	sync *models.Sync,
	probe *Probe,
	formGroup *table.FormGroup,
) (syncFormCallback *SyncFormCallback) {
	syncFormCallback = new(SyncFormCallback)
	syncFormCallback.probe = probe
	syncFormCallback.sync = sync
	syncFormCallback.formGroup = formGroup

	syncFormCallback.CreationMode = (sync == nil)

	return
}

type SyncFormCallback struct {
	sync *models.Sync

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (syncFormCallback *SyncFormCallback) OnSave() {

	log.Println("SyncFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	syncFormCallback.probe.formStage.Checkout()

	if syncFormCallback.sync == nil {
		syncFormCallback.sync = new(models.Sync).Stage(syncFormCallback.probe.stageOfInterest)
	}
	sync_ := syncFormCallback.sync
	_ = sync_

	for _, formDiv := range syncFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sync_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if syncFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sync_.Unstage(syncFormCallback.probe.stageOfInterest)
	}

	syncFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Sync](
		syncFormCallback.probe,
	)
	syncFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if syncFormCallback.CreationMode || syncFormCallback.formGroup.HasSuppressButtonBeenPressed {
		syncFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(syncFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SyncFormCallback(
			nil,
			syncFormCallback.probe,
			newFormGroup,
		)
		sync := new(models.Sync)
		FillUpForm(sync, newFormGroup, syncFormCallback.probe)
		syncFormCallback.probe.formStage.Commit()
	}

	fillUpTree(syncFormCallback.probe)
}
func __gong__New__System_dividersFormCallback(
	system_dividers *models.System_dividers,
	probe *Probe,
	formGroup *table.FormGroup,
) (system_dividersFormCallback *System_dividersFormCallback) {
	system_dividersFormCallback = new(System_dividersFormCallback)
	system_dividersFormCallback.probe = probe
	system_dividersFormCallback.system_dividers = system_dividers
	system_dividersFormCallback.formGroup = formGroup

	system_dividersFormCallback.CreationMode = (system_dividers == nil)

	return
}

type System_dividersFormCallback struct {
	system_dividers *models.System_dividers

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (system_dividersFormCallback *System_dividersFormCallback) OnSave() {

	log.Println("System_dividersFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	system_dividersFormCallback.probe.formStage.Checkout()

	if system_dividersFormCallback.system_dividers == nil {
		system_dividersFormCallback.system_dividers = new(models.System_dividers).Stage(system_dividersFormCallback.probe.stageOfInterest)
	}
	system_dividers_ := system_dividersFormCallback.system_dividers
	_ = system_dividers_

	for _, formDiv := range system_dividersFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(system_dividers_.Name), formDiv)
		case "Left_divider":
			FormDivSelectFieldToField(&(system_dividers_.Left_divider), system_dividersFormCallback.probe.stageOfInterest, formDiv)
		case "Right_divider":
			FormDivSelectFieldToField(&(system_dividers_.Right_divider), system_dividersFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if system_dividersFormCallback.formGroup.HasSuppressButtonBeenPressed {
		system_dividers_.Unstage(system_dividersFormCallback.probe.stageOfInterest)
	}

	system_dividersFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.System_dividers](
		system_dividersFormCallback.probe,
	)
	system_dividersFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if system_dividersFormCallback.CreationMode || system_dividersFormCallback.formGroup.HasSuppressButtonBeenPressed {
		system_dividersFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(system_dividersFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__System_dividersFormCallback(
			nil,
			system_dividersFormCallback.probe,
			newFormGroup,
		)
		system_dividers := new(models.System_dividers)
		FillUpForm(system_dividers, newFormGroup, system_dividersFormCallback.probe)
		system_dividersFormCallback.probe.formStage.Commit()
	}

	fillUpTree(system_dividersFormCallback.probe)
}
func __gong__New__System_layoutFormCallback(
	system_layout *models.System_layout,
	probe *Probe,
	formGroup *table.FormGroup,
) (system_layoutFormCallback *System_layoutFormCallback) {
	system_layoutFormCallback = new(System_layoutFormCallback)
	system_layoutFormCallback.probe = probe
	system_layoutFormCallback.system_layout = system_layout
	system_layoutFormCallback.formGroup = formGroup

	system_layoutFormCallback.CreationMode = (system_layout == nil)

	return
}

type System_layoutFormCallback struct {
	system_layout *models.System_layout

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (system_layoutFormCallback *System_layoutFormCallback) OnSave() {

	log.Println("System_layoutFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	system_layoutFormCallback.probe.formStage.Checkout()

	if system_layoutFormCallback.system_layout == nil {
		system_layoutFormCallback.system_layout = new(models.System_layout).Stage(system_layoutFormCallback.probe.stageOfInterest)
	}
	system_layout_ := system_layoutFormCallback.system_layout
	_ = system_layout_

	for _, formDiv := range system_layoutFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(system_layout_.Name), formDiv)
		case "System_margins":
			FormDivSelectFieldToField(&(system_layout_.System_margins), system_layoutFormCallback.probe.stageOfInterest, formDiv)
		case "System_dividers":
			FormDivSelectFieldToField(&(system_layout_.System_dividers), system_layoutFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if system_layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		system_layout_.Unstage(system_layoutFormCallback.probe.stageOfInterest)
	}

	system_layoutFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.System_layout](
		system_layoutFormCallback.probe,
	)
	system_layoutFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if system_layoutFormCallback.CreationMode || system_layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		system_layoutFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(system_layoutFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__System_layoutFormCallback(
			nil,
			system_layoutFormCallback.probe,
			newFormGroup,
		)
		system_layout := new(models.System_layout)
		FillUpForm(system_layout, newFormGroup, system_layoutFormCallback.probe)
		system_layoutFormCallback.probe.formStage.Commit()
	}

	fillUpTree(system_layoutFormCallback.probe)
}
func __gong__New__System_marginsFormCallback(
	system_margins *models.System_margins,
	probe *Probe,
	formGroup *table.FormGroup,
) (system_marginsFormCallback *System_marginsFormCallback) {
	system_marginsFormCallback = new(System_marginsFormCallback)
	system_marginsFormCallback.probe = probe
	system_marginsFormCallback.system_margins = system_margins
	system_marginsFormCallback.formGroup = formGroup

	system_marginsFormCallback.CreationMode = (system_margins == nil)

	return
}

type System_marginsFormCallback struct {
	system_margins *models.System_margins

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (system_marginsFormCallback *System_marginsFormCallback) OnSave() {

	log.Println("System_marginsFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	system_marginsFormCallback.probe.formStage.Checkout()

	if system_marginsFormCallback.system_margins == nil {
		system_marginsFormCallback.system_margins = new(models.System_margins).Stage(system_marginsFormCallback.probe.stageOfInterest)
	}
	system_margins_ := system_marginsFormCallback.system_margins
	_ = system_margins_

	for _, formDiv := range system_marginsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(system_margins_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if system_marginsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		system_margins_.Unstage(system_marginsFormCallback.probe.stageOfInterest)
	}

	system_marginsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.System_margins](
		system_marginsFormCallback.probe,
	)
	system_marginsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if system_marginsFormCallback.CreationMode || system_marginsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		system_marginsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(system_marginsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__System_marginsFormCallback(
			nil,
			system_marginsFormCallback.probe,
			newFormGroup,
		)
		system_margins := new(models.System_margins)
		FillUpForm(system_margins, newFormGroup, system_marginsFormCallback.probe)
		system_marginsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(system_marginsFormCallback.probe)
}
func __gong__New__TapFormCallback(
	tap *models.Tap,
	probe *Probe,
	formGroup *table.FormGroup,
) (tapFormCallback *TapFormCallback) {
	tapFormCallback = new(TapFormCallback)
	tapFormCallback.probe = probe
	tapFormCallback.tap = tap
	tapFormCallback.formGroup = formGroup

	tapFormCallback.CreationMode = (tap == nil)

	return
}

type TapFormCallback struct {
	tap *models.Tap

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tapFormCallback *TapFormCallback) OnSave() {

	log.Println("TapFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tapFormCallback.probe.formStage.Checkout()

	if tapFormCallback.tap == nil {
		tapFormCallback.tap = new(models.Tap).Stage(tapFormCallback.probe.stageOfInterest)
	}
	tap_ := tapFormCallback.tap
	_ = tap_

	for _, formDiv := range tapFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tap_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(tap_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if tapFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tap_.Unstage(tapFormCallback.probe.stageOfInterest)
	}

	tapFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tap](
		tapFormCallback.probe,
	)
	tapFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tapFormCallback.CreationMode || tapFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tapFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(tapFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TapFormCallback(
			nil,
			tapFormCallback.probe,
			newFormGroup,
		)
		tap := new(models.Tap)
		FillUpForm(tap, newFormGroup, tapFormCallback.probe)
		tapFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tapFormCallback.probe)
}
func __gong__New__TechnicalFormCallback(
	technical *models.Technical,
	probe *Probe,
	formGroup *table.FormGroup,
) (technicalFormCallback *TechnicalFormCallback) {
	technicalFormCallback = new(TechnicalFormCallback)
	technicalFormCallback.probe = probe
	technicalFormCallback.technical = technical
	technicalFormCallback.formGroup = formGroup

	technicalFormCallback.CreationMode = (technical == nil)

	return
}

type TechnicalFormCallback struct {
	technical *models.Technical

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (technicalFormCallback *TechnicalFormCallback) OnSave() {

	log.Println("TechnicalFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	technicalFormCallback.probe.formStage.Checkout()

	if technicalFormCallback.technical == nil {
		technicalFormCallback.technical = new(models.Technical).Stage(technicalFormCallback.probe.stageOfInterest)
	}
	technical_ := technicalFormCallback.technical
	_ = technical_

	for _, formDiv := range technicalFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(technical_.Name), formDiv)
		case "Up_bow":
			FormDivSelectFieldToField(&(technical_.Up_bow), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Down_bow":
			FormDivSelectFieldToField(&(technical_.Down_bow), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Harmonic":
			FormDivSelectFieldToField(&(technical_.Harmonic), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Open_string":
			FormDivSelectFieldToField(&(technical_.Open_string), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Thumb_position":
			FormDivSelectFieldToField(&(technical_.Thumb_position), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Fingering":
			FormDivSelectFieldToField(&(technical_.Fingering), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Double_tongue":
			FormDivSelectFieldToField(&(technical_.Double_tongue), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Triple_tongue":
			FormDivSelectFieldToField(&(technical_.Triple_tongue), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Stopped":
			FormDivSelectFieldToField(&(technical_.Stopped), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Snap_pizzicato":
			FormDivSelectFieldToField(&(technical_.Snap_pizzicato), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Fret":
			FormDivSelectFieldToField(&(technical_.Fret), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Astring":
			FormDivBasicFieldToField(&(technical_.Astring), formDiv)
		case "Hammer_on":
			FormDivSelectFieldToField(&(technical_.Hammer_on), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Pull_off":
			FormDivSelectFieldToField(&(technical_.Pull_off), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Bend":
			FormDivSelectFieldToField(&(technical_.Bend), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Tap":
			FormDivSelectFieldToField(&(technical_.Tap), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Heel":
			FormDivSelectFieldToField(&(technical_.Heel), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Toe":
			FormDivSelectFieldToField(&(technical_.Toe), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Fingernails":
			FormDivSelectFieldToField(&(technical_.Fingernails), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Hole":
			FormDivSelectFieldToField(&(technical_.Hole), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Arrow":
			FormDivSelectFieldToField(&(technical_.Arrow), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Handbell":
			FormDivSelectFieldToField(&(technical_.Handbell), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Brass_bend":
			FormDivSelectFieldToField(&(technical_.Brass_bend), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Flip":
			FormDivSelectFieldToField(&(technical_.Flip), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Smear":
			FormDivSelectFieldToField(&(technical_.Smear), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Open":
			FormDivSelectFieldToField(&(technical_.Open), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Half_muted":
			FormDivSelectFieldToField(&(technical_.Half_muted), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Harmon_mute":
			FormDivSelectFieldToField(&(technical_.Harmon_mute), technicalFormCallback.probe.stageOfInterest, formDiv)
		case "Golpe":
			FormDivSelectFieldToField(&(technical_.Golpe), technicalFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if technicalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		technical_.Unstage(technicalFormCallback.probe.stageOfInterest)
	}

	technicalFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Technical](
		technicalFormCallback.probe,
	)
	technicalFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if technicalFormCallback.CreationMode || technicalFormCallback.formGroup.HasSuppressButtonBeenPressed {
		technicalFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(technicalFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TechnicalFormCallback(
			nil,
			technicalFormCallback.probe,
			newFormGroup,
		)
		technical := new(models.Technical)
		FillUpForm(technical, newFormGroup, technicalFormCallback.probe)
		technicalFormCallback.probe.formStage.Commit()
	}

	fillUpTree(technicalFormCallback.probe)
}
func __gong__New__Text_element_dataFormCallback(
	text_element_data *models.Text_element_data,
	probe *Probe,
	formGroup *table.FormGroup,
) (text_element_dataFormCallback *Text_element_dataFormCallback) {
	text_element_dataFormCallback = new(Text_element_dataFormCallback)
	text_element_dataFormCallback.probe = probe
	text_element_dataFormCallback.text_element_data = text_element_data
	text_element_dataFormCallback.formGroup = formGroup

	text_element_dataFormCallback.CreationMode = (text_element_data == nil)

	return
}

type Text_element_dataFormCallback struct {
	text_element_data *models.Text_element_data

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (text_element_dataFormCallback *Text_element_dataFormCallback) OnSave() {

	log.Println("Text_element_dataFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	text_element_dataFormCallback.probe.formStage.Checkout()

	if text_element_dataFormCallback.text_element_data == nil {
		text_element_dataFormCallback.text_element_data = new(models.Text_element_data).Stage(text_element_dataFormCallback.probe.stageOfInterest)
	}
	text_element_data_ := text_element_dataFormCallback.text_element_data
	_ = text_element_data_

	for _, formDiv := range text_element_dataFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(text_element_data_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(text_element_data_.Value), formDiv)
		case "EmptyString":
			FormDivBasicFieldToField(&(text_element_data_.EmptyString), formDiv)
		}
	}

	// manage the suppress operation
	if text_element_dataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		text_element_data_.Unstage(text_element_dataFormCallback.probe.stageOfInterest)
	}

	text_element_dataFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Text_element_data](
		text_element_dataFormCallback.probe,
	)
	text_element_dataFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if text_element_dataFormCallback.CreationMode || text_element_dataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		text_element_dataFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(text_element_dataFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Text_element_dataFormCallback(
			nil,
			text_element_dataFormCallback.probe,
			newFormGroup,
		)
		text_element_data := new(models.Text_element_data)
		FillUpForm(text_element_data, newFormGroup, text_element_dataFormCallback.probe)
		text_element_dataFormCallback.probe.formStage.Commit()
	}

	fillUpTree(text_element_dataFormCallback.probe)
}
func __gong__New__TieFormCallback(
	tie *models.Tie,
	probe *Probe,
	formGroup *table.FormGroup,
) (tieFormCallback *TieFormCallback) {
	tieFormCallback = new(TieFormCallback)
	tieFormCallback.probe = probe
	tieFormCallback.tie = tie
	tieFormCallback.formGroup = formGroup

	tieFormCallback.CreationMode = (tie == nil)

	return
}

type TieFormCallback struct {
	tie *models.Tie

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tieFormCallback *TieFormCallback) OnSave() {

	log.Println("TieFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tieFormCallback.probe.formStage.Checkout()

	if tieFormCallback.tie == nil {
		tieFormCallback.tie = new(models.Tie).Stage(tieFormCallback.probe.stageOfInterest)
	}
	tie_ := tieFormCallback.tie
	_ = tie_

	for _, formDiv := range tieFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tie_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if tieFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tie_.Unstage(tieFormCallback.probe.stageOfInterest)
	}

	tieFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tie](
		tieFormCallback.probe,
	)
	tieFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tieFormCallback.CreationMode || tieFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tieFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(tieFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TieFormCallback(
			nil,
			tieFormCallback.probe,
			newFormGroup,
		)
		tie := new(models.Tie)
		FillUpForm(tie, newFormGroup, tieFormCallback.probe)
		tieFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tieFormCallback.probe)
}
func __gong__New__TiedFormCallback(
	tied *models.Tied,
	probe *Probe,
	formGroup *table.FormGroup,
) (tiedFormCallback *TiedFormCallback) {
	tiedFormCallback = new(TiedFormCallback)
	tiedFormCallback.probe = probe
	tiedFormCallback.tied = tied
	tiedFormCallback.formGroup = formGroup

	tiedFormCallback.CreationMode = (tied == nil)

	return
}

type TiedFormCallback struct {
	tied *models.Tied

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tiedFormCallback *TiedFormCallback) OnSave() {

	log.Println("TiedFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tiedFormCallback.probe.formStage.Checkout()

	if tiedFormCallback.tied == nil {
		tiedFormCallback.tied = new(models.Tied).Stage(tiedFormCallback.probe.stageOfInterest)
	}
	tied_ := tiedFormCallback.tied
	_ = tied_

	for _, formDiv := range tiedFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tied_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if tiedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tied_.Unstage(tiedFormCallback.probe.stageOfInterest)
	}

	tiedFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tied](
		tiedFormCallback.probe,
	)
	tiedFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tiedFormCallback.CreationMode || tiedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tiedFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(tiedFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TiedFormCallback(
			nil,
			tiedFormCallback.probe,
			newFormGroup,
		)
		tied := new(models.Tied)
		FillUpForm(tied, newFormGroup, tiedFormCallback.probe)
		tiedFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tiedFormCallback.probe)
}
func __gong__New__TimeFormCallback(
	time *models.Time,
	probe *Probe,
	formGroup *table.FormGroup,
) (timeFormCallback *TimeFormCallback) {
	timeFormCallback = new(TimeFormCallback)
	timeFormCallback.probe = probe
	timeFormCallback.time = time
	timeFormCallback.formGroup = formGroup

	timeFormCallback.CreationMode = (time == nil)

	return
}

type TimeFormCallback struct {
	time *models.Time

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (timeFormCallback *TimeFormCallback) OnSave() {

	log.Println("TimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	timeFormCallback.probe.formStage.Checkout()

	if timeFormCallback.time == nil {
		timeFormCallback.time = new(models.Time).Stage(timeFormCallback.probe.stageOfInterest)
	}
	time_ := timeFormCallback.time
	_ = time_

	for _, formDiv := range timeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(time_.Name), formDiv)
		case "Senza_misura":
			FormDivBasicFieldToField(&(time_.Senza_misura), formDiv)
		}
	}

	// manage the suppress operation
	if timeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		time_.Unstage(timeFormCallback.probe.stageOfInterest)
	}

	timeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Time](
		timeFormCallback.probe,
	)
	timeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if timeFormCallback.CreationMode || timeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		timeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(timeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TimeFormCallback(
			nil,
			timeFormCallback.probe,
			newFormGroup,
		)
		time := new(models.Time)
		FillUpForm(time, newFormGroup, timeFormCallback.probe)
		timeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(timeFormCallback.probe)
}
func __gong__New__Time_modificationFormCallback(
	time_modification *models.Time_modification,
	probe *Probe,
	formGroup *table.FormGroup,
) (time_modificationFormCallback *Time_modificationFormCallback) {
	time_modificationFormCallback = new(Time_modificationFormCallback)
	time_modificationFormCallback.probe = probe
	time_modificationFormCallback.time_modification = time_modification
	time_modificationFormCallback.formGroup = formGroup

	time_modificationFormCallback.CreationMode = (time_modification == nil)

	return
}

type Time_modificationFormCallback struct {
	time_modification *models.Time_modification

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (time_modificationFormCallback *Time_modificationFormCallback) OnSave() {

	log.Println("Time_modificationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	time_modificationFormCallback.probe.formStage.Checkout()

	if time_modificationFormCallback.time_modification == nil {
		time_modificationFormCallback.time_modification = new(models.Time_modification).Stage(time_modificationFormCallback.probe.stageOfInterest)
	}
	time_modification_ := time_modificationFormCallback.time_modification
	_ = time_modification_

	for _, formDiv := range time_modificationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(time_modification_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if time_modificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		time_modification_.Unstage(time_modificationFormCallback.probe.stageOfInterest)
	}

	time_modificationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Time_modification](
		time_modificationFormCallback.probe,
	)
	time_modificationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if time_modificationFormCallback.CreationMode || time_modificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		time_modificationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(time_modificationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Time_modificationFormCallback(
			nil,
			time_modificationFormCallback.probe,
			newFormGroup,
		)
		time_modification := new(models.Time_modification)
		FillUpForm(time_modification, newFormGroup, time_modificationFormCallback.probe)
		time_modificationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(time_modificationFormCallback.probe)
}
func __gong__New__TimpaniFormCallback(
	timpani *models.Timpani,
	probe *Probe,
	formGroup *table.FormGroup,
) (timpaniFormCallback *TimpaniFormCallback) {
	timpaniFormCallback = new(TimpaniFormCallback)
	timpaniFormCallback.probe = probe
	timpaniFormCallback.timpani = timpani
	timpaniFormCallback.formGroup = formGroup

	timpaniFormCallback.CreationMode = (timpani == nil)

	return
}

type TimpaniFormCallback struct {
	timpani *models.Timpani

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (timpaniFormCallback *TimpaniFormCallback) OnSave() {

	log.Println("TimpaniFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	timpaniFormCallback.probe.formStage.Checkout()

	if timpaniFormCallback.timpani == nil {
		timpaniFormCallback.timpani = new(models.Timpani).Stage(timpaniFormCallback.probe.stageOfInterest)
	}
	timpani_ := timpaniFormCallback.timpani
	_ = timpani_

	for _, formDiv := range timpaniFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(timpani_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if timpaniFormCallback.formGroup.HasSuppressButtonBeenPressed {
		timpani_.Unstage(timpaniFormCallback.probe.stageOfInterest)
	}

	timpaniFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Timpani](
		timpaniFormCallback.probe,
	)
	timpaniFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if timpaniFormCallback.CreationMode || timpaniFormCallback.formGroup.HasSuppressButtonBeenPressed {
		timpaniFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(timpaniFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TimpaniFormCallback(
			nil,
			timpaniFormCallback.probe,
			newFormGroup,
		)
		timpani := new(models.Timpani)
		FillUpForm(timpani, newFormGroup, timpaniFormCallback.probe)
		timpaniFormCallback.probe.formStage.Commit()
	}

	fillUpTree(timpaniFormCallback.probe)
}
func __gong__New__TransposeFormCallback(
	transpose *models.Transpose,
	probe *Probe,
	formGroup *table.FormGroup,
) (transposeFormCallback *TransposeFormCallback) {
	transposeFormCallback = new(TransposeFormCallback)
	transposeFormCallback.probe = probe
	transposeFormCallback.transpose = transpose
	transposeFormCallback.formGroup = formGroup

	transposeFormCallback.CreationMode = (transpose == nil)

	return
}

type TransposeFormCallback struct {
	transpose *models.Transpose

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (transposeFormCallback *TransposeFormCallback) OnSave() {

	log.Println("TransposeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	transposeFormCallback.probe.formStage.Checkout()

	if transposeFormCallback.transpose == nil {
		transposeFormCallback.transpose = new(models.Transpose).Stage(transposeFormCallback.probe.stageOfInterest)
	}
	transpose_ := transposeFormCallback.transpose
	_ = transpose_

	for _, formDiv := range transposeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(transpose_.Name), formDiv)
		case "Attributes:Transpose":
			// we need to retrieve the field owner before the change
			var pastAttributesOwner *models.Attributes
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Transpose"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				transposeFormCallback.probe.stageOfInterest,
				transposeFormCallback.probe.backRepoOfInterest,
				transpose_,
				&rf)

			if reverseFieldOwner != nil {
				pastAttributesOwner = reverseFieldOwner.(*models.Attributes)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAttributesOwner != nil {
					idx := slices.Index(pastAttributesOwner.Transpose, transpose_)
					pastAttributesOwner.Transpose = slices.Delete(pastAttributesOwner.Transpose, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attributes := range *models.GetGongstructInstancesSet[models.Attributes](transposeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAttributesOwner := _attributes // we have a match
						if pastAttributesOwner != nil {
							if newAttributesOwner != pastAttributesOwner {
								idx := slices.Index(pastAttributesOwner.Transpose, transpose_)
								pastAttributesOwner.Transpose = slices.Delete(pastAttributesOwner.Transpose, idx, idx+1)
								newAttributesOwner.Transpose = append(newAttributesOwner.Transpose, transpose_)
							}
						} else {
							newAttributesOwner.Transpose = append(newAttributesOwner.Transpose, transpose_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if transposeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		transpose_.Unstage(transposeFormCallback.probe.stageOfInterest)
	}

	transposeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Transpose](
		transposeFormCallback.probe,
	)
	transposeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if transposeFormCallback.CreationMode || transposeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		transposeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(transposeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TransposeFormCallback(
			nil,
			transposeFormCallback.probe,
			newFormGroup,
		)
		transpose := new(models.Transpose)
		FillUpForm(transpose, newFormGroup, transposeFormCallback.probe)
		transposeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(transposeFormCallback.probe)
}
func __gong__New__TremoloFormCallback(
	tremolo *models.Tremolo,
	probe *Probe,
	formGroup *table.FormGroup,
) (tremoloFormCallback *TremoloFormCallback) {
	tremoloFormCallback = new(TremoloFormCallback)
	tremoloFormCallback.probe = probe
	tremoloFormCallback.tremolo = tremolo
	tremoloFormCallback.formGroup = formGroup

	tremoloFormCallback.CreationMode = (tremolo == nil)

	return
}

type TremoloFormCallback struct {
	tremolo *models.Tremolo

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tremoloFormCallback *TremoloFormCallback) OnSave() {

	log.Println("TremoloFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tremoloFormCallback.probe.formStage.Checkout()

	if tremoloFormCallback.tremolo == nil {
		tremoloFormCallback.tremolo = new(models.Tremolo).Stage(tremoloFormCallback.probe.stageOfInterest)
	}
	tremolo_ := tremoloFormCallback.tremolo
	_ = tremolo_

	for _, formDiv := range tremoloFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tremolo_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if tremoloFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tremolo_.Unstage(tremoloFormCallback.probe.stageOfInterest)
	}

	tremoloFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tremolo](
		tremoloFormCallback.probe,
	)
	tremoloFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tremoloFormCallback.CreationMode || tremoloFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tremoloFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(tremoloFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TremoloFormCallback(
			nil,
			tremoloFormCallback.probe,
			newFormGroup,
		)
		tremolo := new(models.Tremolo)
		FillUpForm(tremolo, newFormGroup, tremoloFormCallback.probe)
		tremoloFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tremoloFormCallback.probe)
}
func __gong__New__TupletFormCallback(
	tuplet *models.Tuplet,
	probe *Probe,
	formGroup *table.FormGroup,
) (tupletFormCallback *TupletFormCallback) {
	tupletFormCallback = new(TupletFormCallback)
	tupletFormCallback.probe = probe
	tupletFormCallback.tuplet = tuplet
	tupletFormCallback.formGroup = formGroup

	tupletFormCallback.CreationMode = (tuplet == nil)

	return
}

type TupletFormCallback struct {
	tuplet *models.Tuplet

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tupletFormCallback *TupletFormCallback) OnSave() {

	log.Println("TupletFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tupletFormCallback.probe.formStage.Checkout()

	if tupletFormCallback.tuplet == nil {
		tupletFormCallback.tuplet = new(models.Tuplet).Stage(tupletFormCallback.probe.stageOfInterest)
	}
	tuplet_ := tupletFormCallback.tuplet
	_ = tuplet_

	for _, formDiv := range tupletFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tuplet_.Name), formDiv)
		case "Tuplet_actual":
			FormDivSelectFieldToField(&(tuplet_.Tuplet_actual), tupletFormCallback.probe.stageOfInterest, formDiv)
		case "Tuplet_normal":
			FormDivSelectFieldToField(&(tuplet_.Tuplet_normal), tupletFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if tupletFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tuplet_.Unstage(tupletFormCallback.probe.stageOfInterest)
	}

	tupletFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tuplet](
		tupletFormCallback.probe,
	)
	tupletFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tupletFormCallback.CreationMode || tupletFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tupletFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(tupletFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TupletFormCallback(
			nil,
			tupletFormCallback.probe,
			newFormGroup,
		)
		tuplet := new(models.Tuplet)
		FillUpForm(tuplet, newFormGroup, tupletFormCallback.probe)
		tupletFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tupletFormCallback.probe)
}
func __gong__New__Tuplet_dotFormCallback(
	tuplet_dot *models.Tuplet_dot,
	probe *Probe,
	formGroup *table.FormGroup,
) (tuplet_dotFormCallback *Tuplet_dotFormCallback) {
	tuplet_dotFormCallback = new(Tuplet_dotFormCallback)
	tuplet_dotFormCallback.probe = probe
	tuplet_dotFormCallback.tuplet_dot = tuplet_dot
	tuplet_dotFormCallback.formGroup = formGroup

	tuplet_dotFormCallback.CreationMode = (tuplet_dot == nil)

	return
}

type Tuplet_dotFormCallback struct {
	tuplet_dot *models.Tuplet_dot

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tuplet_dotFormCallback *Tuplet_dotFormCallback) OnSave() {

	log.Println("Tuplet_dotFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tuplet_dotFormCallback.probe.formStage.Checkout()

	if tuplet_dotFormCallback.tuplet_dot == nil {
		tuplet_dotFormCallback.tuplet_dot = new(models.Tuplet_dot).Stage(tuplet_dotFormCallback.probe.stageOfInterest)
	}
	tuplet_dot_ := tuplet_dotFormCallback.tuplet_dot
	_ = tuplet_dot_

	for _, formDiv := range tuplet_dotFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tuplet_dot_.Name), formDiv)
		case "Tuplet_portion:Tuplet_dot":
			// we need to retrieve the field owner before the change
			var pastTuplet_portionOwner *models.Tuplet_portion
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Tuplet_portion"
			rf.Fieldname = "Tuplet_dot"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				tuplet_dotFormCallback.probe.stageOfInterest,
				tuplet_dotFormCallback.probe.backRepoOfInterest,
				tuplet_dot_,
				&rf)

			if reverseFieldOwner != nil {
				pastTuplet_portionOwner = reverseFieldOwner.(*models.Tuplet_portion)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTuplet_portionOwner != nil {
					idx := slices.Index(pastTuplet_portionOwner.Tuplet_dot, tuplet_dot_)
					pastTuplet_portionOwner.Tuplet_dot = slices.Delete(pastTuplet_portionOwner.Tuplet_dot, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _tuplet_portion := range *models.GetGongstructInstancesSet[models.Tuplet_portion](tuplet_dotFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _tuplet_portion.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTuplet_portionOwner := _tuplet_portion // we have a match
						if pastTuplet_portionOwner != nil {
							if newTuplet_portionOwner != pastTuplet_portionOwner {
								idx := slices.Index(pastTuplet_portionOwner.Tuplet_dot, tuplet_dot_)
								pastTuplet_portionOwner.Tuplet_dot = slices.Delete(pastTuplet_portionOwner.Tuplet_dot, idx, idx+1)
								newTuplet_portionOwner.Tuplet_dot = append(newTuplet_portionOwner.Tuplet_dot, tuplet_dot_)
							}
						} else {
							newTuplet_portionOwner.Tuplet_dot = append(newTuplet_portionOwner.Tuplet_dot, tuplet_dot_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if tuplet_dotFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tuplet_dot_.Unstage(tuplet_dotFormCallback.probe.stageOfInterest)
	}

	tuplet_dotFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tuplet_dot](
		tuplet_dotFormCallback.probe,
	)
	tuplet_dotFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tuplet_dotFormCallback.CreationMode || tuplet_dotFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tuplet_dotFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(tuplet_dotFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Tuplet_dotFormCallback(
			nil,
			tuplet_dotFormCallback.probe,
			newFormGroup,
		)
		tuplet_dot := new(models.Tuplet_dot)
		FillUpForm(tuplet_dot, newFormGroup, tuplet_dotFormCallback.probe)
		tuplet_dotFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tuplet_dotFormCallback.probe)
}
func __gong__New__Tuplet_numberFormCallback(
	tuplet_number *models.Tuplet_number,
	probe *Probe,
	formGroup *table.FormGroup,
) (tuplet_numberFormCallback *Tuplet_numberFormCallback) {
	tuplet_numberFormCallback = new(Tuplet_numberFormCallback)
	tuplet_numberFormCallback.probe = probe
	tuplet_numberFormCallback.tuplet_number = tuplet_number
	tuplet_numberFormCallback.formGroup = formGroup

	tuplet_numberFormCallback.CreationMode = (tuplet_number == nil)

	return
}

type Tuplet_numberFormCallback struct {
	tuplet_number *models.Tuplet_number

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tuplet_numberFormCallback *Tuplet_numberFormCallback) OnSave() {

	log.Println("Tuplet_numberFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tuplet_numberFormCallback.probe.formStage.Checkout()

	if tuplet_numberFormCallback.tuplet_number == nil {
		tuplet_numberFormCallback.tuplet_number = new(models.Tuplet_number).Stage(tuplet_numberFormCallback.probe.stageOfInterest)
	}
	tuplet_number_ := tuplet_numberFormCallback.tuplet_number
	_ = tuplet_number_

	for _, formDiv := range tuplet_numberFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tuplet_number_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if tuplet_numberFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tuplet_number_.Unstage(tuplet_numberFormCallback.probe.stageOfInterest)
	}

	tuplet_numberFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tuplet_number](
		tuplet_numberFormCallback.probe,
	)
	tuplet_numberFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tuplet_numberFormCallback.CreationMode || tuplet_numberFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tuplet_numberFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(tuplet_numberFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Tuplet_numberFormCallback(
			nil,
			tuplet_numberFormCallback.probe,
			newFormGroup,
		)
		tuplet_number := new(models.Tuplet_number)
		FillUpForm(tuplet_number, newFormGroup, tuplet_numberFormCallback.probe)
		tuplet_numberFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tuplet_numberFormCallback.probe)
}
func __gong__New__Tuplet_portionFormCallback(
	tuplet_portion *models.Tuplet_portion,
	probe *Probe,
	formGroup *table.FormGroup,
) (tuplet_portionFormCallback *Tuplet_portionFormCallback) {
	tuplet_portionFormCallback = new(Tuplet_portionFormCallback)
	tuplet_portionFormCallback.probe = probe
	tuplet_portionFormCallback.tuplet_portion = tuplet_portion
	tuplet_portionFormCallback.formGroup = formGroup

	tuplet_portionFormCallback.CreationMode = (tuplet_portion == nil)

	return
}

type Tuplet_portionFormCallback struct {
	tuplet_portion *models.Tuplet_portion

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tuplet_portionFormCallback *Tuplet_portionFormCallback) OnSave() {

	log.Println("Tuplet_portionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tuplet_portionFormCallback.probe.formStage.Checkout()

	if tuplet_portionFormCallback.tuplet_portion == nil {
		tuplet_portionFormCallback.tuplet_portion = new(models.Tuplet_portion).Stage(tuplet_portionFormCallback.probe.stageOfInterest)
	}
	tuplet_portion_ := tuplet_portionFormCallback.tuplet_portion
	_ = tuplet_portion_

	for _, formDiv := range tuplet_portionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tuplet_portion_.Name), formDiv)
		case "Tuplet_number":
			FormDivSelectFieldToField(&(tuplet_portion_.Tuplet_number), tuplet_portionFormCallback.probe.stageOfInterest, formDiv)
		case "Tuplet_type":
			FormDivSelectFieldToField(&(tuplet_portion_.Tuplet_type), tuplet_portionFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if tuplet_portionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tuplet_portion_.Unstage(tuplet_portionFormCallback.probe.stageOfInterest)
	}

	tuplet_portionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tuplet_portion](
		tuplet_portionFormCallback.probe,
	)
	tuplet_portionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tuplet_portionFormCallback.CreationMode || tuplet_portionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tuplet_portionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(tuplet_portionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Tuplet_portionFormCallback(
			nil,
			tuplet_portionFormCallback.probe,
			newFormGroup,
		)
		tuplet_portion := new(models.Tuplet_portion)
		FillUpForm(tuplet_portion, newFormGroup, tuplet_portionFormCallback.probe)
		tuplet_portionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tuplet_portionFormCallback.probe)
}
func __gong__New__Tuplet_typeFormCallback(
	tuplet_type *models.Tuplet_type,
	probe *Probe,
	formGroup *table.FormGroup,
) (tuplet_typeFormCallback *Tuplet_typeFormCallback) {
	tuplet_typeFormCallback = new(Tuplet_typeFormCallback)
	tuplet_typeFormCallback.probe = probe
	tuplet_typeFormCallback.tuplet_type = tuplet_type
	tuplet_typeFormCallback.formGroup = formGroup

	tuplet_typeFormCallback.CreationMode = (tuplet_type == nil)

	return
}

type Tuplet_typeFormCallback struct {
	tuplet_type *models.Tuplet_type

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tuplet_typeFormCallback *Tuplet_typeFormCallback) OnSave() {

	log.Println("Tuplet_typeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tuplet_typeFormCallback.probe.formStage.Checkout()

	if tuplet_typeFormCallback.tuplet_type == nil {
		tuplet_typeFormCallback.tuplet_type = new(models.Tuplet_type).Stage(tuplet_typeFormCallback.probe.stageOfInterest)
	}
	tuplet_type_ := tuplet_typeFormCallback.tuplet_type
	_ = tuplet_type_

	for _, formDiv := range tuplet_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tuplet_type_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if tuplet_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tuplet_type_.Unstage(tuplet_typeFormCallback.probe.stageOfInterest)
	}

	tuplet_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tuplet_type](
		tuplet_typeFormCallback.probe,
	)
	tuplet_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tuplet_typeFormCallback.CreationMode || tuplet_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tuplet_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(tuplet_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Tuplet_typeFormCallback(
			nil,
			tuplet_typeFormCallback.probe,
			newFormGroup,
		)
		tuplet_type := new(models.Tuplet_type)
		FillUpForm(tuplet_type, newFormGroup, tuplet_typeFormCallback.probe)
		tuplet_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tuplet_typeFormCallback.probe)
}
func __gong__New__Typed_textFormCallback(
	typed_text *models.Typed_text,
	probe *Probe,
	formGroup *table.FormGroup,
) (typed_textFormCallback *Typed_textFormCallback) {
	typed_textFormCallback = new(Typed_textFormCallback)
	typed_textFormCallback.probe = probe
	typed_textFormCallback.typed_text = typed_text
	typed_textFormCallback.formGroup = formGroup

	typed_textFormCallback.CreationMode = (typed_text == nil)

	return
}

type Typed_textFormCallback struct {
	typed_text *models.Typed_text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (typed_textFormCallback *Typed_textFormCallback) OnSave() {

	log.Println("Typed_textFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	typed_textFormCallback.probe.formStage.Checkout()

	if typed_textFormCallback.typed_text == nil {
		typed_textFormCallback.typed_text = new(models.Typed_text).Stage(typed_textFormCallback.probe.stageOfInterest)
	}
	typed_text_ := typed_textFormCallback.typed_text
	_ = typed_text_

	for _, formDiv := range typed_textFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(typed_text_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(typed_text_.Value), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(typed_text_.Type), formDiv)
		case "Identification:Creator":
			// we need to retrieve the field owner before the change
			var pastIdentificationOwner *models.Identification
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Identification"
			rf.Fieldname = "Creator"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				typed_textFormCallback.probe.stageOfInterest,
				typed_textFormCallback.probe.backRepoOfInterest,
				typed_text_,
				&rf)

			if reverseFieldOwner != nil {
				pastIdentificationOwner = reverseFieldOwner.(*models.Identification)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastIdentificationOwner != nil {
					idx := slices.Index(pastIdentificationOwner.Creator, typed_text_)
					pastIdentificationOwner.Creator = slices.Delete(pastIdentificationOwner.Creator, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _identification := range *models.GetGongstructInstancesSet[models.Identification](typed_textFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _identification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newIdentificationOwner := _identification // we have a match
						if pastIdentificationOwner != nil {
							if newIdentificationOwner != pastIdentificationOwner {
								idx := slices.Index(pastIdentificationOwner.Creator, typed_text_)
								pastIdentificationOwner.Creator = slices.Delete(pastIdentificationOwner.Creator, idx, idx+1)
								newIdentificationOwner.Creator = append(newIdentificationOwner.Creator, typed_text_)
							}
						} else {
							newIdentificationOwner.Creator = append(newIdentificationOwner.Creator, typed_text_)
						}
					}
				}
			}
		case "Identification:Rights":
			// we need to retrieve the field owner before the change
			var pastIdentificationOwner *models.Identification
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Identification"
			rf.Fieldname = "Rights"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				typed_textFormCallback.probe.stageOfInterest,
				typed_textFormCallback.probe.backRepoOfInterest,
				typed_text_,
				&rf)

			if reverseFieldOwner != nil {
				pastIdentificationOwner = reverseFieldOwner.(*models.Identification)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastIdentificationOwner != nil {
					idx := slices.Index(pastIdentificationOwner.Rights, typed_text_)
					pastIdentificationOwner.Rights = slices.Delete(pastIdentificationOwner.Rights, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _identification := range *models.GetGongstructInstancesSet[models.Identification](typed_textFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _identification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newIdentificationOwner := _identification // we have a match
						if pastIdentificationOwner != nil {
							if newIdentificationOwner != pastIdentificationOwner {
								idx := slices.Index(pastIdentificationOwner.Rights, typed_text_)
								pastIdentificationOwner.Rights = slices.Delete(pastIdentificationOwner.Rights, idx, idx+1)
								newIdentificationOwner.Rights = append(newIdentificationOwner.Rights, typed_text_)
							}
						} else {
							newIdentificationOwner.Rights = append(newIdentificationOwner.Rights, typed_text_)
						}
					}
				}
			}
		case "Identification:Relation":
			// we need to retrieve the field owner before the change
			var pastIdentificationOwner *models.Identification
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Identification"
			rf.Fieldname = "Relation"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				typed_textFormCallback.probe.stageOfInterest,
				typed_textFormCallback.probe.backRepoOfInterest,
				typed_text_,
				&rf)

			if reverseFieldOwner != nil {
				pastIdentificationOwner = reverseFieldOwner.(*models.Identification)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastIdentificationOwner != nil {
					idx := slices.Index(pastIdentificationOwner.Relation, typed_text_)
					pastIdentificationOwner.Relation = slices.Delete(pastIdentificationOwner.Relation, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _identification := range *models.GetGongstructInstancesSet[models.Identification](typed_textFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _identification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newIdentificationOwner := _identification // we have a match
						if pastIdentificationOwner != nil {
							if newIdentificationOwner != pastIdentificationOwner {
								idx := slices.Index(pastIdentificationOwner.Relation, typed_text_)
								pastIdentificationOwner.Relation = slices.Delete(pastIdentificationOwner.Relation, idx, idx+1)
								newIdentificationOwner.Relation = append(newIdentificationOwner.Relation, typed_text_)
							}
						} else {
							newIdentificationOwner.Relation = append(newIdentificationOwner.Relation, typed_text_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if typed_textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		typed_text_.Unstage(typed_textFormCallback.probe.stageOfInterest)
	}

	typed_textFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Typed_text](
		typed_textFormCallback.probe,
	)
	typed_textFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if typed_textFormCallback.CreationMode || typed_textFormCallback.formGroup.HasSuppressButtonBeenPressed {
		typed_textFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(typed_textFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Typed_textFormCallback(
			nil,
			typed_textFormCallback.probe,
			newFormGroup,
		)
		typed_text := new(models.Typed_text)
		FillUpForm(typed_text, newFormGroup, typed_textFormCallback.probe)
		typed_textFormCallback.probe.formStage.Commit()
	}

	fillUpTree(typed_textFormCallback.probe)
}
func __gong__New__UnpitchedFormCallback(
	unpitched *models.Unpitched,
	probe *Probe,
	formGroup *table.FormGroup,
) (unpitchedFormCallback *UnpitchedFormCallback) {
	unpitchedFormCallback = new(UnpitchedFormCallback)
	unpitchedFormCallback.probe = probe
	unpitchedFormCallback.unpitched = unpitched
	unpitchedFormCallback.formGroup = formGroup

	unpitchedFormCallback.CreationMode = (unpitched == nil)

	return
}

type UnpitchedFormCallback struct {
	unpitched *models.Unpitched

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (unpitchedFormCallback *UnpitchedFormCallback) OnSave() {

	log.Println("UnpitchedFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	unpitchedFormCallback.probe.formStage.Checkout()

	if unpitchedFormCallback.unpitched == nil {
		unpitchedFormCallback.unpitched = new(models.Unpitched).Stage(unpitchedFormCallback.probe.stageOfInterest)
	}
	unpitched_ := unpitchedFormCallback.unpitched
	_ = unpitched_

	for _, formDiv := range unpitchedFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(unpitched_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if unpitchedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		unpitched_.Unstage(unpitchedFormCallback.probe.stageOfInterest)
	}

	unpitchedFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Unpitched](
		unpitchedFormCallback.probe,
	)
	unpitchedFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if unpitchedFormCallback.CreationMode || unpitchedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		unpitchedFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(unpitchedFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UnpitchedFormCallback(
			nil,
			unpitchedFormCallback.probe,
			newFormGroup,
		)
		unpitched := new(models.Unpitched)
		FillUpForm(unpitched, newFormGroup, unpitchedFormCallback.probe)
		unpitchedFormCallback.probe.formStage.Commit()
	}

	fillUpTree(unpitchedFormCallback.probe)
}
func __gong__New__Virtual_instrumentFormCallback(
	virtual_instrument *models.Virtual_instrument,
	probe *Probe,
	formGroup *table.FormGroup,
) (virtual_instrumentFormCallback *Virtual_instrumentFormCallback) {
	virtual_instrumentFormCallback = new(Virtual_instrumentFormCallback)
	virtual_instrumentFormCallback.probe = probe
	virtual_instrumentFormCallback.virtual_instrument = virtual_instrument
	virtual_instrumentFormCallback.formGroup = formGroup

	virtual_instrumentFormCallback.CreationMode = (virtual_instrument == nil)

	return
}

type Virtual_instrumentFormCallback struct {
	virtual_instrument *models.Virtual_instrument

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (virtual_instrumentFormCallback *Virtual_instrumentFormCallback) OnSave() {

	log.Println("Virtual_instrumentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	virtual_instrumentFormCallback.probe.formStage.Checkout()

	if virtual_instrumentFormCallback.virtual_instrument == nil {
		virtual_instrumentFormCallback.virtual_instrument = new(models.Virtual_instrument).Stage(virtual_instrumentFormCallback.probe.stageOfInterest)
	}
	virtual_instrument_ := virtual_instrumentFormCallback.virtual_instrument
	_ = virtual_instrument_

	for _, formDiv := range virtual_instrumentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(virtual_instrument_.Name), formDiv)
		case "Virtual_library":
			FormDivBasicFieldToField(&(virtual_instrument_.Virtual_library), formDiv)
		case "Virtual_name":
			FormDivBasicFieldToField(&(virtual_instrument_.Virtual_name), formDiv)
		}
	}

	// manage the suppress operation
	if virtual_instrumentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		virtual_instrument_.Unstage(virtual_instrumentFormCallback.probe.stageOfInterest)
	}

	virtual_instrumentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Virtual_instrument](
		virtual_instrumentFormCallback.probe,
	)
	virtual_instrumentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if virtual_instrumentFormCallback.CreationMode || virtual_instrumentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		virtual_instrumentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(virtual_instrumentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Virtual_instrumentFormCallback(
			nil,
			virtual_instrumentFormCallback.probe,
			newFormGroup,
		)
		virtual_instrument := new(models.Virtual_instrument)
		FillUpForm(virtual_instrument, newFormGroup, virtual_instrumentFormCallback.probe)
		virtual_instrumentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(virtual_instrumentFormCallback.probe)
}
func __gong__New__WaitFormCallback(
	wait *models.Wait,
	probe *Probe,
	formGroup *table.FormGroup,
) (waitFormCallback *WaitFormCallback) {
	waitFormCallback = new(WaitFormCallback)
	waitFormCallback.probe = probe
	waitFormCallback.wait = wait
	waitFormCallback.formGroup = formGroup

	waitFormCallback.CreationMode = (wait == nil)

	return
}

type WaitFormCallback struct {
	wait *models.Wait

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (waitFormCallback *WaitFormCallback) OnSave() {

	log.Println("WaitFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	waitFormCallback.probe.formStage.Checkout()

	if waitFormCallback.wait == nil {
		waitFormCallback.wait = new(models.Wait).Stage(waitFormCallback.probe.stageOfInterest)
	}
	wait_ := waitFormCallback.wait
	_ = wait_

	for _, formDiv := range waitFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(wait_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if waitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		wait_.Unstage(waitFormCallback.probe.stageOfInterest)
	}

	waitFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Wait](
		waitFormCallback.probe,
	)
	waitFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if waitFormCallback.CreationMode || waitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		waitFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(waitFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__WaitFormCallback(
			nil,
			waitFormCallback.probe,
			newFormGroup,
		)
		wait := new(models.Wait)
		FillUpForm(wait, newFormGroup, waitFormCallback.probe)
		waitFormCallback.probe.formStage.Commit()
	}

	fillUpTree(waitFormCallback.probe)
}
func __gong__New__Wavy_lineFormCallback(
	wavy_line *models.Wavy_line,
	probe *Probe,
	formGroup *table.FormGroup,
) (wavy_lineFormCallback *Wavy_lineFormCallback) {
	wavy_lineFormCallback = new(Wavy_lineFormCallback)
	wavy_lineFormCallback.probe = probe
	wavy_lineFormCallback.wavy_line = wavy_line
	wavy_lineFormCallback.formGroup = formGroup

	wavy_lineFormCallback.CreationMode = (wavy_line == nil)

	return
}

type Wavy_lineFormCallback struct {
	wavy_line *models.Wavy_line

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (wavy_lineFormCallback *Wavy_lineFormCallback) OnSave() {

	log.Println("Wavy_lineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	wavy_lineFormCallback.probe.formStage.Checkout()

	if wavy_lineFormCallback.wavy_line == nil {
		wavy_lineFormCallback.wavy_line = new(models.Wavy_line).Stage(wavy_lineFormCallback.probe.stageOfInterest)
	}
	wavy_line_ := wavy_lineFormCallback.wavy_line
	_ = wavy_line_

	for _, formDiv := range wavy_lineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(wavy_line_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if wavy_lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		wavy_line_.Unstage(wavy_lineFormCallback.probe.stageOfInterest)
	}

	wavy_lineFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Wavy_line](
		wavy_lineFormCallback.probe,
	)
	wavy_lineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if wavy_lineFormCallback.CreationMode || wavy_lineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		wavy_lineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(wavy_lineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Wavy_lineFormCallback(
			nil,
			wavy_lineFormCallback.probe,
			newFormGroup,
		)
		wavy_line := new(models.Wavy_line)
		FillUpForm(wavy_line, newFormGroup, wavy_lineFormCallback.probe)
		wavy_lineFormCallback.probe.formStage.Commit()
	}

	fillUpTree(wavy_lineFormCallback.probe)
}
func __gong__New__WedgeFormCallback(
	wedge *models.Wedge,
	probe *Probe,
	formGroup *table.FormGroup,
) (wedgeFormCallback *WedgeFormCallback) {
	wedgeFormCallback = new(WedgeFormCallback)
	wedgeFormCallback.probe = probe
	wedgeFormCallback.wedge = wedge
	wedgeFormCallback.formGroup = formGroup

	wedgeFormCallback.CreationMode = (wedge == nil)

	return
}

type WedgeFormCallback struct {
	wedge *models.Wedge

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (wedgeFormCallback *WedgeFormCallback) OnSave() {

	log.Println("WedgeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	wedgeFormCallback.probe.formStage.Checkout()

	if wedgeFormCallback.wedge == nil {
		wedgeFormCallback.wedge = new(models.Wedge).Stage(wedgeFormCallback.probe.stageOfInterest)
	}
	wedge_ := wedgeFormCallback.wedge
	_ = wedge_

	for _, formDiv := range wedgeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(wedge_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if wedgeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		wedge_.Unstage(wedgeFormCallback.probe.stageOfInterest)
	}

	wedgeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Wedge](
		wedgeFormCallback.probe,
	)
	wedgeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if wedgeFormCallback.CreationMode || wedgeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		wedgeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(wedgeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__WedgeFormCallback(
			nil,
			wedgeFormCallback.probe,
			newFormGroup,
		)
		wedge := new(models.Wedge)
		FillUpForm(wedge, newFormGroup, wedgeFormCallback.probe)
		wedgeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(wedgeFormCallback.probe)
}
func __gong__New__WoodFormCallback(
	wood *models.Wood,
	probe *Probe,
	formGroup *table.FormGroup,
) (woodFormCallback *WoodFormCallback) {
	woodFormCallback = new(WoodFormCallback)
	woodFormCallback.probe = probe
	woodFormCallback.wood = wood
	woodFormCallback.formGroup = formGroup

	woodFormCallback.CreationMode = (wood == nil)

	return
}

type WoodFormCallback struct {
	wood *models.Wood

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (woodFormCallback *WoodFormCallback) OnSave() {

	log.Println("WoodFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	woodFormCallback.probe.formStage.Checkout()

	if woodFormCallback.wood == nil {
		woodFormCallback.wood = new(models.Wood).Stage(woodFormCallback.probe.stageOfInterest)
	}
	wood_ := woodFormCallback.wood
	_ = wood_

	for _, formDiv := range woodFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(wood_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if woodFormCallback.formGroup.HasSuppressButtonBeenPressed {
		wood_.Unstage(woodFormCallback.probe.stageOfInterest)
	}

	woodFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Wood](
		woodFormCallback.probe,
	)
	woodFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if woodFormCallback.CreationMode || woodFormCallback.formGroup.HasSuppressButtonBeenPressed {
		woodFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(woodFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__WoodFormCallback(
			nil,
			woodFormCallback.probe,
			newFormGroup,
		)
		wood := new(models.Wood)
		FillUpForm(wood, newFormGroup, woodFormCallback.probe)
		woodFormCallback.probe.formStage.Commit()
	}

	fillUpTree(woodFormCallback.probe)
}
func __gong__New__WorkFormCallback(
	work *models.Work,
	probe *Probe,
	formGroup *table.FormGroup,
) (workFormCallback *WorkFormCallback) {
	workFormCallback = new(WorkFormCallback)
	workFormCallback.probe = probe
	workFormCallback.work = work
	workFormCallback.formGroup = formGroup

	workFormCallback.CreationMode = (work == nil)

	return
}

type WorkFormCallback struct {
	work *models.Work

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (workFormCallback *WorkFormCallback) OnSave() {

	log.Println("WorkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	workFormCallback.probe.formStage.Checkout()

	if workFormCallback.work == nil {
		workFormCallback.work = new(models.Work).Stage(workFormCallback.probe.stageOfInterest)
	}
	work_ := workFormCallback.work
	_ = work_

	for _, formDiv := range workFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(work_.Name), formDiv)
		case "Work_number":
			FormDivBasicFieldToField(&(work_.Work_number), formDiv)
		case "Work_title":
			FormDivBasicFieldToField(&(work_.Work_title), formDiv)
		case "Opus":
			FormDivSelectFieldToField(&(work_.Opus), workFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if workFormCallback.formGroup.HasSuppressButtonBeenPressed {
		work_.Unstage(workFormCallback.probe.stageOfInterest)
	}

	workFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Work](
		workFormCallback.probe,
	)
	workFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if workFormCallback.CreationMode || workFormCallback.formGroup.HasSuppressButtonBeenPressed {
		workFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(workFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__WorkFormCallback(
			nil,
			workFormCallback.probe,
			newFormGroup,
		)
		work := new(models.Work)
		FillUpForm(work, newFormGroup, workFormCallback.probe)
		workFormCallback.probe.formStage.Commit()
	}

	fillUpTree(workFormCallback.probe)
}
