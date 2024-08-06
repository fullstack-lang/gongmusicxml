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

	// insertion initialization of objects to stage
	for idx, bookmark := range bookmarkOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bookmark", idx, bookmark.Name)
		map_Bookmark_Identifiers[bookmark] = id

		// Initialisation of values
	}

	for idx, foo := range fooOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Foo", idx, foo.Name)
		map_Foo_Identifiers[foo] = id

		// Initialisation of values
	}

	for idx, link := range linkOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Link", idx, link.Name)
		map_Link_Identifiers[link] = id

		// Initialisation of values
	}

	for idx, lyric := range lyricOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Lyric", idx, lyric.Name)
		map_Lyric_Identifiers[lyric] = id

		// Initialisation of values
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

	for idx, miscellaneous_field := range miscellaneous_fieldOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Miscellaneous_field", idx, miscellaneous_field.Name)
		map_Miscellaneous_field_Identifiers[miscellaneous_field] = id

		// Initialisation of values
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
