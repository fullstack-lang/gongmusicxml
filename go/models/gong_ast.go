// generated code - do not edit
package models

import (
	"bufio"
	"errors"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError
var dummy_time_import time.Time

// swagger:ignore
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(stage *StageStruct, pathToFile string) error {

	ReplaceOldDeclarationsInFile(pathToFile)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *StageStruct, inFile *ast.File, fset *token.FileSet) error {
	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", inFile.Name)
	}
	if len(inFile.Imports) == 3 {
		stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "_")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						// Create an ast.CommentMap from the ast.File's comments.
						// This helps keeping the association between comments
						// and AST nodes.
						cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								stage, &cmap, assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				case *ast.ValueSpec:
					ident := spec.Names[0]
					_ = ident
					if !strings.HasPrefix(ident.Name, "_") {
						continue
					}
					// declaration of a variable without initial value
					if len(spec.Values) == 0 {
						continue
					}
					switch compLit := spec.Values[0].(type) {
					case *ast.CompositeLit:
						var key string
						_ = key
						var value string
						_ = value
						for _, elt := range compLit.Elts {

							// each elt is an expression for struct or for field such as
							// for struct
							//
							//         "dummy.Dummy": &(dummy.Dummy{})
							//
							// or, for field
							//
							//          "dummy.Dummy.Name": (dummy.Dummy{}).Name,
							//
							// first node in the AST is a key value expression
							var ok bool
							var kve *ast.KeyValueExpr
							if kve, ok = elt.(*ast.KeyValueExpr); !ok {
								log.Fatal("Expression should be key value expression" +
									fset.Position(kve.Pos()).String())
							}

							switch bl := kve.Key.(type) {
							case *ast.BasicLit:
								key = bl.Value // "\"dumm.Dummy\"" is the value

								// one remove the ambracing double quotes
								key = strings.TrimPrefix(key, "\"")
								key = strings.TrimSuffix(key, "\"")
							}
							var expressionType GONG__ExpressionType = GONG__STRUCT_INSTANCE
							var docLink GONG__Identifier

							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								expressionType = GONG__FIELD_OR_CONST_VALUE
							}

							var callExpr *ast.CallExpr
							if callExpr, ok = kve.Value.(*ast.CallExpr); ok {

								var se *ast.SelectorExpr
								if se, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(callExpr.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}

								// check the arg type to select wether this is a int or a string enum
								var bl *ast.BasicLit
								if bl, ok = callExpr.Args[0].(*ast.BasicLit); ok {
									switch bl.Kind {
									case token.STRING:
										expressionType = GONG__ENUM_CAST_STRING
									case token.INT:
										expressionType = GONG__ENUM_CAST_INT
									}
								} else {
									log.Fatal("Expression should be a basic lit" +
										fset.Position(se.Pos()).String())
								}

								docLink.Ident = id.Name + "." + se.Sel.Name
								_ = callExpr
							}

							var se2 *ast.SelectorExpr
							switch expressionType {
							case GONG__FIELD_OR_CONST_VALUE:
								if se2, ok = kve.Value.(*ast.SelectorExpr); ok {

									var ident *ast.Ident
									if _, ok = se2.X.(*ast.ParenExpr); ok {
										expressionType = GONG__FIELD_VALUE
										fieldName = se2.Sel.Name
									} else if ident, ok = se2.X.(*ast.Ident); ok {
										expressionType = GONG__IDENTIFIER_CONST
										docLink.Ident = ident.Name + "." + se2.Sel.Name
									} else {
										log.Fatal("Expression should be a selector expression or an ident" +
											fset.Position(kve.Pos()).String())
									}
								} else {

								}
							}

							var pe *ast.ParenExpr
							switch expressionType {
							case GONG__STRUCT_INSTANCE:
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							case GONG__FIELD_VALUE:
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}
							switch expressionType {
							case GONG__FIELD_VALUE, GONG__STRUCT_INSTANCE:
								// expect a Composite Litteral with no Element <type>{}
								var cl *ast.CompositeLit
								if cl, ok = pe.X.(*ast.CompositeLit); !ok {
									log.Fatal("Expression should be a composite lit" +
										fset.Position(pe.Pos()).String())
								}

								var se *ast.SelectorExpr
								if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector" +
										fset.Position(cl.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}
								docLink.Ident = id.Name + "." + se.Sel.Name
							}

							switch expressionType {
							case GONG__FIELD_VALUE:
								docLink.Ident += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink.Ident == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							docLink.Type = expressionType
							stage.Map_DocLink_Renaming[key] = docLink
						}
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_Accidental = make(map[string]*Accidental)
var __gong__map_Accidental_mark = make(map[string]*Accidental_mark)
var __gong__map_Accidental_text = make(map[string]*Accidental_text)
var __gong__map_Accord = make(map[string]*Accord)
var __gong__map_Accordion_registration = make(map[string]*Accordion_registration)
var __gong__map_AnyType = make(map[string]*AnyType)
var __gong__map_Appearance = make(map[string]*Appearance)
var __gong__map_Arpeggiate = make(map[string]*Arpeggiate)
var __gong__map_Arrow = make(map[string]*Arrow)
var __gong__map_Articulations = make(map[string]*Articulations)
var __gong__map_Assess = make(map[string]*Assess)
var __gong__map_Attributes = make(map[string]*Attributes)
var __gong__map_Backup = make(map[string]*Backup)
var __gong__map_Bar_style_color = make(map[string]*Bar_style_color)
var __gong__map_Barline = make(map[string]*Barline)
var __gong__map_Barre = make(map[string]*Barre)
var __gong__map_Bass = make(map[string]*Bass)
var __gong__map_Bass_step = make(map[string]*Bass_step)
var __gong__map_Beam = make(map[string]*Beam)
var __gong__map_Beat_repeat = make(map[string]*Beat_repeat)
var __gong__map_Beat_unit_tied = make(map[string]*Beat_unit_tied)
var __gong__map_Beater = make(map[string]*Beater)
var __gong__map_Bend = make(map[string]*Bend)
var __gong__map_Bookmark = make(map[string]*Bookmark)
var __gong__map_Bracket = make(map[string]*Bracket)
var __gong__map_Breath_mark = make(map[string]*Breath_mark)
var __gong__map_Caesura = make(map[string]*Caesura)
var __gong__map_Cancel = make(map[string]*Cancel)
var __gong__map_Clef = make(map[string]*Clef)
var __gong__map_Coda = make(map[string]*Coda)
var __gong__map_Credit = make(map[string]*Credit)
var __gong__map_Dashes = make(map[string]*Dashes)
var __gong__map_Defaults = make(map[string]*Defaults)
var __gong__map_Degree = make(map[string]*Degree)
var __gong__map_Degree_alter = make(map[string]*Degree_alter)
var __gong__map_Degree_type = make(map[string]*Degree_type)
var __gong__map_Degree_value = make(map[string]*Degree_value)
var __gong__map_Direction = make(map[string]*Direction)
var __gong__map_Direction_type = make(map[string]*Direction_type)
var __gong__map_Distance = make(map[string]*Distance)
var __gong__map_Double = make(map[string]*Double)
var __gong__map_Dynamics = make(map[string]*Dynamics)
var __gong__map_Effect = make(map[string]*Effect)
var __gong__map_Elision = make(map[string]*Elision)
var __gong__map_Empty = make(map[string]*Empty)
var __gong__map_Empty_font = make(map[string]*Empty_font)
var __gong__map_Empty_line = make(map[string]*Empty_line)
var __gong__map_Empty_placement = make(map[string]*Empty_placement)
var __gong__map_Empty_placement_smufl = make(map[string]*Empty_placement_smufl)
var __gong__map_Empty_print_object_style_align = make(map[string]*Empty_print_object_style_align)
var __gong__map_Empty_print_style = make(map[string]*Empty_print_style)
var __gong__map_Empty_print_style_align = make(map[string]*Empty_print_style_align)
var __gong__map_Empty_print_style_align_id = make(map[string]*Empty_print_style_align_id)
var __gong__map_Empty_trill_sound = make(map[string]*Empty_trill_sound)
var __gong__map_Encoding = make(map[string]*Encoding)
var __gong__map_Ending = make(map[string]*Ending)
var __gong__map_Extend = make(map[string]*Extend)
var __gong__map_Feature = make(map[string]*Feature)
var __gong__map_Fermata = make(map[string]*Fermata)
var __gong__map_Figure = make(map[string]*Figure)
var __gong__map_Figured_bass = make(map[string]*Figured_bass)
var __gong__map_Fingering = make(map[string]*Fingering)
var __gong__map_First_fret = make(map[string]*First_fret)
var __gong__map_Foo = make(map[string]*Foo)
var __gong__map_For_part = make(map[string]*For_part)
var __gong__map_Formatted_symbol = make(map[string]*Formatted_symbol)
var __gong__map_Formatted_symbol_id = make(map[string]*Formatted_symbol_id)
var __gong__map_Forward = make(map[string]*Forward)
var __gong__map_Frame = make(map[string]*Frame)
var __gong__map_Frame_note = make(map[string]*Frame_note)
var __gong__map_Fret = make(map[string]*Fret)
var __gong__map_Glass = make(map[string]*Glass)
var __gong__map_Glissando = make(map[string]*Glissando)
var __gong__map_Glyph = make(map[string]*Glyph)
var __gong__map_Grace = make(map[string]*Grace)
var __gong__map_Group_barline = make(map[string]*Group_barline)
var __gong__map_Group_symbol = make(map[string]*Group_symbol)
var __gong__map_Grouping = make(map[string]*Grouping)
var __gong__map_Hammer_on_pull_off = make(map[string]*Hammer_on_pull_off)
var __gong__map_Handbell = make(map[string]*Handbell)
var __gong__map_Harmon_closed = make(map[string]*Harmon_closed)
var __gong__map_Harmon_mute = make(map[string]*Harmon_mute)
var __gong__map_Harmonic = make(map[string]*Harmonic)
var __gong__map_Harmony = make(map[string]*Harmony)
var __gong__map_Harmony_alter = make(map[string]*Harmony_alter)
var __gong__map_Harp_pedals = make(map[string]*Harp_pedals)
var __gong__map_Heel_toe = make(map[string]*Heel_toe)
var __gong__map_Hole = make(map[string]*Hole)
var __gong__map_Hole_closed = make(map[string]*Hole_closed)
var __gong__map_Horizontal_turn = make(map[string]*Horizontal_turn)
var __gong__map_Identification = make(map[string]*Identification)
var __gong__map_Image = make(map[string]*Image)
var __gong__map_Instrument = make(map[string]*Instrument)
var __gong__map_Instrument_change = make(map[string]*Instrument_change)
var __gong__map_Instrument_link = make(map[string]*Instrument_link)
var __gong__map_Interchangeable = make(map[string]*Interchangeable)
var __gong__map_Inversion = make(map[string]*Inversion)
var __gong__map_Key = make(map[string]*Key)
var __gong__map_Key_accidental = make(map[string]*Key_accidental)
var __gong__map_Key_octave = make(map[string]*Key_octave)
var __gong__map_Kind = make(map[string]*Kind)
var __gong__map_Level = make(map[string]*Level)
var __gong__map_Line_detail = make(map[string]*Line_detail)
var __gong__map_Line_width = make(map[string]*Line_width)
var __gong__map_Link = make(map[string]*Link)
var __gong__map_Listen = make(map[string]*Listen)
var __gong__map_Listening = make(map[string]*Listening)
var __gong__map_Lyric = make(map[string]*Lyric)
var __gong__map_Lyric_font = make(map[string]*Lyric_font)
var __gong__map_Lyric_language = make(map[string]*Lyric_language)
var __gong__map_Measure_layout = make(map[string]*Measure_layout)
var __gong__map_Measure_numbering = make(map[string]*Measure_numbering)
var __gong__map_Measure_repeat = make(map[string]*Measure_repeat)
var __gong__map_Measure_style = make(map[string]*Measure_style)
var __gong__map_Membrane = make(map[string]*Membrane)
var __gong__map_Metal = make(map[string]*Metal)
var __gong__map_Metronome = make(map[string]*Metronome)
var __gong__map_Metronome_beam = make(map[string]*Metronome_beam)
var __gong__map_Metronome_note = make(map[string]*Metronome_note)
var __gong__map_Metronome_tied = make(map[string]*Metronome_tied)
var __gong__map_Metronome_tuplet = make(map[string]*Metronome_tuplet)
var __gong__map_Midi_device = make(map[string]*Midi_device)
var __gong__map_Midi_instrument = make(map[string]*Midi_instrument)
var __gong__map_Miscellaneous = make(map[string]*Miscellaneous)
var __gong__map_Miscellaneous_field = make(map[string]*Miscellaneous_field)
var __gong__map_Mordent = make(map[string]*Mordent)
var __gong__map_Multiple_rest = make(map[string]*Multiple_rest)
var __gong__map_Name_display = make(map[string]*Name_display)
var __gong__map_Non_arpeggiate = make(map[string]*Non_arpeggiate)
var __gong__map_Notations = make(map[string]*Notations)
var __gong__map_Note = make(map[string]*Note)
var __gong__map_Note_size = make(map[string]*Note_size)
var __gong__map_Note_type = make(map[string]*Note_type)
var __gong__map_Notehead = make(map[string]*Notehead)
var __gong__map_Notehead_text = make(map[string]*Notehead_text)
var __gong__map_Numeral = make(map[string]*Numeral)
var __gong__map_Numeral_key = make(map[string]*Numeral_key)
var __gong__map_Numeral_root = make(map[string]*Numeral_root)
var __gong__map_Octave_shift = make(map[string]*Octave_shift)
var __gong__map_Offset = make(map[string]*Offset)
var __gong__map_Opus = make(map[string]*Opus)
var __gong__map_Ornaments = make(map[string]*Ornaments)
var __gong__map_Other_appearance = make(map[string]*Other_appearance)
var __gong__map_Other_listening = make(map[string]*Other_listening)
var __gong__map_Other_notation = make(map[string]*Other_notation)
var __gong__map_Other_play = make(map[string]*Other_play)
var __gong__map_Page_layout = make(map[string]*Page_layout)
var __gong__map_Page_margins = make(map[string]*Page_margins)
var __gong__map_Part_clef = make(map[string]*Part_clef)
var __gong__map_Part_group = make(map[string]*Part_group)
var __gong__map_Part_link = make(map[string]*Part_link)
var __gong__map_Part_list = make(map[string]*Part_list)
var __gong__map_Part_symbol = make(map[string]*Part_symbol)
var __gong__map_Part_transpose = make(map[string]*Part_transpose)
var __gong__map_Pedal = make(map[string]*Pedal)
var __gong__map_Pedal_tuning = make(map[string]*Pedal_tuning)
var __gong__map_Percussion = make(map[string]*Percussion)
var __gong__map_Pitch = make(map[string]*Pitch)
var __gong__map_Pitched = make(map[string]*Pitched)
var __gong__map_Play = make(map[string]*Play)
var __gong__map_Player = make(map[string]*Player)
var __gong__map_Principal_voice = make(map[string]*Principal_voice)
var __gong__map_Print = make(map[string]*Print)
var __gong__map_Release = make(map[string]*Release)
var __gong__map_Repeat = make(map[string]*Repeat)
var __gong__map_Rest = make(map[string]*Rest)
var __gong__map_Root = make(map[string]*Root)
var __gong__map_Root_step = make(map[string]*Root_step)
var __gong__map_Scaling = make(map[string]*Scaling)
var __gong__map_Scordatura = make(map[string]*Scordatura)
var __gong__map_Score_instrument = make(map[string]*Score_instrument)
var __gong__map_Score_part = make(map[string]*Score_part)
var __gong__map_Score_partwise = make(map[string]*Score_partwise)
var __gong__map_Score_timewise = make(map[string]*Score_timewise)
var __gong__map_Segno = make(map[string]*Segno)
var __gong__map_Slash = make(map[string]*Slash)
var __gong__map_Slide = make(map[string]*Slide)
var __gong__map_Slur = make(map[string]*Slur)
var __gong__map_Sound = make(map[string]*Sound)
var __gong__map_Staff_details = make(map[string]*Staff_details)
var __gong__map_Staff_divide = make(map[string]*Staff_divide)
var __gong__map_Staff_layout = make(map[string]*Staff_layout)
var __gong__map_Staff_size = make(map[string]*Staff_size)
var __gong__map_Staff_tuning = make(map[string]*Staff_tuning)
var __gong__map_Stem = make(map[string]*Stem)
var __gong__map_Stick = make(map[string]*Stick)
var __gong__map_String_mute = make(map[string]*String_mute)
var __gong__map_Strong_accent = make(map[string]*Strong_accent)
var __gong__map_Supports = make(map[string]*Supports)
var __gong__map_Swing = make(map[string]*Swing)
var __gong__map_Sync = make(map[string]*Sync)
var __gong__map_System_dividers = make(map[string]*System_dividers)
var __gong__map_System_layout = make(map[string]*System_layout)
var __gong__map_System_margins = make(map[string]*System_margins)
var __gong__map_Tap = make(map[string]*Tap)
var __gong__map_Technical = make(map[string]*Technical)
var __gong__map_Text_element_data = make(map[string]*Text_element_data)
var __gong__map_Tie = make(map[string]*Tie)
var __gong__map_Tied = make(map[string]*Tied)
var __gong__map_Time = make(map[string]*Time)
var __gong__map_Time_modification = make(map[string]*Time_modification)
var __gong__map_Timpani = make(map[string]*Timpani)
var __gong__map_Transpose = make(map[string]*Transpose)
var __gong__map_Tremolo = make(map[string]*Tremolo)
var __gong__map_Tuplet = make(map[string]*Tuplet)
var __gong__map_Tuplet_dot = make(map[string]*Tuplet_dot)
var __gong__map_Tuplet_number = make(map[string]*Tuplet_number)
var __gong__map_Tuplet_portion = make(map[string]*Tuplet_portion)
var __gong__map_Tuplet_type = make(map[string]*Tuplet_type)
var __gong__map_Typed_text = make(map[string]*Typed_text)
var __gong__map_Unpitched = make(map[string]*Unpitched)
var __gong__map_Virtual_instrument = make(map[string]*Virtual_instrument)
var __gong__map_Wait = make(map[string]*Wait)
var __gong__map_Wavy_line = make(map[string]*Wavy_line)
var __gong__map_Wedge = make(map[string]*Wedge)
var __gong__map_Wood = make(map[string]*Wood)
var __gong__map_Work = make(map[string]*Work)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}
func lookupSym(recv, name string) (ok bool) {
	if recv == "" {
		return true
	}
	return false
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *StageStruct, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {

	// used for debug purposes
	astCoordinate := "\tAssignStmt: "

	//
	// First parse all comment groups in the assignement
	// if a comment "//gong:ident [DocLink]" is met and is followed by a string assignement.
	// modify the following AST assignement to assigns the DocLink text to the string value
	//

	// get the comment group of the assigStmt
	commentGroups := (*cmap)[assignStmt]
	// get the the prefix
	var hasGongIdentDirective bool
	var commentText string
	var docLinkText string
	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "//gong:ident") {
				hasGongIdentDirective = true
				commentText = comment.Text
			}
		}
	}
	if hasGongIdentDirective {
		// parser configured to find doclinks
		var docLinkFinder comment.Parser
		docLinkFinder.LookupPackage = lookupPackage
		docLinkFinder.LookupSym = lookupSym
		doc := docLinkFinder.Parse(commentText)

		for _, block := range doc.Content {
			switch paragraph := block.(type) {
			case *comment.Paragraph:
				_ = paragraph
				for _, text := range paragraph.Text {
					switch docLink := text.(type) {
					case *comment.DocLink:
						if docLink.Recv == "" {
							docLinkText = docLink.ImportPath + "." + docLink.Name
						} else {
							docLinkText = docLink.ImportPath + "." + docLink.Recv + "." + docLink.Name
						}

						// we check wether the doc link has been renamed
						// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
						if renamed, ok := (stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed.Ident
						}
					}
				}
			}
		}
	}

	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "Accidental":
										instanceAccidental := (&Accidental{Name: instanceName}).Stage(stage)
										instance = any(instanceAccidental)
										__gong__map_Accidental[identifier] = instanceAccidental
									case "Accidental_mark":
										instanceAccidental_mark := (&Accidental_mark{Name: instanceName}).Stage(stage)
										instance = any(instanceAccidental_mark)
										__gong__map_Accidental_mark[identifier] = instanceAccidental_mark
									case "Accidental_text":
										instanceAccidental_text := (&Accidental_text{Name: instanceName}).Stage(stage)
										instance = any(instanceAccidental_text)
										__gong__map_Accidental_text[identifier] = instanceAccidental_text
									case "Accord":
										instanceAccord := (&Accord{Name: instanceName}).Stage(stage)
										instance = any(instanceAccord)
										__gong__map_Accord[identifier] = instanceAccord
									case "Accordion_registration":
										instanceAccordion_registration := (&Accordion_registration{Name: instanceName}).Stage(stage)
										instance = any(instanceAccordion_registration)
										__gong__map_Accordion_registration[identifier] = instanceAccordion_registration
									case "AnyType":
										instanceAnyType := (&AnyType{Name: instanceName}).Stage(stage)
										instance = any(instanceAnyType)
										__gong__map_AnyType[identifier] = instanceAnyType
									case "Appearance":
										instanceAppearance := (&Appearance{Name: instanceName}).Stage(stage)
										instance = any(instanceAppearance)
										__gong__map_Appearance[identifier] = instanceAppearance
									case "Arpeggiate":
										instanceArpeggiate := (&Arpeggiate{Name: instanceName}).Stage(stage)
										instance = any(instanceArpeggiate)
										__gong__map_Arpeggiate[identifier] = instanceArpeggiate
									case "Arrow":
										instanceArrow := (&Arrow{Name: instanceName}).Stage(stage)
										instance = any(instanceArrow)
										__gong__map_Arrow[identifier] = instanceArrow
									case "Articulations":
										instanceArticulations := (&Articulations{Name: instanceName}).Stage(stage)
										instance = any(instanceArticulations)
										__gong__map_Articulations[identifier] = instanceArticulations
									case "Assess":
										instanceAssess := (&Assess{Name: instanceName}).Stage(stage)
										instance = any(instanceAssess)
										__gong__map_Assess[identifier] = instanceAssess
									case "Attributes":
										instanceAttributes := (&Attributes{Name: instanceName}).Stage(stage)
										instance = any(instanceAttributes)
										__gong__map_Attributes[identifier] = instanceAttributes
									case "Backup":
										instanceBackup := (&Backup{Name: instanceName}).Stage(stage)
										instance = any(instanceBackup)
										__gong__map_Backup[identifier] = instanceBackup
									case "Bar_style_color":
										instanceBar_style_color := (&Bar_style_color{Name: instanceName}).Stage(stage)
										instance = any(instanceBar_style_color)
										__gong__map_Bar_style_color[identifier] = instanceBar_style_color
									case "Barline":
										instanceBarline := (&Barline{Name: instanceName}).Stage(stage)
										instance = any(instanceBarline)
										__gong__map_Barline[identifier] = instanceBarline
									case "Barre":
										instanceBarre := (&Barre{Name: instanceName}).Stage(stage)
										instance = any(instanceBarre)
										__gong__map_Barre[identifier] = instanceBarre
									case "Bass":
										instanceBass := (&Bass{Name: instanceName}).Stage(stage)
										instance = any(instanceBass)
										__gong__map_Bass[identifier] = instanceBass
									case "Bass_step":
										instanceBass_step := (&Bass_step{Name: instanceName}).Stage(stage)
										instance = any(instanceBass_step)
										__gong__map_Bass_step[identifier] = instanceBass_step
									case "Beam":
										instanceBeam := (&Beam{Name: instanceName}).Stage(stage)
										instance = any(instanceBeam)
										__gong__map_Beam[identifier] = instanceBeam
									case "Beat_repeat":
										instanceBeat_repeat := (&Beat_repeat{Name: instanceName}).Stage(stage)
										instance = any(instanceBeat_repeat)
										__gong__map_Beat_repeat[identifier] = instanceBeat_repeat
									case "Beat_unit_tied":
										instanceBeat_unit_tied := (&Beat_unit_tied{Name: instanceName}).Stage(stage)
										instance = any(instanceBeat_unit_tied)
										__gong__map_Beat_unit_tied[identifier] = instanceBeat_unit_tied
									case "Beater":
										instanceBeater := (&Beater{Name: instanceName}).Stage(stage)
										instance = any(instanceBeater)
										__gong__map_Beater[identifier] = instanceBeater
									case "Bend":
										instanceBend := (&Bend{Name: instanceName}).Stage(stage)
										instance = any(instanceBend)
										__gong__map_Bend[identifier] = instanceBend
									case "Bookmark":
										instanceBookmark := (&Bookmark{Name: instanceName}).Stage(stage)
										instance = any(instanceBookmark)
										__gong__map_Bookmark[identifier] = instanceBookmark
									case "Bracket":
										instanceBracket := (&Bracket{Name: instanceName}).Stage(stage)
										instance = any(instanceBracket)
										__gong__map_Bracket[identifier] = instanceBracket
									case "Breath_mark":
										instanceBreath_mark := (&Breath_mark{Name: instanceName}).Stage(stage)
										instance = any(instanceBreath_mark)
										__gong__map_Breath_mark[identifier] = instanceBreath_mark
									case "Caesura":
										instanceCaesura := (&Caesura{Name: instanceName}).Stage(stage)
										instance = any(instanceCaesura)
										__gong__map_Caesura[identifier] = instanceCaesura
									case "Cancel":
										instanceCancel := (&Cancel{Name: instanceName}).Stage(stage)
										instance = any(instanceCancel)
										__gong__map_Cancel[identifier] = instanceCancel
									case "Clef":
										instanceClef := (&Clef{Name: instanceName}).Stage(stage)
										instance = any(instanceClef)
										__gong__map_Clef[identifier] = instanceClef
									case "Coda":
										instanceCoda := (&Coda{Name: instanceName}).Stage(stage)
										instance = any(instanceCoda)
										__gong__map_Coda[identifier] = instanceCoda
									case "Credit":
										instanceCredit := (&Credit{Name: instanceName}).Stage(stage)
										instance = any(instanceCredit)
										__gong__map_Credit[identifier] = instanceCredit
									case "Dashes":
										instanceDashes := (&Dashes{Name: instanceName}).Stage(stage)
										instance = any(instanceDashes)
										__gong__map_Dashes[identifier] = instanceDashes
									case "Defaults":
										instanceDefaults := (&Defaults{Name: instanceName}).Stage(stage)
										instance = any(instanceDefaults)
										__gong__map_Defaults[identifier] = instanceDefaults
									case "Degree":
										instanceDegree := (&Degree{Name: instanceName}).Stage(stage)
										instance = any(instanceDegree)
										__gong__map_Degree[identifier] = instanceDegree
									case "Degree_alter":
										instanceDegree_alter := (&Degree_alter{Name: instanceName}).Stage(stage)
										instance = any(instanceDegree_alter)
										__gong__map_Degree_alter[identifier] = instanceDegree_alter
									case "Degree_type":
										instanceDegree_type := (&Degree_type{Name: instanceName}).Stage(stage)
										instance = any(instanceDegree_type)
										__gong__map_Degree_type[identifier] = instanceDegree_type
									case "Degree_value":
										instanceDegree_value := (&Degree_value{Name: instanceName}).Stage(stage)
										instance = any(instanceDegree_value)
										__gong__map_Degree_value[identifier] = instanceDegree_value
									case "Direction":
										instanceDirection := (&Direction{Name: instanceName}).Stage(stage)
										instance = any(instanceDirection)
										__gong__map_Direction[identifier] = instanceDirection
									case "Direction_type":
										instanceDirection_type := (&Direction_type{Name: instanceName}).Stage(stage)
										instance = any(instanceDirection_type)
										__gong__map_Direction_type[identifier] = instanceDirection_type
									case "Distance":
										instanceDistance := (&Distance{Name: instanceName}).Stage(stage)
										instance = any(instanceDistance)
										__gong__map_Distance[identifier] = instanceDistance
									case "Double":
										instanceDouble := (&Double{Name: instanceName}).Stage(stage)
										instance = any(instanceDouble)
										__gong__map_Double[identifier] = instanceDouble
									case "Dynamics":
										instanceDynamics := (&Dynamics{Name: instanceName}).Stage(stage)
										instance = any(instanceDynamics)
										__gong__map_Dynamics[identifier] = instanceDynamics
									case "Effect":
										instanceEffect := (&Effect{Name: instanceName}).Stage(stage)
										instance = any(instanceEffect)
										__gong__map_Effect[identifier] = instanceEffect
									case "Elision":
										instanceElision := (&Elision{Name: instanceName}).Stage(stage)
										instance = any(instanceElision)
										__gong__map_Elision[identifier] = instanceElision
									case "Empty":
										instanceEmpty := (&Empty{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty)
										__gong__map_Empty[identifier] = instanceEmpty
									case "Empty_font":
										instanceEmpty_font := (&Empty_font{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty_font)
										__gong__map_Empty_font[identifier] = instanceEmpty_font
									case "Empty_line":
										instanceEmpty_line := (&Empty_line{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty_line)
										__gong__map_Empty_line[identifier] = instanceEmpty_line
									case "Empty_placement":
										instanceEmpty_placement := (&Empty_placement{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty_placement)
										__gong__map_Empty_placement[identifier] = instanceEmpty_placement
									case "Empty_placement_smufl":
										instanceEmpty_placement_smufl := (&Empty_placement_smufl{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty_placement_smufl)
										__gong__map_Empty_placement_smufl[identifier] = instanceEmpty_placement_smufl
									case "Empty_print_object_style_align":
										instanceEmpty_print_object_style_align := (&Empty_print_object_style_align{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty_print_object_style_align)
										__gong__map_Empty_print_object_style_align[identifier] = instanceEmpty_print_object_style_align
									case "Empty_print_style":
										instanceEmpty_print_style := (&Empty_print_style{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty_print_style)
										__gong__map_Empty_print_style[identifier] = instanceEmpty_print_style
									case "Empty_print_style_align":
										instanceEmpty_print_style_align := (&Empty_print_style_align{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty_print_style_align)
										__gong__map_Empty_print_style_align[identifier] = instanceEmpty_print_style_align
									case "Empty_print_style_align_id":
										instanceEmpty_print_style_align_id := (&Empty_print_style_align_id{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty_print_style_align_id)
										__gong__map_Empty_print_style_align_id[identifier] = instanceEmpty_print_style_align_id
									case "Empty_trill_sound":
										instanceEmpty_trill_sound := (&Empty_trill_sound{Name: instanceName}).Stage(stage)
										instance = any(instanceEmpty_trill_sound)
										__gong__map_Empty_trill_sound[identifier] = instanceEmpty_trill_sound
									case "Encoding":
										instanceEncoding := (&Encoding{Name: instanceName}).Stage(stage)
										instance = any(instanceEncoding)
										__gong__map_Encoding[identifier] = instanceEncoding
									case "Ending":
										instanceEnding := (&Ending{Name: instanceName}).Stage(stage)
										instance = any(instanceEnding)
										__gong__map_Ending[identifier] = instanceEnding
									case "Extend":
										instanceExtend := (&Extend{Name: instanceName}).Stage(stage)
										instance = any(instanceExtend)
										__gong__map_Extend[identifier] = instanceExtend
									case "Feature":
										instanceFeature := (&Feature{Name: instanceName}).Stage(stage)
										instance = any(instanceFeature)
										__gong__map_Feature[identifier] = instanceFeature
									case "Fermata":
										instanceFermata := (&Fermata{Name: instanceName}).Stage(stage)
										instance = any(instanceFermata)
										__gong__map_Fermata[identifier] = instanceFermata
									case "Figure":
										instanceFigure := (&Figure{Name: instanceName}).Stage(stage)
										instance = any(instanceFigure)
										__gong__map_Figure[identifier] = instanceFigure
									case "Figured_bass":
										instanceFigured_bass := (&Figured_bass{Name: instanceName}).Stage(stage)
										instance = any(instanceFigured_bass)
										__gong__map_Figured_bass[identifier] = instanceFigured_bass
									case "Fingering":
										instanceFingering := (&Fingering{Name: instanceName}).Stage(stage)
										instance = any(instanceFingering)
										__gong__map_Fingering[identifier] = instanceFingering
									case "First_fret":
										instanceFirst_fret := (&First_fret{Name: instanceName}).Stage(stage)
										instance = any(instanceFirst_fret)
										__gong__map_First_fret[identifier] = instanceFirst_fret
									case "Foo":
										instanceFoo := (&Foo{Name: instanceName}).Stage(stage)
										instance = any(instanceFoo)
										__gong__map_Foo[identifier] = instanceFoo
									case "For_part":
										instanceFor_part := (&For_part{Name: instanceName}).Stage(stage)
										instance = any(instanceFor_part)
										__gong__map_For_part[identifier] = instanceFor_part
									case "Formatted_symbol":
										instanceFormatted_symbol := (&Formatted_symbol{Name: instanceName}).Stage(stage)
										instance = any(instanceFormatted_symbol)
										__gong__map_Formatted_symbol[identifier] = instanceFormatted_symbol
									case "Formatted_symbol_id":
										instanceFormatted_symbol_id := (&Formatted_symbol_id{Name: instanceName}).Stage(stage)
										instance = any(instanceFormatted_symbol_id)
										__gong__map_Formatted_symbol_id[identifier] = instanceFormatted_symbol_id
									case "Forward":
										instanceForward := (&Forward{Name: instanceName}).Stage(stage)
										instance = any(instanceForward)
										__gong__map_Forward[identifier] = instanceForward
									case "Frame":
										instanceFrame := (&Frame{Name: instanceName}).Stage(stage)
										instance = any(instanceFrame)
										__gong__map_Frame[identifier] = instanceFrame
									case "Frame_note":
										instanceFrame_note := (&Frame_note{Name: instanceName}).Stage(stage)
										instance = any(instanceFrame_note)
										__gong__map_Frame_note[identifier] = instanceFrame_note
									case "Fret":
										instanceFret := (&Fret{Name: instanceName}).Stage(stage)
										instance = any(instanceFret)
										__gong__map_Fret[identifier] = instanceFret
									case "Glass":
										instanceGlass := (&Glass{Name: instanceName}).Stage(stage)
										instance = any(instanceGlass)
										__gong__map_Glass[identifier] = instanceGlass
									case "Glissando":
										instanceGlissando := (&Glissando{Name: instanceName}).Stage(stage)
										instance = any(instanceGlissando)
										__gong__map_Glissando[identifier] = instanceGlissando
									case "Glyph":
										instanceGlyph := (&Glyph{Name: instanceName}).Stage(stage)
										instance = any(instanceGlyph)
										__gong__map_Glyph[identifier] = instanceGlyph
									case "Grace":
										instanceGrace := (&Grace{Name: instanceName}).Stage(stage)
										instance = any(instanceGrace)
										__gong__map_Grace[identifier] = instanceGrace
									case "Group_barline":
										instanceGroup_barline := (&Group_barline{Name: instanceName}).Stage(stage)
										instance = any(instanceGroup_barline)
										__gong__map_Group_barline[identifier] = instanceGroup_barline
									case "Group_symbol":
										instanceGroup_symbol := (&Group_symbol{Name: instanceName}).Stage(stage)
										instance = any(instanceGroup_symbol)
										__gong__map_Group_symbol[identifier] = instanceGroup_symbol
									case "Grouping":
										instanceGrouping := (&Grouping{Name: instanceName}).Stage(stage)
										instance = any(instanceGrouping)
										__gong__map_Grouping[identifier] = instanceGrouping
									case "Hammer_on_pull_off":
										instanceHammer_on_pull_off := (&Hammer_on_pull_off{Name: instanceName}).Stage(stage)
										instance = any(instanceHammer_on_pull_off)
										__gong__map_Hammer_on_pull_off[identifier] = instanceHammer_on_pull_off
									case "Handbell":
										instanceHandbell := (&Handbell{Name: instanceName}).Stage(stage)
										instance = any(instanceHandbell)
										__gong__map_Handbell[identifier] = instanceHandbell
									case "Harmon_closed":
										instanceHarmon_closed := (&Harmon_closed{Name: instanceName}).Stage(stage)
										instance = any(instanceHarmon_closed)
										__gong__map_Harmon_closed[identifier] = instanceHarmon_closed
									case "Harmon_mute":
										instanceHarmon_mute := (&Harmon_mute{Name: instanceName}).Stage(stage)
										instance = any(instanceHarmon_mute)
										__gong__map_Harmon_mute[identifier] = instanceHarmon_mute
									case "Harmonic":
										instanceHarmonic := (&Harmonic{Name: instanceName}).Stage(stage)
										instance = any(instanceHarmonic)
										__gong__map_Harmonic[identifier] = instanceHarmonic
									case "Harmony":
										instanceHarmony := (&Harmony{Name: instanceName}).Stage(stage)
										instance = any(instanceHarmony)
										__gong__map_Harmony[identifier] = instanceHarmony
									case "Harmony_alter":
										instanceHarmony_alter := (&Harmony_alter{Name: instanceName}).Stage(stage)
										instance = any(instanceHarmony_alter)
										__gong__map_Harmony_alter[identifier] = instanceHarmony_alter
									case "Harp_pedals":
										instanceHarp_pedals := (&Harp_pedals{Name: instanceName}).Stage(stage)
										instance = any(instanceHarp_pedals)
										__gong__map_Harp_pedals[identifier] = instanceHarp_pedals
									case "Heel_toe":
										instanceHeel_toe := (&Heel_toe{Name: instanceName}).Stage(stage)
										instance = any(instanceHeel_toe)
										__gong__map_Heel_toe[identifier] = instanceHeel_toe
									case "Hole":
										instanceHole := (&Hole{Name: instanceName}).Stage(stage)
										instance = any(instanceHole)
										__gong__map_Hole[identifier] = instanceHole
									case "Hole_closed":
										instanceHole_closed := (&Hole_closed{Name: instanceName}).Stage(stage)
										instance = any(instanceHole_closed)
										__gong__map_Hole_closed[identifier] = instanceHole_closed
									case "Horizontal_turn":
										instanceHorizontal_turn := (&Horizontal_turn{Name: instanceName}).Stage(stage)
										instance = any(instanceHorizontal_turn)
										__gong__map_Horizontal_turn[identifier] = instanceHorizontal_turn
									case "Identification":
										instanceIdentification := (&Identification{Name: instanceName}).Stage(stage)
										instance = any(instanceIdentification)
										__gong__map_Identification[identifier] = instanceIdentification
									case "Image":
										instanceImage := (&Image{Name: instanceName}).Stage(stage)
										instance = any(instanceImage)
										__gong__map_Image[identifier] = instanceImage
									case "Instrument":
										instanceInstrument := (&Instrument{Name: instanceName}).Stage(stage)
										instance = any(instanceInstrument)
										__gong__map_Instrument[identifier] = instanceInstrument
									case "Instrument_change":
										instanceInstrument_change := (&Instrument_change{Name: instanceName}).Stage(stage)
										instance = any(instanceInstrument_change)
										__gong__map_Instrument_change[identifier] = instanceInstrument_change
									case "Instrument_link":
										instanceInstrument_link := (&Instrument_link{Name: instanceName}).Stage(stage)
										instance = any(instanceInstrument_link)
										__gong__map_Instrument_link[identifier] = instanceInstrument_link
									case "Interchangeable":
										instanceInterchangeable := (&Interchangeable{Name: instanceName}).Stage(stage)
										instance = any(instanceInterchangeable)
										__gong__map_Interchangeable[identifier] = instanceInterchangeable
									case "Inversion":
										instanceInversion := (&Inversion{Name: instanceName}).Stage(stage)
										instance = any(instanceInversion)
										__gong__map_Inversion[identifier] = instanceInversion
									case "Key":
										instanceKey := (&Key{Name: instanceName}).Stage(stage)
										instance = any(instanceKey)
										__gong__map_Key[identifier] = instanceKey
									case "Key_accidental":
										instanceKey_accidental := (&Key_accidental{Name: instanceName}).Stage(stage)
										instance = any(instanceKey_accidental)
										__gong__map_Key_accidental[identifier] = instanceKey_accidental
									case "Key_octave":
										instanceKey_octave := (&Key_octave{Name: instanceName}).Stage(stage)
										instance = any(instanceKey_octave)
										__gong__map_Key_octave[identifier] = instanceKey_octave
									case "Kind":
										instanceKind := (&Kind{Name: instanceName}).Stage(stage)
										instance = any(instanceKind)
										__gong__map_Kind[identifier] = instanceKind
									case "Level":
										instanceLevel := (&Level{Name: instanceName}).Stage(stage)
										instance = any(instanceLevel)
										__gong__map_Level[identifier] = instanceLevel
									case "Line_detail":
										instanceLine_detail := (&Line_detail{Name: instanceName}).Stage(stage)
										instance = any(instanceLine_detail)
										__gong__map_Line_detail[identifier] = instanceLine_detail
									case "Line_width":
										instanceLine_width := (&Line_width{Name: instanceName}).Stage(stage)
										instance = any(instanceLine_width)
										__gong__map_Line_width[identifier] = instanceLine_width
									case "Link":
										instanceLink := (&Link{Name: instanceName}).Stage(stage)
										instance = any(instanceLink)
										__gong__map_Link[identifier] = instanceLink
									case "Listen":
										instanceListen := (&Listen{Name: instanceName}).Stage(stage)
										instance = any(instanceListen)
										__gong__map_Listen[identifier] = instanceListen
									case "Listening":
										instanceListening := (&Listening{Name: instanceName}).Stage(stage)
										instance = any(instanceListening)
										__gong__map_Listening[identifier] = instanceListening
									case "Lyric":
										instanceLyric := (&Lyric{Name: instanceName}).Stage(stage)
										instance = any(instanceLyric)
										__gong__map_Lyric[identifier] = instanceLyric
									case "Lyric_font":
										instanceLyric_font := (&Lyric_font{Name: instanceName}).Stage(stage)
										instance = any(instanceLyric_font)
										__gong__map_Lyric_font[identifier] = instanceLyric_font
									case "Lyric_language":
										instanceLyric_language := (&Lyric_language{Name: instanceName}).Stage(stage)
										instance = any(instanceLyric_language)
										__gong__map_Lyric_language[identifier] = instanceLyric_language
									case "Measure_layout":
										instanceMeasure_layout := (&Measure_layout{Name: instanceName}).Stage(stage)
										instance = any(instanceMeasure_layout)
										__gong__map_Measure_layout[identifier] = instanceMeasure_layout
									case "Measure_numbering":
										instanceMeasure_numbering := (&Measure_numbering{Name: instanceName}).Stage(stage)
										instance = any(instanceMeasure_numbering)
										__gong__map_Measure_numbering[identifier] = instanceMeasure_numbering
									case "Measure_repeat":
										instanceMeasure_repeat := (&Measure_repeat{Name: instanceName}).Stage(stage)
										instance = any(instanceMeasure_repeat)
										__gong__map_Measure_repeat[identifier] = instanceMeasure_repeat
									case "Measure_style":
										instanceMeasure_style := (&Measure_style{Name: instanceName}).Stage(stage)
										instance = any(instanceMeasure_style)
										__gong__map_Measure_style[identifier] = instanceMeasure_style
									case "Membrane":
										instanceMembrane := (&Membrane{Name: instanceName}).Stage(stage)
										instance = any(instanceMembrane)
										__gong__map_Membrane[identifier] = instanceMembrane
									case "Metal":
										instanceMetal := (&Metal{Name: instanceName}).Stage(stage)
										instance = any(instanceMetal)
										__gong__map_Metal[identifier] = instanceMetal
									case "Metronome":
										instanceMetronome := (&Metronome{Name: instanceName}).Stage(stage)
										instance = any(instanceMetronome)
										__gong__map_Metronome[identifier] = instanceMetronome
									case "Metronome_beam":
										instanceMetronome_beam := (&Metronome_beam{Name: instanceName}).Stage(stage)
										instance = any(instanceMetronome_beam)
										__gong__map_Metronome_beam[identifier] = instanceMetronome_beam
									case "Metronome_note":
										instanceMetronome_note := (&Metronome_note{Name: instanceName}).Stage(stage)
										instance = any(instanceMetronome_note)
										__gong__map_Metronome_note[identifier] = instanceMetronome_note
									case "Metronome_tied":
										instanceMetronome_tied := (&Metronome_tied{Name: instanceName}).Stage(stage)
										instance = any(instanceMetronome_tied)
										__gong__map_Metronome_tied[identifier] = instanceMetronome_tied
									case "Metronome_tuplet":
										instanceMetronome_tuplet := (&Metronome_tuplet{Name: instanceName}).Stage(stage)
										instance = any(instanceMetronome_tuplet)
										__gong__map_Metronome_tuplet[identifier] = instanceMetronome_tuplet
									case "Midi_device":
										instanceMidi_device := (&Midi_device{Name: instanceName}).Stage(stage)
										instance = any(instanceMidi_device)
										__gong__map_Midi_device[identifier] = instanceMidi_device
									case "Midi_instrument":
										instanceMidi_instrument := (&Midi_instrument{Name: instanceName}).Stage(stage)
										instance = any(instanceMidi_instrument)
										__gong__map_Midi_instrument[identifier] = instanceMidi_instrument
									case "Miscellaneous":
										instanceMiscellaneous := (&Miscellaneous{Name: instanceName}).Stage(stage)
										instance = any(instanceMiscellaneous)
										__gong__map_Miscellaneous[identifier] = instanceMiscellaneous
									case "Miscellaneous_field":
										instanceMiscellaneous_field := (&Miscellaneous_field{Name: instanceName}).Stage(stage)
										instance = any(instanceMiscellaneous_field)
										__gong__map_Miscellaneous_field[identifier] = instanceMiscellaneous_field
									case "Mordent":
										instanceMordent := (&Mordent{Name: instanceName}).Stage(stage)
										instance = any(instanceMordent)
										__gong__map_Mordent[identifier] = instanceMordent
									case "Multiple_rest":
										instanceMultiple_rest := (&Multiple_rest{Name: instanceName}).Stage(stage)
										instance = any(instanceMultiple_rest)
										__gong__map_Multiple_rest[identifier] = instanceMultiple_rest
									case "Name_display":
										instanceName_display := (&Name_display{Name: instanceName}).Stage(stage)
										instance = any(instanceName_display)
										__gong__map_Name_display[identifier] = instanceName_display
									case "Non_arpeggiate":
										instanceNon_arpeggiate := (&Non_arpeggiate{Name: instanceName}).Stage(stage)
										instance = any(instanceNon_arpeggiate)
										__gong__map_Non_arpeggiate[identifier] = instanceNon_arpeggiate
									case "Notations":
										instanceNotations := (&Notations{Name: instanceName}).Stage(stage)
										instance = any(instanceNotations)
										__gong__map_Notations[identifier] = instanceNotations
									case "Note":
										instanceNote := (&Note{Name: instanceName}).Stage(stage)
										instance = any(instanceNote)
										__gong__map_Note[identifier] = instanceNote
									case "Note_size":
										instanceNote_size := (&Note_size{Name: instanceName}).Stage(stage)
										instance = any(instanceNote_size)
										__gong__map_Note_size[identifier] = instanceNote_size
									case "Note_type":
										instanceNote_type := (&Note_type{Name: instanceName}).Stage(stage)
										instance = any(instanceNote_type)
										__gong__map_Note_type[identifier] = instanceNote_type
									case "Notehead":
										instanceNotehead := (&Notehead{Name: instanceName}).Stage(stage)
										instance = any(instanceNotehead)
										__gong__map_Notehead[identifier] = instanceNotehead
									case "Notehead_text":
										instanceNotehead_text := (&Notehead_text{Name: instanceName}).Stage(stage)
										instance = any(instanceNotehead_text)
										__gong__map_Notehead_text[identifier] = instanceNotehead_text
									case "Numeral":
										instanceNumeral := (&Numeral{Name: instanceName}).Stage(stage)
										instance = any(instanceNumeral)
										__gong__map_Numeral[identifier] = instanceNumeral
									case "Numeral_key":
										instanceNumeral_key := (&Numeral_key{Name: instanceName}).Stage(stage)
										instance = any(instanceNumeral_key)
										__gong__map_Numeral_key[identifier] = instanceNumeral_key
									case "Numeral_root":
										instanceNumeral_root := (&Numeral_root{Name: instanceName}).Stage(stage)
										instance = any(instanceNumeral_root)
										__gong__map_Numeral_root[identifier] = instanceNumeral_root
									case "Octave_shift":
										instanceOctave_shift := (&Octave_shift{Name: instanceName}).Stage(stage)
										instance = any(instanceOctave_shift)
										__gong__map_Octave_shift[identifier] = instanceOctave_shift
									case "Offset":
										instanceOffset := (&Offset{Name: instanceName}).Stage(stage)
										instance = any(instanceOffset)
										__gong__map_Offset[identifier] = instanceOffset
									case "Opus":
										instanceOpus := (&Opus{Name: instanceName}).Stage(stage)
										instance = any(instanceOpus)
										__gong__map_Opus[identifier] = instanceOpus
									case "Ornaments":
										instanceOrnaments := (&Ornaments{Name: instanceName}).Stage(stage)
										instance = any(instanceOrnaments)
										__gong__map_Ornaments[identifier] = instanceOrnaments
									case "Other_appearance":
										instanceOther_appearance := (&Other_appearance{Name: instanceName}).Stage(stage)
										instance = any(instanceOther_appearance)
										__gong__map_Other_appearance[identifier] = instanceOther_appearance
									case "Other_listening":
										instanceOther_listening := (&Other_listening{Name: instanceName}).Stage(stage)
										instance = any(instanceOther_listening)
										__gong__map_Other_listening[identifier] = instanceOther_listening
									case "Other_notation":
										instanceOther_notation := (&Other_notation{Name: instanceName}).Stage(stage)
										instance = any(instanceOther_notation)
										__gong__map_Other_notation[identifier] = instanceOther_notation
									case "Other_play":
										instanceOther_play := (&Other_play{Name: instanceName}).Stage(stage)
										instance = any(instanceOther_play)
										__gong__map_Other_play[identifier] = instanceOther_play
									case "Page_layout":
										instancePage_layout := (&Page_layout{Name: instanceName}).Stage(stage)
										instance = any(instancePage_layout)
										__gong__map_Page_layout[identifier] = instancePage_layout
									case "Page_margins":
										instancePage_margins := (&Page_margins{Name: instanceName}).Stage(stage)
										instance = any(instancePage_margins)
										__gong__map_Page_margins[identifier] = instancePage_margins
									case "Part_clef":
										instancePart_clef := (&Part_clef{Name: instanceName}).Stage(stage)
										instance = any(instancePart_clef)
										__gong__map_Part_clef[identifier] = instancePart_clef
									case "Part_group":
										instancePart_group := (&Part_group{Name: instanceName}).Stage(stage)
										instance = any(instancePart_group)
										__gong__map_Part_group[identifier] = instancePart_group
									case "Part_link":
										instancePart_link := (&Part_link{Name: instanceName}).Stage(stage)
										instance = any(instancePart_link)
										__gong__map_Part_link[identifier] = instancePart_link
									case "Part_list":
										instancePart_list := (&Part_list{Name: instanceName}).Stage(stage)
										instance = any(instancePart_list)
										__gong__map_Part_list[identifier] = instancePart_list
									case "Part_symbol":
										instancePart_symbol := (&Part_symbol{Name: instanceName}).Stage(stage)
										instance = any(instancePart_symbol)
										__gong__map_Part_symbol[identifier] = instancePart_symbol
									case "Part_transpose":
										instancePart_transpose := (&Part_transpose{Name: instanceName}).Stage(stage)
										instance = any(instancePart_transpose)
										__gong__map_Part_transpose[identifier] = instancePart_transpose
									case "Pedal":
										instancePedal := (&Pedal{Name: instanceName}).Stage(stage)
										instance = any(instancePedal)
										__gong__map_Pedal[identifier] = instancePedal
									case "Pedal_tuning":
										instancePedal_tuning := (&Pedal_tuning{Name: instanceName}).Stage(stage)
										instance = any(instancePedal_tuning)
										__gong__map_Pedal_tuning[identifier] = instancePedal_tuning
									case "Percussion":
										instancePercussion := (&Percussion{Name: instanceName}).Stage(stage)
										instance = any(instancePercussion)
										__gong__map_Percussion[identifier] = instancePercussion
									case "Pitch":
										instancePitch := (&Pitch{Name: instanceName}).Stage(stage)
										instance = any(instancePitch)
										__gong__map_Pitch[identifier] = instancePitch
									case "Pitched":
										instancePitched := (&Pitched{Name: instanceName}).Stage(stage)
										instance = any(instancePitched)
										__gong__map_Pitched[identifier] = instancePitched
									case "Play":
										instancePlay := (&Play{Name: instanceName}).Stage(stage)
										instance = any(instancePlay)
										__gong__map_Play[identifier] = instancePlay
									case "Player":
										instancePlayer := (&Player{Name: instanceName}).Stage(stage)
										instance = any(instancePlayer)
										__gong__map_Player[identifier] = instancePlayer
									case "Principal_voice":
										instancePrincipal_voice := (&Principal_voice{Name: instanceName}).Stage(stage)
										instance = any(instancePrincipal_voice)
										__gong__map_Principal_voice[identifier] = instancePrincipal_voice
									case "Print":
										instancePrint := (&Print{Name: instanceName}).Stage(stage)
										instance = any(instancePrint)
										__gong__map_Print[identifier] = instancePrint
									case "Release":
										instanceRelease := (&Release{Name: instanceName}).Stage(stage)
										instance = any(instanceRelease)
										__gong__map_Release[identifier] = instanceRelease
									case "Repeat":
										instanceRepeat := (&Repeat{Name: instanceName}).Stage(stage)
										instance = any(instanceRepeat)
										__gong__map_Repeat[identifier] = instanceRepeat
									case "Rest":
										instanceRest := (&Rest{Name: instanceName}).Stage(stage)
										instance = any(instanceRest)
										__gong__map_Rest[identifier] = instanceRest
									case "Root":
										instanceRoot := (&Root{Name: instanceName}).Stage(stage)
										instance = any(instanceRoot)
										__gong__map_Root[identifier] = instanceRoot
									case "Root_step":
										instanceRoot_step := (&Root_step{Name: instanceName}).Stage(stage)
										instance = any(instanceRoot_step)
										__gong__map_Root_step[identifier] = instanceRoot_step
									case "Scaling":
										instanceScaling := (&Scaling{Name: instanceName}).Stage(stage)
										instance = any(instanceScaling)
										__gong__map_Scaling[identifier] = instanceScaling
									case "Scordatura":
										instanceScordatura := (&Scordatura{Name: instanceName}).Stage(stage)
										instance = any(instanceScordatura)
										__gong__map_Scordatura[identifier] = instanceScordatura
									case "Score_instrument":
										instanceScore_instrument := (&Score_instrument{Name: instanceName}).Stage(stage)
										instance = any(instanceScore_instrument)
										__gong__map_Score_instrument[identifier] = instanceScore_instrument
									case "Score_part":
										instanceScore_part := (&Score_part{Name: instanceName}).Stage(stage)
										instance = any(instanceScore_part)
										__gong__map_Score_part[identifier] = instanceScore_part
									case "Score_partwise":
										instanceScore_partwise := (&Score_partwise{Name: instanceName}).Stage(stage)
										instance = any(instanceScore_partwise)
										__gong__map_Score_partwise[identifier] = instanceScore_partwise
									case "Score_timewise":
										instanceScore_timewise := (&Score_timewise{Name: instanceName}).Stage(stage)
										instance = any(instanceScore_timewise)
										__gong__map_Score_timewise[identifier] = instanceScore_timewise
									case "Segno":
										instanceSegno := (&Segno{Name: instanceName}).Stage(stage)
										instance = any(instanceSegno)
										__gong__map_Segno[identifier] = instanceSegno
									case "Slash":
										instanceSlash := (&Slash{Name: instanceName}).Stage(stage)
										instance = any(instanceSlash)
										__gong__map_Slash[identifier] = instanceSlash
									case "Slide":
										instanceSlide := (&Slide{Name: instanceName}).Stage(stage)
										instance = any(instanceSlide)
										__gong__map_Slide[identifier] = instanceSlide
									case "Slur":
										instanceSlur := (&Slur{Name: instanceName}).Stage(stage)
										instance = any(instanceSlur)
										__gong__map_Slur[identifier] = instanceSlur
									case "Sound":
										instanceSound := (&Sound{Name: instanceName}).Stage(stage)
										instance = any(instanceSound)
										__gong__map_Sound[identifier] = instanceSound
									case "Staff_details":
										instanceStaff_details := (&Staff_details{Name: instanceName}).Stage(stage)
										instance = any(instanceStaff_details)
										__gong__map_Staff_details[identifier] = instanceStaff_details
									case "Staff_divide":
										instanceStaff_divide := (&Staff_divide{Name: instanceName}).Stage(stage)
										instance = any(instanceStaff_divide)
										__gong__map_Staff_divide[identifier] = instanceStaff_divide
									case "Staff_layout":
										instanceStaff_layout := (&Staff_layout{Name: instanceName}).Stage(stage)
										instance = any(instanceStaff_layout)
										__gong__map_Staff_layout[identifier] = instanceStaff_layout
									case "Staff_size":
										instanceStaff_size := (&Staff_size{Name: instanceName}).Stage(stage)
										instance = any(instanceStaff_size)
										__gong__map_Staff_size[identifier] = instanceStaff_size
									case "Staff_tuning":
										instanceStaff_tuning := (&Staff_tuning{Name: instanceName}).Stage(stage)
										instance = any(instanceStaff_tuning)
										__gong__map_Staff_tuning[identifier] = instanceStaff_tuning
									case "Stem":
										instanceStem := (&Stem{Name: instanceName}).Stage(stage)
										instance = any(instanceStem)
										__gong__map_Stem[identifier] = instanceStem
									case "Stick":
										instanceStick := (&Stick{Name: instanceName}).Stage(stage)
										instance = any(instanceStick)
										__gong__map_Stick[identifier] = instanceStick
									case "String_mute":
										instanceString_mute := (&String_mute{Name: instanceName}).Stage(stage)
										instance = any(instanceString_mute)
										__gong__map_String_mute[identifier] = instanceString_mute
									case "Strong_accent":
										instanceStrong_accent := (&Strong_accent{Name: instanceName}).Stage(stage)
										instance = any(instanceStrong_accent)
										__gong__map_Strong_accent[identifier] = instanceStrong_accent
									case "Supports":
										instanceSupports := (&Supports{Name: instanceName}).Stage(stage)
										instance = any(instanceSupports)
										__gong__map_Supports[identifier] = instanceSupports
									case "Swing":
										instanceSwing := (&Swing{Name: instanceName}).Stage(stage)
										instance = any(instanceSwing)
										__gong__map_Swing[identifier] = instanceSwing
									case "Sync":
										instanceSync := (&Sync{Name: instanceName}).Stage(stage)
										instance = any(instanceSync)
										__gong__map_Sync[identifier] = instanceSync
									case "System_dividers":
										instanceSystem_dividers := (&System_dividers{Name: instanceName}).Stage(stage)
										instance = any(instanceSystem_dividers)
										__gong__map_System_dividers[identifier] = instanceSystem_dividers
									case "System_layout":
										instanceSystem_layout := (&System_layout{Name: instanceName}).Stage(stage)
										instance = any(instanceSystem_layout)
										__gong__map_System_layout[identifier] = instanceSystem_layout
									case "System_margins":
										instanceSystem_margins := (&System_margins{Name: instanceName}).Stage(stage)
										instance = any(instanceSystem_margins)
										__gong__map_System_margins[identifier] = instanceSystem_margins
									case "Tap":
										instanceTap := (&Tap{Name: instanceName}).Stage(stage)
										instance = any(instanceTap)
										__gong__map_Tap[identifier] = instanceTap
									case "Technical":
										instanceTechnical := (&Technical{Name: instanceName}).Stage(stage)
										instance = any(instanceTechnical)
										__gong__map_Technical[identifier] = instanceTechnical
									case "Text_element_data":
										instanceText_element_data := (&Text_element_data{Name: instanceName}).Stage(stage)
										instance = any(instanceText_element_data)
										__gong__map_Text_element_data[identifier] = instanceText_element_data
									case "Tie":
										instanceTie := (&Tie{Name: instanceName}).Stage(stage)
										instance = any(instanceTie)
										__gong__map_Tie[identifier] = instanceTie
									case "Tied":
										instanceTied := (&Tied{Name: instanceName}).Stage(stage)
										instance = any(instanceTied)
										__gong__map_Tied[identifier] = instanceTied
									case "Time":
										instanceTime := (&Time{Name: instanceName}).Stage(stage)
										instance = any(instanceTime)
										__gong__map_Time[identifier] = instanceTime
									case "Time_modification":
										instanceTime_modification := (&Time_modification{Name: instanceName}).Stage(stage)
										instance = any(instanceTime_modification)
										__gong__map_Time_modification[identifier] = instanceTime_modification
									case "Timpani":
										instanceTimpani := (&Timpani{Name: instanceName}).Stage(stage)
										instance = any(instanceTimpani)
										__gong__map_Timpani[identifier] = instanceTimpani
									case "Transpose":
										instanceTranspose := (&Transpose{Name: instanceName}).Stage(stage)
										instance = any(instanceTranspose)
										__gong__map_Transpose[identifier] = instanceTranspose
									case "Tremolo":
										instanceTremolo := (&Tremolo{Name: instanceName}).Stage(stage)
										instance = any(instanceTremolo)
										__gong__map_Tremolo[identifier] = instanceTremolo
									case "Tuplet":
										instanceTuplet := (&Tuplet{Name: instanceName}).Stage(stage)
										instance = any(instanceTuplet)
										__gong__map_Tuplet[identifier] = instanceTuplet
									case "Tuplet_dot":
										instanceTuplet_dot := (&Tuplet_dot{Name: instanceName}).Stage(stage)
										instance = any(instanceTuplet_dot)
										__gong__map_Tuplet_dot[identifier] = instanceTuplet_dot
									case "Tuplet_number":
										instanceTuplet_number := (&Tuplet_number{Name: instanceName}).Stage(stage)
										instance = any(instanceTuplet_number)
										__gong__map_Tuplet_number[identifier] = instanceTuplet_number
									case "Tuplet_portion":
										instanceTuplet_portion := (&Tuplet_portion{Name: instanceName}).Stage(stage)
										instance = any(instanceTuplet_portion)
										__gong__map_Tuplet_portion[identifier] = instanceTuplet_portion
									case "Tuplet_type":
										instanceTuplet_type := (&Tuplet_type{Name: instanceName}).Stage(stage)
										instance = any(instanceTuplet_type)
										__gong__map_Tuplet_type[identifier] = instanceTuplet_type
									case "Typed_text":
										instanceTyped_text := (&Typed_text{Name: instanceName}).Stage(stage)
										instance = any(instanceTyped_text)
										__gong__map_Typed_text[identifier] = instanceTyped_text
									case "Unpitched":
										instanceUnpitched := (&Unpitched{Name: instanceName}).Stage(stage)
										instance = any(instanceUnpitched)
										__gong__map_Unpitched[identifier] = instanceUnpitched
									case "Virtual_instrument":
										instanceVirtual_instrument := (&Virtual_instrument{Name: instanceName}).Stage(stage)
										instance = any(instanceVirtual_instrument)
										__gong__map_Virtual_instrument[identifier] = instanceVirtual_instrument
									case "Wait":
										instanceWait := (&Wait{Name: instanceName}).Stage(stage)
										instance = any(instanceWait)
										__gong__map_Wait[identifier] = instanceWait
									case "Wavy_line":
										instanceWavy_line := (&Wavy_line{Name: instanceName}).Stage(stage)
										instance = any(instanceWavy_line)
										__gong__map_Wavy_line[identifier] = instanceWavy_line
									case "Wedge":
										instanceWedge := (&Wedge{Name: instanceName}).Stage(stage)
										instance = any(instanceWedge)
										__gong__map_Wedge[identifier] = instanceWedge
									case "Wood":
										instanceWood := (&Wood{Name: instanceName}).Stage(stage)
										instance = any(instanceWood)
										__gong__map_Wood[identifier] = instanceWood
									case "Work":
										instanceWork := (&Work{Name: instanceName}).Stage(stage)
										instance = any(instanceWork)
										__gong__map_Work[identifier] = instanceWork
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "Accidental":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Accidental_mark":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Accidental_text":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Accord":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Accordion_registration":
							switch fieldName {
							// insertion point for date assign code
							}
						case "AnyType":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Appearance":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Arpeggiate":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Arrow":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Articulations":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Assess":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Attributes":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Backup":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Bar_style_color":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Barline":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Barre":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Bass":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Bass_step":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Beam":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Beat_repeat":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Beat_unit_tied":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Beater":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Bend":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Bookmark":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Bracket":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Breath_mark":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Caesura":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Cancel":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Clef":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Coda":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Credit":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Dashes":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Defaults":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Degree":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Degree_alter":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Degree_type":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Degree_value":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Direction":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Direction_type":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Distance":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Double":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Dynamics":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Effect":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Elision":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty_font":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty_line":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty_placement":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty_placement_smufl":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty_print_object_style_align":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty_print_style":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty_print_style_align":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty_print_style_align_id":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Empty_trill_sound":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Encoding":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Ending":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Extend":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Feature":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Fermata":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Figure":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Figured_bass":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Fingering":
							switch fieldName {
							// insertion point for date assign code
							}
						case "First_fret":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Foo":
							switch fieldName {
							// insertion point for date assign code
							}
						case "For_part":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Formatted_symbol":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Formatted_symbol_id":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Forward":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Frame":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Frame_note":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Fret":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Glass":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Glissando":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Glyph":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Grace":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Group_barline":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Group_symbol":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Grouping":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Hammer_on_pull_off":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Handbell":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Harmon_closed":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Harmon_mute":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Harmonic":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Harmony":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Harmony_alter":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Harp_pedals":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Heel_toe":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Hole":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Hole_closed":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Horizontal_turn":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Identification":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Image":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Instrument":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Instrument_change":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Instrument_link":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Interchangeable":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Inversion":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Key":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Key_accidental":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Key_octave":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Kind":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Level":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Line_detail":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Line_width":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Link":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Listen":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Listening":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Lyric":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Lyric_font":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Lyric_language":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Measure_layout":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Measure_numbering":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Measure_repeat":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Measure_style":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Membrane":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Metal":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Metronome":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Metronome_beam":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Metronome_note":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Metronome_tied":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Metronome_tuplet":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Midi_device":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Midi_instrument":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Miscellaneous":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Miscellaneous_field":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Mordent":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Multiple_rest":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Name_display":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Non_arpeggiate":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Notations":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Note":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Note_size":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Note_type":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Notehead":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Notehead_text":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Numeral":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Numeral_key":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Numeral_root":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Octave_shift":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Offset":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Opus":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Ornaments":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Other_appearance":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Other_listening":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Other_notation":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Other_play":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Page_layout":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Page_margins":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Part_clef":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Part_group":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Part_link":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Part_list":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Part_symbol":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Part_transpose":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Pedal":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Pedal_tuning":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Percussion":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Pitch":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Pitched":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Play":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Player":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Principal_voice":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Print":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Release":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Repeat":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Rest":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Root":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Root_step":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Scaling":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Scordatura":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Score_instrument":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Score_part":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Score_partwise":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Score_timewise":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Segno":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Slash":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Slide":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Slur":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Sound":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Staff_details":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Staff_divide":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Staff_layout":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Staff_size":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Staff_tuning":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Stem":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Stick":
							switch fieldName {
							// insertion point for date assign code
							}
						case "String_mute":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Strong_accent":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Supports":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Swing":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Sync":
							switch fieldName {
							// insertion point for date assign code
							}
						case "System_dividers":
							switch fieldName {
							// insertion point for date assign code
							}
						case "System_layout":
							switch fieldName {
							// insertion point for date assign code
							}
						case "System_margins":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tap":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Technical":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Text_element_data":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tie":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tied":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Time":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Time_modification":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Timpani":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Transpose":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tremolo":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tuplet":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tuplet_dot":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tuplet_number":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tuplet_portion":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tuplet_type":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Typed_text":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Unpitched":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Virtual_instrument":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Wait":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Wavy_line":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Wedge":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Wood":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Work":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "Accidental":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Accidental_mark":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Accidental_text":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Accord":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Accordion_registration":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "AnyType":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Appearance":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Line_width":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Line_width[targetIdentifier]
							__gong__map_Appearance[identifier].Line_width =
								append(__gong__map_Appearance[identifier].Line_width, target)
						case "Note_size":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Note_size[targetIdentifier]
							__gong__map_Appearance[identifier].Note_size =
								append(__gong__map_Appearance[identifier].Note_size, target)
						case "Distance":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Distance[targetIdentifier]
							__gong__map_Appearance[identifier].Distance =
								append(__gong__map_Appearance[identifier].Distance, target)
						case "Glyph":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Glyph[targetIdentifier]
							__gong__map_Appearance[identifier].Glyph =
								append(__gong__map_Appearance[identifier].Glyph, target)
						case "Other_appearance":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Other_appearance[targetIdentifier]
							__gong__map_Appearance[identifier].Other_appearance =
								append(__gong__map_Appearance[identifier].Other_appearance, target)
						}
					case "Arpeggiate":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Arrow":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Articulations":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Assess":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Attributes":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Key":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Key[targetIdentifier]
							__gong__map_Attributes[identifier].Key =
								append(__gong__map_Attributes[identifier].Key, target)
						case "Clef":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Clef[targetIdentifier]
							__gong__map_Attributes[identifier].Clef =
								append(__gong__map_Attributes[identifier].Clef, target)
						case "Staff_details":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Staff_details[targetIdentifier]
							__gong__map_Attributes[identifier].Staff_details =
								append(__gong__map_Attributes[identifier].Staff_details, target)
						case "Measure_style":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Measure_style[targetIdentifier]
							__gong__map_Attributes[identifier].Measure_style =
								append(__gong__map_Attributes[identifier].Measure_style, target)
						case "Transpose":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Transpose[targetIdentifier]
							__gong__map_Attributes[identifier].Transpose =
								append(__gong__map_Attributes[identifier].Transpose, target)
						case "For_part":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_For_part[targetIdentifier]
							__gong__map_Attributes[identifier].For_part =
								append(__gong__map_Attributes[identifier].For_part, target)
						}
					case "Backup":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Bar_style_color":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Barline":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Barre":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Bass":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Bass_step":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Beam":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Beat_repeat":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Beat_unit_tied":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Beater":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Bend":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Bookmark":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Bracket":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Breath_mark":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Caesura":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Cancel":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Clef":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Coda":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Credit":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Link":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Link[targetIdentifier]
							__gong__map_Credit[identifier].Link =
								append(__gong__map_Credit[identifier].Link, target)
						case "Bookmark":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Bookmark[targetIdentifier]
							__gong__map_Credit[identifier].Bookmark =
								append(__gong__map_Credit[identifier].Bookmark, target)
						}
					case "Dashes":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Defaults":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Lyric_font":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Lyric_font[targetIdentifier]
							__gong__map_Defaults[identifier].Lyric_font =
								append(__gong__map_Defaults[identifier].Lyric_font, target)
						case "Lyric_language":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Lyric_language[targetIdentifier]
							__gong__map_Defaults[identifier].Lyric_language =
								append(__gong__map_Defaults[identifier].Lyric_language, target)
						}
					case "Degree":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Degree_alter":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Degree_type":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Degree_value":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Direction":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Direction_type":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Direction_type[targetIdentifier]
							__gong__map_Direction[identifier].Direction_type =
								append(__gong__map_Direction[identifier].Direction_type, target)
						}
					case "Direction_type":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Segno":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Segno[targetIdentifier]
							__gong__map_Direction_type[identifier].Segno =
								append(__gong__map_Direction_type[identifier].Segno, target)
						case "Coda":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Coda[targetIdentifier]
							__gong__map_Direction_type[identifier].Coda =
								append(__gong__map_Direction_type[identifier].Coda, target)
						case "Dynamics":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Dynamics[targetIdentifier]
							__gong__map_Direction_type[identifier].Dynamics =
								append(__gong__map_Direction_type[identifier].Dynamics, target)
						case "Percussion":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Percussion[targetIdentifier]
							__gong__map_Direction_type[identifier].Percussion =
								append(__gong__map_Direction_type[identifier].Percussion, target)
						}
					case "Distance":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Double":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Dynamics":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Effect":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Elision":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty_font":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty_line":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty_placement":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty_placement_smufl":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty_print_object_style_align":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty_print_style":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty_print_style_align":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty_print_style_align_id":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Empty_trill_sound":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Encoding":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Ending":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Extend":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Feature":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Fermata":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Figure":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Figured_bass":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Figure":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Figure[targetIdentifier]
							__gong__map_Figured_bass[identifier].Figure =
								append(__gong__map_Figured_bass[identifier].Figure, target)
						}
					case "Fingering":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "First_fret":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Foo":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "For_part":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Formatted_symbol":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Formatted_symbol_id":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Forward":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Frame":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Frame_note":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Frame_note[targetIdentifier]
							__gong__map_Frame[identifier].Frame_note =
								append(__gong__map_Frame[identifier].Frame_note, target)
						}
					case "Frame_note":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Fret":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Glass":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Glissando":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Glyph":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Grace":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Group_barline":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Group_symbol":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Grouping":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Feature":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Feature[targetIdentifier]
							__gong__map_Grouping[identifier].Feature =
								append(__gong__map_Grouping[identifier].Feature, target)
						}
					case "Hammer_on_pull_off":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Handbell":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Harmon_closed":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Harmon_mute":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Harmonic":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Harmony":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Harmony_alter":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Harp_pedals":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Pedal_tuning":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Pedal_tuning[targetIdentifier]
							__gong__map_Harp_pedals[identifier].Pedal_tuning =
								append(__gong__map_Harp_pedals[identifier].Pedal_tuning, target)
						}
					case "Heel_toe":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Hole":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Hole_closed":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Horizontal_turn":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Identification":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Creator":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Typed_text[targetIdentifier]
							__gong__map_Identification[identifier].Creator =
								append(__gong__map_Identification[identifier].Creator, target)
						case "Rights":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Typed_text[targetIdentifier]
							__gong__map_Identification[identifier].Rights =
								append(__gong__map_Identification[identifier].Rights, target)
						case "Relation":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Typed_text[targetIdentifier]
							__gong__map_Identification[identifier].Relation =
								append(__gong__map_Identification[identifier].Relation, target)
						}
					case "Image":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Instrument":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Instrument_change":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Instrument_link":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Interchangeable":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Inversion":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Key":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Key_octave":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Key_octave[targetIdentifier]
							__gong__map_Key[identifier].Key_octave =
								append(__gong__map_Key[identifier].Key_octave, target)
						}
					case "Key_accidental":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Key_octave":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Kind":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Level":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Line_detail":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Line_width":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Link":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Listen":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Listening":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Lyric":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Lyric_font":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Lyric_language":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Measure_layout":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Measure_numbering":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Measure_repeat":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Measure_style":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Membrane":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Metal":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Metronome":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Metronome_beam":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Metronome_note":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Metronome_dot":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Empty[targetIdentifier]
							__gong__map_Metronome_note[identifier].Metronome_dot =
								append(__gong__map_Metronome_note[identifier].Metronome_dot, target)
						case "Metronome_beam":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Metronome_beam[targetIdentifier]
							__gong__map_Metronome_note[identifier].Metronome_beam =
								append(__gong__map_Metronome_note[identifier].Metronome_beam, target)
						}
					case "Metronome_tied":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Metronome_tuplet":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Midi_device":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Midi_instrument":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Miscellaneous":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Miscellaneous_field":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Miscellaneous_field[targetIdentifier]
							__gong__map_Miscellaneous[identifier].Miscellaneous_field =
								append(__gong__map_Miscellaneous[identifier].Miscellaneous_field, target)
						}
					case "Miscellaneous_field":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Mordent":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Multiple_rest":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Name_display":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Non_arpeggiate":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Notations":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Note":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Instrument":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Instrument[targetIdentifier]
							__gong__map_Note[identifier].Instrument =
								append(__gong__map_Note[identifier].Instrument, target)
						case "Dot":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Empty_placement[targetIdentifier]
							__gong__map_Note[identifier].Dot =
								append(__gong__map_Note[identifier].Dot, target)
						case "Notations":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Notations[targetIdentifier]
							__gong__map_Note[identifier].Notations =
								append(__gong__map_Note[identifier].Notations, target)
						case "Lyric":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Lyric[targetIdentifier]
							__gong__map_Note[identifier].Lyric =
								append(__gong__map_Note[identifier].Lyric, target)
						}
					case "Note_size":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Note_type":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Notehead":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Notehead_text":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Numeral":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Numeral_key":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Numeral_root":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Octave_shift":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Offset":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Opus":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Ornaments":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Accidental_mark":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Accidental_mark[targetIdentifier]
							__gong__map_Ornaments[identifier].Accidental_mark =
								append(__gong__map_Ornaments[identifier].Accidental_mark, target)
						}
					case "Other_appearance":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Other_listening":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Other_notation":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Other_play":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Page_layout":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Page_margins":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Part_clef":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Part_group":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Part_link":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Instrument_link":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Instrument_link[targetIdentifier]
							__gong__map_Part_link[identifier].Instrument_link =
								append(__gong__map_Part_link[identifier].Instrument_link, target)
						}
					case "Part_list":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Part_symbol":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Part_transpose":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Pedal":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Pedal_tuning":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Percussion":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Pitch":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Pitched":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Play":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Player":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Principal_voice":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Print":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Release":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Repeat":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Rest":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Root":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Root_step":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Scaling":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Scordatura":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Accord":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Accord[targetIdentifier]
							__gong__map_Scordatura[identifier].Accord =
								append(__gong__map_Scordatura[identifier].Accord, target)
						}
					case "Score_instrument":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Score_part":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Part_link":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Part_link[targetIdentifier]
							__gong__map_Score_part[identifier].Part_link =
								append(__gong__map_Score_part[identifier].Part_link, target)
						case "Score_instrument":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Score_instrument[targetIdentifier]
							__gong__map_Score_part[identifier].Score_instrument =
								append(__gong__map_Score_part[identifier].Score_instrument, target)
						case "Player":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Player[targetIdentifier]
							__gong__map_Score_part[identifier].Player =
								append(__gong__map_Score_part[identifier].Player, target)
						}
					case "Score_partwise":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Score_timewise":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Segno":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Slash":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Slide":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Slur":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Sound":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Staff_details":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Staff_tuning":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Staff_tuning[targetIdentifier]
							__gong__map_Staff_details[identifier].Staff_tuning =
								append(__gong__map_Staff_details[identifier].Staff_tuning, target)
						}
					case "Staff_divide":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Staff_layout":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Staff_size":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Staff_tuning":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Stem":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Stick":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "String_mute":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Strong_accent":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Supports":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Swing":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Sync":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "System_dividers":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "System_layout":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "System_margins":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tap":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Technical":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Text_element_data":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tie":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tied":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Time":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Time_modification":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Timpani":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Transpose":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tremolo":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tuplet":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tuplet_dot":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tuplet_number":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tuplet_portion":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Tuplet_dot":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Tuplet_dot[targetIdentifier]
							__gong__map_Tuplet_portion[identifier].Tuplet_dot =
								append(__gong__map_Tuplet_portion[identifier].Tuplet_dot, target)
						}
					case "Tuplet_type":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Typed_text":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Unpitched":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Virtual_instrument":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Wait":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Wavy_line":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Wedge":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Wood":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Work":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used

			if bl, ok := expr.(*ast.BasicLit); ok {
				// expression is  for instance ... = 18.000
				basicLit = bl
			} else if ue, ok := expr.(*ast.UnaryExpr); ok {
				// expression is  for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				basicLit = ue.X.(*ast.BasicLit)
				exprSign = -1
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "Accidental":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Accidental[identifier].Name = fielValue
				}
			case "Accidental_mark":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Accidental_mark[identifier].Name = fielValue
				}
			case "Accidental_text":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Accidental_text[identifier].Name = fielValue
				}
			case "Accord":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Accord[identifier].Name = fielValue
				}
			case "Accordion_registration":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Accordion_registration[identifier].Name = fielValue
				}
			case "AnyType":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AnyType[identifier].Name = fielValue
				case "InnerXML":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AnyType[identifier].InnerXML = fielValue
				}
			case "Appearance":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Appearance[identifier].Name = fielValue
				}
			case "Arpeggiate":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Arpeggiate[identifier].Name = fielValue
				}
			case "Arrow":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Arrow[identifier].Name = fielValue
				}
			case "Articulations":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Articulations[identifier].Name = fielValue
				}
			case "Assess":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Assess[identifier].Name = fielValue
				}
			case "Attributes":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attributes[identifier].Name = fielValue
				}
			case "Backup":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Backup[identifier].Name = fielValue
				}
			case "Bar_style_color":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bar_style_color[identifier].Name = fielValue
				}
			case "Barline":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Barline[identifier].Name = fielValue
				case "Segno":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Barline[identifier].Segno = fielValue
				case "Coda":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Barline[identifier].Coda = fielValue
				}
			case "Barre":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Barre[identifier].Name = fielValue
				}
			case "Bass":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bass[identifier].Name = fielValue
				}
			case "Bass_step":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bass_step[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bass_step[identifier].Text = fielValue
				}
			case "Beam":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Beam[identifier].Name = fielValue
				}
			case "Beat_repeat":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Beat_repeat[identifier].Name = fielValue
				}
			case "Beat_unit_tied":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Beat_unit_tied[identifier].Name = fielValue
				}
			case "Beater":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Beater[identifier].Name = fielValue
				}
			case "Bend":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bend[identifier].Name = fielValue
				}
			case "Bookmark":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bookmark[identifier].Name = fielValue
				}
			case "Bracket":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bracket[identifier].Name = fielValue
				}
			case "Breath_mark":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Breath_mark[identifier].Name = fielValue
				}
			case "Caesura":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Caesura[identifier].Name = fielValue
				}
			case "Cancel":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Cancel[identifier].Name = fielValue
				}
			case "Clef":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Clef[identifier].Name = fielValue
				}
			case "Coda":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Coda[identifier].Name = fielValue
				}
			case "Credit":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Credit[identifier].Name = fielValue
				}
			case "Dashes":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Dashes[identifier].Name = fielValue
				}
			case "Defaults":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Defaults[identifier].Name = fielValue
				}
			case "Degree":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Degree[identifier].Name = fielValue
				}
			case "Degree_alter":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Degree_alter[identifier].Name = fielValue
				}
			case "Degree_type":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Degree_type[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Degree_type[identifier].Text = fielValue
				}
			case "Degree_value":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Degree_value[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Degree_value[identifier].Text = fielValue
				}
			case "Direction":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Direction[identifier].Name = fielValue
				}
			case "Direction_type":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Direction_type[identifier].Name = fielValue
				}
			case "Distance":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Distance[identifier].Name = fielValue
				}
			case "Double":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Double[identifier].Name = fielValue
				}
			case "Dynamics":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Dynamics[identifier].Name = fielValue
				}
			case "Effect":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Effect[identifier].Name = fielValue
				}
			case "Elision":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Elision[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Elision[identifier].Value = fielValue
				}
			case "Empty":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty[identifier].Name = fielValue
				}
			case "Empty_font":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty_font[identifier].Name = fielValue
				}
			case "Empty_line":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty_line[identifier].Name = fielValue
				}
			case "Empty_placement":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty_placement[identifier].Name = fielValue
				}
			case "Empty_placement_smufl":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty_placement_smufl[identifier].Name = fielValue
				}
			case "Empty_print_object_style_align":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty_print_object_style_align[identifier].Name = fielValue
				}
			case "Empty_print_style":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty_print_style[identifier].Name = fielValue
				}
			case "Empty_print_style_align":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty_print_style_align[identifier].Name = fielValue
				}
			case "Empty_print_style_align_id":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty_print_style_align_id[identifier].Name = fielValue
				}
			case "Empty_trill_sound":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Empty_trill_sound[identifier].Name = fielValue
				}
			case "Encoding":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Encoding[identifier].Name = fielValue
				case "Software":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Encoding[identifier].Software = fielValue
				case "Encoding_description":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Encoding[identifier].Encoding_description = fielValue
				}
			case "Ending":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ending[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ending[identifier].Value = fielValue
				}
			case "Extend":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Extend[identifier].Name = fielValue
				}
			case "Feature":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Feature[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Feature[identifier].Value = fielValue
				case "Type":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Feature[identifier].Type = fielValue
				}
			case "Fermata":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Fermata[identifier].Name = fielValue
				}
			case "Figure":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Figure[identifier].Name = fielValue
				}
			case "Figured_bass":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Figured_bass[identifier].Name = fielValue
				}
			case "Fingering":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Fingering[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Fingering[identifier].Value = fielValue
				}
			case "First_fret":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_First_fret[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_First_fret[identifier].Text = fielValue
				}
			case "Foo":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Foo[identifier].Name = fielValue
				}
			case "For_part":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_For_part[identifier].Name = fielValue
				}
			case "Formatted_symbol":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Formatted_symbol[identifier].Name = fielValue
				}
			case "Formatted_symbol_id":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Formatted_symbol_id[identifier].Name = fielValue
				}
			case "Forward":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Forward[identifier].Name = fielValue
				}
			case "Frame":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Frame[identifier].Name = fielValue
				case "Unplayed":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Frame[identifier].Unplayed = fielValue
				}
			case "Frame_note":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Frame_note[identifier].Name = fielValue
				case "Astring":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Frame_note[identifier].Astring = fielValue
				}
			case "Fret":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Fret[identifier].Name = fielValue
				}
			case "Glass":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Glass[identifier].Name = fielValue
				}
			case "Glissando":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Glissando[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Glissando[identifier].Value = fielValue
				}
			case "Glyph":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Glyph[identifier].Name = fielValue
				}
			case "Grace":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Grace[identifier].Name = fielValue
				}
			case "Group_barline":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group_barline[identifier].Name = fielValue
				}
			case "Group_symbol":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group_symbol[identifier].Name = fielValue
				}
			case "Grouping":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Grouping[identifier].Name = fielValue
				case "Number":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Grouping[identifier].Number = fielValue
				case "Member_of":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Grouping[identifier].Member_of = fielValue
				}
			case "Hammer_on_pull_off":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Hammer_on_pull_off[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Hammer_on_pull_off[identifier].Value = fielValue
				}
			case "Handbell":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Handbell[identifier].Name = fielValue
				}
			case "Harmon_closed":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Harmon_closed[identifier].Name = fielValue
				}
			case "Harmon_mute":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Harmon_mute[identifier].Name = fielValue
				}
			case "Harmonic":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Harmonic[identifier].Name = fielValue
				}
			case "Harmony":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Harmony[identifier].Name = fielValue
				}
			case "Harmony_alter":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Harmony_alter[identifier].Name = fielValue
				}
			case "Harp_pedals":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Harp_pedals[identifier].Name = fielValue
				}
			case "Heel_toe":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Heel_toe[identifier].Name = fielValue
				}
			case "Hole":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Hole[identifier].Name = fielValue
				case "Hole_type":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Hole[identifier].Hole_type = fielValue
				case "Hole_shape":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Hole[identifier].Hole_shape = fielValue
				}
			case "Hole_closed":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Hole_closed[identifier].Name = fielValue
				}
			case "Horizontal_turn":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Horizontal_turn[identifier].Name = fielValue
				}
			case "Identification":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Identification[identifier].Name = fielValue
				case "Source":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Identification[identifier].Source = fielValue
				}
			case "Image":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Image[identifier].Name = fielValue
				}
			case "Instrument":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Instrument[identifier].Name = fielValue
				}
			case "Instrument_change":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Instrument_change[identifier].Name = fielValue
				}
			case "Instrument_link":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Instrument_link[identifier].Name = fielValue
				}
			case "Interchangeable":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Interchangeable[identifier].Name = fielValue
				}
			case "Inversion":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Inversion[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Inversion[identifier].Text = fielValue
				}
			case "Key":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key[identifier].Name = fielValue
				}
			case "Key_accidental":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key_accidental[identifier].Name = fielValue
				}
			case "Key_octave":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key_octave[identifier].Name = fielValue
				}
			case "Kind":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Kind[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Kind[identifier].Text = fielValue
				}
			case "Level":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Level[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Level[identifier].Value = fielValue
				}
			case "Line_detail":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line_detail[identifier].Name = fielValue
				}
			case "Line_width":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line_width[identifier].Name = fielValue
				}
			case "Link":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Name = fielValue
				}
			case "Listen":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Listen[identifier].Name = fielValue
				}
			case "Listening":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Listening[identifier].Name = fielValue
				}
			case "Lyric":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Lyric[identifier].Name = fielValue
				}
			case "Lyric_font":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Lyric_font[identifier].Name = fielValue
				}
			case "Lyric_language":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Lyric_language[identifier].Name = fielValue
				case "EmptyString":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Lyric_language[identifier].EmptyString = fielValue
				}
			case "Measure_layout":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Measure_layout[identifier].Name = fielValue
				}
			case "Measure_numbering":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Measure_numbering[identifier].Name = fielValue
				}
			case "Measure_repeat":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Measure_repeat[identifier].Name = fielValue
				}
			case "Measure_style":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Measure_style[identifier].Name = fielValue
				}
			case "Membrane":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Membrane[identifier].Name = fielValue
				}
			case "Metal":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Metal[identifier].Name = fielValue
				}
			case "Metronome":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Metronome[identifier].Name = fielValue
				}
			case "Metronome_beam":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Metronome_beam[identifier].Name = fielValue
				}
			case "Metronome_note":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Metronome_note[identifier].Name = fielValue
				}
			case "Metronome_tied":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Metronome_tied[identifier].Name = fielValue
				}
			case "Metronome_tuplet":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Metronome_tuplet[identifier].Name = fielValue
				}
			case "Midi_device":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Midi_device[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Midi_device[identifier].Value = fielValue
				}
			case "Midi_instrument":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Midi_instrument[identifier].Name = fielValue
				case "Midi_name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Midi_instrument[identifier].Midi_name = fielValue
				}
			case "Miscellaneous":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Miscellaneous[identifier].Name = fielValue
				}
			case "Miscellaneous_field":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Miscellaneous_field[identifier].Value = fielValue
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Miscellaneous_field[identifier].Name = fielValue
				}
			case "Mordent":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Mordent[identifier].Name = fielValue
				}
			case "Multiple_rest":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Multiple_rest[identifier].Name = fielValue
				}
			case "Name_display":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Name_display[identifier].Name = fielValue
				}
			case "Non_arpeggiate":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Non_arpeggiate[identifier].Name = fielValue
				}
			case "Notations":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Notations[identifier].Name = fielValue
				}
			case "Note":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Note[identifier].Name = fielValue
				}
			case "Note_size":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Note_size[identifier].Name = fielValue
				}
			case "Note_type":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Note_type[identifier].Name = fielValue
				}
			case "Notehead":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Notehead[identifier].Name = fielValue
				}
			case "Notehead_text":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Notehead_text[identifier].Name = fielValue
				}
			case "Numeral":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Numeral[identifier].Name = fielValue
				}
			case "Numeral_key":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Numeral_key[identifier].Name = fielValue
				}
			case "Numeral_root":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Numeral_root[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Numeral_root[identifier].Text = fielValue
				}
			case "Octave_shift":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Octave_shift[identifier].Name = fielValue
				}
			case "Offset":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Offset[identifier].Name = fielValue
				}
			case "Opus":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Opus[identifier].Name = fielValue
				}
			case "Ornaments":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ornaments[identifier].Name = fielValue
				}
			case "Other_appearance":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_appearance[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_appearance[identifier].Value = fielValue
				case "Type":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_appearance[identifier].Type = fielValue
				}
			case "Other_listening":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_listening[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_listening[identifier].Value = fielValue
				case "Type":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_listening[identifier].Type = fielValue
				}
			case "Other_notation":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_notation[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_notation[identifier].Value = fielValue
				}
			case "Other_play":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_play[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_play[identifier].Value = fielValue
				case "Type":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Other_play[identifier].Type = fielValue
				}
			case "Page_layout":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Page_layout[identifier].Name = fielValue
				}
			case "Page_margins":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Page_margins[identifier].Name = fielValue
				}
			case "Part_clef":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Part_clef[identifier].Name = fielValue
				}
			case "Part_group":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Part_group[identifier].Name = fielValue
				case "Number":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Part_group[identifier].Number = fielValue
				}
			case "Part_link":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Part_link[identifier].Name = fielValue
				}
			case "Part_list":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Part_list[identifier].Name = fielValue
				}
			case "Part_symbol":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Part_symbol[identifier].Name = fielValue
				}
			case "Part_transpose":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Part_transpose[identifier].Name = fielValue
				}
			case "Pedal":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Pedal[identifier].Name = fielValue
				}
			case "Pedal_tuning":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Pedal_tuning[identifier].Name = fielValue
				}
			case "Percussion":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Percussion[identifier].Name = fielValue
				}
			case "Pitch":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Pitch[identifier].Name = fielValue
				}
			case "Pitched":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Pitched[identifier].Name = fielValue
				}
			case "Play":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Play[identifier].Name = fielValue
				case "Ipa":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Play[identifier].Ipa = fielValue
				}
			case "Player":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Player[identifier].Name = fielValue
				case "Player_name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Player[identifier].Player_name = fielValue
				}
			case "Principal_voice":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Principal_voice[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Principal_voice[identifier].Value = fielValue
				}
			case "Print":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Print[identifier].Name = fielValue
				}
			case "Release":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Release[identifier].Name = fielValue
				}
			case "Repeat":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Repeat[identifier].Name = fielValue
				}
			case "Rest":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rest[identifier].Name = fielValue
				}
			case "Root":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Root[identifier].Name = fielValue
				}
			case "Root_step":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Root_step[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Root_step[identifier].Text = fielValue
				}
			case "Scaling":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Scaling[identifier].Name = fielValue
				}
			case "Scordatura":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Scordatura[identifier].Name = fielValue
				}
			case "Score_instrument":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Score_instrument[identifier].Name = fielValue
				case "Instrument_name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Score_instrument[identifier].Instrument_name = fielValue
				case "Instrument_abbreviation":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Score_instrument[identifier].Instrument_abbreviation = fielValue
				}
			case "Score_part":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Score_part[identifier].Name = fielValue
				}
			case "Score_partwise":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Score_partwise[identifier].Name = fielValue
				}
			case "Score_timewise":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Score_timewise[identifier].Name = fielValue
				}
			case "Segno":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Segno[identifier].Name = fielValue
				}
			case "Slash":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Slash[identifier].Name = fielValue
				}
			case "Slide":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Slide[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Slide[identifier].Value = fielValue
				}
			case "Slur":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Slur[identifier].Name = fielValue
				}
			case "Sound":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sound[identifier].Name = fielValue
				case "Segno":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sound[identifier].Segno = fielValue
				case "Dalsegno":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sound[identifier].Dalsegno = fielValue
				case "Coda":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sound[identifier].Coda = fielValue
				case "Tocoda":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sound[identifier].Tocoda = fielValue
				case "Fine":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sound[identifier].Fine = fielValue
				}
			case "Staff_details":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Staff_details[identifier].Name = fielValue
				}
			case "Staff_divide":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Staff_divide[identifier].Name = fielValue
				}
			case "Staff_layout":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Staff_layout[identifier].Name = fielValue
				}
			case "Staff_size":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Staff_size[identifier].Name = fielValue
				}
			case "Staff_tuning":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Staff_tuning[identifier].Name = fielValue
				}
			case "Stem":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Stem[identifier].Name = fielValue
				}
			case "Stick":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Stick[identifier].Name = fielValue
				}
			case "String_mute":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_String_mute[identifier].Name = fielValue
				}
			case "Strong_accent":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Strong_accent[identifier].Name = fielValue
				}
			case "Supports":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Supports[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Supports[identifier].Value = fielValue
				}
			case "Swing":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Swing[identifier].Name = fielValue
				case "Swing_style":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Swing[identifier].Swing_style = fielValue
				}
			case "Sync":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sync[identifier].Name = fielValue
				}
			case "System_dividers":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_System_dividers[identifier].Name = fielValue
				}
			case "System_layout":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_System_layout[identifier].Name = fielValue
				}
			case "System_margins":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_System_margins[identifier].Name = fielValue
				}
			case "Tap":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tap[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tap[identifier].Value = fielValue
				}
			case "Technical":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Technical[identifier].Name = fielValue
				case "Astring":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Technical[identifier].Astring = fielValue
				}
			case "Text_element_data":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text_element_data[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text_element_data[identifier].Value = fielValue
				case "EmptyString":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text_element_data[identifier].EmptyString = fielValue
				}
			case "Tie":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tie[identifier].Name = fielValue
				}
			case "Tied":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tied[identifier].Name = fielValue
				}
			case "Time":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Time[identifier].Name = fielValue
				case "Senza_misura":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Time[identifier].Senza_misura = fielValue
				}
			case "Time_modification":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Time_modification[identifier].Name = fielValue
				}
			case "Timpani":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Timpani[identifier].Name = fielValue
				}
			case "Transpose":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Transpose[identifier].Name = fielValue
				}
			case "Tremolo":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tremolo[identifier].Name = fielValue
				}
			case "Tuplet":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tuplet[identifier].Name = fielValue
				}
			case "Tuplet_dot":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tuplet_dot[identifier].Name = fielValue
				}
			case "Tuplet_number":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tuplet_number[identifier].Name = fielValue
				}
			case "Tuplet_portion":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tuplet_portion[identifier].Name = fielValue
				}
			case "Tuplet_type":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tuplet_type[identifier].Name = fielValue
				}
			case "Typed_text":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Typed_text[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Typed_text[identifier].Value = fielValue
				case "Type":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Typed_text[identifier].Type = fielValue
				}
			case "Unpitched":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Unpitched[identifier].Name = fielValue
				}
			case "Virtual_instrument":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Virtual_instrument[identifier].Name = fielValue
				case "Virtual_library":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Virtual_instrument[identifier].Virtual_library = fielValue
				case "Virtual_name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Virtual_instrument[identifier].Virtual_name = fielValue
				}
			case "Wait":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Wait[identifier].Name = fielValue
				}
			case "Wavy_line":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Wavy_line[identifier].Name = fielValue
				}
			case "Wedge":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Wedge[identifier].Name = fielValue
				}
			case "Wood":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Wood[identifier].Name = fielValue
				}
			case "Work":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Work[identifier].Name = fielValue
				case "Work_number":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Work[identifier].Work_number = fielValue
				case "Work_title":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Work[identifier].Work_title = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "Accidental":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Accidental_mark":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Accidental_text":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Accord":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Accordion_registration":
				switch fieldName {
				// insertion point for field dependant code
				case "Accordion_high":
					targetIdentifier := ident.Name
					__gong__map_Accordion_registration[identifier].Accordion_high = __gong__map_Empty[targetIdentifier]
				case "Accordion_low":
					targetIdentifier := ident.Name
					__gong__map_Accordion_registration[identifier].Accordion_low = __gong__map_Empty[targetIdentifier]
				}
			case "AnyType":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Appearance":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Arpeggiate":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Arrow":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Articulations":
				switch fieldName {
				// insertion point for field dependant code
				case "Accent":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Accent = __gong__map_Empty_placement[targetIdentifier]
				case "Strong_accent":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Strong_accent = __gong__map_Strong_accent[targetIdentifier]
				case "Staccato":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Staccato = __gong__map_Empty_placement[targetIdentifier]
				case "Tenuto":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Tenuto = __gong__map_Empty_placement[targetIdentifier]
				case "Detached_legato":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Detached_legato = __gong__map_Empty_placement[targetIdentifier]
				case "Staccatissimo":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Staccatissimo = __gong__map_Empty_placement[targetIdentifier]
				case "Spiccato":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Spiccato = __gong__map_Empty_placement[targetIdentifier]
				case "Scoop":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Scoop = __gong__map_Empty_line[targetIdentifier]
				case "Plop":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Plop = __gong__map_Empty_line[targetIdentifier]
				case "Doit":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Doit = __gong__map_Empty_line[targetIdentifier]
				case "Falloff":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Falloff = __gong__map_Empty_line[targetIdentifier]
				case "Breath_mark":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Breath_mark = __gong__map_Breath_mark[targetIdentifier]
				case "Caesura":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Caesura = __gong__map_Caesura[targetIdentifier]
				case "Stress":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Stress = __gong__map_Empty_placement[targetIdentifier]
				case "Unstress":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Unstress = __gong__map_Empty_placement[targetIdentifier]
				case "Soft_accent":
					targetIdentifier := ident.Name
					__gong__map_Articulations[identifier].Soft_accent = __gong__map_Empty_placement[targetIdentifier]
				}
			case "Assess":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Attributes":
				switch fieldName {
				// insertion point for field dependant code
				case "Part_symbol":
					targetIdentifier := ident.Name
					__gong__map_Attributes[identifier].Part_symbol = __gong__map_Part_symbol[targetIdentifier]
				}
			case "Backup":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Bar_style_color":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Barline":
				switch fieldName {
				// insertion point for field dependant code
				case "Bar_style":
					targetIdentifier := ident.Name
					__gong__map_Barline[identifier].Bar_style = __gong__map_Bar_style_color[targetIdentifier]
				case "Wavy_line":
					targetIdentifier := ident.Name
					__gong__map_Barline[identifier].Wavy_line = __gong__map_Wavy_line[targetIdentifier]
				case "Fermata":
					targetIdentifier := ident.Name
					__gong__map_Barline[identifier].Fermata = __gong__map_Fermata[targetIdentifier]
				case "Ending":
					targetIdentifier := ident.Name
					__gong__map_Barline[identifier].Ending = __gong__map_Ending[targetIdentifier]
				case "Repeat":
					targetIdentifier := ident.Name
					__gong__map_Barline[identifier].Repeat = __gong__map_Repeat[targetIdentifier]
				}
			case "Barre":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Bass":
				switch fieldName {
				// insertion point for field dependant code
				case "Bass_step":
					targetIdentifier := ident.Name
					__gong__map_Bass[identifier].Bass_step = __gong__map_Bass_step[targetIdentifier]
				case "Bass_alter":
					targetIdentifier := ident.Name
					__gong__map_Bass[identifier].Bass_alter = __gong__map_Harmony_alter[targetIdentifier]
				}
			case "Bass_step":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Beam":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Beat_repeat":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Beat_unit_tied":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Beater":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Bend":
				switch fieldName {
				// insertion point for field dependant code
				case "Pre_bend":
					targetIdentifier := ident.Name
					__gong__map_Bend[identifier].Pre_bend = __gong__map_Empty[targetIdentifier]
				case "Release":
					targetIdentifier := ident.Name
					__gong__map_Bend[identifier].Release = __gong__map_Release[targetIdentifier]
				}
			case "Bookmark":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Bracket":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Breath_mark":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Caesura":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Cancel":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Clef":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Coda":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Credit":
				switch fieldName {
				// insertion point for field dependant code
				case "Credit_image":
					targetIdentifier := ident.Name
					__gong__map_Credit[identifier].Credit_image = __gong__map_Image[targetIdentifier]
				}
			case "Dashes":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Defaults":
				switch fieldName {
				// insertion point for field dependant code
				case "Scaling":
					targetIdentifier := ident.Name
					__gong__map_Defaults[identifier].Scaling = __gong__map_Scaling[targetIdentifier]
				case "Concert_score":
					targetIdentifier := ident.Name
					__gong__map_Defaults[identifier].Concert_score = __gong__map_Empty[targetIdentifier]
				case "Appearance":
					targetIdentifier := ident.Name
					__gong__map_Defaults[identifier].Appearance = __gong__map_Appearance[targetIdentifier]
				case "Music_font":
					targetIdentifier := ident.Name
					__gong__map_Defaults[identifier].Music_font = __gong__map_Empty_font[targetIdentifier]
				case "Word_font":
					targetIdentifier := ident.Name
					__gong__map_Defaults[identifier].Word_font = __gong__map_Empty_font[targetIdentifier]
				}
			case "Degree":
				switch fieldName {
				// insertion point for field dependant code
				case "Degree_value":
					targetIdentifier := ident.Name
					__gong__map_Degree[identifier].Degree_value = __gong__map_Degree_value[targetIdentifier]
				case "Degree_alter":
					targetIdentifier := ident.Name
					__gong__map_Degree[identifier].Degree_alter = __gong__map_Degree_alter[targetIdentifier]
				case "Degree_type":
					targetIdentifier := ident.Name
					__gong__map_Degree[identifier].Degree_type = __gong__map_Degree_type[targetIdentifier]
				}
			case "Degree_alter":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Degree_type":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Degree_value":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Direction":
				switch fieldName {
				// insertion point for field dependant code
				case "Offset":
					targetIdentifier := ident.Name
					__gong__map_Direction[identifier].Offset = __gong__map_Offset[targetIdentifier]
				case "Sound":
					targetIdentifier := ident.Name
					__gong__map_Direction[identifier].Sound = __gong__map_Sound[targetIdentifier]
				case "Listening":
					targetIdentifier := ident.Name
					__gong__map_Direction[identifier].Listening = __gong__map_Listening[targetIdentifier]
				}
			case "Direction_type":
				switch fieldName {
				// insertion point for field dependant code
				case "Wedge":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Wedge = __gong__map_Wedge[targetIdentifier]
				case "Dashes":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Dashes = __gong__map_Dashes[targetIdentifier]
				case "Bracket":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Bracket = __gong__map_Bracket[targetIdentifier]
				case "Pedal":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Pedal = __gong__map_Pedal[targetIdentifier]
				case "Metronome":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Metronome = __gong__map_Metronome[targetIdentifier]
				case "Octave_shift":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Octave_shift = __gong__map_Octave_shift[targetIdentifier]
				case "Harp_pedals":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Harp_pedals = __gong__map_Harp_pedals[targetIdentifier]
				case "Damp":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Damp = __gong__map_Empty_print_style_align_id[targetIdentifier]
				case "Damp_all":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Damp_all = __gong__map_Empty_print_style_align_id[targetIdentifier]
				case "Eyeglasses":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Eyeglasses = __gong__map_Empty_print_style_align_id[targetIdentifier]
				case "String_mute":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].String_mute = __gong__map_String_mute[targetIdentifier]
				case "Scordatura":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Scordatura = __gong__map_Scordatura[targetIdentifier]
				case "Image":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Image = __gong__map_Image[targetIdentifier]
				case "Principal_voice":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Principal_voice = __gong__map_Principal_voice[targetIdentifier]
				case "Accordion_registration":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Accordion_registration = __gong__map_Accordion_registration[targetIdentifier]
				case "Staff_divide":
					targetIdentifier := ident.Name
					__gong__map_Direction_type[identifier].Staff_divide = __gong__map_Staff_divide[targetIdentifier]
				}
			case "Distance":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Double":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Dynamics":
				switch fieldName {
				// insertion point for field dependant code
				case "P":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].P = __gong__map_Empty[targetIdentifier]
				case "Pp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Pp = __gong__map_Empty[targetIdentifier]
				case "Ppp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Ppp = __gong__map_Empty[targetIdentifier]
				case "Pppp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Pppp = __gong__map_Empty[targetIdentifier]
				case "Ppppp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Ppppp = __gong__map_Empty[targetIdentifier]
				case "Pppppp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Pppppp = __gong__map_Empty[targetIdentifier]
				case "F":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].F = __gong__map_Empty[targetIdentifier]
				case "Ff":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Ff = __gong__map_Empty[targetIdentifier]
				case "Fff":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Fff = __gong__map_Empty[targetIdentifier]
				case "Ffff":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Ffff = __gong__map_Empty[targetIdentifier]
				case "Fffff":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Fffff = __gong__map_Empty[targetIdentifier]
				case "Ffffff":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Ffffff = __gong__map_Empty[targetIdentifier]
				case "Mp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Mp = __gong__map_Empty[targetIdentifier]
				case "Mf":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Mf = __gong__map_Empty[targetIdentifier]
				case "Sf":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Sf = __gong__map_Empty[targetIdentifier]
				case "Sfp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Sfp = __gong__map_Empty[targetIdentifier]
				case "Sfpp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Sfpp = __gong__map_Empty[targetIdentifier]
				case "Fp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Fp = __gong__map_Empty[targetIdentifier]
				case "Rf":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Rf = __gong__map_Empty[targetIdentifier]
				case "Rfz":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Rfz = __gong__map_Empty[targetIdentifier]
				case "Sfz":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Sfz = __gong__map_Empty[targetIdentifier]
				case "Sffz":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Sffz = __gong__map_Empty[targetIdentifier]
				case "Fz":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Fz = __gong__map_Empty[targetIdentifier]
				case "N":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].N = __gong__map_Empty[targetIdentifier]
				case "Pf":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Pf = __gong__map_Empty[targetIdentifier]
				case "Sfzp":
					targetIdentifier := ident.Name
					__gong__map_Dynamics[identifier].Sfzp = __gong__map_Empty[targetIdentifier]
				}
			case "Effect":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Elision":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty_font":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty_line":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty_placement":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty_placement_smufl":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty_print_object_style_align":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty_print_style":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty_print_style_align":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty_print_style_align_id":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Empty_trill_sound":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Encoding":
				switch fieldName {
				// insertion point for field dependant code
				case "Encoder":
					targetIdentifier := ident.Name
					__gong__map_Encoding[identifier].Encoder = __gong__map_Typed_text[targetIdentifier]
				case "Supports":
					targetIdentifier := ident.Name
					__gong__map_Encoding[identifier].Supports = __gong__map_Supports[targetIdentifier]
				}
			case "Ending":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Extend":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Feature":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Fermata":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Figure":
				switch fieldName {
				// insertion point for field dependant code
				case "Extend":
					targetIdentifier := ident.Name
					__gong__map_Figure[identifier].Extend = __gong__map_Extend[targetIdentifier]
				}
			case "Figured_bass":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Fingering":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "First_fret":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Foo":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "For_part":
				switch fieldName {
				// insertion point for field dependant code
				case "Part_clef":
					targetIdentifier := ident.Name
					__gong__map_For_part[identifier].Part_clef = __gong__map_Part_clef[targetIdentifier]
				case "Part_transpose":
					targetIdentifier := ident.Name
					__gong__map_For_part[identifier].Part_transpose = __gong__map_Part_transpose[targetIdentifier]
				}
			case "Formatted_symbol":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Formatted_symbol_id":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Forward":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Frame":
				switch fieldName {
				// insertion point for field dependant code
				case "First_fret":
					targetIdentifier := ident.Name
					__gong__map_Frame[identifier].First_fret = __gong__map_First_fret[targetIdentifier]
				}
			case "Frame_note":
				switch fieldName {
				// insertion point for field dependant code
				case "Fret":
					targetIdentifier := ident.Name
					__gong__map_Frame_note[identifier].Fret = __gong__map_Fret[targetIdentifier]
				case "Fingering":
					targetIdentifier := ident.Name
					__gong__map_Frame_note[identifier].Fingering = __gong__map_Fingering[targetIdentifier]
				case "Barre":
					targetIdentifier := ident.Name
					__gong__map_Frame_note[identifier].Barre = __gong__map_Barre[targetIdentifier]
				}
			case "Fret":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Glass":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Glissando":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Glyph":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Grace":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Group_barline":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Group_symbol":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Grouping":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Hammer_on_pull_off":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Handbell":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Harmon_closed":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Harmon_mute":
				switch fieldName {
				// insertion point for field dependant code
				case "Harmon_closed":
					targetIdentifier := ident.Name
					__gong__map_Harmon_mute[identifier].Harmon_closed = __gong__map_Harmon_closed[targetIdentifier]
				}
			case "Harmonic":
				switch fieldName {
				// insertion point for field dependant code
				case "Natural":
					targetIdentifier := ident.Name
					__gong__map_Harmonic[identifier].Natural = __gong__map_Empty[targetIdentifier]
				case "Artificial":
					targetIdentifier := ident.Name
					__gong__map_Harmonic[identifier].Artificial = __gong__map_Empty[targetIdentifier]
				case "Base_pitch":
					targetIdentifier := ident.Name
					__gong__map_Harmonic[identifier].Base_pitch = __gong__map_Empty[targetIdentifier]
				case "Touching_pitch":
					targetIdentifier := ident.Name
					__gong__map_Harmonic[identifier].Touching_pitch = __gong__map_Empty[targetIdentifier]
				case "Sounding_pitch":
					targetIdentifier := ident.Name
					__gong__map_Harmonic[identifier].Sounding_pitch = __gong__map_Empty[targetIdentifier]
				}
			case "Harmony":
				switch fieldName {
				// insertion point for field dependant code
				case "Frame":
					targetIdentifier := ident.Name
					__gong__map_Harmony[identifier].Frame = __gong__map_Frame[targetIdentifier]
				case "Offset":
					targetIdentifier := ident.Name
					__gong__map_Harmony[identifier].Offset = __gong__map_Offset[targetIdentifier]
				}
			case "Harmony_alter":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Harp_pedals":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Heel_toe":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Hole":
				switch fieldName {
				// insertion point for field dependant code
				case "Hole_closed":
					targetIdentifier := ident.Name
					__gong__map_Hole[identifier].Hole_closed = __gong__map_Hole_closed[targetIdentifier]
				}
			case "Hole_closed":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Horizontal_turn":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Identification":
				switch fieldName {
				// insertion point for field dependant code
				case "Encoding":
					targetIdentifier := ident.Name
					__gong__map_Identification[identifier].Encoding = __gong__map_Encoding[targetIdentifier]
				case "Miscellaneous":
					targetIdentifier := ident.Name
					__gong__map_Identification[identifier].Miscellaneous = __gong__map_Miscellaneous[targetIdentifier]
				}
			case "Image":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Instrument":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Instrument_change":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Instrument_link":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Interchangeable":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Inversion":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Key":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Key_accidental":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Key_octave":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Kind":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Level":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Line_detail":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Line_width":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Link":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Listen":
				switch fieldName {
				// insertion point for field dependant code
				case "Assess":
					targetIdentifier := ident.Name
					__gong__map_Listen[identifier].Assess = __gong__map_Assess[targetIdentifier]
				case "Wait":
					targetIdentifier := ident.Name
					__gong__map_Listen[identifier].Wait = __gong__map_Wait[targetIdentifier]
				case "Other_listen":
					targetIdentifier := ident.Name
					__gong__map_Listen[identifier].Other_listen = __gong__map_Other_listening[targetIdentifier]
				}
			case "Listening":
				switch fieldName {
				// insertion point for field dependant code
				case "Offset":
					targetIdentifier := ident.Name
					__gong__map_Listening[identifier].Offset = __gong__map_Offset[targetIdentifier]
				case "Sync":
					targetIdentifier := ident.Name
					__gong__map_Listening[identifier].Sync = __gong__map_Sync[targetIdentifier]
				case "Other_listening":
					targetIdentifier := ident.Name
					__gong__map_Listening[identifier].Other_listening = __gong__map_Other_listening[targetIdentifier]
				}
			case "Lyric":
				switch fieldName {
				// insertion point for field dependant code
				case "End_line":
					targetIdentifier := ident.Name
					__gong__map_Lyric[identifier].End_line = __gong__map_Empty[targetIdentifier]
				case "End_paragraph":
					targetIdentifier := ident.Name
					__gong__map_Lyric[identifier].End_paragraph = __gong__map_Empty[targetIdentifier]
				case "Extend":
					targetIdentifier := ident.Name
					__gong__map_Lyric[identifier].Extend = __gong__map_Extend[targetIdentifier]
				case "Laughing":
					targetIdentifier := ident.Name
					__gong__map_Lyric[identifier].Laughing = __gong__map_Empty[targetIdentifier]
				case "Humming":
					targetIdentifier := ident.Name
					__gong__map_Lyric[identifier].Humming = __gong__map_Empty[targetIdentifier]
				}
			case "Lyric_font":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Lyric_language":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Measure_layout":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Measure_numbering":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Measure_repeat":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Measure_style":
				switch fieldName {
				// insertion point for field dependant code
				case "Multiple_rest":
					targetIdentifier := ident.Name
					__gong__map_Measure_style[identifier].Multiple_rest = __gong__map_Multiple_rest[targetIdentifier]
				case "Measure_repeat":
					targetIdentifier := ident.Name
					__gong__map_Measure_style[identifier].Measure_repeat = __gong__map_Measure_repeat[targetIdentifier]
				case "Beat_repeat":
					targetIdentifier := ident.Name
					__gong__map_Measure_style[identifier].Beat_repeat = __gong__map_Beat_repeat[targetIdentifier]
				case "Slash":
					targetIdentifier := ident.Name
					__gong__map_Measure_style[identifier].Slash = __gong__map_Slash[targetIdentifier]
				}
			case "Membrane":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Metal":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Metronome":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Metronome_beam":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Metronome_note":
				switch fieldName {
				// insertion point for field dependant code
				case "Metronome_tied":
					targetIdentifier := ident.Name
					__gong__map_Metronome_note[identifier].Metronome_tied = __gong__map_Metronome_tied[targetIdentifier]
				case "Metronome_tuplet":
					targetIdentifier := ident.Name
					__gong__map_Metronome_note[identifier].Metronome_tuplet = __gong__map_Metronome_tuplet[targetIdentifier]
				}
			case "Metronome_tied":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Metronome_tuplet":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Midi_device":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Midi_instrument":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Miscellaneous":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Miscellaneous_field":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Mordent":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Multiple_rest":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Name_display":
				switch fieldName {
				// insertion point for field dependant code
				case "Accidental_text":
					targetIdentifier := ident.Name
					__gong__map_Name_display[identifier].Accidental_text = __gong__map_Accidental_text[targetIdentifier]
				}
			case "Non_arpeggiate":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Notations":
				switch fieldName {
				// insertion point for field dependant code
				case "Tied":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Tied = __gong__map_Tied[targetIdentifier]
				case "Slur":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Slur = __gong__map_Slur[targetIdentifier]
				case "Tuplet":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Tuplet = __gong__map_Tuplet[targetIdentifier]
				case "Glissando":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Glissando = __gong__map_Glissando[targetIdentifier]
				case "Slide":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Slide = __gong__map_Slide[targetIdentifier]
				case "Ornaments":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Ornaments = __gong__map_Ornaments[targetIdentifier]
				case "Technical":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Technical = __gong__map_Technical[targetIdentifier]
				case "Articulations":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Articulations = __gong__map_Articulations[targetIdentifier]
				case "Dynamics":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Dynamics = __gong__map_Dynamics[targetIdentifier]
				case "Fermata":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Fermata = __gong__map_Fermata[targetIdentifier]
				case "Arpeggiate":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Arpeggiate = __gong__map_Arpeggiate[targetIdentifier]
				case "Non_arpeggiate":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Non_arpeggiate = __gong__map_Non_arpeggiate[targetIdentifier]
				case "Accidental_mark":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Accidental_mark = __gong__map_Accidental_mark[targetIdentifier]
				case "Other_notation":
					targetIdentifier := ident.Name
					__gong__map_Notations[identifier].Other_notation = __gong__map_Other_notation[targetIdentifier]
				}
			case "Note":
				switch fieldName {
				// insertion point for field dependant code
				case "Type_":
					targetIdentifier := ident.Name
					__gong__map_Note[identifier].Type_ = __gong__map_Note_type[targetIdentifier]
				case "Accidental":
					targetIdentifier := ident.Name
					__gong__map_Note[identifier].Accidental = __gong__map_Accidental[targetIdentifier]
				case "Time_modification":
					targetIdentifier := ident.Name
					__gong__map_Note[identifier].Time_modification = __gong__map_Time_modification[targetIdentifier]
				case "Stem":
					targetIdentifier := ident.Name
					__gong__map_Note[identifier].Stem = __gong__map_Stem[targetIdentifier]
				case "Notehead":
					targetIdentifier := ident.Name
					__gong__map_Note[identifier].Notehead = __gong__map_Notehead[targetIdentifier]
				case "Notehead_text":
					targetIdentifier := ident.Name
					__gong__map_Note[identifier].Notehead_text = __gong__map_Notehead_text[targetIdentifier]
				case "Beam":
					targetIdentifier := ident.Name
					__gong__map_Note[identifier].Beam = __gong__map_Beam[targetIdentifier]
				case "Play":
					targetIdentifier := ident.Name
					__gong__map_Note[identifier].Play = __gong__map_Play[targetIdentifier]
				case "Listen":
					targetIdentifier := ident.Name
					__gong__map_Note[identifier].Listen = __gong__map_Listen[targetIdentifier]
				}
			case "Note_size":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Note_type":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Notehead":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Notehead_text":
				switch fieldName {
				// insertion point for field dependant code
				case "Accidental_text":
					targetIdentifier := ident.Name
					__gong__map_Notehead_text[identifier].Accidental_text = __gong__map_Accidental_text[targetIdentifier]
				}
			case "Numeral":
				switch fieldName {
				// insertion point for field dependant code
				case "Numeral_root":
					targetIdentifier := ident.Name
					__gong__map_Numeral[identifier].Numeral_root = __gong__map_Numeral_root[targetIdentifier]
				case "Numeral_alter":
					targetIdentifier := ident.Name
					__gong__map_Numeral[identifier].Numeral_alter = __gong__map_Harmony_alter[targetIdentifier]
				case "Numeral_key":
					targetIdentifier := ident.Name
					__gong__map_Numeral[identifier].Numeral_key = __gong__map_Numeral_key[targetIdentifier]
				}
			case "Numeral_key":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Numeral_root":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Octave_shift":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Offset":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Opus":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Ornaments":
				switch fieldName {
				// insertion point for field dependant code
				case "Trill_mark":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Trill_mark = __gong__map_Empty_trill_sound[targetIdentifier]
				case "Turn":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Turn = __gong__map_Horizontal_turn[targetIdentifier]
				case "Delayed_turn":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Delayed_turn = __gong__map_Horizontal_turn[targetIdentifier]
				case "Inverted_turn":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Inverted_turn = __gong__map_Horizontal_turn[targetIdentifier]
				case "Delayed_inverted_turn":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Delayed_inverted_turn = __gong__map_Horizontal_turn[targetIdentifier]
				case "Vertical_turn":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Vertical_turn = __gong__map_Empty_trill_sound[targetIdentifier]
				case "Inverted_vertical_turn":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Inverted_vertical_turn = __gong__map_Empty_trill_sound[targetIdentifier]
				case "Shake":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Shake = __gong__map_Empty_trill_sound[targetIdentifier]
				case "Wavy_line":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Wavy_line = __gong__map_Wavy_line[targetIdentifier]
				case "Mordent":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Mordent = __gong__map_Mordent[targetIdentifier]
				case "Inverted_mordent":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Inverted_mordent = __gong__map_Mordent[targetIdentifier]
				case "Schleifer":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Schleifer = __gong__map_Empty_placement[targetIdentifier]
				case "Tremolo":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Tremolo = __gong__map_Tremolo[targetIdentifier]
				case "Haydn":
					targetIdentifier := ident.Name
					__gong__map_Ornaments[identifier].Haydn = __gong__map_Empty_trill_sound[targetIdentifier]
				}
			case "Other_appearance":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Other_listening":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Other_notation":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Other_play":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Page_layout":
				switch fieldName {
				// insertion point for field dependant code
				case "Page_margins":
					targetIdentifier := ident.Name
					__gong__map_Page_layout[identifier].Page_margins = __gong__map_Page_margins[targetIdentifier]
				}
			case "Page_margins":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Part_clef":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Part_group":
				switch fieldName {
				// insertion point for field dependant code
				case "Group_name_display":
					targetIdentifier := ident.Name
					__gong__map_Part_group[identifier].Group_name_display = __gong__map_Name_display[targetIdentifier]
				case "Group_abbreviation_display":
					targetIdentifier := ident.Name
					__gong__map_Part_group[identifier].Group_abbreviation_display = __gong__map_Name_display[targetIdentifier]
				case "Group_symbol":
					targetIdentifier := ident.Name
					__gong__map_Part_group[identifier].Group_symbol = __gong__map_Group_symbol[targetIdentifier]
				case "Group_barline":
					targetIdentifier := ident.Name
					__gong__map_Part_group[identifier].Group_barline = __gong__map_Group_barline[targetIdentifier]
				case "Group_time":
					targetIdentifier := ident.Name
					__gong__map_Part_group[identifier].Group_time = __gong__map_Empty[targetIdentifier]
				}
			case "Part_link":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Part_list":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Part_symbol":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Part_transpose":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Pedal":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Pedal_tuning":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Percussion":
				switch fieldName {
				// insertion point for field dependant code
				case "Glass":
					targetIdentifier := ident.Name
					__gong__map_Percussion[identifier].Glass = __gong__map_Glass[targetIdentifier]
				case "Metal":
					targetIdentifier := ident.Name
					__gong__map_Percussion[identifier].Metal = __gong__map_Metal[targetIdentifier]
				case "Wood":
					targetIdentifier := ident.Name
					__gong__map_Percussion[identifier].Wood = __gong__map_Wood[targetIdentifier]
				case "Pitched":
					targetIdentifier := ident.Name
					__gong__map_Percussion[identifier].Pitched = __gong__map_Pitched[targetIdentifier]
				case "Membrane":
					targetIdentifier := ident.Name
					__gong__map_Percussion[identifier].Membrane = __gong__map_Membrane[targetIdentifier]
				case "Effect":
					targetIdentifier := ident.Name
					__gong__map_Percussion[identifier].Effect = __gong__map_Effect[targetIdentifier]
				case "Timpani":
					targetIdentifier := ident.Name
					__gong__map_Percussion[identifier].Timpani = __gong__map_Timpani[targetIdentifier]
				case "Beater":
					targetIdentifier := ident.Name
					__gong__map_Percussion[identifier].Beater = __gong__map_Beater[targetIdentifier]
				case "Stick":
					targetIdentifier := ident.Name
					__gong__map_Percussion[identifier].Stick = __gong__map_Stick[targetIdentifier]
				}
			case "Pitch":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Pitched":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Play":
				switch fieldName {
				// insertion point for field dependant code
				case "Other_play":
					targetIdentifier := ident.Name
					__gong__map_Play[identifier].Other_play = __gong__map_Other_play[targetIdentifier]
				}
			case "Player":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Principal_voice":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Print":
				switch fieldName {
				// insertion point for field dependant code
				case "Measure_layout":
					targetIdentifier := ident.Name
					__gong__map_Print[identifier].Measure_layout = __gong__map_Measure_layout[targetIdentifier]
				case "Measure_numbering":
					targetIdentifier := ident.Name
					__gong__map_Print[identifier].Measure_numbering = __gong__map_Measure_numbering[targetIdentifier]
				case "Part_name_display":
					targetIdentifier := ident.Name
					__gong__map_Print[identifier].Part_name_display = __gong__map_Name_display[targetIdentifier]
				case "Part_abbreviation_display":
					targetIdentifier := ident.Name
					__gong__map_Print[identifier].Part_abbreviation_display = __gong__map_Name_display[targetIdentifier]
				}
			case "Release":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Repeat":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Rest":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Root":
				switch fieldName {
				// insertion point for field dependant code
				case "Root_step":
					targetIdentifier := ident.Name
					__gong__map_Root[identifier].Root_step = __gong__map_Root_step[targetIdentifier]
				case "Root_alter":
					targetIdentifier := ident.Name
					__gong__map_Root[identifier].Root_alter = __gong__map_Harmony_alter[targetIdentifier]
				}
			case "Root_step":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Scaling":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Scordatura":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Score_instrument":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Score_part":
				switch fieldName {
				// insertion point for field dependant code
				case "Identification":
					targetIdentifier := ident.Name
					__gong__map_Score_part[identifier].Identification = __gong__map_Identification[targetIdentifier]
				case "Part_name_display":
					targetIdentifier := ident.Name
					__gong__map_Score_part[identifier].Part_name_display = __gong__map_Name_display[targetIdentifier]
				case "Part_abbreviation_display":
					targetIdentifier := ident.Name
					__gong__map_Score_part[identifier].Part_abbreviation_display = __gong__map_Name_display[targetIdentifier]
				}
			case "Score_partwise":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Score_timewise":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Segno":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Slash":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Slide":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Slur":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Sound":
				switch fieldName {
				// insertion point for field dependant code
				case "Swing":
					targetIdentifier := ident.Name
					__gong__map_Sound[identifier].Swing = __gong__map_Swing[targetIdentifier]
				case "Offset":
					targetIdentifier := ident.Name
					__gong__map_Sound[identifier].Offset = __gong__map_Offset[targetIdentifier]
				}
			case "Staff_details":
				switch fieldName {
				// insertion point for field dependant code
				case "Staff_size":
					targetIdentifier := ident.Name
					__gong__map_Staff_details[identifier].Staff_size = __gong__map_Staff_size[targetIdentifier]
				}
			case "Staff_divide":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Staff_layout":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Staff_size":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Staff_tuning":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Stem":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Stick":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "String_mute":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Strong_accent":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Supports":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Swing":
				switch fieldName {
				// insertion point for field dependant code
				case "Straight":
					targetIdentifier := ident.Name
					__gong__map_Swing[identifier].Straight = __gong__map_Empty[targetIdentifier]
				}
			case "Sync":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "System_dividers":
				switch fieldName {
				// insertion point for field dependant code
				case "Left_divider":
					targetIdentifier := ident.Name
					__gong__map_System_dividers[identifier].Left_divider = __gong__map_Empty_print_object_style_align[targetIdentifier]
				case "Right_divider":
					targetIdentifier := ident.Name
					__gong__map_System_dividers[identifier].Right_divider = __gong__map_Empty_print_object_style_align[targetIdentifier]
				}
			case "System_layout":
				switch fieldName {
				// insertion point for field dependant code
				case "System_margins":
					targetIdentifier := ident.Name
					__gong__map_System_layout[identifier].System_margins = __gong__map_System_margins[targetIdentifier]
				case "System_dividers":
					targetIdentifier := ident.Name
					__gong__map_System_layout[identifier].System_dividers = __gong__map_System_dividers[targetIdentifier]
				}
			case "System_margins":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tap":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Technical":
				switch fieldName {
				// insertion point for field dependant code
				case "Up_bow":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Up_bow = __gong__map_Empty_placement[targetIdentifier]
				case "Down_bow":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Down_bow = __gong__map_Empty_placement[targetIdentifier]
				case "Harmonic":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Harmonic = __gong__map_Harmonic[targetIdentifier]
				case "Open_string":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Open_string = __gong__map_Empty_placement[targetIdentifier]
				case "Thumb_position":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Thumb_position = __gong__map_Empty_placement[targetIdentifier]
				case "Fingering":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Fingering = __gong__map_Fingering[targetIdentifier]
				case "Double_tongue":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Double_tongue = __gong__map_Empty_placement[targetIdentifier]
				case "Triple_tongue":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Triple_tongue = __gong__map_Empty_placement[targetIdentifier]
				case "Stopped":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Stopped = __gong__map_Empty_placement_smufl[targetIdentifier]
				case "Snap_pizzicato":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Snap_pizzicato = __gong__map_Empty_placement[targetIdentifier]
				case "Fret":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Fret = __gong__map_Fret[targetIdentifier]
				case "Hammer_on":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Hammer_on = __gong__map_Hammer_on_pull_off[targetIdentifier]
				case "Pull_off":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Pull_off = __gong__map_Hammer_on_pull_off[targetIdentifier]
				case "Bend":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Bend = __gong__map_Bend[targetIdentifier]
				case "Tap":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Tap = __gong__map_Tap[targetIdentifier]
				case "Heel":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Heel = __gong__map_Heel_toe[targetIdentifier]
				case "Toe":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Toe = __gong__map_Heel_toe[targetIdentifier]
				case "Fingernails":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Fingernails = __gong__map_Empty_placement[targetIdentifier]
				case "Hole":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Hole = __gong__map_Hole[targetIdentifier]
				case "Arrow":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Arrow = __gong__map_Arrow[targetIdentifier]
				case "Handbell":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Handbell = __gong__map_Handbell[targetIdentifier]
				case "Brass_bend":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Brass_bend = __gong__map_Empty_placement[targetIdentifier]
				case "Flip":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Flip = __gong__map_Empty_placement[targetIdentifier]
				case "Smear":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Smear = __gong__map_Empty_placement[targetIdentifier]
				case "Open":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Open = __gong__map_Empty_placement_smufl[targetIdentifier]
				case "Half_muted":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Half_muted = __gong__map_Empty_placement_smufl[targetIdentifier]
				case "Harmon_mute":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Harmon_mute = __gong__map_Harmon_mute[targetIdentifier]
				case "Golpe":
					targetIdentifier := ident.Name
					__gong__map_Technical[identifier].Golpe = __gong__map_Empty_placement[targetIdentifier]
				}
			case "Text_element_data":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tie":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tied":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Time":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Time_modification":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Timpani":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Transpose":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tremolo":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tuplet":
				switch fieldName {
				// insertion point for field dependant code
				case "Tuplet_actual":
					targetIdentifier := ident.Name
					__gong__map_Tuplet[identifier].Tuplet_actual = __gong__map_Tuplet_portion[targetIdentifier]
				case "Tuplet_normal":
					targetIdentifier := ident.Name
					__gong__map_Tuplet[identifier].Tuplet_normal = __gong__map_Tuplet_portion[targetIdentifier]
				}
			case "Tuplet_dot":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tuplet_number":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tuplet_portion":
				switch fieldName {
				// insertion point for field dependant code
				case "Tuplet_number":
					targetIdentifier := ident.Name
					__gong__map_Tuplet_portion[identifier].Tuplet_number = __gong__map_Tuplet_number[targetIdentifier]
				case "Tuplet_type":
					targetIdentifier := ident.Name
					__gong__map_Tuplet_portion[identifier].Tuplet_type = __gong__map_Tuplet_type[targetIdentifier]
				}
			case "Tuplet_type":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Typed_text":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Unpitched":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Virtual_instrument":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Wait":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Wavy_line":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Wedge":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Wood":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Work":
				switch fieldName {
				// insertion point for field dependant code
				case "Opus":
					targetIdentifier := ident.Name
					__gong__map_Work[identifier].Opus = __gong__map_Opus[targetIdentifier]
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "Accidental":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Accidental_mark":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Accidental_text":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Accord":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Accordion_registration":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "AnyType":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Appearance":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Arpeggiate":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Arrow":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Articulations":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Assess":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Attributes":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Backup":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Bar_style_color":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Barline":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Barre":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Bass":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Bass_step":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Beam":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Beat_repeat":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Beat_unit_tied":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Beater":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Bend":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Bookmark":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Bracket":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Breath_mark":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Caesura":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Cancel":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Clef":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Coda":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Credit":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Dashes":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Defaults":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Degree":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Degree_alter":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Degree_type":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Degree_value":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Direction":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Direction_type":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Distance":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Double":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Dynamics":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Effect":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Elision":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty_font":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty_line":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty_placement":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty_placement_smufl":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty_print_object_style_align":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty_print_style":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty_print_style_align":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty_print_style_align_id":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Empty_trill_sound":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Encoding":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Ending":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Extend":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Feature":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Fermata":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Figure":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Figured_bass":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Fingering":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "First_fret":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Foo":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "For_part":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Formatted_symbol":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Formatted_symbol_id":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Forward":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Frame":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Frame_note":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Fret":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Glass":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Glissando":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Glyph":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Grace":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Group_barline":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Group_symbol":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Grouping":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Hammer_on_pull_off":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Handbell":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Harmon_closed":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Harmon_mute":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Harmonic":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Harmony":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Harmony_alter":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Harp_pedals":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Heel_toe":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Hole":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Hole_closed":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Horizontal_turn":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Identification":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Image":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Instrument":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Instrument_change":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Instrument_link":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Interchangeable":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Inversion":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Key":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Key_accidental":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Key_octave":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Kind":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Level":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Line_detail":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Line_width":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Link":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Listen":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Listening":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Lyric":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Lyric_font":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Lyric_language":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Measure_layout":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Measure_numbering":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Measure_repeat":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Measure_style":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Membrane":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Metal":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Metronome":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Metronome_beam":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Metronome_note":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Metronome_tied":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Metronome_tuplet":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Midi_device":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Midi_instrument":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Miscellaneous":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Miscellaneous_field":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Mordent":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Multiple_rest":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Name_display":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Non_arpeggiate":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Notations":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Note":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Note_size":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Note_type":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Notehead":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Notehead_text":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Numeral":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Numeral_key":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Numeral_root":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Octave_shift":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Offset":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Opus":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Ornaments":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Other_appearance":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Other_listening":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Other_notation":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Other_play":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Page_layout":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Page_margins":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Part_clef":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Part_group":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Part_link":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Part_list":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Part_symbol":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Part_transpose":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Pedal":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Pedal_tuning":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Percussion":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Pitch":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Pitched":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Play":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Player":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Principal_voice":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Print":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Release":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Repeat":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Rest":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Root":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Root_step":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Scaling":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Scordatura":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Score_instrument":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Score_part":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Score_partwise":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Score_timewise":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Segno":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Slash":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Slide":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Slur":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Sound":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Staff_details":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Staff_divide":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Staff_layout":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Staff_size":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Staff_tuning":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Stem":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Stick":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "String_mute":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Strong_accent":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Supports":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Swing":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Sync":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "System_dividers":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "System_layout":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "System_margins":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Tap":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Technical":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Text_element_data":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Tie":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Tied":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Time":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Time_modification":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Timpani":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Transpose":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Tremolo":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Tuplet":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Tuplet_dot":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Tuplet_number":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Tuplet_portion":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Tuplet_type":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Typed_text":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Unpitched":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Virtual_instrument":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Wait":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Wavy_line":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Wedge":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Wood":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Work":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}

// ReplaceOldDeclarationsInFile replaces specific text in a file at the given path.
func ReplaceOldDeclarationsInFile(pathToFile string) error {
	// Open the file for reading.
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// replacing function with Injection
	pattern := regexp.MustCompile(`\b\w*Injection\b`)
	pattern2 := regexp.MustCompile(`\bmap_DocLink_Identifier_\w*\b`)

	// Temporary slice to hold lines from the file.
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Replace the target text with the desired text.
		line := strings.Replace(scanner.Text(), "var ___dummy__Time_stage time.Time", "var _ time.Time", -1)
		line = pattern.ReplaceAllString(line, "_")
		line = pattern2.ReplaceAllString(line, "_")

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Re-open the file for writing.
	file, err = os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the modified lines back to the file.
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
