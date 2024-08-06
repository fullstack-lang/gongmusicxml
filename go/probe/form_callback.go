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
func __gong__New__FoFormCallback(
	fo *models.Fo,
	probe *Probe,
	formGroup *table.FormGroup,
) (foFormCallback *FoFormCallback) {
	foFormCallback = new(FoFormCallback)
	foFormCallback.probe = probe
	foFormCallback.fo = fo
	foFormCallback.formGroup = formGroup

	foFormCallback.CreationMode = (fo == nil)

	return
}

type FoFormCallback struct {
	fo *models.Fo

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (foFormCallback *FoFormCallback) OnSave() {

	log.Println("FoFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	foFormCallback.probe.formStage.Checkout()

	if foFormCallback.fo == nil {
		foFormCallback.fo = new(models.Fo).Stage(foFormCallback.probe.stageOfInterest)
	}
	fo_ := foFormCallback.fo
	_ = fo_

	for _, formDiv := range foFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(fo_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if foFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fo_.Unstage(foFormCallback.probe.stageOfInterest)
	}

	foFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Fo](
		foFormCallback.probe,
	)
	foFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if foFormCallback.CreationMode || foFormCallback.formGroup.HasSuppressButtonBeenPressed {
		foFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(foFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FoFormCallback(
			nil,
			foFormCallback.probe,
			newFormGroup,
		)
		fo := new(models.Fo)
		FillUpForm(fo, newFormGroup, foFormCallback.probe)
		foFormCallback.probe.formStage.Commit()
	}

	fillUpTree(foFormCallback.probe)
}
