// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Accidental_Identifiers := make(map[*Accidental]string)
	_ = map_Accidental_Identifiers

	accidentalOrdered := []*Accidental{}
	for accidental := range stage.Accidentals {
		accidentalOrdered = append(accidentalOrdered, accidental)
	}
	sort.Slice(accidentalOrdered[:], func(i, j int) bool {
		return accidentalOrdered[i].Name < accidentalOrdered[j].Name
	})
	if len(accidentalOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, accidental := range accidentalOrdered {

		id = generatesIdentifier("Accidental", idx, accidental.Name)
		map_Accidental_Identifiers[accidental] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accidental")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", accidental.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(accidental.Name))
		initializerStatements += setValueField

	}

	map_Accidental_mark_Identifiers := make(map[*Accidental_mark]string)
	_ = map_Accidental_mark_Identifiers

	accidental_markOrdered := []*Accidental_mark{}
	for accidental_mark := range stage.Accidental_marks {
		accidental_markOrdered = append(accidental_markOrdered, accidental_mark)
	}
	sort.Slice(accidental_markOrdered[:], func(i, j int) bool {
		return accidental_markOrdered[i].Name < accidental_markOrdered[j].Name
	})
	if len(accidental_markOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, accidental_mark := range accidental_markOrdered {

		id = generatesIdentifier("Accidental_mark", idx, accidental_mark.Name)
		map_Accidental_mark_Identifiers[accidental_mark] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accidental_mark")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", accidental_mark.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(accidental_mark.Name))
		initializerStatements += setValueField

	}

	map_Accidental_text_Identifiers := make(map[*Accidental_text]string)
	_ = map_Accidental_text_Identifiers

	accidental_textOrdered := []*Accidental_text{}
	for accidental_text := range stage.Accidental_texts {
		accidental_textOrdered = append(accidental_textOrdered, accidental_text)
	}
	sort.Slice(accidental_textOrdered[:], func(i, j int) bool {
		return accidental_textOrdered[i].Name < accidental_textOrdered[j].Name
	})
	if len(accidental_textOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, accidental_text := range accidental_textOrdered {

		id = generatesIdentifier("Accidental_text", idx, accidental_text.Name)
		map_Accidental_text_Identifiers[accidental_text] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accidental_text")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", accidental_text.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(accidental_text.Name))
		initializerStatements += setValueField

	}

	map_Accord_Identifiers := make(map[*Accord]string)
	_ = map_Accord_Identifiers

	accordOrdered := []*Accord{}
	for accord := range stage.Accords {
		accordOrdered = append(accordOrdered, accord)
	}
	sort.Slice(accordOrdered[:], func(i, j int) bool {
		return accordOrdered[i].Name < accordOrdered[j].Name
	})
	if len(accordOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, accord := range accordOrdered {

		id = generatesIdentifier("Accord", idx, accord.Name)
		map_Accord_Identifiers[accord] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accord")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", accord.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(accord.Name))
		initializerStatements += setValueField

	}

	map_Accordion_registration_Identifiers := make(map[*Accordion_registration]string)
	_ = map_Accordion_registration_Identifiers

	accordion_registrationOrdered := []*Accordion_registration{}
	for accordion_registration := range stage.Accordion_registrations {
		accordion_registrationOrdered = append(accordion_registrationOrdered, accordion_registration)
	}
	sort.Slice(accordion_registrationOrdered[:], func(i, j int) bool {
		return accordion_registrationOrdered[i].Name < accordion_registrationOrdered[j].Name
	})
	if len(accordion_registrationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, accordion_registration := range accordion_registrationOrdered {

		id = generatesIdentifier("Accordion_registration", idx, accordion_registration.Name)
		map_Accordion_registration_Identifiers[accordion_registration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accordion_registration")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", accordion_registration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(accordion_registration.Name))
		initializerStatements += setValueField

	}

	map_AnyType_Identifiers := make(map[*AnyType]string)
	_ = map_AnyType_Identifiers

	anytypeOrdered := []*AnyType{}
	for anytype := range stage.AnyTypes {
		anytypeOrdered = append(anytypeOrdered, anytype)
	}
	sort.Slice(anytypeOrdered[:], func(i, j int) bool {
		return anytypeOrdered[i].Name < anytypeOrdered[j].Name
	})
	if len(anytypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, anytype := range anytypeOrdered {

		id = generatesIdentifier("AnyType", idx, anytype.Name)
		map_AnyType_Identifiers[anytype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AnyType")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", anytype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(anytype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InnerXML")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(anytype.InnerXML))
		initializerStatements += setValueField

	}

	map_Appearance_Identifiers := make(map[*Appearance]string)
	_ = map_Appearance_Identifiers

	appearanceOrdered := []*Appearance{}
	for appearance := range stage.Appearances {
		appearanceOrdered = append(appearanceOrdered, appearance)
	}
	sort.Slice(appearanceOrdered[:], func(i, j int) bool {
		return appearanceOrdered[i].Name < appearanceOrdered[j].Name
	})
	if len(appearanceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, appearance := range appearanceOrdered {

		id = generatesIdentifier("Appearance", idx, appearance.Name)
		map_Appearance_Identifiers[appearance] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Appearance")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", appearance.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(appearance.Name))
		initializerStatements += setValueField

	}

	map_Arpeggiate_Identifiers := make(map[*Arpeggiate]string)
	_ = map_Arpeggiate_Identifiers

	arpeggiateOrdered := []*Arpeggiate{}
	for arpeggiate := range stage.Arpeggiates {
		arpeggiateOrdered = append(arpeggiateOrdered, arpeggiate)
	}
	sort.Slice(arpeggiateOrdered[:], func(i, j int) bool {
		return arpeggiateOrdered[i].Name < arpeggiateOrdered[j].Name
	})
	if len(arpeggiateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, arpeggiate := range arpeggiateOrdered {

		id = generatesIdentifier("Arpeggiate", idx, arpeggiate.Name)
		map_Arpeggiate_Identifiers[arpeggiate] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Arpeggiate")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", arpeggiate.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(arpeggiate.Name))
		initializerStatements += setValueField

	}

	map_Arrow_Identifiers := make(map[*Arrow]string)
	_ = map_Arrow_Identifiers

	arrowOrdered := []*Arrow{}
	for arrow := range stage.Arrows {
		arrowOrdered = append(arrowOrdered, arrow)
	}
	sort.Slice(arrowOrdered[:], func(i, j int) bool {
		return arrowOrdered[i].Name < arrowOrdered[j].Name
	})
	if len(arrowOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, arrow := range arrowOrdered {

		id = generatesIdentifier("Arrow", idx, arrow.Name)
		map_Arrow_Identifiers[arrow] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Arrow")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", arrow.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(arrow.Name))
		initializerStatements += setValueField

	}

	map_Articulations_Identifiers := make(map[*Articulations]string)
	_ = map_Articulations_Identifiers

	articulationsOrdered := []*Articulations{}
	for articulations := range stage.Articulationss {
		articulationsOrdered = append(articulationsOrdered, articulations)
	}
	sort.Slice(articulationsOrdered[:], func(i, j int) bool {
		return articulationsOrdered[i].Name < articulationsOrdered[j].Name
	})
	if len(articulationsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, articulations := range articulationsOrdered {

		id = generatesIdentifier("Articulations", idx, articulations.Name)
		map_Articulations_Identifiers[articulations] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Articulations")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", articulations.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(articulations.Name))
		initializerStatements += setValueField

	}

	map_Assess_Identifiers := make(map[*Assess]string)
	_ = map_Assess_Identifiers

	assessOrdered := []*Assess{}
	for assess := range stage.Assesss {
		assessOrdered = append(assessOrdered, assess)
	}
	sort.Slice(assessOrdered[:], func(i, j int) bool {
		return assessOrdered[i].Name < assessOrdered[j].Name
	})
	if len(assessOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, assess := range assessOrdered {

		id = generatesIdentifier("Assess", idx, assess.Name)
		map_Assess_Identifiers[assess] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Assess")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", assess.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assess.Name))
		initializerStatements += setValueField

	}

	map_Attributes_Identifiers := make(map[*Attributes]string)
	_ = map_Attributes_Identifiers

	attributesOrdered := []*Attributes{}
	for attributes := range stage.Attributess {
		attributesOrdered = append(attributesOrdered, attributes)
	}
	sort.Slice(attributesOrdered[:], func(i, j int) bool {
		return attributesOrdered[i].Name < attributesOrdered[j].Name
	})
	if len(attributesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributes := range attributesOrdered {

		id = generatesIdentifier("Attributes", idx, attributes.Name)
		map_Attributes_Identifiers[attributes] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Attributes")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributes.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributes.Name))
		initializerStatements += setValueField

	}

	map_Backup_Identifiers := make(map[*Backup]string)
	_ = map_Backup_Identifiers

	backupOrdered := []*Backup{}
	for backup := range stage.Backups {
		backupOrdered = append(backupOrdered, backup)
	}
	sort.Slice(backupOrdered[:], func(i, j int) bool {
		return backupOrdered[i].Name < backupOrdered[j].Name
	})
	if len(backupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, backup := range backupOrdered {

		id = generatesIdentifier("Backup", idx, backup.Name)
		map_Backup_Identifiers[backup] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Backup")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", backup.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(backup.Name))
		initializerStatements += setValueField

	}

	map_Bar_style_color_Identifiers := make(map[*Bar_style_color]string)
	_ = map_Bar_style_color_Identifiers

	bar_style_colorOrdered := []*Bar_style_color{}
	for bar_style_color := range stage.Bar_style_colors {
		bar_style_colorOrdered = append(bar_style_colorOrdered, bar_style_color)
	}
	sort.Slice(bar_style_colorOrdered[:], func(i, j int) bool {
		return bar_style_colorOrdered[i].Name < bar_style_colorOrdered[j].Name
	})
	if len(bar_style_colorOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, bar_style_color := range bar_style_colorOrdered {

		id = generatesIdentifier("Bar_style_color", idx, bar_style_color.Name)
		map_Bar_style_color_Identifiers[bar_style_color] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bar_style_color")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bar_style_color.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bar_style_color.Name))
		initializerStatements += setValueField

	}

	map_Barline_Identifiers := make(map[*Barline]string)
	_ = map_Barline_Identifiers

	barlineOrdered := []*Barline{}
	for barline := range stage.Barlines {
		barlineOrdered = append(barlineOrdered, barline)
	}
	sort.Slice(barlineOrdered[:], func(i, j int) bool {
		return barlineOrdered[i].Name < barlineOrdered[j].Name
	})
	if len(barlineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, barline := range barlineOrdered {

		id = generatesIdentifier("Barline", idx, barline.Name)
		map_Barline_Identifiers[barline] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Barline")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", barline.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(barline.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Segno")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(barline.Segno))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Coda")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(barline.Coda))
		initializerStatements += setValueField

	}

	map_Barre_Identifiers := make(map[*Barre]string)
	_ = map_Barre_Identifiers

	barreOrdered := []*Barre{}
	for barre := range stage.Barres {
		barreOrdered = append(barreOrdered, barre)
	}
	sort.Slice(barreOrdered[:], func(i, j int) bool {
		return barreOrdered[i].Name < barreOrdered[j].Name
	})
	if len(barreOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, barre := range barreOrdered {

		id = generatesIdentifier("Barre", idx, barre.Name)
		map_Barre_Identifiers[barre] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Barre")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", barre.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(barre.Name))
		initializerStatements += setValueField

	}

	map_Bass_Identifiers := make(map[*Bass]string)
	_ = map_Bass_Identifiers

	bassOrdered := []*Bass{}
	for bass := range stage.Basss {
		bassOrdered = append(bassOrdered, bass)
	}
	sort.Slice(bassOrdered[:], func(i, j int) bool {
		return bassOrdered[i].Name < bassOrdered[j].Name
	})
	if len(bassOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, bass := range bassOrdered {

		id = generatesIdentifier("Bass", idx, bass.Name)
		map_Bass_Identifiers[bass] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bass")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bass.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bass.Name))
		initializerStatements += setValueField

	}

	map_Bass_step_Identifiers := make(map[*Bass_step]string)
	_ = map_Bass_step_Identifiers

	bass_stepOrdered := []*Bass_step{}
	for bass_step := range stage.Bass_steps {
		bass_stepOrdered = append(bass_stepOrdered, bass_step)
	}
	sort.Slice(bass_stepOrdered[:], func(i, j int) bool {
		return bass_stepOrdered[i].Name < bass_stepOrdered[j].Name
	})
	if len(bass_stepOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, bass_step := range bass_stepOrdered {

		id = generatesIdentifier("Bass_step", idx, bass_step.Name)
		map_Bass_step_Identifiers[bass_step] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bass_step")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bass_step.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bass_step.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bass_step.Text))
		initializerStatements += setValueField

	}

	map_Beam_Identifiers := make(map[*Beam]string)
	_ = map_Beam_Identifiers

	beamOrdered := []*Beam{}
	for beam := range stage.Beams {
		beamOrdered = append(beamOrdered, beam)
	}
	sort.Slice(beamOrdered[:], func(i, j int) bool {
		return beamOrdered[i].Name < beamOrdered[j].Name
	})
	if len(beamOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, beam := range beamOrdered {

		id = generatesIdentifier("Beam", idx, beam.Name)
		map_Beam_Identifiers[beam] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Beam")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", beam.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(beam.Name))
		initializerStatements += setValueField

	}

	map_Beat_repeat_Identifiers := make(map[*Beat_repeat]string)
	_ = map_Beat_repeat_Identifiers

	beat_repeatOrdered := []*Beat_repeat{}
	for beat_repeat := range stage.Beat_repeats {
		beat_repeatOrdered = append(beat_repeatOrdered, beat_repeat)
	}
	sort.Slice(beat_repeatOrdered[:], func(i, j int) bool {
		return beat_repeatOrdered[i].Name < beat_repeatOrdered[j].Name
	})
	if len(beat_repeatOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, beat_repeat := range beat_repeatOrdered {

		id = generatesIdentifier("Beat_repeat", idx, beat_repeat.Name)
		map_Beat_repeat_Identifiers[beat_repeat] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Beat_repeat")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", beat_repeat.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(beat_repeat.Name))
		initializerStatements += setValueField

	}

	map_Beat_unit_tied_Identifiers := make(map[*Beat_unit_tied]string)
	_ = map_Beat_unit_tied_Identifiers

	beat_unit_tiedOrdered := []*Beat_unit_tied{}
	for beat_unit_tied := range stage.Beat_unit_tieds {
		beat_unit_tiedOrdered = append(beat_unit_tiedOrdered, beat_unit_tied)
	}
	sort.Slice(beat_unit_tiedOrdered[:], func(i, j int) bool {
		return beat_unit_tiedOrdered[i].Name < beat_unit_tiedOrdered[j].Name
	})
	if len(beat_unit_tiedOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, beat_unit_tied := range beat_unit_tiedOrdered {

		id = generatesIdentifier("Beat_unit_tied", idx, beat_unit_tied.Name)
		map_Beat_unit_tied_Identifiers[beat_unit_tied] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Beat_unit_tied")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", beat_unit_tied.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(beat_unit_tied.Name))
		initializerStatements += setValueField

	}

	map_Beater_Identifiers := make(map[*Beater]string)
	_ = map_Beater_Identifiers

	beaterOrdered := []*Beater{}
	for beater := range stage.Beaters {
		beaterOrdered = append(beaterOrdered, beater)
	}
	sort.Slice(beaterOrdered[:], func(i, j int) bool {
		return beaterOrdered[i].Name < beaterOrdered[j].Name
	})
	if len(beaterOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, beater := range beaterOrdered {

		id = generatesIdentifier("Beater", idx, beater.Name)
		map_Beater_Identifiers[beater] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Beater")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", beater.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(beater.Name))
		initializerStatements += setValueField

	}

	map_Bend_Identifiers := make(map[*Bend]string)
	_ = map_Bend_Identifiers

	bendOrdered := []*Bend{}
	for bend := range stage.Bends {
		bendOrdered = append(bendOrdered, bend)
	}
	sort.Slice(bendOrdered[:], func(i, j int) bool {
		return bendOrdered[i].Name < bendOrdered[j].Name
	})
	if len(bendOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, bend := range bendOrdered {

		id = generatesIdentifier("Bend", idx, bend.Name)
		map_Bend_Identifiers[bend] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bend")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bend.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bend.Name))
		initializerStatements += setValueField

	}

	map_Bookmark_Identifiers := make(map[*Bookmark]string)
	_ = map_Bookmark_Identifiers

	bookmarkOrdered := []*Bookmark{}
	for bookmark := range stage.Bookmarks {
		bookmarkOrdered = append(bookmarkOrdered, bookmark)
	}
	sort.Slice(bookmarkOrdered[:], func(i, j int) bool {
		return bookmarkOrdered[i].Name < bookmarkOrdered[j].Name
	})
	if len(bookmarkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, bookmark := range bookmarkOrdered {

		id = generatesIdentifier("Bookmark", idx, bookmark.Name)
		map_Bookmark_Identifiers[bookmark] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bookmark")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bookmark.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bookmark.Name))
		initializerStatements += setValueField

	}

	map_Bracket_Identifiers := make(map[*Bracket]string)
	_ = map_Bracket_Identifiers

	bracketOrdered := []*Bracket{}
	for bracket := range stage.Brackets {
		bracketOrdered = append(bracketOrdered, bracket)
	}
	sort.Slice(bracketOrdered[:], func(i, j int) bool {
		return bracketOrdered[i].Name < bracketOrdered[j].Name
	})
	if len(bracketOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, bracket := range bracketOrdered {

		id = generatesIdentifier("Bracket", idx, bracket.Name)
		map_Bracket_Identifiers[bracket] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bracket")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bracket.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bracket.Name))
		initializerStatements += setValueField

	}

	map_Breath_mark_Identifiers := make(map[*Breath_mark]string)
	_ = map_Breath_mark_Identifiers

	breath_markOrdered := []*Breath_mark{}
	for breath_mark := range stage.Breath_marks {
		breath_markOrdered = append(breath_markOrdered, breath_mark)
	}
	sort.Slice(breath_markOrdered[:], func(i, j int) bool {
		return breath_markOrdered[i].Name < breath_markOrdered[j].Name
	})
	if len(breath_markOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, breath_mark := range breath_markOrdered {

		id = generatesIdentifier("Breath_mark", idx, breath_mark.Name)
		map_Breath_mark_Identifiers[breath_mark] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Breath_mark")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", breath_mark.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(breath_mark.Name))
		initializerStatements += setValueField

	}

	map_Caesura_Identifiers := make(map[*Caesura]string)
	_ = map_Caesura_Identifiers

	caesuraOrdered := []*Caesura{}
	for caesura := range stage.Caesuras {
		caesuraOrdered = append(caesuraOrdered, caesura)
	}
	sort.Slice(caesuraOrdered[:], func(i, j int) bool {
		return caesuraOrdered[i].Name < caesuraOrdered[j].Name
	})
	if len(caesuraOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, caesura := range caesuraOrdered {

		id = generatesIdentifier("Caesura", idx, caesura.Name)
		map_Caesura_Identifiers[caesura] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Caesura")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", caesura.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(caesura.Name))
		initializerStatements += setValueField

	}

	map_Cancel_Identifiers := make(map[*Cancel]string)
	_ = map_Cancel_Identifiers

	cancelOrdered := []*Cancel{}
	for cancel := range stage.Cancels {
		cancelOrdered = append(cancelOrdered, cancel)
	}
	sort.Slice(cancelOrdered[:], func(i, j int) bool {
		return cancelOrdered[i].Name < cancelOrdered[j].Name
	})
	if len(cancelOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, cancel := range cancelOrdered {

		id = generatesIdentifier("Cancel", idx, cancel.Name)
		map_Cancel_Identifiers[cancel] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cancel")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cancel.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cancel.Name))
		initializerStatements += setValueField

	}

	map_Clef_Identifiers := make(map[*Clef]string)
	_ = map_Clef_Identifiers

	clefOrdered := []*Clef{}
	for clef := range stage.Clefs {
		clefOrdered = append(clefOrdered, clef)
	}
	sort.Slice(clefOrdered[:], func(i, j int) bool {
		return clefOrdered[i].Name < clefOrdered[j].Name
	})
	if len(clefOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, clef := range clefOrdered {

		id = generatesIdentifier("Clef", idx, clef.Name)
		map_Clef_Identifiers[clef] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Clef")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", clef.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(clef.Name))
		initializerStatements += setValueField

	}

	map_Coda_Identifiers := make(map[*Coda]string)
	_ = map_Coda_Identifiers

	codaOrdered := []*Coda{}
	for coda := range stage.Codas {
		codaOrdered = append(codaOrdered, coda)
	}
	sort.Slice(codaOrdered[:], func(i, j int) bool {
		return codaOrdered[i].Name < codaOrdered[j].Name
	})
	if len(codaOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, coda := range codaOrdered {

		id = generatesIdentifier("Coda", idx, coda.Name)
		map_Coda_Identifiers[coda] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Coda")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", coda.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(coda.Name))
		initializerStatements += setValueField

	}

	map_Credit_Identifiers := make(map[*Credit]string)
	_ = map_Credit_Identifiers

	creditOrdered := []*Credit{}
	for credit := range stage.Credits {
		creditOrdered = append(creditOrdered, credit)
	}
	sort.Slice(creditOrdered[:], func(i, j int) bool {
		return creditOrdered[i].Name < creditOrdered[j].Name
	})
	if len(creditOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, credit := range creditOrdered {

		id = generatesIdentifier("Credit", idx, credit.Name)
		map_Credit_Identifiers[credit] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Credit")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", credit.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(credit.Name))
		initializerStatements += setValueField

	}

	map_Dashes_Identifiers := make(map[*Dashes]string)
	_ = map_Dashes_Identifiers

	dashesOrdered := []*Dashes{}
	for dashes := range stage.Dashess {
		dashesOrdered = append(dashesOrdered, dashes)
	}
	sort.Slice(dashesOrdered[:], func(i, j int) bool {
		return dashesOrdered[i].Name < dashesOrdered[j].Name
	})
	if len(dashesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, dashes := range dashesOrdered {

		id = generatesIdentifier("Dashes", idx, dashes.Name)
		map_Dashes_Identifiers[dashes] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Dashes")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dashes.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dashes.Name))
		initializerStatements += setValueField

	}

	map_Defaults_Identifiers := make(map[*Defaults]string)
	_ = map_Defaults_Identifiers

	defaultsOrdered := []*Defaults{}
	for defaults := range stage.Defaultss {
		defaultsOrdered = append(defaultsOrdered, defaults)
	}
	sort.Slice(defaultsOrdered[:], func(i, j int) bool {
		return defaultsOrdered[i].Name < defaultsOrdered[j].Name
	})
	if len(defaultsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, defaults := range defaultsOrdered {

		id = generatesIdentifier("Defaults", idx, defaults.Name)
		map_Defaults_Identifiers[defaults] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Defaults")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", defaults.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(defaults.Name))
		initializerStatements += setValueField

	}

	map_Degree_Identifiers := make(map[*Degree]string)
	_ = map_Degree_Identifiers

	degreeOrdered := []*Degree{}
	for degree := range stage.Degrees {
		degreeOrdered = append(degreeOrdered, degree)
	}
	sort.Slice(degreeOrdered[:], func(i, j int) bool {
		return degreeOrdered[i].Name < degreeOrdered[j].Name
	})
	if len(degreeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, degree := range degreeOrdered {

		id = generatesIdentifier("Degree", idx, degree.Name)
		map_Degree_Identifiers[degree] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Degree")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", degree.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(degree.Name))
		initializerStatements += setValueField

	}

	map_Degree_alter_Identifiers := make(map[*Degree_alter]string)
	_ = map_Degree_alter_Identifiers

	degree_alterOrdered := []*Degree_alter{}
	for degree_alter := range stage.Degree_alters {
		degree_alterOrdered = append(degree_alterOrdered, degree_alter)
	}
	sort.Slice(degree_alterOrdered[:], func(i, j int) bool {
		return degree_alterOrdered[i].Name < degree_alterOrdered[j].Name
	})
	if len(degree_alterOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, degree_alter := range degree_alterOrdered {

		id = generatesIdentifier("Degree_alter", idx, degree_alter.Name)
		map_Degree_alter_Identifiers[degree_alter] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Degree_alter")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", degree_alter.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(degree_alter.Name))
		initializerStatements += setValueField

	}

	map_Degree_type_Identifiers := make(map[*Degree_type]string)
	_ = map_Degree_type_Identifiers

	degree_typeOrdered := []*Degree_type{}
	for degree_type := range stage.Degree_types {
		degree_typeOrdered = append(degree_typeOrdered, degree_type)
	}
	sort.Slice(degree_typeOrdered[:], func(i, j int) bool {
		return degree_typeOrdered[i].Name < degree_typeOrdered[j].Name
	})
	if len(degree_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, degree_type := range degree_typeOrdered {

		id = generatesIdentifier("Degree_type", idx, degree_type.Name)
		map_Degree_type_Identifiers[degree_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Degree_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", degree_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(degree_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(degree_type.Text))
		initializerStatements += setValueField

	}

	map_Degree_value_Identifiers := make(map[*Degree_value]string)
	_ = map_Degree_value_Identifiers

	degree_valueOrdered := []*Degree_value{}
	for degree_value := range stage.Degree_values {
		degree_valueOrdered = append(degree_valueOrdered, degree_value)
	}
	sort.Slice(degree_valueOrdered[:], func(i, j int) bool {
		return degree_valueOrdered[i].Name < degree_valueOrdered[j].Name
	})
	if len(degree_valueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, degree_value := range degree_valueOrdered {

		id = generatesIdentifier("Degree_value", idx, degree_value.Name)
		map_Degree_value_Identifiers[degree_value] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Degree_value")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", degree_value.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(degree_value.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(degree_value.Text))
		initializerStatements += setValueField

	}

	map_Direction_Identifiers := make(map[*Direction]string)
	_ = map_Direction_Identifiers

	directionOrdered := []*Direction{}
	for direction := range stage.Directions {
		directionOrdered = append(directionOrdered, direction)
	}
	sort.Slice(directionOrdered[:], func(i, j int) bool {
		return directionOrdered[i].Name < directionOrdered[j].Name
	})
	if len(directionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, direction := range directionOrdered {

		id = generatesIdentifier("Direction", idx, direction.Name)
		map_Direction_Identifiers[direction] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Direction")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", direction.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(direction.Name))
		initializerStatements += setValueField

	}

	map_Direction_type_Identifiers := make(map[*Direction_type]string)
	_ = map_Direction_type_Identifiers

	direction_typeOrdered := []*Direction_type{}
	for direction_type := range stage.Direction_types {
		direction_typeOrdered = append(direction_typeOrdered, direction_type)
	}
	sort.Slice(direction_typeOrdered[:], func(i, j int) bool {
		return direction_typeOrdered[i].Name < direction_typeOrdered[j].Name
	})
	if len(direction_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, direction_type := range direction_typeOrdered {

		id = generatesIdentifier("Direction_type", idx, direction_type.Name)
		map_Direction_type_Identifiers[direction_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Direction_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", direction_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(direction_type.Name))
		initializerStatements += setValueField

	}

	map_Distance_Identifiers := make(map[*Distance]string)
	_ = map_Distance_Identifiers

	distanceOrdered := []*Distance{}
	for distance := range stage.Distances {
		distanceOrdered = append(distanceOrdered, distance)
	}
	sort.Slice(distanceOrdered[:], func(i, j int) bool {
		return distanceOrdered[i].Name < distanceOrdered[j].Name
	})
	if len(distanceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, distance := range distanceOrdered {

		id = generatesIdentifier("Distance", idx, distance.Name)
		map_Distance_Identifiers[distance] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Distance")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", distance.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(distance.Name))
		initializerStatements += setValueField

	}

	map_Double_Identifiers := make(map[*Double]string)
	_ = map_Double_Identifiers

	doubleOrdered := []*Double{}
	for double := range stage.Doubles {
		doubleOrdered = append(doubleOrdered, double)
	}
	sort.Slice(doubleOrdered[:], func(i, j int) bool {
		return doubleOrdered[i].Name < doubleOrdered[j].Name
	})
	if len(doubleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, double := range doubleOrdered {

		id = generatesIdentifier("Double", idx, double.Name)
		map_Double_Identifiers[double] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Double")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", double.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(double.Name))
		initializerStatements += setValueField

	}

	map_Dynamics_Identifiers := make(map[*Dynamics]string)
	_ = map_Dynamics_Identifiers

	dynamicsOrdered := []*Dynamics{}
	for dynamics := range stage.Dynamicss {
		dynamicsOrdered = append(dynamicsOrdered, dynamics)
	}
	sort.Slice(dynamicsOrdered[:], func(i, j int) bool {
		return dynamicsOrdered[i].Name < dynamicsOrdered[j].Name
	})
	if len(dynamicsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, dynamics := range dynamicsOrdered {

		id = generatesIdentifier("Dynamics", idx, dynamics.Name)
		map_Dynamics_Identifiers[dynamics] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Dynamics")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dynamics.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dynamics.Name))
		initializerStatements += setValueField

	}

	map_Effect_Identifiers := make(map[*Effect]string)
	_ = map_Effect_Identifiers

	effectOrdered := []*Effect{}
	for effect := range stage.Effects {
		effectOrdered = append(effectOrdered, effect)
	}
	sort.Slice(effectOrdered[:], func(i, j int) bool {
		return effectOrdered[i].Name < effectOrdered[j].Name
	})
	if len(effectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, effect := range effectOrdered {

		id = generatesIdentifier("Effect", idx, effect.Name)
		map_Effect_Identifiers[effect] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Effect")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", effect.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(effect.Name))
		initializerStatements += setValueField

	}

	map_Elision_Identifiers := make(map[*Elision]string)
	_ = map_Elision_Identifiers

	elisionOrdered := []*Elision{}
	for elision := range stage.Elisions {
		elisionOrdered = append(elisionOrdered, elision)
	}
	sort.Slice(elisionOrdered[:], func(i, j int) bool {
		return elisionOrdered[i].Name < elisionOrdered[j].Name
	})
	if len(elisionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, elision := range elisionOrdered {

		id = generatesIdentifier("Elision", idx, elision.Name)
		map_Elision_Identifiers[elision] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Elision")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", elision.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(elision.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(elision.Value))
		initializerStatements += setValueField

	}

	map_Empty_Identifiers := make(map[*Empty]string)
	_ = map_Empty_Identifiers

	emptyOrdered := []*Empty{}
	for empty := range stage.Emptys {
		emptyOrdered = append(emptyOrdered, empty)
	}
	sort.Slice(emptyOrdered[:], func(i, j int) bool {
		return emptyOrdered[i].Name < emptyOrdered[j].Name
	})
	if len(emptyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty := range emptyOrdered {

		id = generatesIdentifier("Empty", idx, empty.Name)
		map_Empty_Identifiers[empty] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty.Name))
		initializerStatements += setValueField

	}

	map_Empty_font_Identifiers := make(map[*Empty_font]string)
	_ = map_Empty_font_Identifiers

	empty_fontOrdered := []*Empty_font{}
	for empty_font := range stage.Empty_fonts {
		empty_fontOrdered = append(empty_fontOrdered, empty_font)
	}
	sort.Slice(empty_fontOrdered[:], func(i, j int) bool {
		return empty_fontOrdered[i].Name < empty_fontOrdered[j].Name
	})
	if len(empty_fontOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty_font := range empty_fontOrdered {

		id = generatesIdentifier("Empty_font", idx, empty_font.Name)
		map_Empty_font_Identifiers[empty_font] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_font")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty_font.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty_font.Name))
		initializerStatements += setValueField

	}

	map_Empty_line_Identifiers := make(map[*Empty_line]string)
	_ = map_Empty_line_Identifiers

	empty_lineOrdered := []*Empty_line{}
	for empty_line := range stage.Empty_lines {
		empty_lineOrdered = append(empty_lineOrdered, empty_line)
	}
	sort.Slice(empty_lineOrdered[:], func(i, j int) bool {
		return empty_lineOrdered[i].Name < empty_lineOrdered[j].Name
	})
	if len(empty_lineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty_line := range empty_lineOrdered {

		id = generatesIdentifier("Empty_line", idx, empty_line.Name)
		map_Empty_line_Identifiers[empty_line] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_line")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty_line.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty_line.Name))
		initializerStatements += setValueField

	}

	map_Empty_placement_Identifiers := make(map[*Empty_placement]string)
	_ = map_Empty_placement_Identifiers

	empty_placementOrdered := []*Empty_placement{}
	for empty_placement := range stage.Empty_placements {
		empty_placementOrdered = append(empty_placementOrdered, empty_placement)
	}
	sort.Slice(empty_placementOrdered[:], func(i, j int) bool {
		return empty_placementOrdered[i].Name < empty_placementOrdered[j].Name
	})
	if len(empty_placementOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty_placement := range empty_placementOrdered {

		id = generatesIdentifier("Empty_placement", idx, empty_placement.Name)
		map_Empty_placement_Identifiers[empty_placement] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_placement")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty_placement.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty_placement.Name))
		initializerStatements += setValueField

	}

	map_Empty_placement_smufl_Identifiers := make(map[*Empty_placement_smufl]string)
	_ = map_Empty_placement_smufl_Identifiers

	empty_placement_smuflOrdered := []*Empty_placement_smufl{}
	for empty_placement_smufl := range stage.Empty_placement_smufls {
		empty_placement_smuflOrdered = append(empty_placement_smuflOrdered, empty_placement_smufl)
	}
	sort.Slice(empty_placement_smuflOrdered[:], func(i, j int) bool {
		return empty_placement_smuflOrdered[i].Name < empty_placement_smuflOrdered[j].Name
	})
	if len(empty_placement_smuflOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty_placement_smufl := range empty_placement_smuflOrdered {

		id = generatesIdentifier("Empty_placement_smufl", idx, empty_placement_smufl.Name)
		map_Empty_placement_smufl_Identifiers[empty_placement_smufl] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_placement_smufl")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty_placement_smufl.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty_placement_smufl.Name))
		initializerStatements += setValueField

	}

	map_Empty_print_object_style_align_Identifiers := make(map[*Empty_print_object_style_align]string)
	_ = map_Empty_print_object_style_align_Identifiers

	empty_print_object_style_alignOrdered := []*Empty_print_object_style_align{}
	for empty_print_object_style_align := range stage.Empty_print_object_style_aligns {
		empty_print_object_style_alignOrdered = append(empty_print_object_style_alignOrdered, empty_print_object_style_align)
	}
	sort.Slice(empty_print_object_style_alignOrdered[:], func(i, j int) bool {
		return empty_print_object_style_alignOrdered[i].Name < empty_print_object_style_alignOrdered[j].Name
	})
	if len(empty_print_object_style_alignOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty_print_object_style_align := range empty_print_object_style_alignOrdered {

		id = generatesIdentifier("Empty_print_object_style_align", idx, empty_print_object_style_align.Name)
		map_Empty_print_object_style_align_Identifiers[empty_print_object_style_align] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_print_object_style_align")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty_print_object_style_align.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty_print_object_style_align.Name))
		initializerStatements += setValueField

	}

	map_Empty_print_style_Identifiers := make(map[*Empty_print_style]string)
	_ = map_Empty_print_style_Identifiers

	empty_print_styleOrdered := []*Empty_print_style{}
	for empty_print_style := range stage.Empty_print_styles {
		empty_print_styleOrdered = append(empty_print_styleOrdered, empty_print_style)
	}
	sort.Slice(empty_print_styleOrdered[:], func(i, j int) bool {
		return empty_print_styleOrdered[i].Name < empty_print_styleOrdered[j].Name
	})
	if len(empty_print_styleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty_print_style := range empty_print_styleOrdered {

		id = generatesIdentifier("Empty_print_style", idx, empty_print_style.Name)
		map_Empty_print_style_Identifiers[empty_print_style] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_print_style")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty_print_style.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty_print_style.Name))
		initializerStatements += setValueField

	}

	map_Empty_print_style_align_Identifiers := make(map[*Empty_print_style_align]string)
	_ = map_Empty_print_style_align_Identifiers

	empty_print_style_alignOrdered := []*Empty_print_style_align{}
	for empty_print_style_align := range stage.Empty_print_style_aligns {
		empty_print_style_alignOrdered = append(empty_print_style_alignOrdered, empty_print_style_align)
	}
	sort.Slice(empty_print_style_alignOrdered[:], func(i, j int) bool {
		return empty_print_style_alignOrdered[i].Name < empty_print_style_alignOrdered[j].Name
	})
	if len(empty_print_style_alignOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty_print_style_align := range empty_print_style_alignOrdered {

		id = generatesIdentifier("Empty_print_style_align", idx, empty_print_style_align.Name)
		map_Empty_print_style_align_Identifiers[empty_print_style_align] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_print_style_align")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty_print_style_align.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty_print_style_align.Name))
		initializerStatements += setValueField

	}

	map_Empty_print_style_align_id_Identifiers := make(map[*Empty_print_style_align_id]string)
	_ = map_Empty_print_style_align_id_Identifiers

	empty_print_style_align_idOrdered := []*Empty_print_style_align_id{}
	for empty_print_style_align_id := range stage.Empty_print_style_align_ids {
		empty_print_style_align_idOrdered = append(empty_print_style_align_idOrdered, empty_print_style_align_id)
	}
	sort.Slice(empty_print_style_align_idOrdered[:], func(i, j int) bool {
		return empty_print_style_align_idOrdered[i].Name < empty_print_style_align_idOrdered[j].Name
	})
	if len(empty_print_style_align_idOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty_print_style_align_id := range empty_print_style_align_idOrdered {

		id = generatesIdentifier("Empty_print_style_align_id", idx, empty_print_style_align_id.Name)
		map_Empty_print_style_align_id_Identifiers[empty_print_style_align_id] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_print_style_align_id")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty_print_style_align_id.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty_print_style_align_id.Name))
		initializerStatements += setValueField

	}

	map_Empty_trill_sound_Identifiers := make(map[*Empty_trill_sound]string)
	_ = map_Empty_trill_sound_Identifiers

	empty_trill_soundOrdered := []*Empty_trill_sound{}
	for empty_trill_sound := range stage.Empty_trill_sounds {
		empty_trill_soundOrdered = append(empty_trill_soundOrdered, empty_trill_sound)
	}
	sort.Slice(empty_trill_soundOrdered[:], func(i, j int) bool {
		return empty_trill_soundOrdered[i].Name < empty_trill_soundOrdered[j].Name
	})
	if len(empty_trill_soundOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, empty_trill_sound := range empty_trill_soundOrdered {

		id = generatesIdentifier("Empty_trill_sound", idx, empty_trill_sound.Name)
		map_Empty_trill_sound_Identifiers[empty_trill_sound] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_trill_sound")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", empty_trill_sound.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(empty_trill_sound.Name))
		initializerStatements += setValueField

	}

	map_Encoding_Identifiers := make(map[*Encoding]string)
	_ = map_Encoding_Identifiers

	encodingOrdered := []*Encoding{}
	for encoding := range stage.Encodings {
		encodingOrdered = append(encodingOrdered, encoding)
	}
	sort.Slice(encodingOrdered[:], func(i, j int) bool {
		return encodingOrdered[i].Name < encodingOrdered[j].Name
	})
	if len(encodingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, encoding := range encodingOrdered {

		id = generatesIdentifier("Encoding", idx, encoding.Name)
		map_Encoding_Identifiers[encoding] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Encoding")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", encoding.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(encoding.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Software")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(encoding.Software))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Encoding_description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(encoding.Encoding_description))
		initializerStatements += setValueField

	}

	map_Ending_Identifiers := make(map[*Ending]string)
	_ = map_Ending_Identifiers

	endingOrdered := []*Ending{}
	for ending := range stage.Endings {
		endingOrdered = append(endingOrdered, ending)
	}
	sort.Slice(endingOrdered[:], func(i, j int) bool {
		return endingOrdered[i].Name < endingOrdered[j].Name
	})
	if len(endingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, ending := range endingOrdered {

		id = generatesIdentifier("Ending", idx, ending.Name)
		map_Ending_Identifiers[ending] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Ending")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ending.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(ending.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(ending.Value))
		initializerStatements += setValueField

	}

	map_Extend_Identifiers := make(map[*Extend]string)
	_ = map_Extend_Identifiers

	extendOrdered := []*Extend{}
	for extend := range stage.Extends {
		extendOrdered = append(extendOrdered, extend)
	}
	sort.Slice(extendOrdered[:], func(i, j int) bool {
		return extendOrdered[i].Name < extendOrdered[j].Name
	})
	if len(extendOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, extend := range extendOrdered {

		id = generatesIdentifier("Extend", idx, extend.Name)
		map_Extend_Identifiers[extend] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Extend")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", extend.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(extend.Name))
		initializerStatements += setValueField

	}

	map_Feature_Identifiers := make(map[*Feature]string)
	_ = map_Feature_Identifiers

	featureOrdered := []*Feature{}
	for feature := range stage.Features {
		featureOrdered = append(featureOrdered, feature)
	}
	sort.Slice(featureOrdered[:], func(i, j int) bool {
		return featureOrdered[i].Name < featureOrdered[j].Name
	})
	if len(featureOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, feature := range featureOrdered {

		id = generatesIdentifier("Feature", idx, feature.Name)
		map_Feature_Identifiers[feature] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Feature")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", feature.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(feature.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(feature.Value))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(feature.Type))
		initializerStatements += setValueField

	}

	map_Fermata_Identifiers := make(map[*Fermata]string)
	_ = map_Fermata_Identifiers

	fermataOrdered := []*Fermata{}
	for fermata := range stage.Fermatas {
		fermataOrdered = append(fermataOrdered, fermata)
	}
	sort.Slice(fermataOrdered[:], func(i, j int) bool {
		return fermataOrdered[i].Name < fermataOrdered[j].Name
	})
	if len(fermataOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, fermata := range fermataOrdered {

		id = generatesIdentifier("Fermata", idx, fermata.Name)
		map_Fermata_Identifiers[fermata] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Fermata")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", fermata.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(fermata.Name))
		initializerStatements += setValueField

	}

	map_Figure_Identifiers := make(map[*Figure]string)
	_ = map_Figure_Identifiers

	figureOrdered := []*Figure{}
	for figure := range stage.Figures {
		figureOrdered = append(figureOrdered, figure)
	}
	sort.Slice(figureOrdered[:], func(i, j int) bool {
		return figureOrdered[i].Name < figureOrdered[j].Name
	})
	if len(figureOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, figure := range figureOrdered {

		id = generatesIdentifier("Figure", idx, figure.Name)
		map_Figure_Identifiers[figure] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Figure")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", figure.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(figure.Name))
		initializerStatements += setValueField

	}

	map_Figured_bass_Identifiers := make(map[*Figured_bass]string)
	_ = map_Figured_bass_Identifiers

	figured_bassOrdered := []*Figured_bass{}
	for figured_bass := range stage.Figured_basss {
		figured_bassOrdered = append(figured_bassOrdered, figured_bass)
	}
	sort.Slice(figured_bassOrdered[:], func(i, j int) bool {
		return figured_bassOrdered[i].Name < figured_bassOrdered[j].Name
	})
	if len(figured_bassOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, figured_bass := range figured_bassOrdered {

		id = generatesIdentifier("Figured_bass", idx, figured_bass.Name)
		map_Figured_bass_Identifiers[figured_bass] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Figured_bass")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", figured_bass.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(figured_bass.Name))
		initializerStatements += setValueField

	}

	map_Fingering_Identifiers := make(map[*Fingering]string)
	_ = map_Fingering_Identifiers

	fingeringOrdered := []*Fingering{}
	for fingering := range stage.Fingerings {
		fingeringOrdered = append(fingeringOrdered, fingering)
	}
	sort.Slice(fingeringOrdered[:], func(i, j int) bool {
		return fingeringOrdered[i].Name < fingeringOrdered[j].Name
	})
	if len(fingeringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, fingering := range fingeringOrdered {

		id = generatesIdentifier("Fingering", idx, fingering.Name)
		map_Fingering_Identifiers[fingering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Fingering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", fingering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(fingering.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(fingering.Value))
		initializerStatements += setValueField

	}

	map_First_fret_Identifiers := make(map[*First_fret]string)
	_ = map_First_fret_Identifiers

	first_fretOrdered := []*First_fret{}
	for first_fret := range stage.First_frets {
		first_fretOrdered = append(first_fretOrdered, first_fret)
	}
	sort.Slice(first_fretOrdered[:], func(i, j int) bool {
		return first_fretOrdered[i].Name < first_fretOrdered[j].Name
	})
	if len(first_fretOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, first_fret := range first_fretOrdered {

		id = generatesIdentifier("First_fret", idx, first_fret.Name)
		map_First_fret_Identifiers[first_fret] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "First_fret")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", first_fret.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(first_fret.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(first_fret.Text))
		initializerStatements += setValueField

	}

	map_Foo_Identifiers := make(map[*Foo]string)
	_ = map_Foo_Identifiers

	fooOrdered := []*Foo{}
	for foo := range stage.Foos {
		fooOrdered = append(fooOrdered, foo)
	}
	sort.Slice(fooOrdered[:], func(i, j int) bool {
		return fooOrdered[i].Name < fooOrdered[j].Name
	})
	if len(fooOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, foo := range fooOrdered {

		id = generatesIdentifier("Foo", idx, foo.Name)
		map_Foo_Identifiers[foo] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Foo")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", foo.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(foo.Name))
		initializerStatements += setValueField

	}

	map_For_part_Identifiers := make(map[*For_part]string)
	_ = map_For_part_Identifiers

	for_partOrdered := []*For_part{}
	for for_part := range stage.For_parts {
		for_partOrdered = append(for_partOrdered, for_part)
	}
	sort.Slice(for_partOrdered[:], func(i, j int) bool {
		return for_partOrdered[i].Name < for_partOrdered[j].Name
	})
	if len(for_partOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, for_part := range for_partOrdered {

		id = generatesIdentifier("For_part", idx, for_part.Name)
		map_For_part_Identifiers[for_part] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "For_part")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", for_part.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(for_part.Name))
		initializerStatements += setValueField

	}

	map_Formatted_symbol_Identifiers := make(map[*Formatted_symbol]string)
	_ = map_Formatted_symbol_Identifiers

	formatted_symbolOrdered := []*Formatted_symbol{}
	for formatted_symbol := range stage.Formatted_symbols {
		formatted_symbolOrdered = append(formatted_symbolOrdered, formatted_symbol)
	}
	sort.Slice(formatted_symbolOrdered[:], func(i, j int) bool {
		return formatted_symbolOrdered[i].Name < formatted_symbolOrdered[j].Name
	})
	if len(formatted_symbolOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, formatted_symbol := range formatted_symbolOrdered {

		id = generatesIdentifier("Formatted_symbol", idx, formatted_symbol.Name)
		map_Formatted_symbol_Identifiers[formatted_symbol] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Formatted_symbol")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formatted_symbol.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formatted_symbol.Name))
		initializerStatements += setValueField

	}

	map_Formatted_symbol_id_Identifiers := make(map[*Formatted_symbol_id]string)
	_ = map_Formatted_symbol_id_Identifiers

	formatted_symbol_idOrdered := []*Formatted_symbol_id{}
	for formatted_symbol_id := range stage.Formatted_symbol_ids {
		formatted_symbol_idOrdered = append(formatted_symbol_idOrdered, formatted_symbol_id)
	}
	sort.Slice(formatted_symbol_idOrdered[:], func(i, j int) bool {
		return formatted_symbol_idOrdered[i].Name < formatted_symbol_idOrdered[j].Name
	})
	if len(formatted_symbol_idOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, formatted_symbol_id := range formatted_symbol_idOrdered {

		id = generatesIdentifier("Formatted_symbol_id", idx, formatted_symbol_id.Name)
		map_Formatted_symbol_id_Identifiers[formatted_symbol_id] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Formatted_symbol_id")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formatted_symbol_id.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formatted_symbol_id.Name))
		initializerStatements += setValueField

	}

	map_Forward_Identifiers := make(map[*Forward]string)
	_ = map_Forward_Identifiers

	forwardOrdered := []*Forward{}
	for forward := range stage.Forwards {
		forwardOrdered = append(forwardOrdered, forward)
	}
	sort.Slice(forwardOrdered[:], func(i, j int) bool {
		return forwardOrdered[i].Name < forwardOrdered[j].Name
	})
	if len(forwardOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, forward := range forwardOrdered {

		id = generatesIdentifier("Forward", idx, forward.Name)
		map_Forward_Identifiers[forward] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Forward")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", forward.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(forward.Name))
		initializerStatements += setValueField

	}

	map_Frame_Identifiers := make(map[*Frame]string)
	_ = map_Frame_Identifiers

	frameOrdered := []*Frame{}
	for frame := range stage.Frames {
		frameOrdered = append(frameOrdered, frame)
	}
	sort.Slice(frameOrdered[:], func(i, j int) bool {
		return frameOrdered[i].Name < frameOrdered[j].Name
	})
	if len(frameOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, frame := range frameOrdered {

		id = generatesIdentifier("Frame", idx, frame.Name)
		map_Frame_Identifiers[frame] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Frame")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", frame.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frame.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Unplayed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frame.Unplayed))
		initializerStatements += setValueField

	}

	map_Frame_note_Identifiers := make(map[*Frame_note]string)
	_ = map_Frame_note_Identifiers

	frame_noteOrdered := []*Frame_note{}
	for frame_note := range stage.Frame_notes {
		frame_noteOrdered = append(frame_noteOrdered, frame_note)
	}
	sort.Slice(frame_noteOrdered[:], func(i, j int) bool {
		return frame_noteOrdered[i].Name < frame_noteOrdered[j].Name
	})
	if len(frame_noteOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, frame_note := range frame_noteOrdered {

		id = generatesIdentifier("Frame_note", idx, frame_note.Name)
		map_Frame_note_Identifiers[frame_note] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Frame_note")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", frame_note.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frame_note.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Astring")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(frame_note.Astring))
		initializerStatements += setValueField

	}

	map_Fret_Identifiers := make(map[*Fret]string)
	_ = map_Fret_Identifiers

	fretOrdered := []*Fret{}
	for fret := range stage.Frets {
		fretOrdered = append(fretOrdered, fret)
	}
	sort.Slice(fretOrdered[:], func(i, j int) bool {
		return fretOrdered[i].Name < fretOrdered[j].Name
	})
	if len(fretOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, fret := range fretOrdered {

		id = generatesIdentifier("Fret", idx, fret.Name)
		map_Fret_Identifiers[fret] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Fret")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", fret.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(fret.Name))
		initializerStatements += setValueField

	}

	map_Glass_Identifiers := make(map[*Glass]string)
	_ = map_Glass_Identifiers

	glassOrdered := []*Glass{}
	for glass := range stage.Glasss {
		glassOrdered = append(glassOrdered, glass)
	}
	sort.Slice(glassOrdered[:], func(i, j int) bool {
		return glassOrdered[i].Name < glassOrdered[j].Name
	})
	if len(glassOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, glass := range glassOrdered {

		id = generatesIdentifier("Glass", idx, glass.Name)
		map_Glass_Identifiers[glass] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Glass")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", glass.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(glass.Name))
		initializerStatements += setValueField

	}

	map_Glissando_Identifiers := make(map[*Glissando]string)
	_ = map_Glissando_Identifiers

	glissandoOrdered := []*Glissando{}
	for glissando := range stage.Glissandos {
		glissandoOrdered = append(glissandoOrdered, glissando)
	}
	sort.Slice(glissandoOrdered[:], func(i, j int) bool {
		return glissandoOrdered[i].Name < glissandoOrdered[j].Name
	})
	if len(glissandoOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, glissando := range glissandoOrdered {

		id = generatesIdentifier("Glissando", idx, glissando.Name)
		map_Glissando_Identifiers[glissando] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Glissando")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", glissando.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(glissando.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(glissando.Value))
		initializerStatements += setValueField

	}

	map_Glyph_Identifiers := make(map[*Glyph]string)
	_ = map_Glyph_Identifiers

	glyphOrdered := []*Glyph{}
	for glyph := range stage.Glyphs {
		glyphOrdered = append(glyphOrdered, glyph)
	}
	sort.Slice(glyphOrdered[:], func(i, j int) bool {
		return glyphOrdered[i].Name < glyphOrdered[j].Name
	})
	if len(glyphOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, glyph := range glyphOrdered {

		id = generatesIdentifier("Glyph", idx, glyph.Name)
		map_Glyph_Identifiers[glyph] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Glyph")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", glyph.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(glyph.Name))
		initializerStatements += setValueField

	}

	map_Grace_Identifiers := make(map[*Grace]string)
	_ = map_Grace_Identifiers

	graceOrdered := []*Grace{}
	for grace := range stage.Graces {
		graceOrdered = append(graceOrdered, grace)
	}
	sort.Slice(graceOrdered[:], func(i, j int) bool {
		return graceOrdered[i].Name < graceOrdered[j].Name
	})
	if len(graceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, grace := range graceOrdered {

		id = generatesIdentifier("Grace", idx, grace.Name)
		map_Grace_Identifiers[grace] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Grace")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", grace.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(grace.Name))
		initializerStatements += setValueField

	}

	map_Group_barline_Identifiers := make(map[*Group_barline]string)
	_ = map_Group_barline_Identifiers

	group_barlineOrdered := []*Group_barline{}
	for group_barline := range stage.Group_barlines {
		group_barlineOrdered = append(group_barlineOrdered, group_barline)
	}
	sort.Slice(group_barlineOrdered[:], func(i, j int) bool {
		return group_barlineOrdered[i].Name < group_barlineOrdered[j].Name
	})
	if len(group_barlineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, group_barline := range group_barlineOrdered {

		id = generatesIdentifier("Group_barline", idx, group_barline.Name)
		map_Group_barline_Identifiers[group_barline] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group_barline")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group_barline.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group_barline.Name))
		initializerStatements += setValueField

	}

	map_Group_symbol_Identifiers := make(map[*Group_symbol]string)
	_ = map_Group_symbol_Identifiers

	group_symbolOrdered := []*Group_symbol{}
	for group_symbol := range stage.Group_symbols {
		group_symbolOrdered = append(group_symbolOrdered, group_symbol)
	}
	sort.Slice(group_symbolOrdered[:], func(i, j int) bool {
		return group_symbolOrdered[i].Name < group_symbolOrdered[j].Name
	})
	if len(group_symbolOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, group_symbol := range group_symbolOrdered {

		id = generatesIdentifier("Group_symbol", idx, group_symbol.Name)
		map_Group_symbol_Identifiers[group_symbol] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group_symbol")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group_symbol.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group_symbol.Name))
		initializerStatements += setValueField

	}

	map_Grouping_Identifiers := make(map[*Grouping]string)
	_ = map_Grouping_Identifiers

	groupingOrdered := []*Grouping{}
	for grouping := range stage.Groupings {
		groupingOrdered = append(groupingOrdered, grouping)
	}
	sort.Slice(groupingOrdered[:], func(i, j int) bool {
		return groupingOrdered[i].Name < groupingOrdered[j].Name
	})
	if len(groupingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, grouping := range groupingOrdered {

		id = generatesIdentifier("Grouping", idx, grouping.Name)
		map_Grouping_Identifiers[grouping] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Grouping")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", grouping.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(grouping.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Number")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(grouping.Number))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Member_of")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(grouping.Member_of))
		initializerStatements += setValueField

	}

	map_Hammer_on_pull_off_Identifiers := make(map[*Hammer_on_pull_off]string)
	_ = map_Hammer_on_pull_off_Identifiers

	hammer_on_pull_offOrdered := []*Hammer_on_pull_off{}
	for hammer_on_pull_off := range stage.Hammer_on_pull_offs {
		hammer_on_pull_offOrdered = append(hammer_on_pull_offOrdered, hammer_on_pull_off)
	}
	sort.Slice(hammer_on_pull_offOrdered[:], func(i, j int) bool {
		return hammer_on_pull_offOrdered[i].Name < hammer_on_pull_offOrdered[j].Name
	})
	if len(hammer_on_pull_offOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, hammer_on_pull_off := range hammer_on_pull_offOrdered {

		id = generatesIdentifier("Hammer_on_pull_off", idx, hammer_on_pull_off.Name)
		map_Hammer_on_pull_off_Identifiers[hammer_on_pull_off] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Hammer_on_pull_off")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", hammer_on_pull_off.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(hammer_on_pull_off.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(hammer_on_pull_off.Value))
		initializerStatements += setValueField

	}

	map_Handbell_Identifiers := make(map[*Handbell]string)
	_ = map_Handbell_Identifiers

	handbellOrdered := []*Handbell{}
	for handbell := range stage.Handbells {
		handbellOrdered = append(handbellOrdered, handbell)
	}
	sort.Slice(handbellOrdered[:], func(i, j int) bool {
		return handbellOrdered[i].Name < handbellOrdered[j].Name
	})
	if len(handbellOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, handbell := range handbellOrdered {

		id = generatesIdentifier("Handbell", idx, handbell.Name)
		map_Handbell_Identifiers[handbell] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Handbell")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", handbell.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(handbell.Name))
		initializerStatements += setValueField

	}

	map_Harmon_closed_Identifiers := make(map[*Harmon_closed]string)
	_ = map_Harmon_closed_Identifiers

	harmon_closedOrdered := []*Harmon_closed{}
	for harmon_closed := range stage.Harmon_closeds {
		harmon_closedOrdered = append(harmon_closedOrdered, harmon_closed)
	}
	sort.Slice(harmon_closedOrdered[:], func(i, j int) bool {
		return harmon_closedOrdered[i].Name < harmon_closedOrdered[j].Name
	})
	if len(harmon_closedOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, harmon_closed := range harmon_closedOrdered {

		id = generatesIdentifier("Harmon_closed", idx, harmon_closed.Name)
		map_Harmon_closed_Identifiers[harmon_closed] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmon_closed")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", harmon_closed.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(harmon_closed.Name))
		initializerStatements += setValueField

	}

	map_Harmon_mute_Identifiers := make(map[*Harmon_mute]string)
	_ = map_Harmon_mute_Identifiers

	harmon_muteOrdered := []*Harmon_mute{}
	for harmon_mute := range stage.Harmon_mutes {
		harmon_muteOrdered = append(harmon_muteOrdered, harmon_mute)
	}
	sort.Slice(harmon_muteOrdered[:], func(i, j int) bool {
		return harmon_muteOrdered[i].Name < harmon_muteOrdered[j].Name
	})
	if len(harmon_muteOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, harmon_mute := range harmon_muteOrdered {

		id = generatesIdentifier("Harmon_mute", idx, harmon_mute.Name)
		map_Harmon_mute_Identifiers[harmon_mute] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmon_mute")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", harmon_mute.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(harmon_mute.Name))
		initializerStatements += setValueField

	}

	map_Harmonic_Identifiers := make(map[*Harmonic]string)
	_ = map_Harmonic_Identifiers

	harmonicOrdered := []*Harmonic{}
	for harmonic := range stage.Harmonics {
		harmonicOrdered = append(harmonicOrdered, harmonic)
	}
	sort.Slice(harmonicOrdered[:], func(i, j int) bool {
		return harmonicOrdered[i].Name < harmonicOrdered[j].Name
	})
	if len(harmonicOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, harmonic := range harmonicOrdered {

		id = generatesIdentifier("Harmonic", idx, harmonic.Name)
		map_Harmonic_Identifiers[harmonic] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmonic")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", harmonic.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(harmonic.Name))
		initializerStatements += setValueField

	}

	map_Harmony_Identifiers := make(map[*Harmony]string)
	_ = map_Harmony_Identifiers

	harmonyOrdered := []*Harmony{}
	for harmony := range stage.Harmonys {
		harmonyOrdered = append(harmonyOrdered, harmony)
	}
	sort.Slice(harmonyOrdered[:], func(i, j int) bool {
		return harmonyOrdered[i].Name < harmonyOrdered[j].Name
	})
	if len(harmonyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, harmony := range harmonyOrdered {

		id = generatesIdentifier("Harmony", idx, harmony.Name)
		map_Harmony_Identifiers[harmony] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmony")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", harmony.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(harmony.Name))
		initializerStatements += setValueField

	}

	map_Harmony_alter_Identifiers := make(map[*Harmony_alter]string)
	_ = map_Harmony_alter_Identifiers

	harmony_alterOrdered := []*Harmony_alter{}
	for harmony_alter := range stage.Harmony_alters {
		harmony_alterOrdered = append(harmony_alterOrdered, harmony_alter)
	}
	sort.Slice(harmony_alterOrdered[:], func(i, j int) bool {
		return harmony_alterOrdered[i].Name < harmony_alterOrdered[j].Name
	})
	if len(harmony_alterOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, harmony_alter := range harmony_alterOrdered {

		id = generatesIdentifier("Harmony_alter", idx, harmony_alter.Name)
		map_Harmony_alter_Identifiers[harmony_alter] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmony_alter")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", harmony_alter.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(harmony_alter.Name))
		initializerStatements += setValueField

	}

	map_Harp_pedals_Identifiers := make(map[*Harp_pedals]string)
	_ = map_Harp_pedals_Identifiers

	harp_pedalsOrdered := []*Harp_pedals{}
	for harp_pedals := range stage.Harp_pedalss {
		harp_pedalsOrdered = append(harp_pedalsOrdered, harp_pedals)
	}
	sort.Slice(harp_pedalsOrdered[:], func(i, j int) bool {
		return harp_pedalsOrdered[i].Name < harp_pedalsOrdered[j].Name
	})
	if len(harp_pedalsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, harp_pedals := range harp_pedalsOrdered {

		id = generatesIdentifier("Harp_pedals", idx, harp_pedals.Name)
		map_Harp_pedals_Identifiers[harp_pedals] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harp_pedals")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", harp_pedals.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(harp_pedals.Name))
		initializerStatements += setValueField

	}

	map_Heel_toe_Identifiers := make(map[*Heel_toe]string)
	_ = map_Heel_toe_Identifiers

	heel_toeOrdered := []*Heel_toe{}
	for heel_toe := range stage.Heel_toes {
		heel_toeOrdered = append(heel_toeOrdered, heel_toe)
	}
	sort.Slice(heel_toeOrdered[:], func(i, j int) bool {
		return heel_toeOrdered[i].Name < heel_toeOrdered[j].Name
	})
	if len(heel_toeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, heel_toe := range heel_toeOrdered {

		id = generatesIdentifier("Heel_toe", idx, heel_toe.Name)
		map_Heel_toe_Identifiers[heel_toe] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Heel_toe")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", heel_toe.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(heel_toe.Name))
		initializerStatements += setValueField

	}

	map_Hole_Identifiers := make(map[*Hole]string)
	_ = map_Hole_Identifiers

	holeOrdered := []*Hole{}
	for hole := range stage.Holes {
		holeOrdered = append(holeOrdered, hole)
	}
	sort.Slice(holeOrdered[:], func(i, j int) bool {
		return holeOrdered[i].Name < holeOrdered[j].Name
	})
	if len(holeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, hole := range holeOrdered {

		id = generatesIdentifier("Hole", idx, hole.Name)
		map_Hole_Identifiers[hole] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Hole")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", hole.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(hole.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Hole_type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(hole.Hole_type))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Hole_shape")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(hole.Hole_shape))
		initializerStatements += setValueField

	}

	map_Hole_closed_Identifiers := make(map[*Hole_closed]string)
	_ = map_Hole_closed_Identifiers

	hole_closedOrdered := []*Hole_closed{}
	for hole_closed := range stage.Hole_closeds {
		hole_closedOrdered = append(hole_closedOrdered, hole_closed)
	}
	sort.Slice(hole_closedOrdered[:], func(i, j int) bool {
		return hole_closedOrdered[i].Name < hole_closedOrdered[j].Name
	})
	if len(hole_closedOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, hole_closed := range hole_closedOrdered {

		id = generatesIdentifier("Hole_closed", idx, hole_closed.Name)
		map_Hole_closed_Identifiers[hole_closed] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Hole_closed")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", hole_closed.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(hole_closed.Name))
		initializerStatements += setValueField

	}

	map_Horizontal_turn_Identifiers := make(map[*Horizontal_turn]string)
	_ = map_Horizontal_turn_Identifiers

	horizontal_turnOrdered := []*Horizontal_turn{}
	for horizontal_turn := range stage.Horizontal_turns {
		horizontal_turnOrdered = append(horizontal_turnOrdered, horizontal_turn)
	}
	sort.Slice(horizontal_turnOrdered[:], func(i, j int) bool {
		return horizontal_turnOrdered[i].Name < horizontal_turnOrdered[j].Name
	})
	if len(horizontal_turnOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, horizontal_turn := range horizontal_turnOrdered {

		id = generatesIdentifier("Horizontal_turn", idx, horizontal_turn.Name)
		map_Horizontal_turn_Identifiers[horizontal_turn] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Horizontal_turn")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", horizontal_turn.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(horizontal_turn.Name))
		initializerStatements += setValueField

	}

	map_Identification_Identifiers := make(map[*Identification]string)
	_ = map_Identification_Identifiers

	identificationOrdered := []*Identification{}
	for identification := range stage.Identifications {
		identificationOrdered = append(identificationOrdered, identification)
	}
	sort.Slice(identificationOrdered[:], func(i, j int) bool {
		return identificationOrdered[i].Name < identificationOrdered[j].Name
	})
	if len(identificationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, identification := range identificationOrdered {

		id = generatesIdentifier("Identification", idx, identification.Name)
		map_Identification_Identifiers[identification] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Identification")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", identification.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(identification.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Source")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(identification.Source))
		initializerStatements += setValueField

	}

	map_Image_Identifiers := make(map[*Image]string)
	_ = map_Image_Identifiers

	imageOrdered := []*Image{}
	for image := range stage.Images {
		imageOrdered = append(imageOrdered, image)
	}
	sort.Slice(imageOrdered[:], func(i, j int) bool {
		return imageOrdered[i].Name < imageOrdered[j].Name
	})
	if len(imageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, image := range imageOrdered {

		id = generatesIdentifier("Image", idx, image.Name)
		map_Image_Identifiers[image] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Image")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", image.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(image.Name))
		initializerStatements += setValueField

	}

	map_Instrument_Identifiers := make(map[*Instrument]string)
	_ = map_Instrument_Identifiers

	instrumentOrdered := []*Instrument{}
	for instrument := range stage.Instruments {
		instrumentOrdered = append(instrumentOrdered, instrument)
	}
	sort.Slice(instrumentOrdered[:], func(i, j int) bool {
		return instrumentOrdered[i].Name < instrumentOrdered[j].Name
	})
	if len(instrumentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, instrument := range instrumentOrdered {

		id = generatesIdentifier("Instrument", idx, instrument.Name)
		map_Instrument_Identifiers[instrument] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Instrument")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", instrument.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(instrument.Name))
		initializerStatements += setValueField

	}

	map_Instrument_change_Identifiers := make(map[*Instrument_change]string)
	_ = map_Instrument_change_Identifiers

	instrument_changeOrdered := []*Instrument_change{}
	for instrument_change := range stage.Instrument_changes {
		instrument_changeOrdered = append(instrument_changeOrdered, instrument_change)
	}
	sort.Slice(instrument_changeOrdered[:], func(i, j int) bool {
		return instrument_changeOrdered[i].Name < instrument_changeOrdered[j].Name
	})
	if len(instrument_changeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, instrument_change := range instrument_changeOrdered {

		id = generatesIdentifier("Instrument_change", idx, instrument_change.Name)
		map_Instrument_change_Identifiers[instrument_change] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Instrument_change")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", instrument_change.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(instrument_change.Name))
		initializerStatements += setValueField

	}

	map_Instrument_link_Identifiers := make(map[*Instrument_link]string)
	_ = map_Instrument_link_Identifiers

	instrument_linkOrdered := []*Instrument_link{}
	for instrument_link := range stage.Instrument_links {
		instrument_linkOrdered = append(instrument_linkOrdered, instrument_link)
	}
	sort.Slice(instrument_linkOrdered[:], func(i, j int) bool {
		return instrument_linkOrdered[i].Name < instrument_linkOrdered[j].Name
	})
	if len(instrument_linkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, instrument_link := range instrument_linkOrdered {

		id = generatesIdentifier("Instrument_link", idx, instrument_link.Name)
		map_Instrument_link_Identifiers[instrument_link] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Instrument_link")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", instrument_link.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(instrument_link.Name))
		initializerStatements += setValueField

	}

	map_Interchangeable_Identifiers := make(map[*Interchangeable]string)
	_ = map_Interchangeable_Identifiers

	interchangeableOrdered := []*Interchangeable{}
	for interchangeable := range stage.Interchangeables {
		interchangeableOrdered = append(interchangeableOrdered, interchangeable)
	}
	sort.Slice(interchangeableOrdered[:], func(i, j int) bool {
		return interchangeableOrdered[i].Name < interchangeableOrdered[j].Name
	})
	if len(interchangeableOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, interchangeable := range interchangeableOrdered {

		id = generatesIdentifier("Interchangeable", idx, interchangeable.Name)
		map_Interchangeable_Identifiers[interchangeable] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Interchangeable")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", interchangeable.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(interchangeable.Name))
		initializerStatements += setValueField

	}

	map_Inversion_Identifiers := make(map[*Inversion]string)
	_ = map_Inversion_Identifiers

	inversionOrdered := []*Inversion{}
	for inversion := range stage.Inversions {
		inversionOrdered = append(inversionOrdered, inversion)
	}
	sort.Slice(inversionOrdered[:], func(i, j int) bool {
		return inversionOrdered[i].Name < inversionOrdered[j].Name
	})
	if len(inversionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, inversion := range inversionOrdered {

		id = generatesIdentifier("Inversion", idx, inversion.Name)
		map_Inversion_Identifiers[inversion] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Inversion")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", inversion.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(inversion.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(inversion.Text))
		initializerStatements += setValueField

	}

	map_Key_Identifiers := make(map[*Key]string)
	_ = map_Key_Identifiers

	keyOrdered := []*Key{}
	for key := range stage.Keys {
		keyOrdered = append(keyOrdered, key)
	}
	sort.Slice(keyOrdered[:], func(i, j int) bool {
		return keyOrdered[i].Name < keyOrdered[j].Name
	})
	if len(keyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, key := range keyOrdered {

		id = generatesIdentifier("Key", idx, key.Name)
		map_Key_Identifiers[key] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", key.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key.Name))
		initializerStatements += setValueField

	}

	map_Key_accidental_Identifiers := make(map[*Key_accidental]string)
	_ = map_Key_accidental_Identifiers

	key_accidentalOrdered := []*Key_accidental{}
	for key_accidental := range stage.Key_accidentals {
		key_accidentalOrdered = append(key_accidentalOrdered, key_accidental)
	}
	sort.Slice(key_accidentalOrdered[:], func(i, j int) bool {
		return key_accidentalOrdered[i].Name < key_accidentalOrdered[j].Name
	})
	if len(key_accidentalOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, key_accidental := range key_accidentalOrdered {

		id = generatesIdentifier("Key_accidental", idx, key_accidental.Name)
		map_Key_accidental_Identifiers[key_accidental] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key_accidental")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", key_accidental.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key_accidental.Name))
		initializerStatements += setValueField

	}

	map_Key_octave_Identifiers := make(map[*Key_octave]string)
	_ = map_Key_octave_Identifiers

	key_octaveOrdered := []*Key_octave{}
	for key_octave := range stage.Key_octaves {
		key_octaveOrdered = append(key_octaveOrdered, key_octave)
	}
	sort.Slice(key_octaveOrdered[:], func(i, j int) bool {
		return key_octaveOrdered[i].Name < key_octaveOrdered[j].Name
	})
	if len(key_octaveOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, key_octave := range key_octaveOrdered {

		id = generatesIdentifier("Key_octave", idx, key_octave.Name)
		map_Key_octave_Identifiers[key_octave] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key_octave")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", key_octave.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(key_octave.Name))
		initializerStatements += setValueField

	}

	map_Kind_Identifiers := make(map[*Kind]string)
	_ = map_Kind_Identifiers

	kindOrdered := []*Kind{}
	for kind := range stage.Kinds {
		kindOrdered = append(kindOrdered, kind)
	}
	sort.Slice(kindOrdered[:], func(i, j int) bool {
		return kindOrdered[i].Name < kindOrdered[j].Name
	})
	if len(kindOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, kind := range kindOrdered {

		id = generatesIdentifier("Kind", idx, kind.Name)
		map_Kind_Identifiers[kind] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Kind")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", kind.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kind.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kind.Text))
		initializerStatements += setValueField

	}

	map_Level_Identifiers := make(map[*Level]string)
	_ = map_Level_Identifiers

	levelOrdered := []*Level{}
	for level := range stage.Levels {
		levelOrdered = append(levelOrdered, level)
	}
	sort.Slice(levelOrdered[:], func(i, j int) bool {
		return levelOrdered[i].Name < levelOrdered[j].Name
	})
	if len(levelOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, level := range levelOrdered {

		id = generatesIdentifier("Level", idx, level.Name)
		map_Level_Identifiers[level] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Level")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", level.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(level.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(level.Value))
		initializerStatements += setValueField

	}

	map_Line_detail_Identifiers := make(map[*Line_detail]string)
	_ = map_Line_detail_Identifiers

	line_detailOrdered := []*Line_detail{}
	for line_detail := range stage.Line_details {
		line_detailOrdered = append(line_detailOrdered, line_detail)
	}
	sort.Slice(line_detailOrdered[:], func(i, j int) bool {
		return line_detailOrdered[i].Name < line_detailOrdered[j].Name
	})
	if len(line_detailOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, line_detail := range line_detailOrdered {

		id = generatesIdentifier("Line_detail", idx, line_detail.Name)
		map_Line_detail_Identifiers[line_detail] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Line_detail")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", line_detail.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(line_detail.Name))
		initializerStatements += setValueField

	}

	map_Line_width_Identifiers := make(map[*Line_width]string)
	_ = map_Line_width_Identifiers

	line_widthOrdered := []*Line_width{}
	for line_width := range stage.Line_widths {
		line_widthOrdered = append(line_widthOrdered, line_width)
	}
	sort.Slice(line_widthOrdered[:], func(i, j int) bool {
		return line_widthOrdered[i].Name < line_widthOrdered[j].Name
	})
	if len(line_widthOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, line_width := range line_widthOrdered {

		id = generatesIdentifier("Line_width", idx, line_width.Name)
		map_Line_width_Identifiers[line_width] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Line_width")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", line_width.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(line_width.Name))
		initializerStatements += setValueField

	}

	map_Link_Identifiers := make(map[*Link]string)
	_ = map_Link_Identifiers

	linkOrdered := []*Link{}
	for link := range stage.Links {
		linkOrdered = append(linkOrdered, link)
	}
	sort.Slice(linkOrdered[:], func(i, j int) bool {
		return linkOrdered[i].Name < linkOrdered[j].Name
	})
	if len(linkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, link := range linkOrdered {

		id = generatesIdentifier("Link", idx, link.Name)
		map_Link_Identifiers[link] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Link")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", link.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.Name))
		initializerStatements += setValueField

	}

	map_Listen_Identifiers := make(map[*Listen]string)
	_ = map_Listen_Identifiers

	listenOrdered := []*Listen{}
	for listen := range stage.Listens {
		listenOrdered = append(listenOrdered, listen)
	}
	sort.Slice(listenOrdered[:], func(i, j int) bool {
		return listenOrdered[i].Name < listenOrdered[j].Name
	})
	if len(listenOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, listen := range listenOrdered {

		id = generatesIdentifier("Listen", idx, listen.Name)
		map_Listen_Identifiers[listen] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Listen")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", listen.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(listen.Name))
		initializerStatements += setValueField

	}

	map_Listening_Identifiers := make(map[*Listening]string)
	_ = map_Listening_Identifiers

	listeningOrdered := []*Listening{}
	for listening := range stage.Listenings {
		listeningOrdered = append(listeningOrdered, listening)
	}
	sort.Slice(listeningOrdered[:], func(i, j int) bool {
		return listeningOrdered[i].Name < listeningOrdered[j].Name
	})
	if len(listeningOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, listening := range listeningOrdered {

		id = generatesIdentifier("Listening", idx, listening.Name)
		map_Listening_Identifiers[listening] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Listening")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", listening.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(listening.Name))
		initializerStatements += setValueField

	}

	map_Lyric_Identifiers := make(map[*Lyric]string)
	_ = map_Lyric_Identifiers

	lyricOrdered := []*Lyric{}
	for lyric := range stage.Lyrics {
		lyricOrdered = append(lyricOrdered, lyric)
	}
	sort.Slice(lyricOrdered[:], func(i, j int) bool {
		return lyricOrdered[i].Name < lyricOrdered[j].Name
	})
	if len(lyricOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, lyric := range lyricOrdered {

		id = generatesIdentifier("Lyric", idx, lyric.Name)
		map_Lyric_Identifiers[lyric] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lyric")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lyric.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lyric.Name))
		initializerStatements += setValueField

	}

	map_Lyric_font_Identifiers := make(map[*Lyric_font]string)
	_ = map_Lyric_font_Identifiers

	lyric_fontOrdered := []*Lyric_font{}
	for lyric_font := range stage.Lyric_fonts {
		lyric_fontOrdered = append(lyric_fontOrdered, lyric_font)
	}
	sort.Slice(lyric_fontOrdered[:], func(i, j int) bool {
		return lyric_fontOrdered[i].Name < lyric_fontOrdered[j].Name
	})
	if len(lyric_fontOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, lyric_font := range lyric_fontOrdered {

		id = generatesIdentifier("Lyric_font", idx, lyric_font.Name)
		map_Lyric_font_Identifiers[lyric_font] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lyric_font")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lyric_font.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lyric_font.Name))
		initializerStatements += setValueField

	}

	map_Lyric_language_Identifiers := make(map[*Lyric_language]string)
	_ = map_Lyric_language_Identifiers

	lyric_languageOrdered := []*Lyric_language{}
	for lyric_language := range stage.Lyric_languages {
		lyric_languageOrdered = append(lyric_languageOrdered, lyric_language)
	}
	sort.Slice(lyric_languageOrdered[:], func(i, j int) bool {
		return lyric_languageOrdered[i].Name < lyric_languageOrdered[j].Name
	})
	if len(lyric_languageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, lyric_language := range lyric_languageOrdered {

		id = generatesIdentifier("Lyric_language", idx, lyric_language.Name)
		map_Lyric_language_Identifiers[lyric_language] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lyric_language")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lyric_language.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lyric_language.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EmptyString")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lyric_language.EmptyString))
		initializerStatements += setValueField

	}

	map_Measure_layout_Identifiers := make(map[*Measure_layout]string)
	_ = map_Measure_layout_Identifiers

	measure_layoutOrdered := []*Measure_layout{}
	for measure_layout := range stage.Measure_layouts {
		measure_layoutOrdered = append(measure_layoutOrdered, measure_layout)
	}
	sort.Slice(measure_layoutOrdered[:], func(i, j int) bool {
		return measure_layoutOrdered[i].Name < measure_layoutOrdered[j].Name
	})
	if len(measure_layoutOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, measure_layout := range measure_layoutOrdered {

		id = generatesIdentifier("Measure_layout", idx, measure_layout.Name)
		map_Measure_layout_Identifiers[measure_layout] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Measure_layout")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", measure_layout.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(measure_layout.Name))
		initializerStatements += setValueField

	}

	map_Measure_numbering_Identifiers := make(map[*Measure_numbering]string)
	_ = map_Measure_numbering_Identifiers

	measure_numberingOrdered := []*Measure_numbering{}
	for measure_numbering := range stage.Measure_numberings {
		measure_numberingOrdered = append(measure_numberingOrdered, measure_numbering)
	}
	sort.Slice(measure_numberingOrdered[:], func(i, j int) bool {
		return measure_numberingOrdered[i].Name < measure_numberingOrdered[j].Name
	})
	if len(measure_numberingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, measure_numbering := range measure_numberingOrdered {

		id = generatesIdentifier("Measure_numbering", idx, measure_numbering.Name)
		map_Measure_numbering_Identifiers[measure_numbering] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Measure_numbering")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", measure_numbering.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(measure_numbering.Name))
		initializerStatements += setValueField

	}

	map_Measure_repeat_Identifiers := make(map[*Measure_repeat]string)
	_ = map_Measure_repeat_Identifiers

	measure_repeatOrdered := []*Measure_repeat{}
	for measure_repeat := range stage.Measure_repeats {
		measure_repeatOrdered = append(measure_repeatOrdered, measure_repeat)
	}
	sort.Slice(measure_repeatOrdered[:], func(i, j int) bool {
		return measure_repeatOrdered[i].Name < measure_repeatOrdered[j].Name
	})
	if len(measure_repeatOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, measure_repeat := range measure_repeatOrdered {

		id = generatesIdentifier("Measure_repeat", idx, measure_repeat.Name)
		map_Measure_repeat_Identifiers[measure_repeat] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Measure_repeat")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", measure_repeat.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(measure_repeat.Name))
		initializerStatements += setValueField

	}

	map_Measure_style_Identifiers := make(map[*Measure_style]string)
	_ = map_Measure_style_Identifiers

	measure_styleOrdered := []*Measure_style{}
	for measure_style := range stage.Measure_styles {
		measure_styleOrdered = append(measure_styleOrdered, measure_style)
	}
	sort.Slice(measure_styleOrdered[:], func(i, j int) bool {
		return measure_styleOrdered[i].Name < measure_styleOrdered[j].Name
	})
	if len(measure_styleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, measure_style := range measure_styleOrdered {

		id = generatesIdentifier("Measure_style", idx, measure_style.Name)
		map_Measure_style_Identifiers[measure_style] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Measure_style")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", measure_style.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(measure_style.Name))
		initializerStatements += setValueField

	}

	map_Membrane_Identifiers := make(map[*Membrane]string)
	_ = map_Membrane_Identifiers

	membraneOrdered := []*Membrane{}
	for membrane := range stage.Membranes {
		membraneOrdered = append(membraneOrdered, membrane)
	}
	sort.Slice(membraneOrdered[:], func(i, j int) bool {
		return membraneOrdered[i].Name < membraneOrdered[j].Name
	})
	if len(membraneOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, membrane := range membraneOrdered {

		id = generatesIdentifier("Membrane", idx, membrane.Name)
		map_Membrane_Identifiers[membrane] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Membrane")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", membrane.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(membrane.Name))
		initializerStatements += setValueField

	}

	map_Metal_Identifiers := make(map[*Metal]string)
	_ = map_Metal_Identifiers

	metalOrdered := []*Metal{}
	for metal := range stage.Metals {
		metalOrdered = append(metalOrdered, metal)
	}
	sort.Slice(metalOrdered[:], func(i, j int) bool {
		return metalOrdered[i].Name < metalOrdered[j].Name
	})
	if len(metalOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, metal := range metalOrdered {

		id = generatesIdentifier("Metal", idx, metal.Name)
		map_Metal_Identifiers[metal] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metal")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", metal.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(metal.Name))
		initializerStatements += setValueField

	}

	map_Metronome_Identifiers := make(map[*Metronome]string)
	_ = map_Metronome_Identifiers

	metronomeOrdered := []*Metronome{}
	for metronome := range stage.Metronomes {
		metronomeOrdered = append(metronomeOrdered, metronome)
	}
	sort.Slice(metronomeOrdered[:], func(i, j int) bool {
		return metronomeOrdered[i].Name < metronomeOrdered[j].Name
	})
	if len(metronomeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, metronome := range metronomeOrdered {

		id = generatesIdentifier("Metronome", idx, metronome.Name)
		map_Metronome_Identifiers[metronome] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", metronome.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(metronome.Name))
		initializerStatements += setValueField

	}

	map_Metronome_beam_Identifiers := make(map[*Metronome_beam]string)
	_ = map_Metronome_beam_Identifiers

	metronome_beamOrdered := []*Metronome_beam{}
	for metronome_beam := range stage.Metronome_beams {
		metronome_beamOrdered = append(metronome_beamOrdered, metronome_beam)
	}
	sort.Slice(metronome_beamOrdered[:], func(i, j int) bool {
		return metronome_beamOrdered[i].Name < metronome_beamOrdered[j].Name
	})
	if len(metronome_beamOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, metronome_beam := range metronome_beamOrdered {

		id = generatesIdentifier("Metronome_beam", idx, metronome_beam.Name)
		map_Metronome_beam_Identifiers[metronome_beam] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome_beam")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", metronome_beam.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(metronome_beam.Name))
		initializerStatements += setValueField

	}

	map_Metronome_note_Identifiers := make(map[*Metronome_note]string)
	_ = map_Metronome_note_Identifiers

	metronome_noteOrdered := []*Metronome_note{}
	for metronome_note := range stage.Metronome_notes {
		metronome_noteOrdered = append(metronome_noteOrdered, metronome_note)
	}
	sort.Slice(metronome_noteOrdered[:], func(i, j int) bool {
		return metronome_noteOrdered[i].Name < metronome_noteOrdered[j].Name
	})
	if len(metronome_noteOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, metronome_note := range metronome_noteOrdered {

		id = generatesIdentifier("Metronome_note", idx, metronome_note.Name)
		map_Metronome_note_Identifiers[metronome_note] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome_note")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", metronome_note.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(metronome_note.Name))
		initializerStatements += setValueField

	}

	map_Metronome_tied_Identifiers := make(map[*Metronome_tied]string)
	_ = map_Metronome_tied_Identifiers

	metronome_tiedOrdered := []*Metronome_tied{}
	for metronome_tied := range stage.Metronome_tieds {
		metronome_tiedOrdered = append(metronome_tiedOrdered, metronome_tied)
	}
	sort.Slice(metronome_tiedOrdered[:], func(i, j int) bool {
		return metronome_tiedOrdered[i].Name < metronome_tiedOrdered[j].Name
	})
	if len(metronome_tiedOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, metronome_tied := range metronome_tiedOrdered {

		id = generatesIdentifier("Metronome_tied", idx, metronome_tied.Name)
		map_Metronome_tied_Identifiers[metronome_tied] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome_tied")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", metronome_tied.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(metronome_tied.Name))
		initializerStatements += setValueField

	}

	map_Metronome_tuplet_Identifiers := make(map[*Metronome_tuplet]string)
	_ = map_Metronome_tuplet_Identifiers

	metronome_tupletOrdered := []*Metronome_tuplet{}
	for metronome_tuplet := range stage.Metronome_tuplets {
		metronome_tupletOrdered = append(metronome_tupletOrdered, metronome_tuplet)
	}
	sort.Slice(metronome_tupletOrdered[:], func(i, j int) bool {
		return metronome_tupletOrdered[i].Name < metronome_tupletOrdered[j].Name
	})
	if len(metronome_tupletOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, metronome_tuplet := range metronome_tupletOrdered {

		id = generatesIdentifier("Metronome_tuplet", idx, metronome_tuplet.Name)
		map_Metronome_tuplet_Identifiers[metronome_tuplet] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome_tuplet")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", metronome_tuplet.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(metronome_tuplet.Name))
		initializerStatements += setValueField

	}

	map_Midi_device_Identifiers := make(map[*Midi_device]string)
	_ = map_Midi_device_Identifiers

	midi_deviceOrdered := []*Midi_device{}
	for midi_device := range stage.Midi_devices {
		midi_deviceOrdered = append(midi_deviceOrdered, midi_device)
	}
	sort.Slice(midi_deviceOrdered[:], func(i, j int) bool {
		return midi_deviceOrdered[i].Name < midi_deviceOrdered[j].Name
	})
	if len(midi_deviceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, midi_device := range midi_deviceOrdered {

		id = generatesIdentifier("Midi_device", idx, midi_device.Name)
		map_Midi_device_Identifiers[midi_device] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Midi_device")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", midi_device.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(midi_device.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(midi_device.Value))
		initializerStatements += setValueField

	}

	map_Midi_instrument_Identifiers := make(map[*Midi_instrument]string)
	_ = map_Midi_instrument_Identifiers

	midi_instrumentOrdered := []*Midi_instrument{}
	for midi_instrument := range stage.Midi_instruments {
		midi_instrumentOrdered = append(midi_instrumentOrdered, midi_instrument)
	}
	sort.Slice(midi_instrumentOrdered[:], func(i, j int) bool {
		return midi_instrumentOrdered[i].Name < midi_instrumentOrdered[j].Name
	})
	if len(midi_instrumentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, midi_instrument := range midi_instrumentOrdered {

		id = generatesIdentifier("Midi_instrument", idx, midi_instrument.Name)
		map_Midi_instrument_Identifiers[midi_instrument] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Midi_instrument")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", midi_instrument.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(midi_instrument.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Midi_name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(midi_instrument.Midi_name))
		initializerStatements += setValueField

	}

	map_Miscellaneous_Identifiers := make(map[*Miscellaneous]string)
	_ = map_Miscellaneous_Identifiers

	miscellaneousOrdered := []*Miscellaneous{}
	for miscellaneous := range stage.Miscellaneouss {
		miscellaneousOrdered = append(miscellaneousOrdered, miscellaneous)
	}
	sort.Slice(miscellaneousOrdered[:], func(i, j int) bool {
		return miscellaneousOrdered[i].Name < miscellaneousOrdered[j].Name
	})
	if len(miscellaneousOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, miscellaneous := range miscellaneousOrdered {

		id = generatesIdentifier("Miscellaneous", idx, miscellaneous.Name)
		map_Miscellaneous_Identifiers[miscellaneous] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Miscellaneous")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", miscellaneous.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(miscellaneous.Name))
		initializerStatements += setValueField

	}

	map_Miscellaneous_field_Identifiers := make(map[*Miscellaneous_field]string)
	_ = map_Miscellaneous_field_Identifiers

	miscellaneous_fieldOrdered := []*Miscellaneous_field{}
	for miscellaneous_field := range stage.Miscellaneous_fields {
		miscellaneous_fieldOrdered = append(miscellaneous_fieldOrdered, miscellaneous_field)
	}
	sort.Slice(miscellaneous_fieldOrdered[:], func(i, j int) bool {
		return miscellaneous_fieldOrdered[i].Name < miscellaneous_fieldOrdered[j].Name
	})
	if len(miscellaneous_fieldOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, miscellaneous_field := range miscellaneous_fieldOrdered {

		id = generatesIdentifier("Miscellaneous_field", idx, miscellaneous_field.Name)
		map_Miscellaneous_field_Identifiers[miscellaneous_field] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Miscellaneous_field")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", miscellaneous_field.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(miscellaneous_field.Value))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(miscellaneous_field.Name))
		initializerStatements += setValueField

	}

	map_Mordent_Identifiers := make(map[*Mordent]string)
	_ = map_Mordent_Identifiers

	mordentOrdered := []*Mordent{}
	for mordent := range stage.Mordents {
		mordentOrdered = append(mordentOrdered, mordent)
	}
	sort.Slice(mordentOrdered[:], func(i, j int) bool {
		return mordentOrdered[i].Name < mordentOrdered[j].Name
	})
	if len(mordentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, mordent := range mordentOrdered {

		id = generatesIdentifier("Mordent", idx, mordent.Name)
		map_Mordent_Identifiers[mordent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Mordent")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", mordent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mordent.Name))
		initializerStatements += setValueField

	}

	map_Multiple_rest_Identifiers := make(map[*Multiple_rest]string)
	_ = map_Multiple_rest_Identifiers

	multiple_restOrdered := []*Multiple_rest{}
	for multiple_rest := range stage.Multiple_rests {
		multiple_restOrdered = append(multiple_restOrdered, multiple_rest)
	}
	sort.Slice(multiple_restOrdered[:], func(i, j int) bool {
		return multiple_restOrdered[i].Name < multiple_restOrdered[j].Name
	})
	if len(multiple_restOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, multiple_rest := range multiple_restOrdered {

		id = generatesIdentifier("Multiple_rest", idx, multiple_rest.Name)
		map_Multiple_rest_Identifiers[multiple_rest] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Multiple_rest")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", multiple_rest.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(multiple_rest.Name))
		initializerStatements += setValueField

	}

	map_Name_display_Identifiers := make(map[*Name_display]string)
	_ = map_Name_display_Identifiers

	name_displayOrdered := []*Name_display{}
	for name_display := range stage.Name_displays {
		name_displayOrdered = append(name_displayOrdered, name_display)
	}
	sort.Slice(name_displayOrdered[:], func(i, j int) bool {
		return name_displayOrdered[i].Name < name_displayOrdered[j].Name
	})
	if len(name_displayOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, name_display := range name_displayOrdered {

		id = generatesIdentifier("Name_display", idx, name_display.Name)
		map_Name_display_Identifiers[name_display] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Name_display")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", name_display.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(name_display.Name))
		initializerStatements += setValueField

	}

	map_Non_arpeggiate_Identifiers := make(map[*Non_arpeggiate]string)
	_ = map_Non_arpeggiate_Identifiers

	non_arpeggiateOrdered := []*Non_arpeggiate{}
	for non_arpeggiate := range stage.Non_arpeggiates {
		non_arpeggiateOrdered = append(non_arpeggiateOrdered, non_arpeggiate)
	}
	sort.Slice(non_arpeggiateOrdered[:], func(i, j int) bool {
		return non_arpeggiateOrdered[i].Name < non_arpeggiateOrdered[j].Name
	})
	if len(non_arpeggiateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, non_arpeggiate := range non_arpeggiateOrdered {

		id = generatesIdentifier("Non_arpeggiate", idx, non_arpeggiate.Name)
		map_Non_arpeggiate_Identifiers[non_arpeggiate] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Non_arpeggiate")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", non_arpeggiate.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(non_arpeggiate.Name))
		initializerStatements += setValueField

	}

	map_Notations_Identifiers := make(map[*Notations]string)
	_ = map_Notations_Identifiers

	notationsOrdered := []*Notations{}
	for notations := range stage.Notationss {
		notationsOrdered = append(notationsOrdered, notations)
	}
	sort.Slice(notationsOrdered[:], func(i, j int) bool {
		return notationsOrdered[i].Name < notationsOrdered[j].Name
	})
	if len(notationsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, notations := range notationsOrdered {

		id = generatesIdentifier("Notations", idx, notations.Name)
		map_Notations_Identifiers[notations] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Notations")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", notations.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(notations.Name))
		initializerStatements += setValueField

	}

	map_Note_Identifiers := make(map[*Note]string)
	_ = map_Note_Identifiers

	noteOrdered := []*Note{}
	for note := range stage.Notes {
		noteOrdered = append(noteOrdered, note)
	}
	sort.Slice(noteOrdered[:], func(i, j int) bool {
		return noteOrdered[i].Name < noteOrdered[j].Name
	})
	if len(noteOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, note := range noteOrdered {

		id = generatesIdentifier("Note", idx, note.Name)
		map_Note_Identifiers[note] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", note.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(note.Name))
		initializerStatements += setValueField

	}

	map_Note_size_Identifiers := make(map[*Note_size]string)
	_ = map_Note_size_Identifiers

	note_sizeOrdered := []*Note_size{}
	for note_size := range stage.Note_sizes {
		note_sizeOrdered = append(note_sizeOrdered, note_size)
	}
	sort.Slice(note_sizeOrdered[:], func(i, j int) bool {
		return note_sizeOrdered[i].Name < note_sizeOrdered[j].Name
	})
	if len(note_sizeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, note_size := range note_sizeOrdered {

		id = generatesIdentifier("Note_size", idx, note_size.Name)
		map_Note_size_Identifiers[note_size] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note_size")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", note_size.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(note_size.Name))
		initializerStatements += setValueField

	}

	map_Note_type_Identifiers := make(map[*Note_type]string)
	_ = map_Note_type_Identifiers

	note_typeOrdered := []*Note_type{}
	for note_type := range stage.Note_types {
		note_typeOrdered = append(note_typeOrdered, note_type)
	}
	sort.Slice(note_typeOrdered[:], func(i, j int) bool {
		return note_typeOrdered[i].Name < note_typeOrdered[j].Name
	})
	if len(note_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, note_type := range note_typeOrdered {

		id = generatesIdentifier("Note_type", idx, note_type.Name)
		map_Note_type_Identifiers[note_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", note_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(note_type.Name))
		initializerStatements += setValueField

	}

	map_Notehead_Identifiers := make(map[*Notehead]string)
	_ = map_Notehead_Identifiers

	noteheadOrdered := []*Notehead{}
	for notehead := range stage.Noteheads {
		noteheadOrdered = append(noteheadOrdered, notehead)
	}
	sort.Slice(noteheadOrdered[:], func(i, j int) bool {
		return noteheadOrdered[i].Name < noteheadOrdered[j].Name
	})
	if len(noteheadOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, notehead := range noteheadOrdered {

		id = generatesIdentifier("Notehead", idx, notehead.Name)
		map_Notehead_Identifiers[notehead] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Notehead")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", notehead.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(notehead.Name))
		initializerStatements += setValueField

	}

	map_Notehead_text_Identifiers := make(map[*Notehead_text]string)
	_ = map_Notehead_text_Identifiers

	notehead_textOrdered := []*Notehead_text{}
	for notehead_text := range stage.Notehead_texts {
		notehead_textOrdered = append(notehead_textOrdered, notehead_text)
	}
	sort.Slice(notehead_textOrdered[:], func(i, j int) bool {
		return notehead_textOrdered[i].Name < notehead_textOrdered[j].Name
	})
	if len(notehead_textOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, notehead_text := range notehead_textOrdered {

		id = generatesIdentifier("Notehead_text", idx, notehead_text.Name)
		map_Notehead_text_Identifiers[notehead_text] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Notehead_text")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", notehead_text.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(notehead_text.Name))
		initializerStatements += setValueField

	}

	map_Numeral_Identifiers := make(map[*Numeral]string)
	_ = map_Numeral_Identifiers

	numeralOrdered := []*Numeral{}
	for numeral := range stage.Numerals {
		numeralOrdered = append(numeralOrdered, numeral)
	}
	sort.Slice(numeralOrdered[:], func(i, j int) bool {
		return numeralOrdered[i].Name < numeralOrdered[j].Name
	})
	if len(numeralOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, numeral := range numeralOrdered {

		id = generatesIdentifier("Numeral", idx, numeral.Name)
		map_Numeral_Identifiers[numeral] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Numeral")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", numeral.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(numeral.Name))
		initializerStatements += setValueField

	}

	map_Numeral_key_Identifiers := make(map[*Numeral_key]string)
	_ = map_Numeral_key_Identifiers

	numeral_keyOrdered := []*Numeral_key{}
	for numeral_key := range stage.Numeral_keys {
		numeral_keyOrdered = append(numeral_keyOrdered, numeral_key)
	}
	sort.Slice(numeral_keyOrdered[:], func(i, j int) bool {
		return numeral_keyOrdered[i].Name < numeral_keyOrdered[j].Name
	})
	if len(numeral_keyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, numeral_key := range numeral_keyOrdered {

		id = generatesIdentifier("Numeral_key", idx, numeral_key.Name)
		map_Numeral_key_Identifiers[numeral_key] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Numeral_key")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", numeral_key.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(numeral_key.Name))
		initializerStatements += setValueField

	}

	map_Numeral_root_Identifiers := make(map[*Numeral_root]string)
	_ = map_Numeral_root_Identifiers

	numeral_rootOrdered := []*Numeral_root{}
	for numeral_root := range stage.Numeral_roots {
		numeral_rootOrdered = append(numeral_rootOrdered, numeral_root)
	}
	sort.Slice(numeral_rootOrdered[:], func(i, j int) bool {
		return numeral_rootOrdered[i].Name < numeral_rootOrdered[j].Name
	})
	if len(numeral_rootOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, numeral_root := range numeral_rootOrdered {

		id = generatesIdentifier("Numeral_root", idx, numeral_root.Name)
		map_Numeral_root_Identifiers[numeral_root] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Numeral_root")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", numeral_root.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(numeral_root.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(numeral_root.Text))
		initializerStatements += setValueField

	}

	map_Octave_shift_Identifiers := make(map[*Octave_shift]string)
	_ = map_Octave_shift_Identifiers

	octave_shiftOrdered := []*Octave_shift{}
	for octave_shift := range stage.Octave_shifts {
		octave_shiftOrdered = append(octave_shiftOrdered, octave_shift)
	}
	sort.Slice(octave_shiftOrdered[:], func(i, j int) bool {
		return octave_shiftOrdered[i].Name < octave_shiftOrdered[j].Name
	})
	if len(octave_shiftOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, octave_shift := range octave_shiftOrdered {

		id = generatesIdentifier("Octave_shift", idx, octave_shift.Name)
		map_Octave_shift_Identifiers[octave_shift] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Octave_shift")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", octave_shift.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(octave_shift.Name))
		initializerStatements += setValueField

	}

	map_Offset_Identifiers := make(map[*Offset]string)
	_ = map_Offset_Identifiers

	offsetOrdered := []*Offset{}
	for offset := range stage.Offsets {
		offsetOrdered = append(offsetOrdered, offset)
	}
	sort.Slice(offsetOrdered[:], func(i, j int) bool {
		return offsetOrdered[i].Name < offsetOrdered[j].Name
	})
	if len(offsetOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, offset := range offsetOrdered {

		id = generatesIdentifier("Offset", idx, offset.Name)
		map_Offset_Identifiers[offset] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Offset")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", offset.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(offset.Name))
		initializerStatements += setValueField

	}

	map_Opus_Identifiers := make(map[*Opus]string)
	_ = map_Opus_Identifiers

	opusOrdered := []*Opus{}
	for opus := range stage.Opuss {
		opusOrdered = append(opusOrdered, opus)
	}
	sort.Slice(opusOrdered[:], func(i, j int) bool {
		return opusOrdered[i].Name < opusOrdered[j].Name
	})
	if len(opusOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, opus := range opusOrdered {

		id = generatesIdentifier("Opus", idx, opus.Name)
		map_Opus_Identifiers[opus] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Opus")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", opus.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(opus.Name))
		initializerStatements += setValueField

	}

	map_Ornaments_Identifiers := make(map[*Ornaments]string)
	_ = map_Ornaments_Identifiers

	ornamentsOrdered := []*Ornaments{}
	for ornaments := range stage.Ornamentss {
		ornamentsOrdered = append(ornamentsOrdered, ornaments)
	}
	sort.Slice(ornamentsOrdered[:], func(i, j int) bool {
		return ornamentsOrdered[i].Name < ornamentsOrdered[j].Name
	})
	if len(ornamentsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, ornaments := range ornamentsOrdered {

		id = generatesIdentifier("Ornaments", idx, ornaments.Name)
		map_Ornaments_Identifiers[ornaments] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Ornaments")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ornaments.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(ornaments.Name))
		initializerStatements += setValueField

	}

	map_Other_appearance_Identifiers := make(map[*Other_appearance]string)
	_ = map_Other_appearance_Identifiers

	other_appearanceOrdered := []*Other_appearance{}
	for other_appearance := range stage.Other_appearances {
		other_appearanceOrdered = append(other_appearanceOrdered, other_appearance)
	}
	sort.Slice(other_appearanceOrdered[:], func(i, j int) bool {
		return other_appearanceOrdered[i].Name < other_appearanceOrdered[j].Name
	})
	if len(other_appearanceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, other_appearance := range other_appearanceOrdered {

		id = generatesIdentifier("Other_appearance", idx, other_appearance.Name)
		map_Other_appearance_Identifiers[other_appearance] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_appearance")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", other_appearance.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_appearance.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_appearance.Value))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_appearance.Type))
		initializerStatements += setValueField

	}

	map_Other_listening_Identifiers := make(map[*Other_listening]string)
	_ = map_Other_listening_Identifiers

	other_listeningOrdered := []*Other_listening{}
	for other_listening := range stage.Other_listenings {
		other_listeningOrdered = append(other_listeningOrdered, other_listening)
	}
	sort.Slice(other_listeningOrdered[:], func(i, j int) bool {
		return other_listeningOrdered[i].Name < other_listeningOrdered[j].Name
	})
	if len(other_listeningOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, other_listening := range other_listeningOrdered {

		id = generatesIdentifier("Other_listening", idx, other_listening.Name)
		map_Other_listening_Identifiers[other_listening] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_listening")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", other_listening.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_listening.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_listening.Value))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_listening.Type))
		initializerStatements += setValueField

	}

	map_Other_notation_Identifiers := make(map[*Other_notation]string)
	_ = map_Other_notation_Identifiers

	other_notationOrdered := []*Other_notation{}
	for other_notation := range stage.Other_notations {
		other_notationOrdered = append(other_notationOrdered, other_notation)
	}
	sort.Slice(other_notationOrdered[:], func(i, j int) bool {
		return other_notationOrdered[i].Name < other_notationOrdered[j].Name
	})
	if len(other_notationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, other_notation := range other_notationOrdered {

		id = generatesIdentifier("Other_notation", idx, other_notation.Name)
		map_Other_notation_Identifiers[other_notation] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_notation")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", other_notation.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_notation.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_notation.Value))
		initializerStatements += setValueField

	}

	map_Other_play_Identifiers := make(map[*Other_play]string)
	_ = map_Other_play_Identifiers

	other_playOrdered := []*Other_play{}
	for other_play := range stage.Other_plays {
		other_playOrdered = append(other_playOrdered, other_play)
	}
	sort.Slice(other_playOrdered[:], func(i, j int) bool {
		return other_playOrdered[i].Name < other_playOrdered[j].Name
	})
	if len(other_playOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, other_play := range other_playOrdered {

		id = generatesIdentifier("Other_play", idx, other_play.Name)
		map_Other_play_Identifiers[other_play] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_play")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", other_play.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_play.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_play.Value))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(other_play.Type))
		initializerStatements += setValueField

	}

	map_Page_layout_Identifiers := make(map[*Page_layout]string)
	_ = map_Page_layout_Identifiers

	page_layoutOrdered := []*Page_layout{}
	for page_layout := range stage.Page_layouts {
		page_layoutOrdered = append(page_layoutOrdered, page_layout)
	}
	sort.Slice(page_layoutOrdered[:], func(i, j int) bool {
		return page_layoutOrdered[i].Name < page_layoutOrdered[j].Name
	})
	if len(page_layoutOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, page_layout := range page_layoutOrdered {

		id = generatesIdentifier("Page_layout", idx, page_layout.Name)
		map_Page_layout_Identifiers[page_layout] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Page_layout")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", page_layout.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(page_layout.Name))
		initializerStatements += setValueField

	}

	map_Page_margins_Identifiers := make(map[*Page_margins]string)
	_ = map_Page_margins_Identifiers

	page_marginsOrdered := []*Page_margins{}
	for page_margins := range stage.Page_marginss {
		page_marginsOrdered = append(page_marginsOrdered, page_margins)
	}
	sort.Slice(page_marginsOrdered[:], func(i, j int) bool {
		return page_marginsOrdered[i].Name < page_marginsOrdered[j].Name
	})
	if len(page_marginsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, page_margins := range page_marginsOrdered {

		id = generatesIdentifier("Page_margins", idx, page_margins.Name)
		map_Page_margins_Identifiers[page_margins] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Page_margins")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", page_margins.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(page_margins.Name))
		initializerStatements += setValueField

	}

	map_Part_clef_Identifiers := make(map[*Part_clef]string)
	_ = map_Part_clef_Identifiers

	part_clefOrdered := []*Part_clef{}
	for part_clef := range stage.Part_clefs {
		part_clefOrdered = append(part_clefOrdered, part_clef)
	}
	sort.Slice(part_clefOrdered[:], func(i, j int) bool {
		return part_clefOrdered[i].Name < part_clefOrdered[j].Name
	})
	if len(part_clefOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, part_clef := range part_clefOrdered {

		id = generatesIdentifier("Part_clef", idx, part_clef.Name)
		map_Part_clef_Identifiers[part_clef] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_clef")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", part_clef.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(part_clef.Name))
		initializerStatements += setValueField

	}

	map_Part_group_Identifiers := make(map[*Part_group]string)
	_ = map_Part_group_Identifiers

	part_groupOrdered := []*Part_group{}
	for part_group := range stage.Part_groups {
		part_groupOrdered = append(part_groupOrdered, part_group)
	}
	sort.Slice(part_groupOrdered[:], func(i, j int) bool {
		return part_groupOrdered[i].Name < part_groupOrdered[j].Name
	})
	if len(part_groupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, part_group := range part_groupOrdered {

		id = generatesIdentifier("Part_group", idx, part_group.Name)
		map_Part_group_Identifiers[part_group] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_group")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", part_group.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(part_group.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Number")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(part_group.Number))
		initializerStatements += setValueField

	}

	map_Part_link_Identifiers := make(map[*Part_link]string)
	_ = map_Part_link_Identifiers

	part_linkOrdered := []*Part_link{}
	for part_link := range stage.Part_links {
		part_linkOrdered = append(part_linkOrdered, part_link)
	}
	sort.Slice(part_linkOrdered[:], func(i, j int) bool {
		return part_linkOrdered[i].Name < part_linkOrdered[j].Name
	})
	if len(part_linkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, part_link := range part_linkOrdered {

		id = generatesIdentifier("Part_link", idx, part_link.Name)
		map_Part_link_Identifiers[part_link] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_link")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", part_link.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(part_link.Name))
		initializerStatements += setValueField

	}

	map_Part_list_Identifiers := make(map[*Part_list]string)
	_ = map_Part_list_Identifiers

	part_listOrdered := []*Part_list{}
	for part_list := range stage.Part_lists {
		part_listOrdered = append(part_listOrdered, part_list)
	}
	sort.Slice(part_listOrdered[:], func(i, j int) bool {
		return part_listOrdered[i].Name < part_listOrdered[j].Name
	})
	if len(part_listOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, part_list := range part_listOrdered {

		id = generatesIdentifier("Part_list", idx, part_list.Name)
		map_Part_list_Identifiers[part_list] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_list")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", part_list.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(part_list.Name))
		initializerStatements += setValueField

	}

	map_Part_symbol_Identifiers := make(map[*Part_symbol]string)
	_ = map_Part_symbol_Identifiers

	part_symbolOrdered := []*Part_symbol{}
	for part_symbol := range stage.Part_symbols {
		part_symbolOrdered = append(part_symbolOrdered, part_symbol)
	}
	sort.Slice(part_symbolOrdered[:], func(i, j int) bool {
		return part_symbolOrdered[i].Name < part_symbolOrdered[j].Name
	})
	if len(part_symbolOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, part_symbol := range part_symbolOrdered {

		id = generatesIdentifier("Part_symbol", idx, part_symbol.Name)
		map_Part_symbol_Identifiers[part_symbol] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_symbol")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", part_symbol.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(part_symbol.Name))
		initializerStatements += setValueField

	}

	map_Part_transpose_Identifiers := make(map[*Part_transpose]string)
	_ = map_Part_transpose_Identifiers

	part_transposeOrdered := []*Part_transpose{}
	for part_transpose := range stage.Part_transposes {
		part_transposeOrdered = append(part_transposeOrdered, part_transpose)
	}
	sort.Slice(part_transposeOrdered[:], func(i, j int) bool {
		return part_transposeOrdered[i].Name < part_transposeOrdered[j].Name
	})
	if len(part_transposeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, part_transpose := range part_transposeOrdered {

		id = generatesIdentifier("Part_transpose", idx, part_transpose.Name)
		map_Part_transpose_Identifiers[part_transpose] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_transpose")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", part_transpose.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(part_transpose.Name))
		initializerStatements += setValueField

	}

	map_Pedal_Identifiers := make(map[*Pedal]string)
	_ = map_Pedal_Identifiers

	pedalOrdered := []*Pedal{}
	for pedal := range stage.Pedals {
		pedalOrdered = append(pedalOrdered, pedal)
	}
	sort.Slice(pedalOrdered[:], func(i, j int) bool {
		return pedalOrdered[i].Name < pedalOrdered[j].Name
	})
	if len(pedalOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, pedal := range pedalOrdered {

		id = generatesIdentifier("Pedal", idx, pedal.Name)
		map_Pedal_Identifiers[pedal] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pedal")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pedal.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pedal.Name))
		initializerStatements += setValueField

	}

	map_Pedal_tuning_Identifiers := make(map[*Pedal_tuning]string)
	_ = map_Pedal_tuning_Identifiers

	pedal_tuningOrdered := []*Pedal_tuning{}
	for pedal_tuning := range stage.Pedal_tunings {
		pedal_tuningOrdered = append(pedal_tuningOrdered, pedal_tuning)
	}
	sort.Slice(pedal_tuningOrdered[:], func(i, j int) bool {
		return pedal_tuningOrdered[i].Name < pedal_tuningOrdered[j].Name
	})
	if len(pedal_tuningOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, pedal_tuning := range pedal_tuningOrdered {

		id = generatesIdentifier("Pedal_tuning", idx, pedal_tuning.Name)
		map_Pedal_tuning_Identifiers[pedal_tuning] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pedal_tuning")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pedal_tuning.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pedal_tuning.Name))
		initializerStatements += setValueField

	}

	map_Percussion_Identifiers := make(map[*Percussion]string)
	_ = map_Percussion_Identifiers

	percussionOrdered := []*Percussion{}
	for percussion := range stage.Percussions {
		percussionOrdered = append(percussionOrdered, percussion)
	}
	sort.Slice(percussionOrdered[:], func(i, j int) bool {
		return percussionOrdered[i].Name < percussionOrdered[j].Name
	})
	if len(percussionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, percussion := range percussionOrdered {

		id = generatesIdentifier("Percussion", idx, percussion.Name)
		map_Percussion_Identifiers[percussion] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Percussion")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", percussion.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(percussion.Name))
		initializerStatements += setValueField

	}

	map_Pitch_Identifiers := make(map[*Pitch]string)
	_ = map_Pitch_Identifiers

	pitchOrdered := []*Pitch{}
	for pitch := range stage.Pitchs {
		pitchOrdered = append(pitchOrdered, pitch)
	}
	sort.Slice(pitchOrdered[:], func(i, j int) bool {
		return pitchOrdered[i].Name < pitchOrdered[j].Name
	})
	if len(pitchOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, pitch := range pitchOrdered {

		id = generatesIdentifier("Pitch", idx, pitch.Name)
		map_Pitch_Identifiers[pitch] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pitch")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pitch.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pitch.Name))
		initializerStatements += setValueField

	}

	map_Pitched_Identifiers := make(map[*Pitched]string)
	_ = map_Pitched_Identifiers

	pitchedOrdered := []*Pitched{}
	for pitched := range stage.Pitcheds {
		pitchedOrdered = append(pitchedOrdered, pitched)
	}
	sort.Slice(pitchedOrdered[:], func(i, j int) bool {
		return pitchedOrdered[i].Name < pitchedOrdered[j].Name
	})
	if len(pitchedOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, pitched := range pitchedOrdered {

		id = generatesIdentifier("Pitched", idx, pitched.Name)
		map_Pitched_Identifiers[pitched] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pitched")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pitched.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pitched.Name))
		initializerStatements += setValueField

	}

	map_Play_Identifiers := make(map[*Play]string)
	_ = map_Play_Identifiers

	playOrdered := []*Play{}
	for play := range stage.Plays {
		playOrdered = append(playOrdered, play)
	}
	sort.Slice(playOrdered[:], func(i, j int) bool {
		return playOrdered[i].Name < playOrdered[j].Name
	})
	if len(playOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, play := range playOrdered {

		id = generatesIdentifier("Play", idx, play.Name)
		map_Play_Identifiers[play] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Play")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", play.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(play.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Ipa")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(play.Ipa))
		initializerStatements += setValueField

	}

	map_Player_Identifiers := make(map[*Player]string)
	_ = map_Player_Identifiers

	playerOrdered := []*Player{}
	for player := range stage.Players {
		playerOrdered = append(playerOrdered, player)
	}
	sort.Slice(playerOrdered[:], func(i, j int) bool {
		return playerOrdered[i].Name < playerOrdered[j].Name
	})
	if len(playerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, player := range playerOrdered {

		id = generatesIdentifier("Player", idx, player.Name)
		map_Player_Identifiers[player] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Player")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", player.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(player.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Player_name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(player.Player_name))
		initializerStatements += setValueField

	}

	map_Principal_voice_Identifiers := make(map[*Principal_voice]string)
	_ = map_Principal_voice_Identifiers

	principal_voiceOrdered := []*Principal_voice{}
	for principal_voice := range stage.Principal_voices {
		principal_voiceOrdered = append(principal_voiceOrdered, principal_voice)
	}
	sort.Slice(principal_voiceOrdered[:], func(i, j int) bool {
		return principal_voiceOrdered[i].Name < principal_voiceOrdered[j].Name
	})
	if len(principal_voiceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, principal_voice := range principal_voiceOrdered {

		id = generatesIdentifier("Principal_voice", idx, principal_voice.Name)
		map_Principal_voice_Identifiers[principal_voice] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Principal_voice")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", principal_voice.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(principal_voice.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(principal_voice.Value))
		initializerStatements += setValueField

	}

	map_Print_Identifiers := make(map[*Print]string)
	_ = map_Print_Identifiers

	printOrdered := []*Print{}
	for print := range stage.Prints {
		printOrdered = append(printOrdered, print)
	}
	sort.Slice(printOrdered[:], func(i, j int) bool {
		return printOrdered[i].Name < printOrdered[j].Name
	})
	if len(printOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, print := range printOrdered {

		id = generatesIdentifier("Print", idx, print.Name)
		map_Print_Identifiers[print] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Print")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", print.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(print.Name))
		initializerStatements += setValueField

	}

	map_Release_Identifiers := make(map[*Release]string)
	_ = map_Release_Identifiers

	releaseOrdered := []*Release{}
	for release := range stage.Releases {
		releaseOrdered = append(releaseOrdered, release)
	}
	sort.Slice(releaseOrdered[:], func(i, j int) bool {
		return releaseOrdered[i].Name < releaseOrdered[j].Name
	})
	if len(releaseOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, release := range releaseOrdered {

		id = generatesIdentifier("Release", idx, release.Name)
		map_Release_Identifiers[release] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Release")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", release.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(release.Name))
		initializerStatements += setValueField

	}

	map_Repeat_Identifiers := make(map[*Repeat]string)
	_ = map_Repeat_Identifiers

	repeatOrdered := []*Repeat{}
	for repeat := range stage.Repeats {
		repeatOrdered = append(repeatOrdered, repeat)
	}
	sort.Slice(repeatOrdered[:], func(i, j int) bool {
		return repeatOrdered[i].Name < repeatOrdered[j].Name
	})
	if len(repeatOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, repeat := range repeatOrdered {

		id = generatesIdentifier("Repeat", idx, repeat.Name)
		map_Repeat_Identifiers[repeat] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Repeat")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", repeat.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(repeat.Name))
		initializerStatements += setValueField

	}

	map_Rest_Identifiers := make(map[*Rest]string)
	_ = map_Rest_Identifiers

	restOrdered := []*Rest{}
	for rest := range stage.Rests {
		restOrdered = append(restOrdered, rest)
	}
	sort.Slice(restOrdered[:], func(i, j int) bool {
		return restOrdered[i].Name < restOrdered[j].Name
	})
	if len(restOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, rest := range restOrdered {

		id = generatesIdentifier("Rest", idx, rest.Name)
		map_Rest_Identifiers[rest] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rest")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rest.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rest.Name))
		initializerStatements += setValueField

	}

	map_Root_Identifiers := make(map[*Root]string)
	_ = map_Root_Identifiers

	rootOrdered := []*Root{}
	for root := range stage.Roots {
		rootOrdered = append(rootOrdered, root)
	}
	sort.Slice(rootOrdered[:], func(i, j int) bool {
		return rootOrdered[i].Name < rootOrdered[j].Name
	})
	if len(rootOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, root := range rootOrdered {

		id = generatesIdentifier("Root", idx, root.Name)
		map_Root_Identifiers[root] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Root")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", root.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(root.Name))
		initializerStatements += setValueField

	}

	map_Root_step_Identifiers := make(map[*Root_step]string)
	_ = map_Root_step_Identifiers

	root_stepOrdered := []*Root_step{}
	for root_step := range stage.Root_steps {
		root_stepOrdered = append(root_stepOrdered, root_step)
	}
	sort.Slice(root_stepOrdered[:], func(i, j int) bool {
		return root_stepOrdered[i].Name < root_stepOrdered[j].Name
	})
	if len(root_stepOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, root_step := range root_stepOrdered {

		id = generatesIdentifier("Root_step", idx, root_step.Name)
		map_Root_step_Identifiers[root_step] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Root_step")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", root_step.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(root_step.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(root_step.Text))
		initializerStatements += setValueField

	}

	map_Scaling_Identifiers := make(map[*Scaling]string)
	_ = map_Scaling_Identifiers

	scalingOrdered := []*Scaling{}
	for scaling := range stage.Scalings {
		scalingOrdered = append(scalingOrdered, scaling)
	}
	sort.Slice(scalingOrdered[:], func(i, j int) bool {
		return scalingOrdered[i].Name < scalingOrdered[j].Name
	})
	if len(scalingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, scaling := range scalingOrdered {

		id = generatesIdentifier("Scaling", idx, scaling.Name)
		map_Scaling_Identifiers[scaling] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Scaling")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", scaling.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(scaling.Name))
		initializerStatements += setValueField

	}

	map_Scordatura_Identifiers := make(map[*Scordatura]string)
	_ = map_Scordatura_Identifiers

	scordaturaOrdered := []*Scordatura{}
	for scordatura := range stage.Scordaturas {
		scordaturaOrdered = append(scordaturaOrdered, scordatura)
	}
	sort.Slice(scordaturaOrdered[:], func(i, j int) bool {
		return scordaturaOrdered[i].Name < scordaturaOrdered[j].Name
	})
	if len(scordaturaOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, scordatura := range scordaturaOrdered {

		id = generatesIdentifier("Scordatura", idx, scordatura.Name)
		map_Scordatura_Identifiers[scordatura] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Scordatura")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", scordatura.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(scordatura.Name))
		initializerStatements += setValueField

	}

	map_Score_instrument_Identifiers := make(map[*Score_instrument]string)
	_ = map_Score_instrument_Identifiers

	score_instrumentOrdered := []*Score_instrument{}
	for score_instrument := range stage.Score_instruments {
		score_instrumentOrdered = append(score_instrumentOrdered, score_instrument)
	}
	sort.Slice(score_instrumentOrdered[:], func(i, j int) bool {
		return score_instrumentOrdered[i].Name < score_instrumentOrdered[j].Name
	})
	if len(score_instrumentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, score_instrument := range score_instrumentOrdered {

		id = generatesIdentifier("Score_instrument", idx, score_instrument.Name)
		map_Score_instrument_Identifiers[score_instrument] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Score_instrument")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", score_instrument.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(score_instrument.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Instrument_name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(score_instrument.Instrument_name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Instrument_abbreviation")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(score_instrument.Instrument_abbreviation))
		initializerStatements += setValueField

	}

	map_Score_part_Identifiers := make(map[*Score_part]string)
	_ = map_Score_part_Identifiers

	score_partOrdered := []*Score_part{}
	for score_part := range stage.Score_parts {
		score_partOrdered = append(score_partOrdered, score_part)
	}
	sort.Slice(score_partOrdered[:], func(i, j int) bool {
		return score_partOrdered[i].Name < score_partOrdered[j].Name
	})
	if len(score_partOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, score_part := range score_partOrdered {

		id = generatesIdentifier("Score_part", idx, score_part.Name)
		map_Score_part_Identifiers[score_part] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Score_part")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", score_part.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(score_part.Name))
		initializerStatements += setValueField

	}

	map_Score_partwise_Identifiers := make(map[*Score_partwise]string)
	_ = map_Score_partwise_Identifiers

	score_partwiseOrdered := []*Score_partwise{}
	for score_partwise := range stage.Score_partwises {
		score_partwiseOrdered = append(score_partwiseOrdered, score_partwise)
	}
	sort.Slice(score_partwiseOrdered[:], func(i, j int) bool {
		return score_partwiseOrdered[i].Name < score_partwiseOrdered[j].Name
	})
	if len(score_partwiseOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, score_partwise := range score_partwiseOrdered {

		id = generatesIdentifier("Score_partwise", idx, score_partwise.Name)
		map_Score_partwise_Identifiers[score_partwise] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Score_partwise")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", score_partwise.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(score_partwise.Name))
		initializerStatements += setValueField

	}

	map_Score_timewise_Identifiers := make(map[*Score_timewise]string)
	_ = map_Score_timewise_Identifiers

	score_timewiseOrdered := []*Score_timewise{}
	for score_timewise := range stage.Score_timewises {
		score_timewiseOrdered = append(score_timewiseOrdered, score_timewise)
	}
	sort.Slice(score_timewiseOrdered[:], func(i, j int) bool {
		return score_timewiseOrdered[i].Name < score_timewiseOrdered[j].Name
	})
	if len(score_timewiseOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, score_timewise := range score_timewiseOrdered {

		id = generatesIdentifier("Score_timewise", idx, score_timewise.Name)
		map_Score_timewise_Identifiers[score_timewise] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Score_timewise")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", score_timewise.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(score_timewise.Name))
		initializerStatements += setValueField

	}

	map_Segno_Identifiers := make(map[*Segno]string)
	_ = map_Segno_Identifiers

	segnoOrdered := []*Segno{}
	for segno := range stage.Segnos {
		segnoOrdered = append(segnoOrdered, segno)
	}
	sort.Slice(segnoOrdered[:], func(i, j int) bool {
		return segnoOrdered[i].Name < segnoOrdered[j].Name
	})
	if len(segnoOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, segno := range segnoOrdered {

		id = generatesIdentifier("Segno", idx, segno.Name)
		map_Segno_Identifiers[segno] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Segno")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", segno.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(segno.Name))
		initializerStatements += setValueField

	}

	map_Slash_Identifiers := make(map[*Slash]string)
	_ = map_Slash_Identifiers

	slashOrdered := []*Slash{}
	for slash := range stage.Slashs {
		slashOrdered = append(slashOrdered, slash)
	}
	sort.Slice(slashOrdered[:], func(i, j int) bool {
		return slashOrdered[i].Name < slashOrdered[j].Name
	})
	if len(slashOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, slash := range slashOrdered {

		id = generatesIdentifier("Slash", idx, slash.Name)
		map_Slash_Identifiers[slash] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slash")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", slash.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slash.Name))
		initializerStatements += setValueField

	}

	map_Slide_Identifiers := make(map[*Slide]string)
	_ = map_Slide_Identifiers

	slideOrdered := []*Slide{}
	for slide := range stage.Slides {
		slideOrdered = append(slideOrdered, slide)
	}
	sort.Slice(slideOrdered[:], func(i, j int) bool {
		return slideOrdered[i].Name < slideOrdered[j].Name
	})
	if len(slideOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, slide := range slideOrdered {

		id = generatesIdentifier("Slide", idx, slide.Name)
		map_Slide_Identifiers[slide] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slide")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", slide.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slide.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slide.Value))
		initializerStatements += setValueField

	}

	map_Slur_Identifiers := make(map[*Slur]string)
	_ = map_Slur_Identifiers

	slurOrdered := []*Slur{}
	for slur := range stage.Slurs {
		slurOrdered = append(slurOrdered, slur)
	}
	sort.Slice(slurOrdered[:], func(i, j int) bool {
		return slurOrdered[i].Name < slurOrdered[j].Name
	})
	if len(slurOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, slur := range slurOrdered {

		id = generatesIdentifier("Slur", idx, slur.Name)
		map_Slur_Identifiers[slur] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slur")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", slur.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slur.Name))
		initializerStatements += setValueField

	}

	map_Sound_Identifiers := make(map[*Sound]string)
	_ = map_Sound_Identifiers

	soundOrdered := []*Sound{}
	for sound := range stage.Sounds {
		soundOrdered = append(soundOrdered, sound)
	}
	sort.Slice(soundOrdered[:], func(i, j int) bool {
		return soundOrdered[i].Name < soundOrdered[j].Name
	})
	if len(soundOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, sound := range soundOrdered {

		id = generatesIdentifier("Sound", idx, sound.Name)
		map_Sound_Identifiers[sound] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Sound")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", sound.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sound.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Segno")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sound.Segno))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Dalsegno")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sound.Dalsegno))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Coda")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sound.Coda))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Tocoda")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sound.Tocoda))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fine")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sound.Fine))
		initializerStatements += setValueField

	}

	map_Staff_details_Identifiers := make(map[*Staff_details]string)
	_ = map_Staff_details_Identifiers

	staff_detailsOrdered := []*Staff_details{}
	for staff_details := range stage.Staff_detailss {
		staff_detailsOrdered = append(staff_detailsOrdered, staff_details)
	}
	sort.Slice(staff_detailsOrdered[:], func(i, j int) bool {
		return staff_detailsOrdered[i].Name < staff_detailsOrdered[j].Name
	})
	if len(staff_detailsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, staff_details := range staff_detailsOrdered {

		id = generatesIdentifier("Staff_details", idx, staff_details.Name)
		map_Staff_details_Identifiers[staff_details] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_details")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staff_details.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staff_details.Name))
		initializerStatements += setValueField

	}

	map_Staff_divide_Identifiers := make(map[*Staff_divide]string)
	_ = map_Staff_divide_Identifiers

	staff_divideOrdered := []*Staff_divide{}
	for staff_divide := range stage.Staff_divides {
		staff_divideOrdered = append(staff_divideOrdered, staff_divide)
	}
	sort.Slice(staff_divideOrdered[:], func(i, j int) bool {
		return staff_divideOrdered[i].Name < staff_divideOrdered[j].Name
	})
	if len(staff_divideOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, staff_divide := range staff_divideOrdered {

		id = generatesIdentifier("Staff_divide", idx, staff_divide.Name)
		map_Staff_divide_Identifiers[staff_divide] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_divide")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staff_divide.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staff_divide.Name))
		initializerStatements += setValueField

	}

	map_Staff_layout_Identifiers := make(map[*Staff_layout]string)
	_ = map_Staff_layout_Identifiers

	staff_layoutOrdered := []*Staff_layout{}
	for staff_layout := range stage.Staff_layouts {
		staff_layoutOrdered = append(staff_layoutOrdered, staff_layout)
	}
	sort.Slice(staff_layoutOrdered[:], func(i, j int) bool {
		return staff_layoutOrdered[i].Name < staff_layoutOrdered[j].Name
	})
	if len(staff_layoutOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, staff_layout := range staff_layoutOrdered {

		id = generatesIdentifier("Staff_layout", idx, staff_layout.Name)
		map_Staff_layout_Identifiers[staff_layout] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_layout")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staff_layout.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staff_layout.Name))
		initializerStatements += setValueField

	}

	map_Staff_size_Identifiers := make(map[*Staff_size]string)
	_ = map_Staff_size_Identifiers

	staff_sizeOrdered := []*Staff_size{}
	for staff_size := range stage.Staff_sizes {
		staff_sizeOrdered = append(staff_sizeOrdered, staff_size)
	}
	sort.Slice(staff_sizeOrdered[:], func(i, j int) bool {
		return staff_sizeOrdered[i].Name < staff_sizeOrdered[j].Name
	})
	if len(staff_sizeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, staff_size := range staff_sizeOrdered {

		id = generatesIdentifier("Staff_size", idx, staff_size.Name)
		map_Staff_size_Identifiers[staff_size] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_size")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staff_size.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staff_size.Name))
		initializerStatements += setValueField

	}

	map_Staff_tuning_Identifiers := make(map[*Staff_tuning]string)
	_ = map_Staff_tuning_Identifiers

	staff_tuningOrdered := []*Staff_tuning{}
	for staff_tuning := range stage.Staff_tunings {
		staff_tuningOrdered = append(staff_tuningOrdered, staff_tuning)
	}
	sort.Slice(staff_tuningOrdered[:], func(i, j int) bool {
		return staff_tuningOrdered[i].Name < staff_tuningOrdered[j].Name
	})
	if len(staff_tuningOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, staff_tuning := range staff_tuningOrdered {

		id = generatesIdentifier("Staff_tuning", idx, staff_tuning.Name)
		map_Staff_tuning_Identifiers[staff_tuning] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_tuning")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", staff_tuning.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(staff_tuning.Name))
		initializerStatements += setValueField

	}

	map_Stem_Identifiers := make(map[*Stem]string)
	_ = map_Stem_Identifiers

	stemOrdered := []*Stem{}
	for stem := range stage.Stems {
		stemOrdered = append(stemOrdered, stem)
	}
	sort.Slice(stemOrdered[:], func(i, j int) bool {
		return stemOrdered[i].Name < stemOrdered[j].Name
	})
	if len(stemOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, stem := range stemOrdered {

		id = generatesIdentifier("Stem", idx, stem.Name)
		map_Stem_Identifiers[stem] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Stem")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", stem.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(stem.Name))
		initializerStatements += setValueField

	}

	map_Stick_Identifiers := make(map[*Stick]string)
	_ = map_Stick_Identifiers

	stickOrdered := []*Stick{}
	for stick := range stage.Sticks {
		stickOrdered = append(stickOrdered, stick)
	}
	sort.Slice(stickOrdered[:], func(i, j int) bool {
		return stickOrdered[i].Name < stickOrdered[j].Name
	})
	if len(stickOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, stick := range stickOrdered {

		id = generatesIdentifier("Stick", idx, stick.Name)
		map_Stick_Identifiers[stick] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Stick")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", stick.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(stick.Name))
		initializerStatements += setValueField

	}

	map_String_mute_Identifiers := make(map[*String_mute]string)
	_ = map_String_mute_Identifiers

	string_muteOrdered := []*String_mute{}
	for string_mute := range stage.String_mutes {
		string_muteOrdered = append(string_muteOrdered, string_mute)
	}
	sort.Slice(string_muteOrdered[:], func(i, j int) bool {
		return string_muteOrdered[i].Name < string_muteOrdered[j].Name
	})
	if len(string_muteOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, string_mute := range string_muteOrdered {

		id = generatesIdentifier("String_mute", idx, string_mute.Name)
		map_String_mute_Identifiers[string_mute] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "String_mute")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", string_mute.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(string_mute.Name))
		initializerStatements += setValueField

	}

	map_Strong_accent_Identifiers := make(map[*Strong_accent]string)
	_ = map_Strong_accent_Identifiers

	strong_accentOrdered := []*Strong_accent{}
	for strong_accent := range stage.Strong_accents {
		strong_accentOrdered = append(strong_accentOrdered, strong_accent)
	}
	sort.Slice(strong_accentOrdered[:], func(i, j int) bool {
		return strong_accentOrdered[i].Name < strong_accentOrdered[j].Name
	})
	if len(strong_accentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, strong_accent := range strong_accentOrdered {

		id = generatesIdentifier("Strong_accent", idx, strong_accent.Name)
		map_Strong_accent_Identifiers[strong_accent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Strong_accent")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", strong_accent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(strong_accent.Name))
		initializerStatements += setValueField

	}

	map_Supports_Identifiers := make(map[*Supports]string)
	_ = map_Supports_Identifiers

	supportsOrdered := []*Supports{}
	for supports := range stage.Supportss {
		supportsOrdered = append(supportsOrdered, supports)
	}
	sort.Slice(supportsOrdered[:], func(i, j int) bool {
		return supportsOrdered[i].Name < supportsOrdered[j].Name
	})
	if len(supportsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, supports := range supportsOrdered {

		id = generatesIdentifier("Supports", idx, supports.Name)
		map_Supports_Identifiers[supports] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Supports")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", supports.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(supports.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(supports.Value))
		initializerStatements += setValueField

	}

	map_Swing_Identifiers := make(map[*Swing]string)
	_ = map_Swing_Identifiers

	swingOrdered := []*Swing{}
	for swing := range stage.Swings {
		swingOrdered = append(swingOrdered, swing)
	}
	sort.Slice(swingOrdered[:], func(i, j int) bool {
		return swingOrdered[i].Name < swingOrdered[j].Name
	})
	if len(swingOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, swing := range swingOrdered {

		id = generatesIdentifier("Swing", idx, swing.Name)
		map_Swing_Identifiers[swing] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Swing")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", swing.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(swing.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Swing_style")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(swing.Swing_style))
		initializerStatements += setValueField

	}

	map_Sync_Identifiers := make(map[*Sync]string)
	_ = map_Sync_Identifiers

	syncOrdered := []*Sync{}
	for sync := range stage.Syncs {
		syncOrdered = append(syncOrdered, sync)
	}
	sort.Slice(syncOrdered[:], func(i, j int) bool {
		return syncOrdered[i].Name < syncOrdered[j].Name
	})
	if len(syncOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, sync := range syncOrdered {

		id = generatesIdentifier("Sync", idx, sync.Name)
		map_Sync_Identifiers[sync] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Sync")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", sync.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sync.Name))
		initializerStatements += setValueField

	}

	map_System_dividers_Identifiers := make(map[*System_dividers]string)
	_ = map_System_dividers_Identifiers

	system_dividersOrdered := []*System_dividers{}
	for system_dividers := range stage.System_dividerss {
		system_dividersOrdered = append(system_dividersOrdered, system_dividers)
	}
	sort.Slice(system_dividersOrdered[:], func(i, j int) bool {
		return system_dividersOrdered[i].Name < system_dividersOrdered[j].Name
	})
	if len(system_dividersOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, system_dividers := range system_dividersOrdered {

		id = generatesIdentifier("System_dividers", idx, system_dividers.Name)
		map_System_dividers_Identifiers[system_dividers] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System_dividers")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", system_dividers.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(system_dividers.Name))
		initializerStatements += setValueField

	}

	map_System_layout_Identifiers := make(map[*System_layout]string)
	_ = map_System_layout_Identifiers

	system_layoutOrdered := []*System_layout{}
	for system_layout := range stage.System_layouts {
		system_layoutOrdered = append(system_layoutOrdered, system_layout)
	}
	sort.Slice(system_layoutOrdered[:], func(i, j int) bool {
		return system_layoutOrdered[i].Name < system_layoutOrdered[j].Name
	})
	if len(system_layoutOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, system_layout := range system_layoutOrdered {

		id = generatesIdentifier("System_layout", idx, system_layout.Name)
		map_System_layout_Identifiers[system_layout] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System_layout")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", system_layout.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(system_layout.Name))
		initializerStatements += setValueField

	}

	map_System_margins_Identifiers := make(map[*System_margins]string)
	_ = map_System_margins_Identifiers

	system_marginsOrdered := []*System_margins{}
	for system_margins := range stage.System_marginss {
		system_marginsOrdered = append(system_marginsOrdered, system_margins)
	}
	sort.Slice(system_marginsOrdered[:], func(i, j int) bool {
		return system_marginsOrdered[i].Name < system_marginsOrdered[j].Name
	})
	if len(system_marginsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, system_margins := range system_marginsOrdered {

		id = generatesIdentifier("System_margins", idx, system_margins.Name)
		map_System_margins_Identifiers[system_margins] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System_margins")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", system_margins.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(system_margins.Name))
		initializerStatements += setValueField

	}

	map_Tap_Identifiers := make(map[*Tap]string)
	_ = map_Tap_Identifiers

	tapOrdered := []*Tap{}
	for tap := range stage.Taps {
		tapOrdered = append(tapOrdered, tap)
	}
	sort.Slice(tapOrdered[:], func(i, j int) bool {
		return tapOrdered[i].Name < tapOrdered[j].Name
	})
	if len(tapOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tap := range tapOrdered {

		id = generatesIdentifier("Tap", idx, tap.Name)
		map_Tap_Identifiers[tap] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tap")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tap.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tap.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tap.Value))
		initializerStatements += setValueField

	}

	map_Technical_Identifiers := make(map[*Technical]string)
	_ = map_Technical_Identifiers

	technicalOrdered := []*Technical{}
	for technical := range stage.Technicals {
		technicalOrdered = append(technicalOrdered, technical)
	}
	sort.Slice(technicalOrdered[:], func(i, j int) bool {
		return technicalOrdered[i].Name < technicalOrdered[j].Name
	})
	if len(technicalOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, technical := range technicalOrdered {

		id = generatesIdentifier("Technical", idx, technical.Name)
		map_Technical_Identifiers[technical] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Technical")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", technical.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(technical.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Astring")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(technical.Astring))
		initializerStatements += setValueField

	}

	map_Text_element_data_Identifiers := make(map[*Text_element_data]string)
	_ = map_Text_element_data_Identifiers

	text_element_dataOrdered := []*Text_element_data{}
	for text_element_data := range stage.Text_element_datas {
		text_element_dataOrdered = append(text_element_dataOrdered, text_element_data)
	}
	sort.Slice(text_element_dataOrdered[:], func(i, j int) bool {
		return text_element_dataOrdered[i].Name < text_element_dataOrdered[j].Name
	})
	if len(text_element_dataOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, text_element_data := range text_element_dataOrdered {

		id = generatesIdentifier("Text_element_data", idx, text_element_data.Name)
		map_Text_element_data_Identifiers[text_element_data] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Text_element_data")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", text_element_data.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text_element_data.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text_element_data.Value))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EmptyString")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text_element_data.EmptyString))
		initializerStatements += setValueField

	}

	map_Tie_Identifiers := make(map[*Tie]string)
	_ = map_Tie_Identifiers

	tieOrdered := []*Tie{}
	for tie := range stage.Ties {
		tieOrdered = append(tieOrdered, tie)
	}
	sort.Slice(tieOrdered[:], func(i, j int) bool {
		return tieOrdered[i].Name < tieOrdered[j].Name
	})
	if len(tieOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tie := range tieOrdered {

		id = generatesIdentifier("Tie", idx, tie.Name)
		map_Tie_Identifiers[tie] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tie")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tie.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tie.Name))
		initializerStatements += setValueField

	}

	map_Tied_Identifiers := make(map[*Tied]string)
	_ = map_Tied_Identifiers

	tiedOrdered := []*Tied{}
	for tied := range stage.Tieds {
		tiedOrdered = append(tiedOrdered, tied)
	}
	sort.Slice(tiedOrdered[:], func(i, j int) bool {
		return tiedOrdered[i].Name < tiedOrdered[j].Name
	})
	if len(tiedOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tied := range tiedOrdered {

		id = generatesIdentifier("Tied", idx, tied.Name)
		map_Tied_Identifiers[tied] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tied")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tied.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tied.Name))
		initializerStatements += setValueField

	}

	map_Time_Identifiers := make(map[*Time]string)
	_ = map_Time_Identifiers

	timeOrdered := []*Time{}
	for time := range stage.Times {
		timeOrdered = append(timeOrdered, time)
	}
	sort.Slice(timeOrdered[:], func(i, j int) bool {
		return timeOrdered[i].Name < timeOrdered[j].Name
	})
	if len(timeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, time := range timeOrdered {

		id = generatesIdentifier("Time", idx, time.Name)
		map_Time_Identifiers[time] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Time")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", time.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(time.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Senza_misura")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(time.Senza_misura))
		initializerStatements += setValueField

	}

	map_Time_modification_Identifiers := make(map[*Time_modification]string)
	_ = map_Time_modification_Identifiers

	time_modificationOrdered := []*Time_modification{}
	for time_modification := range stage.Time_modifications {
		time_modificationOrdered = append(time_modificationOrdered, time_modification)
	}
	sort.Slice(time_modificationOrdered[:], func(i, j int) bool {
		return time_modificationOrdered[i].Name < time_modificationOrdered[j].Name
	})
	if len(time_modificationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, time_modification := range time_modificationOrdered {

		id = generatesIdentifier("Time_modification", idx, time_modification.Name)
		map_Time_modification_Identifiers[time_modification] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Time_modification")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", time_modification.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(time_modification.Name))
		initializerStatements += setValueField

	}

	map_Timpani_Identifiers := make(map[*Timpani]string)
	_ = map_Timpani_Identifiers

	timpaniOrdered := []*Timpani{}
	for timpani := range stage.Timpanis {
		timpaniOrdered = append(timpaniOrdered, timpani)
	}
	sort.Slice(timpaniOrdered[:], func(i, j int) bool {
		return timpaniOrdered[i].Name < timpaniOrdered[j].Name
	})
	if len(timpaniOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, timpani := range timpaniOrdered {

		id = generatesIdentifier("Timpani", idx, timpani.Name)
		map_Timpani_Identifiers[timpani] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Timpani")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", timpani.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(timpani.Name))
		initializerStatements += setValueField

	}

	map_Transpose_Identifiers := make(map[*Transpose]string)
	_ = map_Transpose_Identifiers

	transposeOrdered := []*Transpose{}
	for transpose := range stage.Transposes {
		transposeOrdered = append(transposeOrdered, transpose)
	}
	sort.Slice(transposeOrdered[:], func(i, j int) bool {
		return transposeOrdered[i].Name < transposeOrdered[j].Name
	})
	if len(transposeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, transpose := range transposeOrdered {

		id = generatesIdentifier("Transpose", idx, transpose.Name)
		map_Transpose_Identifiers[transpose] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transpose")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", transpose.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(transpose.Name))
		initializerStatements += setValueField

	}

	map_Tremolo_Identifiers := make(map[*Tremolo]string)
	_ = map_Tremolo_Identifiers

	tremoloOrdered := []*Tremolo{}
	for tremolo := range stage.Tremolos {
		tremoloOrdered = append(tremoloOrdered, tremolo)
	}
	sort.Slice(tremoloOrdered[:], func(i, j int) bool {
		return tremoloOrdered[i].Name < tremoloOrdered[j].Name
	})
	if len(tremoloOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tremolo := range tremoloOrdered {

		id = generatesIdentifier("Tremolo", idx, tremolo.Name)
		map_Tremolo_Identifiers[tremolo] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tremolo")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tremolo.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tremolo.Name))
		initializerStatements += setValueField

	}

	map_Tuplet_Identifiers := make(map[*Tuplet]string)
	_ = map_Tuplet_Identifiers

	tupletOrdered := []*Tuplet{}
	for tuplet := range stage.Tuplets {
		tupletOrdered = append(tupletOrdered, tuplet)
	}
	sort.Slice(tupletOrdered[:], func(i, j int) bool {
		return tupletOrdered[i].Name < tupletOrdered[j].Name
	})
	if len(tupletOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tuplet := range tupletOrdered {

		id = generatesIdentifier("Tuplet", idx, tuplet.Name)
		map_Tuplet_Identifiers[tuplet] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tuplet.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tuplet.Name))
		initializerStatements += setValueField

	}

	map_Tuplet_dot_Identifiers := make(map[*Tuplet_dot]string)
	_ = map_Tuplet_dot_Identifiers

	tuplet_dotOrdered := []*Tuplet_dot{}
	for tuplet_dot := range stage.Tuplet_dots {
		tuplet_dotOrdered = append(tuplet_dotOrdered, tuplet_dot)
	}
	sort.Slice(tuplet_dotOrdered[:], func(i, j int) bool {
		return tuplet_dotOrdered[i].Name < tuplet_dotOrdered[j].Name
	})
	if len(tuplet_dotOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tuplet_dot := range tuplet_dotOrdered {

		id = generatesIdentifier("Tuplet_dot", idx, tuplet_dot.Name)
		map_Tuplet_dot_Identifiers[tuplet_dot] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet_dot")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tuplet_dot.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tuplet_dot.Name))
		initializerStatements += setValueField

	}

	map_Tuplet_number_Identifiers := make(map[*Tuplet_number]string)
	_ = map_Tuplet_number_Identifiers

	tuplet_numberOrdered := []*Tuplet_number{}
	for tuplet_number := range stage.Tuplet_numbers {
		tuplet_numberOrdered = append(tuplet_numberOrdered, tuplet_number)
	}
	sort.Slice(tuplet_numberOrdered[:], func(i, j int) bool {
		return tuplet_numberOrdered[i].Name < tuplet_numberOrdered[j].Name
	})
	if len(tuplet_numberOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tuplet_number := range tuplet_numberOrdered {

		id = generatesIdentifier("Tuplet_number", idx, tuplet_number.Name)
		map_Tuplet_number_Identifiers[tuplet_number] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet_number")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tuplet_number.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tuplet_number.Name))
		initializerStatements += setValueField

	}

	map_Tuplet_portion_Identifiers := make(map[*Tuplet_portion]string)
	_ = map_Tuplet_portion_Identifiers

	tuplet_portionOrdered := []*Tuplet_portion{}
	for tuplet_portion := range stage.Tuplet_portions {
		tuplet_portionOrdered = append(tuplet_portionOrdered, tuplet_portion)
	}
	sort.Slice(tuplet_portionOrdered[:], func(i, j int) bool {
		return tuplet_portionOrdered[i].Name < tuplet_portionOrdered[j].Name
	})
	if len(tuplet_portionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tuplet_portion := range tuplet_portionOrdered {

		id = generatesIdentifier("Tuplet_portion", idx, tuplet_portion.Name)
		map_Tuplet_portion_Identifiers[tuplet_portion] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet_portion")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tuplet_portion.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tuplet_portion.Name))
		initializerStatements += setValueField

	}

	map_Tuplet_type_Identifiers := make(map[*Tuplet_type]string)
	_ = map_Tuplet_type_Identifiers

	tuplet_typeOrdered := []*Tuplet_type{}
	for tuplet_type := range stage.Tuplet_types {
		tuplet_typeOrdered = append(tuplet_typeOrdered, tuplet_type)
	}
	sort.Slice(tuplet_typeOrdered[:], func(i, j int) bool {
		return tuplet_typeOrdered[i].Name < tuplet_typeOrdered[j].Name
	})
	if len(tuplet_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tuplet_type := range tuplet_typeOrdered {

		id = generatesIdentifier("Tuplet_type", idx, tuplet_type.Name)
		map_Tuplet_type_Identifiers[tuplet_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet_type")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tuplet_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tuplet_type.Name))
		initializerStatements += setValueField

	}

	map_Typed_text_Identifiers := make(map[*Typed_text]string)
	_ = map_Typed_text_Identifiers

	typed_textOrdered := []*Typed_text{}
	for typed_text := range stage.Typed_texts {
		typed_textOrdered = append(typed_textOrdered, typed_text)
	}
	sort.Slice(typed_textOrdered[:], func(i, j int) bool {
		return typed_textOrdered[i].Name < typed_textOrdered[j].Name
	})
	if len(typed_textOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, typed_text := range typed_textOrdered {

		id = generatesIdentifier("Typed_text", idx, typed_text.Name)
		map_Typed_text_Identifiers[typed_text] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Typed_text")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", typed_text.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(typed_text.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(typed_text.Value))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(typed_text.Type))
		initializerStatements += setValueField

	}

	map_Unpitched_Identifiers := make(map[*Unpitched]string)
	_ = map_Unpitched_Identifiers

	unpitchedOrdered := []*Unpitched{}
	for unpitched := range stage.Unpitcheds {
		unpitchedOrdered = append(unpitchedOrdered, unpitched)
	}
	sort.Slice(unpitchedOrdered[:], func(i, j int) bool {
		return unpitchedOrdered[i].Name < unpitchedOrdered[j].Name
	})
	if len(unpitchedOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, unpitched := range unpitchedOrdered {

		id = generatesIdentifier("Unpitched", idx, unpitched.Name)
		map_Unpitched_Identifiers[unpitched] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Unpitched")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", unpitched.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(unpitched.Name))
		initializerStatements += setValueField

	}

	map_Virtual_instrument_Identifiers := make(map[*Virtual_instrument]string)
	_ = map_Virtual_instrument_Identifiers

	virtual_instrumentOrdered := []*Virtual_instrument{}
	for virtual_instrument := range stage.Virtual_instruments {
		virtual_instrumentOrdered = append(virtual_instrumentOrdered, virtual_instrument)
	}
	sort.Slice(virtual_instrumentOrdered[:], func(i, j int) bool {
		return virtual_instrumentOrdered[i].Name < virtual_instrumentOrdered[j].Name
	})
	if len(virtual_instrumentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, virtual_instrument := range virtual_instrumentOrdered {

		id = generatesIdentifier("Virtual_instrument", idx, virtual_instrument.Name)
		map_Virtual_instrument_Identifiers[virtual_instrument] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Virtual_instrument")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", virtual_instrument.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(virtual_instrument.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Virtual_library")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(virtual_instrument.Virtual_library))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Virtual_name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(virtual_instrument.Virtual_name))
		initializerStatements += setValueField

	}

	map_Wait_Identifiers := make(map[*Wait]string)
	_ = map_Wait_Identifiers

	waitOrdered := []*Wait{}
	for wait := range stage.Waits {
		waitOrdered = append(waitOrdered, wait)
	}
	sort.Slice(waitOrdered[:], func(i, j int) bool {
		return waitOrdered[i].Name < waitOrdered[j].Name
	})
	if len(waitOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, wait := range waitOrdered {

		id = generatesIdentifier("Wait", idx, wait.Name)
		map_Wait_Identifiers[wait] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Wait")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", wait.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(wait.Name))
		initializerStatements += setValueField

	}

	map_Wavy_line_Identifiers := make(map[*Wavy_line]string)
	_ = map_Wavy_line_Identifiers

	wavy_lineOrdered := []*Wavy_line{}
	for wavy_line := range stage.Wavy_lines {
		wavy_lineOrdered = append(wavy_lineOrdered, wavy_line)
	}
	sort.Slice(wavy_lineOrdered[:], func(i, j int) bool {
		return wavy_lineOrdered[i].Name < wavy_lineOrdered[j].Name
	})
	if len(wavy_lineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, wavy_line := range wavy_lineOrdered {

		id = generatesIdentifier("Wavy_line", idx, wavy_line.Name)
		map_Wavy_line_Identifiers[wavy_line] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Wavy_line")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", wavy_line.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(wavy_line.Name))
		initializerStatements += setValueField

	}

	map_Wedge_Identifiers := make(map[*Wedge]string)
	_ = map_Wedge_Identifiers

	wedgeOrdered := []*Wedge{}
	for wedge := range stage.Wedges {
		wedgeOrdered = append(wedgeOrdered, wedge)
	}
	sort.Slice(wedgeOrdered[:], func(i, j int) bool {
		return wedgeOrdered[i].Name < wedgeOrdered[j].Name
	})
	if len(wedgeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, wedge := range wedgeOrdered {

		id = generatesIdentifier("Wedge", idx, wedge.Name)
		map_Wedge_Identifiers[wedge] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Wedge")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", wedge.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(wedge.Name))
		initializerStatements += setValueField

	}

	map_Wood_Identifiers := make(map[*Wood]string)
	_ = map_Wood_Identifiers

	woodOrdered := []*Wood{}
	for wood := range stage.Woods {
		woodOrdered = append(woodOrdered, wood)
	}
	sort.Slice(woodOrdered[:], func(i, j int) bool {
		return woodOrdered[i].Name < woodOrdered[j].Name
	})
	if len(woodOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, wood := range woodOrdered {

		id = generatesIdentifier("Wood", idx, wood.Name)
		map_Wood_Identifiers[wood] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Wood")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", wood.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(wood.Name))
		initializerStatements += setValueField

	}

	map_Work_Identifiers := make(map[*Work]string)
	_ = map_Work_Identifiers

	workOrdered := []*Work{}
	for work := range stage.Works {
		workOrdered = append(workOrdered, work)
	}
	sort.Slice(workOrdered[:], func(i, j int) bool {
		return workOrdered[i].Name < workOrdered[j].Name
	})
	if len(workOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, work := range workOrdered {

		id = generatesIdentifier("Work", idx, work.Name)
		map_Work_Identifiers[work] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Work")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", work.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(work.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Work_number")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(work.Work_number))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Work_title")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(work.Work_title))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, accidental := range accidentalOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Accidental", idx, accidental.Name)
		map_Accidental_Identifiers[accidental] = id

		// Initialisation of values
	}

	for idx, accidental_mark := range accidental_markOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Accidental_mark", idx, accidental_mark.Name)
		map_Accidental_mark_Identifiers[accidental_mark] = id

		// Initialisation of values
	}

	for idx, accidental_text := range accidental_textOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Accidental_text", idx, accidental_text.Name)
		map_Accidental_text_Identifiers[accidental_text] = id

		// Initialisation of values
	}

	for idx, accord := range accordOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Accord", idx, accord.Name)
		map_Accord_Identifiers[accord] = id

		// Initialisation of values
	}

	for idx, accordion_registration := range accordion_registrationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Accordion_registration", idx, accordion_registration.Name)
		map_Accordion_registration_Identifiers[accordion_registration] = id

		// Initialisation of values
		if accordion_registration.Accordion_high != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accordion_high")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[accordion_registration.Accordion_high])
			pointersInitializesStatements += setPointerField
		}

		if accordion_registration.Accordion_low != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accordion_low")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[accordion_registration.Accordion_low])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, anytype := range anytypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AnyType", idx, anytype.Name)
		map_AnyType_Identifiers[anytype] = id

		// Initialisation of values
	}

	for idx, appearance := range appearanceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Appearance", idx, appearance.Name)
		map_Appearance_Identifiers[appearance] = id

		// Initialisation of values
		for _, _line_width := range appearance.Line_width {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Line_width")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Line_width_Identifiers[_line_width])
			pointersInitializesStatements += setPointerField
		}

		for _, _note_size := range appearance.Note_size {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Note_size")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Note_size_Identifiers[_note_size])
			pointersInitializesStatements += setPointerField
		}

		for _, _distance := range appearance.Distance {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Distance")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Distance_Identifiers[_distance])
			pointersInitializesStatements += setPointerField
		}

		for _, _glyph := range appearance.Glyph {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Glyph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Glyph_Identifiers[_glyph])
			pointersInitializesStatements += setPointerField
		}

		for _, _other_appearance := range appearance.Other_appearance {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Other_appearance")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Other_appearance_Identifiers[_other_appearance])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, arpeggiate := range arpeggiateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Arpeggiate", idx, arpeggiate.Name)
		map_Arpeggiate_Identifiers[arpeggiate] = id

		// Initialisation of values
	}

	for idx, arrow := range arrowOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Arrow", idx, arrow.Name)
		map_Arrow_Identifiers[arrow] = id

		// Initialisation of values
	}

	for idx, articulations := range articulationsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Articulations", idx, articulations.Name)
		map_Articulations_Identifiers[articulations] = id

		// Initialisation of values
		if articulations.Accent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[articulations.Accent])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Strong_accent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Strong_accent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Strong_accent_Identifiers[articulations.Strong_accent])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Staccato != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Staccato")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[articulations.Staccato])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Tenuto != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tenuto")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[articulations.Tenuto])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Detached_legato != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Detached_legato")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[articulations.Detached_legato])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Staccatissimo != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Staccatissimo")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[articulations.Staccatissimo])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Spiccato != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Spiccato")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[articulations.Spiccato])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Scoop != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Scoop")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_line_Identifiers[articulations.Scoop])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Plop != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Plop")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_line_Identifiers[articulations.Plop])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Doit != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Doit")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_line_Identifiers[articulations.Doit])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Falloff != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Falloff")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_line_Identifiers[articulations.Falloff])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Breath_mark != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Breath_mark")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Breath_mark_Identifiers[articulations.Breath_mark])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Caesura != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Caesura")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Caesura_Identifiers[articulations.Caesura])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Stress != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Stress")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[articulations.Stress])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Unstress != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Unstress")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[articulations.Unstress])
			pointersInitializesStatements += setPointerField
		}

		if articulations.Soft_accent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Soft_accent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[articulations.Soft_accent])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, assess := range assessOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Assess", idx, assess.Name)
		map_Assess_Identifiers[assess] = id

		// Initialisation of values
	}

	for idx, attributes := range attributesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Attributes", idx, attributes.Name)
		map_Attributes_Identifiers[attributes] = id

		// Initialisation of values
		for _, _key := range attributes.Key {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Key")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Key_Identifiers[_key])
			pointersInitializesStatements += setPointerField
		}

		if attributes.Part_symbol != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Part_symbol")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Part_symbol_Identifiers[attributes.Part_symbol])
			pointersInitializesStatements += setPointerField
		}

		for _, _clef := range attributes.Clef {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Clef")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Clef_Identifiers[_clef])
			pointersInitializesStatements += setPointerField
		}

		for _, _staff_details := range attributes.Staff_details {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Staff_details")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Staff_details_Identifiers[_staff_details])
			pointersInitializesStatements += setPointerField
		}

		for _, _measure_style := range attributes.Measure_style {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Measure_style")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Measure_style_Identifiers[_measure_style])
			pointersInitializesStatements += setPointerField
		}

		for _, _transpose := range attributes.Transpose {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Transpose")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Transpose_Identifiers[_transpose])
			pointersInitializesStatements += setPointerField
		}

		for _, _for_part := range attributes.For_part {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "For_part")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_For_part_Identifiers[_for_part])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, backup := range backupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Backup", idx, backup.Name)
		map_Backup_Identifiers[backup] = id

		// Initialisation of values
	}

	for idx, bar_style_color := range bar_style_colorOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bar_style_color", idx, bar_style_color.Name)
		map_Bar_style_color_Identifiers[bar_style_color] = id

		// Initialisation of values
	}

	for idx, barline := range barlineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Barline", idx, barline.Name)
		map_Barline_Identifiers[barline] = id

		// Initialisation of values
		if barline.Bar_style != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bar_style")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bar_style_color_Identifiers[barline.Bar_style])
			pointersInitializesStatements += setPointerField
		}

		if barline.Wavy_line != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Wavy_line")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Wavy_line_Identifiers[barline.Wavy_line])
			pointersInitializesStatements += setPointerField
		}

		if barline.Fermata != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fermata")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Fermata_Identifiers[barline.Fermata])
			pointersInitializesStatements += setPointerField
		}

		if barline.Ending != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Ending")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Ending_Identifiers[barline.Ending])
			pointersInitializesStatements += setPointerField
		}

		if barline.Repeat != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Repeat")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Repeat_Identifiers[barline.Repeat])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, barre := range barreOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Barre", idx, barre.Name)
		map_Barre_Identifiers[barre] = id

		// Initialisation of values
	}

	for idx, bass := range bassOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bass", idx, bass.Name)
		map_Bass_Identifiers[bass] = id

		// Initialisation of values
		if bass.Bass_step != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bass_step")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bass_step_Identifiers[bass.Bass_step])
			pointersInitializesStatements += setPointerField
		}

		if bass.Bass_alter != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bass_alter")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Harmony_alter_Identifiers[bass.Bass_alter])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, bass_step := range bass_stepOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bass_step", idx, bass_step.Name)
		map_Bass_step_Identifiers[bass_step] = id

		// Initialisation of values
	}

	for idx, beam := range beamOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Beam", idx, beam.Name)
		map_Beam_Identifiers[beam] = id

		// Initialisation of values
	}

	for idx, beat_repeat := range beat_repeatOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Beat_repeat", idx, beat_repeat.Name)
		map_Beat_repeat_Identifiers[beat_repeat] = id

		// Initialisation of values
	}

	for idx, beat_unit_tied := range beat_unit_tiedOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Beat_unit_tied", idx, beat_unit_tied.Name)
		map_Beat_unit_tied_Identifiers[beat_unit_tied] = id

		// Initialisation of values
	}

	for idx, beater := range beaterOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Beater", idx, beater.Name)
		map_Beater_Identifiers[beater] = id

		// Initialisation of values
	}

	for idx, bend := range bendOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bend", idx, bend.Name)
		map_Bend_Identifiers[bend] = id

		// Initialisation of values
		if bend.Pre_bend != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pre_bend")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[bend.Pre_bend])
			pointersInitializesStatements += setPointerField
		}

		if bend.Release != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Release")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Release_Identifiers[bend.Release])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, bookmark := range bookmarkOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bookmark", idx, bookmark.Name)
		map_Bookmark_Identifiers[bookmark] = id

		// Initialisation of values
	}

	for idx, bracket := range bracketOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bracket", idx, bracket.Name)
		map_Bracket_Identifiers[bracket] = id

		// Initialisation of values
	}

	for idx, breath_mark := range breath_markOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Breath_mark", idx, breath_mark.Name)
		map_Breath_mark_Identifiers[breath_mark] = id

		// Initialisation of values
	}

	for idx, caesura := range caesuraOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Caesura", idx, caesura.Name)
		map_Caesura_Identifiers[caesura] = id

		// Initialisation of values
	}

	for idx, cancel := range cancelOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Cancel", idx, cancel.Name)
		map_Cancel_Identifiers[cancel] = id

		// Initialisation of values
	}

	for idx, clef := range clefOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Clef", idx, clef.Name)
		map_Clef_Identifiers[clef] = id

		// Initialisation of values
	}

	for idx, coda := range codaOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Coda", idx, coda.Name)
		map_Coda_Identifiers[coda] = id

		// Initialisation of values
	}

	for idx, credit := range creditOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Credit", idx, credit.Name)
		map_Credit_Identifiers[credit] = id

		// Initialisation of values
		for _, _link := range credit.Link {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Link")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Link_Identifiers[_link])
			pointersInitializesStatements += setPointerField
		}

		for _, _bookmark := range credit.Bookmark {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bookmark")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bookmark_Identifiers[_bookmark])
			pointersInitializesStatements += setPointerField
		}

		if credit.Credit_image != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Credit_image")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Image_Identifiers[credit.Credit_image])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, dashes := range dashesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Dashes", idx, dashes.Name)
		map_Dashes_Identifiers[dashes] = id

		// Initialisation of values
	}

	for idx, defaults := range defaultsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Defaults", idx, defaults.Name)
		map_Defaults_Identifiers[defaults] = id

		// Initialisation of values
		if defaults.Scaling != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Scaling")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Scaling_Identifiers[defaults.Scaling])
			pointersInitializesStatements += setPointerField
		}

		if defaults.Concert_score != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Concert_score")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[defaults.Concert_score])
			pointersInitializesStatements += setPointerField
		}

		if defaults.Appearance != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Appearance")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Appearance_Identifiers[defaults.Appearance])
			pointersInitializesStatements += setPointerField
		}

		if defaults.Music_font != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Music_font")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_font_Identifiers[defaults.Music_font])
			pointersInitializesStatements += setPointerField
		}

		if defaults.Word_font != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Word_font")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_font_Identifiers[defaults.Word_font])
			pointersInitializesStatements += setPointerField
		}

		for _, _lyric_font := range defaults.Lyric_font {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lyric_font")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lyric_font_Identifiers[_lyric_font])
			pointersInitializesStatements += setPointerField
		}

		for _, _lyric_language := range defaults.Lyric_language {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lyric_language")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lyric_language_Identifiers[_lyric_language])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, degree := range degreeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Degree", idx, degree.Name)
		map_Degree_Identifiers[degree] = id

		// Initialisation of values
		if degree.Degree_value != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Degree_value")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Degree_value_Identifiers[degree.Degree_value])
			pointersInitializesStatements += setPointerField
		}

		if degree.Degree_alter != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Degree_alter")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Degree_alter_Identifiers[degree.Degree_alter])
			pointersInitializesStatements += setPointerField
		}

		if degree.Degree_type != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Degree_type")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Degree_type_Identifiers[degree.Degree_type])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, degree_alter := range degree_alterOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Degree_alter", idx, degree_alter.Name)
		map_Degree_alter_Identifiers[degree_alter] = id

		// Initialisation of values
	}

	for idx, degree_type := range degree_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Degree_type", idx, degree_type.Name)
		map_Degree_type_Identifiers[degree_type] = id

		// Initialisation of values
	}

	for idx, degree_value := range degree_valueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Degree_value", idx, degree_value.Name)
		map_Degree_value_Identifiers[degree_value] = id

		// Initialisation of values
	}

	for idx, direction := range directionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Direction", idx, direction.Name)
		map_Direction_Identifiers[direction] = id

		// Initialisation of values
		for _, _direction_type := range direction.Direction_type {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Direction_type")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Direction_type_Identifiers[_direction_type])
			pointersInitializesStatements += setPointerField
		}

		if direction.Offset != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Offset")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Offset_Identifiers[direction.Offset])
			pointersInitializesStatements += setPointerField
		}

		if direction.Sound != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sound")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sound_Identifiers[direction.Sound])
			pointersInitializesStatements += setPointerField
		}

		if direction.Listening != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Listening")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Listening_Identifiers[direction.Listening])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, direction_type := range direction_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Direction_type", idx, direction_type.Name)
		map_Direction_type_Identifiers[direction_type] = id

		// Initialisation of values
		for _, _segno := range direction_type.Segno {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Segno")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Segno_Identifiers[_segno])
			pointersInitializesStatements += setPointerField
		}

		for _, _coda := range direction_type.Coda {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Coda")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Coda_Identifiers[_coda])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Wedge != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Wedge")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Wedge_Identifiers[direction_type.Wedge])
			pointersInitializesStatements += setPointerField
		}

		for _, _dynamics := range direction_type.Dynamics {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dynamics")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Dynamics_Identifiers[_dynamics])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Dashes != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dashes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Dashes_Identifiers[direction_type.Dashes])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Bracket != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bracket")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bracket_Identifiers[direction_type.Bracket])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Pedal != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pedal")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Pedal_Identifiers[direction_type.Pedal])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Metronome != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Metronome")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Metronome_Identifiers[direction_type.Metronome])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Octave_shift != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Octave_shift")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Octave_shift_Identifiers[direction_type.Octave_shift])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Harp_pedals != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Harp_pedals")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Harp_pedals_Identifiers[direction_type.Harp_pedals])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Damp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Damp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_print_style_align_id_Identifiers[direction_type.Damp])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Damp_all != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Damp_all")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_print_style_align_id_Identifiers[direction_type.Damp_all])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Eyeglasses != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Eyeglasses")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_print_style_align_id_Identifiers[direction_type.Eyeglasses])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.String_mute != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "String_mute")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_String_mute_Identifiers[direction_type.String_mute])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Scordatura != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Scordatura")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Scordatura_Identifiers[direction_type.Scordatura])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Image != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Image")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Image_Identifiers[direction_type.Image])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Principal_voice != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Principal_voice")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Principal_voice_Identifiers[direction_type.Principal_voice])
			pointersInitializesStatements += setPointerField
		}

		for _, _percussion := range direction_type.Percussion {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Percussion")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Percussion_Identifiers[_percussion])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Accordion_registration != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accordion_registration")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Accordion_registration_Identifiers[direction_type.Accordion_registration])
			pointersInitializesStatements += setPointerField
		}

		if direction_type.Staff_divide != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Staff_divide")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Staff_divide_Identifiers[direction_type.Staff_divide])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, distance := range distanceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Distance", idx, distance.Name)
		map_Distance_Identifiers[distance] = id

		// Initialisation of values
	}

	for idx, double := range doubleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Double", idx, double.Name)
		map_Double_Identifiers[double] = id

		// Initialisation of values
	}

	for idx, dynamics := range dynamicsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Dynamics", idx, dynamics.Name)
		map_Dynamics_Identifiers[dynamics] = id

		// Initialisation of values
		if dynamics.P != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "P")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.P])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Pp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Pp])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Ppp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Ppp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Ppp])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Pppp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pppp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Pppp])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Ppppp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Ppppp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Ppppp])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Pppppp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pppppp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Pppppp])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.F != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "F")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.F])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Ff != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Ff")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Ff])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Fff != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fff")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Fff])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Ffff != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Ffff")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Ffff])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Fffff != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fffff")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Fffff])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Ffffff != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Ffffff")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Ffffff])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Mp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Mp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Mp])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Mf != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Mf")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Mf])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Sf != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sf")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Sf])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Sfp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sfp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Sfp])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Sfpp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sfpp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Sfpp])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Fp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Fp])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Rf != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rf")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Rf])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Rfz != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rfz")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Rfz])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Sfz != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sfz")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Sfz])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Sffz != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sffz")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Sffz])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Fz != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fz")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Fz])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.N != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "N")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.N])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Pf != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pf")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Pf])
			pointersInitializesStatements += setPointerField
		}

		if dynamics.Sfzp != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sfzp")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[dynamics.Sfzp])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, effect := range effectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Effect", idx, effect.Name)
		map_Effect_Identifiers[effect] = id

		// Initialisation of values
	}

	for idx, elision := range elisionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Elision", idx, elision.Name)
		map_Elision_Identifiers[elision] = id

		// Initialisation of values
	}

	for idx, empty := range emptyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty", idx, empty.Name)
		map_Empty_Identifiers[empty] = id

		// Initialisation of values
	}

	for idx, empty_font := range empty_fontOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty_font", idx, empty_font.Name)
		map_Empty_font_Identifiers[empty_font] = id

		// Initialisation of values
	}

	for idx, empty_line := range empty_lineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty_line", idx, empty_line.Name)
		map_Empty_line_Identifiers[empty_line] = id

		// Initialisation of values
	}

	for idx, empty_placement := range empty_placementOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty_placement", idx, empty_placement.Name)
		map_Empty_placement_Identifiers[empty_placement] = id

		// Initialisation of values
	}

	for idx, empty_placement_smufl := range empty_placement_smuflOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty_placement_smufl", idx, empty_placement_smufl.Name)
		map_Empty_placement_smufl_Identifiers[empty_placement_smufl] = id

		// Initialisation of values
	}

	for idx, empty_print_object_style_align := range empty_print_object_style_alignOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty_print_object_style_align", idx, empty_print_object_style_align.Name)
		map_Empty_print_object_style_align_Identifiers[empty_print_object_style_align] = id

		// Initialisation of values
	}

	for idx, empty_print_style := range empty_print_styleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty_print_style", idx, empty_print_style.Name)
		map_Empty_print_style_Identifiers[empty_print_style] = id

		// Initialisation of values
	}

	for idx, empty_print_style_align := range empty_print_style_alignOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty_print_style_align", idx, empty_print_style_align.Name)
		map_Empty_print_style_align_Identifiers[empty_print_style_align] = id

		// Initialisation of values
	}

	for idx, empty_print_style_align_id := range empty_print_style_align_idOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty_print_style_align_id", idx, empty_print_style_align_id.Name)
		map_Empty_print_style_align_id_Identifiers[empty_print_style_align_id] = id

		// Initialisation of values
	}

	for idx, empty_trill_sound := range empty_trill_soundOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Empty_trill_sound", idx, empty_trill_sound.Name)
		map_Empty_trill_sound_Identifiers[empty_trill_sound] = id

		// Initialisation of values
	}

	for idx, encoding := range encodingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Encoding", idx, encoding.Name)
		map_Encoding_Identifiers[encoding] = id

		// Initialisation of values
		if encoding.Encoder != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Encoder")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Typed_text_Identifiers[encoding.Encoder])
			pointersInitializesStatements += setPointerField
		}

		if encoding.Supports != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Supports")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Supports_Identifiers[encoding.Supports])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, ending := range endingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Ending", idx, ending.Name)
		map_Ending_Identifiers[ending] = id

		// Initialisation of values
	}

	for idx, extend := range extendOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Extend", idx, extend.Name)
		map_Extend_Identifiers[extend] = id

		// Initialisation of values
	}

	for idx, feature := range featureOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Feature", idx, feature.Name)
		map_Feature_Identifiers[feature] = id

		// Initialisation of values
	}

	for idx, fermata := range fermataOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Fermata", idx, fermata.Name)
		map_Fermata_Identifiers[fermata] = id

		// Initialisation of values
	}

	for idx, figure := range figureOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Figure", idx, figure.Name)
		map_Figure_Identifiers[figure] = id

		// Initialisation of values
		if figure.Extend != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Extend")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Extend_Identifiers[figure.Extend])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, figured_bass := range figured_bassOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Figured_bass", idx, figured_bass.Name)
		map_Figured_bass_Identifiers[figured_bass] = id

		// Initialisation of values
		for _, _figure := range figured_bass.Figure {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Figure")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Figure_Identifiers[_figure])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, fingering := range fingeringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Fingering", idx, fingering.Name)
		map_Fingering_Identifiers[fingering] = id

		// Initialisation of values
	}

	for idx, first_fret := range first_fretOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("First_fret", idx, first_fret.Name)
		map_First_fret_Identifiers[first_fret] = id

		// Initialisation of values
	}

	for idx, foo := range fooOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Foo", idx, foo.Name)
		map_Foo_Identifiers[foo] = id

		// Initialisation of values
	}

	for idx, for_part := range for_partOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("For_part", idx, for_part.Name)
		map_For_part_Identifiers[for_part] = id

		// Initialisation of values
		if for_part.Part_clef != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Part_clef")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Part_clef_Identifiers[for_part.Part_clef])
			pointersInitializesStatements += setPointerField
		}

		if for_part.Part_transpose != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Part_transpose")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Part_transpose_Identifiers[for_part.Part_transpose])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, formatted_symbol := range formatted_symbolOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Formatted_symbol", idx, formatted_symbol.Name)
		map_Formatted_symbol_Identifiers[formatted_symbol] = id

		// Initialisation of values
	}

	for idx, formatted_symbol_id := range formatted_symbol_idOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Formatted_symbol_id", idx, formatted_symbol_id.Name)
		map_Formatted_symbol_id_Identifiers[formatted_symbol_id] = id

		// Initialisation of values
	}

	for idx, forward := range forwardOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Forward", idx, forward.Name)
		map_Forward_Identifiers[forward] = id

		// Initialisation of values
	}

	for idx, frame := range frameOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Frame", idx, frame.Name)
		map_Frame_Identifiers[frame] = id

		// Initialisation of values
		if frame.First_fret != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "First_fret")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_First_fret_Identifiers[frame.First_fret])
			pointersInitializesStatements += setPointerField
		}

		for _, _frame_note := range frame.Frame_note {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Frame_note")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Frame_note_Identifiers[_frame_note])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, frame_note := range frame_noteOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Frame_note", idx, frame_note.Name)
		map_Frame_note_Identifiers[frame_note] = id

		// Initialisation of values
		if frame_note.Fret != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fret")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Fret_Identifiers[frame_note.Fret])
			pointersInitializesStatements += setPointerField
		}

		if frame_note.Fingering != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fingering")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Fingering_Identifiers[frame_note.Fingering])
			pointersInitializesStatements += setPointerField
		}

		if frame_note.Barre != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Barre")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Barre_Identifiers[frame_note.Barre])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, fret := range fretOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Fret", idx, fret.Name)
		map_Fret_Identifiers[fret] = id

		// Initialisation of values
	}

	for idx, glass := range glassOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Glass", idx, glass.Name)
		map_Glass_Identifiers[glass] = id

		// Initialisation of values
	}

	for idx, glissando := range glissandoOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Glissando", idx, glissando.Name)
		map_Glissando_Identifiers[glissando] = id

		// Initialisation of values
	}

	for idx, glyph := range glyphOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Glyph", idx, glyph.Name)
		map_Glyph_Identifiers[glyph] = id

		// Initialisation of values
	}

	for idx, grace := range graceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Grace", idx, grace.Name)
		map_Grace_Identifiers[grace] = id

		// Initialisation of values
	}

	for idx, group_barline := range group_barlineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Group_barline", idx, group_barline.Name)
		map_Group_barline_Identifiers[group_barline] = id

		// Initialisation of values
	}

	for idx, group_symbol := range group_symbolOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Group_symbol", idx, group_symbol.Name)
		map_Group_symbol_Identifiers[group_symbol] = id

		// Initialisation of values
	}

	for idx, grouping := range groupingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Grouping", idx, grouping.Name)
		map_Grouping_Identifiers[grouping] = id

		// Initialisation of values
		for _, _feature := range grouping.Feature {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Feature")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Feature_Identifiers[_feature])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, hammer_on_pull_off := range hammer_on_pull_offOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Hammer_on_pull_off", idx, hammer_on_pull_off.Name)
		map_Hammer_on_pull_off_Identifiers[hammer_on_pull_off] = id

		// Initialisation of values
	}

	for idx, handbell := range handbellOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Handbell", idx, handbell.Name)
		map_Handbell_Identifiers[handbell] = id

		// Initialisation of values
	}

	for idx, harmon_closed := range harmon_closedOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Harmon_closed", idx, harmon_closed.Name)
		map_Harmon_closed_Identifiers[harmon_closed] = id

		// Initialisation of values
	}

	for idx, harmon_mute := range harmon_muteOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Harmon_mute", idx, harmon_mute.Name)
		map_Harmon_mute_Identifiers[harmon_mute] = id

		// Initialisation of values
		if harmon_mute.Harmon_closed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Harmon_closed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Harmon_closed_Identifiers[harmon_mute.Harmon_closed])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, harmonic := range harmonicOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Harmonic", idx, harmonic.Name)
		map_Harmonic_Identifiers[harmonic] = id

		// Initialisation of values
		if harmonic.Natural != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Natural")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[harmonic.Natural])
			pointersInitializesStatements += setPointerField
		}

		if harmonic.Artificial != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Artificial")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[harmonic.Artificial])
			pointersInitializesStatements += setPointerField
		}

		if harmonic.Base_pitch != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Base_pitch")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[harmonic.Base_pitch])
			pointersInitializesStatements += setPointerField
		}

		if harmonic.Touching_pitch != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Touching_pitch")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[harmonic.Touching_pitch])
			pointersInitializesStatements += setPointerField
		}

		if harmonic.Sounding_pitch != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sounding_pitch")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[harmonic.Sounding_pitch])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, harmony := range harmonyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Harmony", idx, harmony.Name)
		map_Harmony_Identifiers[harmony] = id

		// Initialisation of values
		if harmony.Frame != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Frame")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Frame_Identifiers[harmony.Frame])
			pointersInitializesStatements += setPointerField
		}

		if harmony.Offset != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Offset")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Offset_Identifiers[harmony.Offset])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, harmony_alter := range harmony_alterOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Harmony_alter", idx, harmony_alter.Name)
		map_Harmony_alter_Identifiers[harmony_alter] = id

		// Initialisation of values
	}

	for idx, harp_pedals := range harp_pedalsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Harp_pedals", idx, harp_pedals.Name)
		map_Harp_pedals_Identifiers[harp_pedals] = id

		// Initialisation of values
		for _, _pedal_tuning := range harp_pedals.Pedal_tuning {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pedal_tuning")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Pedal_tuning_Identifiers[_pedal_tuning])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, heel_toe := range heel_toeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Heel_toe", idx, heel_toe.Name)
		map_Heel_toe_Identifiers[heel_toe] = id

		// Initialisation of values
	}

	for idx, hole := range holeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Hole", idx, hole.Name)
		map_Hole_Identifiers[hole] = id

		// Initialisation of values
		if hole.Hole_closed != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Hole_closed")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Hole_closed_Identifiers[hole.Hole_closed])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, hole_closed := range hole_closedOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Hole_closed", idx, hole_closed.Name)
		map_Hole_closed_Identifiers[hole_closed] = id

		// Initialisation of values
	}

	for idx, horizontal_turn := range horizontal_turnOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Horizontal_turn", idx, horizontal_turn.Name)
		map_Horizontal_turn_Identifiers[horizontal_turn] = id

		// Initialisation of values
	}

	for idx, identification := range identificationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Identification", idx, identification.Name)
		map_Identification_Identifiers[identification] = id

		// Initialisation of values
		for _, _typed_text := range identification.Creator {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Creator")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Typed_text_Identifiers[_typed_text])
			pointersInitializesStatements += setPointerField
		}

		for _, _typed_text := range identification.Rights {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rights")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Typed_text_Identifiers[_typed_text])
			pointersInitializesStatements += setPointerField
		}

		if identification.Encoding != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Encoding")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Encoding_Identifiers[identification.Encoding])
			pointersInitializesStatements += setPointerField
		}

		for _, _typed_text := range identification.Relation {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Relation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Typed_text_Identifiers[_typed_text])
			pointersInitializesStatements += setPointerField
		}

		if identification.Miscellaneous != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Miscellaneous")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Miscellaneous_Identifiers[identification.Miscellaneous])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, image := range imageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Image", idx, image.Name)
		map_Image_Identifiers[image] = id

		// Initialisation of values
	}

	for idx, instrument := range instrumentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Instrument", idx, instrument.Name)
		map_Instrument_Identifiers[instrument] = id

		// Initialisation of values
	}

	for idx, instrument_change := range instrument_changeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Instrument_change", idx, instrument_change.Name)
		map_Instrument_change_Identifiers[instrument_change] = id

		// Initialisation of values
	}

	for idx, instrument_link := range instrument_linkOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Instrument_link", idx, instrument_link.Name)
		map_Instrument_link_Identifiers[instrument_link] = id

		// Initialisation of values
	}

	for idx, interchangeable := range interchangeableOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Interchangeable", idx, interchangeable.Name)
		map_Interchangeable_Identifiers[interchangeable] = id

		// Initialisation of values
	}

	for idx, inversion := range inversionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Inversion", idx, inversion.Name)
		map_Inversion_Identifiers[inversion] = id

		// Initialisation of values
	}

	for idx, key := range keyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Key", idx, key.Name)
		map_Key_Identifiers[key] = id

		// Initialisation of values
		for _, _key_octave := range key.Key_octave {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Key_octave")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Key_octave_Identifiers[_key_octave])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, key_accidental := range key_accidentalOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Key_accidental", idx, key_accidental.Name)
		map_Key_accidental_Identifiers[key_accidental] = id

		// Initialisation of values
	}

	for idx, key_octave := range key_octaveOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Key_octave", idx, key_octave.Name)
		map_Key_octave_Identifiers[key_octave] = id

		// Initialisation of values
	}

	for idx, kind := range kindOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Kind", idx, kind.Name)
		map_Kind_Identifiers[kind] = id

		// Initialisation of values
	}

	for idx, level := range levelOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Level", idx, level.Name)
		map_Level_Identifiers[level] = id

		// Initialisation of values
	}

	for idx, line_detail := range line_detailOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Line_detail", idx, line_detail.Name)
		map_Line_detail_Identifiers[line_detail] = id

		// Initialisation of values
	}

	for idx, line_width := range line_widthOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Line_width", idx, line_width.Name)
		map_Line_width_Identifiers[line_width] = id

		// Initialisation of values
	}

	for idx, link := range linkOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Link", idx, link.Name)
		map_Link_Identifiers[link] = id

		// Initialisation of values
	}

	for idx, listen := range listenOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Listen", idx, listen.Name)
		map_Listen_Identifiers[listen] = id

		// Initialisation of values
		if listen.Assess != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Assess")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Assess_Identifiers[listen.Assess])
			pointersInitializesStatements += setPointerField
		}

		if listen.Wait != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Wait")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Wait_Identifiers[listen.Wait])
			pointersInitializesStatements += setPointerField
		}

		if listen.Other_listen != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Other_listen")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Other_listening_Identifiers[listen.Other_listen])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, listening := range listeningOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Listening", idx, listening.Name)
		map_Listening_Identifiers[listening] = id

		// Initialisation of values
		if listening.Offset != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Offset")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Offset_Identifiers[listening.Offset])
			pointersInitializesStatements += setPointerField
		}

		if listening.Sync != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sync")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sync_Identifiers[listening.Sync])
			pointersInitializesStatements += setPointerField
		}

		if listening.Other_listening != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Other_listening")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Other_listening_Identifiers[listening.Other_listening])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, lyric := range lyricOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Lyric", idx, lyric.Name)
		map_Lyric_Identifiers[lyric] = id

		// Initialisation of values
		if lyric.End_line != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "End_line")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[lyric.End_line])
			pointersInitializesStatements += setPointerField
		}

		if lyric.End_paragraph != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "End_paragraph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[lyric.End_paragraph])
			pointersInitializesStatements += setPointerField
		}

		if lyric.Extend != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Extend")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Extend_Identifiers[lyric.Extend])
			pointersInitializesStatements += setPointerField
		}

		if lyric.Laughing != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Laughing")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[lyric.Laughing])
			pointersInitializesStatements += setPointerField
		}

		if lyric.Humming != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Humming")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[lyric.Humming])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, lyric_font := range lyric_fontOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Lyric_font", idx, lyric_font.Name)
		map_Lyric_font_Identifiers[lyric_font] = id

		// Initialisation of values
	}

	for idx, lyric_language := range lyric_languageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Lyric_language", idx, lyric_language.Name)
		map_Lyric_language_Identifiers[lyric_language] = id

		// Initialisation of values
	}

	for idx, measure_layout := range measure_layoutOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Measure_layout", idx, measure_layout.Name)
		map_Measure_layout_Identifiers[measure_layout] = id

		// Initialisation of values
	}

	for idx, measure_numbering := range measure_numberingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Measure_numbering", idx, measure_numbering.Name)
		map_Measure_numbering_Identifiers[measure_numbering] = id

		// Initialisation of values
	}

	for idx, measure_repeat := range measure_repeatOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Measure_repeat", idx, measure_repeat.Name)
		map_Measure_repeat_Identifiers[measure_repeat] = id

		// Initialisation of values
	}

	for idx, measure_style := range measure_styleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Measure_style", idx, measure_style.Name)
		map_Measure_style_Identifiers[measure_style] = id

		// Initialisation of values
		if measure_style.Multiple_rest != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Multiple_rest")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Multiple_rest_Identifiers[measure_style.Multiple_rest])
			pointersInitializesStatements += setPointerField
		}

		if measure_style.Measure_repeat != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Measure_repeat")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Measure_repeat_Identifiers[measure_style.Measure_repeat])
			pointersInitializesStatements += setPointerField
		}

		if measure_style.Beat_repeat != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Beat_repeat")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Beat_repeat_Identifiers[measure_style.Beat_repeat])
			pointersInitializesStatements += setPointerField
		}

		if measure_style.Slash != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Slash")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Slash_Identifiers[measure_style.Slash])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, membrane := range membraneOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Membrane", idx, membrane.Name)
		map_Membrane_Identifiers[membrane] = id

		// Initialisation of values
	}

	for idx, metal := range metalOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Metal", idx, metal.Name)
		map_Metal_Identifiers[metal] = id

		// Initialisation of values
	}

	for idx, metronome := range metronomeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Metronome", idx, metronome.Name)
		map_Metronome_Identifiers[metronome] = id

		// Initialisation of values
	}

	for idx, metronome_beam := range metronome_beamOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Metronome_beam", idx, metronome_beam.Name)
		map_Metronome_beam_Identifiers[metronome_beam] = id

		// Initialisation of values
	}

	for idx, metronome_note := range metronome_noteOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Metronome_note", idx, metronome_note.Name)
		map_Metronome_note_Identifiers[metronome_note] = id

		// Initialisation of values
		for _, _empty := range metronome_note.Metronome_dot {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Metronome_dot")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[_empty])
			pointersInitializesStatements += setPointerField
		}

		for _, _metronome_beam := range metronome_note.Metronome_beam {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Metronome_beam")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Metronome_beam_Identifiers[_metronome_beam])
			pointersInitializesStatements += setPointerField
		}

		if metronome_note.Metronome_tied != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Metronome_tied")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Metronome_tied_Identifiers[metronome_note.Metronome_tied])
			pointersInitializesStatements += setPointerField
		}

		if metronome_note.Metronome_tuplet != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Metronome_tuplet")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Metronome_tuplet_Identifiers[metronome_note.Metronome_tuplet])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, metronome_tied := range metronome_tiedOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Metronome_tied", idx, metronome_tied.Name)
		map_Metronome_tied_Identifiers[metronome_tied] = id

		// Initialisation of values
	}

	for idx, metronome_tuplet := range metronome_tupletOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Metronome_tuplet", idx, metronome_tuplet.Name)
		map_Metronome_tuplet_Identifiers[metronome_tuplet] = id

		// Initialisation of values
	}

	for idx, midi_device := range midi_deviceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Midi_device", idx, midi_device.Name)
		map_Midi_device_Identifiers[midi_device] = id

		// Initialisation of values
	}

	for idx, midi_instrument := range midi_instrumentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Midi_instrument", idx, midi_instrument.Name)
		map_Midi_instrument_Identifiers[midi_instrument] = id

		// Initialisation of values
	}

	for idx, miscellaneous := range miscellaneousOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Miscellaneous", idx, miscellaneous.Name)
		map_Miscellaneous_Identifiers[miscellaneous] = id

		// Initialisation of values
		for _, _miscellaneous_field := range miscellaneous.Miscellaneous_field {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Miscellaneous_field")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Miscellaneous_field_Identifiers[_miscellaneous_field])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, miscellaneous_field := range miscellaneous_fieldOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Miscellaneous_field", idx, miscellaneous_field.Name)
		map_Miscellaneous_field_Identifiers[miscellaneous_field] = id

		// Initialisation of values
	}

	for idx, mordent := range mordentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Mordent", idx, mordent.Name)
		map_Mordent_Identifiers[mordent] = id

		// Initialisation of values
	}

	for idx, multiple_rest := range multiple_restOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Multiple_rest", idx, multiple_rest.Name)
		map_Multiple_rest_Identifiers[multiple_rest] = id

		// Initialisation of values
	}

	for idx, name_display := range name_displayOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Name_display", idx, name_display.Name)
		map_Name_display_Identifiers[name_display] = id

		// Initialisation of values
		if name_display.Accidental_text != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accidental_text")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Accidental_text_Identifiers[name_display.Accidental_text])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, non_arpeggiate := range non_arpeggiateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Non_arpeggiate", idx, non_arpeggiate.Name)
		map_Non_arpeggiate_Identifiers[non_arpeggiate] = id

		// Initialisation of values
	}

	for idx, notations := range notationsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Notations", idx, notations.Name)
		map_Notations_Identifiers[notations] = id

		// Initialisation of values
		if notations.Tied != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tied")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tied_Identifiers[notations.Tied])
			pointersInitializesStatements += setPointerField
		}

		if notations.Slur != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Slur")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Slur_Identifiers[notations.Slur])
			pointersInitializesStatements += setPointerField
		}

		if notations.Tuplet != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tuplet")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tuplet_Identifiers[notations.Tuplet])
			pointersInitializesStatements += setPointerField
		}

		if notations.Glissando != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Glissando")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Glissando_Identifiers[notations.Glissando])
			pointersInitializesStatements += setPointerField
		}

		if notations.Slide != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Slide")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Slide_Identifiers[notations.Slide])
			pointersInitializesStatements += setPointerField
		}

		if notations.Ornaments != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Ornaments")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Ornaments_Identifiers[notations.Ornaments])
			pointersInitializesStatements += setPointerField
		}

		if notations.Technical != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Technical")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Technical_Identifiers[notations.Technical])
			pointersInitializesStatements += setPointerField
		}

		if notations.Articulations != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Articulations")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Articulations_Identifiers[notations.Articulations])
			pointersInitializesStatements += setPointerField
		}

		if notations.Dynamics != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dynamics")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Dynamics_Identifiers[notations.Dynamics])
			pointersInitializesStatements += setPointerField
		}

		if notations.Fermata != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fermata")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Fermata_Identifiers[notations.Fermata])
			pointersInitializesStatements += setPointerField
		}

		if notations.Arpeggiate != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Arpeggiate")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Arpeggiate_Identifiers[notations.Arpeggiate])
			pointersInitializesStatements += setPointerField
		}

		if notations.Non_arpeggiate != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Non_arpeggiate")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Non_arpeggiate_Identifiers[notations.Non_arpeggiate])
			pointersInitializesStatements += setPointerField
		}

		if notations.Accidental_mark != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accidental_mark")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Accidental_mark_Identifiers[notations.Accidental_mark])
			pointersInitializesStatements += setPointerField
		}

		if notations.Other_notation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Other_notation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Other_notation_Identifiers[notations.Other_notation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, note := range noteOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Note", idx, note.Name)
		map_Note_Identifiers[note] = id

		// Initialisation of values
		for _, _instrument := range note.Instrument {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Instrument")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Instrument_Identifiers[_instrument])
			pointersInitializesStatements += setPointerField
		}

		if note.Type_ != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Type_")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Note_type_Identifiers[note.Type_])
			pointersInitializesStatements += setPointerField
		}

		for _, _empty_placement := range note.Dot {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dot")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[_empty_placement])
			pointersInitializesStatements += setPointerField
		}

		if note.Accidental != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accidental")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Accidental_Identifiers[note.Accidental])
			pointersInitializesStatements += setPointerField
		}

		if note.Time_modification != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Time_modification")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Time_modification_Identifiers[note.Time_modification])
			pointersInitializesStatements += setPointerField
		}

		if note.Stem != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Stem")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Stem_Identifiers[note.Stem])
			pointersInitializesStatements += setPointerField
		}

		if note.Notehead != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Notehead")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Notehead_Identifiers[note.Notehead])
			pointersInitializesStatements += setPointerField
		}

		if note.Notehead_text != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Notehead_text")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Notehead_text_Identifiers[note.Notehead_text])
			pointersInitializesStatements += setPointerField
		}

		if note.Beam != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Beam")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Beam_Identifiers[note.Beam])
			pointersInitializesStatements += setPointerField
		}

		for _, _notations := range note.Notations {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Notations")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Notations_Identifiers[_notations])
			pointersInitializesStatements += setPointerField
		}

		for _, _lyric := range note.Lyric {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lyric")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lyric_Identifiers[_lyric])
			pointersInitializesStatements += setPointerField
		}

		if note.Play != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Play")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Play_Identifiers[note.Play])
			pointersInitializesStatements += setPointerField
		}

		if note.Listen != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Listen")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Listen_Identifiers[note.Listen])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, note_size := range note_sizeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Note_size", idx, note_size.Name)
		map_Note_size_Identifiers[note_size] = id

		// Initialisation of values
	}

	for idx, note_type := range note_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Note_type", idx, note_type.Name)
		map_Note_type_Identifiers[note_type] = id

		// Initialisation of values
	}

	for idx, notehead := range noteheadOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Notehead", idx, notehead.Name)
		map_Notehead_Identifiers[notehead] = id

		// Initialisation of values
	}

	for idx, notehead_text := range notehead_textOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Notehead_text", idx, notehead_text.Name)
		map_Notehead_text_Identifiers[notehead_text] = id

		// Initialisation of values
		if notehead_text.Accidental_text != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accidental_text")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Accidental_text_Identifiers[notehead_text.Accidental_text])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, numeral := range numeralOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Numeral", idx, numeral.Name)
		map_Numeral_Identifiers[numeral] = id

		// Initialisation of values
		if numeral.Numeral_root != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Numeral_root")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Numeral_root_Identifiers[numeral.Numeral_root])
			pointersInitializesStatements += setPointerField
		}

		if numeral.Numeral_alter != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Numeral_alter")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Harmony_alter_Identifiers[numeral.Numeral_alter])
			pointersInitializesStatements += setPointerField
		}

		if numeral.Numeral_key != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Numeral_key")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Numeral_key_Identifiers[numeral.Numeral_key])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, numeral_key := range numeral_keyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Numeral_key", idx, numeral_key.Name)
		map_Numeral_key_Identifiers[numeral_key] = id

		// Initialisation of values
	}

	for idx, numeral_root := range numeral_rootOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Numeral_root", idx, numeral_root.Name)
		map_Numeral_root_Identifiers[numeral_root] = id

		// Initialisation of values
	}

	for idx, octave_shift := range octave_shiftOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Octave_shift", idx, octave_shift.Name)
		map_Octave_shift_Identifiers[octave_shift] = id

		// Initialisation of values
	}

	for idx, offset := range offsetOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Offset", idx, offset.Name)
		map_Offset_Identifiers[offset] = id

		// Initialisation of values
	}

	for idx, opus := range opusOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Opus", idx, opus.Name)
		map_Opus_Identifiers[opus] = id

		// Initialisation of values
	}

	for idx, ornaments := range ornamentsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Ornaments", idx, ornaments.Name)
		map_Ornaments_Identifiers[ornaments] = id

		// Initialisation of values
		for _, _accidental_mark := range ornaments.Accidental_mark {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accidental_mark")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Accidental_mark_Identifiers[_accidental_mark])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Trill_mark != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Trill_mark")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_trill_sound_Identifiers[ornaments.Trill_mark])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Turn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Turn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Horizontal_turn_Identifiers[ornaments.Turn])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Delayed_turn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Delayed_turn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Horizontal_turn_Identifiers[ornaments.Delayed_turn])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Inverted_turn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Inverted_turn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Horizontal_turn_Identifiers[ornaments.Inverted_turn])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Delayed_inverted_turn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Delayed_inverted_turn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Horizontal_turn_Identifiers[ornaments.Delayed_inverted_turn])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Vertical_turn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Vertical_turn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_trill_sound_Identifiers[ornaments.Vertical_turn])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Inverted_vertical_turn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Inverted_vertical_turn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_trill_sound_Identifiers[ornaments.Inverted_vertical_turn])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Shake != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Shake")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_trill_sound_Identifiers[ornaments.Shake])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Wavy_line != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Wavy_line")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Wavy_line_Identifiers[ornaments.Wavy_line])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Mordent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Mordent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Mordent_Identifiers[ornaments.Mordent])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Inverted_mordent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Inverted_mordent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Mordent_Identifiers[ornaments.Inverted_mordent])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Schleifer != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Schleifer")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[ornaments.Schleifer])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Tremolo != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tremolo")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tremolo_Identifiers[ornaments.Tremolo])
			pointersInitializesStatements += setPointerField
		}

		if ornaments.Haydn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Haydn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_trill_sound_Identifiers[ornaments.Haydn])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, other_appearance := range other_appearanceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Other_appearance", idx, other_appearance.Name)
		map_Other_appearance_Identifiers[other_appearance] = id

		// Initialisation of values
	}

	for idx, other_listening := range other_listeningOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Other_listening", idx, other_listening.Name)
		map_Other_listening_Identifiers[other_listening] = id

		// Initialisation of values
	}

	for idx, other_notation := range other_notationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Other_notation", idx, other_notation.Name)
		map_Other_notation_Identifiers[other_notation] = id

		// Initialisation of values
	}

	for idx, other_play := range other_playOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Other_play", idx, other_play.Name)
		map_Other_play_Identifiers[other_play] = id

		// Initialisation of values
	}

	for idx, page_layout := range page_layoutOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Page_layout", idx, page_layout.Name)
		map_Page_layout_Identifiers[page_layout] = id

		// Initialisation of values
		if page_layout.Page_margins != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Page_margins")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Page_margins_Identifiers[page_layout.Page_margins])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, page_margins := range page_marginsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Page_margins", idx, page_margins.Name)
		map_Page_margins_Identifiers[page_margins] = id

		// Initialisation of values
	}

	for idx, part_clef := range part_clefOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Part_clef", idx, part_clef.Name)
		map_Part_clef_Identifiers[part_clef] = id

		// Initialisation of values
	}

	for idx, part_group := range part_groupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Part_group", idx, part_group.Name)
		map_Part_group_Identifiers[part_group] = id

		// Initialisation of values
		if part_group.Group_name_display != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Group_name_display")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Name_display_Identifiers[part_group.Group_name_display])
			pointersInitializesStatements += setPointerField
		}

		if part_group.Group_abbreviation_display != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Group_abbreviation_display")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Name_display_Identifiers[part_group.Group_abbreviation_display])
			pointersInitializesStatements += setPointerField
		}

		if part_group.Group_symbol != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Group_symbol")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_symbol_Identifiers[part_group.Group_symbol])
			pointersInitializesStatements += setPointerField
		}

		if part_group.Group_barline != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Group_barline")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_barline_Identifiers[part_group.Group_barline])
			pointersInitializesStatements += setPointerField
		}

		if part_group.Group_time != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Group_time")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[part_group.Group_time])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, part_link := range part_linkOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Part_link", idx, part_link.Name)
		map_Part_link_Identifiers[part_link] = id

		// Initialisation of values
		for _, _instrument_link := range part_link.Instrument_link {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Instrument_link")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Instrument_link_Identifiers[_instrument_link])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, part_list := range part_listOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Part_list", idx, part_list.Name)
		map_Part_list_Identifiers[part_list] = id

		// Initialisation of values
	}

	for idx, part_symbol := range part_symbolOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Part_symbol", idx, part_symbol.Name)
		map_Part_symbol_Identifiers[part_symbol] = id

		// Initialisation of values
	}

	for idx, part_transpose := range part_transposeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Part_transpose", idx, part_transpose.Name)
		map_Part_transpose_Identifiers[part_transpose] = id

		// Initialisation of values
	}

	for idx, pedal := range pedalOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Pedal", idx, pedal.Name)
		map_Pedal_Identifiers[pedal] = id

		// Initialisation of values
	}

	for idx, pedal_tuning := range pedal_tuningOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Pedal_tuning", idx, pedal_tuning.Name)
		map_Pedal_tuning_Identifiers[pedal_tuning] = id

		// Initialisation of values
	}

	for idx, percussion := range percussionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Percussion", idx, percussion.Name)
		map_Percussion_Identifiers[percussion] = id

		// Initialisation of values
		if percussion.Glass != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Glass")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Glass_Identifiers[percussion.Glass])
			pointersInitializesStatements += setPointerField
		}

		if percussion.Metal != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Metal")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Metal_Identifiers[percussion.Metal])
			pointersInitializesStatements += setPointerField
		}

		if percussion.Wood != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Wood")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Wood_Identifiers[percussion.Wood])
			pointersInitializesStatements += setPointerField
		}

		if percussion.Pitched != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pitched")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Pitched_Identifiers[percussion.Pitched])
			pointersInitializesStatements += setPointerField
		}

		if percussion.Membrane != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Membrane")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Membrane_Identifiers[percussion.Membrane])
			pointersInitializesStatements += setPointerField
		}

		if percussion.Effect != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Effect")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Effect_Identifiers[percussion.Effect])
			pointersInitializesStatements += setPointerField
		}

		if percussion.Timpani != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Timpani")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Timpani_Identifiers[percussion.Timpani])
			pointersInitializesStatements += setPointerField
		}

		if percussion.Beater != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Beater")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Beater_Identifiers[percussion.Beater])
			pointersInitializesStatements += setPointerField
		}

		if percussion.Stick != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Stick")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Stick_Identifiers[percussion.Stick])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, pitch := range pitchOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Pitch", idx, pitch.Name)
		map_Pitch_Identifiers[pitch] = id

		// Initialisation of values
	}

	for idx, pitched := range pitchedOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Pitched", idx, pitched.Name)
		map_Pitched_Identifiers[pitched] = id

		// Initialisation of values
	}

	for idx, play := range playOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Play", idx, play.Name)
		map_Play_Identifiers[play] = id

		// Initialisation of values
		if play.Other_play != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Other_play")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Other_play_Identifiers[play.Other_play])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, player := range playerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Player", idx, player.Name)
		map_Player_Identifiers[player] = id

		// Initialisation of values
	}

	for idx, principal_voice := range principal_voiceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Principal_voice", idx, principal_voice.Name)
		map_Principal_voice_Identifiers[principal_voice] = id

		// Initialisation of values
	}

	for idx, print := range printOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Print", idx, print.Name)
		map_Print_Identifiers[print] = id

		// Initialisation of values
		if print.Measure_layout != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Measure_layout")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Measure_layout_Identifiers[print.Measure_layout])
			pointersInitializesStatements += setPointerField
		}

		if print.Measure_numbering != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Measure_numbering")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Measure_numbering_Identifiers[print.Measure_numbering])
			pointersInitializesStatements += setPointerField
		}

		if print.Part_name_display != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Part_name_display")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Name_display_Identifiers[print.Part_name_display])
			pointersInitializesStatements += setPointerField
		}

		if print.Part_abbreviation_display != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Part_abbreviation_display")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Name_display_Identifiers[print.Part_abbreviation_display])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, release := range releaseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Release", idx, release.Name)
		map_Release_Identifiers[release] = id

		// Initialisation of values
	}

	for idx, repeat := range repeatOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Repeat", idx, repeat.Name)
		map_Repeat_Identifiers[repeat] = id

		// Initialisation of values
	}

	for idx, rest := range restOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Rest", idx, rest.Name)
		map_Rest_Identifiers[rest] = id

		// Initialisation of values
	}

	for idx, root := range rootOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Root", idx, root.Name)
		map_Root_Identifiers[root] = id

		// Initialisation of values
		if root.Root_step != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Root_step")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Root_step_Identifiers[root.Root_step])
			pointersInitializesStatements += setPointerField
		}

		if root.Root_alter != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Root_alter")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Harmony_alter_Identifiers[root.Root_alter])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, root_step := range root_stepOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Root_step", idx, root_step.Name)
		map_Root_step_Identifiers[root_step] = id

		// Initialisation of values
	}

	for idx, scaling := range scalingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Scaling", idx, scaling.Name)
		map_Scaling_Identifiers[scaling] = id

		// Initialisation of values
	}

	for idx, scordatura := range scordaturaOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Scordatura", idx, scordatura.Name)
		map_Scordatura_Identifiers[scordatura] = id

		// Initialisation of values
		for _, _accord := range scordatura.Accord {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Accord")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Accord_Identifiers[_accord])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, score_instrument := range score_instrumentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Score_instrument", idx, score_instrument.Name)
		map_Score_instrument_Identifiers[score_instrument] = id

		// Initialisation of values
	}

	for idx, score_part := range score_partOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Score_part", idx, score_part.Name)
		map_Score_part_Identifiers[score_part] = id

		// Initialisation of values
		if score_part.Identification != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Identification")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Identification_Identifiers[score_part.Identification])
			pointersInitializesStatements += setPointerField
		}

		for _, _part_link := range score_part.Part_link {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Part_link")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Part_link_Identifiers[_part_link])
			pointersInitializesStatements += setPointerField
		}

		if score_part.Part_name_display != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Part_name_display")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Name_display_Identifiers[score_part.Part_name_display])
			pointersInitializesStatements += setPointerField
		}

		if score_part.Part_abbreviation_display != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Part_abbreviation_display")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Name_display_Identifiers[score_part.Part_abbreviation_display])
			pointersInitializesStatements += setPointerField
		}

		for _, _score_instrument := range score_part.Score_instrument {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Score_instrument")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Score_instrument_Identifiers[_score_instrument])
			pointersInitializesStatements += setPointerField
		}

		for _, _player := range score_part.Player {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Player")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Player_Identifiers[_player])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, score_partwise := range score_partwiseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Score_partwise", idx, score_partwise.Name)
		map_Score_partwise_Identifiers[score_partwise] = id

		// Initialisation of values
	}

	for idx, score_timewise := range score_timewiseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Score_timewise", idx, score_timewise.Name)
		map_Score_timewise_Identifiers[score_timewise] = id

		// Initialisation of values
	}

	for idx, segno := range segnoOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Segno", idx, segno.Name)
		map_Segno_Identifiers[segno] = id

		// Initialisation of values
	}

	for idx, slash := range slashOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Slash", idx, slash.Name)
		map_Slash_Identifiers[slash] = id

		// Initialisation of values
	}

	for idx, slide := range slideOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Slide", idx, slide.Name)
		map_Slide_Identifiers[slide] = id

		// Initialisation of values
	}

	for idx, slur := range slurOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Slur", idx, slur.Name)
		map_Slur_Identifiers[slur] = id

		// Initialisation of values
	}

	for idx, sound := range soundOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Sound", idx, sound.Name)
		map_Sound_Identifiers[sound] = id

		// Initialisation of values
		if sound.Swing != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Swing")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Swing_Identifiers[sound.Swing])
			pointersInitializesStatements += setPointerField
		}

		if sound.Offset != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Offset")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Offset_Identifiers[sound.Offset])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, staff_details := range staff_detailsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Staff_details", idx, staff_details.Name)
		map_Staff_details_Identifiers[staff_details] = id

		// Initialisation of values
		for _, _staff_tuning := range staff_details.Staff_tuning {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Staff_tuning")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Staff_tuning_Identifiers[_staff_tuning])
			pointersInitializesStatements += setPointerField
		}

		if staff_details.Staff_size != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Staff_size")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Staff_size_Identifiers[staff_details.Staff_size])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, staff_divide := range staff_divideOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Staff_divide", idx, staff_divide.Name)
		map_Staff_divide_Identifiers[staff_divide] = id

		// Initialisation of values
	}

	for idx, staff_layout := range staff_layoutOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Staff_layout", idx, staff_layout.Name)
		map_Staff_layout_Identifiers[staff_layout] = id

		// Initialisation of values
	}

	for idx, staff_size := range staff_sizeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Staff_size", idx, staff_size.Name)
		map_Staff_size_Identifiers[staff_size] = id

		// Initialisation of values
	}

	for idx, staff_tuning := range staff_tuningOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Staff_tuning", idx, staff_tuning.Name)
		map_Staff_tuning_Identifiers[staff_tuning] = id

		// Initialisation of values
	}

	for idx, stem := range stemOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Stem", idx, stem.Name)
		map_Stem_Identifiers[stem] = id

		// Initialisation of values
	}

	for idx, stick := range stickOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Stick", idx, stick.Name)
		map_Stick_Identifiers[stick] = id

		// Initialisation of values
	}

	for idx, string_mute := range string_muteOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("String_mute", idx, string_mute.Name)
		map_String_mute_Identifiers[string_mute] = id

		// Initialisation of values
	}

	for idx, strong_accent := range strong_accentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Strong_accent", idx, strong_accent.Name)
		map_Strong_accent_Identifiers[strong_accent] = id

		// Initialisation of values
	}

	for idx, supports := range supportsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Supports", idx, supports.Name)
		map_Supports_Identifiers[supports] = id

		// Initialisation of values
	}

	for idx, swing := range swingOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Swing", idx, swing.Name)
		map_Swing_Identifiers[swing] = id

		// Initialisation of values
		if swing.Straight != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Straight")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_Identifiers[swing.Straight])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, sync := range syncOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Sync", idx, sync.Name)
		map_Sync_Identifiers[sync] = id

		// Initialisation of values
	}

	for idx, system_dividers := range system_dividersOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("System_dividers", idx, system_dividers.Name)
		map_System_dividers_Identifiers[system_dividers] = id

		// Initialisation of values
		if system_dividers.Left_divider != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Left_divider")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_print_object_style_align_Identifiers[system_dividers.Left_divider])
			pointersInitializesStatements += setPointerField
		}

		if system_dividers.Right_divider != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Right_divider")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_print_object_style_align_Identifiers[system_dividers.Right_divider])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, system_layout := range system_layoutOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("System_layout", idx, system_layout.Name)
		map_System_layout_Identifiers[system_layout] = id

		// Initialisation of values
		if system_layout.System_margins != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "System_margins")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_System_margins_Identifiers[system_layout.System_margins])
			pointersInitializesStatements += setPointerField
		}

		if system_layout.System_dividers != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "System_dividers")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_System_dividers_Identifiers[system_layout.System_dividers])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, system_margins := range system_marginsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("System_margins", idx, system_margins.Name)
		map_System_margins_Identifiers[system_margins] = id

		// Initialisation of values
	}

	for idx, tap := range tapOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tap", idx, tap.Name)
		map_Tap_Identifiers[tap] = id

		// Initialisation of values
	}

	for idx, technical := range technicalOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Technical", idx, technical.Name)
		map_Technical_Identifiers[technical] = id

		// Initialisation of values
		if technical.Up_bow != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Up_bow")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Up_bow])
			pointersInitializesStatements += setPointerField
		}

		if technical.Down_bow != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Down_bow")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Down_bow])
			pointersInitializesStatements += setPointerField
		}

		if technical.Harmonic != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Harmonic")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Harmonic_Identifiers[technical.Harmonic])
			pointersInitializesStatements += setPointerField
		}

		if technical.Open_string != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Open_string")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Open_string])
			pointersInitializesStatements += setPointerField
		}

		if technical.Thumb_position != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Thumb_position")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Thumb_position])
			pointersInitializesStatements += setPointerField
		}

		if technical.Fingering != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fingering")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Fingering_Identifiers[technical.Fingering])
			pointersInitializesStatements += setPointerField
		}

		if technical.Double_tongue != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Double_tongue")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Double_tongue])
			pointersInitializesStatements += setPointerField
		}

		if technical.Triple_tongue != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Triple_tongue")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Triple_tongue])
			pointersInitializesStatements += setPointerField
		}

		if technical.Stopped != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Stopped")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_smufl_Identifiers[technical.Stopped])
			pointersInitializesStatements += setPointerField
		}

		if technical.Snap_pizzicato != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Snap_pizzicato")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Snap_pizzicato])
			pointersInitializesStatements += setPointerField
		}

		if technical.Fret != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fret")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Fret_Identifiers[technical.Fret])
			pointersInitializesStatements += setPointerField
		}

		if technical.Hammer_on != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Hammer_on")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Hammer_on_pull_off_Identifiers[technical.Hammer_on])
			pointersInitializesStatements += setPointerField
		}

		if technical.Pull_off != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pull_off")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Hammer_on_pull_off_Identifiers[technical.Pull_off])
			pointersInitializesStatements += setPointerField
		}

		if technical.Bend != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bend")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bend_Identifiers[technical.Bend])
			pointersInitializesStatements += setPointerField
		}

		if technical.Tap != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tap")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tap_Identifiers[technical.Tap])
			pointersInitializesStatements += setPointerField
		}

		if technical.Heel != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Heel")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Heel_toe_Identifiers[technical.Heel])
			pointersInitializesStatements += setPointerField
		}

		if technical.Toe != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Toe")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Heel_toe_Identifiers[technical.Toe])
			pointersInitializesStatements += setPointerField
		}

		if technical.Fingernails != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fingernails")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Fingernails])
			pointersInitializesStatements += setPointerField
		}

		if technical.Hole != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Hole")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Hole_Identifiers[technical.Hole])
			pointersInitializesStatements += setPointerField
		}

		if technical.Arrow != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Arrow")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Arrow_Identifiers[technical.Arrow])
			pointersInitializesStatements += setPointerField
		}

		if technical.Handbell != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Handbell")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Handbell_Identifiers[technical.Handbell])
			pointersInitializesStatements += setPointerField
		}

		if technical.Brass_bend != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Brass_bend")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Brass_bend])
			pointersInitializesStatements += setPointerField
		}

		if technical.Flip != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Flip")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Flip])
			pointersInitializesStatements += setPointerField
		}

		if technical.Smear != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Smear")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Smear])
			pointersInitializesStatements += setPointerField
		}

		if technical.Open != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Open")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_smufl_Identifiers[technical.Open])
			pointersInitializesStatements += setPointerField
		}

		if technical.Half_muted != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Half_muted")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_smufl_Identifiers[technical.Half_muted])
			pointersInitializesStatements += setPointerField
		}

		if technical.Harmon_mute != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Harmon_mute")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Harmon_mute_Identifiers[technical.Harmon_mute])
			pointersInitializesStatements += setPointerField
		}

		if technical.Golpe != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Golpe")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Empty_placement_Identifiers[technical.Golpe])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, text_element_data := range text_element_dataOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Text_element_data", idx, text_element_data.Name)
		map_Text_element_data_Identifiers[text_element_data] = id

		// Initialisation of values
	}

	for idx, tie := range tieOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tie", idx, tie.Name)
		map_Tie_Identifiers[tie] = id

		// Initialisation of values
	}

	for idx, tied := range tiedOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tied", idx, tied.Name)
		map_Tied_Identifiers[tied] = id

		// Initialisation of values
	}

	for idx, time := range timeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Time", idx, time.Name)
		map_Time_Identifiers[time] = id

		// Initialisation of values
	}

	for idx, time_modification := range time_modificationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Time_modification", idx, time_modification.Name)
		map_Time_modification_Identifiers[time_modification] = id

		// Initialisation of values
	}

	for idx, timpani := range timpaniOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Timpani", idx, timpani.Name)
		map_Timpani_Identifiers[timpani] = id

		// Initialisation of values
	}

	for idx, transpose := range transposeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Transpose", idx, transpose.Name)
		map_Transpose_Identifiers[transpose] = id

		// Initialisation of values
	}

	for idx, tremolo := range tremoloOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tremolo", idx, tremolo.Name)
		map_Tremolo_Identifiers[tremolo] = id

		// Initialisation of values
	}

	for idx, tuplet := range tupletOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tuplet", idx, tuplet.Name)
		map_Tuplet_Identifiers[tuplet] = id

		// Initialisation of values
		if tuplet.Tuplet_actual != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tuplet_actual")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tuplet_portion_Identifiers[tuplet.Tuplet_actual])
			pointersInitializesStatements += setPointerField
		}

		if tuplet.Tuplet_normal != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tuplet_normal")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tuplet_portion_Identifiers[tuplet.Tuplet_normal])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, tuplet_dot := range tuplet_dotOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tuplet_dot", idx, tuplet_dot.Name)
		map_Tuplet_dot_Identifiers[tuplet_dot] = id

		// Initialisation of values
	}

	for idx, tuplet_number := range tuplet_numberOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tuplet_number", idx, tuplet_number.Name)
		map_Tuplet_number_Identifiers[tuplet_number] = id

		// Initialisation of values
	}

	for idx, tuplet_portion := range tuplet_portionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tuplet_portion", idx, tuplet_portion.Name)
		map_Tuplet_portion_Identifiers[tuplet_portion] = id

		// Initialisation of values
		if tuplet_portion.Tuplet_number != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tuplet_number")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tuplet_number_Identifiers[tuplet_portion.Tuplet_number])
			pointersInitializesStatements += setPointerField
		}

		if tuplet_portion.Tuplet_type != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tuplet_type")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tuplet_type_Identifiers[tuplet_portion.Tuplet_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _tuplet_dot := range tuplet_portion.Tuplet_dot {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tuplet_dot")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tuplet_dot_Identifiers[_tuplet_dot])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, tuplet_type := range tuplet_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tuplet_type", idx, tuplet_type.Name)
		map_Tuplet_type_Identifiers[tuplet_type] = id

		// Initialisation of values
	}

	for idx, typed_text := range typed_textOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Typed_text", idx, typed_text.Name)
		map_Typed_text_Identifiers[typed_text] = id

		// Initialisation of values
	}

	for idx, unpitched := range unpitchedOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Unpitched", idx, unpitched.Name)
		map_Unpitched_Identifiers[unpitched] = id

		// Initialisation of values
	}

	for idx, virtual_instrument := range virtual_instrumentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Virtual_instrument", idx, virtual_instrument.Name)
		map_Virtual_instrument_Identifiers[virtual_instrument] = id

		// Initialisation of values
	}

	for idx, wait := range waitOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Wait", idx, wait.Name)
		map_Wait_Identifiers[wait] = id

		// Initialisation of values
	}

	for idx, wavy_line := range wavy_lineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Wavy_line", idx, wavy_line.Name)
		map_Wavy_line_Identifiers[wavy_line] = id

		// Initialisation of values
	}

	for idx, wedge := range wedgeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Wedge", idx, wedge.Name)
		map_Wedge_Identifiers[wedge] = id

		// Initialisation of values
	}

	for idx, wood := range woodOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Wood", idx, wood.Name)
		map_Wood_Identifiers[wood] = id

		// Initialisation of values
	}

	for idx, work := range workOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Work", idx, work.Name)
		map_Work_Identifiers[work] = id

		// Initialisation of values
		if work.Opus != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Opus")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Opus_Identifiers[work.Opus])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
