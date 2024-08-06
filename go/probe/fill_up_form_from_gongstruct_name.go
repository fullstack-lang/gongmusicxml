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
	case "Bookmark":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bookmark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BookmarkFormCallback(
			nil,
			probe,
			formGroup,
		)
		bookmark := new(models.Bookmark)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bookmark, formGroup, probe)
	case "Foo":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Foo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FooFormCallback(
			nil,
			probe,
			formGroup,
		)
		foo := new(models.Foo)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(foo, formGroup, probe)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		link := new(models.Link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(link, formGroup, probe)
	case "Lyric":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Lyric Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LyricFormCallback(
			nil,
			probe,
			formGroup,
		)
		lyric := new(models.Lyric)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lyric, formGroup, probe)
	case "Lyric_font":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Lyric_font Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_fontFormCallback(
			nil,
			probe,
			formGroup,
		)
		lyric_font := new(models.Lyric_font)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lyric_font, formGroup, probe)
	case "Lyric_language":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Lyric_language Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_languageFormCallback(
			nil,
			probe,
			formGroup,
		)
		lyric_language := new(models.Lyric_language)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lyric_language, formGroup, probe)
	case "Miscellaneous_field":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Miscellaneous_field Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Miscellaneous_fieldFormCallback(
			nil,
			probe,
			formGroup,
		)
		miscellaneous_field := new(models.Miscellaneous_field)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(miscellaneous_field, formGroup, probe)
	}
	formStage.Commit()
}
