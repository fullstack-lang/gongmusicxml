// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmusicxml/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Bookmark:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Bookmark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BookmarkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Foo:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Foo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FooFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lyric:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Lyric Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LyricFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lyric_font:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Lyric_font Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_fontFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lyric_language:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Lyric_language Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_languageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Miscellaneous_field:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Miscellaneous_field Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Miscellaneous_fieldFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
