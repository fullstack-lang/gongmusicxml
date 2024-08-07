# gongmusicxml

## fetch the xsd

## generate fo code from xsd

> gowsdl -d go -p models -o models.go musicxml.wsdl


The generated file is 4370 lines long

but there are compilation issues


### issue with import of "context"

add 

```go
var _ context.Context
```

### issue with first string const

```go
type Above_below string

const (
	Above_belowAbove Above_below = "above"

	Above_belowBelow Above_below = "below"
)
```

```
invalid constant type Above_belowcompilerInvalidConstType
type Above_below string
```

The problem is AFTER in the code

It is a problem with missing types definitions. Correction in types.go

```go
package models

type PositiveInteger int
type NMTOKEN string
```

### issue with too many number of enums (cannot handle more than 100 union terms (implementation limitation)


```go
type GongstructEnumStringField interface {
	string | Above_below | Accidental_value | AnyURI | Arrow_direction | Arrow_style | Backward_forward | Bar_style | Beam_value | Beater_value | Bend_shape | Breath_mark_value | Caesura_value | Cancel_location | Circular_arrow | Clef_sign | Color | Comma_separated_text | Css_font_size | Degree_symbol_value | Degree_type_value | Distance_type | Effect_value | Enclosure_shape | Ending_number | Fan | Fermata_shape | Font_size | Font_style | Font_weight | Glass_value | Glyph_type | Group_barline_value | Group_symbol_value | Handbell_value | Harmon_closed_location | Harmon_closed_value | Harmony_arrangement | Harmony_type | Hole_closed_location | Hole_closed_value | Kind_value | Left_center_right | Left_right | Line_end | Line_length | Line_shape | Line_type | Line_width_type | Margin_type | Measure_numbering_value | Measure_text | Membrane_value | Metal_value | Mode | Mute | NCName | NMTOKEN | Note_size_type | Note_type_value | Notehead_value | Number_or_normal | Numeral_mode | On_off | Over_under | Pedal_type | Pitched_value | Positive_integer_or_empty | Principal_voice_symbol | Right_left_middle | Semi_pitched | Show_frets | Show_tuplet | Staff_divide_symbol | Staff_type | Start_note | Start_stop | Start_stop_change_continue | Start_stop_continue | Start_stop_discontinue | Start_stop_single | Stem_value | Step | Stick_location | Stick_material | Stick_type | Syllabic | Symbol_size | Sync_type | System_relation_number | Tap_hand | Text_direction | Tied_type | Time_only | Time_relation | Time_separator | Time_symbol | Tip_direction | Top_bottom | Tremolo_type | Trill_step | Two_note_turn | Up_down | Up_down_stop_continue | Upright_inverted | Valign | Valign_image | Wedge_type | Winged | Wood_value | Yes_no | Yes_no_number
	Codes() []string
	CodeValues() []string
}
```

1. Removing all enums

Problem

go/probe/enum_type_to_form.go

patch to gongc OK (505)    

## compiling the struct

### issue with undefined: IDREF compilerUndeclaredName

missing def

### issue with soap.Datetime

```go
	// Time signatures are represented by the beats element for the
	// numerator and the beat-type element for the denominator.
	Time []soap.XSDTime `xml:"time,omitempty" json:"time,omitempty"`
```

### issue with redeclarations

```go
	Segno *Segno `xml:"segno,omitempty" json:"segno,omitempty"`

	Coda *Coda `xml:"coda,omitempty" json:"coda,omitempty"`

	Segno string `xml:"segno,attr,omitempty" json:"segno,omitempty"`

	Coda string `xml:"coda,attr,omitempty" json:"coda,omitempty"`
```

### general issue

hard to pin down

there a 2296 lines in structs.go

bisect

1000

```go
type string struct {
	Value *String_number `xml:",chardata" json:"-,"`
}
```

Notice that 6 structs have already a Name attributes

### issue with 