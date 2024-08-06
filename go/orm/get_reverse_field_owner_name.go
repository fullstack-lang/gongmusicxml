// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongmusicxml/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Bookmark:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Foo:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Lyric:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Lyric_font:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Lyric_language:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Miscellaneous_field:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Bookmark:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Foo:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Lyric:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Lyric_font:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Lyric_language:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Miscellaneous_field:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
