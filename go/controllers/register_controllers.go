// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fullstack-lang/gongmusicxml/go/orm"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongmusicxml/go")
	{ // insertion point for registrations
		v1.GET("/v1/accidentals", GetController().GetAccidentals)
		v1.GET("/v1/accidentals/:id", GetController().GetAccidental)
		v1.POST("/v1/accidentals", GetController().PostAccidental)
		v1.PATCH("/v1/accidentals/:id", GetController().UpdateAccidental)
		v1.PUT("/v1/accidentals/:id", GetController().UpdateAccidental)
		v1.DELETE("/v1/accidentals/:id", GetController().DeleteAccidental)

		v1.GET("/v1/accidental_marks", GetController().GetAccidental_marks)
		v1.GET("/v1/accidental_marks/:id", GetController().GetAccidental_mark)
		v1.POST("/v1/accidental_marks", GetController().PostAccidental_mark)
		v1.PATCH("/v1/accidental_marks/:id", GetController().UpdateAccidental_mark)
		v1.PUT("/v1/accidental_marks/:id", GetController().UpdateAccidental_mark)
		v1.DELETE("/v1/accidental_marks/:id", GetController().DeleteAccidental_mark)

		v1.GET("/v1/accidental_texts", GetController().GetAccidental_texts)
		v1.GET("/v1/accidental_texts/:id", GetController().GetAccidental_text)
		v1.POST("/v1/accidental_texts", GetController().PostAccidental_text)
		v1.PATCH("/v1/accidental_texts/:id", GetController().UpdateAccidental_text)
		v1.PUT("/v1/accidental_texts/:id", GetController().UpdateAccidental_text)
		v1.DELETE("/v1/accidental_texts/:id", GetController().DeleteAccidental_text)

		v1.GET("/v1/accords", GetController().GetAccords)
		v1.GET("/v1/accords/:id", GetController().GetAccord)
		v1.POST("/v1/accords", GetController().PostAccord)
		v1.PATCH("/v1/accords/:id", GetController().UpdateAccord)
		v1.PUT("/v1/accords/:id", GetController().UpdateAccord)
		v1.DELETE("/v1/accords/:id", GetController().DeleteAccord)

		v1.GET("/v1/accordion_registrations", GetController().GetAccordion_registrations)
		v1.GET("/v1/accordion_registrations/:id", GetController().GetAccordion_registration)
		v1.POST("/v1/accordion_registrations", GetController().PostAccordion_registration)
		v1.PATCH("/v1/accordion_registrations/:id", GetController().UpdateAccordion_registration)
		v1.PUT("/v1/accordion_registrations/:id", GetController().UpdateAccordion_registration)
		v1.DELETE("/v1/accordion_registrations/:id", GetController().DeleteAccordion_registration)

		v1.GET("/v1/anytypes", GetController().GetAnyTypes)
		v1.GET("/v1/anytypes/:id", GetController().GetAnyType)
		v1.POST("/v1/anytypes", GetController().PostAnyType)
		v1.PATCH("/v1/anytypes/:id", GetController().UpdateAnyType)
		v1.PUT("/v1/anytypes/:id", GetController().UpdateAnyType)
		v1.DELETE("/v1/anytypes/:id", GetController().DeleteAnyType)

		v1.GET("/v1/appearances", GetController().GetAppearances)
		v1.GET("/v1/appearances/:id", GetController().GetAppearance)
		v1.POST("/v1/appearances", GetController().PostAppearance)
		v1.PATCH("/v1/appearances/:id", GetController().UpdateAppearance)
		v1.PUT("/v1/appearances/:id", GetController().UpdateAppearance)
		v1.DELETE("/v1/appearances/:id", GetController().DeleteAppearance)

		v1.GET("/v1/arpeggiates", GetController().GetArpeggiates)
		v1.GET("/v1/arpeggiates/:id", GetController().GetArpeggiate)
		v1.POST("/v1/arpeggiates", GetController().PostArpeggiate)
		v1.PATCH("/v1/arpeggiates/:id", GetController().UpdateArpeggiate)
		v1.PUT("/v1/arpeggiates/:id", GetController().UpdateArpeggiate)
		v1.DELETE("/v1/arpeggiates/:id", GetController().DeleteArpeggiate)

		v1.GET("/v1/arrows", GetController().GetArrows)
		v1.GET("/v1/arrows/:id", GetController().GetArrow)
		v1.POST("/v1/arrows", GetController().PostArrow)
		v1.PATCH("/v1/arrows/:id", GetController().UpdateArrow)
		v1.PUT("/v1/arrows/:id", GetController().UpdateArrow)
		v1.DELETE("/v1/arrows/:id", GetController().DeleteArrow)

		v1.GET("/v1/articulationss", GetController().GetArticulationss)
		v1.GET("/v1/articulationss/:id", GetController().GetArticulations)
		v1.POST("/v1/articulationss", GetController().PostArticulations)
		v1.PATCH("/v1/articulationss/:id", GetController().UpdateArticulations)
		v1.PUT("/v1/articulationss/:id", GetController().UpdateArticulations)
		v1.DELETE("/v1/articulationss/:id", GetController().DeleteArticulations)

		v1.GET("/v1/assesss", GetController().GetAssesss)
		v1.GET("/v1/assesss/:id", GetController().GetAssess)
		v1.POST("/v1/assesss", GetController().PostAssess)
		v1.PATCH("/v1/assesss/:id", GetController().UpdateAssess)
		v1.PUT("/v1/assesss/:id", GetController().UpdateAssess)
		v1.DELETE("/v1/assesss/:id", GetController().DeleteAssess)

		v1.GET("/v1/attributess", GetController().GetAttributess)
		v1.GET("/v1/attributess/:id", GetController().GetAttributes)
		v1.POST("/v1/attributess", GetController().PostAttributes)
		v1.PATCH("/v1/attributess/:id", GetController().UpdateAttributes)
		v1.PUT("/v1/attributess/:id", GetController().UpdateAttributes)
		v1.DELETE("/v1/attributess/:id", GetController().DeleteAttributes)

		v1.GET("/v1/backups", GetController().GetBackups)
		v1.GET("/v1/backups/:id", GetController().GetBackup)
		v1.POST("/v1/backups", GetController().PostBackup)
		v1.PATCH("/v1/backups/:id", GetController().UpdateBackup)
		v1.PUT("/v1/backups/:id", GetController().UpdateBackup)
		v1.DELETE("/v1/backups/:id", GetController().DeleteBackup)

		v1.GET("/v1/bar_style_colors", GetController().GetBar_style_colors)
		v1.GET("/v1/bar_style_colors/:id", GetController().GetBar_style_color)
		v1.POST("/v1/bar_style_colors", GetController().PostBar_style_color)
		v1.PATCH("/v1/bar_style_colors/:id", GetController().UpdateBar_style_color)
		v1.PUT("/v1/bar_style_colors/:id", GetController().UpdateBar_style_color)
		v1.DELETE("/v1/bar_style_colors/:id", GetController().DeleteBar_style_color)

		v1.GET("/v1/barlines", GetController().GetBarlines)
		v1.GET("/v1/barlines/:id", GetController().GetBarline)
		v1.POST("/v1/barlines", GetController().PostBarline)
		v1.PATCH("/v1/barlines/:id", GetController().UpdateBarline)
		v1.PUT("/v1/barlines/:id", GetController().UpdateBarline)
		v1.DELETE("/v1/barlines/:id", GetController().DeleteBarline)

		v1.GET("/v1/barres", GetController().GetBarres)
		v1.GET("/v1/barres/:id", GetController().GetBarre)
		v1.POST("/v1/barres", GetController().PostBarre)
		v1.PATCH("/v1/barres/:id", GetController().UpdateBarre)
		v1.PUT("/v1/barres/:id", GetController().UpdateBarre)
		v1.DELETE("/v1/barres/:id", GetController().DeleteBarre)

		v1.GET("/v1/basss", GetController().GetBasss)
		v1.GET("/v1/basss/:id", GetController().GetBass)
		v1.POST("/v1/basss", GetController().PostBass)
		v1.PATCH("/v1/basss/:id", GetController().UpdateBass)
		v1.PUT("/v1/basss/:id", GetController().UpdateBass)
		v1.DELETE("/v1/basss/:id", GetController().DeleteBass)

		v1.GET("/v1/bass_steps", GetController().GetBass_steps)
		v1.GET("/v1/bass_steps/:id", GetController().GetBass_step)
		v1.POST("/v1/bass_steps", GetController().PostBass_step)
		v1.PATCH("/v1/bass_steps/:id", GetController().UpdateBass_step)
		v1.PUT("/v1/bass_steps/:id", GetController().UpdateBass_step)
		v1.DELETE("/v1/bass_steps/:id", GetController().DeleteBass_step)

		v1.GET("/v1/beams", GetController().GetBeams)
		v1.GET("/v1/beams/:id", GetController().GetBeam)
		v1.POST("/v1/beams", GetController().PostBeam)
		v1.PATCH("/v1/beams/:id", GetController().UpdateBeam)
		v1.PUT("/v1/beams/:id", GetController().UpdateBeam)
		v1.DELETE("/v1/beams/:id", GetController().DeleteBeam)

		v1.GET("/v1/beat_repeats", GetController().GetBeat_repeats)
		v1.GET("/v1/beat_repeats/:id", GetController().GetBeat_repeat)
		v1.POST("/v1/beat_repeats", GetController().PostBeat_repeat)
		v1.PATCH("/v1/beat_repeats/:id", GetController().UpdateBeat_repeat)
		v1.PUT("/v1/beat_repeats/:id", GetController().UpdateBeat_repeat)
		v1.DELETE("/v1/beat_repeats/:id", GetController().DeleteBeat_repeat)

		v1.GET("/v1/beat_unit_tieds", GetController().GetBeat_unit_tieds)
		v1.GET("/v1/beat_unit_tieds/:id", GetController().GetBeat_unit_tied)
		v1.POST("/v1/beat_unit_tieds", GetController().PostBeat_unit_tied)
		v1.PATCH("/v1/beat_unit_tieds/:id", GetController().UpdateBeat_unit_tied)
		v1.PUT("/v1/beat_unit_tieds/:id", GetController().UpdateBeat_unit_tied)
		v1.DELETE("/v1/beat_unit_tieds/:id", GetController().DeleteBeat_unit_tied)

		v1.GET("/v1/beaters", GetController().GetBeaters)
		v1.GET("/v1/beaters/:id", GetController().GetBeater)
		v1.POST("/v1/beaters", GetController().PostBeater)
		v1.PATCH("/v1/beaters/:id", GetController().UpdateBeater)
		v1.PUT("/v1/beaters/:id", GetController().UpdateBeater)
		v1.DELETE("/v1/beaters/:id", GetController().DeleteBeater)

		v1.GET("/v1/bends", GetController().GetBends)
		v1.GET("/v1/bends/:id", GetController().GetBend)
		v1.POST("/v1/bends", GetController().PostBend)
		v1.PATCH("/v1/bends/:id", GetController().UpdateBend)
		v1.PUT("/v1/bends/:id", GetController().UpdateBend)
		v1.DELETE("/v1/bends/:id", GetController().DeleteBend)

		v1.GET("/v1/bookmarks", GetController().GetBookmarks)
		v1.GET("/v1/bookmarks/:id", GetController().GetBookmark)
		v1.POST("/v1/bookmarks", GetController().PostBookmark)
		v1.PATCH("/v1/bookmarks/:id", GetController().UpdateBookmark)
		v1.PUT("/v1/bookmarks/:id", GetController().UpdateBookmark)
		v1.DELETE("/v1/bookmarks/:id", GetController().DeleteBookmark)

		v1.GET("/v1/brackets", GetController().GetBrackets)
		v1.GET("/v1/brackets/:id", GetController().GetBracket)
		v1.POST("/v1/brackets", GetController().PostBracket)
		v1.PATCH("/v1/brackets/:id", GetController().UpdateBracket)
		v1.PUT("/v1/brackets/:id", GetController().UpdateBracket)
		v1.DELETE("/v1/brackets/:id", GetController().DeleteBracket)

		v1.GET("/v1/breath_marks", GetController().GetBreath_marks)
		v1.GET("/v1/breath_marks/:id", GetController().GetBreath_mark)
		v1.POST("/v1/breath_marks", GetController().PostBreath_mark)
		v1.PATCH("/v1/breath_marks/:id", GetController().UpdateBreath_mark)
		v1.PUT("/v1/breath_marks/:id", GetController().UpdateBreath_mark)
		v1.DELETE("/v1/breath_marks/:id", GetController().DeleteBreath_mark)

		v1.GET("/v1/caesuras", GetController().GetCaesuras)
		v1.GET("/v1/caesuras/:id", GetController().GetCaesura)
		v1.POST("/v1/caesuras", GetController().PostCaesura)
		v1.PATCH("/v1/caesuras/:id", GetController().UpdateCaesura)
		v1.PUT("/v1/caesuras/:id", GetController().UpdateCaesura)
		v1.DELETE("/v1/caesuras/:id", GetController().DeleteCaesura)

		v1.GET("/v1/cancels", GetController().GetCancels)
		v1.GET("/v1/cancels/:id", GetController().GetCancel)
		v1.POST("/v1/cancels", GetController().PostCancel)
		v1.PATCH("/v1/cancels/:id", GetController().UpdateCancel)
		v1.PUT("/v1/cancels/:id", GetController().UpdateCancel)
		v1.DELETE("/v1/cancels/:id", GetController().DeleteCancel)

		v1.GET("/v1/clefs", GetController().GetClefs)
		v1.GET("/v1/clefs/:id", GetController().GetClef)
		v1.POST("/v1/clefs", GetController().PostClef)
		v1.PATCH("/v1/clefs/:id", GetController().UpdateClef)
		v1.PUT("/v1/clefs/:id", GetController().UpdateClef)
		v1.DELETE("/v1/clefs/:id", GetController().DeleteClef)

		v1.GET("/v1/codas", GetController().GetCodas)
		v1.GET("/v1/codas/:id", GetController().GetCoda)
		v1.POST("/v1/codas", GetController().PostCoda)
		v1.PATCH("/v1/codas/:id", GetController().UpdateCoda)
		v1.PUT("/v1/codas/:id", GetController().UpdateCoda)
		v1.DELETE("/v1/codas/:id", GetController().DeleteCoda)

		v1.GET("/v1/credits", GetController().GetCredits)
		v1.GET("/v1/credits/:id", GetController().GetCredit)
		v1.POST("/v1/credits", GetController().PostCredit)
		v1.PATCH("/v1/credits/:id", GetController().UpdateCredit)
		v1.PUT("/v1/credits/:id", GetController().UpdateCredit)
		v1.DELETE("/v1/credits/:id", GetController().DeleteCredit)

		v1.GET("/v1/dashess", GetController().GetDashess)
		v1.GET("/v1/dashess/:id", GetController().GetDashes)
		v1.POST("/v1/dashess", GetController().PostDashes)
		v1.PATCH("/v1/dashess/:id", GetController().UpdateDashes)
		v1.PUT("/v1/dashess/:id", GetController().UpdateDashes)
		v1.DELETE("/v1/dashess/:id", GetController().DeleteDashes)

		v1.GET("/v1/defaultss", GetController().GetDefaultss)
		v1.GET("/v1/defaultss/:id", GetController().GetDefaults)
		v1.POST("/v1/defaultss", GetController().PostDefaults)
		v1.PATCH("/v1/defaultss/:id", GetController().UpdateDefaults)
		v1.PUT("/v1/defaultss/:id", GetController().UpdateDefaults)
		v1.DELETE("/v1/defaultss/:id", GetController().DeleteDefaults)

		v1.GET("/v1/degrees", GetController().GetDegrees)
		v1.GET("/v1/degrees/:id", GetController().GetDegree)
		v1.POST("/v1/degrees", GetController().PostDegree)
		v1.PATCH("/v1/degrees/:id", GetController().UpdateDegree)
		v1.PUT("/v1/degrees/:id", GetController().UpdateDegree)
		v1.DELETE("/v1/degrees/:id", GetController().DeleteDegree)

		v1.GET("/v1/degree_alters", GetController().GetDegree_alters)
		v1.GET("/v1/degree_alters/:id", GetController().GetDegree_alter)
		v1.POST("/v1/degree_alters", GetController().PostDegree_alter)
		v1.PATCH("/v1/degree_alters/:id", GetController().UpdateDegree_alter)
		v1.PUT("/v1/degree_alters/:id", GetController().UpdateDegree_alter)
		v1.DELETE("/v1/degree_alters/:id", GetController().DeleteDegree_alter)

		v1.GET("/v1/degree_types", GetController().GetDegree_types)
		v1.GET("/v1/degree_types/:id", GetController().GetDegree_type)
		v1.POST("/v1/degree_types", GetController().PostDegree_type)
		v1.PATCH("/v1/degree_types/:id", GetController().UpdateDegree_type)
		v1.PUT("/v1/degree_types/:id", GetController().UpdateDegree_type)
		v1.DELETE("/v1/degree_types/:id", GetController().DeleteDegree_type)

		v1.GET("/v1/degree_values", GetController().GetDegree_values)
		v1.GET("/v1/degree_values/:id", GetController().GetDegree_value)
		v1.POST("/v1/degree_values", GetController().PostDegree_value)
		v1.PATCH("/v1/degree_values/:id", GetController().UpdateDegree_value)
		v1.PUT("/v1/degree_values/:id", GetController().UpdateDegree_value)
		v1.DELETE("/v1/degree_values/:id", GetController().DeleteDegree_value)

		v1.GET("/v1/directions", GetController().GetDirections)
		v1.GET("/v1/directions/:id", GetController().GetDirection)
		v1.POST("/v1/directions", GetController().PostDirection)
		v1.PATCH("/v1/directions/:id", GetController().UpdateDirection)
		v1.PUT("/v1/directions/:id", GetController().UpdateDirection)
		v1.DELETE("/v1/directions/:id", GetController().DeleteDirection)

		v1.GET("/v1/direction_types", GetController().GetDirection_types)
		v1.GET("/v1/direction_types/:id", GetController().GetDirection_type)
		v1.POST("/v1/direction_types", GetController().PostDirection_type)
		v1.PATCH("/v1/direction_types/:id", GetController().UpdateDirection_type)
		v1.PUT("/v1/direction_types/:id", GetController().UpdateDirection_type)
		v1.DELETE("/v1/direction_types/:id", GetController().DeleteDirection_type)

		v1.GET("/v1/distances", GetController().GetDistances)
		v1.GET("/v1/distances/:id", GetController().GetDistance)
		v1.POST("/v1/distances", GetController().PostDistance)
		v1.PATCH("/v1/distances/:id", GetController().UpdateDistance)
		v1.PUT("/v1/distances/:id", GetController().UpdateDistance)
		v1.DELETE("/v1/distances/:id", GetController().DeleteDistance)

		v1.GET("/v1/doubles", GetController().GetDoubles)
		v1.GET("/v1/doubles/:id", GetController().GetDouble)
		v1.POST("/v1/doubles", GetController().PostDouble)
		v1.PATCH("/v1/doubles/:id", GetController().UpdateDouble)
		v1.PUT("/v1/doubles/:id", GetController().UpdateDouble)
		v1.DELETE("/v1/doubles/:id", GetController().DeleteDouble)

		v1.GET("/v1/dynamicss", GetController().GetDynamicss)
		v1.GET("/v1/dynamicss/:id", GetController().GetDynamics)
		v1.POST("/v1/dynamicss", GetController().PostDynamics)
		v1.PATCH("/v1/dynamicss/:id", GetController().UpdateDynamics)
		v1.PUT("/v1/dynamicss/:id", GetController().UpdateDynamics)
		v1.DELETE("/v1/dynamicss/:id", GetController().DeleteDynamics)

		v1.GET("/v1/effects", GetController().GetEffects)
		v1.GET("/v1/effects/:id", GetController().GetEffect)
		v1.POST("/v1/effects", GetController().PostEffect)
		v1.PATCH("/v1/effects/:id", GetController().UpdateEffect)
		v1.PUT("/v1/effects/:id", GetController().UpdateEffect)
		v1.DELETE("/v1/effects/:id", GetController().DeleteEffect)

		v1.GET("/v1/elisions", GetController().GetElisions)
		v1.GET("/v1/elisions/:id", GetController().GetElision)
		v1.POST("/v1/elisions", GetController().PostElision)
		v1.PATCH("/v1/elisions/:id", GetController().UpdateElision)
		v1.PUT("/v1/elisions/:id", GetController().UpdateElision)
		v1.DELETE("/v1/elisions/:id", GetController().DeleteElision)

		v1.GET("/v1/emptys", GetController().GetEmptys)
		v1.GET("/v1/emptys/:id", GetController().GetEmpty)
		v1.POST("/v1/emptys", GetController().PostEmpty)
		v1.PATCH("/v1/emptys/:id", GetController().UpdateEmpty)
		v1.PUT("/v1/emptys/:id", GetController().UpdateEmpty)
		v1.DELETE("/v1/emptys/:id", GetController().DeleteEmpty)

		v1.GET("/v1/empty_fonts", GetController().GetEmpty_fonts)
		v1.GET("/v1/empty_fonts/:id", GetController().GetEmpty_font)
		v1.POST("/v1/empty_fonts", GetController().PostEmpty_font)
		v1.PATCH("/v1/empty_fonts/:id", GetController().UpdateEmpty_font)
		v1.PUT("/v1/empty_fonts/:id", GetController().UpdateEmpty_font)
		v1.DELETE("/v1/empty_fonts/:id", GetController().DeleteEmpty_font)

		v1.GET("/v1/empty_lines", GetController().GetEmpty_lines)
		v1.GET("/v1/empty_lines/:id", GetController().GetEmpty_line)
		v1.POST("/v1/empty_lines", GetController().PostEmpty_line)
		v1.PATCH("/v1/empty_lines/:id", GetController().UpdateEmpty_line)
		v1.PUT("/v1/empty_lines/:id", GetController().UpdateEmpty_line)
		v1.DELETE("/v1/empty_lines/:id", GetController().DeleteEmpty_line)

		v1.GET("/v1/empty_placements", GetController().GetEmpty_placements)
		v1.GET("/v1/empty_placements/:id", GetController().GetEmpty_placement)
		v1.POST("/v1/empty_placements", GetController().PostEmpty_placement)
		v1.PATCH("/v1/empty_placements/:id", GetController().UpdateEmpty_placement)
		v1.PUT("/v1/empty_placements/:id", GetController().UpdateEmpty_placement)
		v1.DELETE("/v1/empty_placements/:id", GetController().DeleteEmpty_placement)

		v1.GET("/v1/empty_placement_smufls", GetController().GetEmpty_placement_smufls)
		v1.GET("/v1/empty_placement_smufls/:id", GetController().GetEmpty_placement_smufl)
		v1.POST("/v1/empty_placement_smufls", GetController().PostEmpty_placement_smufl)
		v1.PATCH("/v1/empty_placement_smufls/:id", GetController().UpdateEmpty_placement_smufl)
		v1.PUT("/v1/empty_placement_smufls/:id", GetController().UpdateEmpty_placement_smufl)
		v1.DELETE("/v1/empty_placement_smufls/:id", GetController().DeleteEmpty_placement_smufl)

		v1.GET("/v1/empty_print_object_style_aligns", GetController().GetEmpty_print_object_style_aligns)
		v1.GET("/v1/empty_print_object_style_aligns/:id", GetController().GetEmpty_print_object_style_align)
		v1.POST("/v1/empty_print_object_style_aligns", GetController().PostEmpty_print_object_style_align)
		v1.PATCH("/v1/empty_print_object_style_aligns/:id", GetController().UpdateEmpty_print_object_style_align)
		v1.PUT("/v1/empty_print_object_style_aligns/:id", GetController().UpdateEmpty_print_object_style_align)
		v1.DELETE("/v1/empty_print_object_style_aligns/:id", GetController().DeleteEmpty_print_object_style_align)

		v1.GET("/v1/empty_print_styles", GetController().GetEmpty_print_styles)
		v1.GET("/v1/empty_print_styles/:id", GetController().GetEmpty_print_style)
		v1.POST("/v1/empty_print_styles", GetController().PostEmpty_print_style)
		v1.PATCH("/v1/empty_print_styles/:id", GetController().UpdateEmpty_print_style)
		v1.PUT("/v1/empty_print_styles/:id", GetController().UpdateEmpty_print_style)
		v1.DELETE("/v1/empty_print_styles/:id", GetController().DeleteEmpty_print_style)

		v1.GET("/v1/empty_print_style_aligns", GetController().GetEmpty_print_style_aligns)
		v1.GET("/v1/empty_print_style_aligns/:id", GetController().GetEmpty_print_style_align)
		v1.POST("/v1/empty_print_style_aligns", GetController().PostEmpty_print_style_align)
		v1.PATCH("/v1/empty_print_style_aligns/:id", GetController().UpdateEmpty_print_style_align)
		v1.PUT("/v1/empty_print_style_aligns/:id", GetController().UpdateEmpty_print_style_align)
		v1.DELETE("/v1/empty_print_style_aligns/:id", GetController().DeleteEmpty_print_style_align)

		v1.GET("/v1/empty_print_style_align_ids", GetController().GetEmpty_print_style_align_ids)
		v1.GET("/v1/empty_print_style_align_ids/:id", GetController().GetEmpty_print_style_align_id)
		v1.POST("/v1/empty_print_style_align_ids", GetController().PostEmpty_print_style_align_id)
		v1.PATCH("/v1/empty_print_style_align_ids/:id", GetController().UpdateEmpty_print_style_align_id)
		v1.PUT("/v1/empty_print_style_align_ids/:id", GetController().UpdateEmpty_print_style_align_id)
		v1.DELETE("/v1/empty_print_style_align_ids/:id", GetController().DeleteEmpty_print_style_align_id)

		v1.GET("/v1/empty_trill_sounds", GetController().GetEmpty_trill_sounds)
		v1.GET("/v1/empty_trill_sounds/:id", GetController().GetEmpty_trill_sound)
		v1.POST("/v1/empty_trill_sounds", GetController().PostEmpty_trill_sound)
		v1.PATCH("/v1/empty_trill_sounds/:id", GetController().UpdateEmpty_trill_sound)
		v1.PUT("/v1/empty_trill_sounds/:id", GetController().UpdateEmpty_trill_sound)
		v1.DELETE("/v1/empty_trill_sounds/:id", GetController().DeleteEmpty_trill_sound)

		v1.GET("/v1/encodings", GetController().GetEncodings)
		v1.GET("/v1/encodings/:id", GetController().GetEncoding)
		v1.POST("/v1/encodings", GetController().PostEncoding)
		v1.PATCH("/v1/encodings/:id", GetController().UpdateEncoding)
		v1.PUT("/v1/encodings/:id", GetController().UpdateEncoding)
		v1.DELETE("/v1/encodings/:id", GetController().DeleteEncoding)

		v1.GET("/v1/endings", GetController().GetEndings)
		v1.GET("/v1/endings/:id", GetController().GetEnding)
		v1.POST("/v1/endings", GetController().PostEnding)
		v1.PATCH("/v1/endings/:id", GetController().UpdateEnding)
		v1.PUT("/v1/endings/:id", GetController().UpdateEnding)
		v1.DELETE("/v1/endings/:id", GetController().DeleteEnding)

		v1.GET("/v1/extends", GetController().GetExtends)
		v1.GET("/v1/extends/:id", GetController().GetExtend)
		v1.POST("/v1/extends", GetController().PostExtend)
		v1.PATCH("/v1/extends/:id", GetController().UpdateExtend)
		v1.PUT("/v1/extends/:id", GetController().UpdateExtend)
		v1.DELETE("/v1/extends/:id", GetController().DeleteExtend)

		v1.GET("/v1/features", GetController().GetFeatures)
		v1.GET("/v1/features/:id", GetController().GetFeature)
		v1.POST("/v1/features", GetController().PostFeature)
		v1.PATCH("/v1/features/:id", GetController().UpdateFeature)
		v1.PUT("/v1/features/:id", GetController().UpdateFeature)
		v1.DELETE("/v1/features/:id", GetController().DeleteFeature)

		v1.GET("/v1/fermatas", GetController().GetFermatas)
		v1.GET("/v1/fermatas/:id", GetController().GetFermata)
		v1.POST("/v1/fermatas", GetController().PostFermata)
		v1.PATCH("/v1/fermatas/:id", GetController().UpdateFermata)
		v1.PUT("/v1/fermatas/:id", GetController().UpdateFermata)
		v1.DELETE("/v1/fermatas/:id", GetController().DeleteFermata)

		v1.GET("/v1/figures", GetController().GetFigures)
		v1.GET("/v1/figures/:id", GetController().GetFigure)
		v1.POST("/v1/figures", GetController().PostFigure)
		v1.PATCH("/v1/figures/:id", GetController().UpdateFigure)
		v1.PUT("/v1/figures/:id", GetController().UpdateFigure)
		v1.DELETE("/v1/figures/:id", GetController().DeleteFigure)

		v1.GET("/v1/figured_basss", GetController().GetFigured_basss)
		v1.GET("/v1/figured_basss/:id", GetController().GetFigured_bass)
		v1.POST("/v1/figured_basss", GetController().PostFigured_bass)
		v1.PATCH("/v1/figured_basss/:id", GetController().UpdateFigured_bass)
		v1.PUT("/v1/figured_basss/:id", GetController().UpdateFigured_bass)
		v1.DELETE("/v1/figured_basss/:id", GetController().DeleteFigured_bass)

		v1.GET("/v1/fingerings", GetController().GetFingerings)
		v1.GET("/v1/fingerings/:id", GetController().GetFingering)
		v1.POST("/v1/fingerings", GetController().PostFingering)
		v1.PATCH("/v1/fingerings/:id", GetController().UpdateFingering)
		v1.PUT("/v1/fingerings/:id", GetController().UpdateFingering)
		v1.DELETE("/v1/fingerings/:id", GetController().DeleteFingering)

		v1.GET("/v1/first_frets", GetController().GetFirst_frets)
		v1.GET("/v1/first_frets/:id", GetController().GetFirst_fret)
		v1.POST("/v1/first_frets", GetController().PostFirst_fret)
		v1.PATCH("/v1/first_frets/:id", GetController().UpdateFirst_fret)
		v1.PUT("/v1/first_frets/:id", GetController().UpdateFirst_fret)
		v1.DELETE("/v1/first_frets/:id", GetController().DeleteFirst_fret)

		v1.GET("/v1/foos", GetController().GetFoos)
		v1.GET("/v1/foos/:id", GetController().GetFoo)
		v1.POST("/v1/foos", GetController().PostFoo)
		v1.PATCH("/v1/foos/:id", GetController().UpdateFoo)
		v1.PUT("/v1/foos/:id", GetController().UpdateFoo)
		v1.DELETE("/v1/foos/:id", GetController().DeleteFoo)

		v1.GET("/v1/for_parts", GetController().GetFor_parts)
		v1.GET("/v1/for_parts/:id", GetController().GetFor_part)
		v1.POST("/v1/for_parts", GetController().PostFor_part)
		v1.PATCH("/v1/for_parts/:id", GetController().UpdateFor_part)
		v1.PUT("/v1/for_parts/:id", GetController().UpdateFor_part)
		v1.DELETE("/v1/for_parts/:id", GetController().DeleteFor_part)

		v1.GET("/v1/formatted_symbols", GetController().GetFormatted_symbols)
		v1.GET("/v1/formatted_symbols/:id", GetController().GetFormatted_symbol)
		v1.POST("/v1/formatted_symbols", GetController().PostFormatted_symbol)
		v1.PATCH("/v1/formatted_symbols/:id", GetController().UpdateFormatted_symbol)
		v1.PUT("/v1/formatted_symbols/:id", GetController().UpdateFormatted_symbol)
		v1.DELETE("/v1/formatted_symbols/:id", GetController().DeleteFormatted_symbol)

		v1.GET("/v1/formatted_symbol_ids", GetController().GetFormatted_symbol_ids)
		v1.GET("/v1/formatted_symbol_ids/:id", GetController().GetFormatted_symbol_id)
		v1.POST("/v1/formatted_symbol_ids", GetController().PostFormatted_symbol_id)
		v1.PATCH("/v1/formatted_symbol_ids/:id", GetController().UpdateFormatted_symbol_id)
		v1.PUT("/v1/formatted_symbol_ids/:id", GetController().UpdateFormatted_symbol_id)
		v1.DELETE("/v1/formatted_symbol_ids/:id", GetController().DeleteFormatted_symbol_id)

		v1.GET("/v1/forwards", GetController().GetForwards)
		v1.GET("/v1/forwards/:id", GetController().GetForward)
		v1.POST("/v1/forwards", GetController().PostForward)
		v1.PATCH("/v1/forwards/:id", GetController().UpdateForward)
		v1.PUT("/v1/forwards/:id", GetController().UpdateForward)
		v1.DELETE("/v1/forwards/:id", GetController().DeleteForward)

		v1.GET("/v1/frames", GetController().GetFrames)
		v1.GET("/v1/frames/:id", GetController().GetFrame)
		v1.POST("/v1/frames", GetController().PostFrame)
		v1.PATCH("/v1/frames/:id", GetController().UpdateFrame)
		v1.PUT("/v1/frames/:id", GetController().UpdateFrame)
		v1.DELETE("/v1/frames/:id", GetController().DeleteFrame)

		v1.GET("/v1/frame_notes", GetController().GetFrame_notes)
		v1.GET("/v1/frame_notes/:id", GetController().GetFrame_note)
		v1.POST("/v1/frame_notes", GetController().PostFrame_note)
		v1.PATCH("/v1/frame_notes/:id", GetController().UpdateFrame_note)
		v1.PUT("/v1/frame_notes/:id", GetController().UpdateFrame_note)
		v1.DELETE("/v1/frame_notes/:id", GetController().DeleteFrame_note)

		v1.GET("/v1/frets", GetController().GetFrets)
		v1.GET("/v1/frets/:id", GetController().GetFret)
		v1.POST("/v1/frets", GetController().PostFret)
		v1.PATCH("/v1/frets/:id", GetController().UpdateFret)
		v1.PUT("/v1/frets/:id", GetController().UpdateFret)
		v1.DELETE("/v1/frets/:id", GetController().DeleteFret)

		v1.GET("/v1/glasss", GetController().GetGlasss)
		v1.GET("/v1/glasss/:id", GetController().GetGlass)
		v1.POST("/v1/glasss", GetController().PostGlass)
		v1.PATCH("/v1/glasss/:id", GetController().UpdateGlass)
		v1.PUT("/v1/glasss/:id", GetController().UpdateGlass)
		v1.DELETE("/v1/glasss/:id", GetController().DeleteGlass)

		v1.GET("/v1/glissandos", GetController().GetGlissandos)
		v1.GET("/v1/glissandos/:id", GetController().GetGlissando)
		v1.POST("/v1/glissandos", GetController().PostGlissando)
		v1.PATCH("/v1/glissandos/:id", GetController().UpdateGlissando)
		v1.PUT("/v1/glissandos/:id", GetController().UpdateGlissando)
		v1.DELETE("/v1/glissandos/:id", GetController().DeleteGlissando)

		v1.GET("/v1/glyphs", GetController().GetGlyphs)
		v1.GET("/v1/glyphs/:id", GetController().GetGlyph)
		v1.POST("/v1/glyphs", GetController().PostGlyph)
		v1.PATCH("/v1/glyphs/:id", GetController().UpdateGlyph)
		v1.PUT("/v1/glyphs/:id", GetController().UpdateGlyph)
		v1.DELETE("/v1/glyphs/:id", GetController().DeleteGlyph)

		v1.GET("/v1/graces", GetController().GetGraces)
		v1.GET("/v1/graces/:id", GetController().GetGrace)
		v1.POST("/v1/graces", GetController().PostGrace)
		v1.PATCH("/v1/graces/:id", GetController().UpdateGrace)
		v1.PUT("/v1/graces/:id", GetController().UpdateGrace)
		v1.DELETE("/v1/graces/:id", GetController().DeleteGrace)

		v1.GET("/v1/group_barlines", GetController().GetGroup_barlines)
		v1.GET("/v1/group_barlines/:id", GetController().GetGroup_barline)
		v1.POST("/v1/group_barlines", GetController().PostGroup_barline)
		v1.PATCH("/v1/group_barlines/:id", GetController().UpdateGroup_barline)
		v1.PUT("/v1/group_barlines/:id", GetController().UpdateGroup_barline)
		v1.DELETE("/v1/group_barlines/:id", GetController().DeleteGroup_barline)

		v1.GET("/v1/group_symbols", GetController().GetGroup_symbols)
		v1.GET("/v1/group_symbols/:id", GetController().GetGroup_symbol)
		v1.POST("/v1/group_symbols", GetController().PostGroup_symbol)
		v1.PATCH("/v1/group_symbols/:id", GetController().UpdateGroup_symbol)
		v1.PUT("/v1/group_symbols/:id", GetController().UpdateGroup_symbol)
		v1.DELETE("/v1/group_symbols/:id", GetController().DeleteGroup_symbol)

		v1.GET("/v1/groupings", GetController().GetGroupings)
		v1.GET("/v1/groupings/:id", GetController().GetGrouping)
		v1.POST("/v1/groupings", GetController().PostGrouping)
		v1.PATCH("/v1/groupings/:id", GetController().UpdateGrouping)
		v1.PUT("/v1/groupings/:id", GetController().UpdateGrouping)
		v1.DELETE("/v1/groupings/:id", GetController().DeleteGrouping)

		v1.GET("/v1/hammer_on_pull_offs", GetController().GetHammer_on_pull_offs)
		v1.GET("/v1/hammer_on_pull_offs/:id", GetController().GetHammer_on_pull_off)
		v1.POST("/v1/hammer_on_pull_offs", GetController().PostHammer_on_pull_off)
		v1.PATCH("/v1/hammer_on_pull_offs/:id", GetController().UpdateHammer_on_pull_off)
		v1.PUT("/v1/hammer_on_pull_offs/:id", GetController().UpdateHammer_on_pull_off)
		v1.DELETE("/v1/hammer_on_pull_offs/:id", GetController().DeleteHammer_on_pull_off)

		v1.GET("/v1/handbells", GetController().GetHandbells)
		v1.GET("/v1/handbells/:id", GetController().GetHandbell)
		v1.POST("/v1/handbells", GetController().PostHandbell)
		v1.PATCH("/v1/handbells/:id", GetController().UpdateHandbell)
		v1.PUT("/v1/handbells/:id", GetController().UpdateHandbell)
		v1.DELETE("/v1/handbells/:id", GetController().DeleteHandbell)

		v1.GET("/v1/harmon_closeds", GetController().GetHarmon_closeds)
		v1.GET("/v1/harmon_closeds/:id", GetController().GetHarmon_closed)
		v1.POST("/v1/harmon_closeds", GetController().PostHarmon_closed)
		v1.PATCH("/v1/harmon_closeds/:id", GetController().UpdateHarmon_closed)
		v1.PUT("/v1/harmon_closeds/:id", GetController().UpdateHarmon_closed)
		v1.DELETE("/v1/harmon_closeds/:id", GetController().DeleteHarmon_closed)

		v1.GET("/v1/harmon_mutes", GetController().GetHarmon_mutes)
		v1.GET("/v1/harmon_mutes/:id", GetController().GetHarmon_mute)
		v1.POST("/v1/harmon_mutes", GetController().PostHarmon_mute)
		v1.PATCH("/v1/harmon_mutes/:id", GetController().UpdateHarmon_mute)
		v1.PUT("/v1/harmon_mutes/:id", GetController().UpdateHarmon_mute)
		v1.DELETE("/v1/harmon_mutes/:id", GetController().DeleteHarmon_mute)

		v1.GET("/v1/harmonics", GetController().GetHarmonics)
		v1.GET("/v1/harmonics/:id", GetController().GetHarmonic)
		v1.POST("/v1/harmonics", GetController().PostHarmonic)
		v1.PATCH("/v1/harmonics/:id", GetController().UpdateHarmonic)
		v1.PUT("/v1/harmonics/:id", GetController().UpdateHarmonic)
		v1.DELETE("/v1/harmonics/:id", GetController().DeleteHarmonic)

		v1.GET("/v1/harmonys", GetController().GetHarmonys)
		v1.GET("/v1/harmonys/:id", GetController().GetHarmony)
		v1.POST("/v1/harmonys", GetController().PostHarmony)
		v1.PATCH("/v1/harmonys/:id", GetController().UpdateHarmony)
		v1.PUT("/v1/harmonys/:id", GetController().UpdateHarmony)
		v1.DELETE("/v1/harmonys/:id", GetController().DeleteHarmony)

		v1.GET("/v1/harmony_alters", GetController().GetHarmony_alters)
		v1.GET("/v1/harmony_alters/:id", GetController().GetHarmony_alter)
		v1.POST("/v1/harmony_alters", GetController().PostHarmony_alter)
		v1.PATCH("/v1/harmony_alters/:id", GetController().UpdateHarmony_alter)
		v1.PUT("/v1/harmony_alters/:id", GetController().UpdateHarmony_alter)
		v1.DELETE("/v1/harmony_alters/:id", GetController().DeleteHarmony_alter)

		v1.GET("/v1/harp_pedalss", GetController().GetHarp_pedalss)
		v1.GET("/v1/harp_pedalss/:id", GetController().GetHarp_pedals)
		v1.POST("/v1/harp_pedalss", GetController().PostHarp_pedals)
		v1.PATCH("/v1/harp_pedalss/:id", GetController().UpdateHarp_pedals)
		v1.PUT("/v1/harp_pedalss/:id", GetController().UpdateHarp_pedals)
		v1.DELETE("/v1/harp_pedalss/:id", GetController().DeleteHarp_pedals)

		v1.GET("/v1/heel_toes", GetController().GetHeel_toes)
		v1.GET("/v1/heel_toes/:id", GetController().GetHeel_toe)
		v1.POST("/v1/heel_toes", GetController().PostHeel_toe)
		v1.PATCH("/v1/heel_toes/:id", GetController().UpdateHeel_toe)
		v1.PUT("/v1/heel_toes/:id", GetController().UpdateHeel_toe)
		v1.DELETE("/v1/heel_toes/:id", GetController().DeleteHeel_toe)

		v1.GET("/v1/holes", GetController().GetHoles)
		v1.GET("/v1/holes/:id", GetController().GetHole)
		v1.POST("/v1/holes", GetController().PostHole)
		v1.PATCH("/v1/holes/:id", GetController().UpdateHole)
		v1.PUT("/v1/holes/:id", GetController().UpdateHole)
		v1.DELETE("/v1/holes/:id", GetController().DeleteHole)

		v1.GET("/v1/hole_closeds", GetController().GetHole_closeds)
		v1.GET("/v1/hole_closeds/:id", GetController().GetHole_closed)
		v1.POST("/v1/hole_closeds", GetController().PostHole_closed)
		v1.PATCH("/v1/hole_closeds/:id", GetController().UpdateHole_closed)
		v1.PUT("/v1/hole_closeds/:id", GetController().UpdateHole_closed)
		v1.DELETE("/v1/hole_closeds/:id", GetController().DeleteHole_closed)

		v1.GET("/v1/horizontal_turns", GetController().GetHorizontal_turns)
		v1.GET("/v1/horizontal_turns/:id", GetController().GetHorizontal_turn)
		v1.POST("/v1/horizontal_turns", GetController().PostHorizontal_turn)
		v1.PATCH("/v1/horizontal_turns/:id", GetController().UpdateHorizontal_turn)
		v1.PUT("/v1/horizontal_turns/:id", GetController().UpdateHorizontal_turn)
		v1.DELETE("/v1/horizontal_turns/:id", GetController().DeleteHorizontal_turn)

		v1.GET("/v1/identifications", GetController().GetIdentifications)
		v1.GET("/v1/identifications/:id", GetController().GetIdentification)
		v1.POST("/v1/identifications", GetController().PostIdentification)
		v1.PATCH("/v1/identifications/:id", GetController().UpdateIdentification)
		v1.PUT("/v1/identifications/:id", GetController().UpdateIdentification)
		v1.DELETE("/v1/identifications/:id", GetController().DeleteIdentification)

		v1.GET("/v1/images", GetController().GetImages)
		v1.GET("/v1/images/:id", GetController().GetImage)
		v1.POST("/v1/images", GetController().PostImage)
		v1.PATCH("/v1/images/:id", GetController().UpdateImage)
		v1.PUT("/v1/images/:id", GetController().UpdateImage)
		v1.DELETE("/v1/images/:id", GetController().DeleteImage)

		v1.GET("/v1/instruments", GetController().GetInstruments)
		v1.GET("/v1/instruments/:id", GetController().GetInstrument)
		v1.POST("/v1/instruments", GetController().PostInstrument)
		v1.PATCH("/v1/instruments/:id", GetController().UpdateInstrument)
		v1.PUT("/v1/instruments/:id", GetController().UpdateInstrument)
		v1.DELETE("/v1/instruments/:id", GetController().DeleteInstrument)

		v1.GET("/v1/instrument_changes", GetController().GetInstrument_changes)
		v1.GET("/v1/instrument_changes/:id", GetController().GetInstrument_change)
		v1.POST("/v1/instrument_changes", GetController().PostInstrument_change)
		v1.PATCH("/v1/instrument_changes/:id", GetController().UpdateInstrument_change)
		v1.PUT("/v1/instrument_changes/:id", GetController().UpdateInstrument_change)
		v1.DELETE("/v1/instrument_changes/:id", GetController().DeleteInstrument_change)

		v1.GET("/v1/instrument_links", GetController().GetInstrument_links)
		v1.GET("/v1/instrument_links/:id", GetController().GetInstrument_link)
		v1.POST("/v1/instrument_links", GetController().PostInstrument_link)
		v1.PATCH("/v1/instrument_links/:id", GetController().UpdateInstrument_link)
		v1.PUT("/v1/instrument_links/:id", GetController().UpdateInstrument_link)
		v1.DELETE("/v1/instrument_links/:id", GetController().DeleteInstrument_link)

		v1.GET("/v1/interchangeables", GetController().GetInterchangeables)
		v1.GET("/v1/interchangeables/:id", GetController().GetInterchangeable)
		v1.POST("/v1/interchangeables", GetController().PostInterchangeable)
		v1.PATCH("/v1/interchangeables/:id", GetController().UpdateInterchangeable)
		v1.PUT("/v1/interchangeables/:id", GetController().UpdateInterchangeable)
		v1.DELETE("/v1/interchangeables/:id", GetController().DeleteInterchangeable)

		v1.GET("/v1/inversions", GetController().GetInversions)
		v1.GET("/v1/inversions/:id", GetController().GetInversion)
		v1.POST("/v1/inversions", GetController().PostInversion)
		v1.PATCH("/v1/inversions/:id", GetController().UpdateInversion)
		v1.PUT("/v1/inversions/:id", GetController().UpdateInversion)
		v1.DELETE("/v1/inversions/:id", GetController().DeleteInversion)

		v1.GET("/v1/keys", GetController().GetKeys)
		v1.GET("/v1/keys/:id", GetController().GetKey)
		v1.POST("/v1/keys", GetController().PostKey)
		v1.PATCH("/v1/keys/:id", GetController().UpdateKey)
		v1.PUT("/v1/keys/:id", GetController().UpdateKey)
		v1.DELETE("/v1/keys/:id", GetController().DeleteKey)

		v1.GET("/v1/key_accidentals", GetController().GetKey_accidentals)
		v1.GET("/v1/key_accidentals/:id", GetController().GetKey_accidental)
		v1.POST("/v1/key_accidentals", GetController().PostKey_accidental)
		v1.PATCH("/v1/key_accidentals/:id", GetController().UpdateKey_accidental)
		v1.PUT("/v1/key_accidentals/:id", GetController().UpdateKey_accidental)
		v1.DELETE("/v1/key_accidentals/:id", GetController().DeleteKey_accidental)

		v1.GET("/v1/key_octaves", GetController().GetKey_octaves)
		v1.GET("/v1/key_octaves/:id", GetController().GetKey_octave)
		v1.POST("/v1/key_octaves", GetController().PostKey_octave)
		v1.PATCH("/v1/key_octaves/:id", GetController().UpdateKey_octave)
		v1.PUT("/v1/key_octaves/:id", GetController().UpdateKey_octave)
		v1.DELETE("/v1/key_octaves/:id", GetController().DeleteKey_octave)

		v1.GET("/v1/kinds", GetController().GetKinds)
		v1.GET("/v1/kinds/:id", GetController().GetKind)
		v1.POST("/v1/kinds", GetController().PostKind)
		v1.PATCH("/v1/kinds/:id", GetController().UpdateKind)
		v1.PUT("/v1/kinds/:id", GetController().UpdateKind)
		v1.DELETE("/v1/kinds/:id", GetController().DeleteKind)

		v1.GET("/v1/levels", GetController().GetLevels)
		v1.GET("/v1/levels/:id", GetController().GetLevel)
		v1.POST("/v1/levels", GetController().PostLevel)
		v1.PATCH("/v1/levels/:id", GetController().UpdateLevel)
		v1.PUT("/v1/levels/:id", GetController().UpdateLevel)
		v1.DELETE("/v1/levels/:id", GetController().DeleteLevel)

		v1.GET("/v1/line_details", GetController().GetLine_details)
		v1.GET("/v1/line_details/:id", GetController().GetLine_detail)
		v1.POST("/v1/line_details", GetController().PostLine_detail)
		v1.PATCH("/v1/line_details/:id", GetController().UpdateLine_detail)
		v1.PUT("/v1/line_details/:id", GetController().UpdateLine_detail)
		v1.DELETE("/v1/line_details/:id", GetController().DeleteLine_detail)

		v1.GET("/v1/line_widths", GetController().GetLine_widths)
		v1.GET("/v1/line_widths/:id", GetController().GetLine_width)
		v1.POST("/v1/line_widths", GetController().PostLine_width)
		v1.PATCH("/v1/line_widths/:id", GetController().UpdateLine_width)
		v1.PUT("/v1/line_widths/:id", GetController().UpdateLine_width)
		v1.DELETE("/v1/line_widths/:id", GetController().DeleteLine_width)

		v1.GET("/v1/links", GetController().GetLinks)
		v1.GET("/v1/links/:id", GetController().GetLink)
		v1.POST("/v1/links", GetController().PostLink)
		v1.PATCH("/v1/links/:id", GetController().UpdateLink)
		v1.PUT("/v1/links/:id", GetController().UpdateLink)
		v1.DELETE("/v1/links/:id", GetController().DeleteLink)

		v1.GET("/v1/listens", GetController().GetListens)
		v1.GET("/v1/listens/:id", GetController().GetListen)
		v1.POST("/v1/listens", GetController().PostListen)
		v1.PATCH("/v1/listens/:id", GetController().UpdateListen)
		v1.PUT("/v1/listens/:id", GetController().UpdateListen)
		v1.DELETE("/v1/listens/:id", GetController().DeleteListen)

		v1.GET("/v1/listenings", GetController().GetListenings)
		v1.GET("/v1/listenings/:id", GetController().GetListening)
		v1.POST("/v1/listenings", GetController().PostListening)
		v1.PATCH("/v1/listenings/:id", GetController().UpdateListening)
		v1.PUT("/v1/listenings/:id", GetController().UpdateListening)
		v1.DELETE("/v1/listenings/:id", GetController().DeleteListening)

		v1.GET("/v1/lyrics", GetController().GetLyrics)
		v1.GET("/v1/lyrics/:id", GetController().GetLyric)
		v1.POST("/v1/lyrics", GetController().PostLyric)
		v1.PATCH("/v1/lyrics/:id", GetController().UpdateLyric)
		v1.PUT("/v1/lyrics/:id", GetController().UpdateLyric)
		v1.DELETE("/v1/lyrics/:id", GetController().DeleteLyric)

		v1.GET("/v1/lyric_fonts", GetController().GetLyric_fonts)
		v1.GET("/v1/lyric_fonts/:id", GetController().GetLyric_font)
		v1.POST("/v1/lyric_fonts", GetController().PostLyric_font)
		v1.PATCH("/v1/lyric_fonts/:id", GetController().UpdateLyric_font)
		v1.PUT("/v1/lyric_fonts/:id", GetController().UpdateLyric_font)
		v1.DELETE("/v1/lyric_fonts/:id", GetController().DeleteLyric_font)

		v1.GET("/v1/lyric_languages", GetController().GetLyric_languages)
		v1.GET("/v1/lyric_languages/:id", GetController().GetLyric_language)
		v1.POST("/v1/lyric_languages", GetController().PostLyric_language)
		v1.PATCH("/v1/lyric_languages/:id", GetController().UpdateLyric_language)
		v1.PUT("/v1/lyric_languages/:id", GetController().UpdateLyric_language)
		v1.DELETE("/v1/lyric_languages/:id", GetController().DeleteLyric_language)

		v1.GET("/v1/measure_layouts", GetController().GetMeasure_layouts)
		v1.GET("/v1/measure_layouts/:id", GetController().GetMeasure_layout)
		v1.POST("/v1/measure_layouts", GetController().PostMeasure_layout)
		v1.PATCH("/v1/measure_layouts/:id", GetController().UpdateMeasure_layout)
		v1.PUT("/v1/measure_layouts/:id", GetController().UpdateMeasure_layout)
		v1.DELETE("/v1/measure_layouts/:id", GetController().DeleteMeasure_layout)

		v1.GET("/v1/measure_numberings", GetController().GetMeasure_numberings)
		v1.GET("/v1/measure_numberings/:id", GetController().GetMeasure_numbering)
		v1.POST("/v1/measure_numberings", GetController().PostMeasure_numbering)
		v1.PATCH("/v1/measure_numberings/:id", GetController().UpdateMeasure_numbering)
		v1.PUT("/v1/measure_numberings/:id", GetController().UpdateMeasure_numbering)
		v1.DELETE("/v1/measure_numberings/:id", GetController().DeleteMeasure_numbering)

		v1.GET("/v1/measure_repeats", GetController().GetMeasure_repeats)
		v1.GET("/v1/measure_repeats/:id", GetController().GetMeasure_repeat)
		v1.POST("/v1/measure_repeats", GetController().PostMeasure_repeat)
		v1.PATCH("/v1/measure_repeats/:id", GetController().UpdateMeasure_repeat)
		v1.PUT("/v1/measure_repeats/:id", GetController().UpdateMeasure_repeat)
		v1.DELETE("/v1/measure_repeats/:id", GetController().DeleteMeasure_repeat)

		v1.GET("/v1/measure_styles", GetController().GetMeasure_styles)
		v1.GET("/v1/measure_styles/:id", GetController().GetMeasure_style)
		v1.POST("/v1/measure_styles", GetController().PostMeasure_style)
		v1.PATCH("/v1/measure_styles/:id", GetController().UpdateMeasure_style)
		v1.PUT("/v1/measure_styles/:id", GetController().UpdateMeasure_style)
		v1.DELETE("/v1/measure_styles/:id", GetController().DeleteMeasure_style)

		v1.GET("/v1/membranes", GetController().GetMembranes)
		v1.GET("/v1/membranes/:id", GetController().GetMembrane)
		v1.POST("/v1/membranes", GetController().PostMembrane)
		v1.PATCH("/v1/membranes/:id", GetController().UpdateMembrane)
		v1.PUT("/v1/membranes/:id", GetController().UpdateMembrane)
		v1.DELETE("/v1/membranes/:id", GetController().DeleteMembrane)

		v1.GET("/v1/metals", GetController().GetMetals)
		v1.GET("/v1/metals/:id", GetController().GetMetal)
		v1.POST("/v1/metals", GetController().PostMetal)
		v1.PATCH("/v1/metals/:id", GetController().UpdateMetal)
		v1.PUT("/v1/metals/:id", GetController().UpdateMetal)
		v1.DELETE("/v1/metals/:id", GetController().DeleteMetal)

		v1.GET("/v1/metronomes", GetController().GetMetronomes)
		v1.GET("/v1/metronomes/:id", GetController().GetMetronome)
		v1.POST("/v1/metronomes", GetController().PostMetronome)
		v1.PATCH("/v1/metronomes/:id", GetController().UpdateMetronome)
		v1.PUT("/v1/metronomes/:id", GetController().UpdateMetronome)
		v1.DELETE("/v1/metronomes/:id", GetController().DeleteMetronome)

		v1.GET("/v1/metronome_beams", GetController().GetMetronome_beams)
		v1.GET("/v1/metronome_beams/:id", GetController().GetMetronome_beam)
		v1.POST("/v1/metronome_beams", GetController().PostMetronome_beam)
		v1.PATCH("/v1/metronome_beams/:id", GetController().UpdateMetronome_beam)
		v1.PUT("/v1/metronome_beams/:id", GetController().UpdateMetronome_beam)
		v1.DELETE("/v1/metronome_beams/:id", GetController().DeleteMetronome_beam)

		v1.GET("/v1/metronome_notes", GetController().GetMetronome_notes)
		v1.GET("/v1/metronome_notes/:id", GetController().GetMetronome_note)
		v1.POST("/v1/metronome_notes", GetController().PostMetronome_note)
		v1.PATCH("/v1/metronome_notes/:id", GetController().UpdateMetronome_note)
		v1.PUT("/v1/metronome_notes/:id", GetController().UpdateMetronome_note)
		v1.DELETE("/v1/metronome_notes/:id", GetController().DeleteMetronome_note)

		v1.GET("/v1/metronome_tieds", GetController().GetMetronome_tieds)
		v1.GET("/v1/metronome_tieds/:id", GetController().GetMetronome_tied)
		v1.POST("/v1/metronome_tieds", GetController().PostMetronome_tied)
		v1.PATCH("/v1/metronome_tieds/:id", GetController().UpdateMetronome_tied)
		v1.PUT("/v1/metronome_tieds/:id", GetController().UpdateMetronome_tied)
		v1.DELETE("/v1/metronome_tieds/:id", GetController().DeleteMetronome_tied)

		v1.GET("/v1/metronome_tuplets", GetController().GetMetronome_tuplets)
		v1.GET("/v1/metronome_tuplets/:id", GetController().GetMetronome_tuplet)
		v1.POST("/v1/metronome_tuplets", GetController().PostMetronome_tuplet)
		v1.PATCH("/v1/metronome_tuplets/:id", GetController().UpdateMetronome_tuplet)
		v1.PUT("/v1/metronome_tuplets/:id", GetController().UpdateMetronome_tuplet)
		v1.DELETE("/v1/metronome_tuplets/:id", GetController().DeleteMetronome_tuplet)

		v1.GET("/v1/midi_devices", GetController().GetMidi_devices)
		v1.GET("/v1/midi_devices/:id", GetController().GetMidi_device)
		v1.POST("/v1/midi_devices", GetController().PostMidi_device)
		v1.PATCH("/v1/midi_devices/:id", GetController().UpdateMidi_device)
		v1.PUT("/v1/midi_devices/:id", GetController().UpdateMidi_device)
		v1.DELETE("/v1/midi_devices/:id", GetController().DeleteMidi_device)

		v1.GET("/v1/midi_instruments", GetController().GetMidi_instruments)
		v1.GET("/v1/midi_instruments/:id", GetController().GetMidi_instrument)
		v1.POST("/v1/midi_instruments", GetController().PostMidi_instrument)
		v1.PATCH("/v1/midi_instruments/:id", GetController().UpdateMidi_instrument)
		v1.PUT("/v1/midi_instruments/:id", GetController().UpdateMidi_instrument)
		v1.DELETE("/v1/midi_instruments/:id", GetController().DeleteMidi_instrument)

		v1.GET("/v1/miscellaneouss", GetController().GetMiscellaneouss)
		v1.GET("/v1/miscellaneouss/:id", GetController().GetMiscellaneous)
		v1.POST("/v1/miscellaneouss", GetController().PostMiscellaneous)
		v1.PATCH("/v1/miscellaneouss/:id", GetController().UpdateMiscellaneous)
		v1.PUT("/v1/miscellaneouss/:id", GetController().UpdateMiscellaneous)
		v1.DELETE("/v1/miscellaneouss/:id", GetController().DeleteMiscellaneous)

		v1.GET("/v1/miscellaneous_fields", GetController().GetMiscellaneous_fields)
		v1.GET("/v1/miscellaneous_fields/:id", GetController().GetMiscellaneous_field)
		v1.POST("/v1/miscellaneous_fields", GetController().PostMiscellaneous_field)
		v1.PATCH("/v1/miscellaneous_fields/:id", GetController().UpdateMiscellaneous_field)
		v1.PUT("/v1/miscellaneous_fields/:id", GetController().UpdateMiscellaneous_field)
		v1.DELETE("/v1/miscellaneous_fields/:id", GetController().DeleteMiscellaneous_field)

		v1.GET("/v1/mordents", GetController().GetMordents)
		v1.GET("/v1/mordents/:id", GetController().GetMordent)
		v1.POST("/v1/mordents", GetController().PostMordent)
		v1.PATCH("/v1/mordents/:id", GetController().UpdateMordent)
		v1.PUT("/v1/mordents/:id", GetController().UpdateMordent)
		v1.DELETE("/v1/mordents/:id", GetController().DeleteMordent)

		v1.GET("/v1/multiple_rests", GetController().GetMultiple_rests)
		v1.GET("/v1/multiple_rests/:id", GetController().GetMultiple_rest)
		v1.POST("/v1/multiple_rests", GetController().PostMultiple_rest)
		v1.PATCH("/v1/multiple_rests/:id", GetController().UpdateMultiple_rest)
		v1.PUT("/v1/multiple_rests/:id", GetController().UpdateMultiple_rest)
		v1.DELETE("/v1/multiple_rests/:id", GetController().DeleteMultiple_rest)

		v1.GET("/v1/name_displays", GetController().GetName_displays)
		v1.GET("/v1/name_displays/:id", GetController().GetName_display)
		v1.POST("/v1/name_displays", GetController().PostName_display)
		v1.PATCH("/v1/name_displays/:id", GetController().UpdateName_display)
		v1.PUT("/v1/name_displays/:id", GetController().UpdateName_display)
		v1.DELETE("/v1/name_displays/:id", GetController().DeleteName_display)

		v1.GET("/v1/non_arpeggiates", GetController().GetNon_arpeggiates)
		v1.GET("/v1/non_arpeggiates/:id", GetController().GetNon_arpeggiate)
		v1.POST("/v1/non_arpeggiates", GetController().PostNon_arpeggiate)
		v1.PATCH("/v1/non_arpeggiates/:id", GetController().UpdateNon_arpeggiate)
		v1.PUT("/v1/non_arpeggiates/:id", GetController().UpdateNon_arpeggiate)
		v1.DELETE("/v1/non_arpeggiates/:id", GetController().DeleteNon_arpeggiate)

		v1.GET("/v1/notationss", GetController().GetNotationss)
		v1.GET("/v1/notationss/:id", GetController().GetNotations)
		v1.POST("/v1/notationss", GetController().PostNotations)
		v1.PATCH("/v1/notationss/:id", GetController().UpdateNotations)
		v1.PUT("/v1/notationss/:id", GetController().UpdateNotations)
		v1.DELETE("/v1/notationss/:id", GetController().DeleteNotations)

		v1.GET("/v1/notes", GetController().GetNotes)
		v1.GET("/v1/notes/:id", GetController().GetNote)
		v1.POST("/v1/notes", GetController().PostNote)
		v1.PATCH("/v1/notes/:id", GetController().UpdateNote)
		v1.PUT("/v1/notes/:id", GetController().UpdateNote)
		v1.DELETE("/v1/notes/:id", GetController().DeleteNote)

		v1.GET("/v1/note_sizes", GetController().GetNote_sizes)
		v1.GET("/v1/note_sizes/:id", GetController().GetNote_size)
		v1.POST("/v1/note_sizes", GetController().PostNote_size)
		v1.PATCH("/v1/note_sizes/:id", GetController().UpdateNote_size)
		v1.PUT("/v1/note_sizes/:id", GetController().UpdateNote_size)
		v1.DELETE("/v1/note_sizes/:id", GetController().DeleteNote_size)

		v1.GET("/v1/note_types", GetController().GetNote_types)
		v1.GET("/v1/note_types/:id", GetController().GetNote_type)
		v1.POST("/v1/note_types", GetController().PostNote_type)
		v1.PATCH("/v1/note_types/:id", GetController().UpdateNote_type)
		v1.PUT("/v1/note_types/:id", GetController().UpdateNote_type)
		v1.DELETE("/v1/note_types/:id", GetController().DeleteNote_type)

		v1.GET("/v1/noteheads", GetController().GetNoteheads)
		v1.GET("/v1/noteheads/:id", GetController().GetNotehead)
		v1.POST("/v1/noteheads", GetController().PostNotehead)
		v1.PATCH("/v1/noteheads/:id", GetController().UpdateNotehead)
		v1.PUT("/v1/noteheads/:id", GetController().UpdateNotehead)
		v1.DELETE("/v1/noteheads/:id", GetController().DeleteNotehead)

		v1.GET("/v1/notehead_texts", GetController().GetNotehead_texts)
		v1.GET("/v1/notehead_texts/:id", GetController().GetNotehead_text)
		v1.POST("/v1/notehead_texts", GetController().PostNotehead_text)
		v1.PATCH("/v1/notehead_texts/:id", GetController().UpdateNotehead_text)
		v1.PUT("/v1/notehead_texts/:id", GetController().UpdateNotehead_text)
		v1.DELETE("/v1/notehead_texts/:id", GetController().DeleteNotehead_text)

		v1.GET("/v1/numerals", GetController().GetNumerals)
		v1.GET("/v1/numerals/:id", GetController().GetNumeral)
		v1.POST("/v1/numerals", GetController().PostNumeral)
		v1.PATCH("/v1/numerals/:id", GetController().UpdateNumeral)
		v1.PUT("/v1/numerals/:id", GetController().UpdateNumeral)
		v1.DELETE("/v1/numerals/:id", GetController().DeleteNumeral)

		v1.GET("/v1/numeral_keys", GetController().GetNumeral_keys)
		v1.GET("/v1/numeral_keys/:id", GetController().GetNumeral_key)
		v1.POST("/v1/numeral_keys", GetController().PostNumeral_key)
		v1.PATCH("/v1/numeral_keys/:id", GetController().UpdateNumeral_key)
		v1.PUT("/v1/numeral_keys/:id", GetController().UpdateNumeral_key)
		v1.DELETE("/v1/numeral_keys/:id", GetController().DeleteNumeral_key)

		v1.GET("/v1/numeral_roots", GetController().GetNumeral_roots)
		v1.GET("/v1/numeral_roots/:id", GetController().GetNumeral_root)
		v1.POST("/v1/numeral_roots", GetController().PostNumeral_root)
		v1.PATCH("/v1/numeral_roots/:id", GetController().UpdateNumeral_root)
		v1.PUT("/v1/numeral_roots/:id", GetController().UpdateNumeral_root)
		v1.DELETE("/v1/numeral_roots/:id", GetController().DeleteNumeral_root)

		v1.GET("/v1/octave_shifts", GetController().GetOctave_shifts)
		v1.GET("/v1/octave_shifts/:id", GetController().GetOctave_shift)
		v1.POST("/v1/octave_shifts", GetController().PostOctave_shift)
		v1.PATCH("/v1/octave_shifts/:id", GetController().UpdateOctave_shift)
		v1.PUT("/v1/octave_shifts/:id", GetController().UpdateOctave_shift)
		v1.DELETE("/v1/octave_shifts/:id", GetController().DeleteOctave_shift)

		v1.GET("/v1/offsets", GetController().GetOffsets)
		v1.GET("/v1/offsets/:id", GetController().GetOffset)
		v1.POST("/v1/offsets", GetController().PostOffset)
		v1.PATCH("/v1/offsets/:id", GetController().UpdateOffset)
		v1.PUT("/v1/offsets/:id", GetController().UpdateOffset)
		v1.DELETE("/v1/offsets/:id", GetController().DeleteOffset)

		v1.GET("/v1/opuss", GetController().GetOpuss)
		v1.GET("/v1/opuss/:id", GetController().GetOpus)
		v1.POST("/v1/opuss", GetController().PostOpus)
		v1.PATCH("/v1/opuss/:id", GetController().UpdateOpus)
		v1.PUT("/v1/opuss/:id", GetController().UpdateOpus)
		v1.DELETE("/v1/opuss/:id", GetController().DeleteOpus)

		v1.GET("/v1/ornamentss", GetController().GetOrnamentss)
		v1.GET("/v1/ornamentss/:id", GetController().GetOrnaments)
		v1.POST("/v1/ornamentss", GetController().PostOrnaments)
		v1.PATCH("/v1/ornamentss/:id", GetController().UpdateOrnaments)
		v1.PUT("/v1/ornamentss/:id", GetController().UpdateOrnaments)
		v1.DELETE("/v1/ornamentss/:id", GetController().DeleteOrnaments)

		v1.GET("/v1/other_appearances", GetController().GetOther_appearances)
		v1.GET("/v1/other_appearances/:id", GetController().GetOther_appearance)
		v1.POST("/v1/other_appearances", GetController().PostOther_appearance)
		v1.PATCH("/v1/other_appearances/:id", GetController().UpdateOther_appearance)
		v1.PUT("/v1/other_appearances/:id", GetController().UpdateOther_appearance)
		v1.DELETE("/v1/other_appearances/:id", GetController().DeleteOther_appearance)

		v1.GET("/v1/other_listenings", GetController().GetOther_listenings)
		v1.GET("/v1/other_listenings/:id", GetController().GetOther_listening)
		v1.POST("/v1/other_listenings", GetController().PostOther_listening)
		v1.PATCH("/v1/other_listenings/:id", GetController().UpdateOther_listening)
		v1.PUT("/v1/other_listenings/:id", GetController().UpdateOther_listening)
		v1.DELETE("/v1/other_listenings/:id", GetController().DeleteOther_listening)

		v1.GET("/v1/other_notations", GetController().GetOther_notations)
		v1.GET("/v1/other_notations/:id", GetController().GetOther_notation)
		v1.POST("/v1/other_notations", GetController().PostOther_notation)
		v1.PATCH("/v1/other_notations/:id", GetController().UpdateOther_notation)
		v1.PUT("/v1/other_notations/:id", GetController().UpdateOther_notation)
		v1.DELETE("/v1/other_notations/:id", GetController().DeleteOther_notation)

		v1.GET("/v1/other_plays", GetController().GetOther_plays)
		v1.GET("/v1/other_plays/:id", GetController().GetOther_play)
		v1.POST("/v1/other_plays", GetController().PostOther_play)
		v1.PATCH("/v1/other_plays/:id", GetController().UpdateOther_play)
		v1.PUT("/v1/other_plays/:id", GetController().UpdateOther_play)
		v1.DELETE("/v1/other_plays/:id", GetController().DeleteOther_play)

		v1.GET("/v1/page_layouts", GetController().GetPage_layouts)
		v1.GET("/v1/page_layouts/:id", GetController().GetPage_layout)
		v1.POST("/v1/page_layouts", GetController().PostPage_layout)
		v1.PATCH("/v1/page_layouts/:id", GetController().UpdatePage_layout)
		v1.PUT("/v1/page_layouts/:id", GetController().UpdatePage_layout)
		v1.DELETE("/v1/page_layouts/:id", GetController().DeletePage_layout)

		v1.GET("/v1/page_marginss", GetController().GetPage_marginss)
		v1.GET("/v1/page_marginss/:id", GetController().GetPage_margins)
		v1.POST("/v1/page_marginss", GetController().PostPage_margins)
		v1.PATCH("/v1/page_marginss/:id", GetController().UpdatePage_margins)
		v1.PUT("/v1/page_marginss/:id", GetController().UpdatePage_margins)
		v1.DELETE("/v1/page_marginss/:id", GetController().DeletePage_margins)

		v1.GET("/v1/part_clefs", GetController().GetPart_clefs)
		v1.GET("/v1/part_clefs/:id", GetController().GetPart_clef)
		v1.POST("/v1/part_clefs", GetController().PostPart_clef)
		v1.PATCH("/v1/part_clefs/:id", GetController().UpdatePart_clef)
		v1.PUT("/v1/part_clefs/:id", GetController().UpdatePart_clef)
		v1.DELETE("/v1/part_clefs/:id", GetController().DeletePart_clef)

		v1.GET("/v1/part_groups", GetController().GetPart_groups)
		v1.GET("/v1/part_groups/:id", GetController().GetPart_group)
		v1.POST("/v1/part_groups", GetController().PostPart_group)
		v1.PATCH("/v1/part_groups/:id", GetController().UpdatePart_group)
		v1.PUT("/v1/part_groups/:id", GetController().UpdatePart_group)
		v1.DELETE("/v1/part_groups/:id", GetController().DeletePart_group)

		v1.GET("/v1/part_links", GetController().GetPart_links)
		v1.GET("/v1/part_links/:id", GetController().GetPart_link)
		v1.POST("/v1/part_links", GetController().PostPart_link)
		v1.PATCH("/v1/part_links/:id", GetController().UpdatePart_link)
		v1.PUT("/v1/part_links/:id", GetController().UpdatePart_link)
		v1.DELETE("/v1/part_links/:id", GetController().DeletePart_link)

		v1.GET("/v1/part_lists", GetController().GetPart_lists)
		v1.GET("/v1/part_lists/:id", GetController().GetPart_list)
		v1.POST("/v1/part_lists", GetController().PostPart_list)
		v1.PATCH("/v1/part_lists/:id", GetController().UpdatePart_list)
		v1.PUT("/v1/part_lists/:id", GetController().UpdatePart_list)
		v1.DELETE("/v1/part_lists/:id", GetController().DeletePart_list)

		v1.GET("/v1/part_symbols", GetController().GetPart_symbols)
		v1.GET("/v1/part_symbols/:id", GetController().GetPart_symbol)
		v1.POST("/v1/part_symbols", GetController().PostPart_symbol)
		v1.PATCH("/v1/part_symbols/:id", GetController().UpdatePart_symbol)
		v1.PUT("/v1/part_symbols/:id", GetController().UpdatePart_symbol)
		v1.DELETE("/v1/part_symbols/:id", GetController().DeletePart_symbol)

		v1.GET("/v1/part_transposes", GetController().GetPart_transposes)
		v1.GET("/v1/part_transposes/:id", GetController().GetPart_transpose)
		v1.POST("/v1/part_transposes", GetController().PostPart_transpose)
		v1.PATCH("/v1/part_transposes/:id", GetController().UpdatePart_transpose)
		v1.PUT("/v1/part_transposes/:id", GetController().UpdatePart_transpose)
		v1.DELETE("/v1/part_transposes/:id", GetController().DeletePart_transpose)

		v1.GET("/v1/pedals", GetController().GetPedals)
		v1.GET("/v1/pedals/:id", GetController().GetPedal)
		v1.POST("/v1/pedals", GetController().PostPedal)
		v1.PATCH("/v1/pedals/:id", GetController().UpdatePedal)
		v1.PUT("/v1/pedals/:id", GetController().UpdatePedal)
		v1.DELETE("/v1/pedals/:id", GetController().DeletePedal)

		v1.GET("/v1/pedal_tunings", GetController().GetPedal_tunings)
		v1.GET("/v1/pedal_tunings/:id", GetController().GetPedal_tuning)
		v1.POST("/v1/pedal_tunings", GetController().PostPedal_tuning)
		v1.PATCH("/v1/pedal_tunings/:id", GetController().UpdatePedal_tuning)
		v1.PUT("/v1/pedal_tunings/:id", GetController().UpdatePedal_tuning)
		v1.DELETE("/v1/pedal_tunings/:id", GetController().DeletePedal_tuning)

		v1.GET("/v1/percussions", GetController().GetPercussions)
		v1.GET("/v1/percussions/:id", GetController().GetPercussion)
		v1.POST("/v1/percussions", GetController().PostPercussion)
		v1.PATCH("/v1/percussions/:id", GetController().UpdatePercussion)
		v1.PUT("/v1/percussions/:id", GetController().UpdatePercussion)
		v1.DELETE("/v1/percussions/:id", GetController().DeletePercussion)

		v1.GET("/v1/pitchs", GetController().GetPitchs)
		v1.GET("/v1/pitchs/:id", GetController().GetPitch)
		v1.POST("/v1/pitchs", GetController().PostPitch)
		v1.PATCH("/v1/pitchs/:id", GetController().UpdatePitch)
		v1.PUT("/v1/pitchs/:id", GetController().UpdatePitch)
		v1.DELETE("/v1/pitchs/:id", GetController().DeletePitch)

		v1.GET("/v1/pitcheds", GetController().GetPitcheds)
		v1.GET("/v1/pitcheds/:id", GetController().GetPitched)
		v1.POST("/v1/pitcheds", GetController().PostPitched)
		v1.PATCH("/v1/pitcheds/:id", GetController().UpdatePitched)
		v1.PUT("/v1/pitcheds/:id", GetController().UpdatePitched)
		v1.DELETE("/v1/pitcheds/:id", GetController().DeletePitched)

		v1.GET("/v1/plays", GetController().GetPlays)
		v1.GET("/v1/plays/:id", GetController().GetPlay)
		v1.POST("/v1/plays", GetController().PostPlay)
		v1.PATCH("/v1/plays/:id", GetController().UpdatePlay)
		v1.PUT("/v1/plays/:id", GetController().UpdatePlay)
		v1.DELETE("/v1/plays/:id", GetController().DeletePlay)

		v1.GET("/v1/players", GetController().GetPlayers)
		v1.GET("/v1/players/:id", GetController().GetPlayer)
		v1.POST("/v1/players", GetController().PostPlayer)
		v1.PATCH("/v1/players/:id", GetController().UpdatePlayer)
		v1.PUT("/v1/players/:id", GetController().UpdatePlayer)
		v1.DELETE("/v1/players/:id", GetController().DeletePlayer)

		v1.GET("/v1/principal_voices", GetController().GetPrincipal_voices)
		v1.GET("/v1/principal_voices/:id", GetController().GetPrincipal_voice)
		v1.POST("/v1/principal_voices", GetController().PostPrincipal_voice)
		v1.PATCH("/v1/principal_voices/:id", GetController().UpdatePrincipal_voice)
		v1.PUT("/v1/principal_voices/:id", GetController().UpdatePrincipal_voice)
		v1.DELETE("/v1/principal_voices/:id", GetController().DeletePrincipal_voice)

		v1.GET("/v1/prints", GetController().GetPrints)
		v1.GET("/v1/prints/:id", GetController().GetPrint)
		v1.POST("/v1/prints", GetController().PostPrint)
		v1.PATCH("/v1/prints/:id", GetController().UpdatePrint)
		v1.PUT("/v1/prints/:id", GetController().UpdatePrint)
		v1.DELETE("/v1/prints/:id", GetController().DeletePrint)

		v1.GET("/v1/releases", GetController().GetReleases)
		v1.GET("/v1/releases/:id", GetController().GetRelease)
		v1.POST("/v1/releases", GetController().PostRelease)
		v1.PATCH("/v1/releases/:id", GetController().UpdateRelease)
		v1.PUT("/v1/releases/:id", GetController().UpdateRelease)
		v1.DELETE("/v1/releases/:id", GetController().DeleteRelease)

		v1.GET("/v1/repeats", GetController().GetRepeats)
		v1.GET("/v1/repeats/:id", GetController().GetRepeat)
		v1.POST("/v1/repeats", GetController().PostRepeat)
		v1.PATCH("/v1/repeats/:id", GetController().UpdateRepeat)
		v1.PUT("/v1/repeats/:id", GetController().UpdateRepeat)
		v1.DELETE("/v1/repeats/:id", GetController().DeleteRepeat)

		v1.GET("/v1/rests", GetController().GetRests)
		v1.GET("/v1/rests/:id", GetController().GetRest)
		v1.POST("/v1/rests", GetController().PostRest)
		v1.PATCH("/v1/rests/:id", GetController().UpdateRest)
		v1.PUT("/v1/rests/:id", GetController().UpdateRest)
		v1.DELETE("/v1/rests/:id", GetController().DeleteRest)

		v1.GET("/v1/roots", GetController().GetRoots)
		v1.GET("/v1/roots/:id", GetController().GetRoot)
		v1.POST("/v1/roots", GetController().PostRoot)
		v1.PATCH("/v1/roots/:id", GetController().UpdateRoot)
		v1.PUT("/v1/roots/:id", GetController().UpdateRoot)
		v1.DELETE("/v1/roots/:id", GetController().DeleteRoot)

		v1.GET("/v1/root_steps", GetController().GetRoot_steps)
		v1.GET("/v1/root_steps/:id", GetController().GetRoot_step)
		v1.POST("/v1/root_steps", GetController().PostRoot_step)
		v1.PATCH("/v1/root_steps/:id", GetController().UpdateRoot_step)
		v1.PUT("/v1/root_steps/:id", GetController().UpdateRoot_step)
		v1.DELETE("/v1/root_steps/:id", GetController().DeleteRoot_step)

		v1.GET("/v1/scalings", GetController().GetScalings)
		v1.GET("/v1/scalings/:id", GetController().GetScaling)
		v1.POST("/v1/scalings", GetController().PostScaling)
		v1.PATCH("/v1/scalings/:id", GetController().UpdateScaling)
		v1.PUT("/v1/scalings/:id", GetController().UpdateScaling)
		v1.DELETE("/v1/scalings/:id", GetController().DeleteScaling)

		v1.GET("/v1/scordaturas", GetController().GetScordaturas)
		v1.GET("/v1/scordaturas/:id", GetController().GetScordatura)
		v1.POST("/v1/scordaturas", GetController().PostScordatura)
		v1.PATCH("/v1/scordaturas/:id", GetController().UpdateScordatura)
		v1.PUT("/v1/scordaturas/:id", GetController().UpdateScordatura)
		v1.DELETE("/v1/scordaturas/:id", GetController().DeleteScordatura)

		v1.GET("/v1/score_instruments", GetController().GetScore_instruments)
		v1.GET("/v1/score_instruments/:id", GetController().GetScore_instrument)
		v1.POST("/v1/score_instruments", GetController().PostScore_instrument)
		v1.PATCH("/v1/score_instruments/:id", GetController().UpdateScore_instrument)
		v1.PUT("/v1/score_instruments/:id", GetController().UpdateScore_instrument)
		v1.DELETE("/v1/score_instruments/:id", GetController().DeleteScore_instrument)

		v1.GET("/v1/score_parts", GetController().GetScore_parts)
		v1.GET("/v1/score_parts/:id", GetController().GetScore_part)
		v1.POST("/v1/score_parts", GetController().PostScore_part)
		v1.PATCH("/v1/score_parts/:id", GetController().UpdateScore_part)
		v1.PUT("/v1/score_parts/:id", GetController().UpdateScore_part)
		v1.DELETE("/v1/score_parts/:id", GetController().DeleteScore_part)

		v1.GET("/v1/score_partwises", GetController().GetScore_partwises)
		v1.GET("/v1/score_partwises/:id", GetController().GetScore_partwise)
		v1.POST("/v1/score_partwises", GetController().PostScore_partwise)
		v1.PATCH("/v1/score_partwises/:id", GetController().UpdateScore_partwise)
		v1.PUT("/v1/score_partwises/:id", GetController().UpdateScore_partwise)
		v1.DELETE("/v1/score_partwises/:id", GetController().DeleteScore_partwise)

		v1.GET("/v1/score_timewises", GetController().GetScore_timewises)
		v1.GET("/v1/score_timewises/:id", GetController().GetScore_timewise)
		v1.POST("/v1/score_timewises", GetController().PostScore_timewise)
		v1.PATCH("/v1/score_timewises/:id", GetController().UpdateScore_timewise)
		v1.PUT("/v1/score_timewises/:id", GetController().UpdateScore_timewise)
		v1.DELETE("/v1/score_timewises/:id", GetController().DeleteScore_timewise)

		v1.GET("/v1/segnos", GetController().GetSegnos)
		v1.GET("/v1/segnos/:id", GetController().GetSegno)
		v1.POST("/v1/segnos", GetController().PostSegno)
		v1.PATCH("/v1/segnos/:id", GetController().UpdateSegno)
		v1.PUT("/v1/segnos/:id", GetController().UpdateSegno)
		v1.DELETE("/v1/segnos/:id", GetController().DeleteSegno)

		v1.GET("/v1/slashs", GetController().GetSlashs)
		v1.GET("/v1/slashs/:id", GetController().GetSlash)
		v1.POST("/v1/slashs", GetController().PostSlash)
		v1.PATCH("/v1/slashs/:id", GetController().UpdateSlash)
		v1.PUT("/v1/slashs/:id", GetController().UpdateSlash)
		v1.DELETE("/v1/slashs/:id", GetController().DeleteSlash)

		v1.GET("/v1/slides", GetController().GetSlides)
		v1.GET("/v1/slides/:id", GetController().GetSlide)
		v1.POST("/v1/slides", GetController().PostSlide)
		v1.PATCH("/v1/slides/:id", GetController().UpdateSlide)
		v1.PUT("/v1/slides/:id", GetController().UpdateSlide)
		v1.DELETE("/v1/slides/:id", GetController().DeleteSlide)

		v1.GET("/v1/slurs", GetController().GetSlurs)
		v1.GET("/v1/slurs/:id", GetController().GetSlur)
		v1.POST("/v1/slurs", GetController().PostSlur)
		v1.PATCH("/v1/slurs/:id", GetController().UpdateSlur)
		v1.PUT("/v1/slurs/:id", GetController().UpdateSlur)
		v1.DELETE("/v1/slurs/:id", GetController().DeleteSlur)

		v1.GET("/v1/sounds", GetController().GetSounds)
		v1.GET("/v1/sounds/:id", GetController().GetSound)
		v1.POST("/v1/sounds", GetController().PostSound)
		v1.PATCH("/v1/sounds/:id", GetController().UpdateSound)
		v1.PUT("/v1/sounds/:id", GetController().UpdateSound)
		v1.DELETE("/v1/sounds/:id", GetController().DeleteSound)

		v1.GET("/v1/staff_detailss", GetController().GetStaff_detailss)
		v1.GET("/v1/staff_detailss/:id", GetController().GetStaff_details)
		v1.POST("/v1/staff_detailss", GetController().PostStaff_details)
		v1.PATCH("/v1/staff_detailss/:id", GetController().UpdateStaff_details)
		v1.PUT("/v1/staff_detailss/:id", GetController().UpdateStaff_details)
		v1.DELETE("/v1/staff_detailss/:id", GetController().DeleteStaff_details)

		v1.GET("/v1/staff_divides", GetController().GetStaff_divides)
		v1.GET("/v1/staff_divides/:id", GetController().GetStaff_divide)
		v1.POST("/v1/staff_divides", GetController().PostStaff_divide)
		v1.PATCH("/v1/staff_divides/:id", GetController().UpdateStaff_divide)
		v1.PUT("/v1/staff_divides/:id", GetController().UpdateStaff_divide)
		v1.DELETE("/v1/staff_divides/:id", GetController().DeleteStaff_divide)

		v1.GET("/v1/staff_layouts", GetController().GetStaff_layouts)
		v1.GET("/v1/staff_layouts/:id", GetController().GetStaff_layout)
		v1.POST("/v1/staff_layouts", GetController().PostStaff_layout)
		v1.PATCH("/v1/staff_layouts/:id", GetController().UpdateStaff_layout)
		v1.PUT("/v1/staff_layouts/:id", GetController().UpdateStaff_layout)
		v1.DELETE("/v1/staff_layouts/:id", GetController().DeleteStaff_layout)

		v1.GET("/v1/staff_sizes", GetController().GetStaff_sizes)
		v1.GET("/v1/staff_sizes/:id", GetController().GetStaff_size)
		v1.POST("/v1/staff_sizes", GetController().PostStaff_size)
		v1.PATCH("/v1/staff_sizes/:id", GetController().UpdateStaff_size)
		v1.PUT("/v1/staff_sizes/:id", GetController().UpdateStaff_size)
		v1.DELETE("/v1/staff_sizes/:id", GetController().DeleteStaff_size)

		v1.GET("/v1/staff_tunings", GetController().GetStaff_tunings)
		v1.GET("/v1/staff_tunings/:id", GetController().GetStaff_tuning)
		v1.POST("/v1/staff_tunings", GetController().PostStaff_tuning)
		v1.PATCH("/v1/staff_tunings/:id", GetController().UpdateStaff_tuning)
		v1.PUT("/v1/staff_tunings/:id", GetController().UpdateStaff_tuning)
		v1.DELETE("/v1/staff_tunings/:id", GetController().DeleteStaff_tuning)

		v1.GET("/v1/stems", GetController().GetStems)
		v1.GET("/v1/stems/:id", GetController().GetStem)
		v1.POST("/v1/stems", GetController().PostStem)
		v1.PATCH("/v1/stems/:id", GetController().UpdateStem)
		v1.PUT("/v1/stems/:id", GetController().UpdateStem)
		v1.DELETE("/v1/stems/:id", GetController().DeleteStem)

		v1.GET("/v1/sticks", GetController().GetSticks)
		v1.GET("/v1/sticks/:id", GetController().GetStick)
		v1.POST("/v1/sticks", GetController().PostStick)
		v1.PATCH("/v1/sticks/:id", GetController().UpdateStick)
		v1.PUT("/v1/sticks/:id", GetController().UpdateStick)
		v1.DELETE("/v1/sticks/:id", GetController().DeleteStick)

		v1.GET("/v1/string_mutes", GetController().GetString_mutes)
		v1.GET("/v1/string_mutes/:id", GetController().GetString_mute)
		v1.POST("/v1/string_mutes", GetController().PostString_mute)
		v1.PATCH("/v1/string_mutes/:id", GetController().UpdateString_mute)
		v1.PUT("/v1/string_mutes/:id", GetController().UpdateString_mute)
		v1.DELETE("/v1/string_mutes/:id", GetController().DeleteString_mute)

		v1.GET("/v1/strong_accents", GetController().GetStrong_accents)
		v1.GET("/v1/strong_accents/:id", GetController().GetStrong_accent)
		v1.POST("/v1/strong_accents", GetController().PostStrong_accent)
		v1.PATCH("/v1/strong_accents/:id", GetController().UpdateStrong_accent)
		v1.PUT("/v1/strong_accents/:id", GetController().UpdateStrong_accent)
		v1.DELETE("/v1/strong_accents/:id", GetController().DeleteStrong_accent)

		v1.GET("/v1/supportss", GetController().GetSupportss)
		v1.GET("/v1/supportss/:id", GetController().GetSupports)
		v1.POST("/v1/supportss", GetController().PostSupports)
		v1.PATCH("/v1/supportss/:id", GetController().UpdateSupports)
		v1.PUT("/v1/supportss/:id", GetController().UpdateSupports)
		v1.DELETE("/v1/supportss/:id", GetController().DeleteSupports)

		v1.GET("/v1/swings", GetController().GetSwings)
		v1.GET("/v1/swings/:id", GetController().GetSwing)
		v1.POST("/v1/swings", GetController().PostSwing)
		v1.PATCH("/v1/swings/:id", GetController().UpdateSwing)
		v1.PUT("/v1/swings/:id", GetController().UpdateSwing)
		v1.DELETE("/v1/swings/:id", GetController().DeleteSwing)

		v1.GET("/v1/syncs", GetController().GetSyncs)
		v1.GET("/v1/syncs/:id", GetController().GetSync)
		v1.POST("/v1/syncs", GetController().PostSync)
		v1.PATCH("/v1/syncs/:id", GetController().UpdateSync)
		v1.PUT("/v1/syncs/:id", GetController().UpdateSync)
		v1.DELETE("/v1/syncs/:id", GetController().DeleteSync)

		v1.GET("/v1/system_dividerss", GetController().GetSystem_dividerss)
		v1.GET("/v1/system_dividerss/:id", GetController().GetSystem_dividers)
		v1.POST("/v1/system_dividerss", GetController().PostSystem_dividers)
		v1.PATCH("/v1/system_dividerss/:id", GetController().UpdateSystem_dividers)
		v1.PUT("/v1/system_dividerss/:id", GetController().UpdateSystem_dividers)
		v1.DELETE("/v1/system_dividerss/:id", GetController().DeleteSystem_dividers)

		v1.GET("/v1/system_layouts", GetController().GetSystem_layouts)
		v1.GET("/v1/system_layouts/:id", GetController().GetSystem_layout)
		v1.POST("/v1/system_layouts", GetController().PostSystem_layout)
		v1.PATCH("/v1/system_layouts/:id", GetController().UpdateSystem_layout)
		v1.PUT("/v1/system_layouts/:id", GetController().UpdateSystem_layout)
		v1.DELETE("/v1/system_layouts/:id", GetController().DeleteSystem_layout)

		v1.GET("/v1/system_marginss", GetController().GetSystem_marginss)
		v1.GET("/v1/system_marginss/:id", GetController().GetSystem_margins)
		v1.POST("/v1/system_marginss", GetController().PostSystem_margins)
		v1.PATCH("/v1/system_marginss/:id", GetController().UpdateSystem_margins)
		v1.PUT("/v1/system_marginss/:id", GetController().UpdateSystem_margins)
		v1.DELETE("/v1/system_marginss/:id", GetController().DeleteSystem_margins)

		v1.GET("/v1/taps", GetController().GetTaps)
		v1.GET("/v1/taps/:id", GetController().GetTap)
		v1.POST("/v1/taps", GetController().PostTap)
		v1.PATCH("/v1/taps/:id", GetController().UpdateTap)
		v1.PUT("/v1/taps/:id", GetController().UpdateTap)
		v1.DELETE("/v1/taps/:id", GetController().DeleteTap)

		v1.GET("/v1/technicals", GetController().GetTechnicals)
		v1.GET("/v1/technicals/:id", GetController().GetTechnical)
		v1.POST("/v1/technicals", GetController().PostTechnical)
		v1.PATCH("/v1/technicals/:id", GetController().UpdateTechnical)
		v1.PUT("/v1/technicals/:id", GetController().UpdateTechnical)
		v1.DELETE("/v1/technicals/:id", GetController().DeleteTechnical)

		v1.GET("/v1/text_element_datas", GetController().GetText_element_datas)
		v1.GET("/v1/text_element_datas/:id", GetController().GetText_element_data)
		v1.POST("/v1/text_element_datas", GetController().PostText_element_data)
		v1.PATCH("/v1/text_element_datas/:id", GetController().UpdateText_element_data)
		v1.PUT("/v1/text_element_datas/:id", GetController().UpdateText_element_data)
		v1.DELETE("/v1/text_element_datas/:id", GetController().DeleteText_element_data)

		v1.GET("/v1/ties", GetController().GetTies)
		v1.GET("/v1/ties/:id", GetController().GetTie)
		v1.POST("/v1/ties", GetController().PostTie)
		v1.PATCH("/v1/ties/:id", GetController().UpdateTie)
		v1.PUT("/v1/ties/:id", GetController().UpdateTie)
		v1.DELETE("/v1/ties/:id", GetController().DeleteTie)

		v1.GET("/v1/tieds", GetController().GetTieds)
		v1.GET("/v1/tieds/:id", GetController().GetTied)
		v1.POST("/v1/tieds", GetController().PostTied)
		v1.PATCH("/v1/tieds/:id", GetController().UpdateTied)
		v1.PUT("/v1/tieds/:id", GetController().UpdateTied)
		v1.DELETE("/v1/tieds/:id", GetController().DeleteTied)

		v1.GET("/v1/times", GetController().GetTimes)
		v1.GET("/v1/times/:id", GetController().GetTime)
		v1.POST("/v1/times", GetController().PostTime)
		v1.PATCH("/v1/times/:id", GetController().UpdateTime)
		v1.PUT("/v1/times/:id", GetController().UpdateTime)
		v1.DELETE("/v1/times/:id", GetController().DeleteTime)

		v1.GET("/v1/time_modifications", GetController().GetTime_modifications)
		v1.GET("/v1/time_modifications/:id", GetController().GetTime_modification)
		v1.POST("/v1/time_modifications", GetController().PostTime_modification)
		v1.PATCH("/v1/time_modifications/:id", GetController().UpdateTime_modification)
		v1.PUT("/v1/time_modifications/:id", GetController().UpdateTime_modification)
		v1.DELETE("/v1/time_modifications/:id", GetController().DeleteTime_modification)

		v1.GET("/v1/timpanis", GetController().GetTimpanis)
		v1.GET("/v1/timpanis/:id", GetController().GetTimpani)
		v1.POST("/v1/timpanis", GetController().PostTimpani)
		v1.PATCH("/v1/timpanis/:id", GetController().UpdateTimpani)
		v1.PUT("/v1/timpanis/:id", GetController().UpdateTimpani)
		v1.DELETE("/v1/timpanis/:id", GetController().DeleteTimpani)

		v1.GET("/v1/transposes", GetController().GetTransposes)
		v1.GET("/v1/transposes/:id", GetController().GetTranspose)
		v1.POST("/v1/transposes", GetController().PostTranspose)
		v1.PATCH("/v1/transposes/:id", GetController().UpdateTranspose)
		v1.PUT("/v1/transposes/:id", GetController().UpdateTranspose)
		v1.DELETE("/v1/transposes/:id", GetController().DeleteTranspose)

		v1.GET("/v1/tremolos", GetController().GetTremolos)
		v1.GET("/v1/tremolos/:id", GetController().GetTremolo)
		v1.POST("/v1/tremolos", GetController().PostTremolo)
		v1.PATCH("/v1/tremolos/:id", GetController().UpdateTremolo)
		v1.PUT("/v1/tremolos/:id", GetController().UpdateTremolo)
		v1.DELETE("/v1/tremolos/:id", GetController().DeleteTremolo)

		v1.GET("/v1/tuplets", GetController().GetTuplets)
		v1.GET("/v1/tuplets/:id", GetController().GetTuplet)
		v1.POST("/v1/tuplets", GetController().PostTuplet)
		v1.PATCH("/v1/tuplets/:id", GetController().UpdateTuplet)
		v1.PUT("/v1/tuplets/:id", GetController().UpdateTuplet)
		v1.DELETE("/v1/tuplets/:id", GetController().DeleteTuplet)

		v1.GET("/v1/tuplet_dots", GetController().GetTuplet_dots)
		v1.GET("/v1/tuplet_dots/:id", GetController().GetTuplet_dot)
		v1.POST("/v1/tuplet_dots", GetController().PostTuplet_dot)
		v1.PATCH("/v1/tuplet_dots/:id", GetController().UpdateTuplet_dot)
		v1.PUT("/v1/tuplet_dots/:id", GetController().UpdateTuplet_dot)
		v1.DELETE("/v1/tuplet_dots/:id", GetController().DeleteTuplet_dot)

		v1.GET("/v1/tuplet_numbers", GetController().GetTuplet_numbers)
		v1.GET("/v1/tuplet_numbers/:id", GetController().GetTuplet_number)
		v1.POST("/v1/tuplet_numbers", GetController().PostTuplet_number)
		v1.PATCH("/v1/tuplet_numbers/:id", GetController().UpdateTuplet_number)
		v1.PUT("/v1/tuplet_numbers/:id", GetController().UpdateTuplet_number)
		v1.DELETE("/v1/tuplet_numbers/:id", GetController().DeleteTuplet_number)

		v1.GET("/v1/tuplet_portions", GetController().GetTuplet_portions)
		v1.GET("/v1/tuplet_portions/:id", GetController().GetTuplet_portion)
		v1.POST("/v1/tuplet_portions", GetController().PostTuplet_portion)
		v1.PATCH("/v1/tuplet_portions/:id", GetController().UpdateTuplet_portion)
		v1.PUT("/v1/tuplet_portions/:id", GetController().UpdateTuplet_portion)
		v1.DELETE("/v1/tuplet_portions/:id", GetController().DeleteTuplet_portion)

		v1.GET("/v1/tuplet_types", GetController().GetTuplet_types)
		v1.GET("/v1/tuplet_types/:id", GetController().GetTuplet_type)
		v1.POST("/v1/tuplet_types", GetController().PostTuplet_type)
		v1.PATCH("/v1/tuplet_types/:id", GetController().UpdateTuplet_type)
		v1.PUT("/v1/tuplet_types/:id", GetController().UpdateTuplet_type)
		v1.DELETE("/v1/tuplet_types/:id", GetController().DeleteTuplet_type)

		v1.GET("/v1/typed_texts", GetController().GetTyped_texts)
		v1.GET("/v1/typed_texts/:id", GetController().GetTyped_text)
		v1.POST("/v1/typed_texts", GetController().PostTyped_text)
		v1.PATCH("/v1/typed_texts/:id", GetController().UpdateTyped_text)
		v1.PUT("/v1/typed_texts/:id", GetController().UpdateTyped_text)
		v1.DELETE("/v1/typed_texts/:id", GetController().DeleteTyped_text)

		v1.GET("/v1/unpitcheds", GetController().GetUnpitcheds)
		v1.GET("/v1/unpitcheds/:id", GetController().GetUnpitched)
		v1.POST("/v1/unpitcheds", GetController().PostUnpitched)
		v1.PATCH("/v1/unpitcheds/:id", GetController().UpdateUnpitched)
		v1.PUT("/v1/unpitcheds/:id", GetController().UpdateUnpitched)
		v1.DELETE("/v1/unpitcheds/:id", GetController().DeleteUnpitched)

		v1.GET("/v1/virtual_instruments", GetController().GetVirtual_instruments)
		v1.GET("/v1/virtual_instruments/:id", GetController().GetVirtual_instrument)
		v1.POST("/v1/virtual_instruments", GetController().PostVirtual_instrument)
		v1.PATCH("/v1/virtual_instruments/:id", GetController().UpdateVirtual_instrument)
		v1.PUT("/v1/virtual_instruments/:id", GetController().UpdateVirtual_instrument)
		v1.DELETE("/v1/virtual_instruments/:id", GetController().DeleteVirtual_instrument)

		v1.GET("/v1/waits", GetController().GetWaits)
		v1.GET("/v1/waits/:id", GetController().GetWait)
		v1.POST("/v1/waits", GetController().PostWait)
		v1.PATCH("/v1/waits/:id", GetController().UpdateWait)
		v1.PUT("/v1/waits/:id", GetController().UpdateWait)
		v1.DELETE("/v1/waits/:id", GetController().DeleteWait)

		v1.GET("/v1/wavy_lines", GetController().GetWavy_lines)
		v1.GET("/v1/wavy_lines/:id", GetController().GetWavy_line)
		v1.POST("/v1/wavy_lines", GetController().PostWavy_line)
		v1.PATCH("/v1/wavy_lines/:id", GetController().UpdateWavy_line)
		v1.PUT("/v1/wavy_lines/:id", GetController().UpdateWavy_line)
		v1.DELETE("/v1/wavy_lines/:id", GetController().DeleteWavy_line)

		v1.GET("/v1/wedges", GetController().GetWedges)
		v1.GET("/v1/wedges/:id", GetController().GetWedge)
		v1.POST("/v1/wedges", GetController().PostWedge)
		v1.PATCH("/v1/wedges/:id", GetController().UpdateWedge)
		v1.PUT("/v1/wedges/:id", GetController().UpdateWedge)
		v1.DELETE("/v1/wedges/:id", GetController().DeleteWedge)

		v1.GET("/v1/woods", GetController().GetWoods)
		v1.GET("/v1/woods/:id", GetController().GetWood)
		v1.POST("/v1/woods", GetController().PostWood)
		v1.PATCH("/v1/woods/:id", GetController().UpdateWood)
		v1.PUT("/v1/woods/:id", GetController().UpdateWood)
		v1.DELETE("/v1/woods/:id", GetController().DeleteWood)

		v1.GET("/v1/works", GetController().GetWorks)
		v1.GET("/v1/works/:id", GetController().GetWork)
		v1.POST("/v1/works", GetController().PostWork)
		v1.PATCH("/v1/works/:id", GetController().UpdateWork)
		v1.PUT("/v1/works/:id", GetController().UpdateWork)
		v1.DELETE("/v1/works/:id", GetController().DeleteWork)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/commitfrombacknb", GetController().onWebSocketRequestForCommitFromBackNb)
		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)

		v1.GET("/v1/stacks", GetController().stacks)
	}
}

func (controller *Controller) stacks(c *gin.Context) {

	var res []string

	for k, _ := range controller.Map_BackRepos {
		res = append(res, k)
	}

	c.JSON(http.StatusOK, res)
}

// onWebSocketRequestForCommitFromBackNb is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForCommitFromBackNb(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gongmusicxml/go, onWebSocketRequestForCommitFromBackNb")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {

		// Send elapsed time as a string over the WebSocket connection
		err = wsConnection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", nbCommitBackRepo)))
		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gongmusicxml/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

	err = wsConnection.WriteJSON(backRepoData)
	// log.Println("Stack github.com/fullstack-lang/gongmusicxml/go, onWebSocketRequestForBackRepoContent, first sent back repo of", stackPath)
	if err != nil {
		log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	}

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
		_ = nbCommitBackRepo

		backRepoData := new(orm.BackRepoData)
		orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

		// Send backRepo data
		err = wsConnection.WriteJSON(backRepoData)

		// log.Println("Stack github.com/fullstack-lang/gongmusicxml/go, onWebSocketRequestForBackRepoContent, sent back repo of", stackPath)

		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
