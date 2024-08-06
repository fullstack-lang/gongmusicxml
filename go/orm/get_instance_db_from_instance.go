// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongmusicxml/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | BookmarkDB | FooDB | LinkDB | LyricDB | Lyric_fontDB | Lyric_languageDB | Miscellaneous_fieldDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Bookmark:
		bookmarkInstance := any(concreteInstance).(*models.Bookmark)
		ret2 := backRepo.BackRepoBookmark.GetBookmarkDBFromBookmarkPtr(bookmarkInstance)
		ret = any(ret2).(*T2)
	case *models.Foo:
		fooInstance := any(concreteInstance).(*models.Foo)
		ret2 := backRepo.BackRepoFoo.GetFooDBFromFooPtr(fooInstance)
		ret = any(ret2).(*T2)
	case *models.Link:
		linkInstance := any(concreteInstance).(*models.Link)
		ret2 := backRepo.BackRepoLink.GetLinkDBFromLinkPtr(linkInstance)
		ret = any(ret2).(*T2)
	case *models.Lyric:
		lyricInstance := any(concreteInstance).(*models.Lyric)
		ret2 := backRepo.BackRepoLyric.GetLyricDBFromLyricPtr(lyricInstance)
		ret = any(ret2).(*T2)
	case *models.Lyric_font:
		lyric_fontInstance := any(concreteInstance).(*models.Lyric_font)
		ret2 := backRepo.BackRepoLyric_font.GetLyric_fontDBFromLyric_fontPtr(lyric_fontInstance)
		ret = any(ret2).(*T2)
	case *models.Lyric_language:
		lyric_languageInstance := any(concreteInstance).(*models.Lyric_language)
		ret2 := backRepo.BackRepoLyric_language.GetLyric_languageDBFromLyric_languagePtr(lyric_languageInstance)
		ret = any(ret2).(*T2)
	case *models.Miscellaneous_field:
		miscellaneous_fieldInstance := any(concreteInstance).(*models.Miscellaneous_field)
		ret2 := backRepo.BackRepoMiscellaneous_field.GetMiscellaneous_fieldDBFromMiscellaneous_fieldPtr(miscellaneous_fieldInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Bookmark:
		tmp := GetInstanceDBFromInstance[models.Bookmark, BookmarkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Foo:
		tmp := GetInstanceDBFromInstance[models.Foo, FooDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric:
		tmp := GetInstanceDBFromInstance[models.Lyric, LyricDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric_font:
		tmp := GetInstanceDBFromInstance[models.Lyric_font, Lyric_fontDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric_language:
		tmp := GetInstanceDBFromInstance[models.Lyric_language, Lyric_languageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Miscellaneous_field:
		tmp := GetInstanceDBFromInstance[models.Miscellaneous_field, Miscellaneous_fieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Bookmark:
		tmp := GetInstanceDBFromInstance[models.Bookmark, BookmarkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Foo:
		tmp := GetInstanceDBFromInstance[models.Foo, FooDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric:
		tmp := GetInstanceDBFromInstance[models.Lyric, LyricDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric_font:
		tmp := GetInstanceDBFromInstance[models.Lyric_font, Lyric_fontDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric_language:
		tmp := GetInstanceDBFromInstance[models.Lyric_language, Lyric_languageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Miscellaneous_field:
		tmp := GetInstanceDBFromInstance[models.Miscellaneous_field, Miscellaneous_fieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
