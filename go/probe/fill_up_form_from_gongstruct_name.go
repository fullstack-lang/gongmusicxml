// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmusicxml/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Fo":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Fo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FoFormCallback(
			nil,
			probe,
			formGroup,
		)
		fo := new(models.Fo)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(fo, formGroup, probe)
	}
	formStage.Commit()
}
